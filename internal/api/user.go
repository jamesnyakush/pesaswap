package api

import (
	"encoding/json"
	"github.com/nyumbapoa/pesaswap/internal/models"
)

func (s Service) NewUser(firstname string, lastname string, email string, phone string) (string, error) {

	userObject := &models.User{
		ConsumerKey: s.consumerKey,
		ApiKey:      s.appKey,
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

	url := s.baseURL() + CreateUser
	return s.newReq(url, body, headers)
}

func (s Service) UserUpdate(firstname string, lastname string, email string, phone string, address1 string, address2, country string, externalID string) (string, error) {

	userObject := &models.User{
		ConsumerKey: s.consumerKey,
		ApiKey:      s.appKey,
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

	url := s.baseURL() + UpdateUser
	return s.newReq(url, body, headers)
}

func (s Service) CustomerQuery(externalID string) (string, error) {

	customerQuery := &models.CustomerQuery{
		ConsumerKey: s.consumerKey,
		ApiKey:      s.appKey,
		ExternalId:  externalID,
	}

	body, err := json.Marshal(customerQuery)

	if err != nil {
		return "", err
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	url := s.baseURL() + QueryCustomer
	return s.newReq(url, body, headers)
}
