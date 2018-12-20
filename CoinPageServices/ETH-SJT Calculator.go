package main

import (
	"log"
	"net/http"
	"encoding/json"
//"bytes"
	"github.com/gorilla/mux"
)

type EthSJTInfo struct {
	ETH string
	SJT string
}

var getEthSJTInfo EthSJTInfo

var EthSJTConvertion = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request){  
	
	//log.Println("=================== API CALLED SUCCESSFULLY !!!!!!!!!!!!! =================")

	params := mux.Vars(req)
	var ETHParams string

	ETHParams = params["eth"]
	//log.Println("ETH PARAMS =========>", ETHParams)
	//ETHinWei := getEthValueInWei(ETHParams)
	//ETHinWei_String := ETHinWei.String()
	param:= ETHParams
	CallingLog("EthSJTConvertion", "ETH-SJTCalculator.go", "NULL", "NULL",param)

	//log.Println("====================get Token price info  ==================")
	
	_, SJTinOneETH := getTokenPriceInfo()
	//log.Println("========================== getTokenPriceInfo ==================")
	
	SJTinOneETH_string := SJTinOneETH.String()
//	log.Println("========================== SJTinOneETH ==================")
	
	SJTinOneETHinWei := getSJTValueInWei(SJTinOneETH_string)
//	log.Println("========================== getSJTValueInWei ==================")
	
	SJTinOneETHinWei_String := SJTinOneETHinWei.String()
//	log.Println("=======================================ETH-SJT Calculator.go==================")

	totalSJTinWei := doMultiplication(SJTinOneETHinWei_String, ETHParams, "1")
	totalSJT := getTokenInSJTFromWei(totalSJTinWei)

    getEthSJTInfo := new(EthSJTInfo)
    getEthSJTInfo.ETH = ETHParams
	getEthSJTInfo.SJT = totalSJT
	
//	log.Println("getEthSJTInfo =====> ", getEthSJTInfo)
	b, err := json.Marshal(getEthSJTInfo)
	
		if err != nil {
			throwException("EthSJTConvertion", "ETH-SJTCalculator.go", "NULL", "NULL","Fatal Error")
			log.Println("Error: %s", err)
	
			return
	
		}
	
		var s5 string=string(b)
		
	ResultLog ("EthSJTConvertion", "ETH-SJTCalculator.go", s5, param)
	json.NewEncoder(w).Encode(getEthSJTInfo)  
	
})