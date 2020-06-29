package api

import (
	"encoding/json"
	"pesaswap/internal/models"
)

func (s Service) QueryTransactionExternalId(TransExtId string) (string, error) {
	queryByTransExtId := &models.QueryTransactionExternalId{
		ConsumerKey:           s.consumerKey,
		ApiKey:                s.appKey,
		TransactionExternalId: TransExtId,
	}

	body, err := json.Marshal(queryByTransExtId)

	if err != nil {
		return "", err
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	url := s.baseURL() + GetTransactionRecord
	return s.newReq(url, body, headers)
}

func (s Service) QueryTransactionId(TransId string) (string, error) {

	queryByTransId := &models.QueryTransactionId{
		ConsumerKey:   s.consumerKey,
		ApiKey:        s.appKey,
		TransactionId: TransId,
	}

	body, err := json.Marshal(queryByTransId)

	if err != nil {
		return "", err
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	url := s.baseURL() + FilterByTransactionId
	return s.newReq(url, body, headers)
}

func (s Service) QueryAll() (string, error) {
	queryAll := &models.QueryAll{
		ConsumerKey: s.consumerKey,
		ApiKey:      s.appKey,
	}

	body, err := json.Marshal(queryAll)

	if err != nil {
		return "", err
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	url := s.baseURL() + ViewAllTransactionsRecord
	return s.newReq(url, body, headers)
}

func (s Service) CardByRequestId(RequestId string) (string, error) {

	cardByRequestId := &models.CardByRequestId{
		ConsumerKey: s.consumerKey,
		ApiKey:      s.appKey,
		RequestId:   RequestId,
	}

	body, err := json.Marshal(cardByRequestId)

	if err != nil {
		return "", err
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	url := s.baseURL() + IveriTransaction
	return s.getReq(url, body, headers)
}

func (s Service) CardByAppId(AppId string) (string, error) {

	cardByAppId := &models.CardByAppId{
		ConsumerKey:   s.consumerKey,
		ApiKey:        s.appKey,
		ApplicationId: AppId,
	}

	body, err := json.Marshal(cardByAppId)

	if err != nil {
		return "", err
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	url := s.baseURL() + FilterByApplicationId
	return s.newReq(url, body, headers)
}
