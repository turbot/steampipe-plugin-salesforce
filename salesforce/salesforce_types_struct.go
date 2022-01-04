package salesforce

// TODO: Handle time fields
type Account struct {
	Id                string
	Name              string
	OwnerId           string
	CreatedById       string
	CreatedDate       string
	Industry          string
	AnnualRevenue     float64
	NumberOfEmployees int64
	LastModifiedById  string
	LastModifiedDate  string
	Website           string
	Phone             string
	AccountSource     string
	Type              string
	BillingAddress    map[string]string
	ShippingAddress   map[string]string
}

type User struct {
	AboutMe                string
	AccountId              string
	CreatedById            string
	CreatedDate            string
	Email                  string
	FirstName              string
	Id                     string
	IsActive               bool
	LastLoginDate          string
	LastModifiedById       string
	LastModifiedDate       string
	LastName               string
	LastPasswordChangeDate string
	Name                   string
	ProfileId              string
	Username               string
	// LastPasswordChangeDate *time.Time
	// LastModifiedDate       *time.Time
	// CreatedDate            *time.Time
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
	LastReferencedDate       string
	LastViewedDate           string
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
	// CurrentGeneratorsC          string `json:"CurrentGenerators__c,omitempty"`
	// DeliveryInstallationStatusC string `json:"DeliveryInstallationStatus__c,omitempty"`
	// MainCompetitorsC            string `json:"MainCompetitors__c,omitempty"`
	// OrderNumberC                string `json:"OrderNumber__c,omitempty"`
	// TrackingNumberC             string `json:"TrackingNumber__c,omitempty"`
}
