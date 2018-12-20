package main

import (
    "encoding/json"
    "net/http"
	//"math/big"
	"log"
	//"bytes"
)

type ICOInfo struct {
	Supply string
	RemainingBalance string
}

type SJTReservedCoinInfo struct {
	Supply string
	RemainingBalance string
}

var getICOInfo ICOInfo
var getSJTReservedInfo SJTReservedCoinInfo

//func GetICOInfo(w http.ResponseWriter, req *http.Request) {
var GetICOInfo = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){   
	CallingLog("GetICOInfo","SJTReservedAndICOInfo.go", "NULL", "NULL","NULL")
	var ICOSupplyInWei = getICOSupplyHelper()
	var ICOSupplyInSJT = getTokenInSJTFromWei(ICOSupplyInWei)
	var ICORemainingBalanceInWei = getICORemainingBalHelper()
	var ICORemainingBalanceInSJT = getTokenInSJTFromWei(ICORemainingBalanceInWei)
	

    getICOInfo := new(ICOInfo)
    getICOInfo.Supply = ICOSupplyInSJT
    getICOInfo.RemainingBalance = ICORemainingBalanceInSJT
	b, err := json.Marshal(getSJTReservedInfo)
	
		if err != nil {
	
			log.Println("Error: %s", err)
			throwException("GetICOInfo", "SJTReservedAndICOInfo.go", "NULL", "NULL","Fatal Error")
			return
	
		}
	
		var s5 string = string(b)
		
	ResultLog ("GetICOInfo","SJTReservedAndICOInfo.go", s5, "NULL")
	json.NewEncoder(w).Encode(getICOInfo)  
	
})

//func GetSJTReservedInfo(w http.ResponseWriter, req *http.Request) {
var GetSJTReservedInfo = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){   
	CallingLog("GetSJTReservedInfo","SJTReservedAndICOInfo.go", "NULL", "NULL","NULL")
	var SJTReservedSupplyInWei = getSJTReservedSupplyHelper()
	var SJTReservedSupplyInSJT = getTokenInSJTFromWei(SJTReservedSupplyInWei)
	var SJTReservedRemainingBalanceInWei = getSJTReservedRemainingBalHelper()
	var SJTReservedRemainingBalanceInSJT = getTokenInSJTFromWei(SJTReservedRemainingBalanceInWei)
	
	getSJTReservedInfo := new(SJTReservedCoinInfo)
	getSJTReservedInfo.Supply = SJTReservedSupplyInSJT
	getSJTReservedInfo.RemainingBalance = SJTReservedRemainingBalanceInSJT
	b, err := json.Marshal(getSJTReservedInfo)
	
		if err != nil {
	
			log.Println("Error: %s", err)
			throwException("GetSJTReservedInfo", "SJTReservedAndICOInfo.go", "NULL", "NULL","Fatal Error")
			return
		}
	
		var s5 string = string(b)
		
	ResultLog ("GetSJTReservedInfo","SJTReservedAndICOInfo.go", s5,"NULL")	
	json.NewEncoder(w).Encode(getSJTReservedInfo)  
	
})