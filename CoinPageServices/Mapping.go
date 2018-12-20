package main

import (
    "encoding/json"
    "net/http"
	"fmt"
	"log"
	//"bytes"
    "github.com/ethereum/go-ethereum/common"
    "github.com/gorilla/mux"
)

type EmailPrivateMapping struct {
	EmailId string
	SJTAddress common.Address
}

type PublicPrivateMapping struct {
	PublicAddress common.Address
	SJTAddress common.Address
}

var getPrivateByEmail EmailPrivateMapping
var getPrivateAddrByPublicAddr PublicPrivateMapping

var GetPvtByEmail = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){   
	
	params := mux.Vars(r)
		
	var e = params["emailAddress"]
		
	decryptEmail := decrypt(e)
	
		param:= e
		CallingLog("GetPvtByEmail", "Mapping.go", "Null", decryptEmail,param)	
	fmt.Println("decryptEmail = ", decryptEmail)
	
	getPrivateByEmail := new(EmailPrivateMapping)
	getPrivateByEmail.EmailId =  decryptEmail
	getPrivateByEmail.SJTAddress = getPrivateAddressByEmailID(decryptEmail)
		
	fmt.Println("getPrivateByEmail = ", getPrivateByEmail)
	b, err := json.Marshal(getPrivateByEmail)
	
		if err != nil {
	
			log.Println("Error: %s", err)
			throwException("GetPvtByEmail", "Mapping.go", "NULL", "NULL","Fatal Error")
			
			return
	
		}
	
		var s5 string= string(b)
		
	ResultLog ("GetPvtByEmail", "Mapping.go", s5,param)
	json.NewEncoder(w).Encode(getPrivateByEmail)  
	
})

var GetPrivateAddressByPublicAddress = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){   
	
	params := mux.Vars(r)
	var publicAdd = params["publicAddress"]
	decryptPublicAdd := decrypt(publicAdd)
	param:= publicAdd	
	CallingLog("GetPrivateAddressByPublicAddress","Mapping.go", decryptPublicAdd, "NULL",param)
	PublicAddress := common.HexToAddress(decryptPublicAdd)
	

	fmt.Println("decryptPublicAdd = ", decryptPublicAdd)
	
	getPrivateAddrByPublicAddr := new(PublicPrivateMapping)
	getPrivateAddrByPublicAddr.PublicAddress =  PublicAddress
	getPrivateAddrByPublicAddr.SJTAddress = getPvtAddrByPubAddr(PublicAddress)
			
	fmt.Println("getPrivateByPublic = ", getPrivateAddrByPublicAddr)
	b, err := json.Marshal(getPrivateAddrByPublicAddr)
	
	if err != nil {
		
				log.Println("Error: %s", err)
				throwException("GetPrivateAddressByPublicAddress", "Mapping.go", decryptPublicAdd, "NULL","Fatal Error")
				
				return
		
			}
	
		var s5 string= string(b)
		
	ResultLog ("GetPrivateAddressByPublicAddress","Mapping.go", s5, param)
	
	json.NewEncoder(w).Encode(getPrivateAddrByPublicAddr)  

})