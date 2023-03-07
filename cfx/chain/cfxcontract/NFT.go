// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cfxcontract

import (
	"math/big"
	"strings"

	"github.com/Conflux-Chain/go-conflux-sdk/bind"
	"github.com/Conflux-Chain/go-conflux-sdk/cfxclient/bulk"

	types "github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethBind "github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = ethBind.Bind
	_ = common.Big1
	_ = ethtypes.BloomLookup
	_ = event.NewSubscription
)

// MetaderbyHorseNFTTokenABI is the input ABI used to generate the binding from.
const MetaderbyHorseNFTTokenABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_baseUrl\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sign\",\"type\":\"bytes\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mintForUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setBackendAddress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseUrl\",\"type\":\"string\"}],\"name\":\"setBaseUrl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sign\",\"type\":\"bytes\"}],\"name\":\"validUser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// MetaderbyHorseNFTToken is an auto generated Go binding around an Conflux contract.
type MetaderbyHorseNFTToken struct {
	MetaderbyHorseNFTTokenCaller     // Read-only binding to the contract
	MetaderbyHorseNFTTokenTransactor // Write-only binding to the contract
	MetaderbyHorseNFTTokenFilterer   // Log filterer for contract events
}

// MetaderbyHorseNFTTokenCaller is an auto generated read-only Go binding around an Conflux contract.
type MetaderbyHorseNFTTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetaderbyHorseNFTTokenBulkCaller is an auto generated read-only Go binding around an Conflux contract.
type MetaderbyHorseNFTTokenBulkCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetaderbyHorseNFTTokenTransactor is an auto generated write-only Go binding around an Conflux contract.
type MetaderbyHorseNFTTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetaderbyHorseNFTTokenBulkTransactor is an auto generated write-only Go binding around an Conflux contract.
type MetaderbyHorseNFTTokenBulkTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetaderbyHorseNFTTokenFilterer is an auto generated log filtering Go binding around an Conflux contract events.
type MetaderbyHorseNFTTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetaderbyHorseNFTTokenSession is an auto generated Go binding around an Conflux contract,
// with pre-set call and transact options.
type MetaderbyHorseNFTTokenSession struct {
	Contract     *MetaderbyHorseNFTToken // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// MetaderbyHorseNFTTokenCallerSession is an auto generated read-only Go binding around an Conflux contract,
// with pre-set call options.
type MetaderbyHorseNFTTokenCallerSession struct {
	Contract *MetaderbyHorseNFTTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// MetaderbyHorseNFTTokenTransactorSession is an auto generated write-only Go binding around an Conflux contract,
// with pre-set transact options.
type MetaderbyHorseNFTTokenTransactorSession struct {
	Contract     *MetaderbyHorseNFTTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// MetaderbyHorseNFTTokenRaw is an auto generated low-level Go binding around an Conflux contract.
type MetaderbyHorseNFTTokenRaw struct {
	Contract *MetaderbyHorseNFTToken // Generic contract binding to access the raw methods on
}

// MetaderbyHorseNFTTokenCallerRaw is an auto generated low-level read-only Go binding around an Conflux contract.
type MetaderbyHorseNFTTokenCallerRaw struct {
	Contract *MetaderbyHorseNFTTokenCaller // Generic read-only contract binding to access the raw methods on
}

// MetaderbyHorseNFTTokenTransactorRaw is an auto generated low-level write-only Go binding around an Conflux contract.
type MetaderbyHorseNFTTokenTransactorRaw struct {
	Contract *MetaderbyHorseNFTTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMetaderbyHorseNFTToken creates a new instance of MetaderbyHorseNFTToken, bound to a specific deployed contract.
func NewMetaderbyHorseNFTToken(address types.Address, backend bind.ContractBackend) (*MetaderbyHorseNFTToken, error) {
	contract, err := bindMetaderbyHorseNFTToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MetaderbyHorseNFTToken{MetaderbyHorseNFTTokenCaller: MetaderbyHorseNFTTokenCaller{contract: contract}, MetaderbyHorseNFTTokenTransactor: MetaderbyHorseNFTTokenTransactor{contract: contract}, MetaderbyHorseNFTTokenFilterer: MetaderbyHorseNFTTokenFilterer{contract: contract}}, nil
}

// NewMetaderbyHorseNFTTokenCaller creates a new read-only instance of MetaderbyHorseNFTToken, bound to a specific deployed contract.
func NewMetaderbyHorseNFTTokenCaller(address types.Address, caller bind.ContractCaller) (*MetaderbyHorseNFTTokenCaller, error) {
	contract, err := bindMetaderbyHorseNFTToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MetaderbyHorseNFTTokenCaller{contract: contract}, nil
}

// NewMetaderbyHorseNFTTokenTransactor creates a new write-only instance of MetaderbyHorseNFTToken, bound to a specific deployed contract.
func NewMetaderbyHorseNFTTokenTransactor(address types.Address, transactor bind.ContractTransactor) (*MetaderbyHorseNFTTokenTransactor, error) {
	contract, err := bindMetaderbyHorseNFTToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MetaderbyHorseNFTTokenTransactor{contract: contract}, nil
}

// NewMetaderbyHorseNFTTokenFilterer creates a new log filterer instance of MetaderbyHorseNFTToken, bound to a specific deployed contract.
func NewMetaderbyHorseNFTTokenFilterer(address types.Address, filterer bind.ContractFilterer) (*MetaderbyHorseNFTTokenFilterer, error) {
	contract, err := bindMetaderbyHorseNFTToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MetaderbyHorseNFTTokenFilterer{contract: contract}, nil
}

// NewMetaderbyHorseNFTTokenCaller creates a new read-only instance of MetaderbyHorseNFTToken, bound to a specific deployed contract.
func NewMetaderbyHorseNFTTokenBulkCaller(address types.Address, caller bind.ContractCaller) (*MetaderbyHorseNFTTokenBulkCaller, error) {
	contract, err := bindMetaderbyHorseNFTToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MetaderbyHorseNFTTokenBulkCaller{contract: contract}, nil
}

// NewMetaderbyHorseNFTTokenBulkTransactor creates a new write-only instance of MetaderbyHorseNFTToken, bound to a specific deployed contract.
func NewMetaderbyHorseNFTTokenBulkTransactor(address types.Address, transactor bind.ContractTransactor) (*MetaderbyHorseNFTTokenBulkTransactor, error) {
	contract, err := bindMetaderbyHorseNFTToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MetaderbyHorseNFTTokenBulkTransactor{contract: contract}, nil
}

// bindMetaderbyHorseNFTToken binds a generic wrapper to an already deployed contract.
func bindMetaderbyHorseNFTToken(address types.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MetaderbyHorseNFTTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MetaderbyHorseNFTToken.Contract.MetaderbyHorseNFTTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenRaw) Transfer(opts *bind.TransactOpts) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MetaderbyHorseNFTToken.Contract.MetaderbyHorseNFTTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MetaderbyHorseNFTToken.Contract.MetaderbyHorseNFTTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MetaderbyHorseNFTToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MetaderbyHorseNFTToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MetaderbyHorseNFTToken.Contract.contract.Transact(opts, method, params...)
}

// BaseUrl is a free data retrieval call binding the contract method 0xf34a39d4.
//
// Solidity: function _baseUrl() view returns(string)
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenCaller) BaseUrl(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	__err := _MetaderbyHorseNFTToken.contract.Call(opts, &out, "_baseUrl")

	if __err != nil {
		return *new(string), __err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, __err

}

// BaseUrl is a free data retrieval call binding the contract method 0xf34a39d4.
//
// Solidity: function _baseUrl() view returns(string)
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenBulkCaller) BaseUrl(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*string, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _MetaderbyHorseNFTToken.contract.GenRequest(opts, "_baseUrl")

	out0 := new(string)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _MetaderbyHorseNFTToken.contract.DecodeOutput(&out, rawOut, "_baseUrl")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(string)).(*string)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// BaseUrl is a free data retrieval call binding the contract method 0xf34a39d4.
//
// Solidity: function _baseUrl() view returns(string)
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenSession) BaseUrl() (string, error) {
	return _MetaderbyHorseNFTToken.Contract.BaseUrl(&_MetaderbyHorseNFTToken.CallOpts)
}

// BaseUrl is a free data retrieval call binding the contract method 0xf34a39d4.
//
// Solidity: function _baseUrl() view returns(string)
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenCallerSession) BaseUrl() (string, error) {
	return _MetaderbyHorseNFTToken.Contract.BaseUrl(&_MetaderbyHorseNFTToken.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	__err := _MetaderbyHorseNFTToken.contract.Call(opts, &out, "balanceOf", owner)

	if __err != nil {
		return *new(*big.Int), __err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, __err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenBulkCaller) BalanceOf(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, owner common.Address) (**big.Int, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _MetaderbyHorseNFTToken.contract.GenRequest(opts, "balanceOf", owner)

	out0 := new(*big.Int)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _MetaderbyHorseNFTToken.contract.DecodeOutput(&out, rawOut, "balanceOf")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _MetaderbyHorseNFTToken.Contract.BalanceOf(&_MetaderbyHorseNFTToken.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _MetaderbyHorseNFTToken.Contract.BalanceOf(&_MetaderbyHorseNFTToken.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	__err := _MetaderbyHorseNFTToken.contract.Call(opts, &out, "getApproved", tokenId)

	if __err != nil {
		return *new(common.Address), __err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, __err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenBulkCaller) GetApproved(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, tokenId *big.Int) (*common.Address, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _MetaderbyHorseNFTToken.contract.GenRequest(opts, "getApproved", tokenId)

	out0 := new(common.Address)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _MetaderbyHorseNFTToken.contract.DecodeOutput(&out, rawOut, "getApproved")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _MetaderbyHorseNFTToken.Contract.GetApproved(&_MetaderbyHorseNFTToken.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _MetaderbyHorseNFTToken.Contract.GetApproved(&_MetaderbyHorseNFTToken.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	__err := _MetaderbyHorseNFTToken.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if __err != nil {
		return *new(bool), __err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, __err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenBulkCaller) IsApprovedForAll(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, owner common.Address, operator common.Address) (*bool, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _MetaderbyHorseNFTToken.contract.GenRequest(opts, "isApprovedForAll", owner, operator)

	out0 := new(bool)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _MetaderbyHorseNFTToken.contract.DecodeOutput(&out, rawOut, "isApprovedForAll")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(bool)).(*bool)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _MetaderbyHorseNFTToken.Contract.IsApprovedForAll(&_MetaderbyHorseNFTToken.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _MetaderbyHorseNFTToken.Contract.IsApprovedForAll(&_MetaderbyHorseNFTToken.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	__err := _MetaderbyHorseNFTToken.contract.Call(opts, &out, "name")

	if __err != nil {
		return *new(string), __err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, __err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenBulkCaller) Name(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*string, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _MetaderbyHorseNFTToken.contract.GenRequest(opts, "name")

	out0 := new(string)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _MetaderbyHorseNFTToken.contract.DecodeOutput(&out, rawOut, "name")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(string)).(*string)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenSession) Name() (string, error) {
	return _MetaderbyHorseNFTToken.Contract.Name(&_MetaderbyHorseNFTToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenCallerSession) Name() (string, error) {
	return _MetaderbyHorseNFTToken.Contract.Name(&_MetaderbyHorseNFTToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	__err := _MetaderbyHorseNFTToken.contract.Call(opts, &out, "owner")

	if __err != nil {
		return *new(common.Address), __err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, __err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenBulkCaller) Owner(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*common.Address, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _MetaderbyHorseNFTToken.contract.GenRequest(opts, "owner")

	out0 := new(common.Address)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _MetaderbyHorseNFTToken.contract.DecodeOutput(&out, rawOut, "owner")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenSession) Owner() (common.Address, error) {
	return _MetaderbyHorseNFTToken.Contract.Owner(&_MetaderbyHorseNFTToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenCallerSession) Owner() (common.Address, error) {
	return _MetaderbyHorseNFTToken.Contract.Owner(&_MetaderbyHorseNFTToken.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	__err := _MetaderbyHorseNFTToken.contract.Call(opts, &out, "ownerOf", tokenId)

	if __err != nil {
		return *new(common.Address), __err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, __err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenBulkCaller) OwnerOf(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, tokenId *big.Int) (*common.Address, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _MetaderbyHorseNFTToken.contract.GenRequest(opts, "ownerOf", tokenId)

	out0 := new(common.Address)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _MetaderbyHorseNFTToken.contract.DecodeOutput(&out, rawOut, "ownerOf")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _MetaderbyHorseNFTToken.Contract.OwnerOf(&_MetaderbyHorseNFTToken.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _MetaderbyHorseNFTToken.Contract.OwnerOf(&_MetaderbyHorseNFTToken.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	__err := _MetaderbyHorseNFTToken.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if __err != nil {
		return *new(bool), __err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, __err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenBulkCaller) SupportsInterface(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, interfaceId [4]byte) (*bool, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _MetaderbyHorseNFTToken.contract.GenRequest(opts, "supportsInterface", interfaceId)

	out0 := new(bool)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _MetaderbyHorseNFTToken.contract.DecodeOutput(&out, rawOut, "supportsInterface")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(bool)).(*bool)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MetaderbyHorseNFTToken.Contract.SupportsInterface(&_MetaderbyHorseNFTToken.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MetaderbyHorseNFTToken.Contract.SupportsInterface(&_MetaderbyHorseNFTToken.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	__err := _MetaderbyHorseNFTToken.contract.Call(opts, &out, "symbol")

	if __err != nil {
		return *new(string), __err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, __err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenBulkCaller) Symbol(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*string, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _MetaderbyHorseNFTToken.contract.GenRequest(opts, "symbol")

	out0 := new(string)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _MetaderbyHorseNFTToken.contract.DecodeOutput(&out, rawOut, "symbol")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(string)).(*string)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenSession) Symbol() (string, error) {
	return _MetaderbyHorseNFTToken.Contract.Symbol(&_MetaderbyHorseNFTToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenCallerSession) Symbol() (string, error) {
	return _MetaderbyHorseNFTToken.Contract.Symbol(&_MetaderbyHorseNFTToken.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	__err := _MetaderbyHorseNFTToken.contract.Call(opts, &out, "tokenURI", tokenId)

	if __err != nil {
		return *new(string), __err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, __err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenBulkCaller) TokenURI(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, tokenId *big.Int) (*string, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _MetaderbyHorseNFTToken.contract.GenRequest(opts, "tokenURI", tokenId)

	out0 := new(string)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _MetaderbyHorseNFTToken.contract.DecodeOutput(&out, rawOut, "tokenURI")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(string)).(*string)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenSession) TokenURI(tokenId *big.Int) (string, error) {
	return _MetaderbyHorseNFTToken.Contract.TokenURI(&_MetaderbyHorseNFTToken.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _MetaderbyHorseNFTToken.Contract.TokenURI(&_MetaderbyHorseNFTToken.CallOpts, tokenId)
}

// ValidUser is a free data retrieval call binding the contract method 0x58870770.
//
// Solidity: function validUser(uint256 num, uint256 timestamp, uint256 amount, bytes sign) view returns(bool)
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenCaller) ValidUser(opts *bind.CallOpts, num *big.Int, timestamp *big.Int, amount *big.Int, sign []byte) (bool, error) {
	var out []interface{}
	__err := _MetaderbyHorseNFTToken.contract.Call(opts, &out, "validUser", num, timestamp, amount, sign)

	if __err != nil {
		return *new(bool), __err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, __err

}

// ValidUser is a free data retrieval call binding the contract method 0x58870770.
//
// Solidity: function validUser(uint256 num, uint256 timestamp, uint256 amount, bytes sign) view returns(bool)
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenBulkCaller) ValidUser(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, num *big.Int, timestamp *big.Int, amount *big.Int, sign []byte) (*bool, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _MetaderbyHorseNFTToken.contract.GenRequest(opts, "validUser", num, timestamp, amount, sign)

	out0 := new(bool)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _MetaderbyHorseNFTToken.contract.DecodeOutput(&out, rawOut, "validUser")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(bool)).(*bool)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// ValidUser is a free data retrieval call binding the contract method 0x58870770.
//
// Solidity: function validUser(uint256 num, uint256 timestamp, uint256 amount, bytes sign) view returns(bool)
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenSession) ValidUser(num *big.Int, timestamp *big.Int, amount *big.Int, sign []byte) (bool, error) {
	return _MetaderbyHorseNFTToken.Contract.ValidUser(&_MetaderbyHorseNFTToken.CallOpts, num, timestamp, amount, sign)
}

// ValidUser is a free data retrieval call binding the contract method 0x58870770.
//
// Solidity: function validUser(uint256 num, uint256 timestamp, uint256 amount, bytes sign) view returns(bool)
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenCallerSession) ValidUser(num *big.Int, timestamp *big.Int, amount *big.Int, sign []byte) (bool, error) {
	return _MetaderbyHorseNFTToken.Contract.ValidUser(&_MetaderbyHorseNFTToken.CallOpts, num, timestamp, amount, sign)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MetaderbyHorseNFTToken.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenBulkTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) types.UnsignedTransaction {
	return _MetaderbyHorseNFTToken.contract.GenUnsignedTransaction(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenSession) Approve(to common.Address, tokenId *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MetaderbyHorseNFTToken.Contract.Approve(&_MetaderbyHorseNFTToken.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MetaderbyHorseNFTToken.Contract.Approve(&_MetaderbyHorseNFTToken.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MetaderbyHorseNFTToken.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenBulkTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) types.UnsignedTransaction {
	return _MetaderbyHorseNFTToken.contract.GenUnsignedTransaction(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenSession) Burn(tokenId *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MetaderbyHorseNFTToken.Contract.Burn(&_MetaderbyHorseNFTToken.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenTransactorSession) Burn(tokenId *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MetaderbyHorseNFTToken.Contract.Burn(&_MetaderbyHorseNFTToken.TransactOpts, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x4a9eee69.
//
// Solidity: function mint(uint256 num, uint256 timestamp, uint256 amount, bytes sign) payable returns()
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenTransactor) Mint(opts *bind.TransactOpts, num *big.Int, timestamp *big.Int, amount *big.Int, sign []byte) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MetaderbyHorseNFTToken.contract.Transact(opts, "mint", num, timestamp, amount, sign)
}

// Mint is a paid mutator transaction binding the contract method 0x4a9eee69.
//
// Solidity: function mint(uint256 num, uint256 timestamp, uint256 amount, bytes sign) payable returns()
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenBulkTransactor) Mint(opts *bind.TransactOpts, num *big.Int, timestamp *big.Int, amount *big.Int, sign []byte) types.UnsignedTransaction {
	return _MetaderbyHorseNFTToken.contract.GenUnsignedTransaction(opts, "mint", num, timestamp, amount, sign)
}

// Mint is a paid mutator transaction binding the contract method 0x4a9eee69.
//
// Solidity: function mint(uint256 num, uint256 timestamp, uint256 amount, bytes sign) payable returns()
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenSession) Mint(num *big.Int, timestamp *big.Int, amount *big.Int, sign []byte) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MetaderbyHorseNFTToken.Contract.Mint(&_MetaderbyHorseNFTToken.TransactOpts, num, timestamp, amount, sign)
}

// Mint is a paid mutator transaction binding the contract method 0x4a9eee69.
//
// Solidity: function mint(uint256 num, uint256 timestamp, uint256 amount, bytes sign) payable returns()
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenTransactorSession) Mint(num *big.Int, timestamp *big.Int, amount *big.Int, sign []byte) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MetaderbyHorseNFTToken.Contract.Mint(&_MetaderbyHorseNFTToken.TransactOpts, num, timestamp, amount, sign)
}

// MintForUser is a paid mutator transaction binding the contract method 0x1746b220.
//
// Solidity: function mintForUser(uint256 num, address to) returns()
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenTransactor) MintForUser(opts *bind.TransactOpts, num *big.Int, to common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MetaderbyHorseNFTToken.contract.Transact(opts, "mintForUser", num, to)
}

// MintForUser is a paid mutator transaction binding the contract method 0x1746b220.
//
// Solidity: function mintForUser(uint256 num, address to) returns()
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenBulkTransactor) MintForUser(opts *bind.TransactOpts, num *big.Int, to common.Address) types.UnsignedTransaction {
	return _MetaderbyHorseNFTToken.contract.GenUnsignedTransaction(opts, "mintForUser", num, to)
}

// MintForUser is a paid mutator transaction binding the contract method 0x1746b220.
//
// Solidity: function mintForUser(uint256 num, address to) returns()
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenSession) MintForUser(num *big.Int, to common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MetaderbyHorseNFTToken.Contract.MintForUser(&_MetaderbyHorseNFTToken.TransactOpts, num, to)
}

// MintForUser is a paid mutator transaction binding the contract method 0x1746b220.
//
// Solidity: function mintForUser(uint256 num, address to) returns()
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenTransactorSession) MintForUser(num *big.Int, to common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MetaderbyHorseNFTToken.Contract.MintForUser(&_MetaderbyHorseNFTToken.TransactOpts, num, to)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MetaderbyHorseNFTToken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenBulkTransactor) RenounceOwnership(opts *bind.TransactOpts) types.UnsignedTransaction {
	return _MetaderbyHorseNFTToken.contract.GenUnsignedTransaction(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenSession) RenounceOwnership() (*types.UnsignedTransaction, *types.Hash, error) {
	return _MetaderbyHorseNFTToken.Contract.RenounceOwnership(&_MetaderbyHorseNFTToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenTransactorSession) RenounceOwnership() (*types.UnsignedTransaction, *types.Hash, error) {
	return _MetaderbyHorseNFTToken.Contract.RenounceOwnership(&_MetaderbyHorseNFTToken.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MetaderbyHorseNFTToken.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenBulkTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) types.UnsignedTransaction {
	return _MetaderbyHorseNFTToken.contract.GenUnsignedTransaction(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MetaderbyHorseNFTToken.Contract.SafeTransferFrom(&_MetaderbyHorseNFTToken.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MetaderbyHorseNFTToken.Contract.SafeTransferFrom(&_MetaderbyHorseNFTToken.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MetaderbyHorseNFTToken.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenBulkTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) types.UnsignedTransaction {
	return _MetaderbyHorseNFTToken.contract.GenUnsignedTransaction(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MetaderbyHorseNFTToken.Contract.SafeTransferFrom0(&_MetaderbyHorseNFTToken.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MetaderbyHorseNFTToken.Contract.SafeTransferFrom0(&_MetaderbyHorseNFTToken.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MetaderbyHorseNFTToken.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenBulkTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) types.UnsignedTransaction {
	return _MetaderbyHorseNFTToken.contract.GenUnsignedTransaction(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenSession) SetApprovalForAll(operator common.Address, approved bool) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MetaderbyHorseNFTToken.Contract.SetApprovalForAll(&_MetaderbyHorseNFTToken.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MetaderbyHorseNFTToken.Contract.SetApprovalForAll(&_MetaderbyHorseNFTToken.TransactOpts, operator, approved)
}

// SetBackendAddress is a paid mutator transaction binding the contract method 0x1815ce7d.
//
// Solidity: function setBackendAddress(address addr) returns(bool)
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenTransactor) SetBackendAddress(opts *bind.TransactOpts, addr common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MetaderbyHorseNFTToken.contract.Transact(opts, "setBackendAddress", addr)
}

// SetBackendAddress is a paid mutator transaction binding the contract method 0x1815ce7d.
//
// Solidity: function setBackendAddress(address addr) returns(bool)
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenBulkTransactor) SetBackendAddress(opts *bind.TransactOpts, addr common.Address) types.UnsignedTransaction {
	return _MetaderbyHorseNFTToken.contract.GenUnsignedTransaction(opts, "setBackendAddress", addr)
}

// SetBackendAddress is a paid mutator transaction binding the contract method 0x1815ce7d.
//
// Solidity: function setBackendAddress(address addr) returns(bool)
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenSession) SetBackendAddress(addr common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MetaderbyHorseNFTToken.Contract.SetBackendAddress(&_MetaderbyHorseNFTToken.TransactOpts, addr)
}

// SetBackendAddress is a paid mutator transaction binding the contract method 0x1815ce7d.
//
// Solidity: function setBackendAddress(address addr) returns(bool)
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenTransactorSession) SetBackendAddress(addr common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MetaderbyHorseNFTToken.Contract.SetBackendAddress(&_MetaderbyHorseNFTToken.TransactOpts, addr)
}

// SetBaseUrl is a paid mutator transaction binding the contract method 0xc7c3268b.
//
// Solidity: function setBaseUrl(string baseUrl) returns()
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenTransactor) SetBaseUrl(opts *bind.TransactOpts, baseUrl string) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MetaderbyHorseNFTToken.contract.Transact(opts, "setBaseUrl", baseUrl)
}

// SetBaseUrl is a paid mutator transaction binding the contract method 0xc7c3268b.
//
// Solidity: function setBaseUrl(string baseUrl) returns()
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenBulkTransactor) SetBaseUrl(opts *bind.TransactOpts, baseUrl string) types.UnsignedTransaction {
	return _MetaderbyHorseNFTToken.contract.GenUnsignedTransaction(opts, "setBaseUrl", baseUrl)
}

// SetBaseUrl is a paid mutator transaction binding the contract method 0xc7c3268b.
//
// Solidity: function setBaseUrl(string baseUrl) returns()
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenSession) SetBaseUrl(baseUrl string) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MetaderbyHorseNFTToken.Contract.SetBaseUrl(&_MetaderbyHorseNFTToken.TransactOpts, baseUrl)
}

// SetBaseUrl is a paid mutator transaction binding the contract method 0xc7c3268b.
//
// Solidity: function setBaseUrl(string baseUrl) returns()
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenTransactorSession) SetBaseUrl(baseUrl string) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MetaderbyHorseNFTToken.Contract.SetBaseUrl(&_MetaderbyHorseNFTToken.TransactOpts, baseUrl)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MetaderbyHorseNFTToken.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenBulkTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) types.UnsignedTransaction {
	return _MetaderbyHorseNFTToken.contract.GenUnsignedTransaction(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MetaderbyHorseNFTToken.Contract.TransferFrom(&_MetaderbyHorseNFTToken.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MetaderbyHorseNFTToken.Contract.TransferFrom(&_MetaderbyHorseNFTToken.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MetaderbyHorseNFTToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenBulkTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) types.UnsignedTransaction {
	return _MetaderbyHorseNFTToken.contract.GenUnsignedTransaction(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenSession) TransferOwnership(newOwner common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MetaderbyHorseNFTToken.Contract.TransferOwnership(&_MetaderbyHorseNFTToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MetaderbyHorseNFTToken.Contract.TransferOwnership(&_MetaderbyHorseNFTToken.TransactOpts, newOwner)
}

// MetaderbyHorseNFTTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MetaderbyHorseNFTToken contract.
type MetaderbyHorseNFTTokenApprovalIterator struct {
	Event *MetaderbyHorseNFTTokenApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MetaderbyHorseNFTTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaderbyHorseNFTTokenApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	if it.sub == nil {
		it.done = true
		return it.Next()
	}

	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MetaderbyHorseNFTTokenApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MetaderbyHorseNFTTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaderbyHorseNFTTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaderbyHorseNFTTokenApproval represents a Approval event raised by the MetaderbyHorseNFTToken contract.
type MetaderbyHorseNFTTokenApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// MetaderbyHorseNFTTokenApprovalOrChainReorg represents a Approval subscription event raised by the MetaderbyHorseNFTToken contract.
type MetaderbyHorseNFTTokenApprovalOrChainReorg struct {
	Event      *MetaderbyHorseNFTTokenApproval
	ChainReorg *types.ChainReorg
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*MetaderbyHorseNFTTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, err := _MetaderbyHorseNFTToken.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MetaderbyHorseNFTTokenApprovalIterator{contract: _MetaderbyHorseNFTToken.contract, event: "Approval", logs: logs}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MetaderbyHorseNFTTokenApprovalOrChainReorg, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _MetaderbyHorseNFTToken.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaderbyHorseNFTTokenApprovalOrChainReorg)
				event.Event = new(MetaderbyHorseNFTTokenApproval)

				if log.ChainReorg == nil {
					if err := _MetaderbyHorseNFTToken.contract.UnpackLog(event.Event, "Approval", *log.Log); err != nil {
						return err
					}
					event.Event.Raw = *log.Log
				} else {
					event.ChainReorg = log.ChainReorg
				}

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenFilterer) ParseApproval(log types.Log) (*MetaderbyHorseNFTTokenApproval, error) {
	event := new(MetaderbyHorseNFTTokenApproval)
	if err := _MetaderbyHorseNFTToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaderbyHorseNFTTokenApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the MetaderbyHorseNFTToken contract.
type MetaderbyHorseNFTTokenApprovalForAllIterator struct {
	Event *MetaderbyHorseNFTTokenApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MetaderbyHorseNFTTokenApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaderbyHorseNFTTokenApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	if it.sub == nil {
		it.done = true
		return it.Next()
	}

	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MetaderbyHorseNFTTokenApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MetaderbyHorseNFTTokenApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaderbyHorseNFTTokenApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaderbyHorseNFTTokenApprovalForAll represents a ApprovalForAll event raised by the MetaderbyHorseNFTToken contract.
type MetaderbyHorseNFTTokenApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// MetaderbyHorseNFTTokenApprovalForAllOrChainReorg represents a ApprovalForAll subscription event raised by the MetaderbyHorseNFTToken contract.
type MetaderbyHorseNFTTokenApprovalForAllOrChainReorg struct {
	Event      *MetaderbyHorseNFTTokenApprovalForAll
	ChainReorg *types.ChainReorg
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*MetaderbyHorseNFTTokenApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, err := _MetaderbyHorseNFTToken.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &MetaderbyHorseNFTTokenApprovalForAllIterator{contract: _MetaderbyHorseNFTToken.contract, event: "ApprovalForAll", logs: logs}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *MetaderbyHorseNFTTokenApprovalForAllOrChainReorg, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _MetaderbyHorseNFTToken.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaderbyHorseNFTTokenApprovalForAllOrChainReorg)
				event.Event = new(MetaderbyHorseNFTTokenApprovalForAll)

				if log.ChainReorg == nil {
					if err := _MetaderbyHorseNFTToken.contract.UnpackLog(event.Event, "ApprovalForAll", *log.Log); err != nil {
						return err
					}
					event.Event.Raw = *log.Log
				} else {
					event.ChainReorg = log.ChainReorg
				}

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenFilterer) ParseApprovalForAll(log types.Log) (*MetaderbyHorseNFTTokenApprovalForAll, error) {
	event := new(MetaderbyHorseNFTTokenApprovalForAll)
	if err := _MetaderbyHorseNFTToken.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaderbyHorseNFTTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MetaderbyHorseNFTToken contract.
type MetaderbyHorseNFTTokenOwnershipTransferredIterator struct {
	Event *MetaderbyHorseNFTTokenOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MetaderbyHorseNFTTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaderbyHorseNFTTokenOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	if it.sub == nil {
		it.done = true
		return it.Next()
	}

	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MetaderbyHorseNFTTokenOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MetaderbyHorseNFTTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaderbyHorseNFTTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaderbyHorseNFTTokenOwnershipTransferred represents a OwnershipTransferred event raised by the MetaderbyHorseNFTToken contract.
type MetaderbyHorseNFTTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// MetaderbyHorseNFTTokenOwnershipTransferredOrChainReorg represents a OwnershipTransferred subscription event raised by the MetaderbyHorseNFTToken contract.
type MetaderbyHorseNFTTokenOwnershipTransferredOrChainReorg struct {
	Event      *MetaderbyHorseNFTTokenOwnershipTransferred
	ChainReorg *types.ChainReorg
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MetaderbyHorseNFTTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, err := _MetaderbyHorseNFTToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MetaderbyHorseNFTTokenOwnershipTransferredIterator{contract: _MetaderbyHorseNFTToken.contract, event: "OwnershipTransferred", logs: logs}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MetaderbyHorseNFTTokenOwnershipTransferredOrChainReorg, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MetaderbyHorseNFTToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaderbyHorseNFTTokenOwnershipTransferredOrChainReorg)
				event.Event = new(MetaderbyHorseNFTTokenOwnershipTransferred)

				if log.ChainReorg == nil {
					if err := _MetaderbyHorseNFTToken.contract.UnpackLog(event.Event, "OwnershipTransferred", *log.Log); err != nil {
						return err
					}
					event.Event.Raw = *log.Log
				} else {
					event.ChainReorg = log.ChainReorg
				}

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenFilterer) ParseOwnershipTransferred(log types.Log) (*MetaderbyHorseNFTTokenOwnershipTransferred, error) {
	event := new(MetaderbyHorseNFTTokenOwnershipTransferred)
	if err := _MetaderbyHorseNFTToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaderbyHorseNFTTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MetaderbyHorseNFTToken contract.
type MetaderbyHorseNFTTokenTransferIterator struct {
	Event *MetaderbyHorseNFTTokenTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MetaderbyHorseNFTTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaderbyHorseNFTTokenTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	if it.sub == nil {
		it.done = true
		return it.Next()
	}

	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MetaderbyHorseNFTTokenTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MetaderbyHorseNFTTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaderbyHorseNFTTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaderbyHorseNFTTokenTransfer represents a Transfer event raised by the MetaderbyHorseNFTToken contract.
type MetaderbyHorseNFTTokenTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// MetaderbyHorseNFTTokenTransferOrChainReorg represents a Transfer subscription event raised by the MetaderbyHorseNFTToken contract.
type MetaderbyHorseNFTTokenTransferOrChainReorg struct {
	Event      *MetaderbyHorseNFTTokenTransfer
	ChainReorg *types.ChainReorg
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*MetaderbyHorseNFTTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, err := _MetaderbyHorseNFTToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MetaderbyHorseNFTTokenTransferIterator{contract: _MetaderbyHorseNFTToken.contract, event: "Transfer", logs: logs}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MetaderbyHorseNFTTokenTransferOrChainReorg, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _MetaderbyHorseNFTToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaderbyHorseNFTTokenTransferOrChainReorg)
				event.Event = new(MetaderbyHorseNFTTokenTransfer)

				if log.ChainReorg == nil {
					if err := _MetaderbyHorseNFTToken.contract.UnpackLog(event.Event, "Transfer", *log.Log); err != nil {
						return err
					}
					event.Event.Raw = *log.Log
				} else {
					event.ChainReorg = log.ChainReorg
				}

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_MetaderbyHorseNFTToken *MetaderbyHorseNFTTokenFilterer) ParseTransfer(log types.Log) (*MetaderbyHorseNFTTokenTransfer, error) {
	event := new(MetaderbyHorseNFTTokenTransfer)
	if err := _MetaderbyHorseNFTToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
