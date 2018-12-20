package main

import (
	"log"
	"net/http"
	"encoding/json"
	"strings"
    "net/url"
    "strconv"
	//"bytes"
 	"github.com/ethereum/go-ethereum/common"
	"github.com/gorilla/mux"
	//"os"
)

var IssueTokensByUSD = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request){ 
//func IssueTokensByUSD(w http.ResponseWriter, req *http.Request) {

	client := getClient()
	objToken := getTokenObject(client)
	objTokenTransaction := getTokenTransactorObject(client)
	auth := getAuthObject()

	params := mux.Vars(req)

	orderID = params["orderID"]		// encrypted
	userAddr = params["addr"]		// encrypted SJT Address
	value = params["USDValue"]
	emailID = params["userEmailID"]		// encrypted Email ID
	param:= orderID + "/" + userAddr + "/" + value  + "/" + emailID
	CallingLog("IssueTokensByUSD","IssueTokens.go",userAddr,emailID,param)
	log.Println("orderID =====> ", orderID)
	log.Println("userAddr =====> ", userAddr)
	log.Println("value =====> ", value)
	log.Println("emailID =====> ", emailID)
	

	respStatus := "Not OK"

	decryptOrderID := decrypt(orderID)
	decryptAddr = decrypt(userAddr)
	decryptEmailID = decrypt(emailID)
	userAddress = common.HexToAddress(decryptAddr)
	order, err := strconv.Atoi(decryptOrderID)
	if (err != nil) {
		log.Println(err)
		//errorlog
	}

	if (order <= 0){
		
		log.Println(order)
		throwException("IssueTokensByUSD", "IssueTokens.go", decryptAddr, decryptEmailID,"OrderId negative or zero")

	} else if (isUserValid(userAddress, decryptEmailID)) {
		
		var usdValueInCents = getUSDValueInCents(value)
		log.Println("usdValueInCents ====== ", usdValueInCents)

		var fundingTokensInWei = getFundingTokensForUSDPayments(usdValueInCents)
		log.Println("fundingTokensInWei =====================> ", fundingTokensInWei)

		successResult, err_result := objToken.CheckBalanceBeforeTransaction(nil, "IssueTokensByUSD", userAddress, fundingTokensInWei)
		if err_result!=nil {
			log.Println(err_result)
			//error log throw exception
		}
		
		if (successResult) {

			releaseTokens, err := objTokenTransaction.ReleaseTokensByUSD(auth, userAddress, fundingTokensInWei)
			if err != nil {
				log.Println("Failed to instantiate a Token contract: %v", err)
				throwException("IssueTokensByUSD", "IssueTokens.go", decryptAddr, decryptEmailID,"Fatal Error")				
			}
			log.Println("Release Token", releaseTokens)

			//****************************************************************************************************
			// Call PHP API To Update Value
			//**************************************************************************************************

			var fundingTokens = getTokenInSJTFromWei(fundingTokensInWei)
			log.Println("fundingTokens ==========> ", fundingTokens)

			SJTReceivedFromAddress := getTokenContractAddress()
			SJTReceivedFromAddress_string := SJTReceivedFromAddress.String()
			
			//encryptOrderID := encrypt(orderID)
			encryptPaidValue := encrypt(value)
			//encryptPrivateAddress := encrypt(userAddr)
			encryptSJTReceived := encrypt(fundingTokens)
			encryptFromAddress := encrypt(SJTReceivedFromAddress_string)

			log.Println("Order ID = ", orderID)
			log.Println("PaidValue = ", encryptPaidValue)
			log.Println("SJTRecieved = ", encryptSJTReceived)
			log.Println("fromAddress = ", encryptFromAddress)
			log.Println("senderAddress = ", userAddr)

			//apiUrl := "https://api.com"
			resource := "/api/purchaseapi.php/"
			data := url.Values{}
			data.Set("orderID", orderID)
			data.Add("SJTRecieved", encryptSJTReceived)
			data.Add("fromAddress", encryptFromAddress)
			data.Add("senderAddress", userAddr)
			data.Add("paidValue", encryptPaidValue)
		
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

			respStatus = resp.Status
			//**************************************************************************************************

			b, err := json.Marshal(respStatus)
			
			if err != nil {
				log.Println("Error: %s", err)
				throwException("IssueTokensByUSD", "IssueTokens.go", decryptAddr, decryptEmailID,"Fatal Error")
				return
			}
			
			var s5 string= string(b)
			ResultLog ("IssueTokensByUSD","IssueTokens.go",s5,param)
			json.NewEncoder(w).Encode(successResult)
			//json.NewEncoder(w).Encode(respStatus)
		} else {
			json.NewEncoder(w).Encode(successResult)			
			//throwException - "success = false" - balance check failed
		}
	} else {
		throwException("IssueTokensByUSD", "IssueTokens.go", decryptAddr, decryptEmailID,"Address Did Not Match with Email Id")
	}
})

