// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bridge

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

// BridgeABI is the input ABI used to generate the binding from.
const BridgeABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_dummy\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"}]"

// BridgeBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const BridgeBinRuntime = `6080604052348015600f57600080fd5b506004361060325760003560e01c806301ffc9a7146037578063ffa1ad7414606f575b600080fd5b605b60048036036020811015604b57600080fd5b50356001600160e01b0319166092565b604080519115158252519081900360200190f35b607560b1565b6040805167ffffffffffffffff9092168252519081900360200190f35b6001600160e01b03191660009081526020819052604090205460ff1690565b60018156fea165627a7a7230582024b2adcbcb791242a7c02d51ab9d186c161d0a8ded27424ed79cab7673e1b65f0029`

// BridgeFuncSigs maps the 4-byte function signature to its string representation.
var BridgeFuncSigs = map[string]string{
	"ffa1ad74": "VERSION()",
	"01ffc9a7": "supportsInterface(bytes4)",
}

// BridgeBin is the compiled bytecode used for deploying new contracts.
var BridgeBin = "0x608060408190526001805463ffffffff19168117600160201b600160601b031916905560208061029d8339810180604052602081101561003e57600080fd5b50516100707f01ffc9a7000000000000000000000000000000000000000000000000000000006100df602090811b901c565b6001546100869060e01b6100df602090811b901c565b6001546100a490640100000000900460e01b6100df602090811b901c565b6001546100c69068010000000000000000900460e01b6100df602090811b901c565b6100d9600460e01b6100df60201b60201c565b506101ad565b7fffffffff00000000000000000000000000000000000000000000000000000000808216141561017057604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f4b495031333a20696e76616c696420696e746572666163652069640000000000604482015290519081900360640190fd5b7fffffffff00000000000000000000000000000000000000000000000000000000166000908152602081905260409020805460ff19166001179055565b60e2806101bb6000396000f3fe6080604052348015600f57600080fd5b506004361060325760003560e01c806301ffc9a7146037578063ffa1ad7414606f575b600080fd5b605b60048036036020811015604b57600080fd5b50356001600160e01b0319166092565b604080519115158252519081900360200190f35b607560b1565b6040805167ffffffffffffffff9092168252519081900360200190f35b6001600160e01b03191660009081526020819052604090205460ff1690565b60018156fea165627a7a7230582024b2adcbcb791242a7c02d51ab9d186c161d0a8ded27424ed79cab7673e1b65f0029"

// DeployBridge deploys a new Klaytn contract, binding an instance of Bridge to it.
func DeployBridge(auth *bind.TransactOpts, backend bind.ContractBackend, _dummy bool) (common.Address, *types.Transaction, *Bridge, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BridgeBin), backend, _dummy)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// Bridge is an auto generated Go binding around a Klaytn contract.
type Bridge struct {
	BridgeCaller     // Read-only binding to the contract
	BridgeTransactor // Write-only binding to the contract
	BridgeFilterer   // Log filterer for contract events
}

