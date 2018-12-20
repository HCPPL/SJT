package main

import (
    "encoding/json"
    "net/http"
	"log"
	//"bytes"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gorilla/mux"
	
)

type SJTBalanceInfo struct {
	SJTAddress common.Address
	SJTBalance string
}

var getSJTBalanceInfo SJTBalanceInfo

var GetSJTAddressBalanceInfo = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request){  
	client := getClient()
	objToken := getTokenObject(client)

	params := mux.Vars(req)

	var SJTAddressParams string

	SJTAddressParams = params["addr"]
	SJTAddress_decrypt := decrypt(SJTAddressParams)
	SJTAddressValue := getTokenAddress(SJTAddress_decrypt)
	param:= SJTAddressParams
CallingLog("GetSJTAddressBalanceInfo","GetBalance.go", "NULL", "NULL",param)
	balance, err := objToken.BalanceOf(nil, SJTAddressValue)
	if err != nil {
		throwException("GetSJTAddressBalanceInfo", "GetBalance.go", "NULL", "NULL","Fatal Error")
	}

	balanceInSJT := getTokenInSJTFromWei(balance)

    getSJTBalanceInfo := new(SJTBalanceInfo)
    getSJTBalanceInfo.SJTAddress = SJTAddressValue
    getSJTBalanceInfo.SJTBalance = balanceInSJT
	
	b, err := json.Marshal(getSJTBalanceInfo)
	
		if err != nil {
			throwException("GetSJTAddressBalanceInfo", "GetBalance.go", "NULL", "NULL","Fatal Error")
			log.Println("Error: %s", err)
	
			return
	
		}
	
		var s5 string= string(b)
		log.Println(string(s5))
		ResultLog ("GetSJTAddressBalanceInfo","GetBalance.go", s5, param)
	json.NewEncoder(w).Encode(getSJTBalanceInfo) 
	
	 
})