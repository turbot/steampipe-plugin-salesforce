package salesforce

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/turbot/go-kit/helpers"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

const pluginName = "steampipe-plugin-salesforce"

type contextKey string

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             pluginName,
		DefaultTransform: transform.FromCamel().NullIfZero(),
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		SchemaMode:   plugin.SchemaModeDynamic,
		TableMapFunc: pluginTableDefinitions,
	}

	return p
}

func pluginTableDefinitions(ctx context.Context, p *plugin.Plugin) (map[string]*plugin.Table, error) {

	// Some fixed tables added to plugin
	salesforceTables := []string{
		"Account",
		"AccountContactRole",
		"Asset",
		"Campaign",
		"Case",
		"Contact",
		"Contract",
		"Lead",
		"Opportunity",
		"OpportunityContactRole",
		"Order",
		"Pricebook2",
		"PricebookEntry",
		"Product2",
		"User",
	}

	// Initialize tables
	tables := map[string]*plugin.Table{
		"salesforce_account":                  SalesforceAccount(ctx),
		"salesforce_account_contact_role":     SalesforceAccountContactRole(ctx),
		"salesforce_contract":                 SalesforceContract(ctx),
		"salesforce_lead":                     SalesforceLead(ctx),
		"salesforce_opportunity":              SalesforceOpportunity(ctx),
		"salesforce_opportunity_contact_role": SalesforceOpportunityContactRole(ctx),
		"salesforce_order":                    SalesforceOrder(ctx),
		"salesforce_product":                  SalesforceProduct(ctx),
		"salesforce_user":                     SalesforceUser(ctx),
	}

	config := GetConfig(p.Connection)
	if config.Tables != nil && len(*config.Tables) > 0 {
		for _, tableName := range *config.Tables {
			if !helpers.StringSliceContains(salesforceTables, tableName) {
				salesforceTables = append(salesforceTables, tableName)
			}
		}
	}

	// If unable to connect to salesforce instance, log warning and abort dynamic table creation
	client, err := connectRaw(ctx, p.ConnectionManager, p.Connection)
	if err != nil {
		plugin.Logger(ctx).Error("salesforce.pluginTableDefinitions", "connection_error: unable to generate dynamic tables because of invalid steampipe salesforce configuration", err)
		return tables, nil
	}
	if client == nil {
		plugin.Logger(ctx).Error("salesforce.pluginTableDefinitions", "connection_error: unable to generate dynamic tables because of invalid steampipe salesforce configuration", err)
		return tables, nil
	}
	var re = regexp.MustCompile(`\d+`)
	var substitution = ``

	for _, table := range salesforceTables {
		ctx = context.WithValue(ctx, contextKey("SalesforceTableName"), table)
		tableName := "salesforce_" + strcase.ToSnake(re.ReplaceAllString(table, substitution))
		ctx = context.WithValue(ctx, contextKey("PluginTableName"), tableName)
		plugin.Logger(ctx).Debug("pluginTableDefinitions Table Names", "SALESFORCE_OBJECT_NAME", table, "STEAMPIPE_TABLE_NAME", tableName)
		delete(tables, tableName)
		tables[tableName] = generateDynamicTables(ctx, p)
	}
	return tables, nil
}

