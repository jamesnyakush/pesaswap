package models

// Sending payment request
type PaymentRequest struct {
	ConsumerKey        string `json:"consumer_key"`
	ApiKey             string `json:"api_key"`
	PaybillDescription string `json:"paybill_description"`
	Description        string `json:"description"`
	Range              string `json:"range"`
	BillingDate        string `json:"billing_date"`
	LastBillingDate    string `json:"last_billing_date"`
	ExternalId         string `json:"external_id"`
	PhoneNumber        string `json:"phone_number"`
	TotalAmount        int    `json:"total_amount"`
}
