package main

import (
	"fmt"
	"github.com/jamesnyakush/pesaswap/pesaswap/api"
	"log"
)

const (
	consumerKey = "dMMcTvqhxl"
	appSecret   = "wjqRX5tTSzFyJbwW33ANzwh9V"
)

func main() {
	fmt.Println("Welcome to pesaswap")

	svc, err := api.New(consumerKey, appSecret, api.SANDBOX)

	if err != nil {
		panic(err)
	}

		res, err := svc.NewUser(
			"James",
			"Nyakush",
			"nyakush@gmail.com",
			"254746445198",
		)


	/*
	   res, err := svc.CardPayment(
	   		"KES",
	   		"10",
	   		"012022",
	   		"111",
	   		"4242424242424242",
	   		"AC312E60-BA4B-11EA-A700-4903572028c",
	   		"AC312E60-BA4B-11EA-A700-4903572028FB",
	   	)
	*/
	/*
	   res, err := svc.PaymentRequests(
	   		"434334",
	   		"Water Bill",
	   		"1",
	   		"10/07/2020",
	   		"10/06/2020",
	   		"AC312E60-BA4B-11EA-A700-4903572028FB",
	   		"0746445198",
	   		100,
	   	)
	*/

	/*
	   res, err := svc.SendToMpesa(
	   	"4242424242424242",
	   	"254746445198",
	   	"test",
	   	"test",
	   	"1",
	   	"KES",
	   	"Test",
	   	"AC312E60-BA4B-11EA-A700-4903572028FB",
	   	"AC312E60-BA4B-11EA-A700-4903572028FB",
	   	)
	*/
	/*	res, err := svc.CouponPayments(
		"smN94",
		"1",
		"AC312E60-BA4B-11EA-A700-4903572028FB",
		"AC312E60-BA4B-11EA-A700-4903572028FA",
	)*/

	//res, err := svc.QueryAll()

	if err != nil {
		log.Println(err)
	}

	log.Println(res)
}

/*
const (
	consumerKey = "uLdedoqenx"
	appSecret   = "IWheLlsxP3ZuAZ1357z3A1zn8"
)
*/
