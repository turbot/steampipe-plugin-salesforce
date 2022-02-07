connection "salesforce" {
  plugin = "salesforce"

  # Salesforce instance URL, e.g., "https://na01.salesforce.com/"
  # url = "YOUR_SALESFORCE_URL"

  # Salesforce account name
  # user = "YOUR_SALESFORCE_USERNAME"

  # Salesforce account password
  # password = "YOUR_SALESFORCE_PASSWORD"

  # Security token, could be omitted if Trusted IP is configured
  # token = "YOUR_SALESFORCE_TOKEN"

  # Salesforce client ID from the connected app
  # client_id = "YOUR_SALESFORCE_CLIENT_ID"

  # List of Salesforce object names to generate additional tables for
  # This argument only accepts exact Salesforce standard and custom object names, e.g., AccountBrand, OpportunityStage, CustomApp__c
  # For a full list of standard object names, please see https://developer.salesforce.com/docs/atlas.en-us.api.meta/api/sforce_api_objects_list.htm)
  # All custom object names should end in "__c", following Salesforce object naming
  # tables = ["AccountBrand", "OpportunityStage", "CustomApp__c"]

  # Salesforce API version to connect to.
  # api_version = "43.0"
}
