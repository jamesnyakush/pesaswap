package models

// Working with card payment
type CardPayment struct {
	ConsumerKey           string `json:"consumer_key"`
	ApiKey                string `json:"api_key"`
	Currency              string `json:"currency"`
	Amount                string `json:"amount"`
	ExpiryDate            string `json:"expiry_date"`
	CardSecurityCode      string `json:"card_security_code"`
	CreditCardNumber      string `json:"credit_card_number"`
	CustomerExternalId    string `json:"customer_external_id"`
	TransactionExternalId string `json:"transaction_external_id"`
	Environment           string `json:"environment"`
}

// Checking for Account Balance
type AccountBalance struct {
	ConsumerKey   string `json:"consumer_key"`
	ApiKey        string `json:"api_key"`
	AccountNumber string `json:"account_number"`
	Environment   string `json:"environment"`
}

// Sending Fund transfer
type FundTransfer struct {
	ConsumerKey              string `json:"consumer_key"`
	ApiKey                   string `json:"api_key"`
	SourceAccountNumber      string `json:"source_account_number"`
	DestinationAccountNumber string `json:"destination_account_number"`
	SourceNarration          string `json:"source_narration"`
	DestinationNarration     string `json:"destination_narration"`
	Amount                   string `json:"amount"`
	Currency                 string `json:"currency"`
	BankCode                 string `json:"bank_code"`
	Environment              string `json:"environment"`
	CustomerExternalId       string `json:"customer_external_id"`
	TransactionExternalId    string `json:"transaction_external_id"`
}

// send to mpesa
type SendToMpesa struct {
	ConsumerKey             string `json:"consumer_key"`
	ApiKey                  string `json:"api_key"`
	SourceAccountNumber     string `json:"source_account_number"`
	DestinationMobileNumber string `json:"destination_mobile_number"`
	SourceNarration         string `json:"source_narration"`
	DestinationNarration    string `json:"destination_narration"`
	Amount                  string `json:"amount"`
	Currency                string `json:"currency"`
	Environment             string `json:"environment"`
	CustomerExternalId      string `json:"customer_external_id"`
	TransactionExternalId   string `json:"transaction_external_id"`
}

// Status Query
type StatusQuery struct {
	ConsumerKey      string `json:"consumer_key"`
	ApiKey           string `json:"api_key"`
	MessageReference string `json:"message_reference"`
	Environment      string `json:"environment"`
}
