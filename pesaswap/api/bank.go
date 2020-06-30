package api

import (
	"encoding/json"
	"github.com/nyumbapoa/pesaswap/pesaswap/models"
)

func (s Service) CardPayment(Currency string, Amount string, ExpiryDate string, CardSecurityCode string, CardNumber string, TransactionId string, CustomerId string) (string, error) {

	cardPayment := &models.CardPayment{
		ConsumerKey:           s.consumerKey,
		ApiKey:                s.appKey,
		Currency:              Currency,
		Amount:                Amount,
		ExpiryDate:            ExpiryDate,
		CardSecurityCode:      CardSecurityCode,
		CardNumber:            CardNumber,
		TransactionExternalId: TransactionId,
		CustomerExternalId:    CustomerId,
	}

	body, err := json.Marshal(cardPayment)

	if err != nil {
		return "", err
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	url := s.baseURL() + CardPayment
	return s.newReq(url, body, headers)
}

func (s Service) AccountBalance(account string, env string) (string, error) {

	accountBalance := &models.AccountBalance{
		ConsumerKey:   s.consumerKey,
		ApiKey:        s.appKey,
		AccountNumber: account,
		Environment:   env,
	}

	body, err := json.Marshal(accountBalance)

	if err != nil {
		return "", err
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	url := s.baseURL() + AccountBalance
	return s.newReq(url, body, headers)
}

func (s Service) FundTransfer(Source string, Destination string, SourceNarration string, DestinationNarration string, Amount string, Currency string, BankCode string, Environment string, CustomerID string, TransactionID string) (string, error) {

	fundTransfer := &models.FundTransfer{
		ConsumerKey:              s.consumerKey,
		ApiKey:                   s.appKey,
		SourceAccountNumber:      Source,
		DestinationAccountNumber: Destination,
		SourceNarration:          SourceNarration,
		DestinationNarration:     DestinationNarration,
		Amount:                   Amount,
		Currency:                 Currency,
		BankCode:                 BankCode,
		Environment:              Environment,
		CustomerExternalId:       CustomerID,
		TransactionExternalId:    TransactionID,
	}

	body, err := json.Marshal(fundTransfer)

	if err != nil {
		return "", err
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	url := s.baseURL() + FundTransfer
	return s.newReq(url, body, headers)
}

func (s Service) SendToMpesa(Source string, Destination string, SourceNarration string, DestinationNarration string, Amount string, Currency string, Env string, CustomerId string, TransactionId string) (string, error) {

	sendToMpesa := &models.SendToMpesa{
		ConsumerKey:             s.consumerKey,
		ApiKey:                  s.appKey,
		SourceAccountNumber:     Source,
		DestinationMobileNumber: Destination,
		SourceNarration:         SourceNarration,
		DestinationNarration:    DestinationNarration,
		Amount:                  Amount,
		Currency:                Currency,
		Environment:             Env,
		CustomerExternalId:      CustomerId,
		TransactionExternalId:   TransactionId,
	}

	body, err := json.Marshal(sendToMpesa)

	if err != nil {
		return "", err
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	url := s.baseURL() + SendToMpesa
	return s.newReq(url, body, headers)
}

func (s Service) StatusQuery(MsgReference string, env string) (string, error) {

	statusQuery := &models.StatusQuery{
		ConsumerKey:      s.consumerKey,
		ApiKey:           s.appKey,
		MessageReference: MsgReference,
		Environment:      env,
	}

	body, err := json.Marshal(statusQuery)

	if err != nil {
		return "", err
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	url := s.baseURL() + StatusQuery
	return s.newReq(url, body, headers)
}
