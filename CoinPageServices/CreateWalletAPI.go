package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"log"
	//"bytes"
	//"reflect"
	
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/gorilla/mux"
)
/************************************************************************************************************* */
// HANDLER FUNCTIONS 
/************************************************************************************************************* */

var CreateAccount = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
//func CreateAccount(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
		
	passphrase = params["passphrase"]
	emailID = params["userEmailID"]
 param:=passphrase + "/" + emailID
 CallingLog("CreateAccount","CreateWalletAPI.go","NULL",emailID,param)
	log.Println("passphrase =========> ", passphrase )
	log.Println("emailID =========> ", emailID )

	var result accounts.Account =  createAccount(passphrase, emailID)
	b, err := json.Marshal(result)
	
		if err != nil {
			throwException("CreateAccount()", "CreateWalletAPI.go", "NULL", emailID,"Fatal Error")
			log.Println("Error: %s", err)
	
			return
	
		}
		var s5 string= string(b)
	
	ResultLog ("CreateAccount","CreateWalletAPI.go", s5, param)
	json.NewEncoder(w).Encode(result)
	
	
})

var UpdateAccount = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
//func UpdateAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	
	userAddr = params["addr"]				//addr = "0x247b9023a2c8fdab874accbc4a81d85a21c46886"
	passphrase = params["passphrase"]
	newPassphrase = params["newPassphrase"]
	emailID = params["userEmailID"]
	
	param:= userAddr + "/" + passphrase + "/" + newPassphrase + "/" + emailID
	CallingLog("UpdateAccount","CreateWalletAPI.go",userAddr,emailID,param)
	log.Println("userAddr =====================>", userAddr)
	log.Println("passphrase =====================>", passphrase)
	log.Println("emailID =====================>", emailID)
	log.Println("newPassphrase =====================>", newPassphrase)
	
    var result error =  changePassphrase(userAddr, passphrase, newPassphrase, emailID)
	b, err := json.Marshal(result)
	
		if err != nil {
			throwException("UpdateAccount()", "CreateWalletAPI.go", userAddr, emailID,"Fatal Error")
			log.Println("Error: %s", err)
	
			return
	
		}
	
		var s5 string= string(b)
		
	ResultLog("UpdateAccount","CreateWalletAPI.go", s5, param)
	json.NewEncoder(w).Encode(result)
	
})

var UnlockAccount = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
//func UnlockAccount(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)

	userAddr = params["addr"]			//addr = "0x247b9023a2c8fdab874accbc4a81d85a21c46886"
	passphrase = params["passphrase"]
	emailID = params["userEmailID"]
	param:= userAddr + "/" + passphrase + "/" + emailID
	CallingLog("UnlockAccount","CreateWalletAPI.go", userAddr, emailID,param)
    var result error =  unlockAccount(userAddr, passphrase, emailID)
	b, err := json.Marshal(result)
	
		if err != nil {
			throwException("UnlockAccount()", "CreateWalletAPI.go", userAddr, emailID,"Fatal Error")
			log.Println("Error: %s", err)
	
			return
	
		}
	
		var s5 string= string(b)
	
	ResultLog ("UnlockAccount","CreateWalletAPI.go", s5,param)
	json.NewEncoder(w).Encode(result)
	
})

/**************************************************************************************************************** */
// METHODS
/*************************************************************************************************************** */

/**************** CREATE NEW ACCOUNT *************************** */
func createAccount(pwd string, email string)  (accounts.Account){
    CallingLog("createAccount()","CreateWalletAPI.go", "NULL", "NULL","NULL")
	
	decryptPassphrase = decrypt(pwd)
	decryptEmailID = decrypt(email)

	//func (ks *KeyStore) NewAccount(passphrase string) (accounts.Account, error)
	ks := keystore.NewKeyStore("/home/ubuntu/eth-dev/keystore", keystore.StandardScryptN, keystore.StandardScryptP)
	newAcc, _ := ks.NewAccount(decryptPassphrase)

	//fmt.Println("Address : ", reflect.TypeOf(newAcc.Address))
	setEmailIDPrivateAddressMapping(decryptEmailID, newAcc.Address)

	fmt.Println("\n ACCOUNT CREATED!!! ")
	ResultLog ("createAccount()","CreateWalletAPI.go", "NULL","NULL")
	
	return newAcc
}

