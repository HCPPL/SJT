package main

import (
    "encoding/json"
    "log"
    "net/http"
    //"fmt"

    "github.com/ethereum/go-ethereum/common"
    "github.com/gorilla/mux"
    
)

type SJTTokenInfo struct {
	Name string
	InitialSupply string
	Symbol string
	Decimals uint8
	ContractAddress common.Address
	ValuePerToken string
	Balance string
	TokenHolder uint
    NoOfTransfers uint
}

var getTokenInfo SJTTokenInfo

// Get Coin Page Informations
//func GetTokenInfo(w http.ResponseWriter, req *http.Request) {
var GetTokenInfo = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){  
  CallingLog("GetTokenInfo","CoinPageAPI.go", "NULL", "NULL","NULL")
    var TotalSupplyInWei = getTokenSupplyHelper()
    var TotalSupplyInSJT = getTokenInSJTFromWei(TotalSupplyInWei)

    var TokenSellPriceInCents = getSellPriceHelper()
    var TokenSellPriceInUSD = getUsdValueFromCents(TokenSellPriceInCents)

    getTokenInfo := new(SJTTokenInfo)
    getTokenInfo.Name = getTokenName()
    getTokenInfo.InitialSupply = TotalSupplyInSJT
    getTokenInfo.Symbol = getTokenSymbol()
    getTokenInfo.Decimals = getTokenDecimals()
    getTokenInfo.ContractAddress = getTokenContractAddress()
    getTokenInfo.ValuePerToken = TokenSellPriceInUSD
    getTokenInfo.Balance = getTokenBalance()
    getTokenInfo.TokenHolder = getTokenHolder()
    getTokenInfo.NoOfTransfers = getTokenTransfers()

    log.Println("getTokenInfo = ", getTokenInfo)
    
    b, err := json.Marshal(getTokenInfo)
	
		if err != nil {
            throwException("GetTokenInfo", "CoinPageAPI.go", decryptAddr, decryptEmailID,"Fatal Error")
			log.Println("Error: %s", err)
	
			return
			
	
		}
	
		var s5 string= string(b)
		
        ResultLog ("GetTokenInfo","CoinPageAPI.go", s5,"NULL")

    json.NewEncoder(w).Encode(getTokenInfo)  
})

func main() {
    router := mux.NewRouter()
    //router.HandleFunc("/getTokenInfo", GetTokenInfo).Methods("GET")
    //router.HandleFunc("/issueTokens/{addr}/{USDValue}", IssueTokensByUSD).Methods("GET")
    router.HandleFunc("/issueTokensByETH/{contractAddress}/{key}/{addr}/{Eth}", IssueTokensByETH).Methods("GET")
    //router.HandleFunc("/transferTokens/{addr}/{userEmailID}/{receiverAddr}/{tokens}", TransferSJTTokens).Methods("GET")
    
    router.Handle("/getTokenInfo", jwtMiddleware.Handler(GetTokenInfo)).Methods("POST")
    router.Handle("/getBalance/{addr}", jwtMiddleware.Handler(GetSJTAddressBalanceInfo)).Methods("POST")
    router.Handle("/encrypt/{EncryptText}", jwtMiddleware.Handler(EncryptParams)).Methods("POST")
    router.Handle("/decrypt/{DecryptText}", jwtMiddleware.Handler(DecryptParams)).Methods("POST")
    router.Handle("/CreateAccount/{passphrase}/{userEmailID}", jwtMiddleware.Handler(CreateAccount)).Methods("POST")
    router.Handle("/UpdateAccount/{addr}/{passphrase}/{newPassphrase}/{userEmailID}", jwtMiddleware.Handler(UpdateAccount)).Methods("POST")
    router.Handle("/UnlockAccount/{addr}/{passphrase}/{userEmailID}", jwtMiddleware.Handler(UnlockAccount)).Methods("POST")
    router.Handle("/issueTokens/{orderID}/{addr}/{userEmailID}/{USDValue}", jwtMiddleware.Handler(IssueTokensByUSD)).Methods("POST")
    router.Handle("/transferTokens/{orderID}/{addr}/{userEmailID}/{receiverAddr}/{tokens}", jwtMiddleware.Handler(TransferSJTTokens)).Methods("POST")
    router.Handle("/issueReservedTokens/{orderID}/{addr}/{userEmailID}/{receiverAddr}/{tokens}", jwtMiddleware.Handler(IssueReservedTokens)).Methods("POST")
    router.Handle("/mintICOTokens/{addr}/{userEmailID}/{value}", jwtMiddleware.Handler(MintICOSupply)).Methods("POST")
    router.Handle("/mintTokens/{addr}/{userEmailID}/{value}", jwtMiddleware.Handler(MintTokens)).Methods("POST")
    router.Handle("/sellPrice/{addr}/{userEmailID}/{sellPrice}", jwtMiddleware.Handler(SetSellPrice)).Methods("POST")
    router.Handle("/buyPrice/{addr}/{userEmailID}/{buyPrice}", jwtMiddleware.Handler(SetBuyPrice)).Methods("POST")
    router.Handle("/setPublicPrivateAddresses/{addr}/{userEmailID}/{publicAddress}", jwtMiddleware.Handler(SetPublicAddressWithPrivateAddress)).Methods("POST")
    router.Handle("/getPrivateAddress/{publicAddress}", jwtMiddleware.Handler(GetPrivateAddressByPublicAddress)).Methods("POST")
    router.Handle("/getPrivateAddr/{emailAddress}", jwtMiddleware.Handler(GetPvtByEmail)).Methods("POST")
    router.Handle("/getICOInfo", jwtMiddleware.Handler(GetICOInfo)).Methods("POST")
    router.Handle("/getSJTReservedInfo", jwtMiddleware.Handler(GetSJTReservedInfo)).Methods("POST")
    router.Handle("/getSellBuyInfo", jwtMiddleware.Handler(GetSellPriceBuyPriceInfo)).Methods("POST")
    //router.Handle("/getTokenInfoSecurity", jwtMiddleware.Handler(GetTokenInfoSecurity)).Methods("POST")
    //router.Handle("/getFundingTokens/{addr}/{userEmailID}", jwtMiddleware.Handler(GetFundingTokens)).Methods("POST")
    //router.Handle("/getUSDTokens/{addr}/{userEmailID}", jwtMiddleware.Handler(GetUSDTokens)).Methods("POST")
    router.Handle("/getTokenPriceInfo", jwtMiddleware.Handler(GetTokenPrice)).Methods("POST")
    //router.Handle("/eth/currentprice", jwtMiddleware.Handler(GetETHPrice)).Methods("POST")
    router.Handle("/getSJTfromETH/{eth}", jwtMiddleware.Handler(EthSJTConvertion)).Methods("POST")
   router.Handle("/setEmailWithPrivate/{emailId}/{privateAddress}", jwtMiddleware.Handler(SetEmailWithPrivateAddress)).Methods("POST")    
    
   log.Fatal(http.ListenAndServe(":8080", router))
}