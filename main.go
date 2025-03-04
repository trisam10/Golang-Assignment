package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Struct to match CoinGecko API response
type CoinGeckoResponse struct {
	Bitcoin struct {
		USD float64 `json:"usd"`
	} `json:"bitcoin"`
}

var cache struct {
	price     float64
	timestamp time.Time
}

func getCryptoPrice(c *gin.Context) {
	// Cache expiry time (e.g., 1 minute)
	cacheExpiry := time.Minute

	// Check if cache is valid
	if time.Since(cache.timestamp) < cacheExpiry {
		c.JSON(http.StatusOK, gin.H{"price": cache.price})
		return
	}

	// Fetch from CoinGecko API
	url := "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=usd"
	resp, err := http.Get(url)
	if err != nil {
		log.Println("❌ Error fetching data:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
		return
	}
	defer resp.Body.Close()

	// Decode JSON response
	var result CoinGeckoResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Println("❌ Error decoding JSON:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse data"})
		return
	}

	// Store in cache
	cache.price = result.Bitcoin.USD
	cache.timestamp = time.Now()

	// Return price
	c.JSON(http.StatusOK, gin.H{"price": cache.price})
}

func main() {
	r := gin.Default()
	r.GET("/crypto-price", getCryptoPrice)
	fmt.Println("🚀 Server running on port 8080")
	r.Run(":8080")
}
