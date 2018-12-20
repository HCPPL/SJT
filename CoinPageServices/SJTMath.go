package main

import (
	"math/big"
//	"log"

	"github.com/decimal"
)

func getFundingTokensForUSDPayments(usdValueInCents *big.Int)(fundingTokensInWei_bigInt *big.Int){
	CallingLog("getFundingTokensForUSDPayments()","SJTMath.go", "NULL", "NULL","NULL")
	var usdValueInCentsDecimal  = decimal.NewFromBigInt(usdValueInCents, 0)
	
	var SJTInWeiDecimal, err1  = decimal.NewFromString(SJTInWei)
	if err1 != nil {
		throwException("getFundingTokensForUSDPayments()", "SJTMath.go", "NULL", "NULL","Fatal Error")
	}
	
	sellPriceInCents_bigInt := getSellPriceHelper()	
	var sellPriceInCentsDecimal = decimal.NewFromBigInt(sellPriceInCents_bigInt, 0)

	var fundingTokens = usdValueInCentsDecimal.Mul(SJTInWeiDecimal)
	fundingTokens =  fundingTokens.Div(sellPriceInCentsDecimal)
	var fundingTokens_string = fundingTokens.String()
	
	fundingTokensInWei_bigInt = big.NewInt(10)
	fundingTokensInWei_bigInt.SetString(fundingTokens_string, 10)
	//bigint := big.NewInt(fundingTokensInWei_bigInt)
	//bigstr := bigint.String()
	ResultLog ("getFundingTokensForUSDPayments()","SJTMath.go", "NULL","NULL")
	return
}

func getEthValueInWei(EthValue string) (*big.Int) {
	CallingLog("getEthValueInWei()","SJTMath.go", "NULL", "NULL","NULL")
	var EthValueInDecimal, err1  = decimal.NewFromString(EthValue)
	if err1 != nil {
		throwException("getEthValueInWei()", "SJTMath.go", "NULL", "NULL","Fatal Error")
	}

	var EthValueWeiInDecimal, err2  = decimal.NewFromString(EthValueInWei)
	if err2 != nil {
		throwException("getEthValueInWei()", "SJTMath.go", "NULL", "NULL","Fatal Error")
	}

	var TotalEthInWei = EthValueInDecimal.Mul(EthValueWeiInDecimal)
	var TotalEthInWei_string = TotalEthInWei.String()
	
	TotalEthInWei_bigInt := big.NewInt(10)
	TotalEthInWei_bigInt.SetString(TotalEthInWei_string, 10)
	//bigint := big.NewInt(TotalEthInWei_bigInt)
//	bigstr := bigint.String()
	ResultLog ("getEthValueInWei()","SJTMath.go", "NULL","NULL")
	return TotalEthInWei_bigInt
}

func getUSDValueInCents(USDValue string) (*big.Int) {
	CallingLog("getUSDValueInCents()","SJTMath.go", "NULL", "NULL","NULL")
	var USDValueInDecimal, err1  = decimal.NewFromString(USDValue)
//	log.Println("==============================================================");

	if err1 != nil {
//		log.Println("Error: ", err1);
		throwException("getUSDValueInCents()", "SJTMath.go", "NULL", "NULL","Fatal Error")
	}

	var USDValueCentsInDecimal, err2  = decimal.NewFromString(USDInCents)
	if err2 != nil {
		throwException("getUSDValueInCents()", "SJTMath.go", "NULL", "NULL","Fatal Error")
	}

	var TotalUSDInCents = USDValueInDecimal.Mul(USDValueCentsInDecimal)
	var TotalUSDInCents_string = TotalUSDInCents.String()
	
	TotalUSDInCents_bigInt := big.NewInt(10)
	TotalUSDInCents_bigInt.SetString(TotalUSDInCents_string, 10)
	//bigint := big.NewInt(TotalUSDInCents_bigInt)
	//bigstr := bigint.String()
	ResultLog ("getUSDValueInCents()","SJTMath.go", "NULL","NULL")
	return TotalUSDInCents_bigInt
}