// Code Modified 

// Issue Tokens From Reserved Balance By Admin
var IssueReservedTokens = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request){ 
//func IssueReservedTokens(w http.ResponseWriter, req *http.Request) {

	log.Println("IssueReservedTokens method is called !!!! ")
    
	client := getClient()
	objToken := getTokenObject(client)
	objTokenTransaction := getTokenTransactorObject(client)
	auth := getAuthObject()

	params := mux.Vars(req)
	orderID = params["orderID"]					// encrypted
	userAddr = params["addr"]					// encrypted
	value = params["tokens"]					
	receiverAddr := params["receiverAddr"]		// encrypted
	emailID = params["userEmailID"]				// encrypted
	param:= orderID + "/" + userAddr + "/" + value  + "/" + emailID + "/" + receiverAddr
	CallingLog("IssueReservedTokens","IssueTokens.go", decryptAddr, emailID,param)
	log.Println("orderID =====> ", orderID)
	log.Println("userAddr =====> ", userAddr)
	log.Println("receiverAddr =====> ", receiverAddr)
	log.Println("value =====> ", value)
	log.Println("emailID =====> ", emailID)
	
	
	respStatus := "Not OK"

	decryptOrderID = decrypt(orderID)
	decryptReceiverAddr := decrypt(receiverAddr)
	decryptAddr = decrypt(userAddr)
	decryptEmailID = decrypt(emailID)
	userAddress = common.HexToAddress(decryptAddr)
	receiverAddress := common.HexToAddress(decryptReceiverAddr)
	

	log.Println("decryptAddr =====> ", decryptAddr)
	log.Println("decryptEmailID =====> ", decryptEmailID)
	log.Println("decryptReceiverAddr =====> ", decryptReceiverAddr)
	
	order, err := strconv.Atoi(decryptOrderID)
	if (err != nil) {
		log.Println(err)
		//errorlog
	}

	if (order <= 0) {
		
		log.Println(order)
		throwException("IssueReservedTokens", "IssueTokens.go", decryptAddr, decryptEmailID,"OrderId negative or zero")
	
	} else if (isUserValid(userAddress,decryptEmailID)) {
		
		var totalSJTvalueInWei = getSJTValueInWei(value)

		successResult, err_result := objToken.CheckBalanceBeforeTransaction( nil, "IssueReservedTokens", receiverAddress, totalSJTvalueInWei)
		if err_result!=nil {
			log.Println(err_result)
			//error log throw exception
		}
		
		if (successResult) {

			releaseTokens, err := objTokenTransaction.ReleaseTokensFromReserved(auth, receiverAddress, totalSJTvalueInWei);
			if err != nil {
				log.Println("Failed to instantiate a Token contract: %v", err)
				throwException("IssueReservedTokens", "IssueTokens.go", decryptAddr, decryptEmailID,"Fatal Error")
				
			}

			log.Println("Release Token", releaseTokens)
			
			//*************************************************************************************************
			
			var fundingTokens = getTokenInSJTFromWei(totalSJTvalueInWei)
			log.Println("fundingTokens ==========> ", fundingTokens)

			//encryptOrderID := encrypt(orderID)
			encryptPaidValue := encrypt(value)
			//encryptPrivateAddress := encrypt(userAddr)
			encryptSJTReceived := encrypt(fundingTokens)
			//encryptFromAddress := encrypt(SJTReceivedFromAddress_string)
			log.Println("Order ID = ", orderID)
			log.Println("PaidValue = ", encryptPaidValue)
			log.Println("SJTRecieved = ", encryptSJTReceived)
			log.Println("fromAddress = ", userAddr)
			log.Println("senderAddress = ", receiverAddr)

			resource := "/api/purchaseapi.php/"
			data := url.Values{}
			data.Set("orderID", orderID)						// encrypted value received
			data.Add("SJTRecieved", encryptSJTReceived)			// encrypted
			data.Add("fromAddress", userAddr)					// encrypted value received
			data.Add("senderAddress", receiverAddr)				// encrypted value received
			data.Add("paidValue", encryptPaidValue)				// encrypted
		
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
			log.Println("resp ===========================> ", resp)
			
			respStatus = resp.Status
			b, err := json.Marshal(respStatus)
			if err != nil {
				log.Println("Error: %s", err)
				throwException("IssueReservedTokens", "IssueTokens.go", decryptAddr, decryptEmailID,"Fatal Error")
				return
			}
			var s5 string= string(b)
			ResultLog ("IssueReservedTokens","IssueTokens.go", s5,param)
			//json.NewEncoder(w).Encode(respStatus)
			json.NewEncoder(w).Encode(successResult)
		} else {
			json.NewEncoder(w).Encode(successResult)
		}
	} else {
		throwException("IssueReservedTokens", "IssueTokens.go", decryptAddr, decryptEmailID,"Address Did Not Match with Email Id")
	}		
})

