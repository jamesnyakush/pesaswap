package models

// Working with card payment
type CardPayment struct {
	ConsumerKey           string `json:"consumer_key"`
	ApiKey                string `json:"api_key"`
	Currency              string `json:"currency"`
	Amount                string `json:"amount"`
	ExpiryDate            string `json:"expiry_date"`
	CardSecurityCode      string `json:"card_security_code"`
	CardNumber            string `json:"card_number"`
	TransactionExternalId string `json:"transaction_external_id"`
	CustomerExternalId    string `json:"customer_external_id"`
}