func generateDynamicTables(ctx context.Context, p *plugin.Plugin) *plugin.Table {

	client, err := connectRaw(ctx, p.ConnectionManager, p.Connection)
	if err != nil {
		plugin.Logger(ctx).Error("salesforce.generateDynamicTables", "connection_error", err)
		return nil
	}

	// Get the query for the metric (required)
	salesforceTableName := ctx.Value(contextKey("SalesforceTableName")).(string)
	tableName := ctx.Value(contextKey("PluginTableName")).(string)

	sObjectMeta := client.SObject(salesforceTableName).Describe()
	if sObjectMeta == nil {
		plugin.Logger(ctx).Error("salesforce.generateDynamicTables", fmt.Sprintf("Table %s not present in salesforce", salesforceTableName))
		return nil
	}

	// Top columns
	cols := []*plugin.Column{}

	data := *sObjectMeta
	data1, err := json.Marshal(data["fields"])
	if err != nil {
		plugin.Logger(ctx).Error("[simpleforce] json marshal error %v", err)
	}

	metadata := []map[string]interface{}{}
	// var queryColumns []string
	err = json.Unmarshal(data1, &metadata)
	if err != nil {
		plugin.Logger(ctx).Error("[simpleforce] json unmarshal error %v", err)
	}

	for _, fields := range metadata {
		if fields["name"] == nil {
			continue
		}
		fieldName := fields["name"].(string)
		compoundFieldName := fields["compoundFieldName"]
		if compoundFieldName != nil && compoundFieldName.(string) != fieldName {
			continue
		}

		// queryColumns = append(queryColumns, fieldName)
		if fields["soapType"] == nil {
			continue
		}
		soapType := strings.Split((fields["soapType"]).(string), ":")
		fieldType := soapType[len(soapType)-1]

		// Coloumn dynamic generation
		columnFieldName := strcase.ToSnake(fieldName)
		column := plugin.Column{
			Name:        columnFieldName,
			Description: fmt.Sprintf("The %s.", fields["label"].(string)),
			Transform:   transform.FromP(getFieldFromSObjectMap, fieldName),
		}

		if columns, ok := columnDescriptions[salesforceTableName]; ok {
			if columnFieldName, ok := columns.Columns[columnFieldName]; ok {
				column.Description = columnFieldName
			}
		}

		// Set column type based on the `soapType` from salesforce schema
		switch fieldType {
		case "string", "ID":
			column.Type = proto.ColumnType_STRING
		case "date", "dateTime":
			column.Type = proto.ColumnType_TIMESTAMP
		case "boolean":
			column.Type = proto.ColumnType_BOOL
		case "double":
			column.Type = proto.ColumnType_DOUBLE
		case "int":
			column.Type = proto.ColumnType_INT
		default:
			column.Type = proto.ColumnType_JSON
		}
		cols = append(cols, &column)

	}

	// query := fmt.Sprintf("SELECT %s FROM %s", strings.Join(queryColumns, ", "), salesforceTableName)

	Table := plugin.Table{
		Name:        tableName,
		Description: fmt.Sprintf("Salesforce %s.", data["label"]),
		List: &plugin.ListConfig{
			// KeyColumns: keyColumns,
			Hydrate: listSalesforceObjectsByTable(salesforceTableName),
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getSalesforceObjectbyID(salesforceTableName),
		},
		Columns: cols,
	}

	if tableDescription, ok := columnDescriptions[salesforceTableName]; ok {
		Table.Description = tableDescription.Description
	}

	return &Table
}

type TableDescription struct {
	Description string
	Columns     map[string]string
}

