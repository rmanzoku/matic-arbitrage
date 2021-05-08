// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package v1

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// V1ABI is the input ABI used to generate the binding from.
const V1ABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"swapper1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"swapper2\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"forth\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"back\",\"type\":\"address[]\"}],\"name\":\"dry\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"swapper1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"swapper2\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"forth\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"back\",\"type\":\"address[]\"}],\"name\":\"swap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// V1 is an auto generated Go binding around an Ethereum contract.
type V1 struct {
	V1Caller     // Read-only binding to the contract
	V1Transactor // Write-only binding to the contract
	V1Filterer   // Log filterer for contract events
}

// V1Caller is an auto generated read-only Go binding around an Ethereum contract.
type V1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type V1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type V1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type V1Session struct {
	Contract     *V1               // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// V1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type V1CallerSession struct {
	Contract *V1Caller     // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// V1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type V1TransactorSession struct {
	Contract     *V1Transactor     // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// V1Raw is an auto generated low-level Go binding around an Ethereum contract.
type V1Raw struct {
	Contract *V1 // Generic contract binding to access the raw methods on
}

// V1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type V1CallerRaw struct {
	Contract *V1Caller // Generic read-only contract binding to access the raw methods on
}

// V1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type V1TransactorRaw struct {
	Contract *V1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewV1 creates a new instance of V1, bound to a specific deployed contract.
func NewV1(address common.Address, backend bind.ContractBackend) (*V1, error) {
	contract, err := bindV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &V1{V1Caller: V1Caller{contract: contract}, V1Transactor: V1Transactor{contract: contract}, V1Filterer: V1Filterer{contract: contract}}, nil
}

// NewV1Caller creates a new read-only instance of V1, bound to a specific deployed contract.
func NewV1Caller(address common.Address, caller bind.ContractCaller) (*V1Caller, error) {
	contract, err := bindV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &V1Caller{contract: contract}, nil
}

// NewV1Transactor creates a new write-only instance of V1, bound to a specific deployed contract.
func NewV1Transactor(address common.Address, transactor bind.ContractTransactor) (*V1Transactor, error) {
	contract, err := bindV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &V1Transactor{contract: contract}, nil
}

// NewV1Filterer creates a new log filterer instance of V1, bound to a specific deployed contract.
func NewV1Filterer(address common.Address, filterer bind.ContractFilterer) (*V1Filterer, error) {
	contract, err := bindV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &V1Filterer{contract: contract}, nil
}

// bindV1 binds a generic wrapper to an already deployed contract.
func bindV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(V1ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_V1 *V1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _V1.Contract.V1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_V1 *V1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V1.Contract.V1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_V1 *V1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _V1.Contract.V1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_V1 *V1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _V1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_V1 *V1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_V1 *V1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _V1.Contract.contract.Transact(opts, method, params...)
}

// Dry is a free data retrieval call binding the contract method 0x32bd510d.
//
// Solidity: function dry(address swapper1, address swapper2, uint256 val, address[] forth, address[] back) view returns(uint256)
func (_V1 *V1Caller) Dry(opts *bind.CallOpts, swapper1 common.Address, swapper2 common.Address, val *big.Int, forth []common.Address, back []common.Address) (*big.Int, error) {
	var out []interface{}
	err := _V1.contract.Call(opts, &out, "dry", swapper1, swapper2, val, forth, back)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dry is a free data retrieval call binding the contract method 0x32bd510d.
//
// Solidity: function dry(address swapper1, address swapper2, uint256 val, address[] forth, address[] back) view returns(uint256)
func (_V1 *V1Session) Dry(swapper1 common.Address, swapper2 common.Address, val *big.Int, forth []common.Address, back []common.Address) (*big.Int, error) {
	return _V1.Contract.Dry(&_V1.CallOpts, swapper1, swapper2, val, forth, back)
}

// Dry is a free data retrieval call binding the contract method 0x32bd510d.
//
// Solidity: function dry(address swapper1, address swapper2, uint256 val, address[] forth, address[] back) view returns(uint256)
func (_V1 *V1CallerSession) Dry(swapper1 common.Address, swapper2 common.Address, val *big.Int, forth []common.Address, back []common.Address) (*big.Int, error) {
	return _V1.Contract.Dry(&_V1.CallOpts, swapper1, swapper2, val, forth, back)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_V1 *V1Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _V1.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_V1 *V1Session) Owner() (common.Address, error) {
	return _V1.Contract.Owner(&_V1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_V1 *V1CallerSession) Owner() (common.Address, error) {
	return _V1.Contract.Owner(&_V1.CallOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_V1 *V1Transactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V1.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_V1 *V1Session) Kill() (*types.Transaction, error) {
	return _V1.Contract.Kill(&_V1.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_V1 *V1TransactorSession) Kill() (*types.Transaction, error) {
	return _V1.Contract.Kill(&_V1.TransactOpts)
}

// Swap is a paid mutator transaction binding the contract method 0xd530dd7a.
//
// Solidity: function swap(address swapper1, address swapper2, uint256 val, address[] forth, address[] back) returns()
func (_V1 *V1Transactor) Swap(opts *bind.TransactOpts, swapper1 common.Address, swapper2 common.Address, val *big.Int, forth []common.Address, back []common.Address) (*types.Transaction, error) {
	return _V1.contract.Transact(opts, "swap", swapper1, swapper2, val, forth, back)
}

// Swap is a paid mutator transaction binding the contract method 0xd530dd7a.
//
// Solidity: function swap(address swapper1, address swapper2, uint256 val, address[] forth, address[] back) returns()
func (_V1 *V1Session) Swap(swapper1 common.Address, swapper2 common.Address, val *big.Int, forth []common.Address, back []common.Address) (*types.Transaction, error) {
	return _V1.Contract.Swap(&_V1.TransactOpts, swapper1, swapper2, val, forth, back)
}

// Swap is a paid mutator transaction binding the contract method 0xd530dd7a.
//
// Solidity: function swap(address swapper1, address swapper2, uint256 val, address[] forth, address[] back) returns()
func (_V1 *V1TransactorSession) Swap(swapper1 common.Address, swapper2 common.Address, val *big.Int, forth []common.Address, back []common.Address) (*types.Transaction, error) {
	return _V1.Contract.Swap(&_V1.TransactOpts, swapper1, swapper2, val, forth, back)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address _token) returns()
func (_V1 *V1Transactor) Withdraw(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _V1.contract.Transact(opts, "withdraw", _token)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address _token) returns()
func (_V1 *V1Session) Withdraw(_token common.Address) (*types.Transaction, error) {
	return _V1.Contract.Withdraw(&_V1.TransactOpts, _token)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address _token) returns()
func (_V1 *V1TransactorSession) Withdraw(_token common.Address) (*types.Transaction, error) {
	return _V1.Contract.Withdraw(&_V1.TransactOpts, _token)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_V1 *V1Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V1.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_V1 *V1Session) Receive() (*types.Transaction, error) {
	return _V1.Contract.Receive(&_V1.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_V1 *V1TransactorSession) Receive() (*types.Transaction, error) {
	return _V1.Contract.Receive(&_V1.TransactOpts)
}

