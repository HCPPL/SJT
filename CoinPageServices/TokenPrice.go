package main

import (
    "encoding/json"
    "net/http"
	"math/big"
	"log"
	//"bytes"
)
type tokenPrice struct {
	TokenPriceInUSD string
	TokenPriceInETH *big.Int
}

var getTokenPrice tokenPrice

//func GetTokenPrice(w http.ResponseWriter, req *http.Request) {
var GetTokenPrice = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){   
	CallingLog("GetTokenPrice","TokenPrice.go", "NULL", "NULL","NULL")
	getTokenPrice := new(tokenPrice)
	
	TotalTokensForUSD, TotalTokensForETH := getTokenPriceInfo()
	
    getTokenPrice.TokenPriceInUSD = TotalTokensForUSD
    getTokenPrice.TokenPriceInETH = TotalTokensForETH
	b, err := json.Marshal(getTokenPrice)
	
		if err != nil {
	
			log.Println("Error: %s", err)
			throwException("GetTokenPrice", "TokenPrice.go", "NULL", "NULL","Fatal Error")
			return
			
	
		}
	
		var s5 string= string(b)
			
	ResultLog ("GetTokenPrice","TokenPrice.go", s5,"NULL")
	json.NewEncoder(w).Encode(getTokenPrice)  
	
})

func getTokenPriceInfo()(NoOfTokensInOneUSD string, NoOfTokensInOneETH *big.Int) {

//	log.Println("=================> getTokenPriceInfo called")

	CallingLog("getTokenPriceInfo","TokenPrice.go", "NULL", "NULL","NULL")
	n1 := "1"

	CurrentTokenPriceInUSD := getSellPrice()
//	log.Println("=================> getSellPrice called", CurrentTokenPriceInUSD)
	
	NoOfTokensInOneUSD = doDivision(n1, CurrentTokenPriceInUSD)
//	log.Println("=================> NoOfTokensInOneUSD called")

	CurrentEthPriceInUSD := price("ETH")
//	log.Println("=================> CurrentEthPriceInUSD called")

	CurrentEthPriceInUSD_strings := CurrentEthPriceInUSD.String()
	
	NoOfTokensInOneETH = doMultiplication(CurrentEthPriceInUSD_strings, NoOfTokensInOneUSD, n1)
//	log.Println("=================> doMultiplication ")
	
	ResultLog ("getTokenPriceInfo","TokenPrice.go", "NULL","NULL")
	return
}

