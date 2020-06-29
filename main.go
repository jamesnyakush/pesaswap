package main

import (
	"fmt"
	"log"
	"pesaswap/internal/api"
)

const (
	consumerKey = "uLdedoqenx"
	appSecret   = "IWheLlsxP3ZuAZ1357z3A1zn8"
)

func main() {
	fmt.Println("Welcome to pesaswap")

	svc, err := api.New(consumerKey, appSecret, api.PRODUCTION)

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
