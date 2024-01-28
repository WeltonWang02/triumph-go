package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"triumph_intern/models"
)

type OrderBook struct {
	Asks [][]interface{} `json:"asks"`
	Bids [][]interface{} `json:"bids"`
}

type KrakenResponse struct {
	Error  []interface{}        `json:"error"`
	Result map[string]OrderBook `json:"result"`
}

type PriceInfo struct {
	Price    float64
	Exchange string
}

func ExecuteOrder(quantity string, symbol string, action string) (models.Order, error) {
	// Logic to choose the best exchange and execute the order
	// Fetch data from APIs, calculate the best price
	// Mock response for demonstration
	amount, err := strconv.ParseFloat(quantity, 64)
	if err != nil {
		return models.Order{}, fmt.Errorf("Invalid amount of currency: %w", err)
	}

	bestPrice, err := getBestPrice(amount, symbol, action)
	if err != nil {
		return models.Order{}, fmt.Errorf("Error fetching best price: %w", err)
	}

	return models.Order{BtcAmount: amount, UsdAmount: bestPrice.Price, Exchange: bestPrice.Exchange}, nil
}

func getBestPrice(amount float64, symbol string, action string) (PriceInfo, error) {
	// Get the best price to buy/sell given a string asset
	// Calls getBookData to get data, merges the respective l2data
	// Sorts by price to get ideal data
	// amount is an unused parameter right now
	// TODO: implement return multiple PriceInfo objects based on the Bid/Ask quantity and provided quantity

	exchanges := []string{"coinbase", "kraken"} // Add more exchanges as needed

	var allPrices []PriceInfo

	for _, exchange := range exchanges {
		book, err := getBookData(symbol, exchange)
		if err != nil {
			continue // check the next source
		}

		if action == "buy" {
			for _, ask := range book.Asks {
				price, err := strconv.ParseFloat(ask[0].(string), 64)
				if err == nil {
					allPrices = append(allPrices, PriceInfo{Price: price, Exchange: exchange})
				}
			}
		} else if action == "sell" {
			for _, bid := range book.Bids {
				price, err := strconv.ParseFloat(bid[0].(string), 64)
				if err == nil {
					allPrices = append(allPrices, PriceInfo{Price: price, Exchange: exchange})
				}
			}
		}
	}

	if len(allPrices) == 0 {
		return PriceInfo{}, fmt.Errorf("no prices available for action %s", action)
	}

	// Sort prices based on the action
	sort.Slice(allPrices, func(i, j int) bool {
		if action == "buy" {
			return allPrices[i].Price < allPrices[j].Price
		}
		return allPrices[i].Price > allPrices[j].Price
	})

	// Return the best price based on the action
	return allPrices[0], nil
}

func getBookData(product string, exchange string) (OrderBook, error) {
	// Determine the URL based on the product and exchange
	// Send HTTP request and return l2 data book
	var url string

	if exchange == "coinbase" {
		url = "https://api.exchange.coinbase.com/products/" + product + "/book"
	} else { // default to krakem
		url = "https://api.kraken.com/0/public/Depth?pair=" + product
	}

	response, err := sendHTTPrequest(url)
	if err != nil {
		return OrderBook{}, err
	}

	var l2data OrderBook

	if exchange == "coinbase" {
		if err := json.Unmarshal([]byte(response), &l2data); err != nil {
			return OrderBook{}, fmt.Errorf("error parsing response JSON: %w", err)
		}
	} else {
		var kraken KrakenResponse
		if err := json.Unmarshal([]byte(response), &kraken); err != nil {
			return OrderBook{}, fmt.Errorf("error parsing response JSON: %w", err)
		}

		var exists bool
		l2data, exists = kraken.Result[product]

		if !exists {
			return OrderBook{}, fmt.Errorf("product data not found in response")
		}
	}

	return l2data, nil
}

func sendHTTPrequest(url string) (string, error) {
	// Send HTTP GET Request to given url
	// Return either string response or error
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("received non-200 status code: %d", res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil

}
