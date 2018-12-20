package main

import (
	"math/big"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

/***************************************************************************************************/
//  GET TOKEN INFO
/**************************************************************************************************/

func getTokenName() (string) {
	CallingLog("getTokenName()","TokenCustom.go", "NULL", "NULL","NULL")
	client := getClient()
	objToken := getTokenObject(client)

	name, err := objToken.Name(nil)
	if err != nil {
		throwException("getTokenName()", "TokenCustom.go", "NULL", "NULL","Fatal Error")
	}
	
		
        ResultLog ("getTokenName()","TokenCustom.go", name,"NULL")
	return name
}

//code modified
func getTokenSupply() (string) {
	CallingLog("getTokenSupply()","TokenCustom.go", "NULL", "NULL","NULL")
	totalSupply:= getTokenSupplyHelper()

	var totalSupplyInStrings = getTokenInSJTFromWei(totalSupply)
	ResultLog("getTokenSupply()","TokenCustom.go", totalSupplyInStrings,"NULL")
	return totalSupplyInStrings
}

func getTokenSupplyHelper() (*big.Int) {
	CallingLog("getTokenSupplyHelper()","TokenCustom.go", "NULL", "NULL","NULL")
	client := getClient()
	objToken := getTokenObject(client)

	totalSupply, err := objToken.TotalSupply(nil)
	if err != nil {
		throwException("getTokenSupplyHelper()", "TokenCustom.go", "NULL", "NULL","Fatal Error")
	}
	//bigstr :=(totalSupply)
	ResultLog("getTokenSupplyHelper()","TokenCustom.go", "NULL","NULL")
	return totalSupply
}

func getTokenSymbol() (string) {
	CallingLog("getTokenSymbol()","TokenCustom.go", "NULL", "NULL","NULL")
	client := getClient()
	objToken := getTokenObject(client)

	symbol, err := objToken.Symbol(nil)
	if err != nil {
		throwException("getTokenSymbol()", "TokenCustom.go", "NULL", "NULL","Fatal Error")
	}
	ResultLog ("getTokenSymbol()","TokenCustom.go", symbol,"NULL")
	return symbol
}

func getTokenDecimals() (uint8) {
	CallingLog("getTokenDecimals()","TokenCustom.go", "NULL", "NULL","NULL")
	
	client := getClient()
	objToken := getTokenObject(client)

	decimals, err := objToken.Decimals(nil)
	if err != nil {
		throwException("getTokenDecimals()", "TokenCustom.go", "NULL", "NULL","Fatal Error")
	}
	ResultLog ("getTokenDecimals()","TokenCustom.go", "NULL","NULL")
	return decimals

}

func getTokenContractAddress() (common.Address) {
	CallingLog("getTokenContractAddress()","TokenCustom.go", "NULL", "NULL","NULL")
	var tokenAddr = getTokenAddress(ContractAddress)
	ResultLog ("getTokenContractAddress()","TokenCustom.go", "NULL","NULL")
	return tokenAddr

}

//code modified
func getTokenBuyPrice() (string) {
	CallingLog("getTokenBuyPrice()","TokenCustom.go", "NULL", "NULL","NULL")
	buyPrice := getTokenBuyPriceHelper()

	var buyPriceInUSD = getUsdValueFromCents(buyPrice)
	ResultLog ("getTokenBuyPrice()","TokenCustom.go", buyPriceInUSD,"NULL")
	return buyPriceInUSD
}

func getTokenBuyPriceHelper() (*big.Int) {
	CallingLog("getTokenBuyPriceHelper()","TokenCustom.go", "NULL", "NULL","NULL")
	client := getClient()
	objToken := getTokenObject(client)
	
	buyPrice, err := objToken.BuyPriceInCentsForSingleToken(nil)
	if err != nil {
		throwException("getTokenBuyPriceHelper()", "TokenCustom.go", "NULL", "NULL","Fatal Error")
	}
	//bigstr :=(buyPrice)
	ResultLog ("getTokenBuyPriceHelper()","TokenCustom.go", "NULL","NULL")
	return buyPrice
}

//code modified
func getTokenBalance() (string) {
	CallingLog("getTokenBalance()","TokenCustom.go", "NULL", "NULL","NULL")
	client := getClient()
	objToken := getTokenObject(client)
	
	tokenAddress = getTokenAddress(CoinbaseAddress)
	
	balance, err := objToken.BalanceOf(nil, tokenAddress)
	if err != nil {
		throwException("getTokenBalance()", "TokenCustom.go", "NULL", "NULL","Fatal Error")
	}

	var balanceInSJT = getTokenInSJTFromWei(balance)
	ResultLog ("getTokenBalance()","TokenCustom.go", balanceInSJT,"NULL")
	return balanceInSJT	
}

func getPvtAddByPubAdd(publicAddress common.Address) (common.Address) {
	CallingLog("getPvtAddByPubAdd()","TokenCustom.go", "NULL", "NULL","NULL")
	client := getClient()
	objToken := getTokenObject(client)

	privateAddr, err := objToken.GetPrivateAddByPublicAdd(nil, publicAddress)
	if err != nil {
		throwException("getPvtAddByPubAdd()", "TokenCustom.go", "NULL", "NULL","Fatal Error")
	}
	ResultLog ("getPvtAddByPubAdd()","TokenCustom.go", "NULL","NULL")
	return privateAddr	
}

func getTokenHolder() (uint) {
	
	return 0
}

func getTokenTransfers() (uint) {
	
	return 0
}
//code modified
func getICOSupply() (string){
	CallingLog("getICOSupply()","TokenCustom.go", "NULL", "NULL","NULL")

	icoSupply := getICOSupplyHelper()

	var icoSupplyInSJT = getTokenInSJTFromWei(icoSupply)
	ResultLog ("getICOSupply()","TokenCustom.go", icoSupplyInSJT,"NULL")
	return icoSupplyInSJT
}

func getICOSupplyHelper() (*big.Int) {
	CallingLog("getICOSupplyHelper()","TokenCustom.go", "NULL", "NULL","NULL")
	client := getClient()
	objToken := getTokenObject(client)
	
	icoSupply, err := objToken.GetICOSupply(nil)
	if err != nil {
		throwException("getICOSupplyHelper()", "TokenCustom.go", "NULL", "NULL","Fatal Error")
	}
	//bigstr :=(icoSupply)
	ResultLog ("getICOSupplyHelper()","TokenCustom.go", "NULL","NULL")
	return icoSupply
}

//code modified
func getICORemainingBal() (string){
	CallingLog("getICORemainingBal()","TokenCustom.go", "NULL", "NULL","NULL")
	icoRemaining:= getICORemainingBalHelper()
	var icoRemainingInSJT = getTokenInSJTFromWei(icoRemaining)
	ResultLog ("getICORemainingBal()","TokenCustom.go", icoRemainingInSJT,"NULL")
	return icoRemainingInSJT
}

func getICORemainingBalHelper() (*big.Int) {
	CallingLog("getICORemainingBalHelper()","TokenCustom.go", "NULL", "NULL","NULL")
	client := getClient()
	objToken := getTokenObject(client)
	
	icoRemaining, err := objToken.GetICORemainingBalance(nil)
	if err != nil {
		throwException("getICORemainingBalHelper()", "TokenCustom.go", "NULL", "NULL","Fatal Error")
	}
	//bigstr :=(icoRemaining)
	ResultLog ("getICORemainingBalHelper()","TokenCustom.go", "NULL","NULL")
	return icoRemaining
}

func getSJTReservedSupply() (string){
	CallingLog("getSJTReservedSupply()","TokenCustom.go", "NULL", "NULL","NULL")
	supply := getSJTReservedSupplyHelper()

	var supplyInSJT = getTokenInSJTFromWei(supply)
	ResultLog ("getSJTReservedSupply()","TokenCustom.go", supplyInSJT,"NULL")
	return supplyInSJT
}

func getSJTReservedSupplyHelper() (*big.Int) {
	CallingLog("getSJTReservedSupplyHelper()","TokenCustom.go", "NULL", "NULL","NULL")
	client := getClient()
	objToken := getTokenObject(client)
	
	supply, err := objToken.GetReservedCoins(nil)
	if err != nil {
		throwException("getSJTReservedSupplyHelper()", "TokenCustom.go", "NULL", "NULL","Fatal Error")
	}
	//bigstr :=(supply)
	ResultLog ("getSJTReservedSupplyHelper()","TokenCustom.go", "NULL","NULL")
	return supply
}

func getSJTReservedRemainingBal() (string){
	CallingLog("getSJTReservedRemainingBal()","TokenCustom.go", "NULL", "NULL","NULL")
	remaining := getSJTReservedRemainingBalHelper()

	var remainingInSJT = getTokenInSJTFromWei(remaining)
	ResultLog ("getSJTReservedRemainingBal()","TokenCustom.go", remainingInSJT,"NULL")
	return remainingInSJT
}

func getSJTReservedRemainingBalHelper() (*big.Int) {
	CallingLog("getSJTReservedRemainingBalHelper()","TokenCustom.go", "NULL", "NULL","NULL")
	client:= getClient()
	objToken := getTokenObject(client)

	remaining, err := objToken.GetReservedRemainingCoins(nil)
	if err != nil {
		throwException("getSJTReservedRemainingBalHelper()", "TokenCustom.go", "NULL", "NULL","Fatal Error")
	}
	//bigstr :=(remaining)
	ResultLog ("getSJTReservedRemainingBalHelper()","TokenCustom.go", "NULL","NULL")
	return remaining
}

//code modified
func getSellPrice() (string){

//	log.Println("================================> getSellPrice method called")
	CallingLog("getSellPrice()","TokenCustom.go", "NULL", "NULL","NULL")
	sellPrice := getSellPriceHelper()
//	log.Println("================================> SELL Price", sellPrice)
	
	var sellPriceInUSD = getUsdValueFromCents(sellPrice)
//	log.Println("================================> sellPriceInUSD", sellPriceInUSD)
	
	ResultLog ("getSellPrice()","TokenCustom.go", sellPriceInUSD,"NULL")
	return sellPriceInUSD
}

func getSellPriceHelper()(*big.Int){
	CallingLog("getSellPriceHelper()","TokenCustom.go", "NULL", "NULL","NULL")
	client := getClient()
	objToken := getTokenObject(client)
	
	sellPrice, err := objToken.GetSellPrice(nil)
	if err != nil {
		throwException("getSellPriceHelper()", "TokenCustom.go", "NULL", "NULL","Fatal Error")
	}
	//bigstr :=(sellPrice)
	ResultLog ("getSellPriceHelper()","TokenCustom.go", "NULL","NULL")
	return sellPrice
}

//code modified
func getBuyPrice() (string){
	CallingLog("getBuyPrice()","TokenCustom.go", "NULL", "NULL","NULL")
	buyPrice := getBuyPriceHelper()

	var buyPriceInUSD = getUsdValueFromCents(buyPrice)
	ResultLog ("getBuyPrice()","TokenCustom.go", buyPriceInUSD,"NULL")
	return buyPriceInUSD
}

func getBuyPriceHelper()(*big.Int){
	CallingLog("getBuyPriceHelper()","TokenCustom.go", "NULL", "NULL","NULL")
	client := getClient()
	objToken := getTokenObject(client)
	
	buyPrice, err := objToken.GetBuyPrice(nil)
	if err != nil {
		throwException("getBuyPriceHelper()", "TokenCustom.go", "NULL", "NULL","Fatal Error")
	}
//	bigstr :=(buyPrice)
	ResultLog ("getBuyPriceHelper()","TokenCustom.go", "NULL","NULL")
	return buyPrice
}

func getPrivateAddressByEmailID(email string) (privateAddr common.Address) {
	CallingLog("getPrivateAddressByEmailID()","TokenCustom.go", "NULL", "NULL","NULL")
	client := getClient()
	objToken := getTokenObject(client)
	
	privateAddr, err := objToken.GetPivateAddByEmailID(nil, email)
	if err != nil {
		throwException("getPrivateAddressByEmailID()", "TokenCustom.go", "NULL", "NULL","Fatal Error")
	}
	ResultLog ("getPrivateAddressByEmailID()","TokenCustom.go", "NULL","NULL")
	return
}

func getPvtAddrByPubAddr(publicAdd common.Address) (privateAddrs common.Address) {
	CallingLog("getPvtAddrByPubAddr()","TokenCustom.go", "NULL", "NULL","NULL")
	client := getClient()
	objToken := getTokenObject(client)
	
	privateAddrs, err := objToken.GetPrivateAddByPublicAdd(nil, publicAdd)
	if err != nil {
		throwException("getPvtAddrByPubAddr()", "TokenCustom.go", "NULL", "NULL","Fatal Error")
	}
	ResultLog ("getPvtAddrByPubAddr()","TokenCustom.go", "NULL","NULL")
	return
}

/***************************************************************************************************/
//  SET TOKEN INFO
/**************************************************************************************************/

func setEmailIDPrivateAddressMapping(email string, privateAddress common.Address) {
	CallingLog("setEmailIDPrivateAddressMapping()","TokenCustom.go", "NULL", "NULL","NULL")
	client := getClient()
	objToken := getTokenObject(client)
	auth := getAuthObject()

	result, err := objToken.MapEmailIDWithPrivateAdd(auth, email, privateAddress)
	if err != nil {
		throwException("setEmailIDPrivateAddressMapping()", "TokenCustom.go", "NULL", "NULL","Fatal Error")
	}
	log.Println("result = ", result)
	ResultLog ("setEmailIDPrivateAddressMapping()","TokenCustom.go", "NULL","NULL")
}

func setPublicAddressPrivateAddressMapping(publicAddress common.Address, privateAddress common.Address) {
	CallingLog("setPublicAddressPrivateAddressMapping()","TokenCustom.go", "NULL", "NULL","NULL")
	client := getClient()
	objToken := getTokenObject(client)
	auth := getAuthObject()

	result, err := objToken.SetPrivateAddByPublicAdd(auth, publicAddress, privateAddress)
	if err != nil {
		throwException("setPublicAddressPrivateAddressMapping()", "TokenCustom.go", "NULL", "NULL","Fatal Error")
	}
	log.Println("Result = ", result)
	ResultLog ("setPublicAddressPrivateAddressMapping()","TokenCustom.go", "NULL","NULL")
}

/*****************************************************************************************************************/
//

func getTokenAddress(s string) (common.Address){
	CallingLog("getTokenAddress()","TokenCustom.go", "NULL", "NULL","NULL")
	ResultLog ("getTokenAddress()","TokenCustom.go", "NULL","NULL")
	return common.HexToAddress(s)
}

func getClient() (client *ethclient.Client) {
	CallingLog("getClient()","TokenCustom.go", "NULL", "NULL","NULL")
	endPoint := "http://ec2-52-14-203-178.us-east-2.compute.amazonaws.com:8545/"
	client, err := ethclient.Dial(endPoint)
	if err != nil {
		throwException("getClient()", "TokenCustom.go", "NULL", "NULL","Fatal Error")
	}
	ResultLog ("getClient()","TokenCustom.go", "NULL","NULL")
	return
}

func getTokenObject(client *ethclient.Client)(objToken *Token){
	CallingLog("getTokenObject()","TokenCustom.go", "NULL", "NULL","NULL")
	objToken, err := NewToken(common.HexToAddress(ContractAddress), client)
	if err != nil {
		throwException("getTokenObject()", "TokenCustom.go", "NULL", "NULL","Fatal Error")
	}
	ResultLog ("getTokenObject()","TokenCustom.go", "NULL","NULL")
	return
}

func getTokenTransactorObject(client *ethclient.Client)(objToken *TokenTransactor){
	CallingLog("getTokenTransactorObject()","TokenCustom.go", "NULL", "NULL","NULL")
	objToken, err := NewTokenTransactor(common.HexToAddress(ContractAddress), client)
	if err != nil {
		throwException("getTokenTransactorObject()", "TokenCustom.go", "NULL", "NULL","Fatal Error")
	}
	ResultLog ("getTokenTransactorObject()","TokenCustom.go", "NULL","NULL")
	return
}

func getAuthObject()(auth *bind.TransactOpts){
	CallingLog("getAuthObject()","TokenCustom.go", "NULL", "NULL","NULL")
	auth, err := bind.NewTransactor(strings.NewReader(key), CoinbasePassword)
	if err != nil {
		throwException("getAuthObject()", "TokenCustom.go", "NULL", "NULL","Fatal Error")
	}
	ResultLog ("getAuthObject()","TokenCustom.go", "NULL","NULL")
	return
}