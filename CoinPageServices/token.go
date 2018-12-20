// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// TokenABI is the input ABI used to generate the binding from.
const TokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getBuyPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"methodType\",\"type\":\"string\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"checkBalanceBeforeTransaction\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_sbalanceOfpender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"SJTCwdSale\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ICOSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"publicAddress\",\"type\":\"address\"},{\"name\":\"privateAddress\",\"type\":\"address\"}],\"name\":\"setPrivateAddByPublicAdd\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"initCoinSupply\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"reservedCoins\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"EmailIDPrivateAddrMapping\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"email\",\"type\":\"string\"},{\"name\":\"privateAddress\",\"type\":\"address\"}],\"name\":\"MapEmailIDWithPrivateAdd\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fundingTokens\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"TokensIssuedFromReservedBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"checkBalanceForTransferTokens\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getReservedCoins\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"coinSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"completedAt\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"releaseTokensForICO\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSellPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"email\",\"type\":\"string\"}],\"name\":\"getPivateAddByEmailID\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"standard\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"SJTinWei\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newBuyPriceInCents\",\"type\":\"uint256\"}],\"name\":\"setBuyPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"releaseTokensFromReserved\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"buyPriceInCentsForSingleToken\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ICORemainingBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"increasedCoins\",\"type\":\"uint256\"}],\"name\":\"mintICOSupply\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getFundingTokensForReserved\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"sellPriceInCentsForSingleToken\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"publicAddr\",\"type\":\"address\"}],\"name\":\"getPrivateAddByPublicAdd\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTokenHolders\",\"outputs\":[{\"name\":\"noOfHolders\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getICORemainingBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"reservedCoinsBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferTokens\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getICOSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\"},{\"name\":\"fundingTokens\",\"type\":\"uint256\"}],\"name\":\"releaseTokensByUSD\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"PublicPrivateAddressesMapping\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"mintedAmount\",\"type\":\"uint256\"}],\"name\":\"mintToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"source\",\"type\":\"string\"}],\"name\":\"stringToBytes32\",\"outputs\":[{\"name\":\"result\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getFundingTokens\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"priceInUSD\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fundingTokensForReserved\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getReservedRemainingCoins\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newSellPriceInCents\",\"type\":\"uint256\"}],\"name\":\"setSellPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LogRefund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"currentTotal\",\"type\":\"uint256\"}],\"name\":\"LogFundingReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"totalRaised\",\"type\":\"uint256\"}],\"name\":\"LogFundingSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"url\",\"type\":\"string\"}],\"name\":\"LogFunderInitialized\",\"type\":\"event\"}]"

// Token is an auto generated Go binding around an Ethereum contract.
type Token struct {
	TokenCaller     // Read-only binding to the contract
	TokenTransactor // Write-only binding to the contract
}

// TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenSession struct {
	Contract     *Token            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenCallerSession struct {
	Contract *TokenCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenTransactorSession struct {
	Contract     *TokenTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRaw struct {
	Contract *Token // Generic contract binding to access the raw methods on
}

// TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenCallerRaw struct {
	Contract *TokenCaller // Generic read-only contract binding to access the raw methods on
}

// TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenTransactorRaw struct {
	Contract *TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewToken creates a new instance of Token, bound to a specific deployed contract.
func NewToken(address common.Address, backend bind.ContractBackend) (*Token, error) {
	contract, err := bindToken(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}}, nil
}

// NewTokenCaller creates a new read-only instance of Token, bound to a specific deployed contract.
func NewTokenCaller(address common.Address, caller bind.ContractCaller) (*TokenCaller, error) {
	contract, err := bindToken(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &TokenCaller{contract: contract}, nil
}

// NewTokenTransactor creates a new write-only instance of Token, bound to a specific deployed contract.
func NewTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenTransactor, error) {
	contract, err := bindToken(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &TokenTransactor{contract: contract}, nil
}

// bindToken binds a generic wrapper to an already deployed contract.
func bindToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.contract.Transact(opts, method, params...)
}

// EmailIDPrivateAddrMapping is a free data retrieval call binding the contract method 0x258e542f.
//
// Solidity: function EmailIDPrivateAddrMapping( bytes32) constant returns(address)
func (_Token *TokenCaller) EmailIDPrivateAddrMapping(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "EmailIDPrivateAddrMapping", arg0)
	return *ret0, err
}

// EmailIDPrivateAddrMapping is a free data retrieval call binding the contract method 0x258e542f.
//
// Solidity: function EmailIDPrivateAddrMapping( bytes32) constant returns(address)
func (_Token *TokenSession) EmailIDPrivateAddrMapping(arg0 [32]byte) (common.Address, error) {
	return _Token.Contract.EmailIDPrivateAddrMapping(&_Token.CallOpts, arg0)
}

// EmailIDPrivateAddrMapping is a free data retrieval call binding the contract method 0x258e542f.
//
// Solidity: function EmailIDPrivateAddrMapping( bytes32) constant returns(address)
func (_Token *TokenCallerSession) EmailIDPrivateAddrMapping(arg0 [32]byte) (common.Address, error) {
	return _Token.Contract.EmailIDPrivateAddrMapping(&_Token.CallOpts, arg0)
}

// ICORemainingBalance is a free data retrieval call binding the contract method 0x7458cc15.
//
// Solidity: function ICORemainingBalance() constant returns(uint256)
func (_Token *TokenCaller) ICORemainingBalance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "ICORemainingBalance")
	return *ret0, err
}

// ICORemainingBalance is a free data retrieval call binding the contract method 0x7458cc15.
//
// Solidity: function ICORemainingBalance() constant returns(uint256)
func (_Token *TokenSession) ICORemainingBalance() (*big.Int, error) {
	return _Token.Contract.ICORemainingBalance(&_Token.CallOpts)
}

// ICORemainingBalance is a free data retrieval call binding the contract method 0x7458cc15.
//
// Solidity: function ICORemainingBalance() constant returns(uint256)
func (_Token *TokenCallerSession) ICORemainingBalance() (*big.Int, error) {
	return _Token.Contract.ICORemainingBalance(&_Token.CallOpts)
}

