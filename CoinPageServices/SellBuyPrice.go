package main

import (
    "encoding/json"
    "net/http"
    "log"
    //"bytes"
)
type SellBuyInfo struct {
	SellPrice string
	BuyPrice string
}

var getSellBuyInfo SellBuyInfo

//func GetSellPriceBuyPriceInfo(w http.ResponseWriter, req *http.Request) {
var GetSellPriceBuyPriceInfo = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){  
CallingLog("GetSellPriceBuyPriceInfo", "SellBuyPrice.go","NULL","NULL","NULL")
    var SellPriceInCents = getSellPriceHelper()
    log.Println("SellPriceInCents ================> ", SellPriceInCents)
    var SellPriceInUSD = getUsdValueFromCents(SellPriceInCents)
    log.Println("SellPriceInUSD ================> ", SellPriceInUSD)
    
    var BuyPriceInCents = getBuyPriceHelper()
    var BuyPriceInUSD = getUsdValueFromCents(BuyPriceInCents)
   // pu:= SellPriceInCents + SellPriceInUSD
   

    getSellBuyInfo := new(SellBuyInfo)
    getSellBuyInfo.SellPrice = SellPriceInUSD
    getSellBuyInfo.BuyPrice = BuyPriceInUSD

    log.Println("getSellBuyInfo = ", getSellBuyInfo)
    b, err := json.Marshal(getSellBuyInfo)
	
		if err != nil {
	
            log.Println("Error: %s", err)
            throwException("GetSellPriceBuyPriceInfo", "SellBuyPrice.go", "NULL", "NULL","Fatal Error")
	return
			
	
		}
        var s5 string= string(b)
		
    ResultLog ("GetSellPriceBuyPriceInfo", "SellBuyPrice.go" ,s5,"NULL")
    json.NewEncoder(w).Encode(getSellBuyInfo)  
    
})