package models

// working with business to customer in pesaswap
type B2CPesaSwap struct {
	Method                string `json:"method"`
	ConsumerKey           string `json:"consumer_key"`
	ApiKey                string `json:"api_key"`
	PaybillDescription    string `json:"paybill_description"`
	CommandID             string `json:"CommandID"`
	PartyB                string `json:"PartyB"`
	Amount                string `json:"Amount"`
	TransactionExternalId string `json:"transaction_external_id"`
	CustomerExternalId    string `json:"customer_external_id"`
}