// ICOSupply is a free data retrieval call binding the contract method 0x15b73a1d.
//
// Solidity: function ICOSupply() constant returns(uint256)
func (_Token *TokenCaller) ICOSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "ICOSupply")
	return *ret0, err
}

// ICOSupply is a free data retrieval call binding the contract method 0x15b73a1d.
//
// Solidity: function ICOSupply() constant returns(uint256)
func (_Token *TokenSession) ICOSupply() (*big.Int, error) {
	return _Token.Contract.ICOSupply(&_Token.CallOpts)
}

// ICOSupply is a free data retrieval call binding the contract method 0x15b73a1d.
//
// Solidity: function ICOSupply() constant returns(uint256)
func (_Token *TokenCallerSession) ICOSupply() (*big.Int, error) {
	return _Token.Contract.ICOSupply(&_Token.CallOpts)
}

// PublicPrivateAddressesMapping is a free data retrieval call binding the contract method 0xc520b0b4.
//
// Solidity: function PublicPrivateAddressesMapping( address) constant returns(address)
func (_Token *TokenCaller) PublicPrivateAddressesMapping(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "PublicPrivateAddressesMapping", arg0)
	return *ret0, err
}

// PublicPrivateAddressesMapping is a free data retrieval call binding the contract method 0xc520b0b4.
//
// Solidity: function PublicPrivateAddressesMapping( address) constant returns(address)
func (_Token *TokenSession) PublicPrivateAddressesMapping(arg0 common.Address) (common.Address, error) {
	return _Token.Contract.PublicPrivateAddressesMapping(&_Token.CallOpts, arg0)
}

// PublicPrivateAddressesMapping is a free data retrieval call binding the contract method 0xc520b0b4.
//
// Solidity: function PublicPrivateAddressesMapping( address) constant returns(address)
func (_Token *TokenCallerSession) PublicPrivateAddressesMapping(arg0 common.Address) (common.Address, error) {
	return _Token.Contract.PublicPrivateAddressesMapping(&_Token.CallOpts, arg0)
}

// SJTinWei is a free data retrieval call binding the contract method 0x5bdd6121.
//
// Solidity: function SJTinWei() constant returns(uint256)
func (_Token *TokenCaller) SJTinWei(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "SJTinWei")
	return *ret0, err
}

// SJTinWei is a free data retrieval call binding the contract method 0x5bdd6121.
//
// Solidity: function SJTinWei() constant returns(uint256)
func (_Token *TokenSession) SJTinWei() (*big.Int, error) {
	return _Token.Contract.SJTinWei(&_Token.CallOpts)
}

// SJTinWei is a free data retrieval call binding the contract method 0x5bdd6121.
//
// Solidity: function SJTinWei() constant returns(uint256)
func (_Token *TokenCallerSession) SJTinWei() (*big.Int, error) {
	return _Token.Contract.SJTinWei(&_Token.CallOpts)
}

// TokensIssuedFromReservedBalance is a free data retrieval call binding the contract method 0x2cb4639c.
//
// Solidity: function TokensIssuedFromReservedBalance() constant returns(uint256)
func (_Token *TokenCaller) TokensIssuedFromReservedBalance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "TokensIssuedFromReservedBalance")
	return *ret0, err
}

// TokensIssuedFromReservedBalance is a free data retrieval call binding the contract method 0x2cb4639c.
//
// Solidity: function TokensIssuedFromReservedBalance() constant returns(uint256)
func (_Token *TokenSession) TokensIssuedFromReservedBalance() (*big.Int, error) {
	return _Token.Contract.TokensIssuedFromReservedBalance(&_Token.CallOpts)
}

// TokensIssuedFromReservedBalance is a free data retrieval call binding the contract method 0x2cb4639c.
//
// Solidity: function TokensIssuedFromReservedBalance() constant returns(uint256)
func (_Token *TokenCallerSession) TokensIssuedFromReservedBalance() (*big.Int, error) {
	return _Token.Contract.TokensIssuedFromReservedBalance(&_Token.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_Token *TokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_Token *TokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Token.Contract.Allowance(&_Token.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_Token *TokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Token.Contract.Allowance(&_Token.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_Token *TokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_Token *TokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_Token *TokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, _owner)
}

// BuyPriceInCentsForSingleToken is a free data retrieval call binding the contract method 0x69e853f2.
//
// Solidity: function buyPriceInCentsForSingleToken() constant returns(uint256)
func (_Token *TokenCaller) BuyPriceInCentsForSingleToken(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "buyPriceInCentsForSingleToken")
	return *ret0, err
}

// BuyPriceInCentsForSingleToken is a free data retrieval call binding the contract method 0x69e853f2.
//
// Solidity: function buyPriceInCentsForSingleToken() constant returns(uint256)
func (_Token *TokenSession) BuyPriceInCentsForSingleToken() (*big.Int, error) {
	return _Token.Contract.BuyPriceInCentsForSingleToken(&_Token.CallOpts)
}

// BuyPriceInCentsForSingleToken is a free data retrieval call binding the contract method 0x69e853f2.
//
// Solidity: function buyPriceInCentsForSingleToken() constant returns(uint256)
func (_Token *TokenCallerSession) BuyPriceInCentsForSingleToken() (*big.Int, error) {
	return _Token.Contract.BuyPriceInCentsForSingleToken(&_Token.CallOpts)
}

// CheckBalanceBeforeTransaction is a free data retrieval call binding the contract method 0x0825ee75.
//
// Solidity: function checkBalanceBeforeTransaction(methodType string, _to address, _value uint256) constant returns(success bool)
func (_Token *TokenCaller) CheckBalanceBeforeTransaction(opts *bind.CallOpts, methodType string, _to common.Address, _value *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "checkBalanceBeforeTransaction", methodType, _to, _value)
	return *ret0, err
}

