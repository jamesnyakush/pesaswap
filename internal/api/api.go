package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"pesaswap/internal/models"
	"time"
)

type Env string

const (
	SANDBOX = iota
	PRODUCTION
)

type Service struct {
	consumer_key string
	app_key      string
	Env          int
}
type Mpesa struct {
	c2b string
}

func New(consumer_key, app_key string, env int) (Service, error) {
	return Service{consumer_key, app_key, env}, nil
}

func (s Service) UserCreate(firstname string, lastname string, email string, phone string) (string, error) {

	userObject := &models.User{
		ConsumerKey: s.consumer_key,
		ApiKey:      s.app_key,
		Firstname:   firstname,
		Lastname:    lastname,
		Email:       email,
		Phone:       phone,
	}

	body, err := json.Marshal(userObject)

	if err != nil {
		return "", err
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	url := s.baseURL() + "api/pesaswap/create/customer"
	return s.newReq(url, body, headers)

}

func (s Service) UpdateUser(firstname string, lastname string, email string, phone string, address1 string, address2, country string, externalID string) (string, error) {

	userObject := &models.User{
		ConsumerKey: s.consumer_key,
		ApiKey:      s.app_key,
		Firstname:   firstname,
		Lastname:    lastname,
		Email:       email,
		Phone:       phone,
		Address1:    address1,
		Address2:    address2,
		Country:     country,
		ExternalId:  externalID,
	}

	body, err := json.Marshal(userObject)

	if err != nil {
		return "", err
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	url := s.baseURL() + "api/pesaswap/update/customer"
	return s.newReq(url, body, headers)
}

func (s Service) C2BPesaSwap(desc string, command string, msisdn string, amount string, customerId string) (string, error) {

	c2bObject := &models.C2BPesaSwap{
		Method:             "c2b",
		ConsumerKey:        s.consumer_key,
		ApiKey:             s.app_key,
		PaybillDescription: desc,
		CommandID:          command,
		Msisdn:             msisdn,
		Amount:             amount,
		CustomerExternalId: customerId,
	}

	body, err := json.Marshal(c2bObject)

	if err != nil {
		return "", err
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	url := s.baseURL() + "api/pesaswap/mobile-payment"
	return s.newReq(url, body, headers)
}

func (s Service) B2CPesaSwap() (string, error) {

	b2cObject := &models.B2CPesaSwap{
		Method: "b2c",
		ConsumerKey: s.consumer_key,
		ApiKey: s.app_key,

	}

	body, err := json.Marshal(b2cObject)

	if err != nil {
		return "", err
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	url := s.baseURL() + "api/pesaswap/mobile-payment"
	return s.newReq(url, body, headers)
}

/**/
func (s Service) newReq(url string, body []byte, headers map[string]string) (string, error) {
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))

	if err != nil {
		return "", err
	}

	for key, value := range headers {
		request.Header.Set(key, value)
	}

	client := &http.Client{Timeout: 60 * time.Second}
	res, err := client.Do(request)

	if res != nil {
		defer res.Body.Close()
	}

	if err != nil {
		return "", err
	}

	stringBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(stringBody), nil
}

/**/
func (s Service) baseURL() string {
	if s.Env == PRODUCTION {
		return "https://www.pesaswap.com/"
	}
	return "https://devpesaswap.azurewebsites.net/"
}