func getSJTValueInWei(SJTValue string) (*big.Int) {

//	log.Println("=============> Get SJT value In WEI =================")

	CallingLog("getSJTValueInWei()","SJTMath.go", "NULL", "NULL","NULL")
	var SJTValueInDecimal, err1  = decimal.NewFromString(SJTValue)
	if err1 != nil {
		throwException("getSJTValueInWei()", "SJTMath.go", "NULL", "NULL","Fatal Error")
	}
//	log.Println("=============> Conversion =================")
	
	var SJTValueWeiInDecimal, err2  = decimal.NewFromString(SJTInWei)
	if err2 != nil {
		throwException("getSJTValueInWei()", "SJTMath.go", "NULL", "NULL","Fatal Error")
	}
//	log.Println("=============> Conversion =================")
	
	var TotalSJTInWei = SJTValueInDecimal.Mul(SJTValueWeiInDecimal)
	var TotalSJTInWei_string = TotalSJTInWei.String()
	
	TotalSJTInWei_bigInt := big.NewInt(10)
	TotalSJTInWei_bigInt.SetString(TotalSJTInWei_string, 10)
	//bigint := big.NewInt(TotalSJTInWei_bigInt)
	//bigstr := bigint.String()
	
	ResultLog ("getSJTValueInWei()","SJTMath.go", "NULL","NULL")
//	log.Println("=============> END Get SJT value In WEI =================")
	
	return TotalSJTInWei_bigInt
}

func getTokenInSJTFromWei(fundingTokensInWei *big.Int) (string){
	CallingLog("getTokenInSJTFromWei()","SJTMath.go", "NULL", "NULL","NULL")
	var fundingTokensInWei_Decimal = decimal.NewFromBigInt(fundingTokensInWei, 0)
	var SJTinWei_Decimal, err = decimal.NewFromString(SJTInWei)
	if err != nil {
		throwException("getTokenInSJTFromWei()", "SJTMath.go", "NULL", "NULL","Fatal Error")
	}

	var fundingTokensInSJT_Decimal = fundingTokensInWei_Decimal.Div(SJTinWei_Decimal)
	var fundingTokensInSJT_string = fundingTokensInSJT_Decimal.String()
	//bigint := big.NewInt(fundingTokensInSJT_string)
	//bigstr := bigint.String()
	ResultLog ("getTokenInSJTFromWei()","SJTMath.go", "NULL","NULL")
	return fundingTokensInSJT_string
}

func getUsdValueFromCents(totalUSDValueInCents *big.Int) (string){
//	log.Println("==========================================get usd value from cent called==================================================")
	CallingLog("getUsdValueFromCents()","SJTMath.go", "NULL", "NULL","NULL")
//	log.Println("PARAMS => ", totalUSDValueInCents)
	var totalUSDValueInCents_Decimal = decimal.NewFromBigInt(totalUSDValueInCents, 0)
	
	var usdValueInCents_Decimal, err = decimal.NewFromString(USDInCents)
	if err != nil {
		throwException("getUsdValueFromCents()", "SJTMath.go", "NULL", "NULL","Fatal Error")
	}
//	log.Println("totalUSDValueInCents_Decimal================>>>>>>>>>>>>", totalUSDValueInCents_Decimal)
//	log.Println("usdValueInCents_Decimal================>>>>>>>>>>>>", usdValueInCents_Decimal)
	
	var TokenValueInUSD_Decimal = totalUSDValueInCents_Decimal.Div(usdValueInCents_Decimal)
//	log.Println("TokenValueInUSD_Decimal================>>>>>>>>>>>>", TokenValueInUSD_Decimal)
	
	var TokenValueInUSD_string = TokenValueInUSD_Decimal.String()

	ResultLog ("getUsdValueFromCents()","SJTMath.go", TokenValueInUSD_string,"NULL")
	return TokenValueInUSD_string
}

func getEthValueFromWei(totalEthValueInWei *big.Int) (string){
	CallingLog("getEthValueFromWei()","SJTMath.go", "NULL", "NULL","NULL")
	var totalEthValueInWei_Decimal = decimal.NewFromBigInt(totalEthValueInWei, 0)
	var EthValueInWei_Decimal, err = decimal.NewFromString(EthValueInWei)
	if err != nil {
		throwException("getEthValueFromWei()", "SJTMath.go", "NULL", "NULL","Fatal Error")
	}
	
	var ETHValue_Decimal = totalEthValueInWei_Decimal.Div(EthValueInWei_Decimal)
	var ETHValue_string = ETHValue_Decimal.String()
	ResultLog ("getEthValueFromWei()","SJTMath.go", ETHValue_string,"NULL")
	return ETHValue_string
}

