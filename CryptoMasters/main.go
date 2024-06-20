package main

import (
	"fmt"
	"learningo.com/go/crypto/api"
	"sync"
)

func main() {
	currencies := []string{"BTC", "ETH", "BCH"}
	var wg sync.WaitGroup
	for _, currency := range currencies {
		wg.Add(1)
		go func() {
			GetCurrencyData(currency)
			wg.Done()
		}()
	}
	wg.Wait()
}

func GetCurrencyData(curr string) {
	rate, err := api.GetRate(curr)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("The rate for %v is %.2f \n", rate.Currency, rate.Price)
	}
}