// CheckBalanceBeforeTransaction is a free data retrieval call binding the contract method 0x0825ee75.
//
// Solidity: function checkBalanceBeforeTransaction(methodType string, _to address, _value uint256) constant returns(success bool)
func (_Token *TokenSession) CheckBalanceBeforeTransaction(methodType string, _to common.Address, _value *big.Int) (bool, error) {
	return _Token.Contract.CheckBalanceBeforeTransaction(&_Token.CallOpts, methodType, _to, _value)
}

// CheckBalanceBeforeTransaction is a free data retrieval call binding the contract method 0x0825ee75.
//
// Solidity: function checkBalanceBeforeTransaction(methodType string, _to address, _value uint256) constant returns(success bool)
func (_Token *TokenCallerSession) CheckBalanceBeforeTransaction(methodType string, _to common.Address, _value *big.Int) (bool, error) {
	return _Token.Contract.CheckBalanceBeforeTransaction(&_Token.CallOpts, methodType, _to, _value)
}

// CheckBalanceForTransferTokens is a free data retrieval call binding the contract method 0x2f1c8628.
//
// Solidity: function checkBalanceForTransferTokens(_from address, _to address, _value uint256) constant returns(success bool)
func (_Token *TokenCaller) CheckBalanceForTransferTokens(opts *bind.CallOpts, _from common.Address, _to common.Address, _value *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "checkBalanceForTransferTokens", _from, _to, _value)
	return *ret0, err
}

// CheckBalanceForTransferTokens is a free data retrieval call binding the contract method 0x2f1c8628.
//
// Solidity: function checkBalanceForTransferTokens(_from address, _to address, _value uint256) constant returns(success bool)
func (_Token *TokenSession) CheckBalanceForTransferTokens(_from common.Address, _to common.Address, _value *big.Int) (bool, error) {
	return _Token.Contract.CheckBalanceForTransferTokens(&_Token.CallOpts, _from, _to, _value)
}

// CheckBalanceForTransferTokens is a free data retrieval call binding the contract method 0x2f1c8628.
//
// Solidity: function checkBalanceForTransferTokens(_from address, _to address, _value uint256) constant returns(success bool)
func (_Token *TokenCallerSession) CheckBalanceForTransferTokens(_from common.Address, _to common.Address, _value *big.Int) (bool, error) {
	return _Token.Contract.CheckBalanceForTransferTokens(&_Token.CallOpts, _from, _to, _value)
}

// CoinSupply is a free data retrieval call binding the contract method 0x31f170c2.
//
// Solidity: function coinSupply() constant returns(uint256)
func (_Token *TokenCaller) CoinSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "coinSupply")
	return *ret0, err
}

// CoinSupply is a free data retrieval call binding the contract method 0x31f170c2.
//
// Solidity: function coinSupply() constant returns(uint256)
func (_Token *TokenSession) CoinSupply() (*big.Int, error) {
	return _Token.Contract.CoinSupply(&_Token.CallOpts)
}

// CoinSupply is a free data retrieval call binding the contract method 0x31f170c2.
//
// Solidity: function coinSupply() constant returns(uint256)
func (_Token *TokenCallerSession) CoinSupply() (*big.Int, error) {
	return _Token.Contract.CoinSupply(&_Token.CallOpts)
}

// CompletedAt is a free data retrieval call binding the contract method 0x38771242.
//
// Solidity: function completedAt() constant returns(uint256)
func (_Token *TokenCaller) CompletedAt(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "completedAt")
	return *ret0, err
}

// CompletedAt is a free data retrieval call binding the contract method 0x38771242.
//
// Solidity: function completedAt() constant returns(uint256)
func (_Token *TokenSession) CompletedAt() (*big.Int, error) {
	return _Token.Contract.CompletedAt(&_Token.CallOpts)
}

// CompletedAt is a free data retrieval call binding the contract method 0x38771242.
//
// Solidity: function completedAt() constant returns(uint256)
func (_Token *TokenCallerSession) CompletedAt() (*big.Int, error) {
	return _Token.Contract.CompletedAt(&_Token.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() constant returns(address)
func (_Token *TokenCaller) Creator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "creator")
	return *ret0, err
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() constant returns(address)
func (_Token *TokenSession) Creator() (common.Address, error) {
	return _Token.Contract.Creator(&_Token.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() constant returns(address)
func (_Token *TokenCallerSession) Creator() (common.Address, error) {
	return _Token.Contract.Creator(&_Token.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Token *TokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Token *TokenSession) Decimals() (uint8, error) {
	return _Token.Contract.Decimals(&_Token.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Token *TokenCallerSession) Decimals() (uint8, error) {
	return _Token.Contract.Decimals(&_Token.CallOpts)
}

// FundingTokens is a free data retrieval call binding the contract method 0x2b359e6b.
//
// Solidity: function fundingTokens() constant returns(uint256)
func (_Token *TokenCaller) FundingTokens(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "fundingTokens")
	return *ret0, err
}

// FundingTokens is a free data retrieval call binding the contract method 0x2b359e6b.
//
// Solidity: function fundingTokens() constant returns(uint256)
func (_Token *TokenSession) FundingTokens() (*big.Int, error) {
	return _Token.Contract.FundingTokens(&_Token.CallOpts)
}

// FundingTokens is a free data retrieval call binding the contract method 0x2b359e6b.
//
// Solidity: function fundingTokens() constant returns(uint256)
func (_Token *TokenCallerSession) FundingTokens() (*big.Int, error) {
	return _Token.Contract.FundingTokens(&_Token.CallOpts)
}

// FundingTokensForReserved is a free data retrieval call binding the contract method 0xf5e32a24.
//
// Solidity: function fundingTokensForReserved() constant returns(uint256)
func (_Token *TokenCaller) FundingTokensForReserved(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "fundingTokensForReserved")
	return *ret0, err
}

// FundingTokensForReserved is a free data retrieval call binding the contract method 0xf5e32a24.
//
// Solidity: function fundingTokensForReserved() constant returns(uint256)
func (_Token *TokenSession) FundingTokensForReserved() (*big.Int, error) {
	return _Token.Contract.FundingTokensForReserved(&_Token.CallOpts)
}

// FundingTokensForReserved is a free data retrieval call binding the contract method 0xf5e32a24.
//
// Solidity: function fundingTokensForReserved() constant returns(uint256)
func (_Token *TokenCallerSession) FundingTokensForReserved() (*big.Int, error) {
	return _Token.Contract.FundingTokensForReserved(&_Token.CallOpts)
}

// GetBuyPrice is a free data retrieval call binding the contract method 0x018a25e8.
//
// Solidity: function getBuyPrice() constant returns(uint256)
func (_Token *TokenCaller) GetBuyPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "getBuyPrice")
	return *ret0, err
}

// GetBuyPrice is a free data retrieval call binding the contract method 0x018a25e8.
//
// Solidity: function getBuyPrice() constant returns(uint256)
func (_Token *TokenSession) GetBuyPrice() (*big.Int, error) {
	return _Token.Contract.GetBuyPrice(&_Token.CallOpts)
}

// GetBuyPrice is a free data retrieval call binding the contract method 0x018a25e8.
//
// Solidity: function getBuyPrice() constant returns(uint256)
func (_Token *TokenCallerSession) GetBuyPrice() (*big.Int, error) {
	return _Token.Contract.GetBuyPrice(&_Token.CallOpts)
}

// GetFundingTokens is a free data retrieval call binding the contract method 0xd079d3e0.
//
// Solidity: function getFundingTokens() constant returns(uint256)
func (_Token *TokenCaller) GetFundingTokens(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "getFundingTokens")
	return *ret0, err
}

