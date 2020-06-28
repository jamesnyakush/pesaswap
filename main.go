package main

import (
	"fmt"
	"log"
	"pesaswap/internal/api"
)

const (
	consumerKey = ""
	appSecret   = ""
)

func main() {
	fmt.Println("Welcome to pesaswap")

	svc, err := api.New(consumerKey, appSecret, api.PRODUCTION)

	if err != nil {
		panic(err)
	}

	/*
		res, err := svc.UserCreate(
				"J",
				"M",
				"j@gmail.com",
				"0746445198",
			)
	*/

	/*
		res, err := svc.UpdateUser(
			"J",
			"B",
			"jec@gmail.com",
			"254770977160",
			"Utalii street",
			"Monrovia",
			"Kenya",
			"2C55D6F0-B90j2-11EA-B7B0-7DC668kllkml",
		)
	*/
	res, err := svc.C2BPesaSwap(
		"Hey there",
		"CustomerPayBillOnline",
		"254723722363",
		"5",
		"2C55D6F0-B90j2-11EA-B7B0-kjknasknkjn",
	)

	if err != nil {
		log.Println(err)
	}

	log.Println(res)
}