var TransferSJTTokens = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request){ 
//func TransferSJTTokens(w http.ResponseWriter, req *http.Request) {	
	client := getClient()
	objToken := getTokenObject(client)
	objTokenTransaction := getTokenTransactorObject(client)
	auth := getAuthObject()

	params := mux.Vars(req)
	orderID = params["orderID"]						// encrypted
	userAddr = params["addr"]						// encrypted
	value = params["tokens"]
	receiverAddr := params["receiverAddr"]			// encrypted
	emailID = params["userEmailID"]					// encrypted
	param:= orderID + "/" + userAddr + "/" + value  + "/" + emailID + "/" + receiverAddr
	CallingLog("TransferSJTTokens","IssueTokens.go", decryptAddr, emailID,param)
	log.Println("orderID =====> ", orderID)
	log.Println("userAddr =====> ", userAddr)
	log.Println("receiverAddr =====> ", receiverAddr)
	log.Println("value =====> ", value)
	log.Println("emailID =====> ", emailID)

	respStatus := "Not OK"

	decryptOrderID = decrypt(orderID)
	decryptReceiverAddr := decrypt(receiverAddr)
	decryptAddr = decrypt(userAddr)
	decryptEmailID = decrypt(emailID)
	userAddress = common.HexToAddress(decryptAddr)
	receiverAddress := common.HexToAddress(decryptReceiverAddr)
	

	log.Println("userAddress =========> ", userAddress)
	log.Println("decryptEmailID =========> ", decryptEmailID)
	order, err := strconv.Atoi(decryptOrderID)
	if err != nil {
		log.Println(err)
		//error log
    }
	if (order <= 0) {
		
		log.Println(order)
		throwException("TransferSJTTokens", "IssueTokens.go", decryptAddr, decryptEmailID,"OrderId negative or zero")
	
	} else if (isUserValid(userAddress,decryptEmailID)) {
		
		var totalSJTvalueInWei = getSJTValueInWei(value)
	
		successResult, err_result := objToken.CheckBalanceForTransferTokens(nil, userAddress, receiverAddress, totalSJTvalueInWei)
		if err_result!=nil {
			log.Println(err_result)
			//error log throw exception
		}
		
		if (successResult) {
			transferTokens, err := objTokenTransaction.TransferTokens(auth, userAddress, receiverAddress, totalSJTvalueInWei);
			if err != nil {
				throwException("TransferSJTTokens", "IssueTokens.go", decryptAddr, decryptEmailID,"Fatal Error")
			}

			log.Println("Release Token", transferTokens)
			
			//****************************************************************************************************

			var fundingTokens = getTokenInSJTFromWei(totalSJTvalueInWei)
			log.Println("fundingTokens ==========> ", fundingTokens)

			//encryptOrderID := encrypt(orderID)
			encryptPaidValue := encrypt(value)
			//encryptToAddress := encrypt(receiverAddr)
			encryptSJTReceived := encrypt(fundingTokens)
			//encryptFromAddress := encrypt(SJTReceivedFromAddress_string)
			log.Println("Order ID = ", orderID)
			log.Println("PaidValue = ", encryptPaidValue)
			log.Println("SJTRecieved = ", encryptSJTReceived)
			log.Println("fromAddress = ", userAddr)
			log.Println("senderAddress = ", receiverAddr)

			resource := "/api/purchaseapi.php/"
			data := url.Values{}
			data.Set("orderID", orderID)						// encrypted value received
			data.Add("SJTRecieved", encryptSJTReceived)			// encrypted 
			data.Add("fromAddress", userAddr)					// encrypted value received
			data.Add("senderAddress", receiverAddr)				// encrypted value received
			data.Add("paidValue", encryptPaidValue)				// encrypted
		
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
			log.Println("resp ===========================> ", resp)		
			respStatus = resp.Status
			//*****************************************************************************************************
			
			b, err := json.Marshal(respStatus)
			if err != nil {
				log.Println("Error: %s", err)
				throwException("TransferSJTTokens", "IssueTokens.go", decryptAddr, decryptEmailID, "Fatal Error")
				return
			}
			var s5 string= string(b)
			ResultLog ("TransferSJTTokens","IssueTokens.go",s5,param)
			log.Println(respStatus)
			
			json.NewEncoder(w).Encode(successResult)			
			//json.NewEncoder(w).Encode(respStatus)
		} else {
			json.NewEncoder(w).Encode(successResult)			
		}
	} else {
		throwException("TransferSJTTokens", "IssueTokens.go", decryptAddr, decryptEmailID, "Address Did Not Match with Email Id")
	}
})