// GetFundingTokens is a free data retrieval call binding the contract method 0xd079d3e0.
//
// Solidity: function getFundingTokens() constant returns(uint256)
func (_Token *TokenSession) GetFundingTokens() (*big.Int, error) {
	return _Token.Contract.GetFundingTokens(&_Token.CallOpts)
}

// GetFundingTokens is a free data retrieval call binding the contract method 0xd079d3e0.
//
// Solidity: function getFundingTokens() constant returns(uint256)
func (_Token *TokenCallerSession) GetFundingTokens() (*big.Int, error) {
	return _Token.Contract.GetFundingTokens(&_Token.CallOpts)
}

// GetFundingTokensForReserved is a free data retrieval call binding the contract method 0x7b62ace6.
//
// Solidity: function getFundingTokensForReserved() constant returns(uint256)
func (_Token *TokenCaller) GetFundingTokensForReserved(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "getFundingTokensForReserved")
	return *ret0, err
}

// GetFundingTokensForReserved is a free data retrieval call binding the contract method 0x7b62ace6.
//
// Solidity: function getFundingTokensForReserved() constant returns(uint256)
func (_Token *TokenSession) GetFundingTokensForReserved() (*big.Int, error) {
	return _Token.Contract.GetFundingTokensForReserved(&_Token.CallOpts)
}

// GetFundingTokensForReserved is a free data retrieval call binding the contract method 0x7b62ace6.
//
// Solidity: function getFundingTokensForReserved() constant returns(uint256)
func (_Token *TokenCallerSession) GetFundingTokensForReserved() (*big.Int, error) {
	return _Token.Contract.GetFundingTokensForReserved(&_Token.CallOpts)
}

// GetICORemainingBalance is a free data retrieval call binding the contract method 0x92cce438.
//
// Solidity: function getICORemainingBalance() constant returns(uint256)
func (_Token *TokenCaller) GetICORemainingBalance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "getICORemainingBalance")
	return *ret0, err
}

// GetICORemainingBalance is a free data retrieval call binding the contract method 0x92cce438.
//
// Solidity: function getICORemainingBalance() constant returns(uint256)
func (_Token *TokenSession) GetICORemainingBalance() (*big.Int, error) {
	return _Token.Contract.GetICORemainingBalance(&_Token.CallOpts)
}

// GetICORemainingBalance is a free data retrieval call binding the contract method 0x92cce438.
//
// Solidity: function getICORemainingBalance() constant returns(uint256)
func (_Token *TokenCallerSession) GetICORemainingBalance() (*big.Int, error) {
	return _Token.Contract.GetICORemainingBalance(&_Token.CallOpts)
}

// GetICOSupply is a free data retrieval call binding the contract method 0xa66fb7d7.
//
// Solidity: function getICOSupply() constant returns(uint256)
func (_Token *TokenCaller) GetICOSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "getICOSupply")
	return *ret0, err
}

// GetICOSupply is a free data retrieval call binding the contract method 0xa66fb7d7.
//
// Solidity: function getICOSupply() constant returns(uint256)
func (_Token *TokenSession) GetICOSupply() (*big.Int, error) {
	return _Token.Contract.GetICOSupply(&_Token.CallOpts)
}

// GetICOSupply is a free data retrieval call binding the contract method 0xa66fb7d7.
//
// Solidity: function getICOSupply() constant returns(uint256)
func (_Token *TokenCallerSession) GetICOSupply() (*big.Int, error) {
	return _Token.Contract.GetICOSupply(&_Token.CallOpts)
}

// GetPivateAddByEmailID is a free data retrieval call binding the contract method 0x4fb4cbde.
//
// Solidity: function getPivateAddByEmailID(email string) constant returns(address)
func (_Token *TokenCaller) GetPivateAddByEmailID(opts *bind.CallOpts, email string) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "getPivateAddByEmailID", email)
	return *ret0, err
}

// GetPivateAddByEmailID is a free data retrieval call binding the contract method 0x4fb4cbde.
//
// Solidity: function getPivateAddByEmailID(email string) constant returns(address)
func (_Token *TokenSession) GetPivateAddByEmailID(email string) (common.Address, error) {
	return _Token.Contract.GetPivateAddByEmailID(&_Token.CallOpts, email)
}

