package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"learningo.com/go/crypto/datatypes"
	"net/http"
	"strconv"
)

const apiUrl = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*datatypes.Rate, error) {
	if len(currency) == 0 {
		return nil, errors.New("currency empty")
	}
	res, err := http.Get(fmt.Sprintf(apiUrl, currency))
	rate := datatypes.Rate{}

	if err != nil {
		return nil, err
	}
	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		var result CEXResponse
		json.Unmarshal(bodyBytes, &result)
		rate.Currency = currency
		price, err := strconv.ParseFloat(result.High, 64)
		if err == nil {
			rate.Price = price
		}
	} else {
		return nil, fmt.Errorf("status code error: %v", res.StatusCode)
	}
	return &rate, nil
}
