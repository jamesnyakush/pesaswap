package api

import (
	"encoding/json"
	"github.com/nyumbapoa/pesaswap/internal/models"
)

func (s Service) PaymentRequests(paybillDesc string, desc string, ranges string, billingDate string, lastBillingDate string, externalID string, phone string, amount string) (string, error) {

	paymentRequest := &models.PaymentRequest{
		ConsumerKey:        s.consumerKey,
		ApiKey:             s.appKey,
		PaybillDescription: paybillDesc,
		Description:        desc,
		Range:              ranges,
		BillingDate:        billingDate,
		LastBillingDate:    lastBillingDate,
		ExternalId:         externalID,
		PhoneNumber:        phone,
		TotalAmount:        amount,
	}

	body, err := json.Marshal(paymentRequest)

	if err != nil {
		return "", err
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	url := s.baseURL() + PaymentRequest
	return s.newReq(url, body, headers)
}