var MintTokens = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request){
//func MintTokens(w http.ResponseWriter, req *http.Request) {
    
    client := getClient()
	objTokenTransaction := getTokenTransactorObject(client)
	auth := getAuthObject()

	params := mux.Vars(req)
	userAddr = params["addr"]
	value = params["value"]
	emailID = params["userEmailID"]
	param:= userAddr + "/" + value + "/" + emailID
	CallingLog("MintTokens","IssueTokens.go", decryptAddr, decryptEmailID,param)
	decryptAddr = decrypt(userAddr)
	decryptEmailID = decrypt(emailID)

	userAddress = common.HexToAddress(decryptAddr)
	
	if (isUserValid(userAddress,decryptEmailID)) {
		
		var SJTValue = getSJTValueInWei(value)
	
		mintTokens, err := objTokenTransaction.MintToken(auth, SJTValue)
		if err != nil {
			log.Println("Failed to instantiate a Token contract: %v", err)
		}
	
		b, err := json.Marshal(mintTokens)
		if err != nil {
			log.Println("Error: %s", err)
			throwException("MintTokens", "IssueTokens.go", decryptAddr, decryptEmailID, "Fatal Error")
			return
		}
		var s5 string= string(b)
		ResultLog ("MintTokens","IssueTokens.go", s5,param)
	
		json.NewEncoder(w).Encode(mintTokens)
	} else {
		throwException("MintTokens", "IssueTokens.go", decryptAddr, decryptEmailID, "Fatal Error")
	}
})

