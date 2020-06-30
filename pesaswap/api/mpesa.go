package api

import (
	"encoding/json"
	"github.com/nyumbapoa/pesaswap/pesaswap/models"
)

var (
	customer2Business = "c2b"
	business2Customer = "b2c"
	mpesaExpress      = "op"
	reversal          = "rp"
)

func (s Service) C2B(desc string, command string, msisdn string, amount string, customerId string) (string, error) {

	c2bObject := &models.C2BPesaSwap{
		Method:             customer2Business,
		ConsumerKey:        s.consumerKey,
		ApiKey:             s.appKey,
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

	url := s.baseURL() + MpesaPesaSwap
	return s.newReq(url, body, headers)
}

func (s Service) STK(desc string, phone string, amount string, transID string, customerID string) (string, error) {

	mpesaExpressObject := &models.MpesaExpress{
		Method:                mpesaExpress,
		ConsumerKey:           s.consumerKey,
		ApiKey:                s.appKey,
		PaybillDescription:    desc,
		PhoneNumber:           phone,
		Amount:                amount,
		TransactionExternalId: transID,
		CustomerExternalId:    customerID,
	}

	body, err := json.Marshal(mpesaExpressObject)

	if err != nil {
		return "", err
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	url := s.baseURL() + MpesaPesaSwap

	return s.newReq(url, body, headers)
}

func (s Service) B2C(desc string, command string, phone string, amout string, externalID string, customerID string) (string, error) {

	b2cObject := &models.B2CPesaSwap{
		Method:                business2Customer,
		ConsumerKey:           s.consumerKey,
		ApiKey:                s.appKey,
		PaybillDescription:    desc,
		CommandID:             command,
		PartyB:                phone,
		Amount:                amout,
		TransactionExternalId: externalID,
		CustomerExternalId:    customerID,
	}

	body, err := json.Marshal(b2cObject)

	if err != nil {
		return "", err
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	url := s.baseURL() + MpesaPesaSwap

	return s.newReq(url, body, headers)
}

func (s Service) Reversal(desc string, amount string, shortcode string, transID string, transExtID string) (string, error) {

	reversalObject := &models.ReversalPesaSwap{
		Method:                reversal,
		ConsumerKey:           s.consumerKey,
		ApiKey:                s.appKey,
		PaybillDescription:    desc,
		Amount:                amount,
		ReceiverParty:         shortcode,
		TransactionID:         transID,
		TransactionExternalId: transExtID,
	}

	body, err := json.Marshal(reversalObject)

	if err != nil {
		return "", err
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	url := s.baseURL() + MpesaPesaSwap

	return s.newReq(url, body, headers)
}