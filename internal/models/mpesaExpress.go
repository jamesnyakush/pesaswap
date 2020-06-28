package models

type MpesaExpress struct {
	Method                string `json:"method"`
	ConsumerKey           string `json:"consumer_key"`
	ApiKey                string `json:"api_key"`
	PaybillDescription    string `json:"paybill_description"`
	PhoneNumber           string `json:"PhoneNumber"`
	Amount                string `json:"Amount"`
	TransactionExternalId string `json:"transaction_external_id"`
	CustomerExternalId    string `json:"customer_external_id"`
	CodeType              string `json:"code_type"`
}