var MintICOSupply = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request){
//func MintICOSupply(w http.ResponseWriter, req *http.Request) {
    
	client := getClient()
	objTokenTransaction := getTokenTransactorObject(client)
	auth := getAuthObject()

	params := mux.Vars(req)
	userAddr = params["addr"]
	value = params["value"]
	emailID = params["userEmailID"]
	param:= userAddr + "/" + value + "/" + emailID
	CallingLog("MintICOSupply","IssueTokens.go", decryptAddr, decryptEmailID,param)
	decryptAddr = decrypt(userAddr)
	decryptEmailID = decrypt(emailID)

	userAddress = common.HexToAddress(decryptAddr)
	
	if (isUserValid(userAddress,decryptEmailID)) {
		
		var SJTValue = getSJTValueInWei(value)
	
		mintTokens, err := objTokenTransaction.MintICOSupply(auth, SJTValue)
		if err != nil {
			log.Println("Failed to instantiate a Token contract: %v", err)
			throwException("MintICOSupply", "IssueTokens.go", decryptAddr, decryptEmailID,"Fatal Error")
		}
	
		b, err := json.Marshal(mintTokens)
		if err != nil {
			log.Println("Error: %s", err)
			throwException("MintICOSupply", "IssueTokens.go", decryptAddr, decryptEmailID,"Fatal Error")
			return
		}

		var s5 string=string(b)
		ResultLog ("MintICOSupply","IssueTokens.go", s5,param)
		
		json.NewEncoder(w).Encode(mintTokens)
	} else {
		throwException("MintICOSupply", "IssueTokens.go", decryptAddr, decryptEmailID,"Fatal Error")
	}
})

var SetBuyPrice = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request){
//func SetBuyPrice(w http.ResponseWriter, req *http.Request) {
    
    client := getClient()
	objTokenTransaction := getTokenTransactorObject(client)
	auth := getAuthObject()

	params := mux.Vars(req)
	
	userAddr = params["addr"]
	emailID = params["userEmailID"]
	value = params["buyPrice"]
	param:= userAddr + "/" + emailID + "/" + value
	CallingLog("SetBuyPrice","IssueTokens.go", decryptAddr, decryptEmailID,param)
	decryptAddr = decrypt(userAddr)
	decryptEmailID = decrypt(emailID)

	userAddress = common.HexToAddress(decryptAddr)
	
	if (isUserValid(userAddress,decryptEmailID)) {
		
		var usdValueInCents = getUSDValueInCents(value)
		setBuyPrice, err := objTokenTransaction.SetBuyPrice(auth, usdValueInCents)
		if err != nil {
			log.Println("Failed to instantiate a Token contract: %v", err)
			throwException("SetBuyPrice", "IssueTokens.go", decryptAddr, decryptEmailID,"Fatal Error")
		}
		b, err := json.Marshal(setBuyPrice)
		
			if err != nil {
				log.Println("Error: %s", err)
		throwException("SetBuyPrice", "IssueTokens.go", decryptAddr, decryptEmailID,"Fatal Error")
				
		
				return
				
		
			}
		
			var s5 string= string(b)
			
		ResultLog ("SetBuyPrice","IssueTokens.go", s5,param)
		
		json.NewEncoder(w).Encode(setBuyPrice)
	} else {
		throwException("SetBuyPrice", "IssueTokens.go", decryptAddr, decryptEmailID,"Fatal Error")
	}
})

var SetSellPrice = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request){
//func SetSellPrice(w http.ResponseWriter, req *http.Request) {
    
    client := getClient()
	objTokenTransaction := getTokenTransactorObject(client)
	auth := getAuthObject()

	params := mux.Vars(req)
	userAddr = params["addr"]
	emailID = params["userEmailID"]
	value = params["sellPrice"]
	param:= userAddr + "/" + emailID + "/" + value
	CallingLog("SetSellPrice","IssueTokens.go", decryptAddr, decryptEmailID,param)
	decryptAddr = decrypt(userAddr)
	decryptEmailID = decrypt(emailID)

	userAddress = common.HexToAddress(decryptAddr)
	

	if (isUserValid(userAddress,decryptEmailID)) {
		
		var usdValueInCents = getUSDValueInCents(value)
	
		setSellPrice, err := objTokenTransaction.SetSellPrice(auth, usdValueInCents)
		if err != nil {
			log.Println("Failed to instantiate a Token contract: %v", err)
			throwException("SetSellPrice", "IssueTokens.go", decryptAddr, decryptEmailID,"Fatal Error")
		}
		b, err := json.Marshal(setSellPrice)
		
			if err != nil {
				log.Println("Error: %s", err)
				throwException("SetSellPrice", "IssueTokens.go", decryptAddr, decryptEmailID,"Fatal Error")
				
		
				return
		
			}
		
			var s5 string= string(b)
		
		ResultLog ("SetSellPrice","IssueTokens.go", s5, param)
		json.NewEncoder(w).Encode(setSellPrice)
	} else {
		throwException("SetSellPrice", "IssueTokens.go", decryptAddr, decryptEmailID,"Fatal Error")
	}

})

