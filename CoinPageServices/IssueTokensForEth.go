package main

import (
	"math/big"
	"log"
	"net/http"
	"encoding/json"
	"strings"
    "net/url"
    "strconv"
	//"bytes"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gorilla/mux"
)

type SuccessValue struct {
	Success bool
}

var getEthTransferResult SuccessValue

//var IssueTokensByETH = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request){
func IssueTokensByETH(w http.ResponseWriter, req *http.Request) {

	client := getClient()
	objToken := getTokenTransactorObject(client)
	auth := getAuthObject()

	getEthTransferResult := new(SuccessValue)
	params := mux.Vars(req)
	var addr string	 
	var value string
	var SJTBridgeContractAddressParams string
	var mySecurityKeyParams string
	var totalTokensIssued *big.Int

	getEthTransferResult.Success = true
 
	orderID = "-1"
	
	SJTBridgeContractAddressParams = params["contractAddress"]
	log.Println("SJTBridgeContractAddressParams ============>>>>> ", SJTBridgeContractAddressParams)
	
	mySecurityKeyParams = params["key"]
	log.Println("mySecurityKeyParams ===================> ", mySecurityKeyParams)
	mySecurityKey_decrpyted := decrypt(mySecurityKeyParams)
	log.Println("mySecurityKey_decrpyted ===================> ", mySecurityKey_decrpyted)
	
	
	addr = params["addr"] 		// public address from which user will send the ETH from.
	value = params["Eth"]
	param:= SJTBridgeContractAddressParams + "/" + mySecurityKeyParams + "/" + mySecurityKey_decrpyted + "/" + addr + "/" + value
CallingLog("IssueTokensByETH()","IssueTokensForEth.go",addr,"Null",param)

	if (mySecurityKey_decrpyted != mySecretKeyForEthTxn){

		getEthTransferResult.Success = false
		log.Println("Security Issue !! Wrong Secret Key ! Not an authenticated user !")

	} else if (SJTBridgeContractAddressParams != SJTBridgeContractAddress){
		
		getEthTransferResult.Success = false
		log.Println("Security Issue !! Wrong Contract Address ! Not an authenticated user !")
	
	} else {
	
		ETHval := new(big.Int)
		ETHval.SetString(value, 10)
		ETHval_string := getEthValueFromWei(ETHval)		

		senderPublicAddress = common.HexToAddress(addr)	

		totalTokensIssued = getTokenValues(ETHval_string) 
		log.Println("totalTokensIssued = ", totalTokensIssued)		

		privateAddr := getPvtAddrByPubAddr(senderPublicAddress)
		privateAddr_string := privateAddr.String()
		

		if (privateAddr != nullAddr) {
			senderPrivateAddress = common.Address(privateAddr)
			releaseTokens, err := objToken.ReleaseTokensForICO(auth, privateAddr, totalTokensIssued);
			if err != nil {
				log.Println("Failed to instantiate a Token contract: %v", err)
				throwException("IssueTokensByETH()", "IssueTokensForEth.go", addr, "NULL","Fatal Error")
				getEthTransferResult.Success = false
			}
			log.Println(releaseTokens)
			
			//*************************************************************************************************
			// Call API

			var fundingTokens = getTokenInSJTFromWei(totalTokensIssued)
			log.Println("fundingTokens ==========> ", fundingTokens)

			//SJTReceivedFromAddress := getTokenContractAddress()						// need to modify to SJT Contract Address!!!
			//SJTReceivedFromAddress_string := SJTReceivedFromAddress.String()
			SJTReceivedFromAddress := SJTBridgeContractAddress_full

			encryptOrderID := encrypt(orderID)
			encryptPublicAddress := encrypt(addr)
			encryptPaidValue := encrypt(ETHval_string)
			encryptPrivateAddress := encrypt(privateAddr_string)
			encryptFromAddress := encrypt(SJTReceivedFromAddress)
			encryptSJTReceived := encrypt(fundingTokens)
			log.Println("Order ID = ", encryptOrderID)
			log.Println("Public Address = ", encryptPublicAddress)
			log.Println("Paid Value = ", encryptPaidValue)
			log.Println("Private Address = ", encryptPrivateAddress)
			log.Println("From Address = ", encryptFromAddress)
			log.Println("SJT Received = ", encryptSJTReceived)

			//apiUrl := "https://api.com"
			resource := "/api/purchaseapi.php/"
			data := url.Values{}
			data.Set("orderID", encryptOrderID)
			data.Add("SJTRecieved", encryptSJTReceived)
			data.Add("fromAddress", encryptFromAddress)
			data.Add("senderAddress", encryptPrivateAddress)
			data.Add("paidValue", encryptPaidValue)
			data.Add("publicAddress", encryptPublicAddress)
		
			u, _ := url.ParseRequestURI(apiToUpdateDB)
			u.Path = resource
			urlStr := u.String() // "https://api.com/user/"
			log.Println("urlStr ==========> ", urlStr)

			client := &http.Client{}
			r, _ := http.NewRequest("POST", urlStr, strings.NewReader(data.Encode())) // <-- URL-encoded payload
			r.Header.Add("Authorization", "auth_token=\"XXXXXXX\"")
			r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
			r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
		
			resp, _ := client.Do(r)
			log.Println("STATUS ===========================> ", resp.Status)
			log.Println("Value ===========================> ", resp)
			
			//*************************************************************************************************
		} else {
			getEthTransferResult.Success = false
		}
	}
	log.Println("Success = ", getEthTransferResult)
	b, err := json.Marshal(getEthTransferResult)
	
		if err != nil {
	
			log.Println("Error: %s", err)
	throwException("IssueTokensByETH()", "IssueTokensForEth.go", decryptAddr, decryptEmailID,"Fatal Error")
		return
	
		}
	
		var s5 string= string(b)
		
	ResultLog ("IssueTokensByETH()","IssueTokensForEth.go", s5,param)
	json.NewEncoder(w).Encode(getEthTransferResult)
	
}

func getTokenValues(valueInEth string ) (FundingTokens *big.Int) {
	CallingLog("getTokenValues()","IssueTokensForEth.go", "NULL", "NULL","NULL")
	var CurrentEthPriceInUSD *big.Int
	var division_result string

	CurrentEthPriceInUSD = price("ETH")										// 330 usd
	//log.Println("CurrentEthPriceInUSD = ", CurrentEthPriceInUSD)			
	
	CurrentEthPriceInUSD_String := CurrentEthPriceInUSD.String()
	//log.Println("CurrentEthPriceInUSD_String = ", CurrentEthPriceInUSD_String)
	
	TokenValueInCents_string := getSellPrice()								// 100 cents
	//log.Println("TokenValueInCents_string = ", TokenValueInCents_string)	

	division_result = doDivision(SJTInWei, TokenValueInCents_string)
	//log.Println("division_result : ",division_result)
	
	FundingTokens = doMultiplication(division_result, CurrentEthPriceInUSD_String, valueInEth)
	//log.Println("funding token in SJT wei: ",FundingTokens)
	//CrLog ("getTokenValues","IssueTokensForEth.go", "admin", " ")
	//bigstr :=(FundingTokens)
	//bigint := big.NewInt(FundingTokens)
	//bigstr := bigint.String()
	ResultLog ("getTokenValues()","IssueTokensForEth.go", "NULL","NULL")
	return
}