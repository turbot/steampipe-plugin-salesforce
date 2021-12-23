package salesforce

import "time"

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
	CreatedDate            *time.Time
	Email                  string
	FirstName              string
	Id                     string
	IsActive               bool
	LastLoginDate          string
	LastModifiedById       string
	LastModifiedDate       *time.Time
	LastName               string
	LastPasswordChangeDate *time.Time
	Name                   string
	ProfileId              string
	Username               string
}