var columnDescriptions = map[string]TableDescription{
	"Account": {
		Description: "Represents an individual account, which is an organization or person involved with business (such as customers, competitors, and partners).",
		Columns: map[string]string{
			"id":                  "Unique identifier of the account in Salesforce.",
			"name":                "Name of the account.",
			"annual_revenue":      "Estimated annual revenue of the account.",
			"industry":            "Primary business of account.",
			"owner_id":            "The ID of the user who currently owns this account. Default value is the user logged in to the API to perform the create.",
			"type":                "Type of account, for example, Customer, Competitor, or Partner.",
			"account_source":      "The source of the account record. For example, Advertisement, Data.com, or Trade Show.",
			"clean_status":        "Indicates the record's clean status as compared with Data.com. Values are: Matched, Different,Acknowledged,NotFound,Inactive,Pending, SelectMatch, or Skipped.",
			"created_by_id":       "The id of the user who created the account.",
			"created_date":        "The creation date and time of the account.",
			"description":         "Text description of the account.",
			"is_deleted":          "Indicates whether the object has been moved to the Recycle Bin (true) or not (false).",
			"last_modified_by_id": "The id of the user who last changed the contact fields, including modification date and time.",
			"last_modified_date":  "The date and time of last modification to account.",
			"number_of_employees": "Number of employees working at the company represented by this account.",
			"ownership":           "Ownership type for the account, for example Private, Public, or Subsidiary.",
			"phone":               "The contact's primary phone number.",
			"rating":              "The account's prospect rating, for example Hot, Warm, or Cold.",
			"sic":                 "Standard Industrial Classification code of the company's main business categorization, for example, 57340 for Electronics.",
			"ticker_symbol":       "The stock market symbol for this account.",
			"tradestyle":          "A name, different from its legal name, that an org may use for conducting business. Similar to “Doing business as” or \"DBA\".",
			"website":             "The website of this account, for example, www.acme.com.",
			"billing_address":     "The billing adress of the account.",
			"shipping_address":    "The shipping adress of the account.",
		},
	},
	"AccountContactRole": {
		Description: "Represents the role that a Contact plays on an Account.",
		Columns: map[string]string{
			"id":                  "Unique identifier of the account contact role in Salesforce.",
			"account_id":          "ID of the Account.",
			"contact_id":          "ID of the Contact associated with this account.",
			"is_primary":          "Specifies whether the Contact plays the primary role on the Account (true) or not (false). Note that each account has only one primary contact role. Label is Primary. Default value is false.",
			"role":                "Name of the role played by the Contact on this Account, such as Decision Maker, Approver, Buyer, and so on.",
			"created_by_id":       "Id of the user who created contact role record.",
			"created_date":        "Date and time of the creation of the contact role record.",
			"last_modified_by_id": "Id of the user who most recently changed the contact role record.",
			"last_modified_date":  "Date of most recent change in the contact role record.",
		},
	},
	"Contract": {
		Description: "Represents a contract (a business agreement) associated with an Account.",
		Columns: map[string]string{
			"id":                      "Unique identifier of the contract in Salesforce.",
			"account_id":              "ID of the Account associated with this contract.",
			"contract_number":         "Number of the contract.",
			"contract_term":           "Number of months that the contract is valid.",
			"end_date":                "Calculated end date of the contract. This value is calculated by adding the ContractTerm to the start_date.",
			"owner_id":                "ID of the user who owns the contract.",
			"start_date":              "Start date for this contract.",
			"status":                  "The picklist of values that indicate order status. A contract can be Draft, InApproval, or Activated.",
			"activated_by_id":         "ID of the User who activated this contract.",
			"activated_date":          "Date and time when this contract was activated.",
			"billing_address":         "Billing address of the account.",
			"company_signed_date":     "Date on which the contract was signed by organization.",
			"company_signed_id":       "ID of the User who signed the contract.",
			"created_by_id":           "Id of the user who created the contract record.",
			"created_date":            "Date and time when contract record was created.",
			"customer_signed_date":    "Date on which the customer signed the contract.",
			"customer_signed_id":      "ID of the Contact who signed this contract.",
			"customer_signed_title":   "Title of the contact who signed the contract.",
			"description":             "Statement describing the contract.",
			"is_deleted":              "Indicates whether the object has been moved to the Recycle Bin (true) or not (false).",
			"last_activity_date":      "Value is one of the following, whichever is the most recent. a) Due date of the most recent event logged against the record. b) Due date of the most recently closed task associated with the record.",
			"last_approved_date":      "Last date the contract was approved.",
			"last_modified_by_id":     "The id of user who most recently changed the contract record.",
			"last_modified_date":      "The date and time of the last change to contract record.",
			"last_referenced_date":    "The timestamp when the current user last accessed this record, a record related to this record, or a list view.",
			"last_viewed_date":        "The timestamp when the current user last viewed this record or list view. If this value is null, the user might have only accessed this record or list view (last_referenced_date) but not viewed it.",
			"owner_expiration_notice": "Number of days ahead of the contract end date (15, 30, 45, 60, 90, and 120). Used to notify the owner in advance that the contract is ending.",
			"pricebook_2_id":          "ID of the pricebook, if any, associated with this contract.",
			"special_terms":           "Special terms that apply to the contract.",
			"system_modstamp":         "The date and time when contract was last modified by a user or by an automated process.",
		},
	},
	"Lead": {
		Description: "Represents a prospect or lead.",
		Columns: map[string]string{
			"id":                  "Unique identifier of the lead in Salesforce.",
			"email":               "The lead's email address.",
			"is_converted":        "Indicates whether the lead has been converted (true) or not (false).",
			"name":                "Name of the lead, as displayed on lead detail page.",
			"phone":               "Lead's primary phone number.",
			"status":              "Status of the lead, for example, Open, Contacted, Qualified.",
			"address":             "Street address for the lead.",
			"annual_revenue":      "Annual revenue for the lead's company.",
			"company":             "The lead's company.",
			"converted_date":      "Date on which this lead was converted.",
			"created_by_id":       "Id of the user who created the lead.",
			"created_date":        "Creation date and time of the lead.",
			"industry":            "Primary business of lead's company.",
			"last_modified_by_id": "Id of the user who last changed the lead record.",
			"last_modified_date":  "Date and time of the last changes to lead record.",
			"lead_source":         "Source of lead, for example, Advertisement, Partner, or Web.",
			"number_of_employees": "Number of employees at the lead's company.",
			"owner_id":            "Id of the assigned owner of the lead.",
			"rating":              "Indicates value or prospect of the lead, for example, Hot, Warm, Cold.",
			"website":             "URL of the lead's company's website.",
		},
	},
	"OpportunityContactRole": {
		Description: "Represents the role that a Contact plays on an Opportunity.",
		Columns: map[string]string{
			"id":                  "Unique identifier of the opportunity contact role in Salesforce.",
			"contact_id":          "ID of an associated Contact.",
			"is_primary":          "Indicates whether the associated Contact plays the primary role on the Opportunity (true) or not (false).",
			"opportunity_id":      "Required. ID of an associated Opportunity.",
			"role":                "Name of the role played by the associated Contact on the Opportunity, such as Business User or Decision Maker.",
			"created_by_id":       "Id of the user who created contact role record.",
			"created_date":        "Date and time of the creation of the contact role record.",
			"last_modified_by_id": "Id of the user who most recently changed the contact role record.",
			"last_modified_date":  "Date of most recent change in the contact role record.",
		},
	},
	"Opportunity": {
		Description: "Represents an opportunity, which is a sale or pending deal.",
		Columns: map[string]string{
			"id":                         "Unique identifier of the opportunity in Salesforce.",
			"account_id":                 "ID of the account associated with this opportunity.",
			"amount":                     "Estimated total sale amount. For opportunities with products, the amount is the sum of the related products.",
			"name":                       "A name for this opportunity.",
			"owner_id":                   "ID of the User who has been assigned to work this opportunity.",
			"campaign_id":                "ID of a related Campaign. This field is defined only for those organizations that have the campaign feature Campaigns enabled.",
			"close_date":                 "Date when the opportunity is expected to close.",
			"created_by_id":              "Id of the user who created the opportunity.",
			"created_date":               "The creation date and time of the opportunity.",
			"description":                "Description of the opportunity.",
			"expected_revenue":           "Calculated revenue based on the Amount and Probability fields.",
			"fiscal_quarter":             "Represents the fiscal quarter. Valid values are 1, 2, 3, or 4.",
			"fiscal_year":                "Represents the fiscal year, for example, 2006.",
			"forecast_category":          "Forecast category name displayed in reports, opportunity detail and edit pages, opportunity searches, and opportunity list views.",
			"forecast_category_name":     "Name of the forecast category.",
			"has_open_activity":          "Indicates whether an opportunity has an open event or task (true) or not (false).",
			"has_opportunity_line_item":  "Indicates whether the opportunity has associated line items. A value of true means that Opportunity line items have been created for the opportunity.",
			"has_overdue_task":           "Indicates whether an opportunity has an overdue task (true) or not (false).",
			"is_closed":                  "Indicates that the opportunity is closed.",
			"is_deleted":                 "Indicates that the opportunity is deleted.",
			"is_private":                 "Indicates that the opportunity is private.",
			"is_won":                     "Indicates that the quote or proposal has been signed or electronically accepted.",
			"last_activity_date":         "Value is one of the following, whichever is the most recent of a) Due date of the most recent event logged against the record or b) Due date of the most recently closed task associated with the record.",
			"last_modified_by_id":        "The id of the user who last modified the oppurtinity record.",
			"last_modified_date":         "The data and time of the last modification of the oppurtinity record.",
			"lead_source":                "Source of this opportunity, such as Advertisement or Trade Show.",
			"next_step":                  "Description of next task in closing opportunity.",
			"pricebook_2_id":             "ID of a related Pricebook2 object. The Pricebook2Id field indicates which Pricebook2 applies to this opportunity. The Pricebook2Id field is defined only for those organizations that have products enabled as a feature.",
			"probability":                "Percentage of estimated confidence in closing the opportunity.",
			"stage_name":                 "Current stage of opportunity.",
			"system_modstamp":            "The date and time when opportunity was last modified by a user or by an automated process.",
			"total_opportunity_quantity": "Number of items included in this opportunity. Used in quantity-based forecasting.",
			"type":                       "Type of opportunity, such as Existing Business or New Business.",
		},
	},
	"Order": {
		Description: "Represents an order associated with a contract or an account.",
		Columns: map[string]string{
			"id":                        "Unique identifier of the order in Salesforce.",
			"name":                      "Title for the order that distinguishes it from other orders.",
			"account_id":                "ID of the Account associated with this order.",
			"order_number":              "Order number assigned to this order.",
			"owner_id":                  " ID of the User or queue that owns this order.",
			"status":                    "The Status field specifies the current state of an order. Status strings represent its current state (Draft or Activated).",
			"total_amount":              "Total amount of the order.",
			"type":                      "Type of order.",
			"activated_by_id":           "ID of the User who activated this order.",
			"activated_date":            "Date and time when the order was activated.",
			"bill_to_contact_id":        "ID of the contact that the order is billed to.",
			"billing_address":           "The billing adress for the order.",
			"company_authorized_by_id":  "ID of the user who authorized the account associated with the order.",
			"company_authorized_date":   "The date on which your organization authorized the order.",
			"contract_id":               "ID of the contract associated with this order. Can only be updated when the order's StatusCode value is Draft.",
			"created_by_id":             "Id of the user who created the order record.",
			"created_date":              "Creation date and time of the order record.",
			"customer_authorized_by_id": "ID of the contact who authorized the order.",
			"customer_authorized_date":  "Date on which the contact authorized the order.",
			"description":               "Description of the order.",
			"effective_date":            "Date at which the order becomes effective.",
			"end_date":                  "Date at which the order ends.",
			"is_deleted":                "Indicates that the order is deleted.",
			"is_reduction_order":        "Determines whether an order is a reduction order.",
			"last_modified_by_id":       "Id of the user who most recently changed the order record.",
			"last_modified_date":        "Date of most recent change in the order record.",
			"last_referenced_date":      "The timestamp when the current user last accessed this record, a record related to this record, or a list view.",
			"last_viewed_date":          "The timestamp when the current user last viewed this record or list view. If this value is null, the user might have only accessed this record or list view (LastReferencedDate) but not viewed it.",
			"order_reference_number":    "Reference number assigned to the order.",
			"original_order_id":         "Optional. ID of the original order that a reduction order is reducing, if the reduction order is reducing a single order.",
			"po_date":                   "Date of the purchase order.",
			"po_number":                 "Number identifying the purchase order.",
			"pricebook_2_id":            "ID of the price book associated with this order.",
			"ship_to_contact_id":        "ID of the contact that the order is shipped to.",
			"status_code":               "Status code of the stage that the order has reached in the order business process.",
			"system_modstamp":           "The date and time when order record was last modified by a user or by an automated process.",
			"shipping_address":          "The shipping adress for the order.",
		},
	},
	"Product2": {
		Description: "Represents a product that organization sells.",
		Columns: map[string]string{
			"id":                       "Unique identifier of the product in Salesforce.",
			"name":                     "The product's name.",
			"family":                   "Name of the product family associated with this record.",
			"product_code":             "The internal code or product number that you use to identify the product.",
			"is_active":                "Indicates that the product is ready for use in a price book, opportunity, or quote, and whether you can see the product in views.",
			"created_by_id":            "The id of the user who created the product record, with the date and time of creation.",
			"created_date":             "Date and time of the creation of the product role record.",
			"description":              "Description of the product.",
			"display_url":              "URL leading to a specific version of a record in the linked external data source.",
			"external_data_source_id":  "The id of the related external data source.",
			"external_id":              "The unique identifier of a record in the linked external data source.",
			"is_archived":              "Describes whether the product is archived. The default value is false.",
			"is_deleted":               "Indicates whether the object has been moved to the Recycle Bin (true) or not (false).",
			"last_modified_by_id":      "Id of the user who most recently changed the product record.",
			"last_modified_date":       "Date of most recent change in the product record.",
			"last_referenced_date":     "The timestamp when the current user last viewed product record.",
			"last_viewed_date":         "The timestamp when the current user last viewed this record. If this value is null, this record might only have been referenced (last_referenced_date) and not viewed by the current user.",
			"quantity_unit_of_measure": "Unit of the product—for example, kilograms, liters, or cases.",
			"stock_keeping_unit":       "The product's SKU, which can be used with or in place of the Product Code field.",
			"system_modstamp":          "The date and time when order record was last modified by a user or by an automated process.",
		},
	},
	"User": {
		Description: "Represents a user in organization.",
		Columns: map[string]string{
			"id":                  "Unique identifier of the user in Salesforce.",
			"alias":               "The user's alias. For example, jsmith.",
			"username":            "Login name of the user.",
			"name":                "Display name of the user.",
			"email":               "The user's email address.",
			"is_active":           "Indicates whether the user has access to log in (true) or not (false).",
			"account_id":          "ID of the Account associated with a Customer Portal user. This field is null for Salesforce users.",
			"created_by_id":       "Id of the user who created the user including creation date and time.",
			"department":          "The company department associated with the user.",
			"employee_number":     "The user's employee number.",
			"forecast_enabled":    "Indicates whether the user is enabled as a forecast manager (true) or not (false).",
			"last_login_date":     "The date and time when the user last successfully logged in. This value is updated if 60 seconds elapses since the user's last login.",
			"last_modified_by_id": "Id of the user who last changed the user fields, including modification date and time.",
			"profile_id":          "ID of the user's Profile.",
			"state":               "The state associated with the User.",
			"user_type":           "The category of user license. Can be one of Standard, PowerPartner, CSPLitePortal, CustomerSuccess, PowerCustomerSuccess, CsnOnly, and Guest.",
		},
	},
}
