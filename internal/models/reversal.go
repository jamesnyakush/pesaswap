package models

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