// GetPivateAddByEmailID is a free data retrieval call binding the contract method 0x4fb4cbde.
//
// Solidity: function getPivateAddByEmailID(email string) constant returns(address)
func (_Token *TokenCallerSession) GetPivateAddByEmailID(email string) (common.Address, error) {
	return _Token.Contract.GetPivateAddByEmailID(&_Token.CallOpts, email)
}

// GetPrivateAddByPublicAdd is a free data retrieval call binding the contract method 0x81ee5536.
//
// Solidity: function getPrivateAddByPublicAdd(publicAddr address) constant returns(address)
func (_Token *TokenCaller) GetPrivateAddByPublicAdd(opts *bind.CallOpts, publicAddr common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "getPrivateAddByPublicAdd", publicAddr)
	return *ret0, err
}

// GetPrivateAddByPublicAdd is a free data retrieval call binding the contract method 0x81ee5536.
//
// Solidity: function getPrivateAddByPublicAdd(publicAddr address) constant returns(address)
func (_Token *TokenSession) GetPrivateAddByPublicAdd(publicAddr common.Address) (common.Address, error) {
	return _Token.Contract.GetPrivateAddByPublicAdd(&_Token.CallOpts, publicAddr)
}

// GetPrivateAddByPublicAdd is a free data retrieval call binding the contract method 0x81ee5536.
//
// Solidity: function getPrivateAddByPublicAdd(publicAddr address) constant returns(address)
func (_Token *TokenCallerSession) GetPrivateAddByPublicAdd(publicAddr common.Address) (common.Address, error) {
	return _Token.Contract.GetPrivateAddByPublicAdd(&_Token.CallOpts, publicAddr)
}

// GetReservedCoins is a free data retrieval call binding the contract method 0x30e52113.
//
// Solidity: function getReservedCoins() constant returns(uint256)
func (_Token *TokenCaller) GetReservedCoins(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "getReservedCoins")
	return *ret0, err
}

// GetReservedCoins is a free data retrieval call binding the contract method 0x30e52113.
//
// Solidity: function getReservedCoins() constant returns(uint256)
func (_Token *TokenSession) GetReservedCoins() (*big.Int, error) {
	return _Token.Contract.GetReservedCoins(&_Token.CallOpts)
}

// GetReservedCoins is a free data retrieval call binding the contract method 0x30e52113.
//
// Solidity: function getReservedCoins() constant returns(uint256)
func (_Token *TokenCallerSession) GetReservedCoins() (*big.Int, error) {
	return _Token.Contract.GetReservedCoins(&_Token.CallOpts)
}

// GetReservedRemainingCoins is a free data retrieval call binding the contract method 0xfbefed74.
//
// Solidity: function getReservedRemainingCoins() constant returns(uint256)
func (_Token *TokenCaller) GetReservedRemainingCoins(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "getReservedRemainingCoins")
	return *ret0, err
}

// GetReservedRemainingCoins is a free data retrieval call binding the contract method 0xfbefed74.
//
// Solidity: function getReservedRemainingCoins() constant returns(uint256)
func (_Token *TokenSession) GetReservedRemainingCoins() (*big.Int, error) {
	return _Token.Contract.GetReservedRemainingCoins(&_Token.CallOpts)
}

// GetReservedRemainingCoins is a free data retrieval call binding the contract method 0xfbefed74.
//
// Solidity: function getReservedRemainingCoins() constant returns(uint256)
func (_Token *TokenCallerSession) GetReservedRemainingCoins() (*big.Int, error) {
	return _Token.Contract.GetReservedRemainingCoins(&_Token.CallOpts)
}

// GetSellPrice is a free data retrieval call binding the contract method 0x43d32e9c.
//
// Solidity: function getSellPrice() constant returns(uint256)
func (_Token *TokenCaller) GetSellPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "getSellPrice")
	return *ret0, err
}

// GetSellPrice is a free data retrieval call binding the contract method 0x43d32e9c.
//
// Solidity: function getSellPrice() constant returns(uint256)
func (_Token *TokenSession) GetSellPrice() (*big.Int, error) {
	return _Token.Contract.GetSellPrice(&_Token.CallOpts)
}

// GetSellPrice is a free data retrieval call binding the contract method 0x43d32e9c.
//
// Solidity: function getSellPrice() constant returns(uint256)
func (_Token *TokenCallerSession) GetSellPrice() (*big.Int, error) {
	return _Token.Contract.GetSellPrice(&_Token.CallOpts)
}

// GetTokenHolders is a free data retrieval call binding the contract method 0x876b1566.
//
// Solidity: function getTokenHolders() constant returns(noOfHolders uint256)
func (_Token *TokenCaller) GetTokenHolders(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "getTokenHolders")
	return *ret0, err
}

// GetTokenHolders is a free data retrieval call binding the contract method 0x876b1566.
//
// Solidity: function getTokenHolders() constant returns(noOfHolders uint256)
func (_Token *TokenSession) GetTokenHolders() (*big.Int, error) {
	return _Token.Contract.GetTokenHolders(&_Token.CallOpts)
}

// GetTokenHolders is a free data retrieval call binding the contract method 0x876b1566.
//
// Solidity: function getTokenHolders() constant returns(noOfHolders uint256)
func (_Token *TokenCallerSession) GetTokenHolders() (*big.Int, error) {
	return _Token.Contract.GetTokenHolders(&_Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Token *TokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Token *TokenSession) Name() (string, error) {
	return _Token.Contract.Name(&_Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Token *TokenCallerSession) Name() (string, error) {
	return _Token.Contract.Name(&_Token.CallOpts)
}

// PriceInUSD is a free data retrieval call binding the contract method 0xd41edb7b.
//
// Solidity: function priceInUSD() constant returns(uint256)
func (_Token *TokenCaller) PriceInUSD(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "priceInUSD")
	return *ret0, err
}

// PriceInUSD is a free data retrieval call binding the contract method 0xd41edb7b.
//
// Solidity: function priceInUSD() constant returns(uint256)
func (_Token *TokenSession) PriceInUSD() (*big.Int, error) {
	return _Token.Contract.PriceInUSD(&_Token.CallOpts)
}

// PriceInUSD is a free data retrieval call binding the contract method 0xd41edb7b.
//
// Solidity: function priceInUSD() constant returns(uint256)
func (_Token *TokenCallerSession) PriceInUSD() (*big.Int, error) {
	return _Token.Contract.PriceInUSD(&_Token.CallOpts)
}

// ReservedCoins is a free data retrieval call binding the contract method 0x18fd6440.
//
// Solidity: function reservedCoins() constant returns(uint256)
func (_Token *TokenCaller) ReservedCoins(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "reservedCoins")
	return *ret0, err
}