/************** UPDATE ACCOUNT PASSPHRASE ************* */
func changePassphrase(addr string, pwd string, newPwd string, email string) error {
	CallingLog("changePassphrase()","CreateWalletAPI.go", "NULL", "NULL","NULL")
	
	decryptAddr = decrypt(addr)
	decryptPassphrase = decrypt(pwd)
	decryptNewPassphrase = decrypt(newPwd)
	decryptEmailID = decrypt(email)

	userAddress = common.HexToAddress(decryptAddr)
	acc := accounts.Account{Address: userAddress}
	log.Println("decryptPassphrase =====================>", decryptPassphrase)	
	log.Println("decryptNewPassphrase =====================>", decryptNewPassphrase)

	if (isUserValid(userAddress,decryptEmailID)) {

		//func (ks *KeyStore) Update(a accounts.Account, passphrase, newPassphrase string) error
		ks := keystore.NewKeyStore("/home/ubuntu/eth-dev/keystore", keystore.StandardScryptN, keystore.StandardScryptP);
		err := ks.Update(acc, decryptPassphrase, decryptNewPassphrase)
		return err
	} else {
		throwException("changePassphrase()", "CreateWalletAPI.go", decryptAddr, decryptEmailID,"Fatal Error")
		return nil
	}
	ResultLog ("changePassphrase()","CreateWalletAPI.go", "NULL","NULL")
	
return nil
}

/*************************** UNLOCK AN ACCOUNT **************************************** */
func unlockAccount(addr string, pwd string, email string) error {
	CallingLog("unlockAccount()","CreateWalletAPI.go", "NULL", "NULL","NULL")
	
	decryptAddr = decrypt(addr)
	decryptPassphrase = decrypt(pwd)
	decryptEmailID = decrypt(email)
	
	userAddress = common.HexToAddress(decryptAddr)
	acc := accounts.Account{Address: userAddress}
	
	if (isUserValid(userAddress,decryptEmailID)) {
		
		//func (ks *KeyStore) Unlock(a accounts.Account, passphrase string) error
		ks := keystore.NewKeyStore("/home/ubuntu/eth-dev/keystore", keystore.StandardScryptN, keystore.StandardScryptP);
		err := ks.Unlock(acc, decryptPassphrase)
		return err
	} else {
		throwException("unlockAccount()", "CreateWalletAPI.go", decryptAddr, decryptEmailID,"Fatal Error")
		return nil
	}
	ResultLog ("unlockAccount()","CreateWalletAPI.go", "NULL","NULL")
	
	return nil
}


/*************************** LOCK AN ACCOUNT *****************************************
func lockAccount(addr common.Address) error {
		
	//func (ks *KeyStore) Lock(addr common.Address) error
	ks := keystore.NewKeyStore("/home/ubuntu/eth-dev/keystore", keystore.StandardScryptN, keystore.StandardScryptP);
	err := ks.Lock(addr)
	return err
}
*/

/*************************** FIND AN ACCOUNT *****************************************
func findAccount(a accounts.Account) (accounts.Account) {

	//Refer: func (ks *KeyStore) Find(a accounts.Account) (accounts.Account, error)
	ks := keystore.NewKeyStore("/home/ubuntu/eth-dev/keystore", keystore.StandardScryptN, keystore.StandardScryptP);
	findAcc, _ := ks.Find(a)
	return findAcc
}
*/
/*************************** EXPORT AN ACCOUNT *****************************************
func exportAccount(a accounts.Account, passphrase, newPassphrase string) (keyJSON []byte) {
	
	//Refer: func (ks *KeyStore) Export(a accounts.Account, passphrase, newPassphrase string) (keyJSON []byte, err error)
	ks := keystore.NewKeyStore("/home/ubuntu/eth-dev/keystore", keystore.StandardScryptN, keystore.StandardScryptP);
	expAcc, err := ks.Export(a, passphrase, newPassphrase)
	fmt.Println("\nERROR ",err)
	return expAcc
}
*/
/*************************** DELETE AN ACCOUNT *****************************************
func deleteAccount(a accounts.Account, passphrase string) error {

	//Refer: func (ks *KeyStore) Delete(a accounts.Account, passphrase string) error
	ks := keystore.NewKeyStore("/home/ubuntu/eth-dev/keystore", keystore.StandardScryptN, keystore.StandardScryptP);
	err := ks.Delete(a, passphrase)
	return err
}
*/