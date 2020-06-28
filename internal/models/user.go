package models

// Creating and updating users in pesaswap
type User struct {
	ConsumerKey string `json:"consumer_key"`
	ApiKey      string `json:"api_key"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Address1    string `json:"address1"`
	Address2    string `json:"address2"`
	State       string `json:"state"`
	Country     string `json:"country"`
	ExternalId  string `json:"external_id"`
}
