package models

// Working with Coupon payment
type CouponPayment struct {
	ConsumerKey           string `json:"consumer_key"`
	ApiKey                string `json:"api_key"`
	Identifier            string `json:"identifier"`
	Amount                string `json:"amount"`
	CustomerExternalId    string `json:"customer_external_id"`
	TransactionExternalId string `json:"transaction_external_id"`
}

type CouponQuery struct {
	ConsumerKey string `json:"consumer_key"`
	ApiKey      string `json:"api_key"`
	Status      string `json:"status"`
	Identifier  string `json:"identifier"`
}
