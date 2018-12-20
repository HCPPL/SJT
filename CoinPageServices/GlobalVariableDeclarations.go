package main

import (
	"github.com/ethereum/go-ethereum/common"
)

var apiToUpdateDB = "https://stjameshdinvtrust.co/api" 			//var apiToUpdateDB = "http://sjtcoin.blockchainpeer.com/api"
//http://sjtcoin.blockchainpeer.com/api/purchaseapi.php/?orderID=27dxedpr1AwfExNxNFFq6-cy&SJTReceived=4shRb1pPfBQ72VrIuLyzw5FpXxr5ng==&senderAddress=RqE4BsxvIpX_N5L-nCRDhzOBYGJhfelPGHY0TCru4feCdNLAWw8IkyJzLqOB7Wa8fA3ZFyVngXirnw==&ethReceivedFrom=RqE4BsxvIpX_N5L-nCRDhzOBYGJhfelPGHY0TCru4feCdNLAWw8IkyJzLqOB7Wa8fA3ZFyVngXirnw==
var tokenAddress common.Address
var ContractAddress = "0x699aecf5e7e46fa742af27df9f56d0cd3a69a620" //"0xd7dbf0c005002c9f3a61fe7b0e14d3507217ebec" // "0x8e1abbccf73741ebc79c16f32802d582a5bcda74"//"0x218665c4b31e28bf9af805008d4c5b682e55223b"
var CoinbaseAddress = "0x6cdb041ac80e0b634a8376b2fe4682d421ee757d"
var CoinbasePassword = "SJTCOIN@2017@100"

var mySecretKeyForEthTxn = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwibmFtZSI6IkFkbyBLdWtpYyJ9.5jw2SUe8LFEW_1Zt2dyX3AyiPWivwFF8B0XumWu0eQ8"
var SJTBridgeContractAddress = "6060a9a876dfbe5a8108ca101cc69dbf018d49a1" //"65dcec8c78c8c271154212a23ce9dcae4cda61dd" //"1bcf1227f2e052058ec4950c4a5bb4e32620d5b2"	//"a2a2d3c41e981d9b9edf63ebae559c24b6f6d06c"		// for 8081: "121b27aba0b3992b058b26b892bec073ef76fd7d"
var SJTBridgeContractAddress_full = "0x"+ SJTBridgeContractAddress
var nullAddr = common.HexToAddress("0x0000000000000000000000000000000000000000")

var value string
var userAddr string
var emailID string
var passphrase string
var newPassphrase string
var publicAddr string
var orderID string
var encryptText string
var decryptText string
var decryptOrderID string
var decryptAddr string
var decryptPublicAddr string
var decryptEmailID string
var decryptPassphrase string
var decryptNewPassphrase string
var userAddress common.Address
var publicAddress common.Address
var senderPublicAddress common.Address
var senderPrivateAddress common.Address

var EthValueInWei = "1000000000000000000"
var SJTInWei = "1000000000"
var USDInCents = "100"

var secretKey []byte = []byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpX")
const key = "{\"address\":\"6cdb041ac80e0b634a8376b2fe4682d421ee757d\",\"crypto\":{\"cipher\":\"aes-128-ctr\",\"ciphertext\":\"5b685a9884071cc479177e5cb2fd98bb85d921b58f6b341eb39f11e44d0b1b38\",\"cipherparams\":{\"iv\":\"b1da14dbae044f0f9fc49b1ce2e86259\"},\"kdf\":\"scrypt\",\"kdfparams\":{\"dklen\":32,\"n\":262144,\"p\":1,\"r\":8,\"salt\":\"1c262c5893d7f1c2cca963d378c76ac829f5e70518c059fbc2211860c1546131\"},\"mac\":\"fe5261f26659c3d4a50ff86e760531b8bc4d0f8c81e142c11e8461d688138161\"},\"id\":\"5d1c96a3-0c8f-4349-aeac-67e0878dda2c\",\"version\":3}"

//*******************************************************************************************************************
//


func isUserValid(userPrivateAddress common.Address, userEmailID string ) (bool) {
	CallingLog("isUserValid()","GlobalVariableDeclarations.go", "NULL", "NULL","NULL")
	
	var result bool = false

	fetchedPrivateAddress := getPrivateAddressByEmailID(userEmailID)

	if (fetchedPrivateAddress == userPrivateAddress) {
		result = true
	}
	ResultLog ("isUserValid()","GlobalVariableDeclarations.go", "NULL","NULL")
	
return result
}
	
func throwException(method string, api string, address string, email string,errors string){
	method1:=method
	api1:=api
	address1:=address
	email1:=email
	errors1:=errors
	ErrorLog(method1, api1, address1, email1, errors1)	
}