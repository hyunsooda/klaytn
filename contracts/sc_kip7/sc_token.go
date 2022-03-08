// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sctoken

import (
	"math/big"
	"strings"

	"github.com/klaytn/klaytn"
	"github.com/klaytn/klaytn/accounts/abi"
	"github.com/klaytn/klaytn/accounts/abi/bind"
	"github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = klaytn.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// AddressABI is the input ABI used to generate the binding from.
const AddressABI = "[]"

// AddressBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const AddressBinRuntime = `73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a72305820b7534b8b203852d9ead56352219b5abb5424533d9bb1b9cf84c0b1d15d11fee50029`

// AddressBin is the compiled bytecode used for deploying new contracts.
var AddressBin = "0x604c6023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a72305820b7534b8b203852d9ead56352219b5abb5424533d9bb1b9cf84c0b1d15d11fee50029"

// DeployAddress deploys a new Klaytn contract, binding an instance of Address to it.
func DeployAddress(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Address, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AddressBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// Address is an auto generated Go binding around a Klaytn contract.
type Address struct {
	AddressCaller     // Read-only binding to the contract
	AddressTransactor // Write-only binding to the contract
	AddressFilterer   // Log filterer for contract events
}

// AddressCaller is an auto generated read-only Go binding around a Klaytn contract.
type AddressCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressTransactor is an auto generated write-only Go binding around a Klaytn contract.
type AddressTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type AddressFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type AddressSession struct {
	Contract     *Address          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type AddressCallerSession struct {
	Contract *AddressCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AddressTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type AddressTransactorSession struct {
	Contract     *AddressTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AddressRaw is an auto generated low-level Go binding around a Klaytn contract.
type AddressRaw struct {
	Contract *Address // Generic contract binding to access the raw methods on
}

// AddressCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type AddressCallerRaw struct {
	Contract *AddressCaller // Generic read-only contract binding to access the raw methods on
}

// AddressTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type AddressTransactorRaw struct {
	Contract *AddressTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddress creates a new instance of Address, bound to a specific deployed contract.
func NewAddress(address common.Address, backend bind.ContractBackend) (*Address, error) {
	contract, err := bindAddress(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// NewAddressCaller creates a new read-only instance of Address, bound to a specific deployed contract.
func NewAddressCaller(address common.Address, caller bind.ContractCaller) (*AddressCaller, error) {
	contract, err := bindAddress(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressCaller{contract: contract}, nil
}

// NewAddressTransactor creates a new write-only instance of Address, bound to a specific deployed contract.
func NewAddressTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressTransactor, error) {
	contract, err := bindAddress(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressTransactor{contract: contract}, nil
}

// NewAddressFilterer creates a new log filterer instance of Address, bound to a specific deployed contract.
func NewAddressFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressFilterer, error) {
	contract, err := bindAddress(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressFilterer{contract: contract}, nil
}

// bindAddress binds a generic wrapper to an already deployed contract.
func bindAddress(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Address.Contract.AddressCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Address.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.contract.Transact(opts, method, params...)
}

// IKIP13ABI is the input ABI used to generate the binding from.
const IKIP13ABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IKIP13BinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const IKIP13BinRuntime = ``

// IKIP13FuncSigs maps the 4-byte function signature to its string representation.
var IKIP13FuncSigs = map[string]string{
	"01ffc9a7": "supportsInterface(bytes4)",
}

// IKIP13 is an auto generated Go binding around a Klaytn contract.
type IKIP13 struct {
	IKIP13Caller     // Read-only binding to the contract
	IKIP13Transactor // Write-only binding to the contract
	IKIP13Filterer   // Log filterer for contract events
}

// IKIP13Caller is an auto generated read-only Go binding around a Klaytn contract.
type IKIP13Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IKIP13Transactor is an auto generated write-only Go binding around a Klaytn contract.
type IKIP13Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IKIP13Filterer is an auto generated log filtering Go binding around a Klaytn contract events.
type IKIP13Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IKIP13Session is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type IKIP13Session struct {
	Contract     *IKIP13           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IKIP13CallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type IKIP13CallerSession struct {
	Contract *IKIP13Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IKIP13TransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type IKIP13TransactorSession struct {
	Contract     *IKIP13Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IKIP13Raw is an auto generated low-level Go binding around a Klaytn contract.
type IKIP13Raw struct {
	Contract *IKIP13 // Generic contract binding to access the raw methods on
}

// IKIP13CallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type IKIP13CallerRaw struct {
	Contract *IKIP13Caller // Generic read-only contract binding to access the raw methods on
}

// IKIP13TransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type IKIP13TransactorRaw struct {
	Contract *IKIP13Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIKIP13 creates a new instance of IKIP13, bound to a specific deployed contract.
func NewIKIP13(address common.Address, backend bind.ContractBackend) (*IKIP13, error) {
	contract, err := bindIKIP13(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IKIP13{IKIP13Caller: IKIP13Caller{contract: contract}, IKIP13Transactor: IKIP13Transactor{contract: contract}, IKIP13Filterer: IKIP13Filterer{contract: contract}}, nil
}

// NewIKIP13Caller creates a new read-only instance of IKIP13, bound to a specific deployed contract.
func NewIKIP13Caller(address common.Address, caller bind.ContractCaller) (*IKIP13Caller, error) {
	contract, err := bindIKIP13(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IKIP13Caller{contract: contract}, nil
}

// NewIKIP13Transactor creates a new write-only instance of IKIP13, bound to a specific deployed contract.
func NewIKIP13Transactor(address common.Address, transactor bind.ContractTransactor) (*IKIP13Transactor, error) {
	contract, err := bindIKIP13(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IKIP13Transactor{contract: contract}, nil
}

// NewIKIP13Filterer creates a new log filterer instance of IKIP13, bound to a specific deployed contract.
func NewIKIP13Filterer(address common.Address, filterer bind.ContractFilterer) (*IKIP13Filterer, error) {
	contract, err := bindIKIP13(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IKIP13Filterer{contract: contract}, nil
}

// bindIKIP13 binds a generic wrapper to an already deployed contract.
func bindIKIP13(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IKIP13ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IKIP13 *IKIP13Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IKIP13.Contract.IKIP13Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IKIP13 *IKIP13Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IKIP13.Contract.IKIP13Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IKIP13 *IKIP13Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IKIP13.Contract.IKIP13Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IKIP13 *IKIP13CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IKIP13.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IKIP13 *IKIP13TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IKIP13.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IKIP13 *IKIP13TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IKIP13.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IKIP13 *IKIP13Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IKIP13.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IKIP13 *IKIP13Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IKIP13.Contract.SupportsInterface(&_IKIP13.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IKIP13 *IKIP13CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IKIP13.Contract.SupportsInterface(&_IKIP13.CallOpts, interfaceId)
}

// IKIP7ABI is the input ABI used to generate the binding from.
const IKIP7ABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"safeTransfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// IKIP7BinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const IKIP7BinRuntime = ``

// IKIP7FuncSigs maps the 4-byte function signature to its string representation.
var IKIP7FuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"423f6cef": "safeTransfer(address,uint256)",
	"eb795549": "safeTransfer(address,uint256,bytes)",
	"42842e0e": "safeTransferFrom(address,address,uint256)",
	"b88d4fde": "safeTransferFrom(address,address,uint256,bytes)",
	"01ffc9a7": "supportsInterface(bytes4)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// IKIP7 is an auto generated Go binding around a Klaytn contract.
type IKIP7 struct {
	IKIP7Caller     // Read-only binding to the contract
	IKIP7Transactor // Write-only binding to the contract
	IKIP7Filterer   // Log filterer for contract events
}

// IKIP7Caller is an auto generated read-only Go binding around a Klaytn contract.
type IKIP7Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IKIP7Transactor is an auto generated write-only Go binding around a Klaytn contract.
type IKIP7Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IKIP7Filterer is an auto generated log filtering Go binding around a Klaytn contract events.
type IKIP7Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IKIP7Session is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type IKIP7Session struct {
	Contract     *IKIP7            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IKIP7CallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type IKIP7CallerSession struct {
	Contract *IKIP7Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IKIP7TransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type IKIP7TransactorSession struct {
	Contract     *IKIP7Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IKIP7Raw is an auto generated low-level Go binding around a Klaytn contract.
type IKIP7Raw struct {
	Contract *IKIP7 // Generic contract binding to access the raw methods on
}

// IKIP7CallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type IKIP7CallerRaw struct {
	Contract *IKIP7Caller // Generic read-only contract binding to access the raw methods on
}

// IKIP7TransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type IKIP7TransactorRaw struct {
	Contract *IKIP7Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIKIP7 creates a new instance of IKIP7, bound to a specific deployed contract.
func NewIKIP7(address common.Address, backend bind.ContractBackend) (*IKIP7, error) {
	contract, err := bindIKIP7(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IKIP7{IKIP7Caller: IKIP7Caller{contract: contract}, IKIP7Transactor: IKIP7Transactor{contract: contract}, IKIP7Filterer: IKIP7Filterer{contract: contract}}, nil
}

// NewIKIP7Caller creates a new read-only instance of IKIP7, bound to a specific deployed contract.
func NewIKIP7Caller(address common.Address, caller bind.ContractCaller) (*IKIP7Caller, error) {
	contract, err := bindIKIP7(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IKIP7Caller{contract: contract}, nil
}

// NewIKIP7Transactor creates a new write-only instance of IKIP7, bound to a specific deployed contract.
func NewIKIP7Transactor(address common.Address, transactor bind.ContractTransactor) (*IKIP7Transactor, error) {
	contract, err := bindIKIP7(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IKIP7Transactor{contract: contract}, nil
}

// NewIKIP7Filterer creates a new log filterer instance of IKIP7, bound to a specific deployed contract.
func NewIKIP7Filterer(address common.Address, filterer bind.ContractFilterer) (*IKIP7Filterer, error) {
	contract, err := bindIKIP7(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IKIP7Filterer{contract: contract}, nil
}

// bindIKIP7 binds a generic wrapper to an already deployed contract.
func bindIKIP7(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IKIP7ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IKIP7 *IKIP7Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IKIP7.Contract.IKIP7Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IKIP7 *IKIP7Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IKIP7.Contract.IKIP7Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IKIP7 *IKIP7Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IKIP7.Contract.IKIP7Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IKIP7 *IKIP7CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IKIP7.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IKIP7 *IKIP7TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IKIP7.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IKIP7 *IKIP7TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IKIP7.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IKIP7 *IKIP7Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IKIP7.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IKIP7 *IKIP7Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IKIP7.Contract.Allowance(&_IKIP7.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IKIP7 *IKIP7CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IKIP7.Contract.Allowance(&_IKIP7.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IKIP7 *IKIP7Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IKIP7.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IKIP7 *IKIP7Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IKIP7.Contract.BalanceOf(&_IKIP7.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IKIP7 *IKIP7CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IKIP7.Contract.BalanceOf(&_IKIP7.CallOpts, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IKIP7 *IKIP7Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IKIP7.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IKIP7 *IKIP7Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IKIP7.Contract.SupportsInterface(&_IKIP7.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IKIP7 *IKIP7CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IKIP7.Contract.SupportsInterface(&_IKIP7.CallOpts, interfaceId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IKIP7 *IKIP7Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IKIP7.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IKIP7 *IKIP7Session) TotalSupply() (*big.Int, error) {
	return _IKIP7.Contract.TotalSupply(&_IKIP7.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IKIP7 *IKIP7CallerSession) TotalSupply() (*big.Int, error) {
	return _IKIP7.Contract.TotalSupply(&_IKIP7.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IKIP7 *IKIP7Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IKIP7.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IKIP7 *IKIP7Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IKIP7.Contract.Approve(&_IKIP7.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IKIP7 *IKIP7TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IKIP7.Contract.Approve(&_IKIP7.TransactOpts, spender, amount)
}

// SafeTransfer is a paid mutator transaction binding the contract method 0x423f6cef.
//
// Solidity: function safeTransfer(address recipient, uint256 amount) returns()
func (_IKIP7 *IKIP7Transactor) SafeTransfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IKIP7.contract.Transact(opts, "safeTransfer", recipient, amount)
}

// SafeTransfer is a paid mutator transaction binding the contract method 0x423f6cef.
//
// Solidity: function safeTransfer(address recipient, uint256 amount) returns()
func (_IKIP7 *IKIP7Session) SafeTransfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IKIP7.Contract.SafeTransfer(&_IKIP7.TransactOpts, recipient, amount)
}

// SafeTransfer is a paid mutator transaction binding the contract method 0x423f6cef.
//
// Solidity: function safeTransfer(address recipient, uint256 amount) returns()
func (_IKIP7 *IKIP7TransactorSession) SafeTransfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IKIP7.Contract.SafeTransfer(&_IKIP7.TransactOpts, recipient, amount)
}

// SafeTransfer0 is a paid mutator transaction binding the contract method 0xeb795549.
//
// Solidity: function safeTransfer(address recipient, uint256 amount, bytes data) returns()
func (_IKIP7 *IKIP7Transactor) SafeTransfer0(opts *bind.TransactOpts, recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IKIP7.contract.Transact(opts, "safeTransfer0", recipient, amount, data)
}

// SafeTransfer0 is a paid mutator transaction binding the contract method 0xeb795549.
//
// Solidity: function safeTransfer(address recipient, uint256 amount, bytes data) returns()
func (_IKIP7 *IKIP7Session) SafeTransfer0(recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IKIP7.Contract.SafeTransfer0(&_IKIP7.TransactOpts, recipient, amount, data)
}

// SafeTransfer0 is a paid mutator transaction binding the contract method 0xeb795549.
//
// Solidity: function safeTransfer(address recipient, uint256 amount, bytes data) returns()
func (_IKIP7 *IKIP7TransactorSession) SafeTransfer0(recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IKIP7.Contract.SafeTransfer0(&_IKIP7.TransactOpts, recipient, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address sender, address recipient, uint256 amount) returns()
func (_IKIP7 *IKIP7Transactor) SafeTransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IKIP7.contract.Transact(opts, "safeTransferFrom", sender, recipient, amount)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address sender, address recipient, uint256 amount) returns()
func (_IKIP7 *IKIP7Session) SafeTransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IKIP7.Contract.SafeTransferFrom(&_IKIP7.TransactOpts, sender, recipient, amount)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address sender, address recipient, uint256 amount) returns()
func (_IKIP7 *IKIP7TransactorSession) SafeTransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IKIP7.Contract.SafeTransferFrom(&_IKIP7.TransactOpts, sender, recipient, amount)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address sender, address recipient, uint256 amount, bytes data) returns()
func (_IKIP7 *IKIP7Transactor) SafeTransferFrom0(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IKIP7.contract.Transact(opts, "safeTransferFrom0", sender, recipient, amount, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address sender, address recipient, uint256 amount, bytes data) returns()
func (_IKIP7 *IKIP7Session) SafeTransferFrom0(sender common.Address, recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IKIP7.Contract.SafeTransferFrom0(&_IKIP7.TransactOpts, sender, recipient, amount, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address sender, address recipient, uint256 amount, bytes data) returns()
func (_IKIP7 *IKIP7TransactorSession) SafeTransferFrom0(sender common.Address, recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IKIP7.Contract.SafeTransferFrom0(&_IKIP7.TransactOpts, sender, recipient, amount, data)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IKIP7 *IKIP7Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IKIP7.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IKIP7 *IKIP7Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IKIP7.Contract.Transfer(&_IKIP7.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IKIP7 *IKIP7TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IKIP7.Contract.Transfer(&_IKIP7.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IKIP7 *IKIP7Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IKIP7.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IKIP7 *IKIP7Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IKIP7.Contract.TransferFrom(&_IKIP7.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IKIP7 *IKIP7TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IKIP7.Contract.TransferFrom(&_IKIP7.TransactOpts, sender, recipient, amount)
}

// IKIP7ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IKIP7 contract.
type IKIP7ApprovalIterator struct {
	Event *IKIP7Approval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IKIP7ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IKIP7Approval)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IKIP7Approval)
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
func (it *IKIP7ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IKIP7ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IKIP7Approval represents a Approval event raised by the IKIP7 contract.
type IKIP7Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IKIP7 *IKIP7Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IKIP7ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IKIP7.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IKIP7ApprovalIterator{contract: _IKIP7.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IKIP7 *IKIP7Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IKIP7Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IKIP7.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IKIP7Approval)
				if err := _IKIP7.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IKIP7 *IKIP7Filterer) ParseApproval(log types.Log) (*IKIP7Approval, error) {
	event := new(IKIP7Approval)
	if err := _IKIP7.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IKIP7TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IKIP7 contract.
type IKIP7TransferIterator struct {
	Event *IKIP7Transfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IKIP7TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IKIP7Transfer)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IKIP7Transfer)
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
func (it *IKIP7TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IKIP7TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IKIP7Transfer represents a Transfer event raised by the IKIP7 contract.
type IKIP7Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IKIP7 *IKIP7Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IKIP7TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IKIP7.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IKIP7TransferIterator{contract: _IKIP7.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IKIP7 *IKIP7Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IKIP7Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IKIP7.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IKIP7Transfer)
				if err := _IKIP7.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IKIP7 *IKIP7Filterer) ParseTransfer(log types.Log) (*IKIP7Transfer, error) {
	event := new(IKIP7Transfer)
	if err := _IKIP7.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IKIP7BridgeReceiverABI is the input ABI used to generate the binding from.
const IKIP7BridgeReceiverABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_feeLimit\",\"type\":\"uint256\"},{\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"onKIP7Received\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IKIP7BridgeReceiverBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const IKIP7BridgeReceiverBinRuntime = ``

// IKIP7BridgeReceiverFuncSigs maps the 4-byte function signature to its string representation.
var IKIP7BridgeReceiverFuncSigs = map[string]string{
	"c80e3a50": "onKIP7Received(address,address,uint256,uint256,bytes)",
}

// IKIP7BridgeReceiver is an auto generated Go binding around a Klaytn contract.
type IKIP7BridgeReceiver struct {
	IKIP7BridgeReceiverCaller     // Read-only binding to the contract
	IKIP7BridgeReceiverTransactor // Write-only binding to the contract
	IKIP7BridgeReceiverFilterer   // Log filterer for contract events
}

// IKIP7BridgeReceiverCaller is an auto generated read-only Go binding around a Klaytn contract.
type IKIP7BridgeReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IKIP7BridgeReceiverTransactor is an auto generated write-only Go binding around a Klaytn contract.
type IKIP7BridgeReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IKIP7BridgeReceiverFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type IKIP7BridgeReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IKIP7BridgeReceiverSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type IKIP7BridgeReceiverSession struct {
	Contract     *IKIP7BridgeReceiver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IKIP7BridgeReceiverCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type IKIP7BridgeReceiverCallerSession struct {
	Contract *IKIP7BridgeReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IKIP7BridgeReceiverTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type IKIP7BridgeReceiverTransactorSession struct {
	Contract     *IKIP7BridgeReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IKIP7BridgeReceiverRaw is an auto generated low-level Go binding around a Klaytn contract.
type IKIP7BridgeReceiverRaw struct {
	Contract *IKIP7BridgeReceiver // Generic contract binding to access the raw methods on
}

// IKIP7BridgeReceiverCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type IKIP7BridgeReceiverCallerRaw struct {
	Contract *IKIP7BridgeReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// IKIP7BridgeReceiverTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type IKIP7BridgeReceiverTransactorRaw struct {
	Contract *IKIP7BridgeReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIKIP7BridgeReceiver creates a new instance of IKIP7BridgeReceiver, bound to a specific deployed contract.
func NewIKIP7BridgeReceiver(address common.Address, backend bind.ContractBackend) (*IKIP7BridgeReceiver, error) {
	contract, err := bindIKIP7BridgeReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IKIP7BridgeReceiver{IKIP7BridgeReceiverCaller: IKIP7BridgeReceiverCaller{contract: contract}, IKIP7BridgeReceiverTransactor: IKIP7BridgeReceiverTransactor{contract: contract}, IKIP7BridgeReceiverFilterer: IKIP7BridgeReceiverFilterer{contract: contract}}, nil
}

// NewIKIP7BridgeReceiverCaller creates a new read-only instance of IKIP7BridgeReceiver, bound to a specific deployed contract.
func NewIKIP7BridgeReceiverCaller(address common.Address, caller bind.ContractCaller) (*IKIP7BridgeReceiverCaller, error) {
	contract, err := bindIKIP7BridgeReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IKIP7BridgeReceiverCaller{contract: contract}, nil
}

// NewIKIP7BridgeReceiverTransactor creates a new write-only instance of IKIP7BridgeReceiver, bound to a specific deployed contract.
func NewIKIP7BridgeReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*IKIP7BridgeReceiverTransactor, error) {
	contract, err := bindIKIP7BridgeReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IKIP7BridgeReceiverTransactor{contract: contract}, nil
}

// NewIKIP7BridgeReceiverFilterer creates a new log filterer instance of IKIP7BridgeReceiver, bound to a specific deployed contract.
func NewIKIP7BridgeReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*IKIP7BridgeReceiverFilterer, error) {
	contract, err := bindIKIP7BridgeReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IKIP7BridgeReceiverFilterer{contract: contract}, nil
}

// bindIKIP7BridgeReceiver binds a generic wrapper to an already deployed contract.
func bindIKIP7BridgeReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IKIP7BridgeReceiverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IKIP7BridgeReceiver *IKIP7BridgeReceiverRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IKIP7BridgeReceiver.Contract.IKIP7BridgeReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IKIP7BridgeReceiver *IKIP7BridgeReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IKIP7BridgeReceiver.Contract.IKIP7BridgeReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IKIP7BridgeReceiver *IKIP7BridgeReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IKIP7BridgeReceiver.Contract.IKIP7BridgeReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IKIP7BridgeReceiver *IKIP7BridgeReceiverCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IKIP7BridgeReceiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IKIP7BridgeReceiver *IKIP7BridgeReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IKIP7BridgeReceiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IKIP7BridgeReceiver *IKIP7BridgeReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IKIP7BridgeReceiver.Contract.contract.Transact(opts, method, params...)
}

// OnKIP7Received is a paid mutator transaction binding the contract method 0xc80e3a50.
//
// Solidity: function onKIP7Received(address _from, address _to, uint256 _amount, uint256 _feeLimit, bytes _extraData) returns()
func (_IKIP7BridgeReceiver *IKIP7BridgeReceiverTransactor) OnKIP7Received(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int, _feeLimit *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _IKIP7BridgeReceiver.contract.Transact(opts, "onKIP7Received", _from, _to, _amount, _feeLimit, _extraData)
}

// OnKIP7Received is a paid mutator transaction binding the contract method 0xc80e3a50.
//
// Solidity: function onKIP7Received(address _from, address _to, uint256 _amount, uint256 _feeLimit, bytes _extraData) returns()
func (_IKIP7BridgeReceiver *IKIP7BridgeReceiverSession) OnKIP7Received(_from common.Address, _to common.Address, _amount *big.Int, _feeLimit *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _IKIP7BridgeReceiver.Contract.OnKIP7Received(&_IKIP7BridgeReceiver.TransactOpts, _from, _to, _amount, _feeLimit, _extraData)
}

// OnKIP7Received is a paid mutator transaction binding the contract method 0xc80e3a50.
//
// Solidity: function onKIP7Received(address _from, address _to, uint256 _amount, uint256 _feeLimit, bytes _extraData) returns()
func (_IKIP7BridgeReceiver *IKIP7BridgeReceiverTransactorSession) OnKIP7Received(_from common.Address, _to common.Address, _amount *big.Int, _feeLimit *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _IKIP7BridgeReceiver.Contract.OnKIP7Received(&_IKIP7BridgeReceiver.TransactOpts, _from, _to, _amount, _feeLimit, _extraData)
}

// IKIP7ReceiverABI is the input ABI used to generate the binding from.
const IKIP7ReceiverABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"onKIP7Received\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IKIP7ReceiverBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const IKIP7ReceiverBinRuntime = ``

// IKIP7ReceiverFuncSigs maps the 4-byte function signature to its string representation.
var IKIP7ReceiverFuncSigs = map[string]string{
	"9d188c22": "onKIP7Received(address,address,uint256,bytes)",
}

// IKIP7Receiver is an auto generated Go binding around a Klaytn contract.
type IKIP7Receiver struct {
	IKIP7ReceiverCaller     // Read-only binding to the contract
	IKIP7ReceiverTransactor // Write-only binding to the contract
	IKIP7ReceiverFilterer   // Log filterer for contract events
}

// IKIP7ReceiverCaller is an auto generated read-only Go binding around a Klaytn contract.
type IKIP7ReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IKIP7ReceiverTransactor is an auto generated write-only Go binding around a Klaytn contract.
type IKIP7ReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IKIP7ReceiverFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type IKIP7ReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IKIP7ReceiverSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type IKIP7ReceiverSession struct {
	Contract     *IKIP7Receiver    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IKIP7ReceiverCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type IKIP7ReceiverCallerSession struct {
	Contract *IKIP7ReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IKIP7ReceiverTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type IKIP7ReceiverTransactorSession struct {
	Contract     *IKIP7ReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IKIP7ReceiverRaw is an auto generated low-level Go binding around a Klaytn contract.
type IKIP7ReceiverRaw struct {
	Contract *IKIP7Receiver // Generic contract binding to access the raw methods on
}

// IKIP7ReceiverCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type IKIP7ReceiverCallerRaw struct {
	Contract *IKIP7ReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// IKIP7ReceiverTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type IKIP7ReceiverTransactorRaw struct {
	Contract *IKIP7ReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIKIP7Receiver creates a new instance of IKIP7Receiver, bound to a specific deployed contract.
func NewIKIP7Receiver(address common.Address, backend bind.ContractBackend) (*IKIP7Receiver, error) {
	contract, err := bindIKIP7Receiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IKIP7Receiver{IKIP7ReceiverCaller: IKIP7ReceiverCaller{contract: contract}, IKIP7ReceiverTransactor: IKIP7ReceiverTransactor{contract: contract}, IKIP7ReceiverFilterer: IKIP7ReceiverFilterer{contract: contract}}, nil
}

// NewIKIP7ReceiverCaller creates a new read-only instance of IKIP7Receiver, bound to a specific deployed contract.
func NewIKIP7ReceiverCaller(address common.Address, caller bind.ContractCaller) (*IKIP7ReceiverCaller, error) {
	contract, err := bindIKIP7Receiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IKIP7ReceiverCaller{contract: contract}, nil
}

// NewIKIP7ReceiverTransactor creates a new write-only instance of IKIP7Receiver, bound to a specific deployed contract.
func NewIKIP7ReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*IKIP7ReceiverTransactor, error) {
	contract, err := bindIKIP7Receiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IKIP7ReceiverTransactor{contract: contract}, nil
}

// NewIKIP7ReceiverFilterer creates a new log filterer instance of IKIP7Receiver, bound to a specific deployed contract.
func NewIKIP7ReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*IKIP7ReceiverFilterer, error) {
	contract, err := bindIKIP7Receiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IKIP7ReceiverFilterer{contract: contract}, nil
}

// bindIKIP7Receiver binds a generic wrapper to an already deployed contract.
func bindIKIP7Receiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IKIP7ReceiverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IKIP7Receiver *IKIP7ReceiverRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IKIP7Receiver.Contract.IKIP7ReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IKIP7Receiver *IKIP7ReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IKIP7Receiver.Contract.IKIP7ReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IKIP7Receiver *IKIP7ReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IKIP7Receiver.Contract.IKIP7ReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IKIP7Receiver *IKIP7ReceiverCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IKIP7Receiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IKIP7Receiver *IKIP7ReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IKIP7Receiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IKIP7Receiver *IKIP7ReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IKIP7Receiver.Contract.contract.Transact(opts, method, params...)
}

// OnKIP7Received is a paid mutator transaction binding the contract method 0x9d188c22.
//
// Solidity: function onKIP7Received(address _operator, address _from, uint256 _amount, bytes _data) returns(bytes4)
func (_IKIP7Receiver *IKIP7ReceiverTransactor) OnKIP7Received(opts *bind.TransactOpts, _operator common.Address, _from common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _IKIP7Receiver.contract.Transact(opts, "onKIP7Received", _operator, _from, _amount, _data)
}

// OnKIP7Received is a paid mutator transaction binding the contract method 0x9d188c22.
//
// Solidity: function onKIP7Received(address _operator, address _from, uint256 _amount, bytes _data) returns(bytes4)
func (_IKIP7Receiver *IKIP7ReceiverSession) OnKIP7Received(_operator common.Address, _from common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _IKIP7Receiver.Contract.OnKIP7Received(&_IKIP7Receiver.TransactOpts, _operator, _from, _amount, _data)
}

// OnKIP7Received is a paid mutator transaction binding the contract method 0x9d188c22.
//
// Solidity: function onKIP7Received(address _operator, address _from, uint256 _amount, bytes _data) returns(bytes4)
func (_IKIP7Receiver *IKIP7ReceiverTransactorSession) OnKIP7Received(_operator common.Address, _from common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _IKIP7Receiver.Contract.OnKIP7Received(&_IKIP7Receiver.TransactOpts, _operator, _from, _amount, _data)
}

// KIP13ABI is the input ABI used to generate the binding from.
const KIP13ABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// KIP13BinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const KIP13BinRuntime = ``

// KIP13FuncSigs maps the 4-byte function signature to its string representation.
var KIP13FuncSigs = map[string]string{
	"01ffc9a7": "supportsInterface(bytes4)",
}

// KIP13 is an auto generated Go binding around a Klaytn contract.
type KIP13 struct {
	KIP13Caller     // Read-only binding to the contract
	KIP13Transactor // Write-only binding to the contract
	KIP13Filterer   // Log filterer for contract events
}

// KIP13Caller is an auto generated read-only Go binding around a Klaytn contract.
type KIP13Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KIP13Transactor is an auto generated write-only Go binding around a Klaytn contract.
type KIP13Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KIP13Filterer is an auto generated log filtering Go binding around a Klaytn contract events.
type KIP13Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KIP13Session is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type KIP13Session struct {
	Contract     *KIP13            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KIP13CallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type KIP13CallerSession struct {
	Contract *KIP13Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// KIP13TransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type KIP13TransactorSession struct {
	Contract     *KIP13Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KIP13Raw is an auto generated low-level Go binding around a Klaytn contract.
type KIP13Raw struct {
	Contract *KIP13 // Generic contract binding to access the raw methods on
}

// KIP13CallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type KIP13CallerRaw struct {
	Contract *KIP13Caller // Generic read-only contract binding to access the raw methods on
}

// KIP13TransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type KIP13TransactorRaw struct {
	Contract *KIP13Transactor // Generic write-only contract binding to access the raw methods on
}

// NewKIP13 creates a new instance of KIP13, bound to a specific deployed contract.
func NewKIP13(address common.Address, backend bind.ContractBackend) (*KIP13, error) {
	contract, err := bindKIP13(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KIP13{KIP13Caller: KIP13Caller{contract: contract}, KIP13Transactor: KIP13Transactor{contract: contract}, KIP13Filterer: KIP13Filterer{contract: contract}}, nil
}

// NewKIP13Caller creates a new read-only instance of KIP13, bound to a specific deployed contract.
func NewKIP13Caller(address common.Address, caller bind.ContractCaller) (*KIP13Caller, error) {
	contract, err := bindKIP13(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KIP13Caller{contract: contract}, nil
}

// NewKIP13Transactor creates a new write-only instance of KIP13, bound to a specific deployed contract.
func NewKIP13Transactor(address common.Address, transactor bind.ContractTransactor) (*KIP13Transactor, error) {
	contract, err := bindKIP13(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KIP13Transactor{contract: contract}, nil
}

// NewKIP13Filterer creates a new log filterer instance of KIP13, bound to a specific deployed contract.
func NewKIP13Filterer(address common.Address, filterer bind.ContractFilterer) (*KIP13Filterer, error) {
	contract, err := bindKIP13(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KIP13Filterer{contract: contract}, nil
}

// bindKIP13 binds a generic wrapper to an already deployed contract.
func bindKIP13(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KIP13ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KIP13 *KIP13Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KIP13.Contract.KIP13Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KIP13 *KIP13Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KIP13.Contract.KIP13Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KIP13 *KIP13Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KIP13.Contract.KIP13Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KIP13 *KIP13CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KIP13.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KIP13 *KIP13TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KIP13.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KIP13 *KIP13TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KIP13.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_KIP13 *KIP13Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KIP13.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_KIP13 *KIP13Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _KIP13.Contract.SupportsInterface(&_KIP13.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_KIP13 *KIP13CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _KIP13.Contract.SupportsInterface(&_KIP13.CallOpts, interfaceId)
}

// KIP7ABI is the input ABI used to generate the binding from.
const KIP7ABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"safeTransfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// KIP7BinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const KIP7BinRuntime = `608060405234801561001057600080fd5b50600436106100a95760003560e01c806342842e0e1161007157806342842e0e1461019357806370a08231146101c9578063a9059cbb146101ef578063b88d4fde1461021b578063dd62ed3e146102e1578063eb7955491461030f576100a9565b806301ffc9a7146100ae578063095ea7b3146100e957806318160ddd1461011557806323b872dd1461012f578063423f6cef14610165575b600080fd5b6100d5600480360360208110156100c457600080fd5b50356001600160e01b0319166103ca565b604080519115158252519081900360200190f35b6100d5600480360360408110156100ff57600080fd5b506001600160a01b0381351690602001356103e9565b61011d6103ff565b60408051918252519081900360200190f35b6100d56004803603606081101561014557600080fd5b506001600160a01b03813581169160208101359091169060400135610405565b6101916004803603604081101561017b57600080fd5b506001600160a01b03813516906020013561045c565b005b610191600480360360608110156101a957600080fd5b506001600160a01b0381358116916020810135909116906040013561047a565b61011d600480360360208110156101df57600080fd5b50356001600160a01b031661049a565b6100d56004803603604081101561020557600080fd5b506001600160a01b0381351690602001356104b5565b6101916004803603608081101561023157600080fd5b6001600160a01b0382358116926020810135909116916040820135919081019060808101606082013564010000000081111561026c57600080fd5b82018360208201111561027e57600080fd5b803590602001918460018302840111640100000000831117156102a057600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506104c2945050505050565b61011d600480360360408110156102f757600080fd5b506001600160a01b038135811691602001351661051e565b6101916004803603606081101561032557600080fd5b6001600160a01b038235169160208101359181019060608101604082013564010000000081111561035557600080fd5b82018360208201111561036757600080fd5b8035906020019184600183028401116401000000008311171561038957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610549945050505050565b6001600160e01b03191660009081526020819052604090205460ff1690565b60006103f633848461059e565b50600192915050565b60035490565b6000610412848484610690565b6001600160a01b03841660009081526002602090815260408083203380855292529091205461045291869161044d908663ffffffff6107da16565b61059e565b5060019392505050565b610476828260405180602001604052806000815250610549565b5050565b610495838383604051806020016040528060008152506104c2565b505050565b6001600160a01b031660009081526001602052604090205490565b60006103f6338484610690565b6104cd848484610405565b506104da8484848461083a565b61051857604051600160e51b62461bcd02815260040180806020018281038252602e815260200180610a00602e913960400191505060405180910390fd5b50505050565b6001600160a01b03918216600090815260026020908152604080832093909416825291909152205490565b61055383836104b5565b506105603384848461083a565b61049557604051600160e51b62461bcd02815260040180806020018281038252602e815260200180610a00602e913960400191505060405180910390fd5b6001600160a01b0383166105e657604051600160e51b62461bcd028152600401808060200182810382526023815260200180610a746023913960400191505060405180910390fd5b6001600160a01b03821661062e57604051600160e51b62461bcd0281526004018080602001828103825260218152602001806109df6021913960400191505060405180910390fd5b6001600160a01b03808416600081815260026020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b6001600160a01b0383166106d857604051600160e51b62461bcd028152600401808060200182810382526024815260200180610a506024913960400191505060405180910390fd5b6001600160a01b03821661072057604051600160e51b62461bcd028152600401808060200182810382526022815260200180610a2e6022913960400191505060405180910390fd5b6001600160a01b038316600090815260016020526040902054610749908263ffffffff6107da16565b6001600160a01b03808516600090815260016020526040808220939093559084168152205461077e908263ffffffff61097416565b6001600160a01b0380841660008181526001602090815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b6000828211156108345760408051600160e51b62461bcd02815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b600061084e846001600160a01b03166109d8565b61085a5750600161096c565b604051600160e11b634e8c461102815233600482018181526001600160a01b03888116602485015260448401879052608060648501908152865160848601528651600095928a1694639d188c229490938c938b938b939260a4019060208501908083838e5b838110156108d75781810151838201526020016108bf565b50505050905090810190601f1680156109045780820380516001836020036101000a031916815260200191505b5095505050505050602060405180830381600087803b15801561092657600080fd5b505af115801561093a573d6000803e3d6000fd5b505050506040513d602081101561095057600080fd5b50516001600160e01b031916600160e11b634e8c461102149150505b949350505050565b6000828201838110156109d15760408051600160e51b62461bcd02815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b3b15159056fe4b4950373a20617070726f766520746f20746865207a65726f20616464726573734b4950373a207472616e7366657220746f206e6f6e204b495037526563656976657220696d706c656d656e7465724b4950373a207472616e7366657220746f20746865207a65726f20616464726573734b4950373a207472616e736665722066726f6d20746865207a65726f20616464726573734b4950373a20617070726f76652066726f6d20746865207a65726f2061646472657373a165627a7a7230582066f4cb0cbcbfaa05f74c58ec5da5bced4ef103d17470fd2c83e6e93a39d4b0130029`

// KIP7FuncSigs maps the 4-byte function signature to its string representation.
var KIP7FuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"423f6cef": "safeTransfer(address,uint256)",
	"eb795549": "safeTransfer(address,uint256,bytes)",
	"42842e0e": "safeTransferFrom(address,address,uint256)",
	"b88d4fde": "safeTransferFrom(address,address,uint256,bytes)",
	"01ffc9a7": "supportsInterface(bytes4)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// KIP7Bin is the compiled bytecode used for deploying new contracts.
var KIP7Bin = "0x608060405234801561001057600080fd5b506100276301ffc9a760e01b61004260201b60201c565b61003d636578737160e01b61004260201b60201c565b610110565b7fffffffff0000000000000000000000000000000000000000000000000000000080821614156100d357604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f4b495031333a20696e76616c696420696e746572666163652069640000000000604482015290519081900360640190fd5b7fffffffff00000000000000000000000000000000000000000000000000000000166000908152602081905260409020805460ff19166001179055565b610ac28061011f6000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c806342842e0e1161007157806342842e0e1461019357806370a08231146101c9578063a9059cbb146101ef578063b88d4fde1461021b578063dd62ed3e146102e1578063eb7955491461030f576100a9565b806301ffc9a7146100ae578063095ea7b3146100e957806318160ddd1461011557806323b872dd1461012f578063423f6cef14610165575b600080fd5b6100d5600480360360208110156100c457600080fd5b50356001600160e01b0319166103ca565b604080519115158252519081900360200190f35b6100d5600480360360408110156100ff57600080fd5b506001600160a01b0381351690602001356103e9565b61011d6103ff565b60408051918252519081900360200190f35b6100d56004803603606081101561014557600080fd5b506001600160a01b03813581169160208101359091169060400135610405565b6101916004803603604081101561017b57600080fd5b506001600160a01b03813516906020013561045c565b005b610191600480360360608110156101a957600080fd5b506001600160a01b0381358116916020810135909116906040013561047a565b61011d600480360360208110156101df57600080fd5b50356001600160a01b031661049a565b6100d56004803603604081101561020557600080fd5b506001600160a01b0381351690602001356104b5565b6101916004803603608081101561023157600080fd5b6001600160a01b0382358116926020810135909116916040820135919081019060808101606082013564010000000081111561026c57600080fd5b82018360208201111561027e57600080fd5b803590602001918460018302840111640100000000831117156102a057600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506104c2945050505050565b61011d600480360360408110156102f757600080fd5b506001600160a01b038135811691602001351661051e565b6101916004803603606081101561032557600080fd5b6001600160a01b038235169160208101359181019060608101604082013564010000000081111561035557600080fd5b82018360208201111561036757600080fd5b8035906020019184600183028401116401000000008311171561038957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610549945050505050565b6001600160e01b03191660009081526020819052604090205460ff1690565b60006103f633848461059e565b50600192915050565b60035490565b6000610412848484610690565b6001600160a01b03841660009081526002602090815260408083203380855292529091205461045291869161044d908663ffffffff6107da16565b61059e565b5060019392505050565b610476828260405180602001604052806000815250610549565b5050565b610495838383604051806020016040528060008152506104c2565b505050565b6001600160a01b031660009081526001602052604090205490565b60006103f6338484610690565b6104cd848484610405565b506104da8484848461083a565b61051857604051600160e51b62461bcd02815260040180806020018281038252602e815260200180610a00602e913960400191505060405180910390fd5b50505050565b6001600160a01b03918216600090815260026020908152604080832093909416825291909152205490565b61055383836104b5565b506105603384848461083a565b61049557604051600160e51b62461bcd02815260040180806020018281038252602e815260200180610a00602e913960400191505060405180910390fd5b6001600160a01b0383166105e657604051600160e51b62461bcd028152600401808060200182810382526023815260200180610a746023913960400191505060405180910390fd5b6001600160a01b03821661062e57604051600160e51b62461bcd0281526004018080602001828103825260218152602001806109df6021913960400191505060405180910390fd5b6001600160a01b03808416600081815260026020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b6001600160a01b0383166106d857604051600160e51b62461bcd028152600401808060200182810382526024815260200180610a506024913960400191505060405180910390fd5b6001600160a01b03821661072057604051600160e51b62461bcd028152600401808060200182810382526022815260200180610a2e6022913960400191505060405180910390fd5b6001600160a01b038316600090815260016020526040902054610749908263ffffffff6107da16565b6001600160a01b03808516600090815260016020526040808220939093559084168152205461077e908263ffffffff61097416565b6001600160a01b0380841660008181526001602090815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b6000828211156108345760408051600160e51b62461bcd02815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b600061084e846001600160a01b03166109d8565b61085a5750600161096c565b604051600160e11b634e8c461102815233600482018181526001600160a01b03888116602485015260448401879052608060648501908152865160848601528651600095928a1694639d188c229490938c938b938b939260a4019060208501908083838e5b838110156108d75781810151838201526020016108bf565b50505050905090810190601f1680156109045780820380516001836020036101000a031916815260200191505b5095505050505050602060405180830381600087803b15801561092657600080fd5b505af115801561093a573d6000803e3d6000fd5b505050506040513d602081101561095057600080fd5b50516001600160e01b031916600160e11b634e8c461102149150505b949350505050565b6000828201838110156109d15760408051600160e51b62461bcd02815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b3b15159056fe4b4950373a20617070726f766520746f20746865207a65726f20616464726573734b4950373a207472616e7366657220746f206e6f6e204b495037526563656976657220696d706c656d656e7465724b4950373a207472616e7366657220746f20746865207a65726f20616464726573734b4950373a207472616e736665722066726f6d20746865207a65726f20616464726573734b4950373a20617070726f76652066726f6d20746865207a65726f2061646472657373a165627a7a7230582066f4cb0cbcbfaa05f74c58ec5da5bced4ef103d17470fd2c83e6e93a39d4b0130029"

// DeployKIP7 deploys a new Klaytn contract, binding an instance of KIP7 to it.
func DeployKIP7(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *KIP7, error) {
	parsed, err := abi.JSON(strings.NewReader(KIP7ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(KIP7Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KIP7{KIP7Caller: KIP7Caller{contract: contract}, KIP7Transactor: KIP7Transactor{contract: contract}, KIP7Filterer: KIP7Filterer{contract: contract}}, nil
}

// KIP7 is an auto generated Go binding around a Klaytn contract.
type KIP7 struct {
	KIP7Caller     // Read-only binding to the contract
	KIP7Transactor // Write-only binding to the contract
	KIP7Filterer   // Log filterer for contract events
}

// KIP7Caller is an auto generated read-only Go binding around a Klaytn contract.
type KIP7Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KIP7Transactor is an auto generated write-only Go binding around a Klaytn contract.
type KIP7Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KIP7Filterer is an auto generated log filtering Go binding around a Klaytn contract events.
type KIP7Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KIP7Session is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type KIP7Session struct {
	Contract     *KIP7             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KIP7CallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type KIP7CallerSession struct {
	Contract *KIP7Caller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// KIP7TransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type KIP7TransactorSession struct {
	Contract     *KIP7Transactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KIP7Raw is an auto generated low-level Go binding around a Klaytn contract.
type KIP7Raw struct {
	Contract *KIP7 // Generic contract binding to access the raw methods on
}

// KIP7CallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type KIP7CallerRaw struct {
	Contract *KIP7Caller // Generic read-only contract binding to access the raw methods on
}

// KIP7TransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type KIP7TransactorRaw struct {
	Contract *KIP7Transactor // Generic write-only contract binding to access the raw methods on
}

// NewKIP7 creates a new instance of KIP7, bound to a specific deployed contract.
func NewKIP7(address common.Address, backend bind.ContractBackend) (*KIP7, error) {
	contract, err := bindKIP7(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KIP7{KIP7Caller: KIP7Caller{contract: contract}, KIP7Transactor: KIP7Transactor{contract: contract}, KIP7Filterer: KIP7Filterer{contract: contract}}, nil
}

// NewKIP7Caller creates a new read-only instance of KIP7, bound to a specific deployed contract.
func NewKIP7Caller(address common.Address, caller bind.ContractCaller) (*KIP7Caller, error) {
	contract, err := bindKIP7(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KIP7Caller{contract: contract}, nil
}

// NewKIP7Transactor creates a new write-only instance of KIP7, bound to a specific deployed contract.
func NewKIP7Transactor(address common.Address, transactor bind.ContractTransactor) (*KIP7Transactor, error) {
	contract, err := bindKIP7(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KIP7Transactor{contract: contract}, nil
}

// NewKIP7Filterer creates a new log filterer instance of KIP7, bound to a specific deployed contract.
func NewKIP7Filterer(address common.Address, filterer bind.ContractFilterer) (*KIP7Filterer, error) {
	contract, err := bindKIP7(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KIP7Filterer{contract: contract}, nil
}

// bindKIP7 binds a generic wrapper to an already deployed contract.
func bindKIP7(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KIP7ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KIP7 *KIP7Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KIP7.Contract.KIP7Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KIP7 *KIP7Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KIP7.Contract.KIP7Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KIP7 *KIP7Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KIP7.Contract.KIP7Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KIP7 *KIP7CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KIP7.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KIP7 *KIP7TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KIP7.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KIP7 *KIP7TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KIP7.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_KIP7 *KIP7Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KIP7.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_KIP7 *KIP7Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _KIP7.Contract.Allowance(&_KIP7.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_KIP7 *KIP7CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _KIP7.Contract.Allowance(&_KIP7.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_KIP7 *KIP7Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KIP7.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_KIP7 *KIP7Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _KIP7.Contract.BalanceOf(&_KIP7.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_KIP7 *KIP7CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _KIP7.Contract.BalanceOf(&_KIP7.CallOpts, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_KIP7 *KIP7Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KIP7.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_KIP7 *KIP7Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _KIP7.Contract.SupportsInterface(&_KIP7.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_KIP7 *KIP7CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _KIP7.Contract.SupportsInterface(&_KIP7.CallOpts, interfaceId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_KIP7 *KIP7Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KIP7.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_KIP7 *KIP7Session) TotalSupply() (*big.Int, error) {
	return _KIP7.Contract.TotalSupply(&_KIP7.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_KIP7 *KIP7CallerSession) TotalSupply() (*big.Int, error) {
	return _KIP7.Contract.TotalSupply(&_KIP7.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_KIP7 *KIP7Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _KIP7.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_KIP7 *KIP7Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _KIP7.Contract.Approve(&_KIP7.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_KIP7 *KIP7TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _KIP7.Contract.Approve(&_KIP7.TransactOpts, spender, value)
}

// SafeTransfer is a paid mutator transaction binding the contract method 0x423f6cef.
//
// Solidity: function safeTransfer(address recipient, uint256 amount) returns()
func (_KIP7 *KIP7Transactor) SafeTransfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KIP7.contract.Transact(opts, "safeTransfer", recipient, amount)
}

// SafeTransfer is a paid mutator transaction binding the contract method 0x423f6cef.
//
// Solidity: function safeTransfer(address recipient, uint256 amount) returns()
func (_KIP7 *KIP7Session) SafeTransfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KIP7.Contract.SafeTransfer(&_KIP7.TransactOpts, recipient, amount)
}

// SafeTransfer is a paid mutator transaction binding the contract method 0x423f6cef.
//
// Solidity: function safeTransfer(address recipient, uint256 amount) returns()
func (_KIP7 *KIP7TransactorSession) SafeTransfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KIP7.Contract.SafeTransfer(&_KIP7.TransactOpts, recipient, amount)
}

// SafeTransfer0 is a paid mutator transaction binding the contract method 0xeb795549.
//
// Solidity: function safeTransfer(address recipient, uint256 amount, bytes data) returns()
func (_KIP7 *KIP7Transactor) SafeTransfer0(opts *bind.TransactOpts, recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _KIP7.contract.Transact(opts, "safeTransfer0", recipient, amount, data)
}

// SafeTransfer0 is a paid mutator transaction binding the contract method 0xeb795549.
//
// Solidity: function safeTransfer(address recipient, uint256 amount, bytes data) returns()
func (_KIP7 *KIP7Session) SafeTransfer0(recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _KIP7.Contract.SafeTransfer0(&_KIP7.TransactOpts, recipient, amount, data)
}

// SafeTransfer0 is a paid mutator transaction binding the contract method 0xeb795549.
//
// Solidity: function safeTransfer(address recipient, uint256 amount, bytes data) returns()
func (_KIP7 *KIP7TransactorSession) SafeTransfer0(recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _KIP7.Contract.SafeTransfer0(&_KIP7.TransactOpts, recipient, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address sender, address recipient, uint256 amount) returns()
func (_KIP7 *KIP7Transactor) SafeTransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KIP7.contract.Transact(opts, "safeTransferFrom", sender, recipient, amount)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address sender, address recipient, uint256 amount) returns()
func (_KIP7 *KIP7Session) SafeTransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KIP7.Contract.SafeTransferFrom(&_KIP7.TransactOpts, sender, recipient, amount)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address sender, address recipient, uint256 amount) returns()
func (_KIP7 *KIP7TransactorSession) SafeTransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KIP7.Contract.SafeTransferFrom(&_KIP7.TransactOpts, sender, recipient, amount)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address sender, address recipient, uint256 amount, bytes data) returns()
func (_KIP7 *KIP7Transactor) SafeTransferFrom0(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _KIP7.contract.Transact(opts, "safeTransferFrom0", sender, recipient, amount, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address sender, address recipient, uint256 amount, bytes data) returns()
func (_KIP7 *KIP7Session) SafeTransferFrom0(sender common.Address, recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _KIP7.Contract.SafeTransferFrom0(&_KIP7.TransactOpts, sender, recipient, amount, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address sender, address recipient, uint256 amount, bytes data) returns()
func (_KIP7 *KIP7TransactorSession) SafeTransferFrom0(sender common.Address, recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _KIP7.Contract.SafeTransferFrom0(&_KIP7.TransactOpts, sender, recipient, amount, data)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_KIP7 *KIP7Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KIP7.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_KIP7 *KIP7Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KIP7.Contract.Transfer(&_KIP7.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_KIP7 *KIP7TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KIP7.Contract.Transfer(&_KIP7.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_KIP7 *KIP7Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KIP7.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_KIP7 *KIP7Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KIP7.Contract.TransferFrom(&_KIP7.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_KIP7 *KIP7TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KIP7.Contract.TransferFrom(&_KIP7.TransactOpts, sender, recipient, amount)
}

// KIP7ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the KIP7 contract.
type KIP7ApprovalIterator struct {
	Event *KIP7Approval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KIP7ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KIP7Approval)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KIP7Approval)
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
func (it *KIP7ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KIP7ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KIP7Approval represents a Approval event raised by the KIP7 contract.
type KIP7Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_KIP7 *KIP7Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*KIP7ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _KIP7.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &KIP7ApprovalIterator{contract: _KIP7.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_KIP7 *KIP7Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *KIP7Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _KIP7.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KIP7Approval)
				if err := _KIP7.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_KIP7 *KIP7Filterer) ParseApproval(log types.Log) (*KIP7Approval, error) {
	event := new(KIP7Approval)
	if err := _KIP7.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KIP7TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the KIP7 contract.
type KIP7TransferIterator struct {
	Event *KIP7Transfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KIP7TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KIP7Transfer)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KIP7Transfer)
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
func (it *KIP7TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KIP7TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KIP7Transfer represents a Transfer event raised by the KIP7 contract.
type KIP7Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_KIP7 *KIP7Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*KIP7TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _KIP7.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &KIP7TransferIterator{contract: _KIP7.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_KIP7 *KIP7Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *KIP7Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _KIP7.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KIP7Transfer)
				if err := _KIP7.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_KIP7 *KIP7Filterer) ParseTransfer(log types.Log) (*KIP7Transfer, error) {
	event := new(KIP7Transfer)
	if err := _KIP7.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KIP7BurnableABI is the input ABI used to generate the binding from.
const KIP7BurnableABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"safeTransfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// KIP7BurnableBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const KIP7BurnableBinRuntime = `608060405234801561001057600080fd5b50600436106100cf5760003560e01c806342966c681161008c578063a9059cbb11610066578063a9059cbb1461025e578063b88d4fde1461028a578063dd62ed3e14610350578063eb7955491461037e576100cf565b806342966c68146101ef57806370a082311461020c57806379cc679014610232576100cf565b806301ffc9a7146100d4578063095ea7b31461010f57806318160ddd1461013b57806323b872dd14610155578063423f6cef1461018b57806342842e0e146101b9575b600080fd5b6100fb600480360360208110156100ea57600080fd5b50356001600160e01b031916610439565b604080519115158252519081900360200190f35b6100fb6004803603604081101561012557600080fd5b506001600160a01b038135169060200135610458565b61014361046e565b60408051918252519081900360200190f35b6100fb6004803603606081101561016b57600080fd5b506001600160a01b03813581169160208101359091169060400135610474565b6101b7600480360360408110156101a157600080fd5b506001600160a01b0381351690602001356104cb565b005b6101b7600480360360608110156101cf57600080fd5b506001600160a01b038135811691602081013590911690604001356104e9565b6101b76004803603602081101561020557600080fd5b5035610509565b6101436004803603602081101561022257600080fd5b50356001600160a01b0316610516565b6101b76004803603604081101561024857600080fd5b506001600160a01b038135169060200135610531565b6100fb6004803603604081101561027457600080fd5b506001600160a01b03813516906020013561053b565b6101b7600480360360808110156102a057600080fd5b6001600160a01b038235811692602081013590911691604082013591908101906080810160608201356401000000008111156102db57600080fd5b8201836020820111156102ed57600080fd5b8035906020019184600183028401116401000000008311171561030f57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610548945050505050565b6101436004803603604081101561036657600080fd5b506001600160a01b03813581169160200135166105a4565b6101b76004803603606081101561039457600080fd5b6001600160a01b03823516916020810135918101906060810160408201356401000000008111156103c457600080fd5b8201836020820111156103d657600080fd5b803590602001918460018302840111640100000000831117156103f857600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506105cf945050505050565b6001600160e01b03191660009081526020819052604090205460ff1690565b6000610465338484610624565b50600192915050565b60035490565b6000610481848484610716565b6001600160a01b0384166000908152600260209081526040808320338085529252909120546104c19186916104bc908663ffffffff61086016565b610624565b5060019392505050565b6104e58282604051806020016040528060008152506105cf565b5050565b61050483838360405180602001604052806000815250610548565b505050565b61051333826108c0565b50565b6001600160a01b031660009081526001602052604090205490565b6104e582826109b4565b6000610465338484610716565b610553848484610474565b50610560848484846109f9565b61059e57604051600160e51b62461bcd02815260040180806020018281038252602e815260200180610bbf602e913960400191505060405180910390fd5b50505050565b6001600160a01b03918216600090815260026020908152604080832093909416825291909152205490565b6105d9838361053b565b506105e6338484846109f9565b61050457604051600160e51b62461bcd02815260040180806020018281038252602e815260200180610bbf602e913960400191505060405180910390fd5b6001600160a01b03831661066c57604051600160e51b62461bcd028152600401808060200182810382526023815260200180610c336023913960400191505060405180910390fd5b6001600160a01b0382166106b457604051600160e51b62461bcd028152600401808060200182810382526021815260200180610b9e6021913960400191505060405180910390fd5b6001600160a01b03808416600081815260026020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b6001600160a01b03831661075e57604051600160e51b62461bcd028152600401808060200182810382526024815260200180610c0f6024913960400191505060405180910390fd5b6001600160a01b0382166107a657604051600160e51b62461bcd028152600401808060200182810382526022815260200180610bed6022913960400191505060405180910390fd5b6001600160a01b0383166000908152600160205260409020546107cf908263ffffffff61086016565b6001600160a01b038085166000908152600160205260408082209390935590841681522054610804908263ffffffff610b3316565b6001600160a01b0380841660008181526001602090815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b6000828211156108ba5760408051600160e51b62461bcd02815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b6001600160a01b03821661091e5760408051600160e51b62461bcd02815260206004820181905260248201527f4b4950373a206275726e2066726f6d20746865207a65726f2061646472657373604482015290519081900360640190fd5b600354610931908263ffffffff61086016565b6003556001600160a01b03821660009081526001602052604090205461095d908263ffffffff61086016565b6001600160a01b0383166000818152600160209081526040808320949094558351858152935191937fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929081900390910190a35050565b6109be82826108c0565b6001600160a01b0382166000908152600260209081526040808320338085529252909120546104e59184916104bc908563ffffffff61086016565b6000610a0d846001600160a01b0316610b97565b610a1957506001610b2b565b604051600160e11b634e8c461102815233600482018181526001600160a01b03888116602485015260448401879052608060648501908152865160848601528651600095928a1694639d188c229490938c938b938b939260a4019060208501908083838e5b83811015610a96578181015183820152602001610a7e565b50505050905090810190601f168015610ac35780820380516001836020036101000a031916815260200191505b5095505050505050602060405180830381600087803b158015610ae557600080fd5b505af1158015610af9573d6000803e3d6000fd5b505050506040513d6020811015610b0f57600080fd5b50516001600160e01b031916600160e11b634e8c461102149150505b949350505050565b600082820183811015610b905760408051600160e51b62461bcd02815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b3b15159056fe4b4950373a20617070726f766520746f20746865207a65726f20616464726573734b4950373a207472616e7366657220746f206e6f6e204b495037526563656976657220696d706c656d656e7465724b4950373a207472616e7366657220746f20746865207a65726f20616464726573734b4950373a207472616e736665722066726f6d20746865207a65726f20616464726573734b4950373a20617070726f76652066726f6d20746865207a65726f2061646472657373a165627a7a72305820eeeb66871396261f9948a28c631ceda88f1d2b4b7c47fbb7e6eb4f48faf1b1a00029`

// KIP7BurnableFuncSigs maps the 4-byte function signature to its string representation.
var KIP7BurnableFuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"42966c68": "burn(uint256)",
	"79cc6790": "burnFrom(address,uint256)",
	"423f6cef": "safeTransfer(address,uint256)",
	"eb795549": "safeTransfer(address,uint256,bytes)",
	"42842e0e": "safeTransferFrom(address,address,uint256)",
	"b88d4fde": "safeTransferFrom(address,address,uint256,bytes)",
	"01ffc9a7": "supportsInterface(bytes4)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// KIP7BurnableBin is the compiled bytecode used for deploying new contracts.
var KIP7BurnableBin = "0x608060405234801561001057600080fd5b506100276301ffc9a760e01b61005860201b60201c565b61003d636578737160e01b61005860201b60201c565b610053633b5a0bf860e01b61005860201b60201c565b610126565b7fffffffff0000000000000000000000000000000000000000000000000000000080821614156100e957604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f4b495031333a20696e76616c696420696e746572666163652069640000000000604482015290519081900360640190fd5b7fffffffff00000000000000000000000000000000000000000000000000000000166000908152602081905260409020805460ff19166001179055565b610c81806101356000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c806342966c681161008c578063a9059cbb11610066578063a9059cbb1461025e578063b88d4fde1461028a578063dd62ed3e14610350578063eb7955491461037e576100cf565b806342966c68146101ef57806370a082311461020c57806379cc679014610232576100cf565b806301ffc9a7146100d4578063095ea7b31461010f57806318160ddd1461013b57806323b872dd14610155578063423f6cef1461018b57806342842e0e146101b9575b600080fd5b6100fb600480360360208110156100ea57600080fd5b50356001600160e01b031916610439565b604080519115158252519081900360200190f35b6100fb6004803603604081101561012557600080fd5b506001600160a01b038135169060200135610458565b61014361046e565b60408051918252519081900360200190f35b6100fb6004803603606081101561016b57600080fd5b506001600160a01b03813581169160208101359091169060400135610474565b6101b7600480360360408110156101a157600080fd5b506001600160a01b0381351690602001356104cb565b005b6101b7600480360360608110156101cf57600080fd5b506001600160a01b038135811691602081013590911690604001356104e9565b6101b76004803603602081101561020557600080fd5b5035610509565b6101436004803603602081101561022257600080fd5b50356001600160a01b0316610516565b6101b76004803603604081101561024857600080fd5b506001600160a01b038135169060200135610531565b6100fb6004803603604081101561027457600080fd5b506001600160a01b03813516906020013561053b565b6101b7600480360360808110156102a057600080fd5b6001600160a01b038235811692602081013590911691604082013591908101906080810160608201356401000000008111156102db57600080fd5b8201836020820111156102ed57600080fd5b8035906020019184600183028401116401000000008311171561030f57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610548945050505050565b6101436004803603604081101561036657600080fd5b506001600160a01b03813581169160200135166105a4565b6101b76004803603606081101561039457600080fd5b6001600160a01b03823516916020810135918101906060810160408201356401000000008111156103c457600080fd5b8201836020820111156103d657600080fd5b803590602001918460018302840111640100000000831117156103f857600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506105cf945050505050565b6001600160e01b03191660009081526020819052604090205460ff1690565b6000610465338484610624565b50600192915050565b60035490565b6000610481848484610716565b6001600160a01b0384166000908152600260209081526040808320338085529252909120546104c19186916104bc908663ffffffff61086016565b610624565b5060019392505050565b6104e58282604051806020016040528060008152506105cf565b5050565b61050483838360405180602001604052806000815250610548565b505050565b61051333826108c0565b50565b6001600160a01b031660009081526001602052604090205490565b6104e582826109b4565b6000610465338484610716565b610553848484610474565b50610560848484846109f9565b61059e57604051600160e51b62461bcd02815260040180806020018281038252602e815260200180610bbf602e913960400191505060405180910390fd5b50505050565b6001600160a01b03918216600090815260026020908152604080832093909416825291909152205490565b6105d9838361053b565b506105e6338484846109f9565b61050457604051600160e51b62461bcd02815260040180806020018281038252602e815260200180610bbf602e913960400191505060405180910390fd5b6001600160a01b03831661066c57604051600160e51b62461bcd028152600401808060200182810382526023815260200180610c336023913960400191505060405180910390fd5b6001600160a01b0382166106b457604051600160e51b62461bcd028152600401808060200182810382526021815260200180610b9e6021913960400191505060405180910390fd5b6001600160a01b03808416600081815260026020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b6001600160a01b03831661075e57604051600160e51b62461bcd028152600401808060200182810382526024815260200180610c0f6024913960400191505060405180910390fd5b6001600160a01b0382166107a657604051600160e51b62461bcd028152600401808060200182810382526022815260200180610bed6022913960400191505060405180910390fd5b6001600160a01b0383166000908152600160205260409020546107cf908263ffffffff61086016565b6001600160a01b038085166000908152600160205260408082209390935590841681522054610804908263ffffffff610b3316565b6001600160a01b0380841660008181526001602090815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b6000828211156108ba5760408051600160e51b62461bcd02815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b6001600160a01b03821661091e5760408051600160e51b62461bcd02815260206004820181905260248201527f4b4950373a206275726e2066726f6d20746865207a65726f2061646472657373604482015290519081900360640190fd5b600354610931908263ffffffff61086016565b6003556001600160a01b03821660009081526001602052604090205461095d908263ffffffff61086016565b6001600160a01b0383166000818152600160209081526040808320949094558351858152935191937fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929081900390910190a35050565b6109be82826108c0565b6001600160a01b0382166000908152600260209081526040808320338085529252909120546104e59184916104bc908563ffffffff61086016565b6000610a0d846001600160a01b0316610b97565b610a1957506001610b2b565b604051600160e11b634e8c461102815233600482018181526001600160a01b03888116602485015260448401879052608060648501908152865160848601528651600095928a1694639d188c229490938c938b938b939260a4019060208501908083838e5b83811015610a96578181015183820152602001610a7e565b50505050905090810190601f168015610ac35780820380516001836020036101000a031916815260200191505b5095505050505050602060405180830381600087803b158015610ae557600080fd5b505af1158015610af9573d6000803e3d6000fd5b505050506040513d6020811015610b0f57600080fd5b50516001600160e01b031916600160e11b634e8c461102149150505b949350505050565b600082820183811015610b905760408051600160e51b62461bcd02815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b3b15159056fe4b4950373a20617070726f766520746f20746865207a65726f20616464726573734b4950373a207472616e7366657220746f206e6f6e204b495037526563656976657220696d706c656d656e7465724b4950373a207472616e7366657220746f20746865207a65726f20616464726573734b4950373a207472616e736665722066726f6d20746865207a65726f20616464726573734b4950373a20617070726f76652066726f6d20746865207a65726f2061646472657373a165627a7a72305820eeeb66871396261f9948a28c631ceda88f1d2b4b7c47fbb7e6eb4f48faf1b1a00029"

// DeployKIP7Burnable deploys a new Klaytn contract, binding an instance of KIP7Burnable to it.
func DeployKIP7Burnable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *KIP7Burnable, error) {
	parsed, err := abi.JSON(strings.NewReader(KIP7BurnableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(KIP7BurnableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KIP7Burnable{KIP7BurnableCaller: KIP7BurnableCaller{contract: contract}, KIP7BurnableTransactor: KIP7BurnableTransactor{contract: contract}, KIP7BurnableFilterer: KIP7BurnableFilterer{contract: contract}}, nil
}

// KIP7Burnable is an auto generated Go binding around a Klaytn contract.
type KIP7Burnable struct {
	KIP7BurnableCaller     // Read-only binding to the contract
	KIP7BurnableTransactor // Write-only binding to the contract
	KIP7BurnableFilterer   // Log filterer for contract events
}

// KIP7BurnableCaller is an auto generated read-only Go binding around a Klaytn contract.
type KIP7BurnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KIP7BurnableTransactor is an auto generated write-only Go binding around a Klaytn contract.
type KIP7BurnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KIP7BurnableFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type KIP7BurnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KIP7BurnableSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type KIP7BurnableSession struct {
	Contract     *KIP7Burnable     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KIP7BurnableCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type KIP7BurnableCallerSession struct {
	Contract *KIP7BurnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// KIP7BurnableTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type KIP7BurnableTransactorSession struct {
	Contract     *KIP7BurnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// KIP7BurnableRaw is an auto generated low-level Go binding around a Klaytn contract.
type KIP7BurnableRaw struct {
	Contract *KIP7Burnable // Generic contract binding to access the raw methods on
}

// KIP7BurnableCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type KIP7BurnableCallerRaw struct {
	Contract *KIP7BurnableCaller // Generic read-only contract binding to access the raw methods on
}

// KIP7BurnableTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type KIP7BurnableTransactorRaw struct {
	Contract *KIP7BurnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKIP7Burnable creates a new instance of KIP7Burnable, bound to a specific deployed contract.
func NewKIP7Burnable(address common.Address, backend bind.ContractBackend) (*KIP7Burnable, error) {
	contract, err := bindKIP7Burnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KIP7Burnable{KIP7BurnableCaller: KIP7BurnableCaller{contract: contract}, KIP7BurnableTransactor: KIP7BurnableTransactor{contract: contract}, KIP7BurnableFilterer: KIP7BurnableFilterer{contract: contract}}, nil
}

// NewKIP7BurnableCaller creates a new read-only instance of KIP7Burnable, bound to a specific deployed contract.
func NewKIP7BurnableCaller(address common.Address, caller bind.ContractCaller) (*KIP7BurnableCaller, error) {
	contract, err := bindKIP7Burnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KIP7BurnableCaller{contract: contract}, nil
}

// NewKIP7BurnableTransactor creates a new write-only instance of KIP7Burnable, bound to a specific deployed contract.
func NewKIP7BurnableTransactor(address common.Address, transactor bind.ContractTransactor) (*KIP7BurnableTransactor, error) {
	contract, err := bindKIP7Burnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KIP7BurnableTransactor{contract: contract}, nil
}

// NewKIP7BurnableFilterer creates a new log filterer instance of KIP7Burnable, bound to a specific deployed contract.
func NewKIP7BurnableFilterer(address common.Address, filterer bind.ContractFilterer) (*KIP7BurnableFilterer, error) {
	contract, err := bindKIP7Burnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KIP7BurnableFilterer{contract: contract}, nil
}

// bindKIP7Burnable binds a generic wrapper to an already deployed contract.
func bindKIP7Burnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KIP7BurnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KIP7Burnable *KIP7BurnableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KIP7Burnable.Contract.KIP7BurnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KIP7Burnable *KIP7BurnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KIP7Burnable.Contract.KIP7BurnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KIP7Burnable *KIP7BurnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KIP7Burnable.Contract.KIP7BurnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KIP7Burnable *KIP7BurnableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KIP7Burnable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KIP7Burnable *KIP7BurnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KIP7Burnable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KIP7Burnable *KIP7BurnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KIP7Burnable.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_KIP7Burnable *KIP7BurnableCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KIP7Burnable.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_KIP7Burnable *KIP7BurnableSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _KIP7Burnable.Contract.Allowance(&_KIP7Burnable.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_KIP7Burnable *KIP7BurnableCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _KIP7Burnable.Contract.Allowance(&_KIP7Burnable.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_KIP7Burnable *KIP7BurnableCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KIP7Burnable.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_KIP7Burnable *KIP7BurnableSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _KIP7Burnable.Contract.BalanceOf(&_KIP7Burnable.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_KIP7Burnable *KIP7BurnableCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _KIP7Burnable.Contract.BalanceOf(&_KIP7Burnable.CallOpts, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_KIP7Burnable *KIP7BurnableCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KIP7Burnable.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_KIP7Burnable *KIP7BurnableSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _KIP7Burnable.Contract.SupportsInterface(&_KIP7Burnable.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_KIP7Burnable *KIP7BurnableCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _KIP7Burnable.Contract.SupportsInterface(&_KIP7Burnable.CallOpts, interfaceId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_KIP7Burnable *KIP7BurnableCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KIP7Burnable.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_KIP7Burnable *KIP7BurnableSession) TotalSupply() (*big.Int, error) {
	return _KIP7Burnable.Contract.TotalSupply(&_KIP7Burnable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_KIP7Burnable *KIP7BurnableCallerSession) TotalSupply() (*big.Int, error) {
	return _KIP7Burnable.Contract.TotalSupply(&_KIP7Burnable.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_KIP7Burnable *KIP7BurnableTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _KIP7Burnable.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_KIP7Burnable *KIP7BurnableSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _KIP7Burnable.Contract.Approve(&_KIP7Burnable.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_KIP7Burnable *KIP7BurnableTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _KIP7Burnable.Contract.Approve(&_KIP7Burnable.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_KIP7Burnable *KIP7BurnableTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _KIP7Burnable.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_KIP7Burnable *KIP7BurnableSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _KIP7Burnable.Contract.Burn(&_KIP7Burnable.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_KIP7Burnable *KIP7BurnableTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _KIP7Burnable.Contract.Burn(&_KIP7Burnable.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_KIP7Burnable *KIP7BurnableTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KIP7Burnable.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_KIP7Burnable *KIP7BurnableSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KIP7Burnable.Contract.BurnFrom(&_KIP7Burnable.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_KIP7Burnable *KIP7BurnableTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KIP7Burnable.Contract.BurnFrom(&_KIP7Burnable.TransactOpts, account, amount)
}

// SafeTransfer is a paid mutator transaction binding the contract method 0x423f6cef.
//
// Solidity: function safeTransfer(address recipient, uint256 amount) returns()
func (_KIP7Burnable *KIP7BurnableTransactor) SafeTransfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KIP7Burnable.contract.Transact(opts, "safeTransfer", recipient, amount)
}

// SafeTransfer is a paid mutator transaction binding the contract method 0x423f6cef.
//
// Solidity: function safeTransfer(address recipient, uint256 amount) returns()
func (_KIP7Burnable *KIP7BurnableSession) SafeTransfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KIP7Burnable.Contract.SafeTransfer(&_KIP7Burnable.TransactOpts, recipient, amount)
}

// SafeTransfer is a paid mutator transaction binding the contract method 0x423f6cef.
//
// Solidity: function safeTransfer(address recipient, uint256 amount) returns()
func (_KIP7Burnable *KIP7BurnableTransactorSession) SafeTransfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KIP7Burnable.Contract.SafeTransfer(&_KIP7Burnable.TransactOpts, recipient, amount)
}

// SafeTransfer0 is a paid mutator transaction binding the contract method 0xeb795549.
//
// Solidity: function safeTransfer(address recipient, uint256 amount, bytes data) returns()
func (_KIP7Burnable *KIP7BurnableTransactor) SafeTransfer0(opts *bind.TransactOpts, recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _KIP7Burnable.contract.Transact(opts, "safeTransfer0", recipient, amount, data)
}

// SafeTransfer0 is a paid mutator transaction binding the contract method 0xeb795549.
//
// Solidity: function safeTransfer(address recipient, uint256 amount, bytes data) returns()
func (_KIP7Burnable *KIP7BurnableSession) SafeTransfer0(recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _KIP7Burnable.Contract.SafeTransfer0(&_KIP7Burnable.TransactOpts, recipient, amount, data)
}

// SafeTransfer0 is a paid mutator transaction binding the contract method 0xeb795549.
//
// Solidity: function safeTransfer(address recipient, uint256 amount, bytes data) returns()
func (_KIP7Burnable *KIP7BurnableTransactorSession) SafeTransfer0(recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _KIP7Burnable.Contract.SafeTransfer0(&_KIP7Burnable.TransactOpts, recipient, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address sender, address recipient, uint256 amount) returns()
func (_KIP7Burnable *KIP7BurnableTransactor) SafeTransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KIP7Burnable.contract.Transact(opts, "safeTransferFrom", sender, recipient, amount)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address sender, address recipient, uint256 amount) returns()
func (_KIP7Burnable *KIP7BurnableSession) SafeTransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KIP7Burnable.Contract.SafeTransferFrom(&_KIP7Burnable.TransactOpts, sender, recipient, amount)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address sender, address recipient, uint256 amount) returns()
func (_KIP7Burnable *KIP7BurnableTransactorSession) SafeTransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KIP7Burnable.Contract.SafeTransferFrom(&_KIP7Burnable.TransactOpts, sender, recipient, amount)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address sender, address recipient, uint256 amount, bytes data) returns()
func (_KIP7Burnable *KIP7BurnableTransactor) SafeTransferFrom0(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _KIP7Burnable.contract.Transact(opts, "safeTransferFrom0", sender, recipient, amount, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address sender, address recipient, uint256 amount, bytes data) returns()
func (_KIP7Burnable *KIP7BurnableSession) SafeTransferFrom0(sender common.Address, recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _KIP7Burnable.Contract.SafeTransferFrom0(&_KIP7Burnable.TransactOpts, sender, recipient, amount, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address sender, address recipient, uint256 amount, bytes data) returns()
func (_KIP7Burnable *KIP7BurnableTransactorSession) SafeTransferFrom0(sender common.Address, recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _KIP7Burnable.Contract.SafeTransferFrom0(&_KIP7Burnable.TransactOpts, sender, recipient, amount, data)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_KIP7Burnable *KIP7BurnableTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KIP7Burnable.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_KIP7Burnable *KIP7BurnableSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KIP7Burnable.Contract.Transfer(&_KIP7Burnable.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_KIP7Burnable *KIP7BurnableTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KIP7Burnable.Contract.Transfer(&_KIP7Burnable.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_KIP7Burnable *KIP7BurnableTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KIP7Burnable.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_KIP7Burnable *KIP7BurnableSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KIP7Burnable.Contract.TransferFrom(&_KIP7Burnable.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_KIP7Burnable *KIP7BurnableTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KIP7Burnable.Contract.TransferFrom(&_KIP7Burnable.TransactOpts, sender, recipient, amount)
}

// KIP7BurnableApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the KIP7Burnable contract.
type KIP7BurnableApprovalIterator struct {
	Event *KIP7BurnableApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KIP7BurnableApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KIP7BurnableApproval)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KIP7BurnableApproval)
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
func (it *KIP7BurnableApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KIP7BurnableApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KIP7BurnableApproval represents a Approval event raised by the KIP7Burnable contract.
type KIP7BurnableApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_KIP7Burnable *KIP7BurnableFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*KIP7BurnableApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _KIP7Burnable.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &KIP7BurnableApprovalIterator{contract: _KIP7Burnable.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_KIP7Burnable *KIP7BurnableFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *KIP7BurnableApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _KIP7Burnable.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KIP7BurnableApproval)
				if err := _KIP7Burnable.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_KIP7Burnable *KIP7BurnableFilterer) ParseApproval(log types.Log) (*KIP7BurnableApproval, error) {
	event := new(KIP7BurnableApproval)
	if err := _KIP7Burnable.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KIP7BurnableTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the KIP7Burnable contract.
type KIP7BurnableTransferIterator struct {
	Event *KIP7BurnableTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KIP7BurnableTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KIP7BurnableTransfer)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KIP7BurnableTransfer)
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
func (it *KIP7BurnableTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KIP7BurnableTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KIP7BurnableTransfer represents a Transfer event raised by the KIP7Burnable contract.
type KIP7BurnableTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_KIP7Burnable *KIP7BurnableFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*KIP7BurnableTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _KIP7Burnable.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &KIP7BurnableTransferIterator{contract: _KIP7Burnable.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_KIP7Burnable *KIP7BurnableFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *KIP7BurnableTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _KIP7Burnable.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KIP7BurnableTransfer)
				if err := _KIP7Burnable.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_KIP7Burnable *KIP7BurnableFilterer) ParseTransfer(log types.Log) (*KIP7BurnableTransfer, error) {
	event := new(KIP7BurnableTransfer)
	if err := _KIP7Burnable.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KIP7MintableABI is the input ABI used to generate the binding from.
const KIP7MintableABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"safeTransfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addMinter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceMinter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isMinter\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MinterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MinterRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// KIP7MintableBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const KIP7MintableBinRuntime = `608060405234801561001057600080fd5b50600436106100f55760003560e01c806370a0823111610097578063aa271e1a11610066578063aa271e1a146102c1578063b88d4fde146102e7578063dd62ed3e146103ad578063eb795549146103db576100f5565b806370a0823114610241578063983b2d5614610267578063986502751461028d578063a9059cbb14610295576100f5565b806323b872dd116100d357806323b872dd1461017b57806340c10f19146101b1578063423f6cef146101dd57806342842e0e1461020b576100f5565b806301ffc9a7146100fa578063095ea7b31461013557806318160ddd14610161575b600080fd5b6101216004803603602081101561011057600080fd5b50356001600160e01b031916610496565b604080519115158252519081900360200190f35b6101216004803603604081101561014b57600080fd5b506001600160a01b0381351690602001356104b5565b6101696104cb565b60408051918252519081900360200190f35b6101216004803603606081101561019157600080fd5b506001600160a01b038135811691602081013590911690604001356104d1565b610121600480360360408110156101c757600080fd5b506001600160a01b038135169060200135610528565b610209600480360360408110156101f357600080fd5b506001600160a01b03813516906020013561057b565b005b6102096004803603606081101561022157600080fd5b506001600160a01b03813581169160208101359091169060400135610599565b6101696004803603602081101561025757600080fd5b50356001600160a01b03166105b9565b6102096004803603602081101561027d57600080fd5b50356001600160a01b03166105d4565b610209610627565b610121600480360360408110156102ab57600080fd5b506001600160a01b038135169060200135610632565b610121600480360360208110156102d757600080fd5b50356001600160a01b031661063f565b610209600480360360808110156102fd57600080fd5b6001600160a01b0382358116926020810135909116916040820135919081019060808101606082013564010000000081111561033857600080fd5b82018360208201111561034a57600080fd5b8035906020019184600183028401116401000000008311171561036c57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610658945050505050565b610169600480360360408110156103c357600080fd5b506001600160a01b03813581169160200135166106b4565b610209600480360360608110156103f157600080fd5b6001600160a01b038235169160208101359181019060608101604082013564010000000081111561042157600080fd5b82018360208201111561043357600080fd5b8035906020019184600183028401116401000000008311171561045557600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506106df945050505050565b6001600160e01b03191660009081526020819052604090205460ff1690565b60006104c2338484610734565b50600192915050565b60035490565b60006104de848484610826565b6001600160a01b03841660009081526002602090815260408083203380855292529091205461051e918691610519908663ffffffff61097016565b610734565b5060019392505050565b60006105333361063f565b61057157604051600160e51b62461bcd028152600401808060200182810382526030815260200180610ec36030913960400191505060405180910390fd5b6104c283836109d0565b6105958282604051806020016040528060008152506106df565b5050565b6105b483838360405180602001604052806000815250610658565b505050565b6001600160a01b031660009081526001602052604090205490565b6105dd3361063f565b61061b57604051600160e51b62461bcd028152600401808060200182810382526030815260200180610ec36030913960400191505060405180910390fd5b61062481610ac5565b50565b61063033610b0d565b565b60006104c2338484610826565b600061065260048363ffffffff610b5516565b92915050565b6106638484846104d1565b5061067084848484610bbf565b6106ae57604051600160e51b62461bcd02815260040180806020018281038252602e815260200180610e73602e913960400191505060405180910390fd5b50505050565b6001600160a01b03918216600090815260026020908152604080832093909416825291909152205490565b6106e98383610632565b506106f633848484610bbf565b6105b457604051600160e51b62461bcd02815260040180806020018281038252602e815260200180610e73602e913960400191505060405180910390fd5b6001600160a01b03831661077c57604051600160e51b62461bcd028152600401808060200182810382526023815260200180610f5a6023913960400191505060405180910390fd5b6001600160a01b0382166107c457604051600160e51b62461bcd028152600401808060200182810382526021815260200180610e526021913960400191505060405180910390fd5b6001600160a01b03808416600081815260026020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b6001600160a01b03831661086e57604051600160e51b62461bcd028152600401808060200182810382526024815260200180610f146024913960400191505060405180910390fd5b6001600160a01b0382166108b657604051600160e51b62461bcd028152600401808060200182810382526022815260200180610ea16022913960400191505060405180910390fd5b6001600160a01b0383166000908152600160205260409020546108df908263ffffffff61097016565b6001600160a01b038085166000908152600160205260408082209390935590841681522054610914908263ffffffff610cf916565b6001600160a01b0380841660008181526001602090815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b6000828211156109ca5760408051600160e51b62461bcd02815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b6001600160a01b038216610a2e5760408051600160e51b62461bcd02815260206004820152601e60248201527f4b4950373a206d696e7420746f20746865207a65726f20616464726573730000604482015290519081900360640190fd5b600354610a41908263ffffffff610cf916565b6003556001600160a01b038216600090815260016020526040902054610a6d908263ffffffff610cf916565b6001600160a01b03831660008181526001602090815260408083209490945583518581529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35050565b610ad660048263ffffffff610d5d16565b6040516001600160a01b038216907f6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f690600090a250565b610b1e60048263ffffffff610de116565b6040516001600160a01b038216907fe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb6669290600090a250565b60006001600160a01b038216610b9f57604051600160e51b62461bcd028152600401808060200182810382526022815260200180610f386022913960400191505060405180910390fd5b506001600160a01b03166000908152602091909152604090205460ff1690565b6000610bd3846001600160a01b0316610e4b565b610bdf57506001610cf1565b604051600160e11b634e8c461102815233600482018181526001600160a01b03888116602485015260448401879052608060648501908152865160848601528651600095928a1694639d188c229490938c938b938b939260a4019060208501908083838e5b83811015610c5c578181015183820152602001610c44565b50505050905090810190601f168015610c895780820380516001836020036101000a031916815260200191505b5095505050505050602060405180830381600087803b158015610cab57600080fd5b505af1158015610cbf573d6000803e3d6000fd5b505050506040513d6020811015610cd557600080fd5b50516001600160e01b031916600160e11b634e8c461102149150505b949350505050565b600082820183811015610d565760408051600160e51b62461bcd02815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b610d678282610b55565b15610dbc5760408051600160e51b62461bcd02815260206004820152601f60248201527f526f6c65733a206163636f756e7420616c72656164792068617320726f6c6500604482015290519081900360640190fd5b6001600160a01b0316600090815260209190915260409020805460ff19166001179055565b610deb8282610b55565b610e2957604051600160e51b62461bcd028152600401808060200182810382526021815260200180610ef36021913960400191505060405180910390fd5b6001600160a01b0316600090815260209190915260409020805460ff19169055565b3b15159056fe4b4950373a20617070726f766520746f20746865207a65726f20616464726573734b4950373a207472616e7366657220746f206e6f6e204b495037526563656976657220696d706c656d656e7465724b4950373a207472616e7366657220746f20746865207a65726f20616464726573734d696e746572526f6c653a2063616c6c657220646f6573206e6f74206861766520746865204d696e74657220726f6c65526f6c65733a206163636f756e7420646f6573206e6f74206861766520726f6c654b4950373a207472616e736665722066726f6d20746865207a65726f2061646472657373526f6c65733a206163636f756e7420697320746865207a65726f20616464726573734b4950373a20617070726f76652066726f6d20746865207a65726f2061646472657373a165627a7a7230582033ff0dee0530628e63af40471df15660957ba0ee8d4db383f3d8c4558d382acc0029`

// KIP7MintableFuncSigs maps the 4-byte function signature to its string representation.
var KIP7MintableFuncSigs = map[string]string{
	"983b2d56": "addMinter(address)",
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"aa271e1a": "isMinter(address)",
	"40c10f19": "mint(address,uint256)",
	"98650275": "renounceMinter()",
	"423f6cef": "safeTransfer(address,uint256)",
	"eb795549": "safeTransfer(address,uint256,bytes)",
	"42842e0e": "safeTransferFrom(address,address,uint256)",
	"b88d4fde": "safeTransferFrom(address,address,uint256,bytes)",
	"01ffc9a7": "supportsInterface(bytes4)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// KIP7MintableBin is the compiled bytecode used for deploying new contracts.
var KIP7MintableBin = "0x60806040523480156200001157600080fd5b506200002a6301ffc9a760e01b6200007160201b60201c565b62000042636578737160e01b6200007160201b60201c565b62000053336200014060201b60201c565b6200006b63eab83e2060e01b6200007160201b60201c565b620002b9565b7fffffffff0000000000000000000000000000000000000000000000000000000080821614156200010357604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f4b495031333a20696e76616c696420696e746572666163652069640000000000604482015290519081900360640190fd5b7fffffffff00000000000000000000000000000000000000000000000000000000166000908152602081905260409020805460ff19166001179055565b6200015b8160046200019260201b62000d5d1790919060201c565b6040516001600160a01b038216907f6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f690600090a250565b620001a482826200023660201b60201c565b156200021157604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f526f6c65733a206163636f756e7420616c72656164792068617320726f6c6500604482015290519081900360640190fd5b6001600160a01b0316600090815260209190915260409020805460ff19166001179055565b60006001600160a01b03821662000299576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526022815260200180620012716022913960400191505060405180910390fd5b506001600160a01b03166000908152602091909152604090205460ff1690565b610fa880620002c96000396000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c806370a0823111610097578063aa271e1a11610066578063aa271e1a146102c1578063b88d4fde146102e7578063dd62ed3e146103ad578063eb795549146103db576100f5565b806370a0823114610241578063983b2d5614610267578063986502751461028d578063a9059cbb14610295576100f5565b806323b872dd116100d357806323b872dd1461017b57806340c10f19146101b1578063423f6cef146101dd57806342842e0e1461020b576100f5565b806301ffc9a7146100fa578063095ea7b31461013557806318160ddd14610161575b600080fd5b6101216004803603602081101561011057600080fd5b50356001600160e01b031916610496565b604080519115158252519081900360200190f35b6101216004803603604081101561014b57600080fd5b506001600160a01b0381351690602001356104b5565b6101696104cb565b60408051918252519081900360200190f35b6101216004803603606081101561019157600080fd5b506001600160a01b038135811691602081013590911690604001356104d1565b610121600480360360408110156101c757600080fd5b506001600160a01b038135169060200135610528565b610209600480360360408110156101f357600080fd5b506001600160a01b03813516906020013561057b565b005b6102096004803603606081101561022157600080fd5b506001600160a01b03813581169160208101359091169060400135610599565b6101696004803603602081101561025757600080fd5b50356001600160a01b03166105b9565b6102096004803603602081101561027d57600080fd5b50356001600160a01b03166105d4565b610209610627565b610121600480360360408110156102ab57600080fd5b506001600160a01b038135169060200135610632565b610121600480360360208110156102d757600080fd5b50356001600160a01b031661063f565b610209600480360360808110156102fd57600080fd5b6001600160a01b0382358116926020810135909116916040820135919081019060808101606082013564010000000081111561033857600080fd5b82018360208201111561034a57600080fd5b8035906020019184600183028401116401000000008311171561036c57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610658945050505050565b610169600480360360408110156103c357600080fd5b506001600160a01b03813581169160200135166106b4565b610209600480360360608110156103f157600080fd5b6001600160a01b038235169160208101359181019060608101604082013564010000000081111561042157600080fd5b82018360208201111561043357600080fd5b8035906020019184600183028401116401000000008311171561045557600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506106df945050505050565b6001600160e01b03191660009081526020819052604090205460ff1690565b60006104c2338484610734565b50600192915050565b60035490565b60006104de848484610826565b6001600160a01b03841660009081526002602090815260408083203380855292529091205461051e918691610519908663ffffffff61097016565b610734565b5060019392505050565b60006105333361063f565b61057157604051600160e51b62461bcd028152600401808060200182810382526030815260200180610ec36030913960400191505060405180910390fd5b6104c283836109d0565b6105958282604051806020016040528060008152506106df565b5050565b6105b483838360405180602001604052806000815250610658565b505050565b6001600160a01b031660009081526001602052604090205490565b6105dd3361063f565b61061b57604051600160e51b62461bcd028152600401808060200182810382526030815260200180610ec36030913960400191505060405180910390fd5b61062481610ac5565b50565b61063033610b0d565b565b60006104c2338484610826565b600061065260048363ffffffff610b5516565b92915050565b6106638484846104d1565b5061067084848484610bbf565b6106ae57604051600160e51b62461bcd02815260040180806020018281038252602e815260200180610e73602e913960400191505060405180910390fd5b50505050565b6001600160a01b03918216600090815260026020908152604080832093909416825291909152205490565b6106e98383610632565b506106f633848484610bbf565b6105b457604051600160e51b62461bcd02815260040180806020018281038252602e815260200180610e73602e913960400191505060405180910390fd5b6001600160a01b03831661077c57604051600160e51b62461bcd028152600401808060200182810382526023815260200180610f5a6023913960400191505060405180910390fd5b6001600160a01b0382166107c457604051600160e51b62461bcd028152600401808060200182810382526021815260200180610e526021913960400191505060405180910390fd5b6001600160a01b03808416600081815260026020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b6001600160a01b03831661086e57604051600160e51b62461bcd028152600401808060200182810382526024815260200180610f146024913960400191505060405180910390fd5b6001600160a01b0382166108b657604051600160e51b62461bcd028152600401808060200182810382526022815260200180610ea16022913960400191505060405180910390fd5b6001600160a01b0383166000908152600160205260409020546108df908263ffffffff61097016565b6001600160a01b038085166000908152600160205260408082209390935590841681522054610914908263ffffffff610cf916565b6001600160a01b0380841660008181526001602090815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b6000828211156109ca5760408051600160e51b62461bcd02815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b6001600160a01b038216610a2e5760408051600160e51b62461bcd02815260206004820152601e60248201527f4b4950373a206d696e7420746f20746865207a65726f20616464726573730000604482015290519081900360640190fd5b600354610a41908263ffffffff610cf916565b6003556001600160a01b038216600090815260016020526040902054610a6d908263ffffffff610cf916565b6001600160a01b03831660008181526001602090815260408083209490945583518581529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35050565b610ad660048263ffffffff610d5d16565b6040516001600160a01b038216907f6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f690600090a250565b610b1e60048263ffffffff610de116565b6040516001600160a01b038216907fe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb6669290600090a250565b60006001600160a01b038216610b9f57604051600160e51b62461bcd028152600401808060200182810382526022815260200180610f386022913960400191505060405180910390fd5b506001600160a01b03166000908152602091909152604090205460ff1690565b6000610bd3846001600160a01b0316610e4b565b610bdf57506001610cf1565b604051600160e11b634e8c461102815233600482018181526001600160a01b03888116602485015260448401879052608060648501908152865160848601528651600095928a1694639d188c229490938c938b938b939260a4019060208501908083838e5b83811015610c5c578181015183820152602001610c44565b50505050905090810190601f168015610c895780820380516001836020036101000a031916815260200191505b5095505050505050602060405180830381600087803b158015610cab57600080fd5b505af1158015610cbf573d6000803e3d6000fd5b505050506040513d6020811015610cd557600080fd5b50516001600160e01b031916600160e11b634e8c461102149150505b949350505050565b600082820183811015610d565760408051600160e51b62461bcd02815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b610d678282610b55565b15610dbc5760408051600160e51b62461bcd02815260206004820152601f60248201527f526f6c65733a206163636f756e7420616c72656164792068617320726f6c6500604482015290519081900360640190fd5b6001600160a01b0316600090815260209190915260409020805460ff19166001179055565b610deb8282610b55565b610e2957604051600160e51b62461bcd028152600401808060200182810382526021815260200180610ef36021913960400191505060405180910390fd5b6001600160a01b0316600090815260209190915260409020805460ff19169055565b3b15159056fe4b4950373a20617070726f766520746f20746865207a65726f20616464726573734b4950373a207472616e7366657220746f206e6f6e204b495037526563656976657220696d706c656d656e7465724b4950373a207472616e7366657220746f20746865207a65726f20616464726573734d696e746572526f6c653a2063616c6c657220646f6573206e6f74206861766520746865204d696e74657220726f6c65526f6c65733a206163636f756e7420646f6573206e6f74206861766520726f6c654b4950373a207472616e736665722066726f6d20746865207a65726f2061646472657373526f6c65733a206163636f756e7420697320746865207a65726f20616464726573734b4950373a20617070726f76652066726f6d20746865207a65726f2061646472657373a165627a7a7230582033ff0dee0530628e63af40471df15660957ba0ee8d4db383f3d8c4558d382acc0029526f6c65733a206163636f756e7420697320746865207a65726f2061646472657373"

// DeployKIP7Mintable deploys a new Klaytn contract, binding an instance of KIP7Mintable to it.
func DeployKIP7Mintable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *KIP7Mintable, error) {
	parsed, err := abi.JSON(strings.NewReader(KIP7MintableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(KIP7MintableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KIP7Mintable{KIP7MintableCaller: KIP7MintableCaller{contract: contract}, KIP7MintableTransactor: KIP7MintableTransactor{contract: contract}, KIP7MintableFilterer: KIP7MintableFilterer{contract: contract}}, nil
}

// KIP7Mintable is an auto generated Go binding around a Klaytn contract.
type KIP7Mintable struct {
	KIP7MintableCaller     // Read-only binding to the contract
	KIP7MintableTransactor // Write-only binding to the contract
	KIP7MintableFilterer   // Log filterer for contract events
}

// KIP7MintableCaller is an auto generated read-only Go binding around a Klaytn contract.
type KIP7MintableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KIP7MintableTransactor is an auto generated write-only Go binding around a Klaytn contract.
type KIP7MintableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KIP7MintableFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type KIP7MintableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KIP7MintableSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type KIP7MintableSession struct {
	Contract     *KIP7Mintable     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KIP7MintableCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type KIP7MintableCallerSession struct {
	Contract *KIP7MintableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// KIP7MintableTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type KIP7MintableTransactorSession struct {
	Contract     *KIP7MintableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// KIP7MintableRaw is an auto generated low-level Go binding around a Klaytn contract.
type KIP7MintableRaw struct {
	Contract *KIP7Mintable // Generic contract binding to access the raw methods on
}

// KIP7MintableCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type KIP7MintableCallerRaw struct {
	Contract *KIP7MintableCaller // Generic read-only contract binding to access the raw methods on
}

// KIP7MintableTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type KIP7MintableTransactorRaw struct {
	Contract *KIP7MintableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKIP7Mintable creates a new instance of KIP7Mintable, bound to a specific deployed contract.
func NewKIP7Mintable(address common.Address, backend bind.ContractBackend) (*KIP7Mintable, error) {
	contract, err := bindKIP7Mintable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KIP7Mintable{KIP7MintableCaller: KIP7MintableCaller{contract: contract}, KIP7MintableTransactor: KIP7MintableTransactor{contract: contract}, KIP7MintableFilterer: KIP7MintableFilterer{contract: contract}}, nil
}

// NewKIP7MintableCaller creates a new read-only instance of KIP7Mintable, bound to a specific deployed contract.
func NewKIP7MintableCaller(address common.Address, caller bind.ContractCaller) (*KIP7MintableCaller, error) {
	contract, err := bindKIP7Mintable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KIP7MintableCaller{contract: contract}, nil
}

// NewKIP7MintableTransactor creates a new write-only instance of KIP7Mintable, bound to a specific deployed contract.
func NewKIP7MintableTransactor(address common.Address, transactor bind.ContractTransactor) (*KIP7MintableTransactor, error) {
	contract, err := bindKIP7Mintable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KIP7MintableTransactor{contract: contract}, nil
}

// NewKIP7MintableFilterer creates a new log filterer instance of KIP7Mintable, bound to a specific deployed contract.
func NewKIP7MintableFilterer(address common.Address, filterer bind.ContractFilterer) (*KIP7MintableFilterer, error) {
	contract, err := bindKIP7Mintable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KIP7MintableFilterer{contract: contract}, nil
}

// bindKIP7Mintable binds a generic wrapper to an already deployed contract.
func bindKIP7Mintable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KIP7MintableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KIP7Mintable *KIP7MintableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KIP7Mintable.Contract.KIP7MintableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KIP7Mintable *KIP7MintableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KIP7Mintable.Contract.KIP7MintableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KIP7Mintable *KIP7MintableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KIP7Mintable.Contract.KIP7MintableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KIP7Mintable *KIP7MintableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KIP7Mintable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KIP7Mintable *KIP7MintableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KIP7Mintable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KIP7Mintable *KIP7MintableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KIP7Mintable.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_KIP7Mintable *KIP7MintableCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KIP7Mintable.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_KIP7Mintable *KIP7MintableSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _KIP7Mintable.Contract.Allowance(&_KIP7Mintable.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_KIP7Mintable *KIP7MintableCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _KIP7Mintable.Contract.Allowance(&_KIP7Mintable.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_KIP7Mintable *KIP7MintableCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KIP7Mintable.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_KIP7Mintable *KIP7MintableSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _KIP7Mintable.Contract.BalanceOf(&_KIP7Mintable.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_KIP7Mintable *KIP7MintableCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _KIP7Mintable.Contract.BalanceOf(&_KIP7Mintable.CallOpts, account)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) view returns(bool)
func (_KIP7Mintable *KIP7MintableCaller) IsMinter(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KIP7Mintable.contract.Call(opts, out, "isMinter", account)
	return *ret0, err
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) view returns(bool)
func (_KIP7Mintable *KIP7MintableSession) IsMinter(account common.Address) (bool, error) {
	return _KIP7Mintable.Contract.IsMinter(&_KIP7Mintable.CallOpts, account)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) view returns(bool)
func (_KIP7Mintable *KIP7MintableCallerSession) IsMinter(account common.Address) (bool, error) {
	return _KIP7Mintable.Contract.IsMinter(&_KIP7Mintable.CallOpts, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_KIP7Mintable *KIP7MintableCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KIP7Mintable.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_KIP7Mintable *KIP7MintableSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _KIP7Mintable.Contract.SupportsInterface(&_KIP7Mintable.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_KIP7Mintable *KIP7MintableCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _KIP7Mintable.Contract.SupportsInterface(&_KIP7Mintable.CallOpts, interfaceId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_KIP7Mintable *KIP7MintableCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KIP7Mintable.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_KIP7Mintable *KIP7MintableSession) TotalSupply() (*big.Int, error) {
	return _KIP7Mintable.Contract.TotalSupply(&_KIP7Mintable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_KIP7Mintable *KIP7MintableCallerSession) TotalSupply() (*big.Int, error) {
	return _KIP7Mintable.Contract.TotalSupply(&_KIP7Mintable.CallOpts)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_KIP7Mintable *KIP7MintableTransactor) AddMinter(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _KIP7Mintable.contract.Transact(opts, "addMinter", account)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_KIP7Mintable *KIP7MintableSession) AddMinter(account common.Address) (*types.Transaction, error) {
	return _KIP7Mintable.Contract.AddMinter(&_KIP7Mintable.TransactOpts, account)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_KIP7Mintable *KIP7MintableTransactorSession) AddMinter(account common.Address) (*types.Transaction, error) {
	return _KIP7Mintable.Contract.AddMinter(&_KIP7Mintable.TransactOpts, account)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_KIP7Mintable *KIP7MintableTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _KIP7Mintable.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_KIP7Mintable *KIP7MintableSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _KIP7Mintable.Contract.Approve(&_KIP7Mintable.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_KIP7Mintable *KIP7MintableTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _KIP7Mintable.Contract.Approve(&_KIP7Mintable.TransactOpts, spender, value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns(bool)
func (_KIP7Mintable *KIP7MintableTransactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KIP7Mintable.contract.Transact(opts, "mint", account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns(bool)
func (_KIP7Mintable *KIP7MintableSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KIP7Mintable.Contract.Mint(&_KIP7Mintable.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns(bool)
func (_KIP7Mintable *KIP7MintableTransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KIP7Mintable.Contract.Mint(&_KIP7Mintable.TransactOpts, account, amount)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_KIP7Mintable *KIP7MintableTransactor) RenounceMinter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KIP7Mintable.contract.Transact(opts, "renounceMinter")
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_KIP7Mintable *KIP7MintableSession) RenounceMinter() (*types.Transaction, error) {
	return _KIP7Mintable.Contract.RenounceMinter(&_KIP7Mintable.TransactOpts)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_KIP7Mintable *KIP7MintableTransactorSession) RenounceMinter() (*types.Transaction, error) {
	return _KIP7Mintable.Contract.RenounceMinter(&_KIP7Mintable.TransactOpts)
}

// SafeTransfer is a paid mutator transaction binding the contract method 0x423f6cef.
//
// Solidity: function safeTransfer(address recipient, uint256 amount) returns()
func (_KIP7Mintable *KIP7MintableTransactor) SafeTransfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KIP7Mintable.contract.Transact(opts, "safeTransfer", recipient, amount)
}

// SafeTransfer is a paid mutator transaction binding the contract method 0x423f6cef.
//
// Solidity: function safeTransfer(address recipient, uint256 amount) returns()
func (_KIP7Mintable *KIP7MintableSession) SafeTransfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KIP7Mintable.Contract.SafeTransfer(&_KIP7Mintable.TransactOpts, recipient, amount)
}

// SafeTransfer is a paid mutator transaction binding the contract method 0x423f6cef.
//
// Solidity: function safeTransfer(address recipient, uint256 amount) returns()
func (_KIP7Mintable *KIP7MintableTransactorSession) SafeTransfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KIP7Mintable.Contract.SafeTransfer(&_KIP7Mintable.TransactOpts, recipient, amount)
}

// SafeTransfer0 is a paid mutator transaction binding the contract method 0xeb795549.
//
// Solidity: function safeTransfer(address recipient, uint256 amount, bytes data) returns()
func (_KIP7Mintable *KIP7MintableTransactor) SafeTransfer0(opts *bind.TransactOpts, recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _KIP7Mintable.contract.Transact(opts, "safeTransfer0", recipient, amount, data)
}

// SafeTransfer0 is a paid mutator transaction binding the contract method 0xeb795549.
//
// Solidity: function safeTransfer(address recipient, uint256 amount, bytes data) returns()
func (_KIP7Mintable *KIP7MintableSession) SafeTransfer0(recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _KIP7Mintable.Contract.SafeTransfer0(&_KIP7Mintable.TransactOpts, recipient, amount, data)
}

// SafeTransfer0 is a paid mutator transaction binding the contract method 0xeb795549.
//
// Solidity: function safeTransfer(address recipient, uint256 amount, bytes data) returns()
func (_KIP7Mintable *KIP7MintableTransactorSession) SafeTransfer0(recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _KIP7Mintable.Contract.SafeTransfer0(&_KIP7Mintable.TransactOpts, recipient, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address sender, address recipient, uint256 amount) returns()
func (_KIP7Mintable *KIP7MintableTransactor) SafeTransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KIP7Mintable.contract.Transact(opts, "safeTransferFrom", sender, recipient, amount)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address sender, address recipient, uint256 amount) returns()
func (_KIP7Mintable *KIP7MintableSession) SafeTransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KIP7Mintable.Contract.SafeTransferFrom(&_KIP7Mintable.TransactOpts, sender, recipient, amount)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address sender, address recipient, uint256 amount) returns()
func (_KIP7Mintable *KIP7MintableTransactorSession) SafeTransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KIP7Mintable.Contract.SafeTransferFrom(&_KIP7Mintable.TransactOpts, sender, recipient, amount)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address sender, address recipient, uint256 amount, bytes data) returns()
func (_KIP7Mintable *KIP7MintableTransactor) SafeTransferFrom0(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _KIP7Mintable.contract.Transact(opts, "safeTransferFrom0", sender, recipient, amount, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address sender, address recipient, uint256 amount, bytes data) returns()
func (_KIP7Mintable *KIP7MintableSession) SafeTransferFrom0(sender common.Address, recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _KIP7Mintable.Contract.SafeTransferFrom0(&_KIP7Mintable.TransactOpts, sender, recipient, amount, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address sender, address recipient, uint256 amount, bytes data) returns()
func (_KIP7Mintable *KIP7MintableTransactorSession) SafeTransferFrom0(sender common.Address, recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _KIP7Mintable.Contract.SafeTransferFrom0(&_KIP7Mintable.TransactOpts, sender, recipient, amount, data)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_KIP7Mintable *KIP7MintableTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KIP7Mintable.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_KIP7Mintable *KIP7MintableSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KIP7Mintable.Contract.Transfer(&_KIP7Mintable.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_KIP7Mintable *KIP7MintableTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KIP7Mintable.Contract.Transfer(&_KIP7Mintable.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_KIP7Mintable *KIP7MintableTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KIP7Mintable.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_KIP7Mintable *KIP7MintableSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KIP7Mintable.Contract.TransferFrom(&_KIP7Mintable.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_KIP7Mintable *KIP7MintableTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KIP7Mintable.Contract.TransferFrom(&_KIP7Mintable.TransactOpts, sender, recipient, amount)
}

// KIP7MintableApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the KIP7Mintable contract.
type KIP7MintableApprovalIterator struct {
	Event *KIP7MintableApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KIP7MintableApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KIP7MintableApproval)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KIP7MintableApproval)
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
func (it *KIP7MintableApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KIP7MintableApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KIP7MintableApproval represents a Approval event raised by the KIP7Mintable contract.
type KIP7MintableApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_KIP7Mintable *KIP7MintableFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*KIP7MintableApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _KIP7Mintable.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &KIP7MintableApprovalIterator{contract: _KIP7Mintable.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_KIP7Mintable *KIP7MintableFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *KIP7MintableApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _KIP7Mintable.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KIP7MintableApproval)
				if err := _KIP7Mintable.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_KIP7Mintable *KIP7MintableFilterer) ParseApproval(log types.Log) (*KIP7MintableApproval, error) {
	event := new(KIP7MintableApproval)
	if err := _KIP7Mintable.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KIP7MintableMinterAddedIterator is returned from FilterMinterAdded and is used to iterate over the raw logs and unpacked data for MinterAdded events raised by the KIP7Mintable contract.
type KIP7MintableMinterAddedIterator struct {
	Event *KIP7MintableMinterAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KIP7MintableMinterAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KIP7MintableMinterAdded)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KIP7MintableMinterAdded)
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
func (it *KIP7MintableMinterAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KIP7MintableMinterAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KIP7MintableMinterAdded represents a MinterAdded event raised by the KIP7Mintable contract.
type KIP7MintableMinterAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinterAdded is a free log retrieval operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed account)
func (_KIP7Mintable *KIP7MintableFilterer) FilterMinterAdded(opts *bind.FilterOpts, account []common.Address) (*KIP7MintableMinterAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _KIP7Mintable.contract.FilterLogs(opts, "MinterAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &KIP7MintableMinterAddedIterator{contract: _KIP7Mintable.contract, event: "MinterAdded", logs: logs, sub: sub}, nil
}

// WatchMinterAdded is a free log subscription operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed account)
func (_KIP7Mintable *KIP7MintableFilterer) WatchMinterAdded(opts *bind.WatchOpts, sink chan<- *KIP7MintableMinterAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _KIP7Mintable.contract.WatchLogs(opts, "MinterAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KIP7MintableMinterAdded)
				if err := _KIP7Mintable.contract.UnpackLog(event, "MinterAdded", log); err != nil {
					return err
				}
				event.Raw = log

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

// ParseMinterAdded is a log parse operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed account)
func (_KIP7Mintable *KIP7MintableFilterer) ParseMinterAdded(log types.Log) (*KIP7MintableMinterAdded, error) {
	event := new(KIP7MintableMinterAdded)
	if err := _KIP7Mintable.contract.UnpackLog(event, "MinterAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KIP7MintableMinterRemovedIterator is returned from FilterMinterRemoved and is used to iterate over the raw logs and unpacked data for MinterRemoved events raised by the KIP7Mintable contract.
type KIP7MintableMinterRemovedIterator struct {
	Event *KIP7MintableMinterRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KIP7MintableMinterRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KIP7MintableMinterRemoved)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KIP7MintableMinterRemoved)
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
func (it *KIP7MintableMinterRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KIP7MintableMinterRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KIP7MintableMinterRemoved represents a MinterRemoved event raised by the KIP7Mintable contract.
type KIP7MintableMinterRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinterRemoved is a free log retrieval operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed account)
func (_KIP7Mintable *KIP7MintableFilterer) FilterMinterRemoved(opts *bind.FilterOpts, account []common.Address) (*KIP7MintableMinterRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _KIP7Mintable.contract.FilterLogs(opts, "MinterRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &KIP7MintableMinterRemovedIterator{contract: _KIP7Mintable.contract, event: "MinterRemoved", logs: logs, sub: sub}, nil
}

// WatchMinterRemoved is a free log subscription operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed account)
func (_KIP7Mintable *KIP7MintableFilterer) WatchMinterRemoved(opts *bind.WatchOpts, sink chan<- *KIP7MintableMinterRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _KIP7Mintable.contract.WatchLogs(opts, "MinterRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KIP7MintableMinterRemoved)
				if err := _KIP7Mintable.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
					return err
				}
				event.Raw = log

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

// ParseMinterRemoved is a log parse operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed account)
func (_KIP7Mintable *KIP7MintableFilterer) ParseMinterRemoved(log types.Log) (*KIP7MintableMinterRemoved, error) {
	event := new(KIP7MintableMinterRemoved)
	if err := _KIP7Mintable.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KIP7MintableTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the KIP7Mintable contract.
type KIP7MintableTransferIterator struct {
	Event *KIP7MintableTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KIP7MintableTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KIP7MintableTransfer)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KIP7MintableTransfer)
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
func (it *KIP7MintableTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KIP7MintableTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KIP7MintableTransfer represents a Transfer event raised by the KIP7Mintable contract.
type KIP7MintableTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_KIP7Mintable *KIP7MintableFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*KIP7MintableTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _KIP7Mintable.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &KIP7MintableTransferIterator{contract: _KIP7Mintable.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_KIP7Mintable *KIP7MintableFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *KIP7MintableTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _KIP7Mintable.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KIP7MintableTransfer)
				if err := _KIP7Mintable.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_KIP7Mintable *KIP7MintableFilterer) ParseTransfer(log types.Log) (*KIP7MintableTransfer, error) {
	event := new(KIP7MintableTransfer)
	if err := _KIP7Mintable.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KIP7ServiceChainABI is the input ABI used to generate the binding from.
const KIP7ServiceChainABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_feeLimit\",\"type\":\"uint256\"},{\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"requestValueTransfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"safeTransfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_bridge\",\"type\":\"address\"}],\"name\":\"setBridge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_bridge\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// KIP7ServiceChainBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const KIP7ServiceChainBinRuntime = ``

// KIP7ServiceChainFuncSigs maps the 4-byte function signature to its string representation.
var KIP7ServiceChainFuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"e78cea92": "bridge()",
	"8f32d59b": "isOwner()",
	"8da5cb5b": "owner()",
	"715018a6": "renounceOwnership()",
	"3dc3c9e1": "requestValueTransfer(uint256,address,uint256,bytes)",
	"423f6cef": "safeTransfer(address,uint256)",
	"eb795549": "safeTransfer(address,uint256,bytes)",
	"42842e0e": "safeTransferFrom(address,address,uint256)",
	"b88d4fde": "safeTransferFrom(address,address,uint256,bytes)",
	"8dd14802": "setBridge(address)",
	"01ffc9a7": "supportsInterface(bytes4)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
	"f2fde38b": "transferOwnership(address)",
}

// KIP7ServiceChain is an auto generated Go binding around a Klaytn contract.
type KIP7ServiceChain struct {
	KIP7ServiceChainCaller     // Read-only binding to the contract
	KIP7ServiceChainTransactor // Write-only binding to the contract
	KIP7ServiceChainFilterer   // Log filterer for contract events
}

// KIP7ServiceChainCaller is an auto generated read-only Go binding around a Klaytn contract.
type KIP7ServiceChainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KIP7ServiceChainTransactor is an auto generated write-only Go binding around a Klaytn contract.
type KIP7ServiceChainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KIP7ServiceChainFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type KIP7ServiceChainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KIP7ServiceChainSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type KIP7ServiceChainSession struct {
	Contract     *KIP7ServiceChain // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KIP7ServiceChainCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type KIP7ServiceChainCallerSession struct {
	Contract *KIP7ServiceChainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// KIP7ServiceChainTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type KIP7ServiceChainTransactorSession struct {
	Contract     *KIP7ServiceChainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// KIP7ServiceChainRaw is an auto generated low-level Go binding around a Klaytn contract.
type KIP7ServiceChainRaw struct {
	Contract *KIP7ServiceChain // Generic contract binding to access the raw methods on
}

// KIP7ServiceChainCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type KIP7ServiceChainCallerRaw struct {
	Contract *KIP7ServiceChainCaller // Generic read-only contract binding to access the raw methods on
}

// KIP7ServiceChainTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type KIP7ServiceChainTransactorRaw struct {
	Contract *KIP7ServiceChainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKIP7ServiceChain creates a new instance of KIP7ServiceChain, bound to a specific deployed contract.
func NewKIP7ServiceChain(address common.Address, backend bind.ContractBackend) (*KIP7ServiceChain, error) {
	contract, err := bindKIP7ServiceChain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KIP7ServiceChain{KIP7ServiceChainCaller: KIP7ServiceChainCaller{contract: contract}, KIP7ServiceChainTransactor: KIP7ServiceChainTransactor{contract: contract}, KIP7ServiceChainFilterer: KIP7ServiceChainFilterer{contract: contract}}, nil
}

// NewKIP7ServiceChainCaller creates a new read-only instance of KIP7ServiceChain, bound to a specific deployed contract.
func NewKIP7ServiceChainCaller(address common.Address, caller bind.ContractCaller) (*KIP7ServiceChainCaller, error) {
	contract, err := bindKIP7ServiceChain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KIP7ServiceChainCaller{contract: contract}, nil
}

// NewKIP7ServiceChainTransactor creates a new write-only instance of KIP7ServiceChain, bound to a specific deployed contract.
func NewKIP7ServiceChainTransactor(address common.Address, transactor bind.ContractTransactor) (*KIP7ServiceChainTransactor, error) {
	contract, err := bindKIP7ServiceChain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KIP7ServiceChainTransactor{contract: contract}, nil
}

// NewKIP7ServiceChainFilterer creates a new log filterer instance of KIP7ServiceChain, bound to a specific deployed contract.
func NewKIP7ServiceChainFilterer(address common.Address, filterer bind.ContractFilterer) (*KIP7ServiceChainFilterer, error) {
	contract, err := bindKIP7ServiceChain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KIP7ServiceChainFilterer{contract: contract}, nil
}

// bindKIP7ServiceChain binds a generic wrapper to an already deployed contract.
func bindKIP7ServiceChain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KIP7ServiceChainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KIP7ServiceChain *KIP7ServiceChainRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KIP7ServiceChain.Contract.KIP7ServiceChainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KIP7ServiceChain *KIP7ServiceChainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KIP7ServiceChain.Contract.KIP7ServiceChainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KIP7ServiceChain *KIP7ServiceChainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KIP7ServiceChain.Contract.KIP7ServiceChainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KIP7ServiceChain *KIP7ServiceChainCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KIP7ServiceChain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KIP7ServiceChain *KIP7ServiceChainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KIP7ServiceChain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KIP7ServiceChain *KIP7ServiceChainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KIP7ServiceChain.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_KIP7ServiceChain *KIP7ServiceChainCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KIP7ServiceChain.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_KIP7ServiceChain *KIP7ServiceChainSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _KIP7ServiceChain.Contract.Allowance(&_KIP7ServiceChain.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_KIP7ServiceChain *KIP7ServiceChainCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _KIP7ServiceChain.Contract.Allowance(&_KIP7ServiceChain.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_KIP7ServiceChain *KIP7ServiceChainCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KIP7ServiceChain.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_KIP7ServiceChain *KIP7ServiceChainSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _KIP7ServiceChain.Contract.BalanceOf(&_KIP7ServiceChain.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_KIP7ServiceChain *KIP7ServiceChainCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _KIP7ServiceChain.Contract.BalanceOf(&_KIP7ServiceChain.CallOpts, account)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_KIP7ServiceChain *KIP7ServiceChainCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KIP7ServiceChain.contract.Call(opts, out, "bridge")
	return *ret0, err
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_KIP7ServiceChain *KIP7ServiceChainSession) Bridge() (common.Address, error) {
	return _KIP7ServiceChain.Contract.Bridge(&_KIP7ServiceChain.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_KIP7ServiceChain *KIP7ServiceChainCallerSession) Bridge() (common.Address, error) {
	return _KIP7ServiceChain.Contract.Bridge(&_KIP7ServiceChain.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_KIP7ServiceChain *KIP7ServiceChainCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KIP7ServiceChain.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_KIP7ServiceChain *KIP7ServiceChainSession) IsOwner() (bool, error) {
	return _KIP7ServiceChain.Contract.IsOwner(&_KIP7ServiceChain.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_KIP7ServiceChain *KIP7ServiceChainCallerSession) IsOwner() (bool, error) {
	return _KIP7ServiceChain.Contract.IsOwner(&_KIP7ServiceChain.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_KIP7ServiceChain *KIP7ServiceChainCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KIP7ServiceChain.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_KIP7ServiceChain *KIP7ServiceChainSession) Owner() (common.Address, error) {
	return _KIP7ServiceChain.Contract.Owner(&_KIP7ServiceChain.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_KIP7ServiceChain *KIP7ServiceChainCallerSession) Owner() (common.Address, error) {
	return _KIP7ServiceChain.Contract.Owner(&_KIP7ServiceChain.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_KIP7ServiceChain *KIP7ServiceChainCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KIP7ServiceChain.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_KIP7ServiceChain *KIP7ServiceChainSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _KIP7ServiceChain.Contract.SupportsInterface(&_KIP7ServiceChain.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_KIP7ServiceChain *KIP7ServiceChainCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _KIP7ServiceChain.Contract.SupportsInterface(&_KIP7ServiceChain.CallOpts, interfaceId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_KIP7ServiceChain *KIP7ServiceChainCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KIP7ServiceChain.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_KIP7ServiceChain *KIP7ServiceChainSession) TotalSupply() (*big.Int, error) {
	return _KIP7ServiceChain.Contract.TotalSupply(&_KIP7ServiceChain.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_KIP7ServiceChain *KIP7ServiceChainCallerSession) TotalSupply() (*big.Int, error) {
	return _KIP7ServiceChain.Contract.TotalSupply(&_KIP7ServiceChain.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_KIP7ServiceChain *KIP7ServiceChainTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _KIP7ServiceChain.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_KIP7ServiceChain *KIP7ServiceChainSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _KIP7ServiceChain.Contract.Approve(&_KIP7ServiceChain.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_KIP7ServiceChain *KIP7ServiceChainTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _KIP7ServiceChain.Contract.Approve(&_KIP7ServiceChain.TransactOpts, spender, value)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_KIP7ServiceChain *KIP7ServiceChainTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KIP7ServiceChain.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_KIP7ServiceChain *KIP7ServiceChainSession) RenounceOwnership() (*types.Transaction, error) {
	return _KIP7ServiceChain.Contract.RenounceOwnership(&_KIP7ServiceChain.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_KIP7ServiceChain *KIP7ServiceChainTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _KIP7ServiceChain.Contract.RenounceOwnership(&_KIP7ServiceChain.TransactOpts)
}

// RequestValueTransfer is a paid mutator transaction binding the contract method 0x3dc3c9e1.
//
// Solidity: function requestValueTransfer(uint256 _amount, address _to, uint256 _feeLimit, bytes _extraData) returns()
func (_KIP7ServiceChain *KIP7ServiceChainTransactor) RequestValueTransfer(opts *bind.TransactOpts, _amount *big.Int, _to common.Address, _feeLimit *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _KIP7ServiceChain.contract.Transact(opts, "requestValueTransfer", _amount, _to, _feeLimit, _extraData)
}

// RequestValueTransfer is a paid mutator transaction binding the contract method 0x3dc3c9e1.
//
// Solidity: function requestValueTransfer(uint256 _amount, address _to, uint256 _feeLimit, bytes _extraData) returns()
func (_KIP7ServiceChain *KIP7ServiceChainSession) RequestValueTransfer(_amount *big.Int, _to common.Address, _feeLimit *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _KIP7ServiceChain.Contract.RequestValueTransfer(&_KIP7ServiceChain.TransactOpts, _amount, _to, _feeLimit, _extraData)
}

// RequestValueTransfer is a paid mutator transaction binding the contract method 0x3dc3c9e1.
//
// Solidity: function requestValueTransfer(uint256 _amount, address _to, uint256 _feeLimit, bytes _extraData) returns()
func (_KIP7ServiceChain *KIP7ServiceChainTransactorSession) RequestValueTransfer(_amount *big.Int, _to common.Address, _feeLimit *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _KIP7ServiceChain.Contract.RequestValueTransfer(&_KIP7ServiceChain.TransactOpts, _amount, _to, _feeLimit, _extraData)
}

// SafeTransfer is a paid mutator transaction binding the contract method 0x423f6cef.
//
// Solidity: function safeTransfer(address recipient, uint256 amount) returns()
func (_KIP7ServiceChain *KIP7ServiceChainTransactor) SafeTransfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KIP7ServiceChain.contract.Transact(opts, "safeTransfer", recipient, amount)
}

// SafeTransfer is a paid mutator transaction binding the contract method 0x423f6cef.
//
// Solidity: function safeTransfer(address recipient, uint256 amount) returns()
func (_KIP7ServiceChain *KIP7ServiceChainSession) SafeTransfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KIP7ServiceChain.Contract.SafeTransfer(&_KIP7ServiceChain.TransactOpts, recipient, amount)
}

// SafeTransfer is a paid mutator transaction binding the contract method 0x423f6cef.
//
// Solidity: function safeTransfer(address recipient, uint256 amount) returns()
func (_KIP7ServiceChain *KIP7ServiceChainTransactorSession) SafeTransfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KIP7ServiceChain.Contract.SafeTransfer(&_KIP7ServiceChain.TransactOpts, recipient, amount)
}

// SafeTransfer0 is a paid mutator transaction binding the contract method 0xeb795549.
//
// Solidity: function safeTransfer(address recipient, uint256 amount, bytes data) returns()
func (_KIP7ServiceChain *KIP7ServiceChainTransactor) SafeTransfer0(opts *bind.TransactOpts, recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _KIP7ServiceChain.contract.Transact(opts, "safeTransfer0", recipient, amount, data)
}

// SafeTransfer0 is a paid mutator transaction binding the contract method 0xeb795549.
//
// Solidity: function safeTransfer(address recipient, uint256 amount, bytes data) returns()
func (_KIP7ServiceChain *KIP7ServiceChainSession) SafeTransfer0(recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _KIP7ServiceChain.Contract.SafeTransfer0(&_KIP7ServiceChain.TransactOpts, recipient, amount, data)
}

// SafeTransfer0 is a paid mutator transaction binding the contract method 0xeb795549.
//
// Solidity: function safeTransfer(address recipient, uint256 amount, bytes data) returns()
func (_KIP7ServiceChain *KIP7ServiceChainTransactorSession) SafeTransfer0(recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _KIP7ServiceChain.Contract.SafeTransfer0(&_KIP7ServiceChain.TransactOpts, recipient, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address sender, address recipient, uint256 amount) returns()
func (_KIP7ServiceChain *KIP7ServiceChainTransactor) SafeTransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KIP7ServiceChain.contract.Transact(opts, "safeTransferFrom", sender, recipient, amount)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address sender, address recipient, uint256 amount) returns()
func (_KIP7ServiceChain *KIP7ServiceChainSession) SafeTransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KIP7ServiceChain.Contract.SafeTransferFrom(&_KIP7ServiceChain.TransactOpts, sender, recipient, amount)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address sender, address recipient, uint256 amount) returns()
func (_KIP7ServiceChain *KIP7ServiceChainTransactorSession) SafeTransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KIP7ServiceChain.Contract.SafeTransferFrom(&_KIP7ServiceChain.TransactOpts, sender, recipient, amount)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address sender, address recipient, uint256 amount, bytes data) returns()
func (_KIP7ServiceChain *KIP7ServiceChainTransactor) SafeTransferFrom0(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _KIP7ServiceChain.contract.Transact(opts, "safeTransferFrom0", sender, recipient, amount, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address sender, address recipient, uint256 amount, bytes data) returns()
func (_KIP7ServiceChain *KIP7ServiceChainSession) SafeTransferFrom0(sender common.Address, recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _KIP7ServiceChain.Contract.SafeTransferFrom0(&_KIP7ServiceChain.TransactOpts, sender, recipient, amount, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address sender, address recipient, uint256 amount, bytes data) returns()
func (_KIP7ServiceChain *KIP7ServiceChainTransactorSession) SafeTransferFrom0(sender common.Address, recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _KIP7ServiceChain.Contract.SafeTransferFrom0(&_KIP7ServiceChain.TransactOpts, sender, recipient, amount, data)
}

// SetBridge is a paid mutator transaction binding the contract method 0x8dd14802.
//
// Solidity: function setBridge(address _bridge) returns()
func (_KIP7ServiceChain *KIP7ServiceChainTransactor) SetBridge(opts *bind.TransactOpts, _bridge common.Address) (*types.Transaction, error) {
	return _KIP7ServiceChain.contract.Transact(opts, "setBridge", _bridge)
}

// SetBridge is a paid mutator transaction binding the contract method 0x8dd14802.
//
// Solidity: function setBridge(address _bridge) returns()
func (_KIP7ServiceChain *KIP7ServiceChainSession) SetBridge(_bridge common.Address) (*types.Transaction, error) {
	return _KIP7ServiceChain.Contract.SetBridge(&_KIP7ServiceChain.TransactOpts, _bridge)
}

// SetBridge is a paid mutator transaction binding the contract method 0x8dd14802.
//
// Solidity: function setBridge(address _bridge) returns()
func (_KIP7ServiceChain *KIP7ServiceChainTransactorSession) SetBridge(_bridge common.Address) (*types.Transaction, error) {
	return _KIP7ServiceChain.Contract.SetBridge(&_KIP7ServiceChain.TransactOpts, _bridge)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_KIP7ServiceChain *KIP7ServiceChainTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KIP7ServiceChain.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_KIP7ServiceChain *KIP7ServiceChainSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KIP7ServiceChain.Contract.Transfer(&_KIP7ServiceChain.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_KIP7ServiceChain *KIP7ServiceChainTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KIP7ServiceChain.Contract.Transfer(&_KIP7ServiceChain.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_KIP7ServiceChain *KIP7ServiceChainTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KIP7ServiceChain.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_KIP7ServiceChain *KIP7ServiceChainSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KIP7ServiceChain.Contract.TransferFrom(&_KIP7ServiceChain.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_KIP7ServiceChain *KIP7ServiceChainTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KIP7ServiceChain.Contract.TransferFrom(&_KIP7ServiceChain.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KIP7ServiceChain *KIP7ServiceChainTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _KIP7ServiceChain.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KIP7ServiceChain *KIP7ServiceChainSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _KIP7ServiceChain.Contract.TransferOwnership(&_KIP7ServiceChain.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KIP7ServiceChain *KIP7ServiceChainTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _KIP7ServiceChain.Contract.TransferOwnership(&_KIP7ServiceChain.TransactOpts, newOwner)
}

// KIP7ServiceChainApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the KIP7ServiceChain contract.
type KIP7ServiceChainApprovalIterator struct {
	Event *KIP7ServiceChainApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KIP7ServiceChainApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KIP7ServiceChainApproval)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KIP7ServiceChainApproval)
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
func (it *KIP7ServiceChainApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KIP7ServiceChainApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KIP7ServiceChainApproval represents a Approval event raised by the KIP7ServiceChain contract.
type KIP7ServiceChainApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_KIP7ServiceChain *KIP7ServiceChainFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*KIP7ServiceChainApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _KIP7ServiceChain.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &KIP7ServiceChainApprovalIterator{contract: _KIP7ServiceChain.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_KIP7ServiceChain *KIP7ServiceChainFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *KIP7ServiceChainApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _KIP7ServiceChain.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KIP7ServiceChainApproval)
				if err := _KIP7ServiceChain.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_KIP7ServiceChain *KIP7ServiceChainFilterer) ParseApproval(log types.Log) (*KIP7ServiceChainApproval, error) {
	event := new(KIP7ServiceChainApproval)
	if err := _KIP7ServiceChain.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KIP7ServiceChainOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the KIP7ServiceChain contract.
type KIP7ServiceChainOwnershipTransferredIterator struct {
	Event *KIP7ServiceChainOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KIP7ServiceChainOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KIP7ServiceChainOwnershipTransferred)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KIP7ServiceChainOwnershipTransferred)
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
func (it *KIP7ServiceChainOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KIP7ServiceChainOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KIP7ServiceChainOwnershipTransferred represents a OwnershipTransferred event raised by the KIP7ServiceChain contract.
type KIP7ServiceChainOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_KIP7ServiceChain *KIP7ServiceChainFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*KIP7ServiceChainOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _KIP7ServiceChain.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &KIP7ServiceChainOwnershipTransferredIterator{contract: _KIP7ServiceChain.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_KIP7ServiceChain *KIP7ServiceChainFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *KIP7ServiceChainOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _KIP7ServiceChain.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KIP7ServiceChainOwnershipTransferred)
				if err := _KIP7ServiceChain.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

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
func (_KIP7ServiceChain *KIP7ServiceChainFilterer) ParseOwnershipTransferred(log types.Log) (*KIP7ServiceChainOwnershipTransferred, error) {
	event := new(KIP7ServiceChainOwnershipTransferred)
	if err := _KIP7ServiceChain.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KIP7ServiceChainTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the KIP7ServiceChain contract.
type KIP7ServiceChainTransferIterator struct {
	Event *KIP7ServiceChainTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KIP7ServiceChainTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KIP7ServiceChainTransfer)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KIP7ServiceChainTransfer)
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
func (it *KIP7ServiceChainTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KIP7ServiceChainTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KIP7ServiceChainTransfer represents a Transfer event raised by the KIP7ServiceChain contract.
type KIP7ServiceChainTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_KIP7ServiceChain *KIP7ServiceChainFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*KIP7ServiceChainTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _KIP7ServiceChain.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &KIP7ServiceChainTransferIterator{contract: _KIP7ServiceChain.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_KIP7ServiceChain *KIP7ServiceChainFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *KIP7ServiceChainTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _KIP7ServiceChain.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KIP7ServiceChainTransfer)
				if err := _KIP7ServiceChain.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_KIP7ServiceChain *KIP7ServiceChainFilterer) ParseTransfer(log types.Log) (*KIP7ServiceChainTransfer, error) {
	event := new(KIP7ServiceChainTransfer)
	if err := _KIP7ServiceChain.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MinterRoleABI is the input ABI used to generate the binding from.
const MinterRoleABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addMinter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceMinter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isMinter\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MinterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MinterRemoved\",\"type\":\"event\"}]"

// MinterRoleBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const MinterRoleBinRuntime = ``

// MinterRoleFuncSigs maps the 4-byte function signature to its string representation.
var MinterRoleFuncSigs = map[string]string{
	"983b2d56": "addMinter(address)",
	"aa271e1a": "isMinter(address)",
	"98650275": "renounceMinter()",
}

// MinterRole is an auto generated Go binding around a Klaytn contract.
type MinterRole struct {
	MinterRoleCaller     // Read-only binding to the contract
	MinterRoleTransactor // Write-only binding to the contract
	MinterRoleFilterer   // Log filterer for contract events
}

// MinterRoleCaller is an auto generated read-only Go binding around a Klaytn contract.
type MinterRoleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinterRoleTransactor is an auto generated write-only Go binding around a Klaytn contract.
type MinterRoleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinterRoleFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type MinterRoleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinterRoleSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type MinterRoleSession struct {
	Contract     *MinterRole       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MinterRoleCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type MinterRoleCallerSession struct {
	Contract *MinterRoleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MinterRoleTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type MinterRoleTransactorSession struct {
	Contract     *MinterRoleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MinterRoleRaw is an auto generated low-level Go binding around a Klaytn contract.
type MinterRoleRaw struct {
	Contract *MinterRole // Generic contract binding to access the raw methods on
}

// MinterRoleCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type MinterRoleCallerRaw struct {
	Contract *MinterRoleCaller // Generic read-only contract binding to access the raw methods on
}

// MinterRoleTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type MinterRoleTransactorRaw struct {
	Contract *MinterRoleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMinterRole creates a new instance of MinterRole, bound to a specific deployed contract.
func NewMinterRole(address common.Address, backend bind.ContractBackend) (*MinterRole, error) {
	contract, err := bindMinterRole(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MinterRole{MinterRoleCaller: MinterRoleCaller{contract: contract}, MinterRoleTransactor: MinterRoleTransactor{contract: contract}, MinterRoleFilterer: MinterRoleFilterer{contract: contract}}, nil
}

// NewMinterRoleCaller creates a new read-only instance of MinterRole, bound to a specific deployed contract.
func NewMinterRoleCaller(address common.Address, caller bind.ContractCaller) (*MinterRoleCaller, error) {
	contract, err := bindMinterRole(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MinterRoleCaller{contract: contract}, nil
}

// NewMinterRoleTransactor creates a new write-only instance of MinterRole, bound to a specific deployed contract.
func NewMinterRoleTransactor(address common.Address, transactor bind.ContractTransactor) (*MinterRoleTransactor, error) {
	contract, err := bindMinterRole(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MinterRoleTransactor{contract: contract}, nil
}

// NewMinterRoleFilterer creates a new log filterer instance of MinterRole, bound to a specific deployed contract.
func NewMinterRoleFilterer(address common.Address, filterer bind.ContractFilterer) (*MinterRoleFilterer, error) {
	contract, err := bindMinterRole(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MinterRoleFilterer{contract: contract}, nil
}

// bindMinterRole binds a generic wrapper to an already deployed contract.
func bindMinterRole(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MinterRoleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MinterRole *MinterRoleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MinterRole.Contract.MinterRoleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MinterRole *MinterRoleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MinterRole.Contract.MinterRoleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MinterRole *MinterRoleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MinterRole.Contract.MinterRoleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MinterRole *MinterRoleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MinterRole.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MinterRole *MinterRoleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MinterRole.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MinterRole *MinterRoleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MinterRole.Contract.contract.Transact(opts, method, params...)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) view returns(bool)
func (_MinterRole *MinterRoleCaller) IsMinter(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MinterRole.contract.Call(opts, out, "isMinter", account)
	return *ret0, err
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) view returns(bool)
func (_MinterRole *MinterRoleSession) IsMinter(account common.Address) (bool, error) {
	return _MinterRole.Contract.IsMinter(&_MinterRole.CallOpts, account)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) view returns(bool)
func (_MinterRole *MinterRoleCallerSession) IsMinter(account common.Address) (bool, error) {
	return _MinterRole.Contract.IsMinter(&_MinterRole.CallOpts, account)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_MinterRole *MinterRoleTransactor) AddMinter(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _MinterRole.contract.Transact(opts, "addMinter", account)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_MinterRole *MinterRoleSession) AddMinter(account common.Address) (*types.Transaction, error) {
	return _MinterRole.Contract.AddMinter(&_MinterRole.TransactOpts, account)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_MinterRole *MinterRoleTransactorSession) AddMinter(account common.Address) (*types.Transaction, error) {
	return _MinterRole.Contract.AddMinter(&_MinterRole.TransactOpts, account)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_MinterRole *MinterRoleTransactor) RenounceMinter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MinterRole.contract.Transact(opts, "renounceMinter")
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_MinterRole *MinterRoleSession) RenounceMinter() (*types.Transaction, error) {
	return _MinterRole.Contract.RenounceMinter(&_MinterRole.TransactOpts)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_MinterRole *MinterRoleTransactorSession) RenounceMinter() (*types.Transaction, error) {
	return _MinterRole.Contract.RenounceMinter(&_MinterRole.TransactOpts)
}

// MinterRoleMinterAddedIterator is returned from FilterMinterAdded and is used to iterate over the raw logs and unpacked data for MinterAdded events raised by the MinterRole contract.
type MinterRoleMinterAddedIterator struct {
	Event *MinterRoleMinterAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MinterRoleMinterAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MinterRoleMinterAdded)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MinterRoleMinterAdded)
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
func (it *MinterRoleMinterAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MinterRoleMinterAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MinterRoleMinterAdded represents a MinterAdded event raised by the MinterRole contract.
type MinterRoleMinterAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinterAdded is a free log retrieval operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed account)
func (_MinterRole *MinterRoleFilterer) FilterMinterAdded(opts *bind.FilterOpts, account []common.Address) (*MinterRoleMinterAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _MinterRole.contract.FilterLogs(opts, "MinterAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &MinterRoleMinterAddedIterator{contract: _MinterRole.contract, event: "MinterAdded", logs: logs, sub: sub}, nil
}

// WatchMinterAdded is a free log subscription operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed account)
func (_MinterRole *MinterRoleFilterer) WatchMinterAdded(opts *bind.WatchOpts, sink chan<- *MinterRoleMinterAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _MinterRole.contract.WatchLogs(opts, "MinterAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MinterRoleMinterAdded)
				if err := _MinterRole.contract.UnpackLog(event, "MinterAdded", log); err != nil {
					return err
				}
				event.Raw = log

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

// ParseMinterAdded is a log parse operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed account)
func (_MinterRole *MinterRoleFilterer) ParseMinterAdded(log types.Log) (*MinterRoleMinterAdded, error) {
	event := new(MinterRoleMinterAdded)
	if err := _MinterRole.contract.UnpackLog(event, "MinterAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MinterRoleMinterRemovedIterator is returned from FilterMinterRemoved and is used to iterate over the raw logs and unpacked data for MinterRemoved events raised by the MinterRole contract.
type MinterRoleMinterRemovedIterator struct {
	Event *MinterRoleMinterRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MinterRoleMinterRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MinterRoleMinterRemoved)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MinterRoleMinterRemoved)
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
func (it *MinterRoleMinterRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MinterRoleMinterRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MinterRoleMinterRemoved represents a MinterRemoved event raised by the MinterRole contract.
type MinterRoleMinterRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinterRemoved is a free log retrieval operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed account)
func (_MinterRole *MinterRoleFilterer) FilterMinterRemoved(opts *bind.FilterOpts, account []common.Address) (*MinterRoleMinterRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _MinterRole.contract.FilterLogs(opts, "MinterRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &MinterRoleMinterRemovedIterator{contract: _MinterRole.contract, event: "MinterRemoved", logs: logs, sub: sub}, nil
}

// WatchMinterRemoved is a free log subscription operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed account)
func (_MinterRole *MinterRoleFilterer) WatchMinterRemoved(opts *bind.WatchOpts, sink chan<- *MinterRoleMinterRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _MinterRole.contract.WatchLogs(opts, "MinterRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MinterRoleMinterRemoved)
				if err := _MinterRole.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
					return err
				}
				event.Raw = log

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

// ParseMinterRemoved is a log parse operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed account)
func (_MinterRole *MinterRoleFilterer) ParseMinterRemoved(log types.Log) (*MinterRoleMinterRemoved, error) {
	event := new(MinterRoleMinterRemoved)
	if err := _MinterRole.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// OwnableBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const OwnableBinRuntime = ``

// OwnableFuncSigs maps the 4-byte function signature to its string representation.
var OwnableFuncSigs = map[string]string{
	"8f32d59b": "isOwner()",
	"8da5cb5b": "owner()",
	"715018a6": "renounceOwnership()",
	"f2fde38b": "transferOwnership(address)",
}

// Ownable is an auto generated Go binding around a Klaytn contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around a Klaytn contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around a Klaytn contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around a Klaytn contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Ownable *OwnableCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Ownable *OwnableSession) IsOwner() (bool, error) {
	return _Ownable.Contract.IsOwner(&_Ownable.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Ownable *OwnableCallerSession) IsOwner() (bool, error) {
	return _Ownable.Contract.IsOwner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OwnableOwnershipTransferred)
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
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

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
func (_Ownable *OwnableFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableOwnershipTransferred, error) {
	event := new(OwnableOwnershipTransferred)
	if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RolesABI is the input ABI used to generate the binding from.
const RolesABI = "[]"

// RolesBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const RolesBinRuntime = `73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a7230582026429ed555f83236ac076363d644804ff5b0baa4e098f1d4ef783a49049f7df70029`

// RolesBin is the compiled bytecode used for deploying new contracts.
var RolesBin = "0x604c6023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a7230582026429ed555f83236ac076363d644804ff5b0baa4e098f1d4ef783a49049f7df70029"

// DeployRoles deploys a new Klaytn contract, binding an instance of Roles to it.
func DeployRoles(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Roles, error) {
	parsed, err := abi.JSON(strings.NewReader(RolesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RolesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Roles{RolesCaller: RolesCaller{contract: contract}, RolesTransactor: RolesTransactor{contract: contract}, RolesFilterer: RolesFilterer{contract: contract}}, nil
}

// Roles is an auto generated Go binding around a Klaytn contract.
type Roles struct {
	RolesCaller     // Read-only binding to the contract
	RolesTransactor // Write-only binding to the contract
	RolesFilterer   // Log filterer for contract events
}

// RolesCaller is an auto generated read-only Go binding around a Klaytn contract.
type RolesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RolesTransactor is an auto generated write-only Go binding around a Klaytn contract.
type RolesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RolesFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type RolesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RolesSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type RolesSession struct {
	Contract     *Roles            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RolesCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type RolesCallerSession struct {
	Contract *RolesCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RolesTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type RolesTransactorSession struct {
	Contract     *RolesTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RolesRaw is an auto generated low-level Go binding around a Klaytn contract.
type RolesRaw struct {
	Contract *Roles // Generic contract binding to access the raw methods on
}

// RolesCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type RolesCallerRaw struct {
	Contract *RolesCaller // Generic read-only contract binding to access the raw methods on
}

// RolesTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type RolesTransactorRaw struct {
	Contract *RolesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRoles creates a new instance of Roles, bound to a specific deployed contract.
func NewRoles(address common.Address, backend bind.ContractBackend) (*Roles, error) {
	contract, err := bindRoles(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Roles{RolesCaller: RolesCaller{contract: contract}, RolesTransactor: RolesTransactor{contract: contract}, RolesFilterer: RolesFilterer{contract: contract}}, nil
}

// NewRolesCaller creates a new read-only instance of Roles, bound to a specific deployed contract.
func NewRolesCaller(address common.Address, caller bind.ContractCaller) (*RolesCaller, error) {
	contract, err := bindRoles(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RolesCaller{contract: contract}, nil
}

// NewRolesTransactor creates a new write-only instance of Roles, bound to a specific deployed contract.
func NewRolesTransactor(address common.Address, transactor bind.ContractTransactor) (*RolesTransactor, error) {
	contract, err := bindRoles(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RolesTransactor{contract: contract}, nil
}

// NewRolesFilterer creates a new log filterer instance of Roles, bound to a specific deployed contract.
func NewRolesFilterer(address common.Address, filterer bind.ContractFilterer) (*RolesFilterer, error) {
	contract, err := bindRoles(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RolesFilterer{contract: contract}, nil
}

// bindRoles binds a generic wrapper to an already deployed contract.
func bindRoles(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RolesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Roles *RolesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Roles.Contract.RolesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Roles *RolesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Roles.Contract.RolesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Roles *RolesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Roles.Contract.RolesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Roles *RolesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Roles.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Roles *RolesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Roles.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Roles *RolesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Roles.Contract.contract.Transact(opts, method, params...)
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const SafeMathBinRuntime = `73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a723058201aa360e6241ae1bc829140fb38ae5ca1672181b2a5c529c4ca3a045ed2896feb0029`

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x604c6023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a723058201aa360e6241ae1bc829140fb38ae5ca1672181b2a5c529c4ca3a045ed2896feb0029"

// DeploySafeMath deploys a new Klaytn contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around a Klaytn contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around a Klaytn contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around a Klaytn contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around a Klaytn contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// ServiceChainKIP7TokenABI is the input ABI used to generate the binding from.
const ServiceChainKIP7TokenABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"DECIMALS\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"INITIAL_SUPPLY\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_feeLimit\",\"type\":\"uint256\"},{\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"requestValueTransfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"safeTransfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_bridge\",\"type\":\"address\"}],\"name\":\"setBridge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addMinter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceMinter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NAME\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isMinter\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"SYMBOL\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_bridge\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MinterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MinterRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// ServiceChainKIP7TokenBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const ServiceChainKIP7TokenBinRuntime = `608060405234801561001057600080fd5b50600436106101c45760003560e01c80638da5cb5b116100f9578063aa271e1a11610097578063e78cea9211610071578063e78cea921461067c578063eb79554914610684578063f2fde38b1461073f578063f76f8d7814610765576101c4565b8063aa271e1a14610562578063b88d4fde14610588578063dd62ed3e1461064e576101c4565b8063983b2d56116100d3578063983b2d561461048b57806398650275146104b1578063a3f4df7e146104b9578063a9059cbb14610536576101c4565b80638da5cb5b146104395780638dd148021461045d5780638f32d59b14610483576101c4565b806340c10f191161016657806342966c681161014057806342966c68146103c257806370a08231146103df578063715018a61461040557806379cc67901461040d576101c4565b806340c10f1914610334578063423f6cef1461036057806342842e0e1461038c576101c4565b806323b872dd116101a257806323b872dd1461024a5780632e0f2625146102805780632ff2e9dc1461029e5780633dc3c9e1146102a6576101c4565b806301ffc9a7146101c9578063095ea7b31461020457806318160ddd14610230575b600080fd5b6101f0600480360360208110156101df57600080fd5b50356001600160e01b03191661076d565b604080519115158252519081900360200190f35b6101f06004803603604081101561021a57600080fd5b506001600160a01b03813516906020013561078c565b6102386107a2565b60408051918252519081900360200190f35b6101f06004803603606081101561026057600080fd5b506001600160a01b038135811691602081013590911690604001356107a8565b6102886107ff565b6040805160ff9092168252519081900360200190f35b610238610804565b610332600480360360808110156102bc57600080fd5b8135916001600160a01b0360208201351691604082013591908101906080810160608201356401000000008111156102f357600080fd5b82018360208201111561030557600080fd5b8035906020019184600183028401116401000000008311171561032757600080fd5b509092509050610814565b005b6101f06004803603604081101561034a57600080fd5b506001600160a01b0381351690602001356108fa565b6103326004803603604081101561037657600080fd5b506001600160a01b03813516906020013561094d565b610332600480360360608110156103a257600080fd5b506001600160a01b0381358116916020810135909116906040013561096b565b610332600480360360208110156103d857600080fd5b503561098b565b610238600480360360208110156103f557600080fd5b50356001600160a01b0316610998565b6103326109b3565b6103326004803603604081101561042357600080fd5b506001600160a01b038135169060200135610a59565b610441610a63565b604080516001600160a01b039092168252519081900360200190f35b6103326004803603602081101561047357600080fd5b50356001600160a01b0316610a72565b6101f0610af0565b610332600480360360208110156104a157600080fd5b50356001600160a01b0316610b01565b610332610b51565b6104c1610b5c565b6040805160208082528351818301528351919283929083019185019080838360005b838110156104fb5781810151838201526020016104e3565b50505050905090810190601f1680156105285780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101f06004803603604081101561054c57600080fd5b506001600160a01b038135169060200135610b95565b6101f06004803603602081101561057857600080fd5b50356001600160a01b0316610ba2565b6103326004803603608081101561059e57600080fd5b6001600160a01b038235811692602081013590911691604082013591908101906080810160608201356401000000008111156105d957600080fd5b8201836020820111156105eb57600080fd5b8035906020019184600183028401116401000000008311171561060d57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610bbb945050505050565b6102386004803603604081101561066457600080fd5b506001600160a01b0381358116916020013516610c17565b610441610c42565b6103326004803603606081101561069a57600080fd5b6001600160a01b03823516916020810135918101906060810160408201356401000000008111156106ca57600080fd5b8201836020820111156106dc57600080fd5b803590602001918460018302840111640100000000831117156106fe57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610c51945050505050565b6103326004803603602081101561075557600080fd5b50356001600160a01b0316610ca6565b6104c1610d0b565b6001600160e01b03191660009081526020819052604090205460ff1690565b6000610799338484610d2d565b50600192915050565b60035490565b60006107b5848484610e1f565b6001600160a01b0384166000908152600260209081526040808320338085529252909120546107f59186916107f0908663ffffffff610f6916565b610d2d565b5060019392505050565b601281565b6b033b2e3c9fd0803ce800000081565b600654610839906001600160a01b0316610834878663ffffffff610fc916565b610b95565b50600654604051600160e41b630c80e3a502815233600482018181526001600160a01b038881166024850152604484018a90526064840188905260a06084850190815260a4850187905294169363c80e3a509389928b928a928a928a92909160c401848480828437600081840152601f19601f820116905080830192505050975050505050505050600060405180830381600087803b1580156108db57600080fd5b505af11580156108ef573d6000803e3d6000fd5b505050505050505050565b600061090533610ba2565b61094357604051600160e51b62461bcd0281526004018080602001828103825260308152602001806116bf6030913960400191505060405180910390fd5b610799838361102d565b610967828260405180602001604052806000815250610c51565b5050565b61098683838360405180602001604052806000815250610bbb565b505050565b6109953382611122565b50565b6001600160a01b031660009081526001602052604090205490565b6109bb610af0565b610a0f5760408051600160e51b62461bcd02815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b6005546040516000916001600160a01b0316907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600580546001600160a01b0319169055565b6109678282611216565b6005546001600160a01b031690565b610a7a610af0565b610ace5760408051600160e51b62461bcd02815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b600680546001600160a01b0319166001600160a01b0392909216919091179055565b6005546001600160a01b0316331490565b610b0a33610ba2565b610b4857604051600160e51b62461bcd0281526004018080602001828103825260308152602001806116bf6030913960400191505060405180910390fd5b6109958161125b565b610b5a336112a3565b565b6040518060400160405280601181526020017f53657276696365436861696e546f6b656e00000000000000000000000000000081525081565b6000610799338484610e1f565b6000610bb560048363ffffffff6112eb16565b92915050565b610bc68484846107a8565b50610bd384848484611355565b610c1157604051600160e51b62461bcd02815260040180806020018281038252602e815260200180611649602e913960400191505060405180910390fd5b50505050565b6001600160a01b03918216600090815260026020908152604080832093909416825291909152205490565b6006546001600160a01b031681565b610c5b8383610b95565b50610c6833848484611355565b61098657604051600160e51b62461bcd02815260040180806020018281038252602e815260200180611649602e913960400191505060405180910390fd5b610cae610af0565b610d025760408051600160e51b62461bcd02815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b6109958161148f565b604051806040016040528060038152602001600160ea1b6214d0d50281525081565b6001600160a01b038316610d7557604051600160e51b62461bcd0281526004018080602001828103825260238152602001806117566023913960400191505060405180910390fd5b6001600160a01b038216610dbd57604051600160e51b62461bcd0281526004018080602001828103825260218152602001806116286021913960400191505060405180910390fd5b6001600160a01b03808416600081815260026020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b6001600160a01b038316610e6757604051600160e51b62461bcd0281526004018080602001828103825260248152602001806117106024913960400191505060405180910390fd5b6001600160a01b038216610eaf57604051600160e51b62461bcd02815260040180806020018281038252602281526020018061169d6022913960400191505060405180910390fd5b6001600160a01b038316600090815260016020526040902054610ed8908263ffffffff610f6916565b6001600160a01b038085166000908152600160205260408082209390935590841681522054610f0d908263ffffffff610fc916565b6001600160a01b0380841660008181526001602090815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b600082821115610fc35760408051600160e51b62461bcd02815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b6000828201838110156110265760408051600160e51b62461bcd02815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b6001600160a01b03821661108b5760408051600160e51b62461bcd02815260206004820152601e60248201527f4b4950373a206d696e7420746f20746865207a65726f20616464726573730000604482015290519081900360640190fd5b60035461109e908263ffffffff610fc916565b6003556001600160a01b0382166000908152600160205260409020546110ca908263ffffffff610fc916565b6001600160a01b03831660008181526001602090815260408083209490945583518581529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35050565b6001600160a01b0382166111805760408051600160e51b62461bcd02815260206004820181905260248201527f4b4950373a206275726e2066726f6d20746865207a65726f2061646472657373604482015290519081900360640190fd5b600354611193908263ffffffff610f6916565b6003556001600160a01b0382166000908152600160205260409020546111bf908263ffffffff610f6916565b6001600160a01b0383166000818152600160209081526040808320949094558351858152935191937fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929081900390910190a35050565b6112208282611122565b6001600160a01b0382166000908152600260209081526040808320338085529252909120546109679184916107f0908563ffffffff610f6916565b61126c60048263ffffffff61153316565b6040516001600160a01b038216907f6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f690600090a250565b6112b460048263ffffffff6115b716565b6040516001600160a01b038216907fe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb6669290600090a250565b60006001600160a01b03821661133557604051600160e51b62461bcd0281526004018080602001828103825260228152602001806117346022913960400191505060405180910390fd5b506001600160a01b03166000908152602091909152604090205460ff1690565b6000611369846001600160a01b0316611621565b61137557506001611487565b604051600160e11b634e8c461102815233600482018181526001600160a01b03888116602485015260448401879052608060648501908152865160848601528651600095928a1694639d188c229490938c938b938b939260a4019060208501908083838e5b838110156113f25781810151838201526020016113da565b50505050905090810190601f16801561141f5780820380516001836020036101000a031916815260200191505b5095505050505050602060405180830381600087803b15801561144157600080fd5b505af1158015611455573d6000803e3d6000fd5b505050506040513d602081101561146b57600080fd5b50516001600160e01b031916600160e11b634e8c461102149150505b949350505050565b6001600160a01b0381166114d757604051600160e51b62461bcd0281526004018080602001828103825260268152602001806116776026913960400191505060405180910390fd5b6005546040516001600160a01b038084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3600580546001600160a01b0319166001600160a01b0392909216919091179055565b61153d82826112eb565b156115925760408051600160e51b62461bcd02815260206004820152601f60248201527f526f6c65733a206163636f756e7420616c72656164792068617320726f6c6500604482015290519081900360640190fd5b6001600160a01b0316600090815260209190915260409020805460ff19166001179055565b6115c182826112eb565b6115ff57604051600160e51b62461bcd0281526004018080602001828103825260218152602001806116ef6021913960400191505060405180910390fd5b6001600160a01b0316600090815260209190915260409020805460ff19169055565b3b15159056fe4b4950373a20617070726f766520746f20746865207a65726f20616464726573734b4950373a207472616e7366657220746f206e6f6e204b495037526563656976657220696d706c656d656e7465724f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573734b4950373a207472616e7366657220746f20746865207a65726f20616464726573734d696e746572526f6c653a2063616c6c657220646f6573206e6f74206861766520746865204d696e74657220726f6c65526f6c65733a206163636f756e7420646f6573206e6f74206861766520726f6c654b4950373a207472616e736665722066726f6d20746865207a65726f2061646472657373526f6c65733a206163636f756e7420697320746865207a65726f20616464726573734b4950373a20617070726f76652066726f6d20746865207a65726f2061646472657373a165627a7a723058205251cb568d61461ffe0f6834f44e8b81dcd8505e9aff90eb24ccbd73d82e0b430029`

// ServiceChainKIP7TokenFuncSigs maps the 4-byte function signature to its string representation.
var ServiceChainKIP7TokenFuncSigs = map[string]string{
	"2e0f2625": "DECIMALS()",
	"2ff2e9dc": "INITIAL_SUPPLY()",
	"a3f4df7e": "NAME()",
	"f76f8d78": "SYMBOL()",
	"983b2d56": "addMinter(address)",
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"e78cea92": "bridge()",
	"42966c68": "burn(uint256)",
	"79cc6790": "burnFrom(address,uint256)",
	"aa271e1a": "isMinter(address)",
	"8f32d59b": "isOwner()",
	"40c10f19": "mint(address,uint256)",
	"8da5cb5b": "owner()",
	"98650275": "renounceMinter()",
	"715018a6": "renounceOwnership()",
	"3dc3c9e1": "requestValueTransfer(uint256,address,uint256,bytes)",
	"423f6cef": "safeTransfer(address,uint256)",
	"eb795549": "safeTransfer(address,uint256,bytes)",
	"42842e0e": "safeTransferFrom(address,address,uint256)",
	"b88d4fde": "safeTransferFrom(address,address,uint256,bytes)",
	"8dd14802": "setBridge(address)",
	"01ffc9a7": "supportsInterface(bytes4)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
	"f2fde38b": "transferOwnership(address)",
}

// ServiceChainKIP7TokenBin is the compiled bytecode used for deploying new contracts.
var ServiceChainKIP7TokenBin = "0x60806040523480156200001157600080fd5b5060405160208062001d9d833981018060405260208110156200003357600080fd5b505180620000687f01ffc9a700000000000000000000000000000000000000000000000000000000620001de602090811b901c565b62000080636578737160e01b620001de60201b60201c565b6200009133620002ad60201b60201c565b620000a963eab83e2060e01b620001de60201b60201c565b620000c1633b5a0bf860e01b620001de60201b60201c565b600580546001600160a01b0319163317908190556040516001600160a01b0391909116906000907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a36200012c816001600160a01b0316620002ff60201b620016211760201c565b6200019857604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f627269646765206973206e6f74206120636f6e74726163740000000000000000604482015290519081900360640190fd5b600680546001600160a01b0319166001600160a01b0392909216919091179055620001d7336b033b2e3c9fd0803ce800000062000305602090811b901c565b50620005c7565b7fffffffff0000000000000000000000000000000000000000000000000000000080821614156200027057604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f4b495031333a20696e76616c696420696e746572666163652069640000000000604482015290519081900360640190fd5b7fffffffff00000000000000000000000000000000000000000000000000000000166000908152602081905260409020805460ff19166001179055565b620002c88160046200042460201b620015331790919060201c565b6040516001600160a01b038216907f6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f690600090a250565b3b151590565b6001600160a01b0382166200037b57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f4b4950373a206d696e7420746f20746865207a65726f20616464726573730000604482015290519081900360640190fd5b6200039781600354620004c860201b62000fc91790919060201c565b6003556001600160a01b038216600090815260016020908152604090912054620003cc91839062000fc9620004c8821b17901c565b6001600160a01b03831660008181526001602090815260408083209490945583518581529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35050565b6200043682826200054460201b60201c565b15620004a357604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f526f6c65733a206163636f756e7420616c72656164792068617320726f6c6500604482015290519081900360640190fd5b6001600160a01b0316600090815260209190915260409020805460ff19166001179055565b6000828201838110156200053d57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b60006001600160a01b038216620005a7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602281526020018062001d7b6022913960400191505060405180910390fd5b506001600160a01b03166000908152602091909152604090205460ff1690565b6117a480620005d76000396000f3fe608060405234801561001057600080fd5b50600436106101c45760003560e01c80638da5cb5b116100f9578063aa271e1a11610097578063e78cea9211610071578063e78cea921461067c578063eb79554914610684578063f2fde38b1461073f578063f76f8d7814610765576101c4565b8063aa271e1a14610562578063b88d4fde14610588578063dd62ed3e1461064e576101c4565b8063983b2d56116100d3578063983b2d561461048b57806398650275146104b1578063a3f4df7e146104b9578063a9059cbb14610536576101c4565b80638da5cb5b146104395780638dd148021461045d5780638f32d59b14610483576101c4565b806340c10f191161016657806342966c681161014057806342966c68146103c257806370a08231146103df578063715018a61461040557806379cc67901461040d576101c4565b806340c10f1914610334578063423f6cef1461036057806342842e0e1461038c576101c4565b806323b872dd116101a257806323b872dd1461024a5780632e0f2625146102805780632ff2e9dc1461029e5780633dc3c9e1146102a6576101c4565b806301ffc9a7146101c9578063095ea7b31461020457806318160ddd14610230575b600080fd5b6101f0600480360360208110156101df57600080fd5b50356001600160e01b03191661076d565b604080519115158252519081900360200190f35b6101f06004803603604081101561021a57600080fd5b506001600160a01b03813516906020013561078c565b6102386107a2565b60408051918252519081900360200190f35b6101f06004803603606081101561026057600080fd5b506001600160a01b038135811691602081013590911690604001356107a8565b6102886107ff565b6040805160ff9092168252519081900360200190f35b610238610804565b610332600480360360808110156102bc57600080fd5b8135916001600160a01b0360208201351691604082013591908101906080810160608201356401000000008111156102f357600080fd5b82018360208201111561030557600080fd5b8035906020019184600183028401116401000000008311171561032757600080fd5b509092509050610814565b005b6101f06004803603604081101561034a57600080fd5b506001600160a01b0381351690602001356108fa565b6103326004803603604081101561037657600080fd5b506001600160a01b03813516906020013561094d565b610332600480360360608110156103a257600080fd5b506001600160a01b0381358116916020810135909116906040013561096b565b610332600480360360208110156103d857600080fd5b503561098b565b610238600480360360208110156103f557600080fd5b50356001600160a01b0316610998565b6103326109b3565b6103326004803603604081101561042357600080fd5b506001600160a01b038135169060200135610a59565b610441610a63565b604080516001600160a01b039092168252519081900360200190f35b6103326004803603602081101561047357600080fd5b50356001600160a01b0316610a72565b6101f0610af0565b610332600480360360208110156104a157600080fd5b50356001600160a01b0316610b01565b610332610b51565b6104c1610b5c565b6040805160208082528351818301528351919283929083019185019080838360005b838110156104fb5781810151838201526020016104e3565b50505050905090810190601f1680156105285780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101f06004803603604081101561054c57600080fd5b506001600160a01b038135169060200135610b95565b6101f06004803603602081101561057857600080fd5b50356001600160a01b0316610ba2565b6103326004803603608081101561059e57600080fd5b6001600160a01b038235811692602081013590911691604082013591908101906080810160608201356401000000008111156105d957600080fd5b8201836020820111156105eb57600080fd5b8035906020019184600183028401116401000000008311171561060d57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610bbb945050505050565b6102386004803603604081101561066457600080fd5b506001600160a01b0381358116916020013516610c17565b610441610c42565b6103326004803603606081101561069a57600080fd5b6001600160a01b03823516916020810135918101906060810160408201356401000000008111156106ca57600080fd5b8201836020820111156106dc57600080fd5b803590602001918460018302840111640100000000831117156106fe57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610c51945050505050565b6103326004803603602081101561075557600080fd5b50356001600160a01b0316610ca6565b6104c1610d0b565b6001600160e01b03191660009081526020819052604090205460ff1690565b6000610799338484610d2d565b50600192915050565b60035490565b60006107b5848484610e1f565b6001600160a01b0384166000908152600260209081526040808320338085529252909120546107f59186916107f0908663ffffffff610f6916565b610d2d565b5060019392505050565b601281565b6b033b2e3c9fd0803ce800000081565b600654610839906001600160a01b0316610834878663ffffffff610fc916565b610b95565b50600654604051600160e41b630c80e3a502815233600482018181526001600160a01b038881166024850152604484018a90526064840188905260a06084850190815260a4850187905294169363c80e3a509389928b928a928a928a92909160c401848480828437600081840152601f19601f820116905080830192505050975050505050505050600060405180830381600087803b1580156108db57600080fd5b505af11580156108ef573d6000803e3d6000fd5b505050505050505050565b600061090533610ba2565b61094357604051600160e51b62461bcd0281526004018080602001828103825260308152602001806116bf6030913960400191505060405180910390fd5b610799838361102d565b610967828260405180602001604052806000815250610c51565b5050565b61098683838360405180602001604052806000815250610bbb565b505050565b6109953382611122565b50565b6001600160a01b031660009081526001602052604090205490565b6109bb610af0565b610a0f5760408051600160e51b62461bcd02815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b6005546040516000916001600160a01b0316907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600580546001600160a01b0319169055565b6109678282611216565b6005546001600160a01b031690565b610a7a610af0565b610ace5760408051600160e51b62461bcd02815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b600680546001600160a01b0319166001600160a01b0392909216919091179055565b6005546001600160a01b0316331490565b610b0a33610ba2565b610b4857604051600160e51b62461bcd0281526004018080602001828103825260308152602001806116bf6030913960400191505060405180910390fd5b6109958161125b565b610b5a336112a3565b565b6040518060400160405280601181526020017f53657276696365436861696e546f6b656e00000000000000000000000000000081525081565b6000610799338484610e1f565b6000610bb560048363ffffffff6112eb16565b92915050565b610bc68484846107a8565b50610bd384848484611355565b610c1157604051600160e51b62461bcd02815260040180806020018281038252602e815260200180611649602e913960400191505060405180910390fd5b50505050565b6001600160a01b03918216600090815260026020908152604080832093909416825291909152205490565b6006546001600160a01b031681565b610c5b8383610b95565b50610c6833848484611355565b61098657604051600160e51b62461bcd02815260040180806020018281038252602e815260200180611649602e913960400191505060405180910390fd5b610cae610af0565b610d025760408051600160e51b62461bcd02815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b6109958161148f565b604051806040016040528060038152602001600160ea1b6214d0d50281525081565b6001600160a01b038316610d7557604051600160e51b62461bcd0281526004018080602001828103825260238152602001806117566023913960400191505060405180910390fd5b6001600160a01b038216610dbd57604051600160e51b62461bcd0281526004018080602001828103825260218152602001806116286021913960400191505060405180910390fd5b6001600160a01b03808416600081815260026020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b6001600160a01b038316610e6757604051600160e51b62461bcd0281526004018080602001828103825260248152602001806117106024913960400191505060405180910390fd5b6001600160a01b038216610eaf57604051600160e51b62461bcd02815260040180806020018281038252602281526020018061169d6022913960400191505060405180910390fd5b6001600160a01b038316600090815260016020526040902054610ed8908263ffffffff610f6916565b6001600160a01b038085166000908152600160205260408082209390935590841681522054610f0d908263ffffffff610fc916565b6001600160a01b0380841660008181526001602090815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b600082821115610fc35760408051600160e51b62461bcd02815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b6000828201838110156110265760408051600160e51b62461bcd02815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b6001600160a01b03821661108b5760408051600160e51b62461bcd02815260206004820152601e60248201527f4b4950373a206d696e7420746f20746865207a65726f20616464726573730000604482015290519081900360640190fd5b60035461109e908263ffffffff610fc916565b6003556001600160a01b0382166000908152600160205260409020546110ca908263ffffffff610fc916565b6001600160a01b03831660008181526001602090815260408083209490945583518581529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35050565b6001600160a01b0382166111805760408051600160e51b62461bcd02815260206004820181905260248201527f4b4950373a206275726e2066726f6d20746865207a65726f2061646472657373604482015290519081900360640190fd5b600354611193908263ffffffff610f6916565b6003556001600160a01b0382166000908152600160205260409020546111bf908263ffffffff610f6916565b6001600160a01b0383166000818152600160209081526040808320949094558351858152935191937fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929081900390910190a35050565b6112208282611122565b6001600160a01b0382166000908152600260209081526040808320338085529252909120546109679184916107f0908563ffffffff610f6916565b61126c60048263ffffffff61153316565b6040516001600160a01b038216907f6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f690600090a250565b6112b460048263ffffffff6115b716565b6040516001600160a01b038216907fe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb6669290600090a250565b60006001600160a01b03821661133557604051600160e51b62461bcd0281526004018080602001828103825260228152602001806117346022913960400191505060405180910390fd5b506001600160a01b03166000908152602091909152604090205460ff1690565b6000611369846001600160a01b0316611621565b61137557506001611487565b604051600160e11b634e8c461102815233600482018181526001600160a01b03888116602485015260448401879052608060648501908152865160848601528651600095928a1694639d188c229490938c938b938b939260a4019060208501908083838e5b838110156113f25781810151838201526020016113da565b50505050905090810190601f16801561141f5780820380516001836020036101000a031916815260200191505b5095505050505050602060405180830381600087803b15801561144157600080fd5b505af1158015611455573d6000803e3d6000fd5b505050506040513d602081101561146b57600080fd5b50516001600160e01b031916600160e11b634e8c461102149150505b949350505050565b6001600160a01b0381166114d757604051600160e51b62461bcd0281526004018080602001828103825260268152602001806116776026913960400191505060405180910390fd5b6005546040516001600160a01b038084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3600580546001600160a01b0319166001600160a01b0392909216919091179055565b61153d82826112eb565b156115925760408051600160e51b62461bcd02815260206004820152601f60248201527f526f6c65733a206163636f756e7420616c72656164792068617320726f6c6500604482015290519081900360640190fd5b6001600160a01b0316600090815260209190915260409020805460ff19166001179055565b6115c182826112eb565b6115ff57604051600160e51b62461bcd0281526004018080602001828103825260218152602001806116ef6021913960400191505060405180910390fd5b6001600160a01b0316600090815260209190915260409020805460ff19169055565b3b15159056fe4b4950373a20617070726f766520746f20746865207a65726f20616464726573734b4950373a207472616e7366657220746f206e6f6e204b495037526563656976657220696d706c656d656e7465724f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573734b4950373a207472616e7366657220746f20746865207a65726f20616464726573734d696e746572526f6c653a2063616c6c657220646f6573206e6f74206861766520746865204d696e74657220726f6c65526f6c65733a206163636f756e7420646f6573206e6f74206861766520726f6c654b4950373a207472616e736665722066726f6d20746865207a65726f2061646472657373526f6c65733a206163636f756e7420697320746865207a65726f20616464726573734b4950373a20617070726f76652066726f6d20746865207a65726f2061646472657373a165627a7a723058205251cb568d61461ffe0f6834f44e8b81dcd8505e9aff90eb24ccbd73d82e0b430029526f6c65733a206163636f756e7420697320746865207a65726f2061646472657373"

// DeployServiceChainKIP7Token deploys a new Klaytn contract, binding an instance of ServiceChainKIP7Token to it.
func DeployServiceChainKIP7Token(auth *bind.TransactOpts, backend bind.ContractBackend, _bridge common.Address) (common.Address, *types.Transaction, *ServiceChainKIP7Token, error) {
	parsed, err := abi.JSON(strings.NewReader(ServiceChainKIP7TokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ServiceChainKIP7TokenBin), backend, _bridge)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ServiceChainKIP7Token{ServiceChainKIP7TokenCaller: ServiceChainKIP7TokenCaller{contract: contract}, ServiceChainKIP7TokenTransactor: ServiceChainKIP7TokenTransactor{contract: contract}, ServiceChainKIP7TokenFilterer: ServiceChainKIP7TokenFilterer{contract: contract}}, nil
}

// ServiceChainKIP7Token is an auto generated Go binding around a Klaytn contract.
type ServiceChainKIP7Token struct {
	ServiceChainKIP7TokenCaller     // Read-only binding to the contract
	ServiceChainKIP7TokenTransactor // Write-only binding to the contract
	ServiceChainKIP7TokenFilterer   // Log filterer for contract events
}

// ServiceChainKIP7TokenCaller is an auto generated read-only Go binding around a Klaytn contract.
type ServiceChainKIP7TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ServiceChainKIP7TokenTransactor is an auto generated write-only Go binding around a Klaytn contract.
type ServiceChainKIP7TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ServiceChainKIP7TokenFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type ServiceChainKIP7TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ServiceChainKIP7TokenSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type ServiceChainKIP7TokenSession struct {
	Contract     *ServiceChainKIP7Token // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ServiceChainKIP7TokenCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type ServiceChainKIP7TokenCallerSession struct {
	Contract *ServiceChainKIP7TokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// ServiceChainKIP7TokenTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type ServiceChainKIP7TokenTransactorSession struct {
	Contract     *ServiceChainKIP7TokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// ServiceChainKIP7TokenRaw is an auto generated low-level Go binding around a Klaytn contract.
type ServiceChainKIP7TokenRaw struct {
	Contract *ServiceChainKIP7Token // Generic contract binding to access the raw methods on
}

// ServiceChainKIP7TokenCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type ServiceChainKIP7TokenCallerRaw struct {
	Contract *ServiceChainKIP7TokenCaller // Generic read-only contract binding to access the raw methods on
}

// ServiceChainKIP7TokenTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type ServiceChainKIP7TokenTransactorRaw struct {
	Contract *ServiceChainKIP7TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewServiceChainKIP7Token creates a new instance of ServiceChainKIP7Token, bound to a specific deployed contract.
func NewServiceChainKIP7Token(address common.Address, backend bind.ContractBackend) (*ServiceChainKIP7Token, error) {
	contract, err := bindServiceChainKIP7Token(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ServiceChainKIP7Token{ServiceChainKIP7TokenCaller: ServiceChainKIP7TokenCaller{contract: contract}, ServiceChainKIP7TokenTransactor: ServiceChainKIP7TokenTransactor{contract: contract}, ServiceChainKIP7TokenFilterer: ServiceChainKIP7TokenFilterer{contract: contract}}, nil
}

// NewServiceChainKIP7TokenCaller creates a new read-only instance of ServiceChainKIP7Token, bound to a specific deployed contract.
func NewServiceChainKIP7TokenCaller(address common.Address, caller bind.ContractCaller) (*ServiceChainKIP7TokenCaller, error) {
	contract, err := bindServiceChainKIP7Token(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ServiceChainKIP7TokenCaller{contract: contract}, nil
}

// NewServiceChainKIP7TokenTransactor creates a new write-only instance of ServiceChainKIP7Token, bound to a specific deployed contract.
func NewServiceChainKIP7TokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ServiceChainKIP7TokenTransactor, error) {
	contract, err := bindServiceChainKIP7Token(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ServiceChainKIP7TokenTransactor{contract: contract}, nil
}

// NewServiceChainKIP7TokenFilterer creates a new log filterer instance of ServiceChainKIP7Token, bound to a specific deployed contract.
func NewServiceChainKIP7TokenFilterer(address common.Address, filterer bind.ContractFilterer) (*ServiceChainKIP7TokenFilterer, error) {
	contract, err := bindServiceChainKIP7Token(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ServiceChainKIP7TokenFilterer{contract: contract}, nil
}

// bindServiceChainKIP7Token binds a generic wrapper to an already deployed contract.
func bindServiceChainKIP7Token(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ServiceChainKIP7TokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ServiceChainKIP7Token.Contract.ServiceChainKIP7TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ServiceChainKIP7Token.Contract.ServiceChainKIP7TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ServiceChainKIP7Token.Contract.ServiceChainKIP7TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ServiceChainKIP7Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ServiceChainKIP7Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ServiceChainKIP7Token.Contract.contract.Transact(opts, method, params...)
}

// DECIMALS is a free data retrieval call binding the contract method 0x2e0f2625.
//
// Solidity: function DECIMALS() view returns(uint8)
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenCaller) DECIMALS(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _ServiceChainKIP7Token.contract.Call(opts, out, "DECIMALS")
	return *ret0, err
}

// DECIMALS is a free data retrieval call binding the contract method 0x2e0f2625.
//
// Solidity: function DECIMALS() view returns(uint8)
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenSession) DECIMALS() (uint8, error) {
	return _ServiceChainKIP7Token.Contract.DECIMALS(&_ServiceChainKIP7Token.CallOpts)
}

// DECIMALS is a free data retrieval call binding the contract method 0x2e0f2625.
//
// Solidity: function DECIMALS() view returns(uint8)
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenCallerSession) DECIMALS() (uint8, error) {
	return _ServiceChainKIP7Token.Contract.DECIMALS(&_ServiceChainKIP7Token.CallOpts)
}

// INITIALSUPPLY is a free data retrieval call binding the contract method 0x2ff2e9dc.
//
// Solidity: function INITIAL_SUPPLY() view returns(uint256)
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenCaller) INITIALSUPPLY(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ServiceChainKIP7Token.contract.Call(opts, out, "INITIAL_SUPPLY")
	return *ret0, err
}

// INITIALSUPPLY is a free data retrieval call binding the contract method 0x2ff2e9dc.
//
// Solidity: function INITIAL_SUPPLY() view returns(uint256)
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenSession) INITIALSUPPLY() (*big.Int, error) {
	return _ServiceChainKIP7Token.Contract.INITIALSUPPLY(&_ServiceChainKIP7Token.CallOpts)
}

// INITIALSUPPLY is a free data retrieval call binding the contract method 0x2ff2e9dc.
//
// Solidity: function INITIAL_SUPPLY() view returns(uint256)
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenCallerSession) INITIALSUPPLY() (*big.Int, error) {
	return _ServiceChainKIP7Token.Contract.INITIALSUPPLY(&_ServiceChainKIP7Token.CallOpts)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenCaller) NAME(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ServiceChainKIP7Token.contract.Call(opts, out, "NAME")
	return *ret0, err
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenSession) NAME() (string, error) {
	return _ServiceChainKIP7Token.Contract.NAME(&_ServiceChainKIP7Token.CallOpts)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenCallerSession) NAME() (string, error) {
	return _ServiceChainKIP7Token.Contract.NAME(&_ServiceChainKIP7Token.CallOpts)
}

// SYMBOL is a free data retrieval call binding the contract method 0xf76f8d78.
//
// Solidity: function SYMBOL() view returns(string)
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenCaller) SYMBOL(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ServiceChainKIP7Token.contract.Call(opts, out, "SYMBOL")
	return *ret0, err
}

// SYMBOL is a free data retrieval call binding the contract method 0xf76f8d78.
//
// Solidity: function SYMBOL() view returns(string)
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenSession) SYMBOL() (string, error) {
	return _ServiceChainKIP7Token.Contract.SYMBOL(&_ServiceChainKIP7Token.CallOpts)
}

// SYMBOL is a free data retrieval call binding the contract method 0xf76f8d78.
//
// Solidity: function SYMBOL() view returns(string)
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenCallerSession) SYMBOL() (string, error) {
	return _ServiceChainKIP7Token.Contract.SYMBOL(&_ServiceChainKIP7Token.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ServiceChainKIP7Token.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ServiceChainKIP7Token.Contract.Allowance(&_ServiceChainKIP7Token.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ServiceChainKIP7Token.Contract.Allowance(&_ServiceChainKIP7Token.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ServiceChainKIP7Token.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ServiceChainKIP7Token.Contract.BalanceOf(&_ServiceChainKIP7Token.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ServiceChainKIP7Token.Contract.BalanceOf(&_ServiceChainKIP7Token.CallOpts, account)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ServiceChainKIP7Token.contract.Call(opts, out, "bridge")
	return *ret0, err
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenSession) Bridge() (common.Address, error) {
	return _ServiceChainKIP7Token.Contract.Bridge(&_ServiceChainKIP7Token.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenCallerSession) Bridge() (common.Address, error) {
	return _ServiceChainKIP7Token.Contract.Bridge(&_ServiceChainKIP7Token.CallOpts)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) view returns(bool)
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenCaller) IsMinter(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ServiceChainKIP7Token.contract.Call(opts, out, "isMinter", account)
	return *ret0, err
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) view returns(bool)
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenSession) IsMinter(account common.Address) (bool, error) {
	return _ServiceChainKIP7Token.Contract.IsMinter(&_ServiceChainKIP7Token.CallOpts, account)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) view returns(bool)
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenCallerSession) IsMinter(account common.Address) (bool, error) {
	return _ServiceChainKIP7Token.Contract.IsMinter(&_ServiceChainKIP7Token.CallOpts, account)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ServiceChainKIP7Token.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenSession) IsOwner() (bool, error) {
	return _ServiceChainKIP7Token.Contract.IsOwner(&_ServiceChainKIP7Token.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenCallerSession) IsOwner() (bool, error) {
	return _ServiceChainKIP7Token.Contract.IsOwner(&_ServiceChainKIP7Token.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ServiceChainKIP7Token.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenSession) Owner() (common.Address, error) {
	return _ServiceChainKIP7Token.Contract.Owner(&_ServiceChainKIP7Token.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenCallerSession) Owner() (common.Address, error) {
	return _ServiceChainKIP7Token.Contract.Owner(&_ServiceChainKIP7Token.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ServiceChainKIP7Token.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ServiceChainKIP7Token.Contract.SupportsInterface(&_ServiceChainKIP7Token.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ServiceChainKIP7Token.Contract.SupportsInterface(&_ServiceChainKIP7Token.CallOpts, interfaceId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ServiceChainKIP7Token.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenSession) TotalSupply() (*big.Int, error) {
	return _ServiceChainKIP7Token.Contract.TotalSupply(&_ServiceChainKIP7Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenCallerSession) TotalSupply() (*big.Int, error) {
	return _ServiceChainKIP7Token.Contract.TotalSupply(&_ServiceChainKIP7Token.CallOpts)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenTransactor) AddMinter(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _ServiceChainKIP7Token.contract.Transact(opts, "addMinter", account)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenSession) AddMinter(account common.Address) (*types.Transaction, error) {
	return _ServiceChainKIP7Token.Contract.AddMinter(&_ServiceChainKIP7Token.TransactOpts, account)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenTransactorSession) AddMinter(account common.Address) (*types.Transaction, error) {
	return _ServiceChainKIP7Token.Contract.AddMinter(&_ServiceChainKIP7Token.TransactOpts, account)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ServiceChainKIP7Token.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ServiceChainKIP7Token.Contract.Approve(&_ServiceChainKIP7Token.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ServiceChainKIP7Token.Contract.Approve(&_ServiceChainKIP7Token.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ServiceChainKIP7Token.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _ServiceChainKIP7Token.Contract.Burn(&_ServiceChainKIP7Token.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _ServiceChainKIP7Token.Contract.Burn(&_ServiceChainKIP7Token.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ServiceChainKIP7Token.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ServiceChainKIP7Token.Contract.BurnFrom(&_ServiceChainKIP7Token.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ServiceChainKIP7Token.Contract.BurnFrom(&_ServiceChainKIP7Token.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns(bool)
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenTransactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ServiceChainKIP7Token.contract.Transact(opts, "mint", account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns(bool)
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ServiceChainKIP7Token.Contract.Mint(&_ServiceChainKIP7Token.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns(bool)
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenTransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ServiceChainKIP7Token.Contract.Mint(&_ServiceChainKIP7Token.TransactOpts, account, amount)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenTransactor) RenounceMinter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ServiceChainKIP7Token.contract.Transact(opts, "renounceMinter")
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenSession) RenounceMinter() (*types.Transaction, error) {
	return _ServiceChainKIP7Token.Contract.RenounceMinter(&_ServiceChainKIP7Token.TransactOpts)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenTransactorSession) RenounceMinter() (*types.Transaction, error) {
	return _ServiceChainKIP7Token.Contract.RenounceMinter(&_ServiceChainKIP7Token.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ServiceChainKIP7Token.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _ServiceChainKIP7Token.Contract.RenounceOwnership(&_ServiceChainKIP7Token.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ServiceChainKIP7Token.Contract.RenounceOwnership(&_ServiceChainKIP7Token.TransactOpts)
}

// RequestValueTransfer is a paid mutator transaction binding the contract method 0x3dc3c9e1.
//
// Solidity: function requestValueTransfer(uint256 _amount, address _to, uint256 _feeLimit, bytes _extraData) returns()
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenTransactor) RequestValueTransfer(opts *bind.TransactOpts, _amount *big.Int, _to common.Address, _feeLimit *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _ServiceChainKIP7Token.contract.Transact(opts, "requestValueTransfer", _amount, _to, _feeLimit, _extraData)
}

// RequestValueTransfer is a paid mutator transaction binding the contract method 0x3dc3c9e1.
//
// Solidity: function requestValueTransfer(uint256 _amount, address _to, uint256 _feeLimit, bytes _extraData) returns()
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenSession) RequestValueTransfer(_amount *big.Int, _to common.Address, _feeLimit *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _ServiceChainKIP7Token.Contract.RequestValueTransfer(&_ServiceChainKIP7Token.TransactOpts, _amount, _to, _feeLimit, _extraData)
}

// RequestValueTransfer is a paid mutator transaction binding the contract method 0x3dc3c9e1.
//
// Solidity: function requestValueTransfer(uint256 _amount, address _to, uint256 _feeLimit, bytes _extraData) returns()
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenTransactorSession) RequestValueTransfer(_amount *big.Int, _to common.Address, _feeLimit *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _ServiceChainKIP7Token.Contract.RequestValueTransfer(&_ServiceChainKIP7Token.TransactOpts, _amount, _to, _feeLimit, _extraData)
}

// SafeTransfer is a paid mutator transaction binding the contract method 0x423f6cef.
//
// Solidity: function safeTransfer(address recipient, uint256 amount) returns()
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenTransactor) SafeTransfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ServiceChainKIP7Token.contract.Transact(opts, "safeTransfer", recipient, amount)
}

// SafeTransfer is a paid mutator transaction binding the contract method 0x423f6cef.
//
// Solidity: function safeTransfer(address recipient, uint256 amount) returns()
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenSession) SafeTransfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ServiceChainKIP7Token.Contract.SafeTransfer(&_ServiceChainKIP7Token.TransactOpts, recipient, amount)
}

// SafeTransfer is a paid mutator transaction binding the contract method 0x423f6cef.
//
// Solidity: function safeTransfer(address recipient, uint256 amount) returns()
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenTransactorSession) SafeTransfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ServiceChainKIP7Token.Contract.SafeTransfer(&_ServiceChainKIP7Token.TransactOpts, recipient, amount)
}

// SafeTransfer0 is a paid mutator transaction binding the contract method 0xeb795549.
//
// Solidity: function safeTransfer(address recipient, uint256 amount, bytes data) returns()
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenTransactor) SafeTransfer0(opts *bind.TransactOpts, recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ServiceChainKIP7Token.contract.Transact(opts, "safeTransfer0", recipient, amount, data)
}

// SafeTransfer0 is a paid mutator transaction binding the contract method 0xeb795549.
//
// Solidity: function safeTransfer(address recipient, uint256 amount, bytes data) returns()
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenSession) SafeTransfer0(recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ServiceChainKIP7Token.Contract.SafeTransfer0(&_ServiceChainKIP7Token.TransactOpts, recipient, amount, data)
}

// SafeTransfer0 is a paid mutator transaction binding the contract method 0xeb795549.
//
// Solidity: function safeTransfer(address recipient, uint256 amount, bytes data) returns()
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenTransactorSession) SafeTransfer0(recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ServiceChainKIP7Token.Contract.SafeTransfer0(&_ServiceChainKIP7Token.TransactOpts, recipient, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address sender, address recipient, uint256 amount) returns()
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenTransactor) SafeTransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ServiceChainKIP7Token.contract.Transact(opts, "safeTransferFrom", sender, recipient, amount)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address sender, address recipient, uint256 amount) returns()
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenSession) SafeTransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ServiceChainKIP7Token.Contract.SafeTransferFrom(&_ServiceChainKIP7Token.TransactOpts, sender, recipient, amount)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address sender, address recipient, uint256 amount) returns()
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenTransactorSession) SafeTransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ServiceChainKIP7Token.Contract.SafeTransferFrom(&_ServiceChainKIP7Token.TransactOpts, sender, recipient, amount)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address sender, address recipient, uint256 amount, bytes data) returns()
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenTransactor) SafeTransferFrom0(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ServiceChainKIP7Token.contract.Transact(opts, "safeTransferFrom0", sender, recipient, amount, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address sender, address recipient, uint256 amount, bytes data) returns()
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenSession) SafeTransferFrom0(sender common.Address, recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ServiceChainKIP7Token.Contract.SafeTransferFrom0(&_ServiceChainKIP7Token.TransactOpts, sender, recipient, amount, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address sender, address recipient, uint256 amount, bytes data) returns()
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenTransactorSession) SafeTransferFrom0(sender common.Address, recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ServiceChainKIP7Token.Contract.SafeTransferFrom0(&_ServiceChainKIP7Token.TransactOpts, sender, recipient, amount, data)
}

// SetBridge is a paid mutator transaction binding the contract method 0x8dd14802.
//
// Solidity: function setBridge(address _bridge) returns()
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenTransactor) SetBridge(opts *bind.TransactOpts, _bridge common.Address) (*types.Transaction, error) {
	return _ServiceChainKIP7Token.contract.Transact(opts, "setBridge", _bridge)
}

// SetBridge is a paid mutator transaction binding the contract method 0x8dd14802.
//
// Solidity: function setBridge(address _bridge) returns()
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenSession) SetBridge(_bridge common.Address) (*types.Transaction, error) {
	return _ServiceChainKIP7Token.Contract.SetBridge(&_ServiceChainKIP7Token.TransactOpts, _bridge)
}

// SetBridge is a paid mutator transaction binding the contract method 0x8dd14802.
//
// Solidity: function setBridge(address _bridge) returns()
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenTransactorSession) SetBridge(_bridge common.Address) (*types.Transaction, error) {
	return _ServiceChainKIP7Token.Contract.SetBridge(&_ServiceChainKIP7Token.TransactOpts, _bridge)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ServiceChainKIP7Token.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ServiceChainKIP7Token.Contract.Transfer(&_ServiceChainKIP7Token.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ServiceChainKIP7Token.Contract.Transfer(&_ServiceChainKIP7Token.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ServiceChainKIP7Token.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ServiceChainKIP7Token.Contract.TransferFrom(&_ServiceChainKIP7Token.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ServiceChainKIP7Token.Contract.TransferFrom(&_ServiceChainKIP7Token.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ServiceChainKIP7Token.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ServiceChainKIP7Token.Contract.TransferOwnership(&_ServiceChainKIP7Token.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ServiceChainKIP7Token.Contract.TransferOwnership(&_ServiceChainKIP7Token.TransactOpts, newOwner)
}

// ServiceChainKIP7TokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ServiceChainKIP7Token contract.
type ServiceChainKIP7TokenApprovalIterator struct {
	Event *ServiceChainKIP7TokenApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ServiceChainKIP7TokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ServiceChainKIP7TokenApproval)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ServiceChainKIP7TokenApproval)
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
func (it *ServiceChainKIP7TokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ServiceChainKIP7TokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ServiceChainKIP7TokenApproval represents a Approval event raised by the ServiceChainKIP7Token contract.
type ServiceChainKIP7TokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ServiceChainKIP7TokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ServiceChainKIP7Token.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ServiceChainKIP7TokenApprovalIterator{contract: _ServiceChainKIP7Token.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ServiceChainKIP7TokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ServiceChainKIP7Token.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ServiceChainKIP7TokenApproval)
				if err := _ServiceChainKIP7Token.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenFilterer) ParseApproval(log types.Log) (*ServiceChainKIP7TokenApproval, error) {
	event := new(ServiceChainKIP7TokenApproval)
	if err := _ServiceChainKIP7Token.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ServiceChainKIP7TokenMinterAddedIterator is returned from FilterMinterAdded and is used to iterate over the raw logs and unpacked data for MinterAdded events raised by the ServiceChainKIP7Token contract.
type ServiceChainKIP7TokenMinterAddedIterator struct {
	Event *ServiceChainKIP7TokenMinterAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ServiceChainKIP7TokenMinterAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ServiceChainKIP7TokenMinterAdded)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ServiceChainKIP7TokenMinterAdded)
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
func (it *ServiceChainKIP7TokenMinterAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ServiceChainKIP7TokenMinterAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ServiceChainKIP7TokenMinterAdded represents a MinterAdded event raised by the ServiceChainKIP7Token contract.
type ServiceChainKIP7TokenMinterAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinterAdded is a free log retrieval operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed account)
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenFilterer) FilterMinterAdded(opts *bind.FilterOpts, account []common.Address) (*ServiceChainKIP7TokenMinterAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ServiceChainKIP7Token.contract.FilterLogs(opts, "MinterAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &ServiceChainKIP7TokenMinterAddedIterator{contract: _ServiceChainKIP7Token.contract, event: "MinterAdded", logs: logs, sub: sub}, nil
}

// WatchMinterAdded is a free log subscription operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed account)
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenFilterer) WatchMinterAdded(opts *bind.WatchOpts, sink chan<- *ServiceChainKIP7TokenMinterAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ServiceChainKIP7Token.contract.WatchLogs(opts, "MinterAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ServiceChainKIP7TokenMinterAdded)
				if err := _ServiceChainKIP7Token.contract.UnpackLog(event, "MinterAdded", log); err != nil {
					return err
				}
				event.Raw = log

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

// ParseMinterAdded is a log parse operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed account)
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenFilterer) ParseMinterAdded(log types.Log) (*ServiceChainKIP7TokenMinterAdded, error) {
	event := new(ServiceChainKIP7TokenMinterAdded)
	if err := _ServiceChainKIP7Token.contract.UnpackLog(event, "MinterAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ServiceChainKIP7TokenMinterRemovedIterator is returned from FilterMinterRemoved and is used to iterate over the raw logs and unpacked data for MinterRemoved events raised by the ServiceChainKIP7Token contract.
type ServiceChainKIP7TokenMinterRemovedIterator struct {
	Event *ServiceChainKIP7TokenMinterRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ServiceChainKIP7TokenMinterRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ServiceChainKIP7TokenMinterRemoved)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ServiceChainKIP7TokenMinterRemoved)
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
func (it *ServiceChainKIP7TokenMinterRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ServiceChainKIP7TokenMinterRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ServiceChainKIP7TokenMinterRemoved represents a MinterRemoved event raised by the ServiceChainKIP7Token contract.
type ServiceChainKIP7TokenMinterRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinterRemoved is a free log retrieval operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed account)
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenFilterer) FilterMinterRemoved(opts *bind.FilterOpts, account []common.Address) (*ServiceChainKIP7TokenMinterRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ServiceChainKIP7Token.contract.FilterLogs(opts, "MinterRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &ServiceChainKIP7TokenMinterRemovedIterator{contract: _ServiceChainKIP7Token.contract, event: "MinterRemoved", logs: logs, sub: sub}, nil
}

// WatchMinterRemoved is a free log subscription operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed account)
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenFilterer) WatchMinterRemoved(opts *bind.WatchOpts, sink chan<- *ServiceChainKIP7TokenMinterRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ServiceChainKIP7Token.contract.WatchLogs(opts, "MinterRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ServiceChainKIP7TokenMinterRemoved)
				if err := _ServiceChainKIP7Token.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
					return err
				}
				event.Raw = log

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

// ParseMinterRemoved is a log parse operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed account)
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenFilterer) ParseMinterRemoved(log types.Log) (*ServiceChainKIP7TokenMinterRemoved, error) {
	event := new(ServiceChainKIP7TokenMinterRemoved)
	if err := _ServiceChainKIP7Token.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ServiceChainKIP7TokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ServiceChainKIP7Token contract.
type ServiceChainKIP7TokenOwnershipTransferredIterator struct {
	Event *ServiceChainKIP7TokenOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ServiceChainKIP7TokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ServiceChainKIP7TokenOwnershipTransferred)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ServiceChainKIP7TokenOwnershipTransferred)
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
func (it *ServiceChainKIP7TokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ServiceChainKIP7TokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ServiceChainKIP7TokenOwnershipTransferred represents a OwnershipTransferred event raised by the ServiceChainKIP7Token contract.
type ServiceChainKIP7TokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ServiceChainKIP7TokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ServiceChainKIP7Token.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ServiceChainKIP7TokenOwnershipTransferredIterator{contract: _ServiceChainKIP7Token.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ServiceChainKIP7TokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ServiceChainKIP7Token.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ServiceChainKIP7TokenOwnershipTransferred)
				if err := _ServiceChainKIP7Token.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

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
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenFilterer) ParseOwnershipTransferred(log types.Log) (*ServiceChainKIP7TokenOwnershipTransferred, error) {
	event := new(ServiceChainKIP7TokenOwnershipTransferred)
	if err := _ServiceChainKIP7Token.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ServiceChainKIP7TokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ServiceChainKIP7Token contract.
type ServiceChainKIP7TokenTransferIterator struct {
	Event *ServiceChainKIP7TokenTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ServiceChainKIP7TokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ServiceChainKIP7TokenTransfer)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ServiceChainKIP7TokenTransfer)
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
func (it *ServiceChainKIP7TokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ServiceChainKIP7TokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ServiceChainKIP7TokenTransfer represents a Transfer event raised by the ServiceChainKIP7Token contract.
type ServiceChainKIP7TokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ServiceChainKIP7TokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ServiceChainKIP7Token.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ServiceChainKIP7TokenTransferIterator{contract: _ServiceChainKIP7Token.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ServiceChainKIP7TokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ServiceChainKIP7Token.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ServiceChainKIP7TokenTransfer)
				if err := _ServiceChainKIP7Token.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ServiceChainKIP7Token *ServiceChainKIP7TokenFilterer) ParseTransfer(log types.Log) (*ServiceChainKIP7TokenTransfer, error) {
	event := new(ServiceChainKIP7TokenTransfer)
	if err := _ServiceChainKIP7Token.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
