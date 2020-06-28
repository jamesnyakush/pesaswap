# Pesaswap :moneybag: :credit_card:
This is unofficial Golang wrapper to pesaswap Api

- [Customer](#customer)

# Customer

### Creating a new user


```go
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

	
    res, err := svc.UserCreate(
            "J",
            "M",
            "j@gmail.com",
            "0746445198",
    )
	
	if err != nil {
		log.Println(err)
	}

	log.Println(res)
}
```

## Updating user

```go
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
	

	if err != nil {
		log.Println(err)
	}

	log.Println(res)
}

```