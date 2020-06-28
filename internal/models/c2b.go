package models

// Working with Mpesa C2B with pesaswap
type C2BPesaSwap struct {
	Method                string `json:"method"`
	ConsumerKey           string `json:"consumer_key"`
	ApiKey                string `json:"api_key"`
	PaybillDescription    string `json:"paybill_description"`
	CommandID             string `json:"CommandID"`
	Msisdn                string `json:"Msisdn"`
	Amount                string `json:"Amount"`
	TransactionExternalId string `json:"transaction_external_id"`
	BillRefNumber         string `json:"BillRefNumber"`
	CustomerExternalId    string `json:"customer_external_id"`
}
