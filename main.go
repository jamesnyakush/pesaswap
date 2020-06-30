package main

import (
	"fmt"
	"github.com/nyumbapoa/pesaswap/pesaswap/api"
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
		"Norman",
		"Nuthu",
		"norman.nuthu@gmail.com",
		"254746445198",
	)

	//res, err := svc.QueryAll()

	if err != nil {
		log.Println(err)
	}

	log.Println(res)
}