// ReservedCoins is a free data retrieval call binding the contract method 0x18fd6440.
//
// Solidity: function reservedCoins() constant returns(uint256)
func (_Token *TokenSession) ReservedCoins() (*big.Int, error) {
	return _Token.Contract.ReservedCoins(&_Token.CallOpts)
}

// ReservedCoins is a free data retrieval call binding the contract method 0x18fd6440.
//
// Solidity: function reservedCoins() constant returns(uint256)
func (_Token *TokenCallerSession) ReservedCoins() (*big.Int, error) {
	return _Token.Contract.ReservedCoins(&_Token.CallOpts)
}

// ReservedCoinsBalance is a free data retrieval call binding the contract method 0xa2a89c90.
//
// Solidity: function reservedCoinsBalance() constant returns(uint256)
func (_Token *TokenCaller) ReservedCoinsBalance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "reservedCoinsBalance")
	return *ret0, err
}

// ReservedCoinsBalance is a free data retrieval call binding the contract method 0xa2a89c90.
//
// Solidity: function reservedCoinsBalance() constant returns(uint256)
func (_Token *TokenSession) ReservedCoinsBalance() (*big.Int, error) {
	return _Token.Contract.ReservedCoinsBalance(&_Token.CallOpts)
}

// ReservedCoinsBalance is a free data retrieval call binding the contract method 0xa2a89c90.
//
// Solidity: function reservedCoinsBalance() constant returns(uint256)
func (_Token *TokenCallerSession) ReservedCoinsBalance() (*big.Int, error) {
	return _Token.Contract.ReservedCoinsBalance(&_Token.CallOpts)
}

// SellPriceInCentsForSingleToken is a free data retrieval call binding the contract method 0x80a903d0.
//
// Solidity: function sellPriceInCentsForSingleToken() constant returns(uint256)
func (_Token *TokenCaller) SellPriceInCentsForSingleToken(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "sellPriceInCentsForSingleToken")
	return *ret0, err
}

// SellPriceInCentsForSingleToken is a free data retrieval call binding the contract method 0x80a903d0.
//
// Solidity: function sellPriceInCentsForSingleToken() constant returns(uint256)
func (_Token *TokenSession) SellPriceInCentsForSingleToken() (*big.Int, error) {
	return _Token.Contract.SellPriceInCentsForSingleToken(&_Token.CallOpts)
}

// SellPriceInCentsForSingleToken is a free data retrieval call binding the contract method 0x80a903d0.
//
// Solidity: function sellPriceInCentsForSingleToken() constant returns(uint256)
func (_Token *TokenCallerSession) SellPriceInCentsForSingleToken() (*big.Int, error) {
	return _Token.Contract.SellPriceInCentsForSingleToken(&_Token.CallOpts)
}

// Standard is a free data retrieval call binding the contract method 0x5a3b7e42.
//
// Solidity: function standard() constant returns(string)
func (_Token *TokenCaller) Standard(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "standard")
	return *ret0, err
}

// Standard is a free data retrieval call binding the contract method 0x5a3b7e42.
//
// Solidity: function standard() constant returns(string)
func (_Token *TokenSession) Standard() (string, error) {
	return _Token.Contract.Standard(&_Token.CallOpts)
}

