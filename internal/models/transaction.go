package models

// Query by transaction external id
type QueryTransactionExternalId struct {
	ConsumerKey           string `json:"consumer_key"`
	ApiKey                string `json:"api_key"`
	TransactionExternalId string `json:"transaction_external_id"`
}

// Query by transaction id
type QueryTransactionId struct {
	ConsumerKey   string `json:"consumer_key"`
	ApiKey        string `json:"api_key"`
	TransactionId string `json:"transaction_id"`
}

// Query All
type QueryAll struct {
	ConsumerKey string `json:"consumer_key"`
	ApiKey      string `json:"api_key"`
}

// Card by Request Id
type CardByRequestId struct {
	ConsumerKey string `json:"consumer_key"`
	ApiKey      string `json:"api_key"`
	RequestId   string `json:"request_id"`
}

// Card By App Id
type CardByAppId struct {
	ConsumerKey   string `json:"consumer_key"`
	ApiKey        string `json:"api_key"`
	ApplicationId string `json:"application_id"`
}
