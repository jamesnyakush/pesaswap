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

// Lipa na mpesa
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

// Working with Reversal in PesaSwap
type ReversalPesaSwap struct {
	Method                string `json:"method"`
	ConsumerKey           string `json:"consumer_key"`
	ApiKey                string `json:"api_key"`
	PaybillDescription    string `json:"paybill_description"`
	Amount                string `json:"Amount"`
	ReceiverParty         string `json:"ReceiverParty"`
	TransactionID         string `json:"TransactionID"`
	TransactionExternalId string `json:"transaction_external_id"`
}

