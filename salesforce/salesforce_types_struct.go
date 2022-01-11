package salesforce

// TODO: Handle time fields
type Account struct {
	AccountSource     string
	AnnualRevenue     float64
	BillingAddress    map[string]string
	CleanStatus       string
	CreatedById       string
	CreatedDate       string
	Description       string
	Id                string
	Industry          string
	IsDeleted         bool
	LastModifiedById  string
	LastModifiedDate  string
	Name              string
	NumberOfEmployees int64
	OwnerId           string
	Ownership         string
	Phone             string
	Rating            string
	ShippingAddress   map[string]string
	Sic               string
	TickerSymbol      string
	Tradestyle        string
	Type              string
	Website           string
}

type User struct {
	AboutMe                string
	AccountId              string
	Alias                  string
	CompanyName            string
	CreatedById            string
	CreatedDate            string
	Department             string
	Email                  string
	EmployeeNumber         string
	FirstName              string
	ForecastEnabled        bool
	Id                     string
	IsActive               bool
	IsPortalEnabled        bool
	IsPartner              bool
	LastLoginDate          string
	LastModifiedById       string
	LastModifiedDate       string
	LastName               string
	LastPasswordChangeDate string
	Name                   string
	ProfileId              string
	State                  string
	UserType               string
	Username               string
}

type Contract struct {
	AccountId             string
	ActivatedById         string
	ActivatedDate         string
	BillingAddress        interface{}
	CompanySignedDate     string
	CompanySignedId       string
	ContractNumber        string
	ContractTerm          int64
	CreatedById           string
	CreatedDate           string
	CustomerSignedDate    string
	CustomerSignedId      string
	CustomerSignedTitle   string
	Description           string
	EndDate               string
	Id                    string
	IsDeleted             bool
	LastActivityDate      string
	LastApprovedDate      string
	LastModifiedById      string
	LastModifiedDate      string
	LastReferencedDate    string
	LastViewedDate        string
	OwnerExpirationNotice string
	OwnerId               string
	Pricebook2Id          string
	SpecialTerms          string
	StartDate             string
	Status                string
	SystemModstamp        string
}

type AccountContactRole struct {
	Id               string
	ContactId        string
	AccountId        string
	CreatedById      string
	CreatedDate      string
	Role             string
	LastModifiedById string
	LastModifiedDate string
	IsPrimary        bool
}

type Opportunity struct {
	AccountId                string
	Amount                   float64
	CampaignId               string
	CloseDate                string
	CreatedById              string
	CreatedDate              string
	Description              string
	ExpectedRevenue          float64
	FiscalQuarter            int64
	FiscalYear               int64
	ForecastCategory         string
	ForecastCategoryName     string
	HasOpenActivity          bool
	HasOpportunityLineItem   bool
	HasOverdueTask           bool
	Id                       string
	IsClosed                 bool
	IsDeleted                bool
	IsPrivate                bool
	IsWon                    bool
	LastActivityDate         string
	LastModifiedById         string
	LastModifiedDate         string
	LeadSource               string
	Name                     string
	NextStep                 string
	OwnerId                  string
	Pricebook2Id             string
	Probability              float64
	StageName                string
	SystemModstamp           string
	TotalOpportunityQuantity string
	Type                     string
}

type OpportunityContactRole struct {
	Id               string `json:"Id"`
	ContactId        string `json:"ContactId"`
	OpportunityId    string `json:"OpportunityId"`
	CreatedById      string `json:"CreatedById"`
	CreatedDate      string `json:"CreatedDate"`
	Role             string `json:"Role"`
	LastModifiedById string `json:"LastModifiedById"`
	LastModifiedDate string `json:"LastModifiedDate"`
	IsPrimary        bool   `json:"IsPrimary"`
}

type Order struct {
	AccountId              string
	ActivatedById          string
	ActivatedDate          string
	BillToContactId        string
	CompanyAuthorizedById  string
	CompanyAuthorizedDate  string
	ContractId             string
	CreatedById            string
	CreatedDate            string
	CustomerAuthorizedById string
	CustomerAuthorizedDate string
	Description            string
	EffectiveDate          string
	EndDate                string
	Id                     string
	IsDeleted              bool
	IsReductionOrder       bool
	LastModifiedById       string
	LastModifiedDate       string
	LastReferencedDate     string
	LastViewedDate         string
	Name                   string
	OrderNumber            string
	OrderReferenceNumber   string
	OriginalOrderId        string
	OwnerId                string
	PoDate                 string
	PoNumber               string
	Pricebook2Id           string
	ShipToContactId        string
	Status                 string
	StatusCode             string
	SystemModstamp         string
	TotalAmount            float64
	Type                   string
	ShippingAddress        interface{}
	BillingAddress         interface{}
}

type Lead struct {
	Address           interface{}
	AnnualRevenue     float64
	Company           string
	ConvertedDate     string
	CreatedById       string
	CreatedDate       string
	Email             string
	Id                string
	Industry          string
	IsConverted       bool
	LastModifiedById  string
	LastModifiedDate  string
	LeadSource        string
	Name              string
	NumberOfEmployees int64
	OwnerId           string
	Phone             string
	Rating            string
	Status            string
	Website           string
}

type Product struct {
	CreatedById           string
	CreatedDate           string
	Description           string
	DisplayUrl            string
	ExternalDataSourceId  string
	ExternalId            string
	Family                string
	Id                    string
	IsActive              bool
	IsArchived            bool
	IsDeleted             bool
	LastModifiedById      string
	LastModifiedDate      string
	LastReferencedDate    string
	LastViewedDate        string
	Name                  string
	ProductCode           string
	QuantityUnitOfMeasure string
	StockKeepingUnit      string
	SystemModstamp        string
}
