package api

import (
	"encoding/json"
	"github.com/jamesnyakush/pesaswap/pesaswap/models"
)

func (s Service) CouponPayments(identifier string, amount string, customerID string, transID string) (string, error) {

	couponPayment := &models.CouponPayment{
		ConsumerKey:           s.consumerKey,
		ApiKey:                s.appKey,
		Identifier:            identifier,
		Amount:                amount,
		CustomerExternalId:    customerID,
		TransactionExternalId: transID,
	}

	body, err := json.Marshal(couponPayment)

	if err != nil {
		return "", err
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	url := s.baseURL() + CouponPayment
	return s.newReq(url, body, headers)
}

func (s Service) CouponQuerys(status string, identifier string) (string, error) {

	couponQuery := &models.CouponQuery{
		ConsumerKey: s.consumerKey,
		ApiKey:      s.appKey,
		Status:      status,
		Identifier:  identifier,
	}

	body, err := json.Marshal(couponQuery)

	if err != nil {
		return "", err
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	url := s.baseURL() + CouponQueryEndpoint
	return s.newReq(url, body, headers)
}