var SetPublicAddressWithPrivateAddress = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request){
//func SetBuyPrice(w http.ResponseWriter, req *http.Request) {
	
	client := getClient()
	objTokenTransaction := getTokenTransactorObject(client)
	auth := getAuthObject()

	params := mux.Vars(req)
	userAddr = params["addr"] 		//private address
	emailID = params["userEmailID"]
	publicAddr = params["publicAddress"]
	param:= userAddr + "/" + emailID + "/" + publicAddr
	CallingLog("SetPublicAddressWithPrivateAddress","IssueTokens.go", decryptAddr, decryptEmailID,param)
	log.Println("userAddr =============> ", userAddr)
	log.Println("emailID =============> ", emailID)
	log.Println("publicAddr =============> ", publicAddr)

	decryptAddr = decrypt(userAddr)
	decryptEmailID = decrypt(emailID)
	decryptPublicAddr = decrypt(publicAddr)
	
	userAddress = common.HexToAddress(decryptAddr) 		//sjt address
	publicAddress = common.HexToAddress(decryptPublicAddr) //public address
	
	if (isUserValid(userAddress,decryptEmailID)) {
		
		setPublicWithPrivate, err := objTokenTransaction.SetPrivateAddByPublicAdd(auth, publicAddress, userAddress)
		if err != nil {
			throwException("SetPublicAddressWithPrivateAddress", "IssueTokens.go", decryptAddr, decryptEmailID,"Fatal Error")
		}
		log.Println("setPublicWithPrivate= ", setPublicWithPrivate)
		b, err := json.Marshal(setPublicWithPrivate)
		
			if err != nil {
				log.Println("Error: %s", err)
				throwException("SetPublicAddressWithPrivateAddress", "IssueTokens.go", decryptAddr, decryptEmailID,"Fatal Error")
				
		
				return
		
			}
		
			var s5 string= string(b)
		
		ResultLog ("SetPublicAddressWithPrivateAddress","IssueTokens.go", s5, param)
		json.NewEncoder(w).Encode(setPublicWithPrivate)
	} else {
		throwException("SetPublicAddressWithPrivateAddress", "IssueTokens.go", decryptAddr, decryptEmailID,"Fatal Error")
	}
})

var SetEmailWithPrivateAddress = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request){
	
	client := getClient()
	objToken := getTokenObject(client)
	auth := getAuthObject()

	params := mux.Vars(req)
	userAddr = params["privateAddress"] 		//private address
	emailID = params["emailId"]
	param:= userAddr + "/" + emailID
	CallingLog("SetEmailWithPrivateAddress","IssueTokens.go", decryptAddr, decryptEmailID,param)
	log.Println("userAddr =============> ", userAddr)
	log.Println("emailID =============> ", emailID)

	Success := false
	decryptAddr = decrypt(userAddr)
	decryptEmailID = decrypt(emailID)
	
	userAddress = common.HexToAddress(decryptAddr) 		//sjt address
	
	result, err := objToken.MapEmailIDWithPrivateAdd(auth, decryptEmailID, userAddress)
	if err != nil {
		throwException("SetEmailWithPrivateAddress", "IssueTokens.go", decryptAddr, decryptEmailID,"Fatal Error")
	}
	log.Println("result = ", result)
	Success = true
	log.Println("Success= ", Success)
	b, err := json.Marshal(result)
	
		if err != nil {
			log.Println("Error: %s", err)
			throwException("SetEmailWithPrivateAddress", "IssueTokens.go", decryptAddr, decryptEmailID,"Fatal Error")
			
			return
	
		}
	
		var s5 string= string(b)
		
	ResultLog ("SetEmailWithPrivateAddress","IssueTokens.go", s5,param)

	json.NewEncoder(w).Encode(result)
})