// Standard is a free data retrieval call binding the contract method 0x5a3b7e42.
//
// Solidity: function standard() constant returns(string)
func (_Token *TokenCallerSession) Standard() (string, error) {
	return _Token.Contract.Standard(&_Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Token *TokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Token *TokenSession) Symbol() (string, error) {
	return _Token.Contract.Symbol(&_Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Token *TokenCallerSession) Symbol() (string, error) {
	return _Token.Contract.Symbol(&_Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(initCoinSupply uint256)
func (_Token *TokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(initCoinSupply uint256)
func (_Token *TokenSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(initCoinSupply uint256)
func (_Token *TokenCallerSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// MapEmailIDWithPrivateAdd is a paid mutator transaction binding the contract method 0x290cc32a.
//
// Solidity: function MapEmailIDWithPrivateAdd(email string, privateAddress address) returns(bool)
func (_Token *TokenTransactor) MapEmailIDWithPrivateAdd(opts *bind.TransactOpts, email string, privateAddress common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "MapEmailIDWithPrivateAdd", email, privateAddress)
}

// MapEmailIDWithPrivateAdd is a paid mutator transaction binding the contract method 0x290cc32a.
//
// Solidity: function MapEmailIDWithPrivateAdd(email string, privateAddress address) returns(bool)
func (_Token *TokenSession) MapEmailIDWithPrivateAdd(email string, privateAddress common.Address) (*types.Transaction, error) {
	return _Token.Contract.MapEmailIDWithPrivateAdd(&_Token.TransactOpts, email, privateAddress)
}

// MapEmailIDWithPrivateAdd is a paid mutator transaction binding the contract method 0x290cc32a.
//
// Solidity: function MapEmailIDWithPrivateAdd(email string, privateAddress address) returns(bool)
func (_Token *TokenTransactorSession) MapEmailIDWithPrivateAdd(email string, privateAddress common.Address) (*types.Transaction, error) {
	return _Token.Contract.MapEmailIDWithPrivateAdd(&_Token.TransactOpts, email, privateAddress)
}

// SJTCwdSale is a paid mutator transaction binding the contract method 0x14daffba.
//
// Solidity: function SJTCwdSale() returns()
func (_Token *TokenTransactor) SJTCwdSale(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "SJTCwdSale")
}

// SJTCwdSale is a paid mutator transaction binding the contract method 0x14daffba.
//
// Solidity: function SJTCwdSale() returns()
func (_Token *TokenSession) SJTCwdSale() (*types.Transaction, error) {
	return _Token.Contract.SJTCwdSale(&_Token.TransactOpts)
}

// SJTCwdSale is a paid mutator transaction binding the contract method 0x14daffba.
//
// Solidity: function SJTCwdSale() returns()
func (_Token *TokenTransactorSession) SJTCwdSale() (*types.Transaction, error) {
	return _Token.Contract.SJTCwdSale(&_Token.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_sbalanceOfpender address, _value uint256) returns(success bool)
func (_Token *TokenTransactor) Approve(opts *bind.TransactOpts, _sbalanceOfpender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "approve", _sbalanceOfpender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_sbalanceOfpender address, _value uint256) returns(success bool)
func (_Token *TokenSession) Approve(_sbalanceOfpender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, _sbalanceOfpender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_sbalanceOfpender address, _value uint256) returns(success bool)
func (_Token *TokenTransactorSession) Approve(_sbalanceOfpender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, _sbalanceOfpender, _value)
}

// MintICOSupply is a paid mutator transaction binding the contract method 0x75355574.
//
// Solidity: function mintICOSupply(increasedCoins uint256) returns()
func (_Token *TokenTransactor) MintICOSupply(opts *bind.TransactOpts, increasedCoins *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "mintICOSupply", increasedCoins)
}

// MintICOSupply is a paid mutator transaction binding the contract method 0x75355574.
//
// Solidity: function mintICOSupply(increasedCoins uint256) returns()
func (_Token *TokenSession) MintICOSupply(increasedCoins *big.Int) (*types.Transaction, error) {
	return _Token.Contract.MintICOSupply(&_Token.TransactOpts, increasedCoins)
}

// MintICOSupply is a paid mutator transaction binding the contract method 0x75355574.
//
// Solidity: function mintICOSupply(increasedCoins uint256) returns()
func (_Token *TokenTransactorSession) MintICOSupply(increasedCoins *big.Int) (*types.Transaction, error) {
	return _Token.Contract.MintICOSupply(&_Token.TransactOpts, increasedCoins)
}

// MintToken is a paid mutator transaction binding the contract method 0xc634d032.
//
// Solidity: function mintToken(mintedAmount uint256) returns()
func (_Token *TokenTransactor) MintToken(opts *bind.TransactOpts, mintedAmount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "mintToken", mintedAmount)
}

// MintToken is a paid mutator transaction binding the contract method 0xc634d032.
//
// Solidity: function mintToken(mintedAmount uint256) returns()
func (_Token *TokenSession) MintToken(mintedAmount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.MintToken(&_Token.TransactOpts, mintedAmount)
}

// MintToken is a paid mutator transaction binding the contract method 0xc634d032.
//
// Solidity: function mintToken(mintedAmount uint256) returns()
func (_Token *TokenTransactorSession) MintToken(mintedAmount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.MintToken(&_Token.TransactOpts, mintedAmount)
}

// ReleaseTokensByUSD is a paid mutator transaction binding the contract method 0xc11c0832.
//
// Solidity: function releaseTokensByUSD(_sender address, fundingTokens uint256) returns()
func (_Token *TokenTransactor) ReleaseTokensByUSD(opts *bind.TransactOpts, _sender common.Address, fundingTokens *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "releaseTokensByUSD", _sender, fundingTokens)
}

// ReleaseTokensByUSD is a paid mutator transaction binding the contract method 0xc11c0832.
//
// Solidity: function releaseTokensByUSD(_sender address, fundingTokens uint256) returns()
func (_Token *TokenSession) ReleaseTokensByUSD(_sender common.Address, fundingTokens *big.Int) (*types.Transaction, error) {
	return _Token.Contract.ReleaseTokensByUSD(&_Token.TransactOpts, _sender, fundingTokens)
}

// ReleaseTokensByUSD is a paid mutator transaction binding the contract method 0xc11c0832.
//
// Solidity: function releaseTokensByUSD(_sender address, fundingTokens uint256) returns()
func (_Token *TokenTransactorSession) ReleaseTokensByUSD(_sender common.Address, fundingTokens *big.Int) (*types.Transaction, error) {
	return _Token.Contract.ReleaseTokensByUSD(&_Token.TransactOpts, _sender, fundingTokens)
}

// ReleaseTokensForICO is a paid mutator transaction binding the contract method 0x3caf2d79.
//
// Solidity: function releaseTokensForICO(_to address, _value uint256) returns(success bool)
func (_Token *TokenTransactor) ReleaseTokensForICO(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "releaseTokensForICO", _to, _value)
}

// ReleaseTokensForICO is a paid mutator transaction binding the contract method 0x3caf2d79.
//
// Solidity: function releaseTokensForICO(_to address, _value uint256) returns(success bool)
func (_Token *TokenSession) ReleaseTokensForICO(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.ReleaseTokensForICO(&_Token.TransactOpts, _to, _value)
}

// ReleaseTokensForICO is a paid mutator transaction binding the contract method 0x3caf2d79.
//
// Solidity: function releaseTokensForICO(_to address, _value uint256) returns(success bool)
func (_Token *TokenTransactorSession) ReleaseTokensForICO(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.ReleaseTokensForICO(&_Token.TransactOpts, _to, _value)
}

// ReleaseTokensFromReserved is a paid mutator transaction binding the contract method 0x67da2632.
//
// Solidity: function releaseTokensFromReserved(_to address, _value uint256) returns(success bool)
func (_Token *TokenTransactor) ReleaseTokensFromReserved(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "releaseTokensFromReserved", _to, _value)
}

// ReleaseTokensFromReserved is a paid mutator transaction binding the contract method 0x67da2632.
//
// Solidity: function releaseTokensFromReserved(_to address, _value uint256) returns(success bool)
func (_Token *TokenSession) ReleaseTokensFromReserved(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.ReleaseTokensFromReserved(&_Token.TransactOpts, _to, _value)
}

