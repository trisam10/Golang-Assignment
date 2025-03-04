package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	url := "https://api.coindesk.com/v1/bpi/currentprice.json"

	// Make a GET request
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("❌ Error connecting to CoinDesk API: %v", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("❌ Error reading response body: %v", err)
	}

	// Print the response
	fmt.Println("✅ Successfully connected to CoinDesk API")
	fmt.Println("Response:")
	fmt.Println(string(body))
}