// BridgeCaller is an auto generated read-only Go binding around a Klaytn contract.
type BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTransactor is an auto generated write-only Go binding around a Klaytn contract.
type BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type BridgeSession struct {
	Contract     *Bridge           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type BridgeCallerSession struct {
	Contract *BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BridgeTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type BridgeTransactorSession struct {
	Contract     *BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeRaw is an auto generated low-level Go binding around a Klaytn contract.
type BridgeRaw struct {
	Contract *Bridge // Generic contract binding to access the raw methods on
}

// BridgeCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type BridgeCallerRaw struct {
	Contract *BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type BridgeTransactorRaw struct {
	Contract *BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridge creates a new instance of Bridge, bound to a specific deployed contract.
func NewBridge(address common.Address, backend bind.ContractBackend) (*Bridge, error) {
	contract, err := bindBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// NewBridgeCaller creates a new read-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeCaller(address common.Address, caller bind.ContractCaller) (*BridgeCaller, error) {
	contract, err := bindBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeCaller{contract: contract}, nil
}

// NewBridgeTransactor creates a new write-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeTransactor, error) {
	contract, err := bindBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTransactor{contract: contract}, nil
}

// NewBridgeFilterer creates a new log filterer instance of Bridge, bound to a specific deployed contract.
func NewBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeFilterer, error) {
	contract, err := bindBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeFilterer{contract: contract}, nil
}

// bindBridge binds a generic wrapper to an already deployed contract.
func bindBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transact(opts, method, params...)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint64)
func (_Bridge *BridgeCaller) VERSION(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _Bridge.contract.Call(opts, out, "VERSION")
	return *ret0, err
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint64)
func (_Bridge *BridgeSession) VERSION() (uint64, error) {
	return _Bridge.Contract.VERSION(&_Bridge.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint64)
func (_Bridge *BridgeCallerSession) VERSION() (uint64, error) {
	return _Bridge.Contract.VERSION(&_Bridge.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bridge *BridgeCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Bridge.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bridge *BridgeSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Bridge.Contract.SupportsInterface(&_Bridge.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bridge *BridgeCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Bridge.Contract.SupportsInterface(&_Bridge.CallOpts, interfaceId)
}

// BridgeTransferABI is the input ABI used to generate the binding from.
const BridgeTransferABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// BridgeTransferBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const BridgeTransferBinRuntime = ``

// BridgeTransferFuncSigs maps the 4-byte function signature to its string representation.
var BridgeTransferFuncSigs = map[string]string{
	"01ffc9a7": "supportsInterface(bytes4)",
}

// BridgeTransfer is an auto generated Go binding around a Klaytn contract.
type BridgeTransfer struct {
	BridgeTransferCaller     // Read-only binding to the contract
	BridgeTransferTransactor // Write-only binding to the contract
	BridgeTransferFilterer   // Log filterer for contract events
}

// BridgeTransferCaller is an auto generated read-only Go binding around a Klaytn contract.
type BridgeTransferCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTransferTransactor is an auto generated write-only Go binding around a Klaytn contract.
type BridgeTransferTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTransferFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type BridgeTransferFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTransferSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type BridgeTransferSession struct {
	Contract     *BridgeTransfer   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeTransferCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type BridgeTransferCallerSession struct {
	Contract *BridgeTransferCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// BridgeTransferTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type BridgeTransferTransactorSession struct {
	Contract     *BridgeTransferTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// BridgeTransferRaw is an auto generated low-level Go binding around a Klaytn contract.
type BridgeTransferRaw struct {
	Contract *BridgeTransfer // Generic contract binding to access the raw methods on
}

// BridgeTransferCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type BridgeTransferCallerRaw struct {
	Contract *BridgeTransferCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeTransferTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type BridgeTransferTransactorRaw struct {
	Contract *BridgeTransferTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeTransfer creates a new instance of BridgeTransfer, bound to a specific deployed contract.
func NewBridgeTransfer(address common.Address, backend bind.ContractBackend) (*BridgeTransfer, error) {
	contract, err := bindBridgeTransfer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeTransfer{BridgeTransferCaller: BridgeTransferCaller{contract: contract}, BridgeTransferTransactor: BridgeTransferTransactor{contract: contract}, BridgeTransferFilterer: BridgeTransferFilterer{contract: contract}}, nil
}

// NewBridgeTransferCaller creates a new read-only instance of BridgeTransfer, bound to a specific deployed contract.
func NewBridgeTransferCaller(address common.Address, caller bind.ContractCaller) (*BridgeTransferCaller, error) {
	contract, err := bindBridgeTransfer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTransferCaller{contract: contract}, nil
}

// NewBridgeTransferTransactor creates a new write-only instance of BridgeTransfer, bound to a specific deployed contract.
func NewBridgeTransferTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeTransferTransactor, error) {
	contract, err := bindBridgeTransfer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTransferTransactor{contract: contract}, nil
}

// NewBridgeTransferFilterer creates a new log filterer instance of BridgeTransfer, bound to a specific deployed contract.
func NewBridgeTransferFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeTransferFilterer, error) {
	contract, err := bindBridgeTransfer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeTransferFilterer{contract: contract}, nil
}

// bindBridgeTransfer binds a generic wrapper to an already deployed contract.
func bindBridgeTransfer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeTransferABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeTransfer *BridgeTransferRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BridgeTransfer.Contract.BridgeTransferCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeTransfer *BridgeTransferRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeTransfer.Contract.BridgeTransferTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeTransfer *BridgeTransferRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeTransfer.Contract.BridgeTransferTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeTransfer *BridgeTransferCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BridgeTransfer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeTransfer *BridgeTransferTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeTransfer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeTransfer *BridgeTransferTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeTransfer.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BridgeTransfer *BridgeTransferCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BridgeTransfer.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BridgeTransfer *BridgeTransferSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BridgeTransfer.Contract.SupportsInterface(&_BridgeTransfer.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BridgeTransfer *BridgeTransferCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BridgeTransfer.Contract.SupportsInterface(&_BridgeTransfer.CallOpts, interfaceId)
}

// BridgeTransferERC20ABI is the input ABI used to generate the binding from.
const BridgeTransferERC20ABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// BridgeTransferERC20BinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const BridgeTransferERC20BinRuntime = ``

// BridgeTransferERC20FuncSigs maps the 4-byte function signature to its string representation.
var BridgeTransferERC20FuncSigs = map[string]string{
	"01ffc9a7": "supportsInterface(bytes4)",
}

// BridgeTransferERC20 is an auto generated Go binding around a Klaytn contract.
type BridgeTransferERC20 struct {
	BridgeTransferERC20Caller     // Read-only binding to the contract
	BridgeTransferERC20Transactor // Write-only binding to the contract
	BridgeTransferERC20Filterer   // Log filterer for contract events
}

// BridgeTransferERC20Caller is an auto generated read-only Go binding around a Klaytn contract.
type BridgeTransferERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTransferERC20Transactor is an auto generated write-only Go binding around a Klaytn contract.
type BridgeTransferERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTransferERC20Filterer is an auto generated log filtering Go binding around a Klaytn contract events.
type BridgeTransferERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTransferERC20Session is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type BridgeTransferERC20Session struct {
	Contract     *BridgeTransferERC20 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BridgeTransferERC20CallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type BridgeTransferERC20CallerSession struct {
	Contract *BridgeTransferERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// BridgeTransferERC20TransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type BridgeTransferERC20TransactorSession struct {
	Contract     *BridgeTransferERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// BridgeTransferERC20Raw is an auto generated low-level Go binding around a Klaytn contract.
type BridgeTransferERC20Raw struct {
	Contract *BridgeTransferERC20 // Generic contract binding to access the raw methods on
}

// BridgeTransferERC20CallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type BridgeTransferERC20CallerRaw struct {
	Contract *BridgeTransferERC20Caller // Generic read-only contract binding to access the raw methods on
}

// BridgeTransferERC20TransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type BridgeTransferERC20TransactorRaw struct {
	Contract *BridgeTransferERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeTransferERC20 creates a new instance of BridgeTransferERC20, bound to a specific deployed contract.
func NewBridgeTransferERC20(address common.Address, backend bind.ContractBackend) (*BridgeTransferERC20, error) {
	contract, err := bindBridgeTransferERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeTransferERC20{BridgeTransferERC20Caller: BridgeTransferERC20Caller{contract: contract}, BridgeTransferERC20Transactor: BridgeTransferERC20Transactor{contract: contract}, BridgeTransferERC20Filterer: BridgeTransferERC20Filterer{contract: contract}}, nil
}

// NewBridgeTransferERC20Caller creates a new read-only instance of BridgeTransferERC20, bound to a specific deployed contract.
func NewBridgeTransferERC20Caller(address common.Address, caller bind.ContractCaller) (*BridgeTransferERC20Caller, error) {
	contract, err := bindBridgeTransferERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTransferERC20Caller{contract: contract}, nil
}

// NewBridgeTransferERC20Transactor creates a new write-only instance of BridgeTransferERC20, bound to a specific deployed contract.
func NewBridgeTransferERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*BridgeTransferERC20Transactor, error) {
	contract, err := bindBridgeTransferERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTransferERC20Transactor{contract: contract}, nil
}

// NewBridgeTransferERC20Filterer creates a new log filterer instance of BridgeTransferERC20, bound to a specific deployed contract.
func NewBridgeTransferERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*BridgeTransferERC20Filterer, error) {
	contract, err := bindBridgeTransferERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeTransferERC20Filterer{contract: contract}, nil
}

// bindBridgeTransferERC20 binds a generic wrapper to an already deployed contract.
func bindBridgeTransferERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeTransferERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeTransferERC20 *BridgeTransferERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BridgeTransferERC20.Contract.BridgeTransferERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeTransferERC20 *BridgeTransferERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeTransferERC20.Contract.BridgeTransferERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeTransferERC20 *BridgeTransferERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeTransferERC20.Contract.BridgeTransferERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeTransferERC20 *BridgeTransferERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BridgeTransferERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeTransferERC20 *BridgeTransferERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeTransferERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeTransferERC20 *BridgeTransferERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeTransferERC20.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BridgeTransferERC20 *BridgeTransferERC20Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BridgeTransferERC20.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BridgeTransferERC20 *BridgeTransferERC20Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BridgeTransferERC20.Contract.SupportsInterface(&_BridgeTransferERC20.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BridgeTransferERC20 *BridgeTransferERC20CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BridgeTransferERC20.Contract.SupportsInterface(&_BridgeTransferERC20.CallOpts, interfaceId)
}

// BridgeTransferERC721ABI is the input ABI used to generate the binding from.
const BridgeTransferERC721ABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// BridgeTransferERC721BinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const BridgeTransferERC721BinRuntime = ``

// BridgeTransferERC721FuncSigs maps the 4-byte function signature to its string representation.
var BridgeTransferERC721FuncSigs = map[string]string{
	"01ffc9a7": "supportsInterface(bytes4)",
}

// BridgeTransferERC721 is an auto generated Go binding around a Klaytn contract.
type BridgeTransferERC721 struct {
	BridgeTransferERC721Caller     // Read-only binding to the contract
	BridgeTransferERC721Transactor // Write-only binding to the contract
	BridgeTransferERC721Filterer   // Log filterer for contract events
}

// BridgeTransferERC721Caller is an auto generated read-only Go binding around a Klaytn contract.
type BridgeTransferERC721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTransferERC721Transactor is an auto generated write-only Go binding around a Klaytn contract.
type BridgeTransferERC721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTransferERC721Filterer is an auto generated log filtering Go binding around a Klaytn contract events.
type BridgeTransferERC721Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTransferERC721Session is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type BridgeTransferERC721Session struct {
	Contract     *BridgeTransferERC721 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BridgeTransferERC721CallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type BridgeTransferERC721CallerSession struct {
	Contract *BridgeTransferERC721Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// BridgeTransferERC721TransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type BridgeTransferERC721TransactorSession struct {
	Contract     *BridgeTransferERC721Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// BridgeTransferERC721Raw is an auto generated low-level Go binding around a Klaytn contract.
type BridgeTransferERC721Raw struct {
	Contract *BridgeTransferERC721 // Generic contract binding to access the raw methods on
}

// BridgeTransferERC721CallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type BridgeTransferERC721CallerRaw struct {
	Contract *BridgeTransferERC721Caller // Generic read-only contract binding to access the raw methods on
}

// BridgeTransferERC721TransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type BridgeTransferERC721TransactorRaw struct {
	Contract *BridgeTransferERC721Transactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeTransferERC721 creates a new instance of BridgeTransferERC721, bound to a specific deployed contract.
func NewBridgeTransferERC721(address common.Address, backend bind.ContractBackend) (*BridgeTransferERC721, error) {
	contract, err := bindBridgeTransferERC721(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeTransferERC721{BridgeTransferERC721Caller: BridgeTransferERC721Caller{contract: contract}, BridgeTransferERC721Transactor: BridgeTransferERC721Transactor{contract: contract}, BridgeTransferERC721Filterer: BridgeTransferERC721Filterer{contract: contract}}, nil
}

// NewBridgeTransferERC721Caller creates a new read-only instance of BridgeTransferERC721, bound to a specific deployed contract.
func NewBridgeTransferERC721Caller(address common.Address, caller bind.ContractCaller) (*BridgeTransferERC721Caller, error) {
	contract, err := bindBridgeTransferERC721(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTransferERC721Caller{contract: contract}, nil
}

// NewBridgeTransferERC721Transactor creates a new write-only instance of BridgeTransferERC721, bound to a specific deployed contract.
func NewBridgeTransferERC721Transactor(address common.Address, transactor bind.ContractTransactor) (*BridgeTransferERC721Transactor, error) {
	contract, err := bindBridgeTransferERC721(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTransferERC721Transactor{contract: contract}, nil
}

// NewBridgeTransferERC721Filterer creates a new log filterer instance of BridgeTransferERC721, bound to a specific deployed contract.
func NewBridgeTransferERC721Filterer(address common.Address, filterer bind.ContractFilterer) (*BridgeTransferERC721Filterer, error) {
	contract, err := bindBridgeTransferERC721(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeTransferERC721Filterer{contract: contract}, nil
}

// bindBridgeTransferERC721 binds a generic wrapper to an already deployed contract.
func bindBridgeTransferERC721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeTransferERC721ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeTransferERC721 *BridgeTransferERC721Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BridgeTransferERC721.Contract.BridgeTransferERC721Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeTransferERC721 *BridgeTransferERC721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeTransferERC721.Contract.BridgeTransferERC721Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeTransferERC721 *BridgeTransferERC721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeTransferERC721.Contract.BridgeTransferERC721Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeTransferERC721 *BridgeTransferERC721CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BridgeTransferERC721.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeTransferERC721 *BridgeTransferERC721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeTransferERC721.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeTransferERC721 *BridgeTransferERC721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeTransferERC721.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BridgeTransferERC721 *BridgeTransferERC721Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BridgeTransferERC721.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BridgeTransferERC721 *BridgeTransferERC721Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BridgeTransferERC721.Contract.SupportsInterface(&_BridgeTransferERC721.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BridgeTransferERC721 *BridgeTransferERC721CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BridgeTransferERC721.Contract.SupportsInterface(&_BridgeTransferERC721.CallOpts, interfaceId)
}

// BridgeTransferKLAYABI is the input ABI used to generate the binding from.
const BridgeTransferKLAYABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// BridgeTransferKLAYBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const BridgeTransferKLAYBinRuntime = ``

// BridgeTransferKLAYFuncSigs maps the 4-byte function signature to its string representation.
var BridgeTransferKLAYFuncSigs = map[string]string{
	"01ffc9a7": "supportsInterface(bytes4)",
}

// BridgeTransferKLAY is an auto generated Go binding around a Klaytn contract.
type BridgeTransferKLAY struct {
	BridgeTransferKLAYCaller     // Read-only binding to the contract
	BridgeTransferKLAYTransactor // Write-only binding to the contract
	BridgeTransferKLAYFilterer   // Log filterer for contract events
}

// BridgeTransferKLAYCaller is an auto generated read-only Go binding around a Klaytn contract.
type BridgeTransferKLAYCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTransferKLAYTransactor is an auto generated write-only Go binding around a Klaytn contract.
type BridgeTransferKLAYTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTransferKLAYFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type BridgeTransferKLAYFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTransferKLAYSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type BridgeTransferKLAYSession struct {
	Contract     *BridgeTransferKLAY // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BridgeTransferKLAYCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type BridgeTransferKLAYCallerSession struct {
	Contract *BridgeTransferKLAYCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// BridgeTransferKLAYTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type BridgeTransferKLAYTransactorSession struct {
	Contract     *BridgeTransferKLAYTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// BridgeTransferKLAYRaw is an auto generated low-level Go binding around a Klaytn contract.
type BridgeTransferKLAYRaw struct {
	Contract *BridgeTransferKLAY // Generic contract binding to access the raw methods on
}

// BridgeTransferKLAYCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type BridgeTransferKLAYCallerRaw struct {
	Contract *BridgeTransferKLAYCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeTransferKLAYTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type BridgeTransferKLAYTransactorRaw struct {
	Contract *BridgeTransferKLAYTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeTransferKLAY creates a new instance of BridgeTransferKLAY, bound to a specific deployed contract.
func NewBridgeTransferKLAY(address common.Address, backend bind.ContractBackend) (*BridgeTransferKLAY, error) {
	contract, err := bindBridgeTransferKLAY(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeTransferKLAY{BridgeTransferKLAYCaller: BridgeTransferKLAYCaller{contract: contract}, BridgeTransferKLAYTransactor: BridgeTransferKLAYTransactor{contract: contract}, BridgeTransferKLAYFilterer: BridgeTransferKLAYFilterer{contract: contract}}, nil
}

// NewBridgeTransferKLAYCaller creates a new read-only instance of BridgeTransferKLAY, bound to a specific deployed contract.
func NewBridgeTransferKLAYCaller(address common.Address, caller bind.ContractCaller) (*BridgeTransferKLAYCaller, error) {
	contract, err := bindBridgeTransferKLAY(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTransferKLAYCaller{contract: contract}, nil
}

// NewBridgeTransferKLAYTransactor creates a new write-only instance of BridgeTransferKLAY, bound to a specific deployed contract.
func NewBridgeTransferKLAYTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeTransferKLAYTransactor, error) {
	contract, err := bindBridgeTransferKLAY(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTransferKLAYTransactor{contract: contract}, nil
}

// NewBridgeTransferKLAYFilterer creates a new log filterer instance of BridgeTransferKLAY, bound to a specific deployed contract.
func NewBridgeTransferKLAYFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeTransferKLAYFilterer, error) {
	contract, err := bindBridgeTransferKLAY(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeTransferKLAYFilterer{contract: contract}, nil
}

// bindBridgeTransferKLAY binds a generic wrapper to an already deployed contract.
func bindBridgeTransferKLAY(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeTransferKLAYABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeTransferKLAY *BridgeTransferKLAYRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BridgeTransferKLAY.Contract.BridgeTransferKLAYCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeTransferKLAY *BridgeTransferKLAYRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeTransferKLAY.Contract.BridgeTransferKLAYTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeTransferKLAY *BridgeTransferKLAYRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeTransferKLAY.Contract.BridgeTransferKLAYTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeTransferKLAY *BridgeTransferKLAYCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BridgeTransferKLAY.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeTransferKLAY *BridgeTransferKLAYTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeTransferKLAY.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeTransferKLAY *BridgeTransferKLAYTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeTransferKLAY.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BridgeTransferKLAY *BridgeTransferKLAYCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BridgeTransferKLAY.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BridgeTransferKLAY *BridgeTransferKLAYSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BridgeTransferKLAY.Contract.SupportsInterface(&_BridgeTransferKLAY.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BridgeTransferKLAY *BridgeTransferKLAYCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BridgeTransferKLAY.Contract.SupportsInterface(&_BridgeTransferKLAY.CallOpts, interfaceId)
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