// ReleaseTokensFromReserved is a paid mutator transaction binding the contract method 0x67da2632.
//
// Solidity: function releaseTokensFromReserved(_to address, _value uint256) returns(success bool)
func (_Token *TokenTransactorSession) ReleaseTokensFromReserved(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.ReleaseTokensFromReserved(&_Token.TransactOpts, _to, _value)
}

// SetBuyPrice is a paid mutator transaction binding the contract method 0x63ae8d6c.
//
// Solidity: function setBuyPrice(newBuyPriceInCents uint256) returns()
func (_Token *TokenTransactor) SetBuyPrice(opts *bind.TransactOpts, newBuyPriceInCents *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setBuyPrice", newBuyPriceInCents)
}

// SetBuyPrice is a paid mutator transaction binding the contract method 0x63ae8d6c.
//
// Solidity: function setBuyPrice(newBuyPriceInCents uint256) returns()
func (_Token *TokenSession) SetBuyPrice(newBuyPriceInCents *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetBuyPrice(&_Token.TransactOpts, newBuyPriceInCents)
}

// SetBuyPrice is a paid mutator transaction binding the contract method 0x63ae8d6c.
//
// Solidity: function setBuyPrice(newBuyPriceInCents uint256) returns()
func (_Token *TokenTransactorSession) SetBuyPrice(newBuyPriceInCents *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetBuyPrice(&_Token.TransactOpts, newBuyPriceInCents)
}

// SetPrivateAddByPublicAdd is a paid mutator transaction binding the contract method 0x15f6f4d9.
//
// Solidity: function setPrivateAddByPublicAdd(publicAddress address, privateAddress address) returns(bool)
func (_Token *TokenTransactor) SetPrivateAddByPublicAdd(opts *bind.TransactOpts, publicAddress common.Address, privateAddress common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setPrivateAddByPublicAdd", publicAddress, privateAddress)
}

// SetPrivateAddByPublicAdd is a paid mutator transaction binding the contract method 0x15f6f4d9.
//
// Solidity: function setPrivateAddByPublicAdd(publicAddress address, privateAddress address) returns(bool)
func (_Token *TokenSession) SetPrivateAddByPublicAdd(publicAddress common.Address, privateAddress common.Address) (*types.Transaction, error) {
	return _Token.Contract.SetPrivateAddByPublicAdd(&_Token.TransactOpts, publicAddress, privateAddress)
}

// SetPrivateAddByPublicAdd is a paid mutator transaction binding the contract method 0x15f6f4d9.
//
// Solidity: function setPrivateAddByPublicAdd(publicAddress address, privateAddress address) returns(bool)
func (_Token *TokenTransactorSession) SetPrivateAddByPublicAdd(publicAddress common.Address, privateAddress common.Address) (*types.Transaction, error) {
	return _Token.Contract.SetPrivateAddByPublicAdd(&_Token.TransactOpts, publicAddress, privateAddress)
}

// SetSellPrice is a paid mutator transaction binding the contract method 0xfc6634b9.
//
// Solidity: function setSellPrice(newSellPriceInCents uint256) returns()
func (_Token *TokenTransactor) SetSellPrice(opts *bind.TransactOpts, newSellPriceInCents *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setSellPrice", newSellPriceInCents)
}

// SetSellPrice is a paid mutator transaction binding the contract method 0xfc6634b9.
//
// Solidity: function setSellPrice(newSellPriceInCents uint256) returns()
func (_Token *TokenSession) SetSellPrice(newSellPriceInCents *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetSellPrice(&_Token.TransactOpts, newSellPriceInCents)
}

// SetSellPrice is a paid mutator transaction binding the contract method 0xfc6634b9.
//
// Solidity: function setSellPrice(newSellPriceInCents uint256) returns()
func (_Token *TokenTransactorSession) SetSellPrice(newSellPriceInCents *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetSellPrice(&_Token.TransactOpts, newSellPriceInCents)
}

// StringToBytes32 is a paid mutator transaction binding the contract method 0xcfb51928.
//
// Solidity: function stringToBytes32(source string) returns(result bytes32)
func (_Token *TokenTransactor) StringToBytes32(opts *bind.TransactOpts, source string) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "stringToBytes32", source)
}

// StringToBytes32 is a paid mutator transaction binding the contract method 0xcfb51928.
//
// Solidity: function stringToBytes32(source string) returns(result bytes32)
func (_Token *TokenSession) StringToBytes32(source string) (*types.Transaction, error) {
	return _Token.Contract.StringToBytes32(&_Token.TransactOpts, source)
}

// StringToBytes32 is a paid mutator transaction binding the contract method 0xcfb51928.
//
// Solidity: function stringToBytes32(source string) returns(result bytes32)
func (_Token *TokenTransactorSession) StringToBytes32(source string) (*types.Transaction, error) {
	return _Token.Contract.StringToBytes32(&_Token.TransactOpts, source)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_Token *TokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_Token *TokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_Token *TokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_Token *TokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_Token *TokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_Token *TokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, _from, _to, _value)
}

// TransferTokens is a paid mutator transaction binding the contract method 0xa64b6e5f.
//
// Solidity: function transferTokens(_from address, _to address, _value uint256) returns(success bool)
func (_Token *TokenTransactor) TransferTokens(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transferTokens", _from, _to, _value)
}

// TransferTokens is a paid mutator transaction binding the contract method 0xa64b6e5f.
//
// Solidity: function transferTokens(_from address, _to address, _value uint256) returns(success bool)
func (_Token *TokenSession) TransferTokens(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferTokens(&_Token.TransactOpts, _from, _to, _value)
}

// TransferTokens is a paid mutator transaction binding the contract method 0xa64b6e5f.
//
// Solidity: function transferTokens(_from address, _to address, _value uint256) returns(success bool)
func (_Token *TokenTransactorSession) TransferTokens(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferTokens(&_Token.TransactOpts, _from, _to, _value)
}
