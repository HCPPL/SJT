package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"math/big"
	//"bytes"
)

type AmountResponse struct {
	Amount   string `json:amount`
	Currency string `json:currency`
}

type DataResponse struct {
	Data AmountResponse `json:data`
}

var data = DataResponse{}
var AmtResponse AmountResponse

func price(coin string) (*big.Int){
	CallingLog("price()","GetEthPriceInUSD.go", "NULL", "NULL","NULL")
	
	client := &http.Client{}
	req, reqErr := http.NewRequest("GET", fmt.Sprintf("https://api.coinbase.com/v2/prices/%s-USD/spot", coin), nil)
	if reqErr != nil {
		log.Fatal(reqErr)
	}

	req.Header.Set("CB-VERSION", "2017-03-25")
	res, resErr := client.Do(req)
	if resErr != nil {
		log.Fatal(resErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if readErr != nil {
		log.Fatal(readErr)
	}

	jsonErr := json.Unmarshal(body, &data)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	ip := new (big.Int)
	ip.SetString(data.Data.Amount, 10)
	//bigint := big.NewInt(ip)
	//bigstr := bigint.String()
	ResultLog ("price()","GetEthPriceInUSD.go", "NULL","NULL")
	//str := fmt.Sprint(ip)
	return ip
}

/*
var GetETHPrice = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request){
//func GetETHPrice(w http.ResponseWriter, req *http.Request) {
	
	price("ETH")
	//fmt.Println(data)
	json.NewEncoder(w).Encode(data)
	
})
*/