func doMultiply(number1 uint64, number2 uint64, number3 uint64) *big.Int {
	CallingLog("doMultiply()","SJTMath.go", "NULL", "NULL","NULL")
	var temp = big.NewInt(10)
    n1 := new(big.Int).SetUint64(number1)
	n2 := new(big.Int).SetUint64(number2)
	n3 := new(big.Int).SetUint64(number3)

	temp.Mul(n1, n2)
	temp.Mul(temp, n3)
	//bigint := big.NewInt(temp)
	//bigstr := bigint.String()
	ResultLog ("doMultiply()","SJTMath.go", "NULL","NULL")
	return temp
}

func doDivide(number1 *big.Int, number2 uint64) *big.Int {
	//log.Println("================= do Division method called =============")
	CallingLog("doDivision()","SJTMath.go", "NULL", "NULL","NULL")
	var temp = big.NewInt(10)
    n1 := new(big.Int).SetUint64(number2)
//	log.Println("=============> n1", n1)
	
	temp.Div(number1, n1)
//	log.Println("================= temp =============", temp)
	
	//bigint := big.NewInt(temp)
	//bigstr := bigint.String()
	ResultLog ("doDivision()","SJTMath.go", "NULL","NULL")
	return temp
}

func doDivision(number1 string, number2 string) string {
//	log.Println("================= do Division method called =============")
	
	CallingLog("doDivision()","SJTMath.go", "NULL", "NULL","NULL")
	var number1_Decimal, err1 = decimal.NewFromString(number1)
//	log.Println("================= number1_Decimal =============", number1_Decimal)
	
	if err1 != nil {
		throwException("doDivision()", "SJTMath.go", "NULL", "NULL","Fatal Error")
	}
//	log.Println("================= no error found =============")
	
	var number2_Decimal, err2 = decimal.NewFromString(number2)
//	log.Println("================= number2_Decimal =============", number2_Decimal)
	
	if err2 != nil {
		throwException("doDivision()", "SJTMath.go", "NULL", "NULL","Fatal Error")
	}
//	log.Println("================= no error found =============")
	

	var result_Decimal = number1_Decimal.Div(number2_Decimal)
//	log.Println("================= DIV =============")
	
	var result_string = result_Decimal.String()
//	log.Println("================= END =============")
	
	ResultLog ("doDivision()","SJTMath.go", result_string,"NULL")

	return result_string
}

func doMultiplication(number1 string, number2 string, number3 string) *big.Int {

//	log.Println("================DO MULTIPLICATION =====================")


	CallingLog("doMultiplication()","SJTMath.go", "NULL", "NULL","NULL")
	var number1_Decimal, err1  = decimal.NewFromString(number1)
	if err1 != nil {
		throwException("doMultiplication()", "SJTMath.go", "NULL", "NULL","Fatal Error")
	}

	var number2_Decimal, err2  = decimal.NewFromString(number2)
	if err2 != nil {
		throwException("doMultiplication()", "SJTMath.go", "NULL", "NULL","Fatal Error")
	}

	var number3_Decimal, err3  = decimal.NewFromString(number3)
	if err3 != nil {
		throwException("doMultiplication()", "SJTMath.go", "NULL", "NULL","Fatal Error")
	}

	var result_Decimal = number1_Decimal.Mul(number2_Decimal)
	result_Decimal = result_Decimal.Mul(number3_Decimal)
	var result_string = result_Decimal.String()
	
	result_bigInt := big.NewInt(10)
	result_bigInt.SetString(result_string, 10)
	//bigint := big.NewInt(result_bigInt)
//bigstr := bigint.String()
	ResultLog ("doMultiplication()","SJTMath.go", "NULL","NULL")
//	log.Println("================DO MULTIPLICATION ENDED=====================")
	
	return result_bigInt
}