package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kubaj/iota/config"

	"github.com/iotaledger/giota"
)

func main() {
	config, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatalln(err)
	}

	api := giota.NewAPI(config.NodeAddress, &http.Client{})
	seedTrytes, _ := giota.ToTrytes(config.Seed)
	addr, _ := giota.NewAddress(seedTrytes, 0, 1)
	balances, err := api.Balances([]giota.Address{addr})
	fmt.Println(balances)
	fmt.Println(err)
}
