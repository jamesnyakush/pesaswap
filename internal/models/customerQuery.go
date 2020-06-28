package models

// Querying Customers in the db
type CustomerQuery struct {
	ConsumerKey string `json:"consumer_key"`
	ApiKey	string `json:"api_key"`
	ExternalId string `json:"external_id"`
}
