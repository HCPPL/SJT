package main

import (
	"net/http"
	"encoding/json"
	"log"

	"github.com/gorilla/mux"
)

type encryptDecrypt struct {
	Result string
}

var getEncryptDecrypt encryptDecrypt

var EncryptParams = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request){ 

	params := mux.Vars(req)
	encryptText = params["EncryptText"]
	param:= encryptText
	CallingLog("EncryptParams","EncryptDecrypt.go", "NULL", "NULL",param)
	textEncrypted := encrypt(encryptText)

	getEncryptDecrypt := new(encryptDecrypt)
	getEncryptDecrypt.Result = textEncrypted
	b, err := json.Marshal(getEncryptDecrypt)
	
		if err != nil {
			throwException("EncryptParams", "EncryptDecrypt.go", "NULL", "NULL","Fatal Error")
			log.Println("Error: %s", err)
	
			return
	
		}
		var s5 string= string(b)
		
	ResultLog ("EncryptParams","EncryptDecrypt.go", s5, param)

	json.NewEncoder(w).Encode(getEncryptDecrypt)
	
})

var DecryptParams = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request){ 

	params := mux.Vars(req)
	decryptText = params["DecryptText"]
	param:= decryptText
	CallingLog("DecryptParams", "EncryptDecrypt.go", "NULL", "NULL" ,param)
	textDecrypted := decrypt(decryptText)

	getEncryptDecrypt := new(encryptDecrypt)
	getEncryptDecrypt.Result = textDecrypted
	b, err := json.Marshal(getEncryptDecrypt)
	
		if err != nil {
			throwException("DecryptParams", "EncryptDecrypt.go", "NULL", "NULL","Fatal Error")
			log.Println("Error: %s", err)
	
			return
	
		}
		var s5 string= string(b)
		
	ResultLog ("DecryptParams", "EncryptDecrypt.go", s5, param)
		
	json.NewEncoder(w).Encode(getEncryptDecrypt)
	
})