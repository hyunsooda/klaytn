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

// CountersABI is the input ABI used to generate the binding from.
const CountersABI = "[]"

// CountersBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const CountersBinRuntime = `73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a723058207329cab25de591c35947480820e66bd7e2d232979872616f21ac6e5783a017700029`

// CountersBin is the compiled bytecode used for deploying new contracts.
var CountersBin = "0x604c6023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a723058207329cab25de591c35947480820e66bd7e2d232979872616f21ac6e5783a017700029"

// DeployCounters deploys a new Klaytn contract, binding an instance of Counters to it.
func DeployCounters(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Counters, error) {
	parsed, err := abi.JSON(strings.NewReader(CountersABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CountersBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Counters{CountersCaller: CountersCaller{contract: contract}, CountersTransactor: CountersTransactor{contract: contract}, CountersFilterer: CountersFilterer{contract: contract}}, nil
}

// Counters is an auto generated Go binding around a Klaytn contract.
type Counters struct {
	CountersCaller     // Read-only binding to the contract
	CountersTransactor // Write-only binding to the contract
	CountersFilterer   // Log filterer for contract events
}

// CountersCaller is an auto generated read-only Go binding around a Klaytn contract.
type CountersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CountersTransactor is an auto generated write-only Go binding around a Klaytn contract.
type CountersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CountersFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type CountersFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CountersSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type CountersSession struct {
	Contract     *Counters         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CountersCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type CountersCallerSession struct {
	Contract *CountersCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// CountersTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type CountersTransactorSession struct {
	Contract     *CountersTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CountersRaw is an auto generated low-level Go binding around a Klaytn contract.
type CountersRaw struct {
	Contract *Counters // Generic contract binding to access the raw methods on
}

// CountersCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type CountersCallerRaw struct {
	Contract *CountersCaller // Generic read-only contract binding to access the raw methods on
}

// CountersTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type CountersTransactorRaw struct {
	Contract *CountersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCounters creates a new instance of Counters, bound to a specific deployed contract.
func NewCounters(address common.Address, backend bind.ContractBackend) (*Counters, error) {
	contract, err := bindCounters(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Counters{CountersCaller: CountersCaller{contract: contract}, CountersTransactor: CountersTransactor{contract: contract}, CountersFilterer: CountersFilterer{contract: contract}}, nil
}

// NewCountersCaller creates a new read-only instance of Counters, bound to a specific deployed contract.
func NewCountersCaller(address common.Address, caller bind.ContractCaller) (*CountersCaller, error) {
	contract, err := bindCounters(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CountersCaller{contract: contract}, nil
}

// NewCountersTransactor creates a new write-only instance of Counters, bound to a specific deployed contract.
func NewCountersTransactor(address common.Address, transactor bind.ContractTransactor) (*CountersTransactor, error) {
	contract, err := bindCounters(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CountersTransactor{contract: contract}, nil
}

// NewCountersFilterer creates a new log filterer instance of Counters, bound to a specific deployed contract.
func NewCountersFilterer(address common.Address, filterer bind.ContractFilterer) (*CountersFilterer, error) {
	contract, err := bindCounters(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CountersFilterer{contract: contract}, nil
}

// bindCounters binds a generic wrapper to an already deployed contract.
func bindCounters(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CountersABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Counters *CountersRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Counters.Contract.CountersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Counters *CountersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counters.Contract.CountersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Counters *CountersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Counters.Contract.CountersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Counters *CountersCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Counters.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Counters *CountersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counters.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Counters *CountersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Counters.Contract.contract.Transact(opts, method, params...)
}

// IERC721ReceiverABI is the input ABI used to generate the binding from.
const IERC721ReceiverABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"operator\",\"type\":\"address\"},{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC721ReceiverBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const IERC721ReceiverBinRuntime = ``

// IERC721ReceiverFuncSigs maps the 4-byte function signature to its string representation.
var IERC721ReceiverFuncSigs = map[string]string{
	"150b7a02": "onERC721Received(address,address,uint256,bytes)",
}

// IERC721Receiver is an auto generated Go binding around a Klaytn contract.
type IERC721Receiver struct {
	IERC721ReceiverCaller     // Read-only binding to the contract
	IERC721ReceiverTransactor // Write-only binding to the contract
	IERC721ReceiverFilterer   // Log filterer for contract events
}

// IERC721ReceiverCaller is an auto generated read-only Go binding around a Klaytn contract.
type IERC721ReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721ReceiverTransactor is an auto generated write-only Go binding around a Klaytn contract.
type IERC721ReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721ReceiverFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type IERC721ReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721ReceiverSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type IERC721ReceiverSession struct {
	Contract     *IERC721Receiver  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC721ReceiverCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type IERC721ReceiverCallerSession struct {
	Contract *IERC721ReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IERC721ReceiverTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type IERC721ReceiverTransactorSession struct {
	Contract     *IERC721ReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IERC721ReceiverRaw is an auto generated low-level Go binding around a Klaytn contract.
type IERC721ReceiverRaw struct {
	Contract *IERC721Receiver // Generic contract binding to access the raw methods on
}

// IERC721ReceiverCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type IERC721ReceiverCallerRaw struct {
	Contract *IERC721ReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// IERC721ReceiverTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type IERC721ReceiverTransactorRaw struct {
	Contract *IERC721ReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC721Receiver creates a new instance of IERC721Receiver, bound to a specific deployed contract.
func NewIERC721Receiver(address common.Address, backend bind.ContractBackend) (*IERC721Receiver, error) {
	contract, err := bindIERC721Receiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC721Receiver{IERC721ReceiverCaller: IERC721ReceiverCaller{contract: contract}, IERC721ReceiverTransactor: IERC721ReceiverTransactor{contract: contract}, IERC721ReceiverFilterer: IERC721ReceiverFilterer{contract: contract}}, nil
}

// NewIERC721ReceiverCaller creates a new read-only instance of IERC721Receiver, bound to a specific deployed contract.
func NewIERC721ReceiverCaller(address common.Address, caller bind.ContractCaller) (*IERC721ReceiverCaller, error) {
	contract, err := bindIERC721Receiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721ReceiverCaller{contract: contract}, nil
}

// NewIERC721ReceiverTransactor creates a new write-only instance of IERC721Receiver, bound to a specific deployed contract.
func NewIERC721ReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC721ReceiverTransactor, error) {
	contract, err := bindIERC721Receiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721ReceiverTransactor{contract: contract}, nil
}

// NewIERC721ReceiverFilterer creates a new log filterer instance of IERC721Receiver, bound to a specific deployed contract.
func NewIERC721ReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC721ReceiverFilterer, error) {
	contract, err := bindIERC721Receiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC721ReceiverFilterer{contract: contract}, nil
}

// bindIERC721Receiver binds a generic wrapper to an already deployed contract.
func bindIERC721Receiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC721ReceiverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721Receiver *IERC721ReceiverRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC721Receiver.Contract.IERC721ReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721Receiver *IERC721ReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC721Receiver.Contract.IERC721ReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721Receiver *IERC721ReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC721Receiver.Contract.IERC721ReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721Receiver *IERC721ReceiverCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC721Receiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721Receiver *IERC721ReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC721Receiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721Receiver *IERC721ReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC721Receiver.Contract.contract.Transact(opts, method, params...)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_IERC721Receiver *IERC721ReceiverTransactor) OnERC721Received(opts *bind.TransactOpts, operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721Receiver.contract.Transact(opts, "onERC721Received", operator, from, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_IERC721Receiver *IERC721ReceiverSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721Receiver.Contract.OnERC721Received(&_IERC721Receiver.TransactOpts, operator, from, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_IERC721Receiver *IERC721ReceiverTransactorSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721Receiver.Contract.OnERC721Received(&_IERC721Receiver.TransactOpts, operator, from, tokenId, data)
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

// IKIP17ABI is the input ABI used to generate the binding from.
const IKIP17ABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"operator\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"operator\",\"type\":\"address\"},{\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"}]"

// IKIP17BinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const IKIP17BinRuntime = ``

// IKIP17FuncSigs maps the 4-byte function signature to its string representation.
var IKIP17FuncSigs = map[string]string{
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"081812fc": "getApproved(uint256)",
	"e985e9c5": "isApprovedForAll(address,address)",
	"6352211e": "ownerOf(uint256)",
	"42842e0e": "safeTransferFrom(address,address,uint256)",
	"b88d4fde": "safeTransferFrom(address,address,uint256,bytes)",
	"a22cb465": "setApprovalForAll(address,bool)",
	"01ffc9a7": "supportsInterface(bytes4)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// IKIP17 is an auto generated Go binding around a Klaytn contract.
type IKIP17 struct {
	IKIP17Caller     // Read-only binding to the contract
	IKIP17Transactor // Write-only binding to the contract
	IKIP17Filterer   // Log filterer for contract events
}

// IKIP17Caller is an auto generated read-only Go binding around a Klaytn contract.
type IKIP17Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IKIP17Transactor is an auto generated write-only Go binding around a Klaytn contract.
type IKIP17Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IKIP17Filterer is an auto generated log filtering Go binding around a Klaytn contract events.
type IKIP17Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IKIP17Session is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type IKIP17Session struct {
	Contract     *IKIP17           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IKIP17CallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type IKIP17CallerSession struct {
	Contract *IKIP17Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IKIP17TransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type IKIP17TransactorSession struct {
	Contract     *IKIP17Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IKIP17Raw is an auto generated low-level Go binding around a Klaytn contract.
type IKIP17Raw struct {
	Contract *IKIP17 // Generic contract binding to access the raw methods on
}

// IKIP17CallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type IKIP17CallerRaw struct {
	Contract *IKIP17Caller // Generic read-only contract binding to access the raw methods on
}

// IKIP17TransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type IKIP17TransactorRaw struct {
	Contract *IKIP17Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIKIP17 creates a new instance of IKIP17, bound to a specific deployed contract.
func NewIKIP17(address common.Address, backend bind.ContractBackend) (*IKIP17, error) {
	contract, err := bindIKIP17(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IKIP17{IKIP17Caller: IKIP17Caller{contract: contract}, IKIP17Transactor: IKIP17Transactor{contract: contract}, IKIP17Filterer: IKIP17Filterer{contract: contract}}, nil
}

// NewIKIP17Caller creates a new read-only instance of IKIP17, bound to a specific deployed contract.
func NewIKIP17Caller(address common.Address, caller bind.ContractCaller) (*IKIP17Caller, error) {
	contract, err := bindIKIP17(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IKIP17Caller{contract: contract}, nil
}

// NewIKIP17Transactor creates a new write-only instance of IKIP17, bound to a specific deployed contract.
func NewIKIP17Transactor(address common.Address, transactor bind.ContractTransactor) (*IKIP17Transactor, error) {
	contract, err := bindIKIP17(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IKIP17Transactor{contract: contract}, nil
}

// NewIKIP17Filterer creates a new log filterer instance of IKIP17, bound to a specific deployed contract.
func NewIKIP17Filterer(address common.Address, filterer bind.ContractFilterer) (*IKIP17Filterer, error) {
	contract, err := bindIKIP17(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IKIP17Filterer{contract: contract}, nil
}

// bindIKIP17 binds a generic wrapper to an already deployed contract.
func bindIKIP17(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IKIP17ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IKIP17 *IKIP17Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IKIP17.Contract.IKIP17Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IKIP17 *IKIP17Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IKIP17.Contract.IKIP17Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IKIP17 *IKIP17Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IKIP17.Contract.IKIP17Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IKIP17 *IKIP17CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IKIP17.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IKIP17 *IKIP17TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IKIP17.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IKIP17 *IKIP17TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IKIP17.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_IKIP17 *IKIP17Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IKIP17.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_IKIP17 *IKIP17Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IKIP17.Contract.BalanceOf(&_IKIP17.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_IKIP17 *IKIP17CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IKIP17.Contract.BalanceOf(&_IKIP17.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_IKIP17 *IKIP17Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IKIP17.contract.Call(opts, out, "getApproved", tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_IKIP17 *IKIP17Session) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _IKIP17.Contract.GetApproved(&_IKIP17.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_IKIP17 *IKIP17CallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _IKIP17.Contract.GetApproved(&_IKIP17.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IKIP17 *IKIP17Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IKIP17.contract.Call(opts, out, "isApprovedForAll", owner, operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IKIP17 *IKIP17Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _IKIP17.Contract.IsApprovedForAll(&_IKIP17.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IKIP17 *IKIP17CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _IKIP17.Contract.IsApprovedForAll(&_IKIP17.CallOpts, owner, operator)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_IKIP17 *IKIP17Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IKIP17.contract.Call(opts, out, "ownerOf", tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_IKIP17 *IKIP17Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _IKIP17.Contract.OwnerOf(&_IKIP17.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_IKIP17 *IKIP17CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _IKIP17.Contract.OwnerOf(&_IKIP17.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IKIP17 *IKIP17Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IKIP17.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IKIP17 *IKIP17Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IKIP17.Contract.SupportsInterface(&_IKIP17.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IKIP17 *IKIP17CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IKIP17.Contract.SupportsInterface(&_IKIP17.CallOpts, interfaceId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IKIP17 *IKIP17Transactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IKIP17.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IKIP17 *IKIP17Session) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IKIP17.Contract.Approve(&_IKIP17.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IKIP17 *IKIP17TransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IKIP17.Contract.Approve(&_IKIP17.TransactOpts, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IKIP17 *IKIP17Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IKIP17.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IKIP17 *IKIP17Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IKIP17.Contract.SafeTransferFrom(&_IKIP17.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IKIP17 *IKIP17TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IKIP17.Contract.SafeTransferFrom(&_IKIP17.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IKIP17 *IKIP17Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IKIP17.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IKIP17 *IKIP17Session) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IKIP17.Contract.SafeTransferFrom0(&_IKIP17.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IKIP17 *IKIP17TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IKIP17.Contract.SafeTransferFrom0(&_IKIP17.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IKIP17 *IKIP17Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IKIP17.contract.Transact(opts, "setApprovalForAll", operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IKIP17 *IKIP17Session) SetApprovalForAll(operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IKIP17.Contract.SetApprovalForAll(&_IKIP17.TransactOpts, operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IKIP17 *IKIP17TransactorSession) SetApprovalForAll(operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IKIP17.Contract.SetApprovalForAll(&_IKIP17.TransactOpts, operator, _approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IKIP17 *IKIP17Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IKIP17.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IKIP17 *IKIP17Session) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IKIP17.Contract.TransferFrom(&_IKIP17.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IKIP17 *IKIP17TransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IKIP17.Contract.TransferFrom(&_IKIP17.TransactOpts, from, to, tokenId)
}

// IKIP17ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IKIP17 contract.
type IKIP17ApprovalIterator struct {
	Event *IKIP17Approval // Event containing the contract specifics and raw log

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
func (it *IKIP17ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IKIP17Approval)
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
		it.Event = new(IKIP17Approval)
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
func (it *IKIP17ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IKIP17ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IKIP17Approval represents a Approval event raised by the IKIP17 contract.
type IKIP17Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IKIP17 *IKIP17Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*IKIP17ApprovalIterator, error) {

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

	logs, sub, err := _IKIP17.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IKIP17ApprovalIterator{contract: _IKIP17.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IKIP17 *IKIP17Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IKIP17Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _IKIP17.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IKIP17Approval)
				if err := _IKIP17.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IKIP17 *IKIP17Filterer) ParseApproval(log types.Log) (*IKIP17Approval, error) {
	event := new(IKIP17Approval)
	if err := _IKIP17.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IKIP17ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the IKIP17 contract.
type IKIP17ApprovalForAllIterator struct {
	Event *IKIP17ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *IKIP17ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IKIP17ApprovalForAll)
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
		it.Event = new(IKIP17ApprovalForAll)
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
func (it *IKIP17ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IKIP17ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IKIP17ApprovalForAll represents a ApprovalForAll event raised by the IKIP17 contract.
type IKIP17ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IKIP17 *IKIP17Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*IKIP17ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IKIP17.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &IKIP17ApprovalForAllIterator{contract: _IKIP17.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IKIP17 *IKIP17Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *IKIP17ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IKIP17.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IKIP17ApprovalForAll)
				if err := _IKIP17.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IKIP17 *IKIP17Filterer) ParseApprovalForAll(log types.Log) (*IKIP17ApprovalForAll, error) {
	event := new(IKIP17ApprovalForAll)
	if err := _IKIP17.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IKIP17TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IKIP17 contract.
type IKIP17TransferIterator struct {
	Event *IKIP17Transfer // Event containing the contract specifics and raw log

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
func (it *IKIP17TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IKIP17Transfer)
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
		it.Event = new(IKIP17Transfer)
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
func (it *IKIP17TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IKIP17TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IKIP17Transfer represents a Transfer event raised by the IKIP17 contract.
type IKIP17Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IKIP17 *IKIP17Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*IKIP17TransferIterator, error) {

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

	logs, sub, err := _IKIP17.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IKIP17TransferIterator{contract: _IKIP17.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IKIP17 *IKIP17Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IKIP17Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _IKIP17.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IKIP17Transfer)
				if err := _IKIP17.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IKIP17 *IKIP17Filterer) ParseTransfer(log types.Log) (*IKIP17Transfer, error) {
	event := new(IKIP17Transfer)
	if err := _IKIP17.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IKIP17BridgeReceiverABI is the input ABI used to generate the binding from.
const IKIP17BridgeReceiverABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"onKIP17Received\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IKIP17BridgeReceiverBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const IKIP17BridgeReceiverBinRuntime = ``

// IKIP17BridgeReceiverFuncSigs maps the 4-byte function signature to its string representation.
var IKIP17BridgeReceiverFuncSigs = map[string]string{
	"15c28177": "onKIP17Received(address,uint256,address,bytes)",
}

// IKIP17BridgeReceiver is an auto generated Go binding around a Klaytn contract.
type IKIP17BridgeReceiver struct {
	IKIP17BridgeReceiverCaller     // Read-only binding to the contract
	IKIP17BridgeReceiverTransactor // Write-only binding to the contract
	IKIP17BridgeReceiverFilterer   // Log filterer for contract events
}

// IKIP17BridgeReceiverCaller is an auto generated read-only Go binding around a Klaytn contract.
type IKIP17BridgeReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IKIP17BridgeReceiverTransactor is an auto generated write-only Go binding around a Klaytn contract.
type IKIP17BridgeReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IKIP17BridgeReceiverFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type IKIP17BridgeReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IKIP17BridgeReceiverSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type IKIP17BridgeReceiverSession struct {
	Contract     *IKIP17BridgeReceiver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IKIP17BridgeReceiverCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type IKIP17BridgeReceiverCallerSession struct {
	Contract *IKIP17BridgeReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// IKIP17BridgeReceiverTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type IKIP17BridgeReceiverTransactorSession struct {
	Contract     *IKIP17BridgeReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// IKIP17BridgeReceiverRaw is an auto generated low-level Go binding around a Klaytn contract.
type IKIP17BridgeReceiverRaw struct {
	Contract *IKIP17BridgeReceiver // Generic contract binding to access the raw methods on
}

// IKIP17BridgeReceiverCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type IKIP17BridgeReceiverCallerRaw struct {
	Contract *IKIP17BridgeReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// IKIP17BridgeReceiverTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type IKIP17BridgeReceiverTransactorRaw struct {
	Contract *IKIP17BridgeReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIKIP17BridgeReceiver creates a new instance of IKIP17BridgeReceiver, bound to a specific deployed contract.
func NewIKIP17BridgeReceiver(address common.Address, backend bind.ContractBackend) (*IKIP17BridgeReceiver, error) {
	contract, err := bindIKIP17BridgeReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IKIP17BridgeReceiver{IKIP17BridgeReceiverCaller: IKIP17BridgeReceiverCaller{contract: contract}, IKIP17BridgeReceiverTransactor: IKIP17BridgeReceiverTransactor{contract: contract}, IKIP17BridgeReceiverFilterer: IKIP17BridgeReceiverFilterer{contract: contract}}, nil
}

// NewIKIP17BridgeReceiverCaller creates a new read-only instance of IKIP17BridgeReceiver, bound to a specific deployed contract.
func NewIKIP17BridgeReceiverCaller(address common.Address, caller bind.ContractCaller) (*IKIP17BridgeReceiverCaller, error) {
	contract, err := bindIKIP17BridgeReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IKIP17BridgeReceiverCaller{contract: contract}, nil
}

// NewIKIP17BridgeReceiverTransactor creates a new write-only instance of IKIP17BridgeReceiver, bound to a specific deployed contract.
func NewIKIP17BridgeReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*IKIP17BridgeReceiverTransactor, error) {
	contract, err := bindIKIP17BridgeReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IKIP17BridgeReceiverTransactor{contract: contract}, nil
}

// NewIKIP17BridgeReceiverFilterer creates a new log filterer instance of IKIP17BridgeReceiver, bound to a specific deployed contract.
func NewIKIP17BridgeReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*IKIP17BridgeReceiverFilterer, error) {
	contract, err := bindIKIP17BridgeReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IKIP17BridgeReceiverFilterer{contract: contract}, nil
}

// bindIKIP17BridgeReceiver binds a generic wrapper to an already deployed contract.
func bindIKIP17BridgeReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IKIP17BridgeReceiverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IKIP17BridgeReceiver *IKIP17BridgeReceiverRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IKIP17BridgeReceiver.Contract.IKIP17BridgeReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IKIP17BridgeReceiver *IKIP17BridgeReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IKIP17BridgeReceiver.Contract.IKIP17BridgeReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IKIP17BridgeReceiver *IKIP17BridgeReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IKIP17BridgeReceiver.Contract.IKIP17BridgeReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IKIP17BridgeReceiver *IKIP17BridgeReceiverCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IKIP17BridgeReceiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IKIP17BridgeReceiver *IKIP17BridgeReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IKIP17BridgeReceiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IKIP17BridgeReceiver *IKIP17BridgeReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IKIP17BridgeReceiver.Contract.contract.Transact(opts, method, params...)
}

// OnKIP17Received is a paid mutator transaction binding the contract method 0x15c28177.
//
// Solidity: function onKIP17Received(address _from, uint256 _tokenId, address _to, bytes _extraData) returns()
func (_IKIP17BridgeReceiver *IKIP17BridgeReceiverTransactor) OnKIP17Received(opts *bind.TransactOpts, _from common.Address, _tokenId *big.Int, _to common.Address, _extraData []byte) (*types.Transaction, error) {
	return _IKIP17BridgeReceiver.contract.Transact(opts, "onKIP17Received", _from, _tokenId, _to, _extraData)
}

// OnKIP17Received is a paid mutator transaction binding the contract method 0x15c28177.
//
// Solidity: function onKIP17Received(address _from, uint256 _tokenId, address _to, bytes _extraData) returns()
func (_IKIP17BridgeReceiver *IKIP17BridgeReceiverSession) OnKIP17Received(_from common.Address, _tokenId *big.Int, _to common.Address, _extraData []byte) (*types.Transaction, error) {
	return _IKIP17BridgeReceiver.Contract.OnKIP17Received(&_IKIP17BridgeReceiver.TransactOpts, _from, _tokenId, _to, _extraData)
}

// OnKIP17Received is a paid mutator transaction binding the contract method 0x15c28177.
//
// Solidity: function onKIP17Received(address _from, uint256 _tokenId, address _to, bytes _extraData) returns()
func (_IKIP17BridgeReceiver *IKIP17BridgeReceiverTransactorSession) OnKIP17Received(_from common.Address, _tokenId *big.Int, _to common.Address, _extraData []byte) (*types.Transaction, error) {
	return _IKIP17BridgeReceiver.Contract.OnKIP17Received(&_IKIP17BridgeReceiver.TransactOpts, _from, _tokenId, _to, _extraData)
}

// IKIP17EnumerableABI is the input ABI used to generate the binding from.
const IKIP17EnumerableABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"operator\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"operator\",\"type\":\"address\"},{\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"}]"

// IKIP17EnumerableBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const IKIP17EnumerableBinRuntime = ``

// IKIP17EnumerableFuncSigs maps the 4-byte function signature to its string representation.
var IKIP17EnumerableFuncSigs = map[string]string{
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"081812fc": "getApproved(uint256)",
	"e985e9c5": "isApprovedForAll(address,address)",
	"6352211e": "ownerOf(uint256)",
	"42842e0e": "safeTransferFrom(address,address,uint256)",
	"b88d4fde": "safeTransferFrom(address,address,uint256,bytes)",
	"a22cb465": "setApprovalForAll(address,bool)",
	"01ffc9a7": "supportsInterface(bytes4)",
	"4f6ccce7": "tokenByIndex(uint256)",
	"2f745c59": "tokenOfOwnerByIndex(address,uint256)",
	"18160ddd": "totalSupply()",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// IKIP17Enumerable is an auto generated Go binding around a Klaytn contract.
type IKIP17Enumerable struct {
	IKIP17EnumerableCaller     // Read-only binding to the contract
	IKIP17EnumerableTransactor // Write-only binding to the contract
	IKIP17EnumerableFilterer   // Log filterer for contract events
}

// IKIP17EnumerableCaller is an auto generated read-only Go binding around a Klaytn contract.
type IKIP17EnumerableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IKIP17EnumerableTransactor is an auto generated write-only Go binding around a Klaytn contract.
type IKIP17EnumerableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IKIP17EnumerableFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type IKIP17EnumerableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IKIP17EnumerableSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type IKIP17EnumerableSession struct {
	Contract     *IKIP17Enumerable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IKIP17EnumerableCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type IKIP17EnumerableCallerSession struct {
	Contract *IKIP17EnumerableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IKIP17EnumerableTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type IKIP17EnumerableTransactorSession struct {
	Contract     *IKIP17EnumerableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IKIP17EnumerableRaw is an auto generated low-level Go binding around a Klaytn contract.
type IKIP17EnumerableRaw struct {
	Contract *IKIP17Enumerable // Generic contract binding to access the raw methods on
}

// IKIP17EnumerableCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type IKIP17EnumerableCallerRaw struct {
	Contract *IKIP17EnumerableCaller // Generic read-only contract binding to access the raw methods on
}

// IKIP17EnumerableTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type IKIP17EnumerableTransactorRaw struct {
	Contract *IKIP17EnumerableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIKIP17Enumerable creates a new instance of IKIP17Enumerable, bound to a specific deployed contract.
func NewIKIP17Enumerable(address common.Address, backend bind.ContractBackend) (*IKIP17Enumerable, error) {
	contract, err := bindIKIP17Enumerable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IKIP17Enumerable{IKIP17EnumerableCaller: IKIP17EnumerableCaller{contract: contract}, IKIP17EnumerableTransactor: IKIP17EnumerableTransactor{contract: contract}, IKIP17EnumerableFilterer: IKIP17EnumerableFilterer{contract: contract}}, nil
}

// NewIKIP17EnumerableCaller creates a new read-only instance of IKIP17Enumerable, bound to a specific deployed contract.
func NewIKIP17EnumerableCaller(address common.Address, caller bind.ContractCaller) (*IKIP17EnumerableCaller, error) {
	contract, err := bindIKIP17Enumerable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IKIP17EnumerableCaller{contract: contract}, nil
}

// NewIKIP17EnumerableTransactor creates a new write-only instance of IKIP17Enumerable, bound to a specific deployed contract.
func NewIKIP17EnumerableTransactor(address common.Address, transactor bind.ContractTransactor) (*IKIP17EnumerableTransactor, error) {
	contract, err := bindIKIP17Enumerable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IKIP17EnumerableTransactor{contract: contract}, nil
}

// NewIKIP17EnumerableFilterer creates a new log filterer instance of IKIP17Enumerable, bound to a specific deployed contract.
func NewIKIP17EnumerableFilterer(address common.Address, filterer bind.ContractFilterer) (*IKIP17EnumerableFilterer, error) {
	contract, err := bindIKIP17Enumerable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IKIP17EnumerableFilterer{contract: contract}, nil
}

// bindIKIP17Enumerable binds a generic wrapper to an already deployed contract.
func bindIKIP17Enumerable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IKIP17EnumerableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IKIP17Enumerable *IKIP17EnumerableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IKIP17Enumerable.Contract.IKIP17EnumerableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IKIP17Enumerable *IKIP17EnumerableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IKIP17Enumerable.Contract.IKIP17EnumerableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IKIP17Enumerable *IKIP17EnumerableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IKIP17Enumerable.Contract.IKIP17EnumerableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IKIP17Enumerable *IKIP17EnumerableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IKIP17Enumerable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IKIP17Enumerable *IKIP17EnumerableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IKIP17Enumerable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IKIP17Enumerable *IKIP17EnumerableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IKIP17Enumerable.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_IKIP17Enumerable *IKIP17EnumerableCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IKIP17Enumerable.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_IKIP17Enumerable *IKIP17EnumerableSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IKIP17Enumerable.Contract.BalanceOf(&_IKIP17Enumerable.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_IKIP17Enumerable *IKIP17EnumerableCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IKIP17Enumerable.Contract.BalanceOf(&_IKIP17Enumerable.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_IKIP17Enumerable *IKIP17EnumerableCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IKIP17Enumerable.contract.Call(opts, out, "getApproved", tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_IKIP17Enumerable *IKIP17EnumerableSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _IKIP17Enumerable.Contract.GetApproved(&_IKIP17Enumerable.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_IKIP17Enumerable *IKIP17EnumerableCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _IKIP17Enumerable.Contract.GetApproved(&_IKIP17Enumerable.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IKIP17Enumerable *IKIP17EnumerableCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IKIP17Enumerable.contract.Call(opts, out, "isApprovedForAll", owner, operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IKIP17Enumerable *IKIP17EnumerableSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _IKIP17Enumerable.Contract.IsApprovedForAll(&_IKIP17Enumerable.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IKIP17Enumerable *IKIP17EnumerableCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _IKIP17Enumerable.Contract.IsApprovedForAll(&_IKIP17Enumerable.CallOpts, owner, operator)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_IKIP17Enumerable *IKIP17EnumerableCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IKIP17Enumerable.contract.Call(opts, out, "ownerOf", tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_IKIP17Enumerable *IKIP17EnumerableSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _IKIP17Enumerable.Contract.OwnerOf(&_IKIP17Enumerable.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_IKIP17Enumerable *IKIP17EnumerableCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _IKIP17Enumerable.Contract.OwnerOf(&_IKIP17Enumerable.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IKIP17Enumerable *IKIP17EnumerableCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IKIP17Enumerable.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IKIP17Enumerable *IKIP17EnumerableSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IKIP17Enumerable.Contract.SupportsInterface(&_IKIP17Enumerable.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IKIP17Enumerable *IKIP17EnumerableCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IKIP17Enumerable.Contract.SupportsInterface(&_IKIP17Enumerable.CallOpts, interfaceId)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_IKIP17Enumerable *IKIP17EnumerableCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IKIP17Enumerable.contract.Call(opts, out, "tokenByIndex", index)
	return *ret0, err
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_IKIP17Enumerable *IKIP17EnumerableSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _IKIP17Enumerable.Contract.TokenByIndex(&_IKIP17Enumerable.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_IKIP17Enumerable *IKIP17EnumerableCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _IKIP17Enumerable.Contract.TokenByIndex(&_IKIP17Enumerable.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256 tokenId)
func (_IKIP17Enumerable *IKIP17EnumerableCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IKIP17Enumerable.contract.Call(opts, out, "tokenOfOwnerByIndex", owner, index)
	return *ret0, err
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256 tokenId)
func (_IKIP17Enumerable *IKIP17EnumerableSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _IKIP17Enumerable.Contract.TokenOfOwnerByIndex(&_IKIP17Enumerable.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256 tokenId)
func (_IKIP17Enumerable *IKIP17EnumerableCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _IKIP17Enumerable.Contract.TokenOfOwnerByIndex(&_IKIP17Enumerable.CallOpts, owner, index)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IKIP17Enumerable *IKIP17EnumerableCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IKIP17Enumerable.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IKIP17Enumerable *IKIP17EnumerableSession) TotalSupply() (*big.Int, error) {
	return _IKIP17Enumerable.Contract.TotalSupply(&_IKIP17Enumerable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IKIP17Enumerable *IKIP17EnumerableCallerSession) TotalSupply() (*big.Int, error) {
	return _IKIP17Enumerable.Contract.TotalSupply(&_IKIP17Enumerable.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IKIP17Enumerable *IKIP17EnumerableTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IKIP17Enumerable.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IKIP17Enumerable *IKIP17EnumerableSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IKIP17Enumerable.Contract.Approve(&_IKIP17Enumerable.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IKIP17Enumerable *IKIP17EnumerableTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IKIP17Enumerable.Contract.Approve(&_IKIP17Enumerable.TransactOpts, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IKIP17Enumerable *IKIP17EnumerableTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IKIP17Enumerable.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IKIP17Enumerable *IKIP17EnumerableSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IKIP17Enumerable.Contract.SafeTransferFrom(&_IKIP17Enumerable.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IKIP17Enumerable *IKIP17EnumerableTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IKIP17Enumerable.Contract.SafeTransferFrom(&_IKIP17Enumerable.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IKIP17Enumerable *IKIP17EnumerableTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IKIP17Enumerable.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IKIP17Enumerable *IKIP17EnumerableSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IKIP17Enumerable.Contract.SafeTransferFrom0(&_IKIP17Enumerable.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IKIP17Enumerable *IKIP17EnumerableTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IKIP17Enumerable.Contract.SafeTransferFrom0(&_IKIP17Enumerable.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IKIP17Enumerable *IKIP17EnumerableTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IKIP17Enumerable.contract.Transact(opts, "setApprovalForAll", operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IKIP17Enumerable *IKIP17EnumerableSession) SetApprovalForAll(operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IKIP17Enumerable.Contract.SetApprovalForAll(&_IKIP17Enumerable.TransactOpts, operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IKIP17Enumerable *IKIP17EnumerableTransactorSession) SetApprovalForAll(operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IKIP17Enumerable.Contract.SetApprovalForAll(&_IKIP17Enumerable.TransactOpts, operator, _approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IKIP17Enumerable *IKIP17EnumerableTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IKIP17Enumerable.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IKIP17Enumerable *IKIP17EnumerableSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IKIP17Enumerable.Contract.TransferFrom(&_IKIP17Enumerable.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IKIP17Enumerable *IKIP17EnumerableTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IKIP17Enumerable.Contract.TransferFrom(&_IKIP17Enumerable.TransactOpts, from, to, tokenId)
}

// IKIP17EnumerableApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IKIP17Enumerable contract.
type IKIP17EnumerableApprovalIterator struct {
	Event *IKIP17EnumerableApproval // Event containing the contract specifics and raw log

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
func (it *IKIP17EnumerableApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IKIP17EnumerableApproval)
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
		it.Event = new(IKIP17EnumerableApproval)
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
func (it *IKIP17EnumerableApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IKIP17EnumerableApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IKIP17EnumerableApproval represents a Approval event raised by the IKIP17Enumerable contract.
type IKIP17EnumerableApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IKIP17Enumerable *IKIP17EnumerableFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*IKIP17EnumerableApprovalIterator, error) {

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

	logs, sub, err := _IKIP17Enumerable.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IKIP17EnumerableApprovalIterator{contract: _IKIP17Enumerable.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IKIP17Enumerable *IKIP17EnumerableFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IKIP17EnumerableApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _IKIP17Enumerable.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IKIP17EnumerableApproval)
				if err := _IKIP17Enumerable.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IKIP17Enumerable *IKIP17EnumerableFilterer) ParseApproval(log types.Log) (*IKIP17EnumerableApproval, error) {
	event := new(IKIP17EnumerableApproval)
	if err := _IKIP17Enumerable.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IKIP17EnumerableApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the IKIP17Enumerable contract.
type IKIP17EnumerableApprovalForAllIterator struct {
	Event *IKIP17EnumerableApprovalForAll // Event containing the contract specifics and raw log

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
func (it *IKIP17EnumerableApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IKIP17EnumerableApprovalForAll)
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
		it.Event = new(IKIP17EnumerableApprovalForAll)
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
func (it *IKIP17EnumerableApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IKIP17EnumerableApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IKIP17EnumerableApprovalForAll represents a ApprovalForAll event raised by the IKIP17Enumerable contract.
type IKIP17EnumerableApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IKIP17Enumerable *IKIP17EnumerableFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*IKIP17EnumerableApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IKIP17Enumerable.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &IKIP17EnumerableApprovalForAllIterator{contract: _IKIP17Enumerable.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IKIP17Enumerable *IKIP17EnumerableFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *IKIP17EnumerableApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IKIP17Enumerable.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IKIP17EnumerableApprovalForAll)
				if err := _IKIP17Enumerable.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IKIP17Enumerable *IKIP17EnumerableFilterer) ParseApprovalForAll(log types.Log) (*IKIP17EnumerableApprovalForAll, error) {
	event := new(IKIP17EnumerableApprovalForAll)
	if err := _IKIP17Enumerable.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IKIP17EnumerableTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IKIP17Enumerable contract.
type IKIP17EnumerableTransferIterator struct {
	Event *IKIP17EnumerableTransfer // Event containing the contract specifics and raw log

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
func (it *IKIP17EnumerableTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IKIP17EnumerableTransfer)
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
		it.Event = new(IKIP17EnumerableTransfer)
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
func (it *IKIP17EnumerableTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IKIP17EnumerableTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IKIP17EnumerableTransfer represents a Transfer event raised by the IKIP17Enumerable contract.
type IKIP17EnumerableTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IKIP17Enumerable *IKIP17EnumerableFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*IKIP17EnumerableTransferIterator, error) {

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

	logs, sub, err := _IKIP17Enumerable.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IKIP17EnumerableTransferIterator{contract: _IKIP17Enumerable.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IKIP17Enumerable *IKIP17EnumerableFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IKIP17EnumerableTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _IKIP17Enumerable.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IKIP17EnumerableTransfer)
				if err := _IKIP17Enumerable.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IKIP17Enumerable *IKIP17EnumerableFilterer) ParseTransfer(log types.Log) (*IKIP17EnumerableTransfer, error) {
	event := new(IKIP17EnumerableTransfer)
	if err := _IKIP17Enumerable.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IKIP17MetadataABI is the input ABI used to generate the binding from.
const IKIP17MetadataABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"operator\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"operator\",\"type\":\"address\"},{\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"}]"

// IKIP17MetadataBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const IKIP17MetadataBinRuntime = ``

// IKIP17MetadataFuncSigs maps the 4-byte function signature to its string representation.
var IKIP17MetadataFuncSigs = map[string]string{
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"081812fc": "getApproved(uint256)",
	"e985e9c5": "isApprovedForAll(address,address)",
	"06fdde03": "name()",
	"6352211e": "ownerOf(uint256)",
	"42842e0e": "safeTransferFrom(address,address,uint256)",
	"b88d4fde": "safeTransferFrom(address,address,uint256,bytes)",
	"a22cb465": "setApprovalForAll(address,bool)",
	"01ffc9a7": "supportsInterface(bytes4)",
	"95d89b41": "symbol()",
	"c87b56dd": "tokenURI(uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// IKIP17Metadata is an auto generated Go binding around a Klaytn contract.
type IKIP17Metadata struct {
	IKIP17MetadataCaller     // Read-only binding to the contract
	IKIP17MetadataTransactor // Write-only binding to the contract
	IKIP17MetadataFilterer   // Log filterer for contract events
}

// IKIP17MetadataCaller is an auto generated read-only Go binding around a Klaytn contract.
type IKIP17MetadataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IKIP17MetadataTransactor is an auto generated write-only Go binding around a Klaytn contract.
type IKIP17MetadataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IKIP17MetadataFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type IKIP17MetadataFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IKIP17MetadataSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type IKIP17MetadataSession struct {
	Contract     *IKIP17Metadata   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IKIP17MetadataCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type IKIP17MetadataCallerSession struct {
	Contract *IKIP17MetadataCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IKIP17MetadataTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type IKIP17MetadataTransactorSession struct {
	Contract     *IKIP17MetadataTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IKIP17MetadataRaw is an auto generated low-level Go binding around a Klaytn contract.
type IKIP17MetadataRaw struct {
	Contract *IKIP17Metadata // Generic contract binding to access the raw methods on
}

// IKIP17MetadataCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type IKIP17MetadataCallerRaw struct {
	Contract *IKIP17MetadataCaller // Generic read-only contract binding to access the raw methods on
}

// IKIP17MetadataTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type IKIP17MetadataTransactorRaw struct {
	Contract *IKIP17MetadataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIKIP17Metadata creates a new instance of IKIP17Metadata, bound to a specific deployed contract.
func NewIKIP17Metadata(address common.Address, backend bind.ContractBackend) (*IKIP17Metadata, error) {
	contract, err := bindIKIP17Metadata(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IKIP17Metadata{IKIP17MetadataCaller: IKIP17MetadataCaller{contract: contract}, IKIP17MetadataTransactor: IKIP17MetadataTransactor{contract: contract}, IKIP17MetadataFilterer: IKIP17MetadataFilterer{contract: contract}}, nil
}

// NewIKIP17MetadataCaller creates a new read-only instance of IKIP17Metadata, bound to a specific deployed contract.
func NewIKIP17MetadataCaller(address common.Address, caller bind.ContractCaller) (*IKIP17MetadataCaller, error) {
	contract, err := bindIKIP17Metadata(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IKIP17MetadataCaller{contract: contract}, nil
}

// NewIKIP17MetadataTransactor creates a new write-only instance of IKIP17Metadata, bound to a specific deployed contract.
func NewIKIP17MetadataTransactor(address common.Address, transactor bind.ContractTransactor) (*IKIP17MetadataTransactor, error) {
	contract, err := bindIKIP17Metadata(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IKIP17MetadataTransactor{contract: contract}, nil
}

// NewIKIP17MetadataFilterer creates a new log filterer instance of IKIP17Metadata, bound to a specific deployed contract.
func NewIKIP17MetadataFilterer(address common.Address, filterer bind.ContractFilterer) (*IKIP17MetadataFilterer, error) {
	contract, err := bindIKIP17Metadata(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IKIP17MetadataFilterer{contract: contract}, nil
}

// bindIKIP17Metadata binds a generic wrapper to an already deployed contract.
func bindIKIP17Metadata(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IKIP17MetadataABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IKIP17Metadata *IKIP17MetadataRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IKIP17Metadata.Contract.IKIP17MetadataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IKIP17Metadata *IKIP17MetadataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IKIP17Metadata.Contract.IKIP17MetadataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IKIP17Metadata *IKIP17MetadataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IKIP17Metadata.Contract.IKIP17MetadataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IKIP17Metadata *IKIP17MetadataCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IKIP17Metadata.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IKIP17Metadata *IKIP17MetadataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IKIP17Metadata.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IKIP17Metadata *IKIP17MetadataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IKIP17Metadata.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_IKIP17Metadata *IKIP17MetadataCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IKIP17Metadata.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_IKIP17Metadata *IKIP17MetadataSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IKIP17Metadata.Contract.BalanceOf(&_IKIP17Metadata.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_IKIP17Metadata *IKIP17MetadataCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IKIP17Metadata.Contract.BalanceOf(&_IKIP17Metadata.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_IKIP17Metadata *IKIP17MetadataCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IKIP17Metadata.contract.Call(opts, out, "getApproved", tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_IKIP17Metadata *IKIP17MetadataSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _IKIP17Metadata.Contract.GetApproved(&_IKIP17Metadata.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_IKIP17Metadata *IKIP17MetadataCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _IKIP17Metadata.Contract.GetApproved(&_IKIP17Metadata.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IKIP17Metadata *IKIP17MetadataCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IKIP17Metadata.contract.Call(opts, out, "isApprovedForAll", owner, operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IKIP17Metadata *IKIP17MetadataSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _IKIP17Metadata.Contract.IsApprovedForAll(&_IKIP17Metadata.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IKIP17Metadata *IKIP17MetadataCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _IKIP17Metadata.Contract.IsApprovedForAll(&_IKIP17Metadata.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IKIP17Metadata *IKIP17MetadataCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _IKIP17Metadata.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IKIP17Metadata *IKIP17MetadataSession) Name() (string, error) {
	return _IKIP17Metadata.Contract.Name(&_IKIP17Metadata.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IKIP17Metadata *IKIP17MetadataCallerSession) Name() (string, error) {
	return _IKIP17Metadata.Contract.Name(&_IKIP17Metadata.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_IKIP17Metadata *IKIP17MetadataCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IKIP17Metadata.contract.Call(opts, out, "ownerOf", tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_IKIP17Metadata *IKIP17MetadataSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _IKIP17Metadata.Contract.OwnerOf(&_IKIP17Metadata.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_IKIP17Metadata *IKIP17MetadataCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _IKIP17Metadata.Contract.OwnerOf(&_IKIP17Metadata.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IKIP17Metadata *IKIP17MetadataCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IKIP17Metadata.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IKIP17Metadata *IKIP17MetadataSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IKIP17Metadata.Contract.SupportsInterface(&_IKIP17Metadata.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IKIP17Metadata *IKIP17MetadataCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IKIP17Metadata.Contract.SupportsInterface(&_IKIP17Metadata.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IKIP17Metadata *IKIP17MetadataCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _IKIP17Metadata.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IKIP17Metadata *IKIP17MetadataSession) Symbol() (string, error) {
	return _IKIP17Metadata.Contract.Symbol(&_IKIP17Metadata.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IKIP17Metadata *IKIP17MetadataCallerSession) Symbol() (string, error) {
	return _IKIP17Metadata.Contract.Symbol(&_IKIP17Metadata.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_IKIP17Metadata *IKIP17MetadataCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _IKIP17Metadata.contract.Call(opts, out, "tokenURI", tokenId)
	return *ret0, err
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_IKIP17Metadata *IKIP17MetadataSession) TokenURI(tokenId *big.Int) (string, error) {
	return _IKIP17Metadata.Contract.TokenURI(&_IKIP17Metadata.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_IKIP17Metadata *IKIP17MetadataCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _IKIP17Metadata.Contract.TokenURI(&_IKIP17Metadata.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IKIP17Metadata *IKIP17MetadataTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IKIP17Metadata.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IKIP17Metadata *IKIP17MetadataSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IKIP17Metadata.Contract.Approve(&_IKIP17Metadata.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IKIP17Metadata *IKIP17MetadataTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IKIP17Metadata.Contract.Approve(&_IKIP17Metadata.TransactOpts, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IKIP17Metadata *IKIP17MetadataTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IKIP17Metadata.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IKIP17Metadata *IKIP17MetadataSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IKIP17Metadata.Contract.SafeTransferFrom(&_IKIP17Metadata.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IKIP17Metadata *IKIP17MetadataTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IKIP17Metadata.Contract.SafeTransferFrom(&_IKIP17Metadata.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IKIP17Metadata *IKIP17MetadataTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IKIP17Metadata.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IKIP17Metadata *IKIP17MetadataSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IKIP17Metadata.Contract.SafeTransferFrom0(&_IKIP17Metadata.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IKIP17Metadata *IKIP17MetadataTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IKIP17Metadata.Contract.SafeTransferFrom0(&_IKIP17Metadata.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IKIP17Metadata *IKIP17MetadataTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IKIP17Metadata.contract.Transact(opts, "setApprovalForAll", operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IKIP17Metadata *IKIP17MetadataSession) SetApprovalForAll(operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IKIP17Metadata.Contract.SetApprovalForAll(&_IKIP17Metadata.TransactOpts, operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IKIP17Metadata *IKIP17MetadataTransactorSession) SetApprovalForAll(operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IKIP17Metadata.Contract.SetApprovalForAll(&_IKIP17Metadata.TransactOpts, operator, _approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IKIP17Metadata *IKIP17MetadataTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IKIP17Metadata.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IKIP17Metadata *IKIP17MetadataSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IKIP17Metadata.Contract.TransferFrom(&_IKIP17Metadata.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IKIP17Metadata *IKIP17MetadataTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IKIP17Metadata.Contract.TransferFrom(&_IKIP17Metadata.TransactOpts, from, to, tokenId)
}

// IKIP17MetadataApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IKIP17Metadata contract.
type IKIP17MetadataApprovalIterator struct {
	Event *IKIP17MetadataApproval // Event containing the contract specifics and raw log

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
func (it *IKIP17MetadataApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IKIP17MetadataApproval)
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
		it.Event = new(IKIP17MetadataApproval)
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
func (it *IKIP17MetadataApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IKIP17MetadataApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IKIP17MetadataApproval represents a Approval event raised by the IKIP17Metadata contract.
type IKIP17MetadataApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IKIP17Metadata *IKIP17MetadataFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*IKIP17MetadataApprovalIterator, error) {

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

	logs, sub, err := _IKIP17Metadata.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IKIP17MetadataApprovalIterator{contract: _IKIP17Metadata.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IKIP17Metadata *IKIP17MetadataFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IKIP17MetadataApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _IKIP17Metadata.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IKIP17MetadataApproval)
				if err := _IKIP17Metadata.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IKIP17Metadata *IKIP17MetadataFilterer) ParseApproval(log types.Log) (*IKIP17MetadataApproval, error) {
	event := new(IKIP17MetadataApproval)
	if err := _IKIP17Metadata.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IKIP17MetadataApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the IKIP17Metadata contract.
type IKIP17MetadataApprovalForAllIterator struct {
	Event *IKIP17MetadataApprovalForAll // Event containing the contract specifics and raw log

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
func (it *IKIP17MetadataApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IKIP17MetadataApprovalForAll)
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
		it.Event = new(IKIP17MetadataApprovalForAll)
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
func (it *IKIP17MetadataApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IKIP17MetadataApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IKIP17MetadataApprovalForAll represents a ApprovalForAll event raised by the IKIP17Metadata contract.
type IKIP17MetadataApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IKIP17Metadata *IKIP17MetadataFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*IKIP17MetadataApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IKIP17Metadata.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &IKIP17MetadataApprovalForAllIterator{contract: _IKIP17Metadata.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IKIP17Metadata *IKIP17MetadataFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *IKIP17MetadataApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IKIP17Metadata.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IKIP17MetadataApprovalForAll)
				if err := _IKIP17Metadata.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IKIP17Metadata *IKIP17MetadataFilterer) ParseApprovalForAll(log types.Log) (*IKIP17MetadataApprovalForAll, error) {
	event := new(IKIP17MetadataApprovalForAll)
	if err := _IKIP17Metadata.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IKIP17MetadataTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IKIP17Metadata contract.
type IKIP17MetadataTransferIterator struct {
	Event *IKIP17MetadataTransfer // Event containing the contract specifics and raw log

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
func (it *IKIP17MetadataTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IKIP17MetadataTransfer)
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
		it.Event = new(IKIP17MetadataTransfer)
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
func (it *IKIP17MetadataTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IKIP17MetadataTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IKIP17MetadataTransfer represents a Transfer event raised by the IKIP17Metadata contract.
type IKIP17MetadataTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IKIP17Metadata *IKIP17MetadataFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*IKIP17MetadataTransferIterator, error) {

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

	logs, sub, err := _IKIP17Metadata.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IKIP17MetadataTransferIterator{contract: _IKIP17Metadata.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IKIP17Metadata *IKIP17MetadataFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IKIP17MetadataTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _IKIP17Metadata.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IKIP17MetadataTransfer)
				if err := _IKIP17Metadata.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IKIP17Metadata *IKIP17MetadataFilterer) ParseTransfer(log types.Log) (*IKIP17MetadataTransfer, error) {
	event := new(IKIP17MetadataTransfer)
	if err := _IKIP17Metadata.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IKIP17ReceiverABI is the input ABI used to generate the binding from.
const IKIP17ReceiverABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"operator\",\"type\":\"address\"},{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onKIP17Received\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IKIP17ReceiverBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const IKIP17ReceiverBinRuntime = ``

// IKIP17ReceiverFuncSigs maps the 4-byte function signature to its string representation.
var IKIP17ReceiverFuncSigs = map[string]string{
	"6745782b": "onKIP17Received(address,address,uint256,bytes)",
}

// IKIP17Receiver is an auto generated Go binding around a Klaytn contract.
type IKIP17Receiver struct {
	IKIP17ReceiverCaller     // Read-only binding to the contract
	IKIP17ReceiverTransactor // Write-only binding to the contract
	IKIP17ReceiverFilterer   // Log filterer for contract events
}

// IKIP17ReceiverCaller is an auto generated read-only Go binding around a Klaytn contract.
type IKIP17ReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IKIP17ReceiverTransactor is an auto generated write-only Go binding around a Klaytn contract.
type IKIP17ReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IKIP17ReceiverFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type IKIP17ReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IKIP17ReceiverSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type IKIP17ReceiverSession struct {
	Contract     *IKIP17Receiver   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IKIP17ReceiverCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type IKIP17ReceiverCallerSession struct {
	Contract *IKIP17ReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IKIP17ReceiverTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type IKIP17ReceiverTransactorSession struct {
	Contract     *IKIP17ReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IKIP17ReceiverRaw is an auto generated low-level Go binding around a Klaytn contract.
type IKIP17ReceiverRaw struct {
	Contract *IKIP17Receiver // Generic contract binding to access the raw methods on
}

// IKIP17ReceiverCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type IKIP17ReceiverCallerRaw struct {
	Contract *IKIP17ReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// IKIP17ReceiverTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type IKIP17ReceiverTransactorRaw struct {
	Contract *IKIP17ReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIKIP17Receiver creates a new instance of IKIP17Receiver, bound to a specific deployed contract.
func NewIKIP17Receiver(address common.Address, backend bind.ContractBackend) (*IKIP17Receiver, error) {
	contract, err := bindIKIP17Receiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IKIP17Receiver{IKIP17ReceiverCaller: IKIP17ReceiverCaller{contract: contract}, IKIP17ReceiverTransactor: IKIP17ReceiverTransactor{contract: contract}, IKIP17ReceiverFilterer: IKIP17ReceiverFilterer{contract: contract}}, nil
}

// NewIKIP17ReceiverCaller creates a new read-only instance of IKIP17Receiver, bound to a specific deployed contract.
func NewIKIP17ReceiverCaller(address common.Address, caller bind.ContractCaller) (*IKIP17ReceiverCaller, error) {
	contract, err := bindIKIP17Receiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IKIP17ReceiverCaller{contract: contract}, nil
}

// NewIKIP17ReceiverTransactor creates a new write-only instance of IKIP17Receiver, bound to a specific deployed contract.
func NewIKIP17ReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*IKIP17ReceiverTransactor, error) {
	contract, err := bindIKIP17Receiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IKIP17ReceiverTransactor{contract: contract}, nil
}

// NewIKIP17ReceiverFilterer creates a new log filterer instance of IKIP17Receiver, bound to a specific deployed contract.
func NewIKIP17ReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*IKIP17ReceiverFilterer, error) {
	contract, err := bindIKIP17Receiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IKIP17ReceiverFilterer{contract: contract}, nil
}

// bindIKIP17Receiver binds a generic wrapper to an already deployed contract.
func bindIKIP17Receiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IKIP17ReceiverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IKIP17Receiver *IKIP17ReceiverRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IKIP17Receiver.Contract.IKIP17ReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IKIP17Receiver *IKIP17ReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IKIP17Receiver.Contract.IKIP17ReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IKIP17Receiver *IKIP17ReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IKIP17Receiver.Contract.IKIP17ReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IKIP17Receiver *IKIP17ReceiverCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IKIP17Receiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IKIP17Receiver *IKIP17ReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IKIP17Receiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IKIP17Receiver *IKIP17ReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IKIP17Receiver.Contract.contract.Transact(opts, method, params...)
}

// OnKIP17Received is a paid mutator transaction binding the contract method 0x6745782b.
//
// Solidity: function onKIP17Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_IKIP17Receiver *IKIP17ReceiverTransactor) OnKIP17Received(opts *bind.TransactOpts, operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IKIP17Receiver.contract.Transact(opts, "onKIP17Received", operator, from, tokenId, data)
}

// OnKIP17Received is a paid mutator transaction binding the contract method 0x6745782b.
//
// Solidity: function onKIP17Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_IKIP17Receiver *IKIP17ReceiverSession) OnKIP17Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IKIP17Receiver.Contract.OnKIP17Received(&_IKIP17Receiver.TransactOpts, operator, from, tokenId, data)
}

// OnKIP17Received is a paid mutator transaction binding the contract method 0x6745782b.
//
// Solidity: function onKIP17Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_IKIP17Receiver *IKIP17ReceiverTransactorSession) OnKIP17Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IKIP17Receiver.Contract.OnKIP17Received(&_IKIP17Receiver.TransactOpts, operator, from, tokenId, data)
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

// KIP17ABI is the input ABI used to generate the binding from.
const KIP17ABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"}]"

// KIP17BinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const KIP17BinRuntime = `608060405234801561001057600080fd5b506004361061009e5760003560e01c80636352211e116100665780636352211e146101b157806370a08231146101ce578063a22cb46514610206578063b88d4fde14610234578063e985e9c5146102fa5761009e565b806301ffc9a7146100a3578063081812fc146100de578063095ea7b31461011757806323b872dd1461014557806342842e0e1461017b575b600080fd5b6100ca600480360360208110156100b957600080fd5b50356001600160e01b031916610328565b604080519115158252519081900360200190f35b6100fb600480360360208110156100f457600080fd5b5035610347565b604080516001600160a01b039092168252519081900360200190f35b6101436004803603604081101561012d57600080fd5b506001600160a01b0381351690602001356103ac565b005b6101436004803603606081101561015b57600080fd5b506001600160a01b038135811691602081013590911690604001356104d9565b6101436004803603606081101561019157600080fd5b506001600160a01b03813581169160208101359091169060400135610531565b6100fb600480360360208110156101c757600080fd5b503561054c565b6101f4600480360360208110156101e457600080fd5b50356001600160a01b03166105a9565b60408051918252519081900360200190f35b6101436004803603604081101561021c57600080fd5b506001600160a01b0381351690602001351515610614565b6101436004803603608081101561024a57600080fd5b6001600160a01b0382358116926020810135909116916040820135919081019060808101606082013564010000000081111561028557600080fd5b82018360208201111561029757600080fd5b803590602001918460018302840111640100000000831117156102b957600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506106e3945050505050565b6100ca6004803603604081101561031057600080fd5b506001600160a01b038135811691602001351661073e565b6001600160e01b03191660009081526020819052604090205460ff1690565b60006103528261076c565b61039057604051600160e51b62461bcd02815260040180806020018281038252602b815260200180610f21602b913960400191505060405180910390fd5b506000908152600260205260409020546001600160a01b031690565b60006103b78261054c565b9050806001600160a01b0316836001600160a01b031614156104235760408051600160e51b62461bcd02815260206004820181905260248201527f4b495031373a20617070726f76616c20746f2063757272656e74206f776e6572604482015290519081900360640190fd5b336001600160a01b038216148061043f575061043f813361073e565b61047d57604051600160e51b62461bcd028152600401808060200182810382526037815260200180610f4c6037913960400191505060405180910390fd5b60008281526002602052604080822080546001600160a01b0319166001600160a01b0387811691821790925591518593918516917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591a4505050565b6104e33382610789565b61052157604051600160e51b62461bcd028152600401808060200182810382526030815260200180610e706030913960400191505060405180910390fd5b61052c838383610830565b505050565b61052c838383604051806020016040528060008152506106e3565b6000818152600160205260408120546001600160a01b0316806105a357604051600160e51b62461bcd028152600401808060200182810382526028815260200180610e486028913960400191505060405180910390fd5b92915050565b60006001600160a01b0382166105f357604051600160e51b62461bcd028152600401808060200182810382526029815260200180610ed06029913960400191505060405180910390fd5b6001600160a01b03821660009081526003602052604090206105a39061097a565b6001600160a01b0382163314156106755760408051600160e51b62461bcd02815260206004820152601860248201527f4b495031373a20617070726f766520746f2063616c6c65720000000000000000604482015290519081900360640190fd5b3360008181526004602090815260408083206001600160a01b03871680855290835292819020805460ff1916861515908117909155815190815290519293927f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31929181900390910190a35050565b6106ee8484846104d9565b6106fa8484848461097e565b61073857604051600160e51b62461bcd028152600401808060200182810382526030815260200180610ea06030913960400191505060405180910390fd5b50505050565b6001600160a01b03918216600090815260046020908152604080832093909416825291909152205460ff1690565b6000908152600160205260409020546001600160a01b0316151590565b60006107948261076c565b6107d257604051600160e51b62461bcd02815260040180806020018281038252602b815260200180610f83602b913960400191505060405180910390fd5b60006107dd8361054c565b9050806001600160a01b0316846001600160a01b031614806108185750836001600160a01b031661080d84610347565b6001600160a01b0316145b806108285750610828818561073e565b949350505050565b826001600160a01b03166108438261054c565b6001600160a01b03161461088b57604051600160e51b62461bcd028152600401808060200182810382526028815260200180610ef96028913960400191505060405180910390fd5b6001600160a01b0382166108d357604051600160e51b62461bcd028152600401808060200182810382526023815260200180610e256023913960400191505060405180910390fd5b6108dc81610d61565b6001600160a01b03831660009081526003602052604090206108fd90610d9e565b6001600160a01b038216600090815260036020526040902061091e90610db5565b60008181526001602052604080822080546001600160a01b0319166001600160a01b0386811691821790925591518493918716917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b5490565b6000806060610995866001600160a01b0316610dbe565b6109a457600192505050610828565b856001600160a01b031663150b7a0260e01b3389888860405160240180856001600160a01b03166001600160a01b03168152602001846001600160a01b03166001600160a01b0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610a30578181015183820152602001610a18565b50505050905090810190601f168015610a5d5780820380516001836020036101000a031916815260200191505b5060408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909a16999099178952518151919890975087965094509250829150849050835b60208310610ac55780518252601f199092019160209182019101610aa6565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114610b27576040519150601f19603f3d011682016040523d82523d6000602084013e610b2c565b606091505b508051919350915015801590610b6c57508051600160e11b630a85bd01029060208084019190811015610b5e57600080fd5b50516001600160e01b031916145b15610b7c57600192505050610828565b856001600160a01b0316636745782b60e01b3389888860405160240180856001600160a01b03166001600160a01b03168152602001846001600160a01b03166001600160a01b0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610c08578181015183820152602001610bf0565b50505050905090810190601f168015610c355780820380516001836020036101000a031916815260200191505b5060408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909a16999099178952518151919890975087965094509250829150849050835b60208310610c9d5780518252601f199092019160209182019101610c7e565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114610cff576040519150601f19603f3d011682016040523d82523d6000602084013e610d04565b606091505b508051919350915015801590610d4457508051600160e01b636745782b029060208084019190811015610d3657600080fd5b50516001600160e01b031916145b15610d5457600192505050610828565b5060009695505050505050565b6000818152600260205260409020546001600160a01b031615610d9b57600081815260026020526040902080546001600160a01b03191690555b50565b8054610db190600163ffffffff610dc416565b9055565b80546001019055565b3b151590565b600082821115610e1e5760408051600160e51b62461bcd02815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b5090039056fe4b495031373a207472616e7366657220746f20746865207a65726f20616464726573734b495031373a206f776e657220717565727920666f72206e6f6e6578697374656e7420746f6b656e4b495031373a207472616e736665722063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f7665644b495031373a207472616e7366657220746f206e6f6e204b49503137526563656976657220696d706c656d656e7465724b495031373a2062616c616e636520717565727920666f7220746865207a65726f20616464726573734b495031373a207472616e73666572206f6620746f6b656e2074686174206973206e6f74206f776e4b495031373a20617070726f76656420717565727920666f72206e6f6e6578697374656e7420746f6b656e4b495031373a20617070726f76652063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f76656420666f7220616c6c4b495031373a206f70657261746f7220717565727920666f72206e6f6e6578697374656e7420746f6b656ea165627a7a72305820b5e6a0f9975602c403a5605cb8bab015e95727b9bf736067787cbb95941531230029`

// KIP17FuncSigs maps the 4-byte function signature to its string representation.
var KIP17FuncSigs = map[string]string{
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"081812fc": "getApproved(uint256)",
	"e985e9c5": "isApprovedForAll(address,address)",
	"6352211e": "ownerOf(uint256)",
	"42842e0e": "safeTransferFrom(address,address,uint256)",
	"b88d4fde": "safeTransferFrom(address,address,uint256,bytes)",
	"a22cb465": "setApprovalForAll(address,bool)",
	"01ffc9a7": "supportsInterface(bytes4)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// KIP17Bin is the compiled bytecode used for deploying new contracts.
var KIP17Bin = "0x608060405234801561001057600080fd5b506100276301ffc9a760e01b61004260201b60201c565b61003d6380ac58cd60e01b61004260201b60201c565b610110565b7fffffffff0000000000000000000000000000000000000000000000000000000080821614156100d357604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f4b495031333a20696e76616c696420696e746572666163652069640000000000604482015290519081900360640190fd5b7fffffffff00000000000000000000000000000000000000000000000000000000166000908152602081905260409020805460ff19166001179055565b610fd98061011f6000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c80636352211e116100665780636352211e146101b157806370a08231146101ce578063a22cb46514610206578063b88d4fde14610234578063e985e9c5146102fa5761009e565b806301ffc9a7146100a3578063081812fc146100de578063095ea7b31461011757806323b872dd1461014557806342842e0e1461017b575b600080fd5b6100ca600480360360208110156100b957600080fd5b50356001600160e01b031916610328565b604080519115158252519081900360200190f35b6100fb600480360360208110156100f457600080fd5b5035610347565b604080516001600160a01b039092168252519081900360200190f35b6101436004803603604081101561012d57600080fd5b506001600160a01b0381351690602001356103ac565b005b6101436004803603606081101561015b57600080fd5b506001600160a01b038135811691602081013590911690604001356104d9565b6101436004803603606081101561019157600080fd5b506001600160a01b03813581169160208101359091169060400135610531565b6100fb600480360360208110156101c757600080fd5b503561054c565b6101f4600480360360208110156101e457600080fd5b50356001600160a01b03166105a9565b60408051918252519081900360200190f35b6101436004803603604081101561021c57600080fd5b506001600160a01b0381351690602001351515610614565b6101436004803603608081101561024a57600080fd5b6001600160a01b0382358116926020810135909116916040820135919081019060808101606082013564010000000081111561028557600080fd5b82018360208201111561029757600080fd5b803590602001918460018302840111640100000000831117156102b957600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506106e3945050505050565b6100ca6004803603604081101561031057600080fd5b506001600160a01b038135811691602001351661073e565b6001600160e01b03191660009081526020819052604090205460ff1690565b60006103528261076c565b61039057604051600160e51b62461bcd02815260040180806020018281038252602b815260200180610f21602b913960400191505060405180910390fd5b506000908152600260205260409020546001600160a01b031690565b60006103b78261054c565b9050806001600160a01b0316836001600160a01b031614156104235760408051600160e51b62461bcd02815260206004820181905260248201527f4b495031373a20617070726f76616c20746f2063757272656e74206f776e6572604482015290519081900360640190fd5b336001600160a01b038216148061043f575061043f813361073e565b61047d57604051600160e51b62461bcd028152600401808060200182810382526037815260200180610f4c6037913960400191505060405180910390fd5b60008281526002602052604080822080546001600160a01b0319166001600160a01b0387811691821790925591518593918516917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591a4505050565b6104e33382610789565b61052157604051600160e51b62461bcd028152600401808060200182810382526030815260200180610e706030913960400191505060405180910390fd5b61052c838383610830565b505050565b61052c838383604051806020016040528060008152506106e3565b6000818152600160205260408120546001600160a01b0316806105a357604051600160e51b62461bcd028152600401808060200182810382526028815260200180610e486028913960400191505060405180910390fd5b92915050565b60006001600160a01b0382166105f357604051600160e51b62461bcd028152600401808060200182810382526029815260200180610ed06029913960400191505060405180910390fd5b6001600160a01b03821660009081526003602052604090206105a39061097a565b6001600160a01b0382163314156106755760408051600160e51b62461bcd02815260206004820152601860248201527f4b495031373a20617070726f766520746f2063616c6c65720000000000000000604482015290519081900360640190fd5b3360008181526004602090815260408083206001600160a01b03871680855290835292819020805460ff1916861515908117909155815190815290519293927f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31929181900390910190a35050565b6106ee8484846104d9565b6106fa8484848461097e565b61073857604051600160e51b62461bcd028152600401808060200182810382526030815260200180610ea06030913960400191505060405180910390fd5b50505050565b6001600160a01b03918216600090815260046020908152604080832093909416825291909152205460ff1690565b6000908152600160205260409020546001600160a01b0316151590565b60006107948261076c565b6107d257604051600160e51b62461bcd02815260040180806020018281038252602b815260200180610f83602b913960400191505060405180910390fd5b60006107dd8361054c565b9050806001600160a01b0316846001600160a01b031614806108185750836001600160a01b031661080d84610347565b6001600160a01b0316145b806108285750610828818561073e565b949350505050565b826001600160a01b03166108438261054c565b6001600160a01b03161461088b57604051600160e51b62461bcd028152600401808060200182810382526028815260200180610ef96028913960400191505060405180910390fd5b6001600160a01b0382166108d357604051600160e51b62461bcd028152600401808060200182810382526023815260200180610e256023913960400191505060405180910390fd5b6108dc81610d61565b6001600160a01b03831660009081526003602052604090206108fd90610d9e565b6001600160a01b038216600090815260036020526040902061091e90610db5565b60008181526001602052604080822080546001600160a01b0319166001600160a01b0386811691821790925591518493918716917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b5490565b6000806060610995866001600160a01b0316610dbe565b6109a457600192505050610828565b856001600160a01b031663150b7a0260e01b3389888860405160240180856001600160a01b03166001600160a01b03168152602001846001600160a01b03166001600160a01b0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610a30578181015183820152602001610a18565b50505050905090810190601f168015610a5d5780820380516001836020036101000a031916815260200191505b5060408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909a16999099178952518151919890975087965094509250829150849050835b60208310610ac55780518252601f199092019160209182019101610aa6565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114610b27576040519150601f19603f3d011682016040523d82523d6000602084013e610b2c565b606091505b508051919350915015801590610b6c57508051600160e11b630a85bd01029060208084019190811015610b5e57600080fd5b50516001600160e01b031916145b15610b7c57600192505050610828565b856001600160a01b0316636745782b60e01b3389888860405160240180856001600160a01b03166001600160a01b03168152602001846001600160a01b03166001600160a01b0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610c08578181015183820152602001610bf0565b50505050905090810190601f168015610c355780820380516001836020036101000a031916815260200191505b5060408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909a16999099178952518151919890975087965094509250829150849050835b60208310610c9d5780518252601f199092019160209182019101610c7e565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114610cff576040519150601f19603f3d011682016040523d82523d6000602084013e610d04565b606091505b508051919350915015801590610d4457508051600160e01b636745782b029060208084019190811015610d3657600080fd5b50516001600160e01b031916145b15610d5457600192505050610828565b5060009695505050505050565b6000818152600260205260409020546001600160a01b031615610d9b57600081815260026020526040902080546001600160a01b03191690555b50565b8054610db190600163ffffffff610dc416565b9055565b80546001019055565b3b151590565b600082821115610e1e5760408051600160e51b62461bcd02815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b5090039056fe4b495031373a207472616e7366657220746f20746865207a65726f20616464726573734b495031373a206f776e657220717565727920666f72206e6f6e6578697374656e7420746f6b656e4b495031373a207472616e736665722063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f7665644b495031373a207472616e7366657220746f206e6f6e204b49503137526563656976657220696d706c656d656e7465724b495031373a2062616c616e636520717565727920666f7220746865207a65726f20616464726573734b495031373a207472616e73666572206f6620746f6b656e2074686174206973206e6f74206f776e4b495031373a20617070726f76656420717565727920666f72206e6f6e6578697374656e7420746f6b656e4b495031373a20617070726f76652063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f76656420666f7220616c6c4b495031373a206f70657261746f7220717565727920666f72206e6f6e6578697374656e7420746f6b656ea165627a7a72305820b5e6a0f9975602c403a5605cb8bab015e95727b9bf736067787cbb95941531230029"

// DeployKIP17 deploys a new Klaytn contract, binding an instance of KIP17 to it.
func DeployKIP17(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *KIP17, error) {
	parsed, err := abi.JSON(strings.NewReader(KIP17ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(KIP17Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KIP17{KIP17Caller: KIP17Caller{contract: contract}, KIP17Transactor: KIP17Transactor{contract: contract}, KIP17Filterer: KIP17Filterer{contract: contract}}, nil
}

// KIP17 is an auto generated Go binding around a Klaytn contract.
type KIP17 struct {
	KIP17Caller     // Read-only binding to the contract
	KIP17Transactor // Write-only binding to the contract
	KIP17Filterer   // Log filterer for contract events
}

// KIP17Caller is an auto generated read-only Go binding around a Klaytn contract.
type KIP17Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KIP17Transactor is an auto generated write-only Go binding around a Klaytn contract.
type KIP17Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KIP17Filterer is an auto generated log filtering Go binding around a Klaytn contract events.
type KIP17Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KIP17Session is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type KIP17Session struct {
	Contract     *KIP17            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KIP17CallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type KIP17CallerSession struct {
	Contract *KIP17Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// KIP17TransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type KIP17TransactorSession struct {
	Contract     *KIP17Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KIP17Raw is an auto generated low-level Go binding around a Klaytn contract.
type KIP17Raw struct {
	Contract *KIP17 // Generic contract binding to access the raw methods on
}

// KIP17CallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type KIP17CallerRaw struct {
	Contract *KIP17Caller // Generic read-only contract binding to access the raw methods on
}

// KIP17TransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type KIP17TransactorRaw struct {
	Contract *KIP17Transactor // Generic write-only contract binding to access the raw methods on
}

// NewKIP17 creates a new instance of KIP17, bound to a specific deployed contract.
func NewKIP17(address common.Address, backend bind.ContractBackend) (*KIP17, error) {
	contract, err := bindKIP17(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KIP17{KIP17Caller: KIP17Caller{contract: contract}, KIP17Transactor: KIP17Transactor{contract: contract}, KIP17Filterer: KIP17Filterer{contract: contract}}, nil
}

// NewKIP17Caller creates a new read-only instance of KIP17, bound to a specific deployed contract.
func NewKIP17Caller(address common.Address, caller bind.ContractCaller) (*KIP17Caller, error) {
	contract, err := bindKIP17(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KIP17Caller{contract: contract}, nil
}

// NewKIP17Transactor creates a new write-only instance of KIP17, bound to a specific deployed contract.
func NewKIP17Transactor(address common.Address, transactor bind.ContractTransactor) (*KIP17Transactor, error) {
	contract, err := bindKIP17(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KIP17Transactor{contract: contract}, nil
}

// NewKIP17Filterer creates a new log filterer instance of KIP17, bound to a specific deployed contract.
func NewKIP17Filterer(address common.Address, filterer bind.ContractFilterer) (*KIP17Filterer, error) {
	contract, err := bindKIP17(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KIP17Filterer{contract: contract}, nil
}

// bindKIP17 binds a generic wrapper to an already deployed contract.
func bindKIP17(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KIP17ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KIP17 *KIP17Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KIP17.Contract.KIP17Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KIP17 *KIP17Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KIP17.Contract.KIP17Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KIP17 *KIP17Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KIP17.Contract.KIP17Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KIP17 *KIP17CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KIP17.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KIP17 *KIP17TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KIP17.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KIP17 *KIP17TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KIP17.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_KIP17 *KIP17Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KIP17.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_KIP17 *KIP17Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _KIP17.Contract.BalanceOf(&_KIP17.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_KIP17 *KIP17CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _KIP17.Contract.BalanceOf(&_KIP17.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_KIP17 *KIP17Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KIP17.contract.Call(opts, out, "getApproved", tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_KIP17 *KIP17Session) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _KIP17.Contract.GetApproved(&_KIP17.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_KIP17 *KIP17CallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _KIP17.Contract.GetApproved(&_KIP17.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_KIP17 *KIP17Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KIP17.contract.Call(opts, out, "isApprovedForAll", owner, operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_KIP17 *KIP17Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _KIP17.Contract.IsApprovedForAll(&_KIP17.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_KIP17 *KIP17CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _KIP17.Contract.IsApprovedForAll(&_KIP17.CallOpts, owner, operator)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_KIP17 *KIP17Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KIP17.contract.Call(opts, out, "ownerOf", tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_KIP17 *KIP17Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _KIP17.Contract.OwnerOf(&_KIP17.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_KIP17 *KIP17CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _KIP17.Contract.OwnerOf(&_KIP17.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_KIP17 *KIP17Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KIP17.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_KIP17 *KIP17Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _KIP17.Contract.SupportsInterface(&_KIP17.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_KIP17 *KIP17CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _KIP17.Contract.SupportsInterface(&_KIP17.CallOpts, interfaceId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_KIP17 *KIP17Transactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_KIP17 *KIP17Session) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17.Contract.Approve(&_KIP17.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_KIP17 *KIP17TransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17.Contract.Approve(&_KIP17.TransactOpts, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_KIP17 *KIP17Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_KIP17 *KIP17Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17.Contract.SafeTransferFrom(&_KIP17.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_KIP17 *KIP17TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17.Contract.SafeTransferFrom(&_KIP17.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_KIP17 *KIP17Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _KIP17.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_KIP17 *KIP17Session) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _KIP17.Contract.SafeTransferFrom0(&_KIP17.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_KIP17 *KIP17TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _KIP17.Contract.SafeTransferFrom0(&_KIP17.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_KIP17 *KIP17Transactor) SetApprovalForAll(opts *bind.TransactOpts, to common.Address, approved bool) (*types.Transaction, error) {
	return _KIP17.contract.Transact(opts, "setApprovalForAll", to, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_KIP17 *KIP17Session) SetApprovalForAll(to common.Address, approved bool) (*types.Transaction, error) {
	return _KIP17.Contract.SetApprovalForAll(&_KIP17.TransactOpts, to, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_KIP17 *KIP17TransactorSession) SetApprovalForAll(to common.Address, approved bool) (*types.Transaction, error) {
	return _KIP17.Contract.SetApprovalForAll(&_KIP17.TransactOpts, to, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_KIP17 *KIP17Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_KIP17 *KIP17Session) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17.Contract.TransferFrom(&_KIP17.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_KIP17 *KIP17TransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17.Contract.TransferFrom(&_KIP17.TransactOpts, from, to, tokenId)
}

// KIP17ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the KIP17 contract.
type KIP17ApprovalIterator struct {
	Event *KIP17Approval // Event containing the contract specifics and raw log

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
func (it *KIP17ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KIP17Approval)
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
		it.Event = new(KIP17Approval)
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
func (it *KIP17ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KIP17ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KIP17Approval represents a Approval event raised by the KIP17 contract.
type KIP17Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_KIP17 *KIP17Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*KIP17ApprovalIterator, error) {

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

	logs, sub, err := _KIP17.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &KIP17ApprovalIterator{contract: _KIP17.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_KIP17 *KIP17Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *KIP17Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _KIP17.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KIP17Approval)
				if err := _KIP17.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_KIP17 *KIP17Filterer) ParseApproval(log types.Log) (*KIP17Approval, error) {
	event := new(KIP17Approval)
	if err := _KIP17.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KIP17ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the KIP17 contract.
type KIP17ApprovalForAllIterator struct {
	Event *KIP17ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *KIP17ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KIP17ApprovalForAll)
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
		it.Event = new(KIP17ApprovalForAll)
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
func (it *KIP17ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KIP17ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KIP17ApprovalForAll represents a ApprovalForAll event raised by the KIP17 contract.
type KIP17ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_KIP17 *KIP17Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*KIP17ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _KIP17.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &KIP17ApprovalForAllIterator{contract: _KIP17.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_KIP17 *KIP17Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *KIP17ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _KIP17.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KIP17ApprovalForAll)
				if err := _KIP17.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_KIP17 *KIP17Filterer) ParseApprovalForAll(log types.Log) (*KIP17ApprovalForAll, error) {
	event := new(KIP17ApprovalForAll)
	if err := _KIP17.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KIP17TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the KIP17 contract.
type KIP17TransferIterator struct {
	Event *KIP17Transfer // Event containing the contract specifics and raw log

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
func (it *KIP17TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KIP17Transfer)
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
		it.Event = new(KIP17Transfer)
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
func (it *KIP17TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KIP17TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KIP17Transfer represents a Transfer event raised by the KIP17 contract.
type KIP17Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_KIP17 *KIP17Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*KIP17TransferIterator, error) {

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

	logs, sub, err := _KIP17.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &KIP17TransferIterator{contract: _KIP17.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_KIP17 *KIP17Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *KIP17Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _KIP17.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KIP17Transfer)
				if err := _KIP17.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_KIP17 *KIP17Filterer) ParseTransfer(log types.Log) (*KIP17Transfer, error) {
	event := new(KIP17Transfer)
	if err := _KIP17.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KIP17BurnableABI is the input ABI used to generate the binding from.
const KIP17BurnableABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"}]"

// KIP17BurnableBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const KIP17BurnableBinRuntime = `608060405234801561001057600080fd5b50600436106100a95760003560e01c806342966c681161007157806342966c68146101bc5780636352211e146101d957806370a08231146101f6578063a22cb4651461022e578063b88d4fde1461025c578063e985e9c514610322576100a9565b806301ffc9a7146100ae578063081812fc146100e9578063095ea7b31461012257806323b872dd1461015057806342842e0e14610186575b600080fd5b6100d5600480360360208110156100c457600080fd5b50356001600160e01b031916610350565b604080519115158252519081900360200190f35b610106600480360360208110156100ff57600080fd5b503561036f565b604080516001600160a01b039092168252519081900360200190f35b61014e6004803603604081101561013857600080fd5b506001600160a01b0381351690602001356103d4565b005b61014e6004803603606081101561016657600080fd5b506001600160a01b03813581169160208101359091169060400135610501565b61014e6004803603606081101561019c57600080fd5b506001600160a01b03813581169160208101359091169060400135610559565b61014e600480360360208110156101d257600080fd5b5035610574565b610106600480360360208110156101ef57600080fd5b50356105c8565b61021c6004803603602081101561020c57600080fd5b50356001600160a01b0316610625565b60408051918252519081900360200190f35b61014e6004803603604081101561024457600080fd5b506001600160a01b0381351690602001351515610690565b61014e6004803603608081101561027257600080fd5b6001600160a01b038235811692602081013590911691604082013591908101906080810160608201356401000000008111156102ad57600080fd5b8201836020820111156102bf57600080fd5b803590602001918460018302840111640100000000831117156102e157600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061075f945050505050565b6100d56004803603604081101561033857600080fd5b506001600160a01b03813581169160200135166107ba565b6001600160e01b03191660009081526020819052604090205460ff1690565b600061037a826107e8565b6103b857604051600160e51b62461bcd02815260040180806020018281038252602b8152602001806110da602b913960400191505060405180910390fd5b506000908152600260205260409020546001600160a01b031690565b60006103df826105c8565b9050806001600160a01b0316836001600160a01b0316141561044b5760408051600160e51b62461bcd02815260206004820181905260248201527f4b495031373a20617070726f76616c20746f2063757272656e74206f776e6572604482015290519081900360640190fd5b336001600160a01b0382161480610467575061046781336107ba565b6104a557604051600160e51b62461bcd0281526004018080602001828103825260378152602001806111056037913960400191505060405180910390fd5b60008281526002602052604080822080546001600160a01b0319166001600160a01b0387811691821790925591518593918516917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591a4505050565b61050b3382610805565b61054957604051600160e51b62461bcd028152600401808060200182810382526030815260200180610fd66030913960400191505060405180910390fd5b6105548383836108ac565b505050565b6105548383836040518060200160405280600081525061075f565b61057e3382610805565b6105bc57604051600160e51b62461bcd02815260040180806020018281038252602f815260200180611036602f913960400191505060405180910390fd5b6105c5816109f6565b50565b6000818152600160205260408120546001600160a01b03168061061f57604051600160e51b62461bcd028152600401808060200182810382526028815260200180610fae6028913960400191505060405180910390fd5b92915050565b60006001600160a01b03821661066f57604051600160e51b62461bcd0281526004018080602001828103825260298152602001806110656029913960400191505060405180910390fd5b6001600160a01b038216600090815260036020526040902061061f90610a08565b6001600160a01b0382163314156106f15760408051600160e51b62461bcd02815260206004820152601860248201527f4b495031373a20617070726f766520746f2063616c6c65720000000000000000604482015290519081900360640190fd5b3360008181526004602090815260408083206001600160a01b03871680855290835292819020805460ff1916861515908117909155815190815290519293927f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31929181900390910190a35050565b61076a848484610501565b61077684848484610a0c565b6107b457604051600160e51b62461bcd0281526004018080602001828103825260308152602001806110066030913960400191505060405180910390fd5b50505050565b6001600160a01b03918216600090815260046020908152604080832093909416825291909152205460ff1690565b6000908152600160205260409020546001600160a01b0316151590565b6000610810826107e8565b61084e57604051600160e51b62461bcd02815260040180806020018281038252602b81526020018061113c602b913960400191505060405180910390fd5b6000610859836105c8565b9050806001600160a01b0316846001600160a01b031614806108945750836001600160a01b03166108898461036f565b6001600160a01b0316145b806108a457506108a481856107ba565b949350505050565b826001600160a01b03166108bf826105c8565b6001600160a01b03161461090757604051600160e51b62461bcd02815260040180806020018281038252602881526020018061108e6028913960400191505060405180910390fd5b6001600160a01b03821661094f57604051600160e51b62461bcd028152600401808060200182810382526023815260200180610f8b6023913960400191505060405180910390fd5b61095881610def565b6001600160a01b038316600090815260036020526040902061097990610e2a565b6001600160a01b038216600090815260036020526040902061099a90610e41565b60008181526001602052604080822080546001600160a01b0319166001600160a01b0386811691821790925591518493918716917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b6105c5610a02826105c8565b82610e4a565b5490565b6000806060610a23866001600160a01b0316610f24565b610a32576001925050506108a4565b856001600160a01b031663150b7a0260e01b3389888860405160240180856001600160a01b03166001600160a01b03168152602001846001600160a01b03166001600160a01b0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610abe578181015183820152602001610aa6565b50505050905090810190601f168015610aeb5780820380516001836020036101000a031916815260200191505b5060408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909a16999099178952518151919890975087965094509250829150849050835b60208310610b535780518252601f199092019160209182019101610b34565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114610bb5576040519150601f19603f3d011682016040523d82523d6000602084013e610bba565b606091505b508051919350915015801590610bfa57508051600160e11b630a85bd01029060208084019190811015610bec57600080fd5b50516001600160e01b031916145b15610c0a576001925050506108a4565b856001600160a01b0316636745782b60e01b3389888860405160240180856001600160a01b03166001600160a01b03168152602001846001600160a01b03166001600160a01b0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610c96578181015183820152602001610c7e565b50505050905090810190601f168015610cc35780820380516001836020036101000a031916815260200191505b5060408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909a16999099178952518151919890975087965094509250829150849050835b60208310610d2b5780518252601f199092019160209182019101610d0c565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114610d8d576040519150601f19603f3d011682016040523d82523d6000602084013e610d92565b606091505b508051919350915015801590610dd257508051600160e01b636745782b029060208084019190811015610dc457600080fd5b50516001600160e01b031916145b15610de2576001925050506108a4565b5060009695505050505050565b6000818152600260205260409020546001600160a01b0316156105c557600090815260026020526040902080546001600160a01b0319169055565b8054610e3d90600163ffffffff610f2a16565b9055565b80546001019055565b816001600160a01b0316610e5d826105c8565b6001600160a01b031614610ea557604051600160e51b62461bcd0281526004018080602001828103825260248152602001806110b66024913960400191505060405180910390fd5b610eae81610def565b6001600160a01b0382166000908152600360205260409020610ecf90610e2a565b60008181526001602052604080822080546001600160a01b0319169055518291906001600160a01b038516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908390a45050565b3b151590565b600082821115610f845760408051600160e51b62461bcd02815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b5090039056fe4b495031373a207472616e7366657220746f20746865207a65726f20616464726573734b495031373a206f776e657220717565727920666f72206e6f6e6578697374656e7420746f6b656e4b495031373a207472616e736665722063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f7665644b495031373a207472616e7366657220746f206e6f6e204b49503137526563656976657220696d706c656d656e7465724b495031374275726e61626c653a2063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f7665644b495031373a2062616c616e636520717565727920666f7220746865207a65726f20616464726573734b495031373a207472616e73666572206f6620746f6b656e2074686174206973206e6f74206f776e4b495031373a206275726e206f6620746f6b656e2074686174206973206e6f74206f776e4b495031373a20617070726f76656420717565727920666f72206e6f6e6578697374656e7420746f6b656e4b495031373a20617070726f76652063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f76656420666f7220616c6c4b495031373a206f70657261746f7220717565727920666f72206e6f6e6578697374656e7420746f6b656ea165627a7a723058200193132c025d025b3d80badcdd11141728c3d345b411150b212c7ba32460fc260029`

// KIP17BurnableFuncSigs maps the 4-byte function signature to its string representation.
var KIP17BurnableFuncSigs = map[string]string{
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"42966c68": "burn(uint256)",
	"081812fc": "getApproved(uint256)",
	"e985e9c5": "isApprovedForAll(address,address)",
	"6352211e": "ownerOf(uint256)",
	"42842e0e": "safeTransferFrom(address,address,uint256)",
	"b88d4fde": "safeTransferFrom(address,address,uint256,bytes)",
	"a22cb465": "setApprovalForAll(address,bool)",
	"01ffc9a7": "supportsInterface(bytes4)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// KIP17BurnableBin is the compiled bytecode used for deploying new contracts.
var KIP17BurnableBin = "0x608060405234801561001057600080fd5b506100276301ffc9a760e01b61005860201b60201c565b61003d6380ac58cd60e01b61005860201b60201c565b6100536342966c6860e01b61005860201b60201c565b610126565b7fffffffff0000000000000000000000000000000000000000000000000000000080821614156100e957604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f4b495031333a20696e76616c696420696e746572666163652069640000000000604482015290519081900360640190fd5b7fffffffff00000000000000000000000000000000000000000000000000000000166000908152602081905260409020805460ff19166001179055565b611192806101356000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c806342966c681161007157806342966c68146101bc5780636352211e146101d957806370a08231146101f6578063a22cb4651461022e578063b88d4fde1461025c578063e985e9c514610322576100a9565b806301ffc9a7146100ae578063081812fc146100e9578063095ea7b31461012257806323b872dd1461015057806342842e0e14610186575b600080fd5b6100d5600480360360208110156100c457600080fd5b50356001600160e01b031916610350565b604080519115158252519081900360200190f35b610106600480360360208110156100ff57600080fd5b503561036f565b604080516001600160a01b039092168252519081900360200190f35b61014e6004803603604081101561013857600080fd5b506001600160a01b0381351690602001356103d4565b005b61014e6004803603606081101561016657600080fd5b506001600160a01b03813581169160208101359091169060400135610501565b61014e6004803603606081101561019c57600080fd5b506001600160a01b03813581169160208101359091169060400135610559565b61014e600480360360208110156101d257600080fd5b5035610574565b610106600480360360208110156101ef57600080fd5b50356105c8565b61021c6004803603602081101561020c57600080fd5b50356001600160a01b0316610625565b60408051918252519081900360200190f35b61014e6004803603604081101561024457600080fd5b506001600160a01b0381351690602001351515610690565b61014e6004803603608081101561027257600080fd5b6001600160a01b038235811692602081013590911691604082013591908101906080810160608201356401000000008111156102ad57600080fd5b8201836020820111156102bf57600080fd5b803590602001918460018302840111640100000000831117156102e157600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061075f945050505050565b6100d56004803603604081101561033857600080fd5b506001600160a01b03813581169160200135166107ba565b6001600160e01b03191660009081526020819052604090205460ff1690565b600061037a826107e8565b6103b857604051600160e51b62461bcd02815260040180806020018281038252602b8152602001806110da602b913960400191505060405180910390fd5b506000908152600260205260409020546001600160a01b031690565b60006103df826105c8565b9050806001600160a01b0316836001600160a01b0316141561044b5760408051600160e51b62461bcd02815260206004820181905260248201527f4b495031373a20617070726f76616c20746f2063757272656e74206f776e6572604482015290519081900360640190fd5b336001600160a01b0382161480610467575061046781336107ba565b6104a557604051600160e51b62461bcd0281526004018080602001828103825260378152602001806111056037913960400191505060405180910390fd5b60008281526002602052604080822080546001600160a01b0319166001600160a01b0387811691821790925591518593918516917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591a4505050565b61050b3382610805565b61054957604051600160e51b62461bcd028152600401808060200182810382526030815260200180610fd66030913960400191505060405180910390fd5b6105548383836108ac565b505050565b6105548383836040518060200160405280600081525061075f565b61057e3382610805565b6105bc57604051600160e51b62461bcd02815260040180806020018281038252602f815260200180611036602f913960400191505060405180910390fd5b6105c5816109f6565b50565b6000818152600160205260408120546001600160a01b03168061061f57604051600160e51b62461bcd028152600401808060200182810382526028815260200180610fae6028913960400191505060405180910390fd5b92915050565b60006001600160a01b03821661066f57604051600160e51b62461bcd0281526004018080602001828103825260298152602001806110656029913960400191505060405180910390fd5b6001600160a01b038216600090815260036020526040902061061f90610a08565b6001600160a01b0382163314156106f15760408051600160e51b62461bcd02815260206004820152601860248201527f4b495031373a20617070726f766520746f2063616c6c65720000000000000000604482015290519081900360640190fd5b3360008181526004602090815260408083206001600160a01b03871680855290835292819020805460ff1916861515908117909155815190815290519293927f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31929181900390910190a35050565b61076a848484610501565b61077684848484610a0c565b6107b457604051600160e51b62461bcd0281526004018080602001828103825260308152602001806110066030913960400191505060405180910390fd5b50505050565b6001600160a01b03918216600090815260046020908152604080832093909416825291909152205460ff1690565b6000908152600160205260409020546001600160a01b0316151590565b6000610810826107e8565b61084e57604051600160e51b62461bcd02815260040180806020018281038252602b81526020018061113c602b913960400191505060405180910390fd5b6000610859836105c8565b9050806001600160a01b0316846001600160a01b031614806108945750836001600160a01b03166108898461036f565b6001600160a01b0316145b806108a457506108a481856107ba565b949350505050565b826001600160a01b03166108bf826105c8565b6001600160a01b03161461090757604051600160e51b62461bcd02815260040180806020018281038252602881526020018061108e6028913960400191505060405180910390fd5b6001600160a01b03821661094f57604051600160e51b62461bcd028152600401808060200182810382526023815260200180610f8b6023913960400191505060405180910390fd5b61095881610def565b6001600160a01b038316600090815260036020526040902061097990610e2a565b6001600160a01b038216600090815260036020526040902061099a90610e41565b60008181526001602052604080822080546001600160a01b0319166001600160a01b0386811691821790925591518493918716917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b6105c5610a02826105c8565b82610e4a565b5490565b6000806060610a23866001600160a01b0316610f24565b610a32576001925050506108a4565b856001600160a01b031663150b7a0260e01b3389888860405160240180856001600160a01b03166001600160a01b03168152602001846001600160a01b03166001600160a01b0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610abe578181015183820152602001610aa6565b50505050905090810190601f168015610aeb5780820380516001836020036101000a031916815260200191505b5060408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909a16999099178952518151919890975087965094509250829150849050835b60208310610b535780518252601f199092019160209182019101610b34565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114610bb5576040519150601f19603f3d011682016040523d82523d6000602084013e610bba565b606091505b508051919350915015801590610bfa57508051600160e11b630a85bd01029060208084019190811015610bec57600080fd5b50516001600160e01b031916145b15610c0a576001925050506108a4565b856001600160a01b0316636745782b60e01b3389888860405160240180856001600160a01b03166001600160a01b03168152602001846001600160a01b03166001600160a01b0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610c96578181015183820152602001610c7e565b50505050905090810190601f168015610cc35780820380516001836020036101000a031916815260200191505b5060408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909a16999099178952518151919890975087965094509250829150849050835b60208310610d2b5780518252601f199092019160209182019101610d0c565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114610d8d576040519150601f19603f3d011682016040523d82523d6000602084013e610d92565b606091505b508051919350915015801590610dd257508051600160e01b636745782b029060208084019190811015610dc457600080fd5b50516001600160e01b031916145b15610de2576001925050506108a4565b5060009695505050505050565b6000818152600260205260409020546001600160a01b0316156105c557600090815260026020526040902080546001600160a01b0319169055565b8054610e3d90600163ffffffff610f2a16565b9055565b80546001019055565b816001600160a01b0316610e5d826105c8565b6001600160a01b031614610ea557604051600160e51b62461bcd0281526004018080602001828103825260248152602001806110b66024913960400191505060405180910390fd5b610eae81610def565b6001600160a01b0382166000908152600360205260409020610ecf90610e2a565b60008181526001602052604080822080546001600160a01b0319169055518291906001600160a01b038516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908390a45050565b3b151590565b600082821115610f845760408051600160e51b62461bcd02815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b5090039056fe4b495031373a207472616e7366657220746f20746865207a65726f20616464726573734b495031373a206f776e657220717565727920666f72206e6f6e6578697374656e7420746f6b656e4b495031373a207472616e736665722063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f7665644b495031373a207472616e7366657220746f206e6f6e204b49503137526563656976657220696d706c656d656e7465724b495031374275726e61626c653a2063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f7665644b495031373a2062616c616e636520717565727920666f7220746865207a65726f20616464726573734b495031373a207472616e73666572206f6620746f6b656e2074686174206973206e6f74206f776e4b495031373a206275726e206f6620746f6b656e2074686174206973206e6f74206f776e4b495031373a20617070726f76656420717565727920666f72206e6f6e6578697374656e7420746f6b656e4b495031373a20617070726f76652063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f76656420666f7220616c6c4b495031373a206f70657261746f7220717565727920666f72206e6f6e6578697374656e7420746f6b656ea165627a7a723058200193132c025d025b3d80badcdd11141728c3d345b411150b212c7ba32460fc260029"

// DeployKIP17Burnable deploys a new Klaytn contract, binding an instance of KIP17Burnable to it.
func DeployKIP17Burnable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *KIP17Burnable, error) {
	parsed, err := abi.JSON(strings.NewReader(KIP17BurnableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(KIP17BurnableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KIP17Burnable{KIP17BurnableCaller: KIP17BurnableCaller{contract: contract}, KIP17BurnableTransactor: KIP17BurnableTransactor{contract: contract}, KIP17BurnableFilterer: KIP17BurnableFilterer{contract: contract}}, nil
}

// KIP17Burnable is an auto generated Go binding around a Klaytn contract.
type KIP17Burnable struct {
	KIP17BurnableCaller     // Read-only binding to the contract
	KIP17BurnableTransactor // Write-only binding to the contract
	KIP17BurnableFilterer   // Log filterer for contract events
}

// KIP17BurnableCaller is an auto generated read-only Go binding around a Klaytn contract.
type KIP17BurnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KIP17BurnableTransactor is an auto generated write-only Go binding around a Klaytn contract.
type KIP17BurnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KIP17BurnableFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type KIP17BurnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KIP17BurnableSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type KIP17BurnableSession struct {
	Contract     *KIP17Burnable    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KIP17BurnableCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type KIP17BurnableCallerSession struct {
	Contract *KIP17BurnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// KIP17BurnableTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type KIP17BurnableTransactorSession struct {
	Contract     *KIP17BurnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// KIP17BurnableRaw is an auto generated low-level Go binding around a Klaytn contract.
type KIP17BurnableRaw struct {
	Contract *KIP17Burnable // Generic contract binding to access the raw methods on
}

// KIP17BurnableCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type KIP17BurnableCallerRaw struct {
	Contract *KIP17BurnableCaller // Generic read-only contract binding to access the raw methods on
}

// KIP17BurnableTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type KIP17BurnableTransactorRaw struct {
	Contract *KIP17BurnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKIP17Burnable creates a new instance of KIP17Burnable, bound to a specific deployed contract.
func NewKIP17Burnable(address common.Address, backend bind.ContractBackend) (*KIP17Burnable, error) {
	contract, err := bindKIP17Burnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KIP17Burnable{KIP17BurnableCaller: KIP17BurnableCaller{contract: contract}, KIP17BurnableTransactor: KIP17BurnableTransactor{contract: contract}, KIP17BurnableFilterer: KIP17BurnableFilterer{contract: contract}}, nil
}

// NewKIP17BurnableCaller creates a new read-only instance of KIP17Burnable, bound to a specific deployed contract.
func NewKIP17BurnableCaller(address common.Address, caller bind.ContractCaller) (*KIP17BurnableCaller, error) {
	contract, err := bindKIP17Burnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KIP17BurnableCaller{contract: contract}, nil
}

// NewKIP17BurnableTransactor creates a new write-only instance of KIP17Burnable, bound to a specific deployed contract.
func NewKIP17BurnableTransactor(address common.Address, transactor bind.ContractTransactor) (*KIP17BurnableTransactor, error) {
	contract, err := bindKIP17Burnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KIP17BurnableTransactor{contract: contract}, nil
}

// NewKIP17BurnableFilterer creates a new log filterer instance of KIP17Burnable, bound to a specific deployed contract.
func NewKIP17BurnableFilterer(address common.Address, filterer bind.ContractFilterer) (*KIP17BurnableFilterer, error) {
	contract, err := bindKIP17Burnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KIP17BurnableFilterer{contract: contract}, nil
}

// bindKIP17Burnable binds a generic wrapper to an already deployed contract.
func bindKIP17Burnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KIP17BurnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KIP17Burnable *KIP17BurnableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KIP17Burnable.Contract.KIP17BurnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KIP17Burnable *KIP17BurnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KIP17Burnable.Contract.KIP17BurnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KIP17Burnable *KIP17BurnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KIP17Burnable.Contract.KIP17BurnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KIP17Burnable *KIP17BurnableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KIP17Burnable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KIP17Burnable *KIP17BurnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KIP17Burnable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KIP17Burnable *KIP17BurnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KIP17Burnable.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_KIP17Burnable *KIP17BurnableCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KIP17Burnable.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_KIP17Burnable *KIP17BurnableSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _KIP17Burnable.Contract.BalanceOf(&_KIP17Burnable.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_KIP17Burnable *KIP17BurnableCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _KIP17Burnable.Contract.BalanceOf(&_KIP17Burnable.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_KIP17Burnable *KIP17BurnableCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KIP17Burnable.contract.Call(opts, out, "getApproved", tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_KIP17Burnable *KIP17BurnableSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _KIP17Burnable.Contract.GetApproved(&_KIP17Burnable.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_KIP17Burnable *KIP17BurnableCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _KIP17Burnable.Contract.GetApproved(&_KIP17Burnable.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_KIP17Burnable *KIP17BurnableCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KIP17Burnable.contract.Call(opts, out, "isApprovedForAll", owner, operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_KIP17Burnable *KIP17BurnableSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _KIP17Burnable.Contract.IsApprovedForAll(&_KIP17Burnable.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_KIP17Burnable *KIP17BurnableCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _KIP17Burnable.Contract.IsApprovedForAll(&_KIP17Burnable.CallOpts, owner, operator)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_KIP17Burnable *KIP17BurnableCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KIP17Burnable.contract.Call(opts, out, "ownerOf", tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_KIP17Burnable *KIP17BurnableSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _KIP17Burnable.Contract.OwnerOf(&_KIP17Burnable.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_KIP17Burnable *KIP17BurnableCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _KIP17Burnable.Contract.OwnerOf(&_KIP17Burnable.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_KIP17Burnable *KIP17BurnableCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KIP17Burnable.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_KIP17Burnable *KIP17BurnableSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _KIP17Burnable.Contract.SupportsInterface(&_KIP17Burnable.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_KIP17Burnable *KIP17BurnableCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _KIP17Burnable.Contract.SupportsInterface(&_KIP17Burnable.CallOpts, interfaceId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_KIP17Burnable *KIP17BurnableTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17Burnable.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_KIP17Burnable *KIP17BurnableSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17Burnable.Contract.Approve(&_KIP17Burnable.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_KIP17Burnable *KIP17BurnableTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17Burnable.Contract.Approve(&_KIP17Burnable.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_KIP17Burnable *KIP17BurnableTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17Burnable.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_KIP17Burnable *KIP17BurnableSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17Burnable.Contract.Burn(&_KIP17Burnable.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_KIP17Burnable *KIP17BurnableTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17Burnable.Contract.Burn(&_KIP17Burnable.TransactOpts, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_KIP17Burnable *KIP17BurnableTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17Burnable.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_KIP17Burnable *KIP17BurnableSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17Burnable.Contract.SafeTransferFrom(&_KIP17Burnable.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_KIP17Burnable *KIP17BurnableTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17Burnable.Contract.SafeTransferFrom(&_KIP17Burnable.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_KIP17Burnable *KIP17BurnableTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _KIP17Burnable.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_KIP17Burnable *KIP17BurnableSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _KIP17Burnable.Contract.SafeTransferFrom0(&_KIP17Burnable.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_KIP17Burnable *KIP17BurnableTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _KIP17Burnable.Contract.SafeTransferFrom0(&_KIP17Burnable.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_KIP17Burnable *KIP17BurnableTransactor) SetApprovalForAll(opts *bind.TransactOpts, to common.Address, approved bool) (*types.Transaction, error) {
	return _KIP17Burnable.contract.Transact(opts, "setApprovalForAll", to, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_KIP17Burnable *KIP17BurnableSession) SetApprovalForAll(to common.Address, approved bool) (*types.Transaction, error) {
	return _KIP17Burnable.Contract.SetApprovalForAll(&_KIP17Burnable.TransactOpts, to, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_KIP17Burnable *KIP17BurnableTransactorSession) SetApprovalForAll(to common.Address, approved bool) (*types.Transaction, error) {
	return _KIP17Burnable.Contract.SetApprovalForAll(&_KIP17Burnable.TransactOpts, to, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_KIP17Burnable *KIP17BurnableTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17Burnable.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_KIP17Burnable *KIP17BurnableSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17Burnable.Contract.TransferFrom(&_KIP17Burnable.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_KIP17Burnable *KIP17BurnableTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17Burnable.Contract.TransferFrom(&_KIP17Burnable.TransactOpts, from, to, tokenId)
}

// KIP17BurnableApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the KIP17Burnable contract.
type KIP17BurnableApprovalIterator struct {
	Event *KIP17BurnableApproval // Event containing the contract specifics and raw log

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
func (it *KIP17BurnableApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KIP17BurnableApproval)
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
		it.Event = new(KIP17BurnableApproval)
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
func (it *KIP17BurnableApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KIP17BurnableApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KIP17BurnableApproval represents a Approval event raised by the KIP17Burnable contract.
type KIP17BurnableApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_KIP17Burnable *KIP17BurnableFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*KIP17BurnableApprovalIterator, error) {

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

	logs, sub, err := _KIP17Burnable.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &KIP17BurnableApprovalIterator{contract: _KIP17Burnable.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_KIP17Burnable *KIP17BurnableFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *KIP17BurnableApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _KIP17Burnable.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KIP17BurnableApproval)
				if err := _KIP17Burnable.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_KIP17Burnable *KIP17BurnableFilterer) ParseApproval(log types.Log) (*KIP17BurnableApproval, error) {
	event := new(KIP17BurnableApproval)
	if err := _KIP17Burnable.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KIP17BurnableApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the KIP17Burnable contract.
type KIP17BurnableApprovalForAllIterator struct {
	Event *KIP17BurnableApprovalForAll // Event containing the contract specifics and raw log

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
func (it *KIP17BurnableApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KIP17BurnableApprovalForAll)
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
		it.Event = new(KIP17BurnableApprovalForAll)
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
func (it *KIP17BurnableApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KIP17BurnableApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KIP17BurnableApprovalForAll represents a ApprovalForAll event raised by the KIP17Burnable contract.
type KIP17BurnableApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_KIP17Burnable *KIP17BurnableFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*KIP17BurnableApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _KIP17Burnable.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &KIP17BurnableApprovalForAllIterator{contract: _KIP17Burnable.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_KIP17Burnable *KIP17BurnableFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *KIP17BurnableApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _KIP17Burnable.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KIP17BurnableApprovalForAll)
				if err := _KIP17Burnable.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_KIP17Burnable *KIP17BurnableFilterer) ParseApprovalForAll(log types.Log) (*KIP17BurnableApprovalForAll, error) {
	event := new(KIP17BurnableApprovalForAll)
	if err := _KIP17Burnable.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KIP17BurnableTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the KIP17Burnable contract.
type KIP17BurnableTransferIterator struct {
	Event *KIP17BurnableTransfer // Event containing the contract specifics and raw log

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
func (it *KIP17BurnableTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KIP17BurnableTransfer)
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
		it.Event = new(KIP17BurnableTransfer)
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
func (it *KIP17BurnableTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KIP17BurnableTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KIP17BurnableTransfer represents a Transfer event raised by the KIP17Burnable contract.
type KIP17BurnableTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_KIP17Burnable *KIP17BurnableFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*KIP17BurnableTransferIterator, error) {

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

	logs, sub, err := _KIP17Burnable.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &KIP17BurnableTransferIterator{contract: _KIP17Burnable.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_KIP17Burnable *KIP17BurnableFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *KIP17BurnableTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _KIP17Burnable.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KIP17BurnableTransfer)
				if err := _KIP17Burnable.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_KIP17Burnable *KIP17BurnableFilterer) ParseTransfer(log types.Log) (*KIP17BurnableTransfer, error) {
	event := new(KIP17BurnableTransfer)
	if err := _KIP17Burnable.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KIP17EnumerableABI is the input ABI used to generate the binding from.
const KIP17EnumerableABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"}]"

// KIP17EnumerableBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const KIP17EnumerableBinRuntime = `608060405234801561001057600080fd5b50600436106100cf5760003560e01c806342842e0e1161008c57806370a082311161006657806370a0823114610262578063a22cb46514610288578063b88d4fde146102b6578063e985e9c51461037c576100cf565b806342842e0e146101f25780634f6ccce7146102285780636352211e14610245576100cf565b806301ffc9a7146100d4578063081812fc1461010f578063095ea7b31461014857806318160ddd1461017657806323b872dd146101905780632f745c59146101c6575b600080fd5b6100fb600480360360208110156100ea57600080fd5b50356001600160e01b0319166103aa565b604080519115158252519081900360200190f35b61012c6004803603602081101561012557600080fd5b50356103c9565b604080516001600160a01b039092168252519081900360200190f35b6101746004803603604081101561015e57600080fd5b506001600160a01b03813516906020013561042e565b005b61017e61055b565b60408051918252519081900360200190f35b610174600480360360608110156101a657600080fd5b506001600160a01b03813581169160208101359091169060400135610562565b61017e600480360360408110156101dc57600080fd5b506001600160a01b0381351690602001356105ba565b6101746004803603606081101561020857600080fd5b506001600160a01b0381358116916020810135909116906040013561063c565b61017e6004803603602081101561023e57600080fd5b5035610657565b61012c6004803603602081101561025b57600080fd5b50356106c0565b61017e6004803603602081101561027857600080fd5b50356001600160a01b031661071d565b6101746004803603604081101561029e57600080fd5b506001600160a01b0381351690602001351515610788565b610174600480360360808110156102cc57600080fd5b6001600160a01b0382358116926020810135909116916040820135919081019060808101606082013564010000000081111561030757600080fd5b82018360208201111561031957600080fd5b8035906020019184600183028401116401000000008311171561033b57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610857945050505050565b6100fb6004803603604081101561039257600080fd5b506001600160a01b03813581169160200135166108b2565b6001600160e01b03191660009081526020819052604090205460ff1690565b60006103d4826108e0565b61041257604051600160e51b62461bcd02815260040180806020018281038252602b815260200180611279602b913960400191505060405180910390fd5b506000908152600260205260409020546001600160a01b031690565b6000610439826106c0565b9050806001600160a01b0316836001600160a01b031614156104a55760408051600160e51b62461bcd02815260206004820181905260248201527f4b495031373a20617070726f76616c20746f2063757272656e74206f776e6572604482015290519081900360640190fd5b336001600160a01b03821614806104c157506104c181336108b2565b6104ff57604051600160e51b62461bcd0281526004018080602001828103825260378152602001806112a46037913960400191505060405180910390fd5b60008281526002602052604080822080546001600160a01b0319166001600160a01b0387811691821790925591518593918516917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591a4505050565b6007545b90565b61056c33826108fd565b6105aa57604051600160e51b62461bcd02815260040180806020018281038252603081526020018061119d6030913960400191505060405180910390fd5b6105b58383836109a4565b505050565b60006105c58361071d565b821061060557604051600160e51b62461bcd02815260040180806020018281038252602a81526020018061114b602a913960400191505060405180910390fd5b6001600160a01b038316600090815260056020526040902080548390811061062957fe5b9060005260206000200154905092915050565b6105b583838360405180602001604052806000815250610857565b600061066161055b565b82106106a157604051600160e51b62461bcd02815260040180806020018281038252602b81526020018061124e602b913960400191505060405180910390fd5b600782815481106106ae57fe5b90600052602060002001549050919050565b6000818152600160205260408120546001600160a01b03168061071757604051600160e51b62461bcd0281526004018080602001828103825260288152602001806111756028913960400191505060405180910390fd5b92915050565b60006001600160a01b03821661076757604051600160e51b62461bcd0281526004018080602001828103825260298152602001806111fd6029913960400191505060405180910390fd5b6001600160a01b0382166000908152600360205260409020610717906109c3565b6001600160a01b0382163314156107e95760408051600160e51b62461bcd02815260206004820152601860248201527f4b495031373a20617070726f766520746f2063616c6c65720000000000000000604482015290519081900360640190fd5b3360008181526004602090815260408083206001600160a01b03871680855290835292819020805460ff1916861515908117909155815190815290519293927f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31929181900390910190a35050565b610862848484610562565b61086e848484846109c7565b6108ac57604051600160e51b62461bcd0281526004018080602001828103825260308152602001806111cd6030913960400191505060405180910390fd5b50505050565b6001600160a01b03918216600090815260046020908152604080832093909416825291909152205460ff1690565b6000908152600160205260409020546001600160a01b0316151590565b6000610908826108e0565b61094657604051600160e51b62461bcd02815260040180806020018281038252602b8152602001806112db602b913960400191505060405180910390fd5b6000610951836106c0565b9050806001600160a01b0316846001600160a01b0316148061098c5750836001600160a01b0316610981846103c9565b6001600160a01b0316145b8061099c575061099c81856108b2565b949350505050565b6109af838383610daa565b6109b98382610ef4565b6105b58282610fe9565b5490565b60008060606109de866001600160a01b0316611027565b6109ed5760019250505061099c565b856001600160a01b031663150b7a0260e01b3389888860405160240180856001600160a01b03166001600160a01b03168152602001846001600160a01b03166001600160a01b0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610a79578181015183820152602001610a61565b50505050905090810190601f168015610aa65780820380516001836020036101000a031916815260200191505b5060408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909a16999099178952518151919890975087965094509250829150849050835b60208310610b0e5780518252601f199092019160209182019101610aef565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114610b70576040519150601f19603f3d011682016040523d82523d6000602084013e610b75565b606091505b508051919350915015801590610bb557508051600160e11b630a85bd01029060208084019190811015610ba757600080fd5b50516001600160e01b031916145b15610bc55760019250505061099c565b856001600160a01b0316636745782b60e01b3389888860405160240180856001600160a01b03166001600160a01b03168152602001846001600160a01b03166001600160a01b0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610c51578181015183820152602001610c39565b50505050905090810190601f168015610c7e5780820380516001836020036101000a031916815260200191505b5060408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909a16999099178952518151919890975087965094509250829150849050835b60208310610ce65780518252601f199092019160209182019101610cc7565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114610d48576040519150601f19603f3d011682016040523d82523d6000602084013e610d4d565b606091505b508051919350915015801590610d8d57508051600160e01b636745782b029060208084019190811015610d7f57600080fd5b50516001600160e01b031916145b15610d9d5760019250505061099c565b5060009695505050505050565b826001600160a01b0316610dbd826106c0565b6001600160a01b031614610e0557604051600160e51b62461bcd0281526004018080602001828103825260288152602001806112266028913960400191505060405180910390fd5b6001600160a01b038216610e4d57604051600160e51b62461bcd0281526004018080602001828103825260238152602001806111286023913960400191505060405180910390fd5b610e568161102d565b6001600160a01b0383166000908152600360205260409020610e779061106a565b6001600160a01b0382166000908152600360205260409020610e9890611081565b60008181526001602052604080822080546001600160a01b0319166001600160a01b0386811691821790925591518493918716917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b6001600160a01b038216600090815260056020526040812054610f1e90600163ffffffff61108a16565b600083815260066020526040902054909150808214610fb9576001600160a01b0384166000908152600560205260408120805484908110610f5b57fe5b906000526020600020015490508060056000876001600160a01b03166001600160a01b031681526020019081526020016000208381548110610f9957fe5b600091825260208083209091019290925591825260069052604090208190555b6001600160a01b0384166000908152600560205260409020805490610fe29060001983016110ea565b5050505050565b6001600160a01b0390911660009081526005602081815260408084208054868652600684529185208290559282526001810183559183529091200155565b3b151590565b6000818152600260205260409020546001600160a01b03161561106757600081815260026020526040902080546001600160a01b03191690555b50565b805461107d90600163ffffffff61108a16565b9055565b80546001019055565b6000828211156110e45760408051600160e51b62461bcd02815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b8154818355818111156105b5576000838152602090206105b591810190830161055f91905b80821115611123576000815560010161110f565b509056fe4b495031373a207472616e7366657220746f20746865207a65726f20616464726573734b49503137456e756d657261626c653a206f776e657220696e646578206f7574206f6620626f756e64734b495031373a206f776e657220717565727920666f72206e6f6e6578697374656e7420746f6b656e4b495031373a207472616e736665722063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f7665644b495031373a207472616e7366657220746f206e6f6e204b49503137526563656976657220696d706c656d656e7465724b495031373a2062616c616e636520717565727920666f7220746865207a65726f20616464726573734b495031373a207472616e73666572206f6620746f6b656e2074686174206973206e6f74206f776e4b49503137456e756d657261626c653a20676c6f62616c20696e646578206f7574206f6620626f756e64734b495031373a20617070726f76656420717565727920666f72206e6f6e6578697374656e7420746f6b656e4b495031373a20617070726f76652063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f76656420666f7220616c6c4b495031373a206f70657261746f7220717565727920666f72206e6f6e6578697374656e7420746f6b656ea165627a7a7230582084a9aef7b39035ef5bd635b015444f9015ec7142d48b43aa878f9c449fb6e7030029`

// KIP17EnumerableFuncSigs maps the 4-byte function signature to its string representation.
var KIP17EnumerableFuncSigs = map[string]string{
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"081812fc": "getApproved(uint256)",
	"e985e9c5": "isApprovedForAll(address,address)",
	"6352211e": "ownerOf(uint256)",
	"42842e0e": "safeTransferFrom(address,address,uint256)",
	"b88d4fde": "safeTransferFrom(address,address,uint256,bytes)",
	"a22cb465": "setApprovalForAll(address,bool)",
	"01ffc9a7": "supportsInterface(bytes4)",
	"4f6ccce7": "tokenByIndex(uint256)",
	"2f745c59": "tokenOfOwnerByIndex(address,uint256)",
	"18160ddd": "totalSupply()",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// KIP17EnumerableBin is the compiled bytecode used for deploying new contracts.
var KIP17EnumerableBin = "0x608060405234801561001057600080fd5b506100276301ffc9a760e01b61005860201b60201c565b61003d6380ac58cd60e01b61005860201b60201c565b61005363780e9d6360e01b61005860201b60201c565b610126565b7fffffffff0000000000000000000000000000000000000000000000000000000080821614156100e957604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f4b495031333a20696e76616c696420696e746572666163652069640000000000604482015290519081900360640190fd5b7fffffffff00000000000000000000000000000000000000000000000000000000166000908152602081905260409020805460ff19166001179055565b611331806101356000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c806342842e0e1161008c57806370a082311161006657806370a0823114610262578063a22cb46514610288578063b88d4fde146102b6578063e985e9c51461037c576100cf565b806342842e0e146101f25780634f6ccce7146102285780636352211e14610245576100cf565b806301ffc9a7146100d4578063081812fc1461010f578063095ea7b31461014857806318160ddd1461017657806323b872dd146101905780632f745c59146101c6575b600080fd5b6100fb600480360360208110156100ea57600080fd5b50356001600160e01b0319166103aa565b604080519115158252519081900360200190f35b61012c6004803603602081101561012557600080fd5b50356103c9565b604080516001600160a01b039092168252519081900360200190f35b6101746004803603604081101561015e57600080fd5b506001600160a01b03813516906020013561042e565b005b61017e61055b565b60408051918252519081900360200190f35b610174600480360360608110156101a657600080fd5b506001600160a01b03813581169160208101359091169060400135610562565b61017e600480360360408110156101dc57600080fd5b506001600160a01b0381351690602001356105ba565b6101746004803603606081101561020857600080fd5b506001600160a01b0381358116916020810135909116906040013561063c565b61017e6004803603602081101561023e57600080fd5b5035610657565b61012c6004803603602081101561025b57600080fd5b50356106c0565b61017e6004803603602081101561027857600080fd5b50356001600160a01b031661071d565b6101746004803603604081101561029e57600080fd5b506001600160a01b0381351690602001351515610788565b610174600480360360808110156102cc57600080fd5b6001600160a01b0382358116926020810135909116916040820135919081019060808101606082013564010000000081111561030757600080fd5b82018360208201111561031957600080fd5b8035906020019184600183028401116401000000008311171561033b57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610857945050505050565b6100fb6004803603604081101561039257600080fd5b506001600160a01b03813581169160200135166108b2565b6001600160e01b03191660009081526020819052604090205460ff1690565b60006103d4826108e0565b61041257604051600160e51b62461bcd02815260040180806020018281038252602b815260200180611279602b913960400191505060405180910390fd5b506000908152600260205260409020546001600160a01b031690565b6000610439826106c0565b9050806001600160a01b0316836001600160a01b031614156104a55760408051600160e51b62461bcd02815260206004820181905260248201527f4b495031373a20617070726f76616c20746f2063757272656e74206f776e6572604482015290519081900360640190fd5b336001600160a01b03821614806104c157506104c181336108b2565b6104ff57604051600160e51b62461bcd0281526004018080602001828103825260378152602001806112a46037913960400191505060405180910390fd5b60008281526002602052604080822080546001600160a01b0319166001600160a01b0387811691821790925591518593918516917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591a4505050565b6007545b90565b61056c33826108fd565b6105aa57604051600160e51b62461bcd02815260040180806020018281038252603081526020018061119d6030913960400191505060405180910390fd5b6105b58383836109a4565b505050565b60006105c58361071d565b821061060557604051600160e51b62461bcd02815260040180806020018281038252602a81526020018061114b602a913960400191505060405180910390fd5b6001600160a01b038316600090815260056020526040902080548390811061062957fe5b9060005260206000200154905092915050565b6105b583838360405180602001604052806000815250610857565b600061066161055b565b82106106a157604051600160e51b62461bcd02815260040180806020018281038252602b81526020018061124e602b913960400191505060405180910390fd5b600782815481106106ae57fe5b90600052602060002001549050919050565b6000818152600160205260408120546001600160a01b03168061071757604051600160e51b62461bcd0281526004018080602001828103825260288152602001806111756028913960400191505060405180910390fd5b92915050565b60006001600160a01b03821661076757604051600160e51b62461bcd0281526004018080602001828103825260298152602001806111fd6029913960400191505060405180910390fd5b6001600160a01b0382166000908152600360205260409020610717906109c3565b6001600160a01b0382163314156107e95760408051600160e51b62461bcd02815260206004820152601860248201527f4b495031373a20617070726f766520746f2063616c6c65720000000000000000604482015290519081900360640190fd5b3360008181526004602090815260408083206001600160a01b03871680855290835292819020805460ff1916861515908117909155815190815290519293927f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31929181900390910190a35050565b610862848484610562565b61086e848484846109c7565b6108ac57604051600160e51b62461bcd0281526004018080602001828103825260308152602001806111cd6030913960400191505060405180910390fd5b50505050565b6001600160a01b03918216600090815260046020908152604080832093909416825291909152205460ff1690565b6000908152600160205260409020546001600160a01b0316151590565b6000610908826108e0565b61094657604051600160e51b62461bcd02815260040180806020018281038252602b8152602001806112db602b913960400191505060405180910390fd5b6000610951836106c0565b9050806001600160a01b0316846001600160a01b0316148061098c5750836001600160a01b0316610981846103c9565b6001600160a01b0316145b8061099c575061099c81856108b2565b949350505050565b6109af838383610daa565b6109b98382610ef4565b6105b58282610fe9565b5490565b60008060606109de866001600160a01b0316611027565b6109ed5760019250505061099c565b856001600160a01b031663150b7a0260e01b3389888860405160240180856001600160a01b03166001600160a01b03168152602001846001600160a01b03166001600160a01b0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610a79578181015183820152602001610a61565b50505050905090810190601f168015610aa65780820380516001836020036101000a031916815260200191505b5060408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909a16999099178952518151919890975087965094509250829150849050835b60208310610b0e5780518252601f199092019160209182019101610aef565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114610b70576040519150601f19603f3d011682016040523d82523d6000602084013e610b75565b606091505b508051919350915015801590610bb557508051600160e11b630a85bd01029060208084019190811015610ba757600080fd5b50516001600160e01b031916145b15610bc55760019250505061099c565b856001600160a01b0316636745782b60e01b3389888860405160240180856001600160a01b03166001600160a01b03168152602001846001600160a01b03166001600160a01b0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610c51578181015183820152602001610c39565b50505050905090810190601f168015610c7e5780820380516001836020036101000a031916815260200191505b5060408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909a16999099178952518151919890975087965094509250829150849050835b60208310610ce65780518252601f199092019160209182019101610cc7565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114610d48576040519150601f19603f3d011682016040523d82523d6000602084013e610d4d565b606091505b508051919350915015801590610d8d57508051600160e01b636745782b029060208084019190811015610d7f57600080fd5b50516001600160e01b031916145b15610d9d5760019250505061099c565b5060009695505050505050565b826001600160a01b0316610dbd826106c0565b6001600160a01b031614610e0557604051600160e51b62461bcd0281526004018080602001828103825260288152602001806112266028913960400191505060405180910390fd5b6001600160a01b038216610e4d57604051600160e51b62461bcd0281526004018080602001828103825260238152602001806111286023913960400191505060405180910390fd5b610e568161102d565b6001600160a01b0383166000908152600360205260409020610e779061106a565b6001600160a01b0382166000908152600360205260409020610e9890611081565b60008181526001602052604080822080546001600160a01b0319166001600160a01b0386811691821790925591518493918716917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b6001600160a01b038216600090815260056020526040812054610f1e90600163ffffffff61108a16565b600083815260066020526040902054909150808214610fb9576001600160a01b0384166000908152600560205260408120805484908110610f5b57fe5b906000526020600020015490508060056000876001600160a01b03166001600160a01b031681526020019081526020016000208381548110610f9957fe5b600091825260208083209091019290925591825260069052604090208190555b6001600160a01b0384166000908152600560205260409020805490610fe29060001983016110ea565b5050505050565b6001600160a01b0390911660009081526005602081815260408084208054868652600684529185208290559282526001810183559183529091200155565b3b151590565b6000818152600260205260409020546001600160a01b03161561106757600081815260026020526040902080546001600160a01b03191690555b50565b805461107d90600163ffffffff61108a16565b9055565b80546001019055565b6000828211156110e45760408051600160e51b62461bcd02815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b8154818355818111156105b5576000838152602090206105b591810190830161055f91905b80821115611123576000815560010161110f565b509056fe4b495031373a207472616e7366657220746f20746865207a65726f20616464726573734b49503137456e756d657261626c653a206f776e657220696e646578206f7574206f6620626f756e64734b495031373a206f776e657220717565727920666f72206e6f6e6578697374656e7420746f6b656e4b495031373a207472616e736665722063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f7665644b495031373a207472616e7366657220746f206e6f6e204b49503137526563656976657220696d706c656d656e7465724b495031373a2062616c616e636520717565727920666f7220746865207a65726f20616464726573734b495031373a207472616e73666572206f6620746f6b656e2074686174206973206e6f74206f776e4b49503137456e756d657261626c653a20676c6f62616c20696e646578206f7574206f6620626f756e64734b495031373a20617070726f76656420717565727920666f72206e6f6e6578697374656e7420746f6b656e4b495031373a20617070726f76652063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f76656420666f7220616c6c4b495031373a206f70657261746f7220717565727920666f72206e6f6e6578697374656e7420746f6b656ea165627a7a7230582084a9aef7b39035ef5bd635b015444f9015ec7142d48b43aa878f9c449fb6e7030029"

// DeployKIP17Enumerable deploys a new Klaytn contract, binding an instance of KIP17Enumerable to it.
func DeployKIP17Enumerable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *KIP17Enumerable, error) {
	parsed, err := abi.JSON(strings.NewReader(KIP17EnumerableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(KIP17EnumerableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KIP17Enumerable{KIP17EnumerableCaller: KIP17EnumerableCaller{contract: contract}, KIP17EnumerableTransactor: KIP17EnumerableTransactor{contract: contract}, KIP17EnumerableFilterer: KIP17EnumerableFilterer{contract: contract}}, nil
}

// KIP17Enumerable is an auto generated Go binding around a Klaytn contract.
type KIP17Enumerable struct {
	KIP17EnumerableCaller     // Read-only binding to the contract
	KIP17EnumerableTransactor // Write-only binding to the contract
	KIP17EnumerableFilterer   // Log filterer for contract events
}

// KIP17EnumerableCaller is an auto generated read-only Go binding around a Klaytn contract.
type KIP17EnumerableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KIP17EnumerableTransactor is an auto generated write-only Go binding around a Klaytn contract.
type KIP17EnumerableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KIP17EnumerableFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type KIP17EnumerableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KIP17EnumerableSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type KIP17EnumerableSession struct {
	Contract     *KIP17Enumerable  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KIP17EnumerableCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type KIP17EnumerableCallerSession struct {
	Contract *KIP17EnumerableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// KIP17EnumerableTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type KIP17EnumerableTransactorSession struct {
	Contract     *KIP17EnumerableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// KIP17EnumerableRaw is an auto generated low-level Go binding around a Klaytn contract.
type KIP17EnumerableRaw struct {
	Contract *KIP17Enumerable // Generic contract binding to access the raw methods on
}

// KIP17EnumerableCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type KIP17EnumerableCallerRaw struct {
	Contract *KIP17EnumerableCaller // Generic read-only contract binding to access the raw methods on
}

// KIP17EnumerableTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type KIP17EnumerableTransactorRaw struct {
	Contract *KIP17EnumerableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKIP17Enumerable creates a new instance of KIP17Enumerable, bound to a specific deployed contract.
func NewKIP17Enumerable(address common.Address, backend bind.ContractBackend) (*KIP17Enumerable, error) {
	contract, err := bindKIP17Enumerable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KIP17Enumerable{KIP17EnumerableCaller: KIP17EnumerableCaller{contract: contract}, KIP17EnumerableTransactor: KIP17EnumerableTransactor{contract: contract}, KIP17EnumerableFilterer: KIP17EnumerableFilterer{contract: contract}}, nil
}

// NewKIP17EnumerableCaller creates a new read-only instance of KIP17Enumerable, bound to a specific deployed contract.
func NewKIP17EnumerableCaller(address common.Address, caller bind.ContractCaller) (*KIP17EnumerableCaller, error) {
	contract, err := bindKIP17Enumerable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KIP17EnumerableCaller{contract: contract}, nil
}

// NewKIP17EnumerableTransactor creates a new write-only instance of KIP17Enumerable, bound to a specific deployed contract.
func NewKIP17EnumerableTransactor(address common.Address, transactor bind.ContractTransactor) (*KIP17EnumerableTransactor, error) {
	contract, err := bindKIP17Enumerable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KIP17EnumerableTransactor{contract: contract}, nil
}

// NewKIP17EnumerableFilterer creates a new log filterer instance of KIP17Enumerable, bound to a specific deployed contract.
func NewKIP17EnumerableFilterer(address common.Address, filterer bind.ContractFilterer) (*KIP17EnumerableFilterer, error) {
	contract, err := bindKIP17Enumerable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KIP17EnumerableFilterer{contract: contract}, nil
}

// bindKIP17Enumerable binds a generic wrapper to an already deployed contract.
func bindKIP17Enumerable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KIP17EnumerableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KIP17Enumerable *KIP17EnumerableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KIP17Enumerable.Contract.KIP17EnumerableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KIP17Enumerable *KIP17EnumerableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KIP17Enumerable.Contract.KIP17EnumerableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KIP17Enumerable *KIP17EnumerableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KIP17Enumerable.Contract.KIP17EnumerableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KIP17Enumerable *KIP17EnumerableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KIP17Enumerable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KIP17Enumerable *KIP17EnumerableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KIP17Enumerable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KIP17Enumerable *KIP17EnumerableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KIP17Enumerable.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_KIP17Enumerable *KIP17EnumerableCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KIP17Enumerable.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_KIP17Enumerable *KIP17EnumerableSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _KIP17Enumerable.Contract.BalanceOf(&_KIP17Enumerable.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_KIP17Enumerable *KIP17EnumerableCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _KIP17Enumerable.Contract.BalanceOf(&_KIP17Enumerable.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_KIP17Enumerable *KIP17EnumerableCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KIP17Enumerable.contract.Call(opts, out, "getApproved", tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_KIP17Enumerable *KIP17EnumerableSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _KIP17Enumerable.Contract.GetApproved(&_KIP17Enumerable.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_KIP17Enumerable *KIP17EnumerableCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _KIP17Enumerable.Contract.GetApproved(&_KIP17Enumerable.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_KIP17Enumerable *KIP17EnumerableCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KIP17Enumerable.contract.Call(opts, out, "isApprovedForAll", owner, operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_KIP17Enumerable *KIP17EnumerableSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _KIP17Enumerable.Contract.IsApprovedForAll(&_KIP17Enumerable.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_KIP17Enumerable *KIP17EnumerableCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _KIP17Enumerable.Contract.IsApprovedForAll(&_KIP17Enumerable.CallOpts, owner, operator)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_KIP17Enumerable *KIP17EnumerableCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KIP17Enumerable.contract.Call(opts, out, "ownerOf", tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_KIP17Enumerable *KIP17EnumerableSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _KIP17Enumerable.Contract.OwnerOf(&_KIP17Enumerable.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_KIP17Enumerable *KIP17EnumerableCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _KIP17Enumerable.Contract.OwnerOf(&_KIP17Enumerable.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_KIP17Enumerable *KIP17EnumerableCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KIP17Enumerable.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_KIP17Enumerable *KIP17EnumerableSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _KIP17Enumerable.Contract.SupportsInterface(&_KIP17Enumerable.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_KIP17Enumerable *KIP17EnumerableCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _KIP17Enumerable.Contract.SupportsInterface(&_KIP17Enumerable.CallOpts, interfaceId)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_KIP17Enumerable *KIP17EnumerableCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KIP17Enumerable.contract.Call(opts, out, "tokenByIndex", index)
	return *ret0, err
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_KIP17Enumerable *KIP17EnumerableSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _KIP17Enumerable.Contract.TokenByIndex(&_KIP17Enumerable.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_KIP17Enumerable *KIP17EnumerableCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _KIP17Enumerable.Contract.TokenByIndex(&_KIP17Enumerable.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_KIP17Enumerable *KIP17EnumerableCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KIP17Enumerable.contract.Call(opts, out, "tokenOfOwnerByIndex", owner, index)
	return *ret0, err
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_KIP17Enumerable *KIP17EnumerableSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _KIP17Enumerable.Contract.TokenOfOwnerByIndex(&_KIP17Enumerable.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_KIP17Enumerable *KIP17EnumerableCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _KIP17Enumerable.Contract.TokenOfOwnerByIndex(&_KIP17Enumerable.CallOpts, owner, index)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_KIP17Enumerable *KIP17EnumerableCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KIP17Enumerable.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_KIP17Enumerable *KIP17EnumerableSession) TotalSupply() (*big.Int, error) {
	return _KIP17Enumerable.Contract.TotalSupply(&_KIP17Enumerable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_KIP17Enumerable *KIP17EnumerableCallerSession) TotalSupply() (*big.Int, error) {
	return _KIP17Enumerable.Contract.TotalSupply(&_KIP17Enumerable.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_KIP17Enumerable *KIP17EnumerableTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17Enumerable.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_KIP17Enumerable *KIP17EnumerableSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17Enumerable.Contract.Approve(&_KIP17Enumerable.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_KIP17Enumerable *KIP17EnumerableTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17Enumerable.Contract.Approve(&_KIP17Enumerable.TransactOpts, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_KIP17Enumerable *KIP17EnumerableTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17Enumerable.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_KIP17Enumerable *KIP17EnumerableSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17Enumerable.Contract.SafeTransferFrom(&_KIP17Enumerable.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_KIP17Enumerable *KIP17EnumerableTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17Enumerable.Contract.SafeTransferFrom(&_KIP17Enumerable.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_KIP17Enumerable *KIP17EnumerableTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _KIP17Enumerable.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_KIP17Enumerable *KIP17EnumerableSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _KIP17Enumerable.Contract.SafeTransferFrom0(&_KIP17Enumerable.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_KIP17Enumerable *KIP17EnumerableTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _KIP17Enumerable.Contract.SafeTransferFrom0(&_KIP17Enumerable.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_KIP17Enumerable *KIP17EnumerableTransactor) SetApprovalForAll(opts *bind.TransactOpts, to common.Address, approved bool) (*types.Transaction, error) {
	return _KIP17Enumerable.contract.Transact(opts, "setApprovalForAll", to, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_KIP17Enumerable *KIP17EnumerableSession) SetApprovalForAll(to common.Address, approved bool) (*types.Transaction, error) {
	return _KIP17Enumerable.Contract.SetApprovalForAll(&_KIP17Enumerable.TransactOpts, to, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_KIP17Enumerable *KIP17EnumerableTransactorSession) SetApprovalForAll(to common.Address, approved bool) (*types.Transaction, error) {
	return _KIP17Enumerable.Contract.SetApprovalForAll(&_KIP17Enumerable.TransactOpts, to, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_KIP17Enumerable *KIP17EnumerableTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17Enumerable.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_KIP17Enumerable *KIP17EnumerableSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17Enumerable.Contract.TransferFrom(&_KIP17Enumerable.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_KIP17Enumerable *KIP17EnumerableTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17Enumerable.Contract.TransferFrom(&_KIP17Enumerable.TransactOpts, from, to, tokenId)
}

// KIP17EnumerableApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the KIP17Enumerable contract.
type KIP17EnumerableApprovalIterator struct {
	Event *KIP17EnumerableApproval // Event containing the contract specifics and raw log

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
func (it *KIP17EnumerableApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KIP17EnumerableApproval)
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
		it.Event = new(KIP17EnumerableApproval)
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
func (it *KIP17EnumerableApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KIP17EnumerableApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KIP17EnumerableApproval represents a Approval event raised by the KIP17Enumerable contract.
type KIP17EnumerableApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_KIP17Enumerable *KIP17EnumerableFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*KIP17EnumerableApprovalIterator, error) {

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

	logs, sub, err := _KIP17Enumerable.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &KIP17EnumerableApprovalIterator{contract: _KIP17Enumerable.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_KIP17Enumerable *KIP17EnumerableFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *KIP17EnumerableApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _KIP17Enumerable.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KIP17EnumerableApproval)
				if err := _KIP17Enumerable.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_KIP17Enumerable *KIP17EnumerableFilterer) ParseApproval(log types.Log) (*KIP17EnumerableApproval, error) {
	event := new(KIP17EnumerableApproval)
	if err := _KIP17Enumerable.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KIP17EnumerableApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the KIP17Enumerable contract.
type KIP17EnumerableApprovalForAllIterator struct {
	Event *KIP17EnumerableApprovalForAll // Event containing the contract specifics and raw log

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
func (it *KIP17EnumerableApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KIP17EnumerableApprovalForAll)
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
		it.Event = new(KIP17EnumerableApprovalForAll)
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
func (it *KIP17EnumerableApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KIP17EnumerableApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KIP17EnumerableApprovalForAll represents a ApprovalForAll event raised by the KIP17Enumerable contract.
type KIP17EnumerableApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_KIP17Enumerable *KIP17EnumerableFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*KIP17EnumerableApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _KIP17Enumerable.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &KIP17EnumerableApprovalForAllIterator{contract: _KIP17Enumerable.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_KIP17Enumerable *KIP17EnumerableFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *KIP17EnumerableApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _KIP17Enumerable.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KIP17EnumerableApprovalForAll)
				if err := _KIP17Enumerable.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_KIP17Enumerable *KIP17EnumerableFilterer) ParseApprovalForAll(log types.Log) (*KIP17EnumerableApprovalForAll, error) {
	event := new(KIP17EnumerableApprovalForAll)
	if err := _KIP17Enumerable.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KIP17EnumerableTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the KIP17Enumerable contract.
type KIP17EnumerableTransferIterator struct {
	Event *KIP17EnumerableTransfer // Event containing the contract specifics and raw log

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
func (it *KIP17EnumerableTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KIP17EnumerableTransfer)
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
		it.Event = new(KIP17EnumerableTransfer)
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
func (it *KIP17EnumerableTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KIP17EnumerableTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KIP17EnumerableTransfer represents a Transfer event raised by the KIP17Enumerable contract.
type KIP17EnumerableTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_KIP17Enumerable *KIP17EnumerableFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*KIP17EnumerableTransferIterator, error) {

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

	logs, sub, err := _KIP17Enumerable.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &KIP17EnumerableTransferIterator{contract: _KIP17Enumerable.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_KIP17Enumerable *KIP17EnumerableFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *KIP17EnumerableTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _KIP17Enumerable.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KIP17EnumerableTransfer)
				if err := _KIP17Enumerable.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_KIP17Enumerable *KIP17EnumerableFilterer) ParseTransfer(log types.Log) (*KIP17EnumerableTransfer, error) {
	event := new(KIP17EnumerableTransfer)
	if err := _KIP17Enumerable.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KIP17FullABI is the input ABI used to generate the binding from.
const KIP17FullABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"}]"

// KIP17FullBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const KIP17FullBinRuntime = `608060405234801561001057600080fd5b50600436106101005760003560e01c80634f6ccce711610097578063a22cb46511610066578063a22cb4651461033e578063b88d4fde1461036c578063c87b56dd14610432578063e985e9c51461044f57610100565b80634f6ccce7146102d65780636352211e146102f357806370a082311461031057806395d89b411461033657610100565b806318160ddd116100d357806318160ddd1461022457806323b872dd1461023e5780632f745c591461027457806342842e0e146102a057610100565b806301ffc9a71461010557806306fdde0314610140578063081812fc146101bd578063095ea7b3146101f6575b600080fd5b61012c6004803603602081101561011b57600080fd5b50356001600160e01b03191661047d565b604080519115158252519081900360200190f35b61014861049c565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561018257818101518382015260200161016a565b50505050905090810190601f1680156101af5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101da600480360360208110156101d357600080fd5b5035610533565b604080516001600160a01b039092168252519081900360200190f35b6102226004803603604081101561020c57600080fd5b506001600160a01b038135169060200135610598565b005b61022c6106c5565b60408051918252519081900360200190f35b6102226004803603606081101561025457600080fd5b506001600160a01b038135811691602081013590911690604001356106cb565b61022c6004803603604081101561028a57600080fd5b506001600160a01b038135169060200135610723565b610222600480360360608110156102b657600080fd5b506001600160a01b038135811691602081013590911690604001356107a5565b61022c600480360360208110156102ec57600080fd5b50356107c0565b6101da6004803603602081101561030957600080fd5b5035610829565b61022c6004803603602081101561032657600080fd5b50356001600160a01b0316610886565b6101486108f1565b6102226004803603604081101561035457600080fd5b506001600160a01b0381351690602001351515610952565b6102226004803603608081101561038257600080fd5b6001600160a01b038235811692602081013590911691604082013591908101906080810160608201356401000000008111156103bd57600080fd5b8201836020820111156103cf57600080fd5b803590602001918460018302840111640100000000831117156103f157600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610a21945050505050565b6101486004803603602081101561044857600080fd5b5035610a7c565b61012c6004803603604081101561046557600080fd5b506001600160a01b0381358116916020013516610b64565b6001600160e01b03191660009081526020819052604090205460ff1690565b60098054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156105285780601f106104fd57610100808354040283529160200191610528565b820191906000526020600020905b81548152906001019060200180831161050b57829003601f168201915b505050505090505b90565b600061053e82610b92565b61057c57604051600160e51b62461bcd02815260040180806020018281038252602b815260200180611559602b913960400191505060405180910390fd5b506000908152600260205260409020546001600160a01b031690565b60006105a382610829565b9050806001600160a01b0316836001600160a01b0316141561060f5760408051600160e51b62461bcd02815260206004820181905260248201527f4b495031373a20617070726f76616c20746f2063757272656e74206f776e6572604482015290519081900360640190fd5b336001600160a01b038216148061062b575061062b8133610b64565b61066957604051600160e51b62461bcd0281526004018080602001828103825260378152602001806115846037913960400191505060405180910390fd5b60008281526002602052604080822080546001600160a01b0319166001600160a01b0387811691821790925591518593918516917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591a4505050565b60075490565b6106d53382610baf565b61071357604051600160e51b62461bcd02815260040180806020018281038252603081526020018061147d6030913960400191505060405180910390fd5b61071e838383610c56565b505050565b600061072e83610886565b821061076e57604051600160e51b62461bcd02815260040180806020018281038252602a81526020018061142b602a913960400191505060405180910390fd5b6001600160a01b038316600090815260056020526040902080548390811061079257fe5b9060005260206000200154905092915050565b61071e83838360405180602001604052806000815250610a21565b60006107ca6106c5565b821061080a57604051600160e51b62461bcd02815260040180806020018281038252602b81526020018061152e602b913960400191505060405180910390fd5b6007828154811061081757fe5b90600052602060002001549050919050565b6000818152600160205260408120546001600160a01b03168061088057604051600160e51b62461bcd0281526004018080602001828103825260288152602001806114556028913960400191505060405180910390fd5b92915050565b60006001600160a01b0382166108d057604051600160e51b62461bcd0281526004018080602001828103825260298152602001806114dd6029913960400191505060405180910390fd5b6001600160a01b038216600090815260036020526040902061088090610c75565b600a8054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156105285780601f106104fd57610100808354040283529160200191610528565b6001600160a01b0382163314156109b35760408051600160e51b62461bcd02815260206004820152601860248201527f4b495031373a20617070726f766520746f2063616c6c65720000000000000000604482015290519081900360640190fd5b3360008181526004602090815260408083206001600160a01b03871680855290835292819020805460ff1916861515908117909155815190815290519293927f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31929181900390910190a35050565b610a2c8484846106cb565b610a3884848484610c79565b610a7657604051600160e51b62461bcd0281526004018080602001828103825260308152602001806114ad6030913960400191505060405180910390fd5b50505050565b6060610a8782610b92565b610ac557604051600160e51b62461bcd02815260040180806020018281038252602e8152602001806113da602e913960400191505060405180910390fd5b6000828152600b602090815260409182902080548351601f600260001961010060018616150201909316929092049182018490048402810184019094528084529091830182828015610b585780601f10610b2d57610100808354040283529160200191610b58565b820191906000526020600020905b815481529060010190602001808311610b3b57829003601f168201915b50505050509050919050565b6001600160a01b03918216600090815260046020908152604080832093909416825291909152205460ff1690565b6000908152600160205260409020546001600160a01b0316151590565b6000610bba82610b92565b610bf857604051600160e51b62461bcd02815260040180806020018281038252602b8152602001806115bb602b913960400191505060405180910390fd5b6000610c0383610829565b9050806001600160a01b0316846001600160a01b03161480610c3e5750836001600160a01b0316610c3384610533565b6001600160a01b0316145b80610c4e5750610c4e8185610b64565b949350505050565b610c6183838361105c565b610c6b83826111a6565b61071e828261129b565b5490565b6000806060610c90866001600160a01b03166112d9565b610c9f57600192505050610c4e565b856001600160a01b031663150b7a0260e01b3389888860405160240180856001600160a01b03166001600160a01b03168152602001846001600160a01b03166001600160a01b0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610d2b578181015183820152602001610d13565b50505050905090810190601f168015610d585780820380516001836020036101000a031916815260200191505b5060408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909a16999099178952518151919890975087965094509250829150849050835b60208310610dc05780518252601f199092019160209182019101610da1565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114610e22576040519150601f19603f3d011682016040523d82523d6000602084013e610e27565b606091505b508051919350915015801590610e6757508051600160e11b630a85bd01029060208084019190811015610e5957600080fd5b50516001600160e01b031916145b15610e7757600192505050610c4e565b856001600160a01b0316636745782b60e01b3389888860405160240180856001600160a01b03166001600160a01b03168152602001846001600160a01b03166001600160a01b0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610f03578181015183820152602001610eeb565b50505050905090810190601f168015610f305780820380516001836020036101000a031916815260200191505b5060408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909a16999099178952518151919890975087965094509250829150849050835b60208310610f985780518252601f199092019160209182019101610f79565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114610ffa576040519150601f19603f3d011682016040523d82523d6000602084013e610fff565b606091505b50805191935091501580159061103f57508051600160e01b636745782b02906020808401919081101561103157600080fd5b50516001600160e01b031916145b1561104f57600192505050610c4e565b5060009695505050505050565b826001600160a01b031661106f82610829565b6001600160a01b0316146110b757604051600160e51b62461bcd0281526004018080602001828103825260288152602001806115066028913960400191505060405180910390fd5b6001600160a01b0382166110ff57604051600160e51b62461bcd0281526004018080602001828103825260238152602001806114086023913960400191505060405180910390fd5b611108816112df565b6001600160a01b03831660009081526003602052604090206111299061131c565b6001600160a01b038216600090815260036020526040902061114a90611333565b60008181526001602052604080822080546001600160a01b0319166001600160a01b0386811691821790925591518493918716917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b6001600160a01b0382166000908152600560205260408120546111d090600163ffffffff61133c16565b60008381526006602052604090205490915080821461126b576001600160a01b038416600090815260056020526040812080548490811061120d57fe5b906000526020600020015490508060056000876001600160a01b03166001600160a01b03168152602001908152602001600020838154811061124b57fe5b600091825260208083209091019290925591825260069052604090208190555b6001600160a01b038416600090815260056020526040902080549061129490600019830161139c565b5050505050565b6001600160a01b0390911660009081526005602081815260408084208054868652600684529185208290559282526001810183559183529091200155565b3b151590565b6000818152600260205260409020546001600160a01b03161561131957600081815260026020526040902080546001600160a01b03191690555b50565b805461132f90600163ffffffff61133c16565b9055565b80546001019055565b6000828211156113965760408051600160e51b62461bcd02815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b81548183558181111561071e5760008381526020902061071e91810190830161053091905b808211156113d557600081556001016113c1565b509056fe4b495031374d657461646174613a2055524920717565727920666f72206e6f6e6578697374656e7420746f6b656e4b495031373a207472616e7366657220746f20746865207a65726f20616464726573734b49503137456e756d657261626c653a206f776e657220696e646578206f7574206f6620626f756e64734b495031373a206f776e657220717565727920666f72206e6f6e6578697374656e7420746f6b656e4b495031373a207472616e736665722063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f7665644b495031373a207472616e7366657220746f206e6f6e204b49503137526563656976657220696d706c656d656e7465724b495031373a2062616c616e636520717565727920666f7220746865207a65726f20616464726573734b495031373a207472616e73666572206f6620746f6b656e2074686174206973206e6f74206f776e4b49503137456e756d657261626c653a20676c6f62616c20696e646578206f7574206f6620626f756e64734b495031373a20617070726f76656420717565727920666f72206e6f6e6578697374656e7420746f6b656e4b495031373a20617070726f76652063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f76656420666f7220616c6c4b495031373a206f70657261746f7220717565727920666f72206e6f6e6578697374656e7420746f6b656ea165627a7a72305820ed7b15db174cbbafb0ffe4d8598c28d541404a3c946362ff10e4dca4b285d3750029`

// KIP17FullFuncSigs maps the 4-byte function signature to its string representation.
var KIP17FullFuncSigs = map[string]string{
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"081812fc": "getApproved(uint256)",
	"e985e9c5": "isApprovedForAll(address,address)",
	"06fdde03": "name()",
	"6352211e": "ownerOf(uint256)",
	"42842e0e": "safeTransferFrom(address,address,uint256)",
	"b88d4fde": "safeTransferFrom(address,address,uint256,bytes)",
	"a22cb465": "setApprovalForAll(address,bool)",
	"01ffc9a7": "supportsInterface(bytes4)",
	"95d89b41": "symbol()",
	"4f6ccce7": "tokenByIndex(uint256)",
	"2f745c59": "tokenOfOwnerByIndex(address,uint256)",
	"c87b56dd": "tokenURI(uint256)",
	"18160ddd": "totalSupply()",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// KIP17FullBin is the compiled bytecode used for deploying new contracts.
var KIP17FullBin = "0x60806040523480156200001157600080fd5b506040516200190038038062001900833981018060405260408110156200003757600080fd5b8101908080516401000000008111156200005057600080fd5b820160208101848111156200006457600080fd5b81516401000000008111828201871017156200007f57600080fd5b505092919060200180516401000000008111156200009c57600080fd5b82016020810184811115620000b057600080fd5b8151640100000000811182820187101715620000cb57600080fd5b50509291905050508181620000ed6301ffc9a760e01b6200016b60201b60201c565b620001056380ac58cd60e01b6200016b60201b60201c565b6200011d63780e9d6360e01b6200016b60201b60201c565b8151620001329060099060208501906200023a565b5080516200014890600a9060208401906200023a565b5062000161635b5e139f60e01b6200016b60201b60201c565b50505050620002df565b7fffffffff000000000000000000000000000000000000000000000000000000008082161415620001fd57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f4b495031333a20696e76616c696420696e746572666163652069640000000000604482015290519081900360640190fd5b7fffffffff00000000000000000000000000000000000000000000000000000000166000908152602081905260409020805460ff19166001179055565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200027d57805160ff1916838001178555620002ad565b82800160010185558215620002ad579182015b82811115620002ad57825182559160200191906001019062000290565b50620002bb929150620002bf565b5090565b620002dc91905b80821115620002bb5760008155600101620002c6565b90565b61161180620002ef6000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c80634f6ccce711610097578063a22cb46511610066578063a22cb4651461033e578063b88d4fde1461036c578063c87b56dd14610432578063e985e9c51461044f57610100565b80634f6ccce7146102d65780636352211e146102f357806370a082311461031057806395d89b411461033657610100565b806318160ddd116100d357806318160ddd1461022457806323b872dd1461023e5780632f745c591461027457806342842e0e146102a057610100565b806301ffc9a71461010557806306fdde0314610140578063081812fc146101bd578063095ea7b3146101f6575b600080fd5b61012c6004803603602081101561011b57600080fd5b50356001600160e01b03191661047d565b604080519115158252519081900360200190f35b61014861049c565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561018257818101518382015260200161016a565b50505050905090810190601f1680156101af5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101da600480360360208110156101d357600080fd5b5035610533565b604080516001600160a01b039092168252519081900360200190f35b6102226004803603604081101561020c57600080fd5b506001600160a01b038135169060200135610598565b005b61022c6106c5565b60408051918252519081900360200190f35b6102226004803603606081101561025457600080fd5b506001600160a01b038135811691602081013590911690604001356106cb565b61022c6004803603604081101561028a57600080fd5b506001600160a01b038135169060200135610723565b610222600480360360608110156102b657600080fd5b506001600160a01b038135811691602081013590911690604001356107a5565b61022c600480360360208110156102ec57600080fd5b50356107c0565b6101da6004803603602081101561030957600080fd5b5035610829565b61022c6004803603602081101561032657600080fd5b50356001600160a01b0316610886565b6101486108f1565b6102226004803603604081101561035457600080fd5b506001600160a01b0381351690602001351515610952565b6102226004803603608081101561038257600080fd5b6001600160a01b038235811692602081013590911691604082013591908101906080810160608201356401000000008111156103bd57600080fd5b8201836020820111156103cf57600080fd5b803590602001918460018302840111640100000000831117156103f157600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610a21945050505050565b6101486004803603602081101561044857600080fd5b5035610a7c565b61012c6004803603604081101561046557600080fd5b506001600160a01b0381358116916020013516610b64565b6001600160e01b03191660009081526020819052604090205460ff1690565b60098054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156105285780601f106104fd57610100808354040283529160200191610528565b820191906000526020600020905b81548152906001019060200180831161050b57829003601f168201915b505050505090505b90565b600061053e82610b92565b61057c57604051600160e51b62461bcd02815260040180806020018281038252602b815260200180611559602b913960400191505060405180910390fd5b506000908152600260205260409020546001600160a01b031690565b60006105a382610829565b9050806001600160a01b0316836001600160a01b0316141561060f5760408051600160e51b62461bcd02815260206004820181905260248201527f4b495031373a20617070726f76616c20746f2063757272656e74206f776e6572604482015290519081900360640190fd5b336001600160a01b038216148061062b575061062b8133610b64565b61066957604051600160e51b62461bcd0281526004018080602001828103825260378152602001806115846037913960400191505060405180910390fd5b60008281526002602052604080822080546001600160a01b0319166001600160a01b0387811691821790925591518593918516917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591a4505050565b60075490565b6106d53382610baf565b61071357604051600160e51b62461bcd02815260040180806020018281038252603081526020018061147d6030913960400191505060405180910390fd5b61071e838383610c56565b505050565b600061072e83610886565b821061076e57604051600160e51b62461bcd02815260040180806020018281038252602a81526020018061142b602a913960400191505060405180910390fd5b6001600160a01b038316600090815260056020526040902080548390811061079257fe5b9060005260206000200154905092915050565b61071e83838360405180602001604052806000815250610a21565b60006107ca6106c5565b821061080a57604051600160e51b62461bcd02815260040180806020018281038252602b81526020018061152e602b913960400191505060405180910390fd5b6007828154811061081757fe5b90600052602060002001549050919050565b6000818152600160205260408120546001600160a01b03168061088057604051600160e51b62461bcd0281526004018080602001828103825260288152602001806114556028913960400191505060405180910390fd5b92915050565b60006001600160a01b0382166108d057604051600160e51b62461bcd0281526004018080602001828103825260298152602001806114dd6029913960400191505060405180910390fd5b6001600160a01b038216600090815260036020526040902061088090610c75565b600a8054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156105285780601f106104fd57610100808354040283529160200191610528565b6001600160a01b0382163314156109b35760408051600160e51b62461bcd02815260206004820152601860248201527f4b495031373a20617070726f766520746f2063616c6c65720000000000000000604482015290519081900360640190fd5b3360008181526004602090815260408083206001600160a01b03871680855290835292819020805460ff1916861515908117909155815190815290519293927f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31929181900390910190a35050565b610a2c8484846106cb565b610a3884848484610c79565b610a7657604051600160e51b62461bcd0281526004018080602001828103825260308152602001806114ad6030913960400191505060405180910390fd5b50505050565b6060610a8782610b92565b610ac557604051600160e51b62461bcd02815260040180806020018281038252602e8152602001806113da602e913960400191505060405180910390fd5b6000828152600b602090815260409182902080548351601f600260001961010060018616150201909316929092049182018490048402810184019094528084529091830182828015610b585780601f10610b2d57610100808354040283529160200191610b58565b820191906000526020600020905b815481529060010190602001808311610b3b57829003601f168201915b50505050509050919050565b6001600160a01b03918216600090815260046020908152604080832093909416825291909152205460ff1690565b6000908152600160205260409020546001600160a01b0316151590565b6000610bba82610b92565b610bf857604051600160e51b62461bcd02815260040180806020018281038252602b8152602001806115bb602b913960400191505060405180910390fd5b6000610c0383610829565b9050806001600160a01b0316846001600160a01b03161480610c3e5750836001600160a01b0316610c3384610533565b6001600160a01b0316145b80610c4e5750610c4e8185610b64565b949350505050565b610c6183838361105c565b610c6b83826111a6565b61071e828261129b565b5490565b6000806060610c90866001600160a01b03166112d9565b610c9f57600192505050610c4e565b856001600160a01b031663150b7a0260e01b3389888860405160240180856001600160a01b03166001600160a01b03168152602001846001600160a01b03166001600160a01b0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610d2b578181015183820152602001610d13565b50505050905090810190601f168015610d585780820380516001836020036101000a031916815260200191505b5060408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909a16999099178952518151919890975087965094509250829150849050835b60208310610dc05780518252601f199092019160209182019101610da1565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114610e22576040519150601f19603f3d011682016040523d82523d6000602084013e610e27565b606091505b508051919350915015801590610e6757508051600160e11b630a85bd01029060208084019190811015610e5957600080fd5b50516001600160e01b031916145b15610e7757600192505050610c4e565b856001600160a01b0316636745782b60e01b3389888860405160240180856001600160a01b03166001600160a01b03168152602001846001600160a01b03166001600160a01b0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610f03578181015183820152602001610eeb565b50505050905090810190601f168015610f305780820380516001836020036101000a031916815260200191505b5060408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909a16999099178952518151919890975087965094509250829150849050835b60208310610f985780518252601f199092019160209182019101610f79565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114610ffa576040519150601f19603f3d011682016040523d82523d6000602084013e610fff565b606091505b50805191935091501580159061103f57508051600160e01b636745782b02906020808401919081101561103157600080fd5b50516001600160e01b031916145b1561104f57600192505050610c4e565b5060009695505050505050565b826001600160a01b031661106f82610829565b6001600160a01b0316146110b757604051600160e51b62461bcd0281526004018080602001828103825260288152602001806115066028913960400191505060405180910390fd5b6001600160a01b0382166110ff57604051600160e51b62461bcd0281526004018080602001828103825260238152602001806114086023913960400191505060405180910390fd5b611108816112df565b6001600160a01b03831660009081526003602052604090206111299061131c565b6001600160a01b038216600090815260036020526040902061114a90611333565b60008181526001602052604080822080546001600160a01b0319166001600160a01b0386811691821790925591518493918716917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b6001600160a01b0382166000908152600560205260408120546111d090600163ffffffff61133c16565b60008381526006602052604090205490915080821461126b576001600160a01b038416600090815260056020526040812080548490811061120d57fe5b906000526020600020015490508060056000876001600160a01b03166001600160a01b03168152602001908152602001600020838154811061124b57fe5b600091825260208083209091019290925591825260069052604090208190555b6001600160a01b038416600090815260056020526040902080549061129490600019830161139c565b5050505050565b6001600160a01b0390911660009081526005602081815260408084208054868652600684529185208290559282526001810183559183529091200155565b3b151590565b6000818152600260205260409020546001600160a01b03161561131957600081815260026020526040902080546001600160a01b03191690555b50565b805461132f90600163ffffffff61133c16565b9055565b80546001019055565b6000828211156113965760408051600160e51b62461bcd02815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b81548183558181111561071e5760008381526020902061071e91810190830161053091905b808211156113d557600081556001016113c1565b509056fe4b495031374d657461646174613a2055524920717565727920666f72206e6f6e6578697374656e7420746f6b656e4b495031373a207472616e7366657220746f20746865207a65726f20616464726573734b49503137456e756d657261626c653a206f776e657220696e646578206f7574206f6620626f756e64734b495031373a206f776e657220717565727920666f72206e6f6e6578697374656e7420746f6b656e4b495031373a207472616e736665722063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f7665644b495031373a207472616e7366657220746f206e6f6e204b49503137526563656976657220696d706c656d656e7465724b495031373a2062616c616e636520717565727920666f7220746865207a65726f20616464726573734b495031373a207472616e73666572206f6620746f6b656e2074686174206973206e6f74206f776e4b49503137456e756d657261626c653a20676c6f62616c20696e646578206f7574206f6620626f756e64734b495031373a20617070726f76656420717565727920666f72206e6f6e6578697374656e7420746f6b656e4b495031373a20617070726f76652063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f76656420666f7220616c6c4b495031373a206f70657261746f7220717565727920666f72206e6f6e6578697374656e7420746f6b656ea165627a7a72305820ed7b15db174cbbafb0ffe4d8598c28d541404a3c946362ff10e4dca4b285d3750029"

// DeployKIP17Full deploys a new Klaytn contract, binding an instance of KIP17Full to it.
func DeployKIP17Full(auth *bind.TransactOpts, backend bind.ContractBackend, name string, symbol string) (common.Address, *types.Transaction, *KIP17Full, error) {
	parsed, err := abi.JSON(strings.NewReader(KIP17FullABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(KIP17FullBin), backend, name, symbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KIP17Full{KIP17FullCaller: KIP17FullCaller{contract: contract}, KIP17FullTransactor: KIP17FullTransactor{contract: contract}, KIP17FullFilterer: KIP17FullFilterer{contract: contract}}, nil
}

// KIP17Full is an auto generated Go binding around a Klaytn contract.
type KIP17Full struct {
	KIP17FullCaller     // Read-only binding to the contract
	KIP17FullTransactor // Write-only binding to the contract
	KIP17FullFilterer   // Log filterer for contract events
}

// KIP17FullCaller is an auto generated read-only Go binding around a Klaytn contract.
type KIP17FullCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KIP17FullTransactor is an auto generated write-only Go binding around a Klaytn contract.
type KIP17FullTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KIP17FullFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type KIP17FullFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KIP17FullSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type KIP17FullSession struct {
	Contract     *KIP17Full        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KIP17FullCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type KIP17FullCallerSession struct {
	Contract *KIP17FullCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// KIP17FullTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type KIP17FullTransactorSession struct {
	Contract     *KIP17FullTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// KIP17FullRaw is an auto generated low-level Go binding around a Klaytn contract.
type KIP17FullRaw struct {
	Contract *KIP17Full // Generic contract binding to access the raw methods on
}

// KIP17FullCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type KIP17FullCallerRaw struct {
	Contract *KIP17FullCaller // Generic read-only contract binding to access the raw methods on
}

// KIP17FullTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type KIP17FullTransactorRaw struct {
	Contract *KIP17FullTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKIP17Full creates a new instance of KIP17Full, bound to a specific deployed contract.
func NewKIP17Full(address common.Address, backend bind.ContractBackend) (*KIP17Full, error) {
	contract, err := bindKIP17Full(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KIP17Full{KIP17FullCaller: KIP17FullCaller{contract: contract}, KIP17FullTransactor: KIP17FullTransactor{contract: contract}, KIP17FullFilterer: KIP17FullFilterer{contract: contract}}, nil
}

// NewKIP17FullCaller creates a new read-only instance of KIP17Full, bound to a specific deployed contract.
func NewKIP17FullCaller(address common.Address, caller bind.ContractCaller) (*KIP17FullCaller, error) {
	contract, err := bindKIP17Full(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KIP17FullCaller{contract: contract}, nil
}

// NewKIP17FullTransactor creates a new write-only instance of KIP17Full, bound to a specific deployed contract.
func NewKIP17FullTransactor(address common.Address, transactor bind.ContractTransactor) (*KIP17FullTransactor, error) {
	contract, err := bindKIP17Full(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KIP17FullTransactor{contract: contract}, nil
}

// NewKIP17FullFilterer creates a new log filterer instance of KIP17Full, bound to a specific deployed contract.
func NewKIP17FullFilterer(address common.Address, filterer bind.ContractFilterer) (*KIP17FullFilterer, error) {
	contract, err := bindKIP17Full(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KIP17FullFilterer{contract: contract}, nil
}

// bindKIP17Full binds a generic wrapper to an already deployed contract.
func bindKIP17Full(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KIP17FullABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KIP17Full *KIP17FullRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KIP17Full.Contract.KIP17FullCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KIP17Full *KIP17FullRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KIP17Full.Contract.KIP17FullTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KIP17Full *KIP17FullRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KIP17Full.Contract.KIP17FullTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KIP17Full *KIP17FullCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KIP17Full.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KIP17Full *KIP17FullTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KIP17Full.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KIP17Full *KIP17FullTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KIP17Full.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_KIP17Full *KIP17FullCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KIP17Full.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_KIP17Full *KIP17FullSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _KIP17Full.Contract.BalanceOf(&_KIP17Full.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_KIP17Full *KIP17FullCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _KIP17Full.Contract.BalanceOf(&_KIP17Full.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_KIP17Full *KIP17FullCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KIP17Full.contract.Call(opts, out, "getApproved", tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_KIP17Full *KIP17FullSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _KIP17Full.Contract.GetApproved(&_KIP17Full.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_KIP17Full *KIP17FullCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _KIP17Full.Contract.GetApproved(&_KIP17Full.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_KIP17Full *KIP17FullCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KIP17Full.contract.Call(opts, out, "isApprovedForAll", owner, operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_KIP17Full *KIP17FullSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _KIP17Full.Contract.IsApprovedForAll(&_KIP17Full.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_KIP17Full *KIP17FullCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _KIP17Full.Contract.IsApprovedForAll(&_KIP17Full.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_KIP17Full *KIP17FullCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _KIP17Full.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_KIP17Full *KIP17FullSession) Name() (string, error) {
	return _KIP17Full.Contract.Name(&_KIP17Full.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_KIP17Full *KIP17FullCallerSession) Name() (string, error) {
	return _KIP17Full.Contract.Name(&_KIP17Full.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_KIP17Full *KIP17FullCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KIP17Full.contract.Call(opts, out, "ownerOf", tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_KIP17Full *KIP17FullSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _KIP17Full.Contract.OwnerOf(&_KIP17Full.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_KIP17Full *KIP17FullCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _KIP17Full.Contract.OwnerOf(&_KIP17Full.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_KIP17Full *KIP17FullCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KIP17Full.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_KIP17Full *KIP17FullSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _KIP17Full.Contract.SupportsInterface(&_KIP17Full.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_KIP17Full *KIP17FullCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _KIP17Full.Contract.SupportsInterface(&_KIP17Full.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_KIP17Full *KIP17FullCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _KIP17Full.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_KIP17Full *KIP17FullSession) Symbol() (string, error) {
	return _KIP17Full.Contract.Symbol(&_KIP17Full.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_KIP17Full *KIP17FullCallerSession) Symbol() (string, error) {
	return _KIP17Full.Contract.Symbol(&_KIP17Full.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_KIP17Full *KIP17FullCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KIP17Full.contract.Call(opts, out, "tokenByIndex", index)
	return *ret0, err
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_KIP17Full *KIP17FullSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _KIP17Full.Contract.TokenByIndex(&_KIP17Full.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_KIP17Full *KIP17FullCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _KIP17Full.Contract.TokenByIndex(&_KIP17Full.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_KIP17Full *KIP17FullCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KIP17Full.contract.Call(opts, out, "tokenOfOwnerByIndex", owner, index)
	return *ret0, err
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_KIP17Full *KIP17FullSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _KIP17Full.Contract.TokenOfOwnerByIndex(&_KIP17Full.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_KIP17Full *KIP17FullCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _KIP17Full.Contract.TokenOfOwnerByIndex(&_KIP17Full.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_KIP17Full *KIP17FullCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _KIP17Full.contract.Call(opts, out, "tokenURI", tokenId)
	return *ret0, err
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_KIP17Full *KIP17FullSession) TokenURI(tokenId *big.Int) (string, error) {
	return _KIP17Full.Contract.TokenURI(&_KIP17Full.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_KIP17Full *KIP17FullCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _KIP17Full.Contract.TokenURI(&_KIP17Full.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_KIP17Full *KIP17FullCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KIP17Full.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_KIP17Full *KIP17FullSession) TotalSupply() (*big.Int, error) {
	return _KIP17Full.Contract.TotalSupply(&_KIP17Full.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_KIP17Full *KIP17FullCallerSession) TotalSupply() (*big.Int, error) {
	return _KIP17Full.Contract.TotalSupply(&_KIP17Full.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_KIP17Full *KIP17FullTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17Full.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_KIP17Full *KIP17FullSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17Full.Contract.Approve(&_KIP17Full.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_KIP17Full *KIP17FullTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17Full.Contract.Approve(&_KIP17Full.TransactOpts, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_KIP17Full *KIP17FullTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17Full.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_KIP17Full *KIP17FullSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17Full.Contract.SafeTransferFrom(&_KIP17Full.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_KIP17Full *KIP17FullTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17Full.Contract.SafeTransferFrom(&_KIP17Full.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_KIP17Full *KIP17FullTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _KIP17Full.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_KIP17Full *KIP17FullSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _KIP17Full.Contract.SafeTransferFrom0(&_KIP17Full.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_KIP17Full *KIP17FullTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _KIP17Full.Contract.SafeTransferFrom0(&_KIP17Full.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_KIP17Full *KIP17FullTransactor) SetApprovalForAll(opts *bind.TransactOpts, to common.Address, approved bool) (*types.Transaction, error) {
	return _KIP17Full.contract.Transact(opts, "setApprovalForAll", to, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_KIP17Full *KIP17FullSession) SetApprovalForAll(to common.Address, approved bool) (*types.Transaction, error) {
	return _KIP17Full.Contract.SetApprovalForAll(&_KIP17Full.TransactOpts, to, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_KIP17Full *KIP17FullTransactorSession) SetApprovalForAll(to common.Address, approved bool) (*types.Transaction, error) {
	return _KIP17Full.Contract.SetApprovalForAll(&_KIP17Full.TransactOpts, to, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_KIP17Full *KIP17FullTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17Full.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_KIP17Full *KIP17FullSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17Full.Contract.TransferFrom(&_KIP17Full.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_KIP17Full *KIP17FullTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17Full.Contract.TransferFrom(&_KIP17Full.TransactOpts, from, to, tokenId)
}

// KIP17FullApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the KIP17Full contract.
type KIP17FullApprovalIterator struct {
	Event *KIP17FullApproval // Event containing the contract specifics and raw log

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
func (it *KIP17FullApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KIP17FullApproval)
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
		it.Event = new(KIP17FullApproval)
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
func (it *KIP17FullApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KIP17FullApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KIP17FullApproval represents a Approval event raised by the KIP17Full contract.
type KIP17FullApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_KIP17Full *KIP17FullFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*KIP17FullApprovalIterator, error) {

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

	logs, sub, err := _KIP17Full.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &KIP17FullApprovalIterator{contract: _KIP17Full.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_KIP17Full *KIP17FullFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *KIP17FullApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _KIP17Full.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KIP17FullApproval)
				if err := _KIP17Full.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_KIP17Full *KIP17FullFilterer) ParseApproval(log types.Log) (*KIP17FullApproval, error) {
	event := new(KIP17FullApproval)
	if err := _KIP17Full.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KIP17FullApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the KIP17Full contract.
type KIP17FullApprovalForAllIterator struct {
	Event *KIP17FullApprovalForAll // Event containing the contract specifics and raw log

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
func (it *KIP17FullApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KIP17FullApprovalForAll)
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
		it.Event = new(KIP17FullApprovalForAll)
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
func (it *KIP17FullApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KIP17FullApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KIP17FullApprovalForAll represents a ApprovalForAll event raised by the KIP17Full contract.
type KIP17FullApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_KIP17Full *KIP17FullFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*KIP17FullApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _KIP17Full.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &KIP17FullApprovalForAllIterator{contract: _KIP17Full.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_KIP17Full *KIP17FullFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *KIP17FullApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _KIP17Full.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KIP17FullApprovalForAll)
				if err := _KIP17Full.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_KIP17Full *KIP17FullFilterer) ParseApprovalForAll(log types.Log) (*KIP17FullApprovalForAll, error) {
	event := new(KIP17FullApprovalForAll)
	if err := _KIP17Full.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KIP17FullTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the KIP17Full contract.
type KIP17FullTransferIterator struct {
	Event *KIP17FullTransfer // Event containing the contract specifics and raw log

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
func (it *KIP17FullTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KIP17FullTransfer)
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
		it.Event = new(KIP17FullTransfer)
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
func (it *KIP17FullTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KIP17FullTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KIP17FullTransfer represents a Transfer event raised by the KIP17Full contract.
type KIP17FullTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_KIP17Full *KIP17FullFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*KIP17FullTransferIterator, error) {

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

	logs, sub, err := _KIP17Full.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &KIP17FullTransferIterator{contract: _KIP17Full.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_KIP17Full *KIP17FullFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *KIP17FullTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _KIP17Full.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KIP17FullTransfer)
				if err := _KIP17Full.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_KIP17Full *KIP17FullFilterer) ParseTransfer(log types.Log) (*KIP17FullTransfer, error) {
	event := new(KIP17FullTransfer)
	if err := _KIP17Full.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KIP17MetadataABI is the input ABI used to generate the binding from.
const KIP17MetadataABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"}]"

// KIP17MetadataBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const KIP17MetadataBinRuntime = `608060405234801561001057600080fd5b50600436106100cf5760003560e01c80636352211e1161008c578063a22cb46511610066578063a22cb465146102bc578063b88d4fde146102ea578063c87b56dd146103b0578063e985e9c5146103cd576100cf565b80636352211e1461025f57806370a082311461027c57806395d89b41146102b4576100cf565b806301ffc9a7146100d457806306fdde031461010f578063081812fc1461018c578063095ea7b3146101c557806323b872dd146101f357806342842e0e14610229575b600080fd5b6100fb600480360360208110156100ea57600080fd5b50356001600160e01b0319166103fb565b604080519115158252519081900360200190f35b61011761041a565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610151578181015183820152602001610139565b50505050905090810190601f16801561017e5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101a9600480360360208110156101a257600080fd5b50356104b0565b604080516001600160a01b039092168252519081900360200190f35b6101f1600480360360408110156101db57600080fd5b506001600160a01b038135169060200135610515565b005b6101f16004803603606081101561020957600080fd5b506001600160a01b03813581169160208101359091169060400135610642565b6101f16004803603606081101561023f57600080fd5b506001600160a01b0381358116916020810135909116906040013561069a565b6101a96004803603602081101561027557600080fd5b50356106b5565b6102a26004803603602081101561029257600080fd5b50356001600160a01b0316610712565b60408051918252519081900360200190f35b61011761077d565b6101f1600480360360408110156102d257600080fd5b506001600160a01b03813516906020013515156107de565b6101f16004803603608081101561030057600080fd5b6001600160a01b0382358116926020810135909116916040820135919081019060808101606082013564010000000081111561033b57600080fd5b82018360208201111561034d57600080fd5b8035906020019184600183028401116401000000008311171561036f57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506108ad945050505050565b610117600480360360208110156103c657600080fd5b5035610908565b6100fb600480360360408110156103e357600080fd5b506001600160a01b03813581169160200135166109f0565b6001600160e01b03191660009081526020819052604090205460ff1690565b60058054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156104a65780601f1061047b576101008083540402835291602001916104a6565b820191906000526020600020905b81548152906001019060200180831161048957829003601f168201915b5050505050905090565b60006104bb82610a1e565b6104f957604051600160e51b62461bcd02815260040180806020018281038252602b815260200180611201602b913960400191505060405180910390fd5b506000908152600260205260409020546001600160a01b031690565b6000610520826106b5565b9050806001600160a01b0316836001600160a01b0316141561058c5760408051600160e51b62461bcd02815260206004820181905260248201527f4b495031373a20617070726f76616c20746f2063757272656e74206f776e6572604482015290519081900360640190fd5b336001600160a01b03821614806105a857506105a881336109f0565b6105e657604051600160e51b62461bcd02815260040180806020018281038252603781526020018061122c6037913960400191505060405180910390fd5b60008281526002602052604080822080546001600160a01b0319166001600160a01b0387811691821790925591518593918516917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591a4505050565b61064c3382610a3b565b61068a57604051600160e51b62461bcd0281526004018080602001828103825260308152602001806111506030913960400191505060405180910390fd5b610695838383610ae2565b505050565b610695838383604051806020016040528060008152506108ad565b6000818152600160205260408120546001600160a01b03168061070c57604051600160e51b62461bcd0281526004018080602001828103825260288152602001806111286028913960400191505060405180910390fd5b92915050565b60006001600160a01b03821661075c57604051600160e51b62461bcd0281526004018080602001828103825260298152602001806111b06029913960400191505060405180910390fd5b6001600160a01b038216600090815260036020526040902061070c90610c2c565b60068054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156104a65780601f1061047b576101008083540402835291602001916104a6565b6001600160a01b03821633141561083f5760408051600160e51b62461bcd02815260206004820152601860248201527f4b495031373a20617070726f766520746f2063616c6c65720000000000000000604482015290519081900360640190fd5b3360008181526004602090815260408083206001600160a01b03871680855290835292819020805460ff1916861515908117909155815190815290519293927f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31929181900390910190a35050565b6108b8848484610642565b6108c484848484610c30565b61090257604051600160e51b62461bcd0281526004018080602001828103825260308152602001806111806030913960400191505060405180910390fd5b50505050565b606061091382610a1e565b61095157604051600160e51b62461bcd02815260040180806020018281038252602e8152602001806110d7602e913960400191505060405180910390fd5b60008281526007602090815260409182902080548351601f6002600019610100600186161502019093169290920491820184900484028101840190945280845290918301828280156109e45780601f106109b9576101008083540402835291602001916109e4565b820191906000526020600020905b8154815290600101906020018083116109c757829003601f168201915b50505050509050919050565b6001600160a01b03918216600090815260046020908152604080832093909416825291909152205460ff1690565b6000908152600160205260409020546001600160a01b0316151590565b6000610a4682610a1e565b610a8457604051600160e51b62461bcd02815260040180806020018281038252602b815260200180611263602b913960400191505060405180910390fd5b6000610a8f836106b5565b9050806001600160a01b0316846001600160a01b03161480610aca5750836001600160a01b0316610abf846104b0565b6001600160a01b0316145b80610ada5750610ada81856109f0565b949350505050565b826001600160a01b0316610af5826106b5565b6001600160a01b031614610b3d57604051600160e51b62461bcd0281526004018080602001828103825260288152602001806111d96028913960400191505060405180910390fd5b6001600160a01b038216610b8557604051600160e51b62461bcd0281526004018080602001828103825260238152602001806111056023913960400191505060405180910390fd5b610b8e81611013565b6001600160a01b0383166000908152600360205260409020610baf90611050565b6001600160a01b0382166000908152600360205260409020610bd090611067565b60008181526001602052604080822080546001600160a01b0319166001600160a01b0386811691821790925591518493918716917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b5490565b6000806060610c47866001600160a01b0316611070565b610c5657600192505050610ada565b856001600160a01b031663150b7a0260e01b3389888860405160240180856001600160a01b03166001600160a01b03168152602001846001600160a01b03166001600160a01b0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610ce2578181015183820152602001610cca565b50505050905090810190601f168015610d0f5780820380516001836020036101000a031916815260200191505b5060408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909a16999099178952518151919890975087965094509250829150849050835b60208310610d775780518252601f199092019160209182019101610d58565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114610dd9576040519150601f19603f3d011682016040523d82523d6000602084013e610dde565b606091505b508051919350915015801590610e1e57508051600160e11b630a85bd01029060208084019190811015610e1057600080fd5b50516001600160e01b031916145b15610e2e57600192505050610ada565b856001600160a01b0316636745782b60e01b3389888860405160240180856001600160a01b03166001600160a01b03168152602001846001600160a01b03166001600160a01b0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610eba578181015183820152602001610ea2565b50505050905090810190601f168015610ee75780820380516001836020036101000a031916815260200191505b5060408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909a16999099178952518151919890975087965094509250829150849050835b60208310610f4f5780518252601f199092019160209182019101610f30565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114610fb1576040519150601f19603f3d011682016040523d82523d6000602084013e610fb6565b606091505b508051919350915015801590610ff657508051600160e01b636745782b029060208084019190811015610fe857600080fd5b50516001600160e01b031916145b1561100657600192505050610ada565b5060009695505050505050565b6000818152600260205260409020546001600160a01b03161561104d57600081815260026020526040902080546001600160a01b03191690555b50565b805461106390600163ffffffff61107616565b9055565b80546001019055565b3b151590565b6000828211156110d05760408051600160e51b62461bcd02815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b5090039056fe4b495031374d657461646174613a2055524920717565727920666f72206e6f6e6578697374656e7420746f6b656e4b495031373a207472616e7366657220746f20746865207a65726f20616464726573734b495031373a206f776e657220717565727920666f72206e6f6e6578697374656e7420746f6b656e4b495031373a207472616e736665722063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f7665644b495031373a207472616e7366657220746f206e6f6e204b49503137526563656976657220696d706c656d656e7465724b495031373a2062616c616e636520717565727920666f7220746865207a65726f20616464726573734b495031373a207472616e73666572206f6620746f6b656e2074686174206973206e6f74206f776e4b495031373a20617070726f76656420717565727920666f72206e6f6e6578697374656e7420746f6b656e4b495031373a20617070726f76652063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f76656420666f7220616c6c4b495031373a206f70657261746f7220717565727920666f72206e6f6e6578697374656e7420746f6b656ea165627a7a72305820826c02912612c85f0cbc51c045aab41ba224242a8c5fae83df2eaa0940fe960a0029`

// KIP17MetadataFuncSigs maps the 4-byte function signature to its string representation.
var KIP17MetadataFuncSigs = map[string]string{
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"081812fc": "getApproved(uint256)",
	"e985e9c5": "isApprovedForAll(address,address)",
	"06fdde03": "name()",
	"6352211e": "ownerOf(uint256)",
	"42842e0e": "safeTransferFrom(address,address,uint256)",
	"b88d4fde": "safeTransferFrom(address,address,uint256,bytes)",
	"a22cb465": "setApprovalForAll(address,bool)",
	"01ffc9a7": "supportsInterface(bytes4)",
	"95d89b41": "symbol()",
	"c87b56dd": "tokenURI(uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// KIP17MetadataBin is the compiled bytecode used for deploying new contracts.
var KIP17MetadataBin = "0x60806040523480156200001157600080fd5b506040516200158c3803806200158c833981018060405260408110156200003757600080fd5b8101908080516401000000008111156200005057600080fd5b820160208101848111156200006457600080fd5b81516401000000008111828201871017156200007f57600080fd5b505092919060200180516401000000008111156200009c57600080fd5b82016020810184811115620000b057600080fd5b8151640100000000811182820187101715620000cb57600080fd5b5050929190505050620000eb6301ffc9a760e01b6200014f60201b60201c565b620001036380ac58cd60e01b6200014f60201b60201c565b8151620001189060059060208501906200021e565b5080516200012e9060069060208401906200021e565b5062000147635b5e139f60e01b6200014f60201b60201c565b5050620002c3565b7fffffffff000000000000000000000000000000000000000000000000000000008082161415620001e157604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f4b495031333a20696e76616c696420696e746572666163652069640000000000604482015290519081900360640190fd5b7fffffffff00000000000000000000000000000000000000000000000000000000166000908152602081905260409020805460ff19166001179055565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200026157805160ff191683800117855562000291565b8280016001018555821562000291579182015b828111156200029157825182559160200191906001019062000274565b506200029f929150620002a3565b5090565b620002c091905b808211156200029f5760008155600101620002aa565b90565b6112b980620002d36000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c80636352211e1161008c578063a22cb46511610066578063a22cb465146102bc578063b88d4fde146102ea578063c87b56dd146103b0578063e985e9c5146103cd576100cf565b80636352211e1461025f57806370a082311461027c57806395d89b41146102b4576100cf565b806301ffc9a7146100d457806306fdde031461010f578063081812fc1461018c578063095ea7b3146101c557806323b872dd146101f357806342842e0e14610229575b600080fd5b6100fb600480360360208110156100ea57600080fd5b50356001600160e01b0319166103fb565b604080519115158252519081900360200190f35b61011761041a565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610151578181015183820152602001610139565b50505050905090810190601f16801561017e5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101a9600480360360208110156101a257600080fd5b50356104b0565b604080516001600160a01b039092168252519081900360200190f35b6101f1600480360360408110156101db57600080fd5b506001600160a01b038135169060200135610515565b005b6101f16004803603606081101561020957600080fd5b506001600160a01b03813581169160208101359091169060400135610642565b6101f16004803603606081101561023f57600080fd5b506001600160a01b0381358116916020810135909116906040013561069a565b6101a96004803603602081101561027557600080fd5b50356106b5565b6102a26004803603602081101561029257600080fd5b50356001600160a01b0316610712565b60408051918252519081900360200190f35b61011761077d565b6101f1600480360360408110156102d257600080fd5b506001600160a01b03813516906020013515156107de565b6101f16004803603608081101561030057600080fd5b6001600160a01b0382358116926020810135909116916040820135919081019060808101606082013564010000000081111561033b57600080fd5b82018360208201111561034d57600080fd5b8035906020019184600183028401116401000000008311171561036f57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506108ad945050505050565b610117600480360360208110156103c657600080fd5b5035610908565b6100fb600480360360408110156103e357600080fd5b506001600160a01b03813581169160200135166109f0565b6001600160e01b03191660009081526020819052604090205460ff1690565b60058054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156104a65780601f1061047b576101008083540402835291602001916104a6565b820191906000526020600020905b81548152906001019060200180831161048957829003601f168201915b5050505050905090565b60006104bb82610a1e565b6104f957604051600160e51b62461bcd02815260040180806020018281038252602b815260200180611201602b913960400191505060405180910390fd5b506000908152600260205260409020546001600160a01b031690565b6000610520826106b5565b9050806001600160a01b0316836001600160a01b0316141561058c5760408051600160e51b62461bcd02815260206004820181905260248201527f4b495031373a20617070726f76616c20746f2063757272656e74206f776e6572604482015290519081900360640190fd5b336001600160a01b03821614806105a857506105a881336109f0565b6105e657604051600160e51b62461bcd02815260040180806020018281038252603781526020018061122c6037913960400191505060405180910390fd5b60008281526002602052604080822080546001600160a01b0319166001600160a01b0387811691821790925591518593918516917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591a4505050565b61064c3382610a3b565b61068a57604051600160e51b62461bcd0281526004018080602001828103825260308152602001806111506030913960400191505060405180910390fd5b610695838383610ae2565b505050565b610695838383604051806020016040528060008152506108ad565b6000818152600160205260408120546001600160a01b03168061070c57604051600160e51b62461bcd0281526004018080602001828103825260288152602001806111286028913960400191505060405180910390fd5b92915050565b60006001600160a01b03821661075c57604051600160e51b62461bcd0281526004018080602001828103825260298152602001806111b06029913960400191505060405180910390fd5b6001600160a01b038216600090815260036020526040902061070c90610c2c565b60068054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156104a65780601f1061047b576101008083540402835291602001916104a6565b6001600160a01b03821633141561083f5760408051600160e51b62461bcd02815260206004820152601860248201527f4b495031373a20617070726f766520746f2063616c6c65720000000000000000604482015290519081900360640190fd5b3360008181526004602090815260408083206001600160a01b03871680855290835292819020805460ff1916861515908117909155815190815290519293927f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31929181900390910190a35050565b6108b8848484610642565b6108c484848484610c30565b61090257604051600160e51b62461bcd0281526004018080602001828103825260308152602001806111806030913960400191505060405180910390fd5b50505050565b606061091382610a1e565b61095157604051600160e51b62461bcd02815260040180806020018281038252602e8152602001806110d7602e913960400191505060405180910390fd5b60008281526007602090815260409182902080548351601f6002600019610100600186161502019093169290920491820184900484028101840190945280845290918301828280156109e45780601f106109b9576101008083540402835291602001916109e4565b820191906000526020600020905b8154815290600101906020018083116109c757829003601f168201915b50505050509050919050565b6001600160a01b03918216600090815260046020908152604080832093909416825291909152205460ff1690565b6000908152600160205260409020546001600160a01b0316151590565b6000610a4682610a1e565b610a8457604051600160e51b62461bcd02815260040180806020018281038252602b815260200180611263602b913960400191505060405180910390fd5b6000610a8f836106b5565b9050806001600160a01b0316846001600160a01b03161480610aca5750836001600160a01b0316610abf846104b0565b6001600160a01b0316145b80610ada5750610ada81856109f0565b949350505050565b826001600160a01b0316610af5826106b5565b6001600160a01b031614610b3d57604051600160e51b62461bcd0281526004018080602001828103825260288152602001806111d96028913960400191505060405180910390fd5b6001600160a01b038216610b8557604051600160e51b62461bcd0281526004018080602001828103825260238152602001806111056023913960400191505060405180910390fd5b610b8e81611013565b6001600160a01b0383166000908152600360205260409020610baf90611050565b6001600160a01b0382166000908152600360205260409020610bd090611067565b60008181526001602052604080822080546001600160a01b0319166001600160a01b0386811691821790925591518493918716917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b5490565b6000806060610c47866001600160a01b0316611070565b610c5657600192505050610ada565b856001600160a01b031663150b7a0260e01b3389888860405160240180856001600160a01b03166001600160a01b03168152602001846001600160a01b03166001600160a01b0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610ce2578181015183820152602001610cca565b50505050905090810190601f168015610d0f5780820380516001836020036101000a031916815260200191505b5060408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909a16999099178952518151919890975087965094509250829150849050835b60208310610d775780518252601f199092019160209182019101610d58565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114610dd9576040519150601f19603f3d011682016040523d82523d6000602084013e610dde565b606091505b508051919350915015801590610e1e57508051600160e11b630a85bd01029060208084019190811015610e1057600080fd5b50516001600160e01b031916145b15610e2e57600192505050610ada565b856001600160a01b0316636745782b60e01b3389888860405160240180856001600160a01b03166001600160a01b03168152602001846001600160a01b03166001600160a01b0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610eba578181015183820152602001610ea2565b50505050905090810190601f168015610ee75780820380516001836020036101000a031916815260200191505b5060408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909a16999099178952518151919890975087965094509250829150849050835b60208310610f4f5780518252601f199092019160209182019101610f30565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114610fb1576040519150601f19603f3d011682016040523d82523d6000602084013e610fb6565b606091505b508051919350915015801590610ff657508051600160e01b636745782b029060208084019190811015610fe857600080fd5b50516001600160e01b031916145b1561100657600192505050610ada565b5060009695505050505050565b6000818152600260205260409020546001600160a01b03161561104d57600081815260026020526040902080546001600160a01b03191690555b50565b805461106390600163ffffffff61107616565b9055565b80546001019055565b3b151590565b6000828211156110d05760408051600160e51b62461bcd02815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b5090039056fe4b495031374d657461646174613a2055524920717565727920666f72206e6f6e6578697374656e7420746f6b656e4b495031373a207472616e7366657220746f20746865207a65726f20616464726573734b495031373a206f776e657220717565727920666f72206e6f6e6578697374656e7420746f6b656e4b495031373a207472616e736665722063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f7665644b495031373a207472616e7366657220746f206e6f6e204b49503137526563656976657220696d706c656d656e7465724b495031373a2062616c616e636520717565727920666f7220746865207a65726f20616464726573734b495031373a207472616e73666572206f6620746f6b656e2074686174206973206e6f74206f776e4b495031373a20617070726f76656420717565727920666f72206e6f6e6578697374656e7420746f6b656e4b495031373a20617070726f76652063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f76656420666f7220616c6c4b495031373a206f70657261746f7220717565727920666f72206e6f6e6578697374656e7420746f6b656ea165627a7a72305820826c02912612c85f0cbc51c045aab41ba224242a8c5fae83df2eaa0940fe960a0029"

// DeployKIP17Metadata deploys a new Klaytn contract, binding an instance of KIP17Metadata to it.
func DeployKIP17Metadata(auth *bind.TransactOpts, backend bind.ContractBackend, name string, symbol string) (common.Address, *types.Transaction, *KIP17Metadata, error) {
	parsed, err := abi.JSON(strings.NewReader(KIP17MetadataABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(KIP17MetadataBin), backend, name, symbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KIP17Metadata{KIP17MetadataCaller: KIP17MetadataCaller{contract: contract}, KIP17MetadataTransactor: KIP17MetadataTransactor{contract: contract}, KIP17MetadataFilterer: KIP17MetadataFilterer{contract: contract}}, nil
}

// KIP17Metadata is an auto generated Go binding around a Klaytn contract.
type KIP17Metadata struct {
	KIP17MetadataCaller     // Read-only binding to the contract
	KIP17MetadataTransactor // Write-only binding to the contract
	KIP17MetadataFilterer   // Log filterer for contract events
}

// KIP17MetadataCaller is an auto generated read-only Go binding around a Klaytn contract.
type KIP17MetadataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KIP17MetadataTransactor is an auto generated write-only Go binding around a Klaytn contract.
type KIP17MetadataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KIP17MetadataFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type KIP17MetadataFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KIP17MetadataSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type KIP17MetadataSession struct {
	Contract     *KIP17Metadata    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KIP17MetadataCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type KIP17MetadataCallerSession struct {
	Contract *KIP17MetadataCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// KIP17MetadataTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type KIP17MetadataTransactorSession struct {
	Contract     *KIP17MetadataTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// KIP17MetadataRaw is an auto generated low-level Go binding around a Klaytn contract.
type KIP17MetadataRaw struct {
	Contract *KIP17Metadata // Generic contract binding to access the raw methods on
}

// KIP17MetadataCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type KIP17MetadataCallerRaw struct {
	Contract *KIP17MetadataCaller // Generic read-only contract binding to access the raw methods on
}

// KIP17MetadataTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type KIP17MetadataTransactorRaw struct {
	Contract *KIP17MetadataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKIP17Metadata creates a new instance of KIP17Metadata, bound to a specific deployed contract.
func NewKIP17Metadata(address common.Address, backend bind.ContractBackend) (*KIP17Metadata, error) {
	contract, err := bindKIP17Metadata(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KIP17Metadata{KIP17MetadataCaller: KIP17MetadataCaller{contract: contract}, KIP17MetadataTransactor: KIP17MetadataTransactor{contract: contract}, KIP17MetadataFilterer: KIP17MetadataFilterer{contract: contract}}, nil
}

// NewKIP17MetadataCaller creates a new read-only instance of KIP17Metadata, bound to a specific deployed contract.
func NewKIP17MetadataCaller(address common.Address, caller bind.ContractCaller) (*KIP17MetadataCaller, error) {
	contract, err := bindKIP17Metadata(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KIP17MetadataCaller{contract: contract}, nil
}

// NewKIP17MetadataTransactor creates a new write-only instance of KIP17Metadata, bound to a specific deployed contract.
func NewKIP17MetadataTransactor(address common.Address, transactor bind.ContractTransactor) (*KIP17MetadataTransactor, error) {
	contract, err := bindKIP17Metadata(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KIP17MetadataTransactor{contract: contract}, nil
}

// NewKIP17MetadataFilterer creates a new log filterer instance of KIP17Metadata, bound to a specific deployed contract.
func NewKIP17MetadataFilterer(address common.Address, filterer bind.ContractFilterer) (*KIP17MetadataFilterer, error) {
	contract, err := bindKIP17Metadata(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KIP17MetadataFilterer{contract: contract}, nil
}

// bindKIP17Metadata binds a generic wrapper to an already deployed contract.
func bindKIP17Metadata(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KIP17MetadataABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KIP17Metadata *KIP17MetadataRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KIP17Metadata.Contract.KIP17MetadataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KIP17Metadata *KIP17MetadataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KIP17Metadata.Contract.KIP17MetadataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KIP17Metadata *KIP17MetadataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KIP17Metadata.Contract.KIP17MetadataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KIP17Metadata *KIP17MetadataCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KIP17Metadata.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KIP17Metadata *KIP17MetadataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KIP17Metadata.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KIP17Metadata *KIP17MetadataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KIP17Metadata.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_KIP17Metadata *KIP17MetadataCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KIP17Metadata.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_KIP17Metadata *KIP17MetadataSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _KIP17Metadata.Contract.BalanceOf(&_KIP17Metadata.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_KIP17Metadata *KIP17MetadataCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _KIP17Metadata.Contract.BalanceOf(&_KIP17Metadata.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_KIP17Metadata *KIP17MetadataCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KIP17Metadata.contract.Call(opts, out, "getApproved", tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_KIP17Metadata *KIP17MetadataSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _KIP17Metadata.Contract.GetApproved(&_KIP17Metadata.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_KIP17Metadata *KIP17MetadataCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _KIP17Metadata.Contract.GetApproved(&_KIP17Metadata.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_KIP17Metadata *KIP17MetadataCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KIP17Metadata.contract.Call(opts, out, "isApprovedForAll", owner, operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_KIP17Metadata *KIP17MetadataSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _KIP17Metadata.Contract.IsApprovedForAll(&_KIP17Metadata.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_KIP17Metadata *KIP17MetadataCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _KIP17Metadata.Contract.IsApprovedForAll(&_KIP17Metadata.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_KIP17Metadata *KIP17MetadataCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _KIP17Metadata.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_KIP17Metadata *KIP17MetadataSession) Name() (string, error) {
	return _KIP17Metadata.Contract.Name(&_KIP17Metadata.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_KIP17Metadata *KIP17MetadataCallerSession) Name() (string, error) {
	return _KIP17Metadata.Contract.Name(&_KIP17Metadata.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_KIP17Metadata *KIP17MetadataCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KIP17Metadata.contract.Call(opts, out, "ownerOf", tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_KIP17Metadata *KIP17MetadataSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _KIP17Metadata.Contract.OwnerOf(&_KIP17Metadata.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_KIP17Metadata *KIP17MetadataCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _KIP17Metadata.Contract.OwnerOf(&_KIP17Metadata.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_KIP17Metadata *KIP17MetadataCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KIP17Metadata.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_KIP17Metadata *KIP17MetadataSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _KIP17Metadata.Contract.SupportsInterface(&_KIP17Metadata.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_KIP17Metadata *KIP17MetadataCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _KIP17Metadata.Contract.SupportsInterface(&_KIP17Metadata.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_KIP17Metadata *KIP17MetadataCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _KIP17Metadata.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_KIP17Metadata *KIP17MetadataSession) Symbol() (string, error) {
	return _KIP17Metadata.Contract.Symbol(&_KIP17Metadata.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_KIP17Metadata *KIP17MetadataCallerSession) Symbol() (string, error) {
	return _KIP17Metadata.Contract.Symbol(&_KIP17Metadata.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_KIP17Metadata *KIP17MetadataCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _KIP17Metadata.contract.Call(opts, out, "tokenURI", tokenId)
	return *ret0, err
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_KIP17Metadata *KIP17MetadataSession) TokenURI(tokenId *big.Int) (string, error) {
	return _KIP17Metadata.Contract.TokenURI(&_KIP17Metadata.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_KIP17Metadata *KIP17MetadataCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _KIP17Metadata.Contract.TokenURI(&_KIP17Metadata.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_KIP17Metadata *KIP17MetadataTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17Metadata.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_KIP17Metadata *KIP17MetadataSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17Metadata.Contract.Approve(&_KIP17Metadata.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_KIP17Metadata *KIP17MetadataTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17Metadata.Contract.Approve(&_KIP17Metadata.TransactOpts, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_KIP17Metadata *KIP17MetadataTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17Metadata.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_KIP17Metadata *KIP17MetadataSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17Metadata.Contract.SafeTransferFrom(&_KIP17Metadata.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_KIP17Metadata *KIP17MetadataTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17Metadata.Contract.SafeTransferFrom(&_KIP17Metadata.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_KIP17Metadata *KIP17MetadataTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _KIP17Metadata.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_KIP17Metadata *KIP17MetadataSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _KIP17Metadata.Contract.SafeTransferFrom0(&_KIP17Metadata.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_KIP17Metadata *KIP17MetadataTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _KIP17Metadata.Contract.SafeTransferFrom0(&_KIP17Metadata.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_KIP17Metadata *KIP17MetadataTransactor) SetApprovalForAll(opts *bind.TransactOpts, to common.Address, approved bool) (*types.Transaction, error) {
	return _KIP17Metadata.contract.Transact(opts, "setApprovalForAll", to, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_KIP17Metadata *KIP17MetadataSession) SetApprovalForAll(to common.Address, approved bool) (*types.Transaction, error) {
	return _KIP17Metadata.Contract.SetApprovalForAll(&_KIP17Metadata.TransactOpts, to, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_KIP17Metadata *KIP17MetadataTransactorSession) SetApprovalForAll(to common.Address, approved bool) (*types.Transaction, error) {
	return _KIP17Metadata.Contract.SetApprovalForAll(&_KIP17Metadata.TransactOpts, to, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_KIP17Metadata *KIP17MetadataTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17Metadata.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_KIP17Metadata *KIP17MetadataSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17Metadata.Contract.TransferFrom(&_KIP17Metadata.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_KIP17Metadata *KIP17MetadataTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17Metadata.Contract.TransferFrom(&_KIP17Metadata.TransactOpts, from, to, tokenId)
}

// KIP17MetadataApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the KIP17Metadata contract.
type KIP17MetadataApprovalIterator struct {
	Event *KIP17MetadataApproval // Event containing the contract specifics and raw log

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
func (it *KIP17MetadataApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KIP17MetadataApproval)
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
		it.Event = new(KIP17MetadataApproval)
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
func (it *KIP17MetadataApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KIP17MetadataApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KIP17MetadataApproval represents a Approval event raised by the KIP17Metadata contract.
type KIP17MetadataApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_KIP17Metadata *KIP17MetadataFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*KIP17MetadataApprovalIterator, error) {

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

	logs, sub, err := _KIP17Metadata.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &KIP17MetadataApprovalIterator{contract: _KIP17Metadata.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_KIP17Metadata *KIP17MetadataFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *KIP17MetadataApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _KIP17Metadata.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KIP17MetadataApproval)
				if err := _KIP17Metadata.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_KIP17Metadata *KIP17MetadataFilterer) ParseApproval(log types.Log) (*KIP17MetadataApproval, error) {
	event := new(KIP17MetadataApproval)
	if err := _KIP17Metadata.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KIP17MetadataApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the KIP17Metadata contract.
type KIP17MetadataApprovalForAllIterator struct {
	Event *KIP17MetadataApprovalForAll // Event containing the contract specifics and raw log

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
func (it *KIP17MetadataApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KIP17MetadataApprovalForAll)
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
		it.Event = new(KIP17MetadataApprovalForAll)
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
func (it *KIP17MetadataApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KIP17MetadataApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KIP17MetadataApprovalForAll represents a ApprovalForAll event raised by the KIP17Metadata contract.
type KIP17MetadataApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_KIP17Metadata *KIP17MetadataFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*KIP17MetadataApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _KIP17Metadata.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &KIP17MetadataApprovalForAllIterator{contract: _KIP17Metadata.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_KIP17Metadata *KIP17MetadataFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *KIP17MetadataApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _KIP17Metadata.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KIP17MetadataApprovalForAll)
				if err := _KIP17Metadata.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_KIP17Metadata *KIP17MetadataFilterer) ParseApprovalForAll(log types.Log) (*KIP17MetadataApprovalForAll, error) {
	event := new(KIP17MetadataApprovalForAll)
	if err := _KIP17Metadata.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KIP17MetadataTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the KIP17Metadata contract.
type KIP17MetadataTransferIterator struct {
	Event *KIP17MetadataTransfer // Event containing the contract specifics and raw log

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
func (it *KIP17MetadataTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KIP17MetadataTransfer)
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
		it.Event = new(KIP17MetadataTransfer)
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
func (it *KIP17MetadataTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KIP17MetadataTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KIP17MetadataTransfer represents a Transfer event raised by the KIP17Metadata contract.
type KIP17MetadataTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_KIP17Metadata *KIP17MetadataFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*KIP17MetadataTransferIterator, error) {

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

	logs, sub, err := _KIP17Metadata.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &KIP17MetadataTransferIterator{contract: _KIP17Metadata.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_KIP17Metadata *KIP17MetadataFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *KIP17MetadataTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _KIP17Metadata.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KIP17MetadataTransfer)
				if err := _KIP17Metadata.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_KIP17Metadata *KIP17MetadataFilterer) ParseTransfer(log types.Log) (*KIP17MetadataTransfer, error) {
	event := new(KIP17MetadataTransfer)
	if err := _KIP17Metadata.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KIP17MetadataMintableABI is the input ABI used to generate the binding from.
const KIP17MetadataMintableABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"},{\"name\":\"tokenURI\",\"type\":\"string\"}],\"name\":\"mintWithTokenURI\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addMinter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceMinter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isMinter\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MinterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MinterRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"}]"

// KIP17MetadataMintableBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const KIP17MetadataMintableBinRuntime = ``

// KIP17MetadataMintableFuncSigs maps the 4-byte function signature to its string representation.
var KIP17MetadataMintableFuncSigs = map[string]string{
	"983b2d56": "addMinter(address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"081812fc": "getApproved(uint256)",
	"e985e9c5": "isApprovedForAll(address,address)",
	"aa271e1a": "isMinter(address)",
	"50bb4e7f": "mintWithTokenURI(address,uint256,string)",
	"06fdde03": "name()",
	"6352211e": "ownerOf(uint256)",
	"98650275": "renounceMinter()",
	"42842e0e": "safeTransferFrom(address,address,uint256)",
	"b88d4fde": "safeTransferFrom(address,address,uint256,bytes)",
	"a22cb465": "setApprovalForAll(address,bool)",
	"01ffc9a7": "supportsInterface(bytes4)",
	"95d89b41": "symbol()",
	"c87b56dd": "tokenURI(uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// KIP17MetadataMintable is an auto generated Go binding around a Klaytn contract.
type KIP17MetadataMintable struct {
	KIP17MetadataMintableCaller     // Read-only binding to the contract
	KIP17MetadataMintableTransactor // Write-only binding to the contract
	KIP17MetadataMintableFilterer   // Log filterer for contract events
}

// KIP17MetadataMintableCaller is an auto generated read-only Go binding around a Klaytn contract.
type KIP17MetadataMintableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KIP17MetadataMintableTransactor is an auto generated write-only Go binding around a Klaytn contract.
type KIP17MetadataMintableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KIP17MetadataMintableFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type KIP17MetadataMintableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KIP17MetadataMintableSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type KIP17MetadataMintableSession struct {
	Contract     *KIP17MetadataMintable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// KIP17MetadataMintableCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type KIP17MetadataMintableCallerSession struct {
	Contract *KIP17MetadataMintableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// KIP17MetadataMintableTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type KIP17MetadataMintableTransactorSession struct {
	Contract     *KIP17MetadataMintableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// KIP17MetadataMintableRaw is an auto generated low-level Go binding around a Klaytn contract.
type KIP17MetadataMintableRaw struct {
	Contract *KIP17MetadataMintable // Generic contract binding to access the raw methods on
}

// KIP17MetadataMintableCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type KIP17MetadataMintableCallerRaw struct {
	Contract *KIP17MetadataMintableCaller // Generic read-only contract binding to access the raw methods on
}

// KIP17MetadataMintableTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type KIP17MetadataMintableTransactorRaw struct {
	Contract *KIP17MetadataMintableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKIP17MetadataMintable creates a new instance of KIP17MetadataMintable, bound to a specific deployed contract.
func NewKIP17MetadataMintable(address common.Address, backend bind.ContractBackend) (*KIP17MetadataMintable, error) {
	contract, err := bindKIP17MetadataMintable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KIP17MetadataMintable{KIP17MetadataMintableCaller: KIP17MetadataMintableCaller{contract: contract}, KIP17MetadataMintableTransactor: KIP17MetadataMintableTransactor{contract: contract}, KIP17MetadataMintableFilterer: KIP17MetadataMintableFilterer{contract: contract}}, nil
}

// NewKIP17MetadataMintableCaller creates a new read-only instance of KIP17MetadataMintable, bound to a specific deployed contract.
func NewKIP17MetadataMintableCaller(address common.Address, caller bind.ContractCaller) (*KIP17MetadataMintableCaller, error) {
	contract, err := bindKIP17MetadataMintable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KIP17MetadataMintableCaller{contract: contract}, nil
}

// NewKIP17MetadataMintableTransactor creates a new write-only instance of KIP17MetadataMintable, bound to a specific deployed contract.
func NewKIP17MetadataMintableTransactor(address common.Address, transactor bind.ContractTransactor) (*KIP17MetadataMintableTransactor, error) {
	contract, err := bindKIP17MetadataMintable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KIP17MetadataMintableTransactor{contract: contract}, nil
}

// NewKIP17MetadataMintableFilterer creates a new log filterer instance of KIP17MetadataMintable, bound to a specific deployed contract.
func NewKIP17MetadataMintableFilterer(address common.Address, filterer bind.ContractFilterer) (*KIP17MetadataMintableFilterer, error) {
	contract, err := bindKIP17MetadataMintable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KIP17MetadataMintableFilterer{contract: contract}, nil
}

// bindKIP17MetadataMintable binds a generic wrapper to an already deployed contract.
func bindKIP17MetadataMintable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KIP17MetadataMintableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KIP17MetadataMintable *KIP17MetadataMintableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KIP17MetadataMintable.Contract.KIP17MetadataMintableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KIP17MetadataMintable *KIP17MetadataMintableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KIP17MetadataMintable.Contract.KIP17MetadataMintableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KIP17MetadataMintable *KIP17MetadataMintableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KIP17MetadataMintable.Contract.KIP17MetadataMintableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KIP17MetadataMintable *KIP17MetadataMintableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KIP17MetadataMintable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KIP17MetadataMintable *KIP17MetadataMintableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KIP17MetadataMintable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KIP17MetadataMintable *KIP17MetadataMintableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KIP17MetadataMintable.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_KIP17MetadataMintable *KIP17MetadataMintableCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KIP17MetadataMintable.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_KIP17MetadataMintable *KIP17MetadataMintableSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _KIP17MetadataMintable.Contract.BalanceOf(&_KIP17MetadataMintable.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_KIP17MetadataMintable *KIP17MetadataMintableCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _KIP17MetadataMintable.Contract.BalanceOf(&_KIP17MetadataMintable.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_KIP17MetadataMintable *KIP17MetadataMintableCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KIP17MetadataMintable.contract.Call(opts, out, "getApproved", tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_KIP17MetadataMintable *KIP17MetadataMintableSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _KIP17MetadataMintable.Contract.GetApproved(&_KIP17MetadataMintable.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_KIP17MetadataMintable *KIP17MetadataMintableCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _KIP17MetadataMintable.Contract.GetApproved(&_KIP17MetadataMintable.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_KIP17MetadataMintable *KIP17MetadataMintableCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KIP17MetadataMintable.contract.Call(opts, out, "isApprovedForAll", owner, operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_KIP17MetadataMintable *KIP17MetadataMintableSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _KIP17MetadataMintable.Contract.IsApprovedForAll(&_KIP17MetadataMintable.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_KIP17MetadataMintable *KIP17MetadataMintableCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _KIP17MetadataMintable.Contract.IsApprovedForAll(&_KIP17MetadataMintable.CallOpts, owner, operator)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) view returns(bool)
func (_KIP17MetadataMintable *KIP17MetadataMintableCaller) IsMinter(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KIP17MetadataMintable.contract.Call(opts, out, "isMinter", account)
	return *ret0, err
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) view returns(bool)
func (_KIP17MetadataMintable *KIP17MetadataMintableSession) IsMinter(account common.Address) (bool, error) {
	return _KIP17MetadataMintable.Contract.IsMinter(&_KIP17MetadataMintable.CallOpts, account)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) view returns(bool)
func (_KIP17MetadataMintable *KIP17MetadataMintableCallerSession) IsMinter(account common.Address) (bool, error) {
	return _KIP17MetadataMintable.Contract.IsMinter(&_KIP17MetadataMintable.CallOpts, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_KIP17MetadataMintable *KIP17MetadataMintableCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _KIP17MetadataMintable.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_KIP17MetadataMintable *KIP17MetadataMintableSession) Name() (string, error) {
	return _KIP17MetadataMintable.Contract.Name(&_KIP17MetadataMintable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_KIP17MetadataMintable *KIP17MetadataMintableCallerSession) Name() (string, error) {
	return _KIP17MetadataMintable.Contract.Name(&_KIP17MetadataMintable.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_KIP17MetadataMintable *KIP17MetadataMintableCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KIP17MetadataMintable.contract.Call(opts, out, "ownerOf", tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_KIP17MetadataMintable *KIP17MetadataMintableSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _KIP17MetadataMintable.Contract.OwnerOf(&_KIP17MetadataMintable.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_KIP17MetadataMintable *KIP17MetadataMintableCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _KIP17MetadataMintable.Contract.OwnerOf(&_KIP17MetadataMintable.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_KIP17MetadataMintable *KIP17MetadataMintableCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KIP17MetadataMintable.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_KIP17MetadataMintable *KIP17MetadataMintableSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _KIP17MetadataMintable.Contract.SupportsInterface(&_KIP17MetadataMintable.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_KIP17MetadataMintable *KIP17MetadataMintableCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _KIP17MetadataMintable.Contract.SupportsInterface(&_KIP17MetadataMintable.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_KIP17MetadataMintable *KIP17MetadataMintableCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _KIP17MetadataMintable.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_KIP17MetadataMintable *KIP17MetadataMintableSession) Symbol() (string, error) {
	return _KIP17MetadataMintable.Contract.Symbol(&_KIP17MetadataMintable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_KIP17MetadataMintable *KIP17MetadataMintableCallerSession) Symbol() (string, error) {
	return _KIP17MetadataMintable.Contract.Symbol(&_KIP17MetadataMintable.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_KIP17MetadataMintable *KIP17MetadataMintableCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _KIP17MetadataMintable.contract.Call(opts, out, "tokenURI", tokenId)
	return *ret0, err
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_KIP17MetadataMintable *KIP17MetadataMintableSession) TokenURI(tokenId *big.Int) (string, error) {
	return _KIP17MetadataMintable.Contract.TokenURI(&_KIP17MetadataMintable.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_KIP17MetadataMintable *KIP17MetadataMintableCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _KIP17MetadataMintable.Contract.TokenURI(&_KIP17MetadataMintable.CallOpts, tokenId)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_KIP17MetadataMintable *KIP17MetadataMintableTransactor) AddMinter(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _KIP17MetadataMintable.contract.Transact(opts, "addMinter", account)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_KIP17MetadataMintable *KIP17MetadataMintableSession) AddMinter(account common.Address) (*types.Transaction, error) {
	return _KIP17MetadataMintable.Contract.AddMinter(&_KIP17MetadataMintable.TransactOpts, account)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_KIP17MetadataMintable *KIP17MetadataMintableTransactorSession) AddMinter(account common.Address) (*types.Transaction, error) {
	return _KIP17MetadataMintable.Contract.AddMinter(&_KIP17MetadataMintable.TransactOpts, account)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_KIP17MetadataMintable *KIP17MetadataMintableTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17MetadataMintable.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_KIP17MetadataMintable *KIP17MetadataMintableSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17MetadataMintable.Contract.Approve(&_KIP17MetadataMintable.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_KIP17MetadataMintable *KIP17MetadataMintableTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17MetadataMintable.Contract.Approve(&_KIP17MetadataMintable.TransactOpts, to, tokenId)
}

// MintWithTokenURI is a paid mutator transaction binding the contract method 0x50bb4e7f.
//
// Solidity: function mintWithTokenURI(address to, uint256 tokenId, string tokenURI) returns(bool)
func (_KIP17MetadataMintable *KIP17MetadataMintableTransactor) MintWithTokenURI(opts *bind.TransactOpts, to common.Address, tokenId *big.Int, tokenURI string) (*types.Transaction, error) {
	return _KIP17MetadataMintable.contract.Transact(opts, "mintWithTokenURI", to, tokenId, tokenURI)
}

// MintWithTokenURI is a paid mutator transaction binding the contract method 0x50bb4e7f.
//
// Solidity: function mintWithTokenURI(address to, uint256 tokenId, string tokenURI) returns(bool)
func (_KIP17MetadataMintable *KIP17MetadataMintableSession) MintWithTokenURI(to common.Address, tokenId *big.Int, tokenURI string) (*types.Transaction, error) {
	return _KIP17MetadataMintable.Contract.MintWithTokenURI(&_KIP17MetadataMintable.TransactOpts, to, tokenId, tokenURI)
}

// MintWithTokenURI is a paid mutator transaction binding the contract method 0x50bb4e7f.
//
// Solidity: function mintWithTokenURI(address to, uint256 tokenId, string tokenURI) returns(bool)
func (_KIP17MetadataMintable *KIP17MetadataMintableTransactorSession) MintWithTokenURI(to common.Address, tokenId *big.Int, tokenURI string) (*types.Transaction, error) {
	return _KIP17MetadataMintable.Contract.MintWithTokenURI(&_KIP17MetadataMintable.TransactOpts, to, tokenId, tokenURI)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_KIP17MetadataMintable *KIP17MetadataMintableTransactor) RenounceMinter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KIP17MetadataMintable.contract.Transact(opts, "renounceMinter")
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_KIP17MetadataMintable *KIP17MetadataMintableSession) RenounceMinter() (*types.Transaction, error) {
	return _KIP17MetadataMintable.Contract.RenounceMinter(&_KIP17MetadataMintable.TransactOpts)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_KIP17MetadataMintable *KIP17MetadataMintableTransactorSession) RenounceMinter() (*types.Transaction, error) {
	return _KIP17MetadataMintable.Contract.RenounceMinter(&_KIP17MetadataMintable.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_KIP17MetadataMintable *KIP17MetadataMintableTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17MetadataMintable.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_KIP17MetadataMintable *KIP17MetadataMintableSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17MetadataMintable.Contract.SafeTransferFrom(&_KIP17MetadataMintable.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_KIP17MetadataMintable *KIP17MetadataMintableTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17MetadataMintable.Contract.SafeTransferFrom(&_KIP17MetadataMintable.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_KIP17MetadataMintable *KIP17MetadataMintableTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _KIP17MetadataMintable.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_KIP17MetadataMintable *KIP17MetadataMintableSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _KIP17MetadataMintable.Contract.SafeTransferFrom0(&_KIP17MetadataMintable.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_KIP17MetadataMintable *KIP17MetadataMintableTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _KIP17MetadataMintable.Contract.SafeTransferFrom0(&_KIP17MetadataMintable.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_KIP17MetadataMintable *KIP17MetadataMintableTransactor) SetApprovalForAll(opts *bind.TransactOpts, to common.Address, approved bool) (*types.Transaction, error) {
	return _KIP17MetadataMintable.contract.Transact(opts, "setApprovalForAll", to, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_KIP17MetadataMintable *KIP17MetadataMintableSession) SetApprovalForAll(to common.Address, approved bool) (*types.Transaction, error) {
	return _KIP17MetadataMintable.Contract.SetApprovalForAll(&_KIP17MetadataMintable.TransactOpts, to, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_KIP17MetadataMintable *KIP17MetadataMintableTransactorSession) SetApprovalForAll(to common.Address, approved bool) (*types.Transaction, error) {
	return _KIP17MetadataMintable.Contract.SetApprovalForAll(&_KIP17MetadataMintable.TransactOpts, to, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_KIP17MetadataMintable *KIP17MetadataMintableTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17MetadataMintable.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_KIP17MetadataMintable *KIP17MetadataMintableSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17MetadataMintable.Contract.TransferFrom(&_KIP17MetadataMintable.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_KIP17MetadataMintable *KIP17MetadataMintableTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17MetadataMintable.Contract.TransferFrom(&_KIP17MetadataMintable.TransactOpts, from, to, tokenId)
}

// KIP17MetadataMintableApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the KIP17MetadataMintable contract.
type KIP17MetadataMintableApprovalIterator struct {
	Event *KIP17MetadataMintableApproval // Event containing the contract specifics and raw log

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
func (it *KIP17MetadataMintableApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KIP17MetadataMintableApproval)
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
		it.Event = new(KIP17MetadataMintableApproval)
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
func (it *KIP17MetadataMintableApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KIP17MetadataMintableApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KIP17MetadataMintableApproval represents a Approval event raised by the KIP17MetadataMintable contract.
type KIP17MetadataMintableApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_KIP17MetadataMintable *KIP17MetadataMintableFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*KIP17MetadataMintableApprovalIterator, error) {

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

	logs, sub, err := _KIP17MetadataMintable.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &KIP17MetadataMintableApprovalIterator{contract: _KIP17MetadataMintable.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_KIP17MetadataMintable *KIP17MetadataMintableFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *KIP17MetadataMintableApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _KIP17MetadataMintable.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KIP17MetadataMintableApproval)
				if err := _KIP17MetadataMintable.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_KIP17MetadataMintable *KIP17MetadataMintableFilterer) ParseApproval(log types.Log) (*KIP17MetadataMintableApproval, error) {
	event := new(KIP17MetadataMintableApproval)
	if err := _KIP17MetadataMintable.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KIP17MetadataMintableApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the KIP17MetadataMintable contract.
type KIP17MetadataMintableApprovalForAllIterator struct {
	Event *KIP17MetadataMintableApprovalForAll // Event containing the contract specifics and raw log

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
func (it *KIP17MetadataMintableApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KIP17MetadataMintableApprovalForAll)
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
		it.Event = new(KIP17MetadataMintableApprovalForAll)
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
func (it *KIP17MetadataMintableApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KIP17MetadataMintableApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KIP17MetadataMintableApprovalForAll represents a ApprovalForAll event raised by the KIP17MetadataMintable contract.
type KIP17MetadataMintableApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_KIP17MetadataMintable *KIP17MetadataMintableFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*KIP17MetadataMintableApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _KIP17MetadataMintable.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &KIP17MetadataMintableApprovalForAllIterator{contract: _KIP17MetadataMintable.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_KIP17MetadataMintable *KIP17MetadataMintableFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *KIP17MetadataMintableApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _KIP17MetadataMintable.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KIP17MetadataMintableApprovalForAll)
				if err := _KIP17MetadataMintable.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_KIP17MetadataMintable *KIP17MetadataMintableFilterer) ParseApprovalForAll(log types.Log) (*KIP17MetadataMintableApprovalForAll, error) {
	event := new(KIP17MetadataMintableApprovalForAll)
	if err := _KIP17MetadataMintable.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KIP17MetadataMintableMinterAddedIterator is returned from FilterMinterAdded and is used to iterate over the raw logs and unpacked data for MinterAdded events raised by the KIP17MetadataMintable contract.
type KIP17MetadataMintableMinterAddedIterator struct {
	Event *KIP17MetadataMintableMinterAdded // Event containing the contract specifics and raw log

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
func (it *KIP17MetadataMintableMinterAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KIP17MetadataMintableMinterAdded)
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
		it.Event = new(KIP17MetadataMintableMinterAdded)
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
func (it *KIP17MetadataMintableMinterAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KIP17MetadataMintableMinterAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KIP17MetadataMintableMinterAdded represents a MinterAdded event raised by the KIP17MetadataMintable contract.
type KIP17MetadataMintableMinterAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinterAdded is a free log retrieval operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed account)
func (_KIP17MetadataMintable *KIP17MetadataMintableFilterer) FilterMinterAdded(opts *bind.FilterOpts, account []common.Address) (*KIP17MetadataMintableMinterAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _KIP17MetadataMintable.contract.FilterLogs(opts, "MinterAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &KIP17MetadataMintableMinterAddedIterator{contract: _KIP17MetadataMintable.contract, event: "MinterAdded", logs: logs, sub: sub}, nil
}

// WatchMinterAdded is a free log subscription operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed account)
func (_KIP17MetadataMintable *KIP17MetadataMintableFilterer) WatchMinterAdded(opts *bind.WatchOpts, sink chan<- *KIP17MetadataMintableMinterAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _KIP17MetadataMintable.contract.WatchLogs(opts, "MinterAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KIP17MetadataMintableMinterAdded)
				if err := _KIP17MetadataMintable.contract.UnpackLog(event, "MinterAdded", log); err != nil {
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
func (_KIP17MetadataMintable *KIP17MetadataMintableFilterer) ParseMinterAdded(log types.Log) (*KIP17MetadataMintableMinterAdded, error) {
	event := new(KIP17MetadataMintableMinterAdded)
	if err := _KIP17MetadataMintable.contract.UnpackLog(event, "MinterAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KIP17MetadataMintableMinterRemovedIterator is returned from FilterMinterRemoved and is used to iterate over the raw logs and unpacked data for MinterRemoved events raised by the KIP17MetadataMintable contract.
type KIP17MetadataMintableMinterRemovedIterator struct {
	Event *KIP17MetadataMintableMinterRemoved // Event containing the contract specifics and raw log

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
func (it *KIP17MetadataMintableMinterRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KIP17MetadataMintableMinterRemoved)
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
		it.Event = new(KIP17MetadataMintableMinterRemoved)
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
func (it *KIP17MetadataMintableMinterRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KIP17MetadataMintableMinterRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KIP17MetadataMintableMinterRemoved represents a MinterRemoved event raised by the KIP17MetadataMintable contract.
type KIP17MetadataMintableMinterRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinterRemoved is a free log retrieval operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed account)
func (_KIP17MetadataMintable *KIP17MetadataMintableFilterer) FilterMinterRemoved(opts *bind.FilterOpts, account []common.Address) (*KIP17MetadataMintableMinterRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _KIP17MetadataMintable.contract.FilterLogs(opts, "MinterRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &KIP17MetadataMintableMinterRemovedIterator{contract: _KIP17MetadataMintable.contract, event: "MinterRemoved", logs: logs, sub: sub}, nil
}

// WatchMinterRemoved is a free log subscription operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed account)
func (_KIP17MetadataMintable *KIP17MetadataMintableFilterer) WatchMinterRemoved(opts *bind.WatchOpts, sink chan<- *KIP17MetadataMintableMinterRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _KIP17MetadataMintable.contract.WatchLogs(opts, "MinterRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KIP17MetadataMintableMinterRemoved)
				if err := _KIP17MetadataMintable.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
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
func (_KIP17MetadataMintable *KIP17MetadataMintableFilterer) ParseMinterRemoved(log types.Log) (*KIP17MetadataMintableMinterRemoved, error) {
	event := new(KIP17MetadataMintableMinterRemoved)
	if err := _KIP17MetadataMintable.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KIP17MetadataMintableTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the KIP17MetadataMintable contract.
type KIP17MetadataMintableTransferIterator struct {
	Event *KIP17MetadataMintableTransfer // Event containing the contract specifics and raw log

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
func (it *KIP17MetadataMintableTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KIP17MetadataMintableTransfer)
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
		it.Event = new(KIP17MetadataMintableTransfer)
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
func (it *KIP17MetadataMintableTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KIP17MetadataMintableTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KIP17MetadataMintableTransfer represents a Transfer event raised by the KIP17MetadataMintable contract.
type KIP17MetadataMintableTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_KIP17MetadataMintable *KIP17MetadataMintableFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*KIP17MetadataMintableTransferIterator, error) {

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

	logs, sub, err := _KIP17MetadataMintable.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &KIP17MetadataMintableTransferIterator{contract: _KIP17MetadataMintable.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_KIP17MetadataMintable *KIP17MetadataMintableFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *KIP17MetadataMintableTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _KIP17MetadataMintable.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KIP17MetadataMintableTransfer)
				if err := _KIP17MetadataMintable.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_KIP17MetadataMintable *KIP17MetadataMintableFilterer) ParseTransfer(log types.Log) (*KIP17MetadataMintableTransfer, error) {
	event := new(KIP17MetadataMintableTransfer)
	if err := _KIP17MetadataMintable.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KIP17ServiceChainABI is the input ABI used to generate the binding from.
const KIP17ServiceChainABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_uid\",\"type\":\"uint256\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"requestValueTransfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_bridge\",\"type\":\"address\"}],\"name\":\"setBridge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_bridge\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"}]"

// KIP17ServiceChainBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const KIP17ServiceChainBinRuntime = ``

// KIP17ServiceChainFuncSigs maps the 4-byte function signature to its string representation.
var KIP17ServiceChainFuncSigs = map[string]string{
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"e78cea92": "bridge()",
	"081812fc": "getApproved(uint256)",
	"e985e9c5": "isApprovedForAll(address,address)",
	"8f32d59b": "isOwner()",
	"8da5cb5b": "owner()",
	"6352211e": "ownerOf(uint256)",
	"715018a6": "renounceOwnership()",
	"3f4c4e3d": "requestValueTransfer(uint256,address,bytes)",
	"42842e0e": "safeTransferFrom(address,address,uint256)",
	"b88d4fde": "safeTransferFrom(address,address,uint256,bytes)",
	"a22cb465": "setApprovalForAll(address,bool)",
	"8dd14802": "setBridge(address)",
	"01ffc9a7": "supportsInterface(bytes4)",
	"23b872dd": "transferFrom(address,address,uint256)",
	"f2fde38b": "transferOwnership(address)",
}

// KIP17ServiceChain is an auto generated Go binding around a Klaytn contract.
type KIP17ServiceChain struct {
	KIP17ServiceChainCaller     // Read-only binding to the contract
	KIP17ServiceChainTransactor // Write-only binding to the contract
	KIP17ServiceChainFilterer   // Log filterer for contract events
}

// KIP17ServiceChainCaller is an auto generated read-only Go binding around a Klaytn contract.
type KIP17ServiceChainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KIP17ServiceChainTransactor is an auto generated write-only Go binding around a Klaytn contract.
type KIP17ServiceChainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KIP17ServiceChainFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type KIP17ServiceChainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KIP17ServiceChainSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type KIP17ServiceChainSession struct {
	Contract     *KIP17ServiceChain // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// KIP17ServiceChainCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type KIP17ServiceChainCallerSession struct {
	Contract *KIP17ServiceChainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// KIP17ServiceChainTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type KIP17ServiceChainTransactorSession struct {
	Contract     *KIP17ServiceChainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// KIP17ServiceChainRaw is an auto generated low-level Go binding around a Klaytn contract.
type KIP17ServiceChainRaw struct {
	Contract *KIP17ServiceChain // Generic contract binding to access the raw methods on
}

// KIP17ServiceChainCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type KIP17ServiceChainCallerRaw struct {
	Contract *KIP17ServiceChainCaller // Generic read-only contract binding to access the raw methods on
}

// KIP17ServiceChainTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type KIP17ServiceChainTransactorRaw struct {
	Contract *KIP17ServiceChainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKIP17ServiceChain creates a new instance of KIP17ServiceChain, bound to a specific deployed contract.
func NewKIP17ServiceChain(address common.Address, backend bind.ContractBackend) (*KIP17ServiceChain, error) {
	contract, err := bindKIP17ServiceChain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KIP17ServiceChain{KIP17ServiceChainCaller: KIP17ServiceChainCaller{contract: contract}, KIP17ServiceChainTransactor: KIP17ServiceChainTransactor{contract: contract}, KIP17ServiceChainFilterer: KIP17ServiceChainFilterer{contract: contract}}, nil
}

// NewKIP17ServiceChainCaller creates a new read-only instance of KIP17ServiceChain, bound to a specific deployed contract.
func NewKIP17ServiceChainCaller(address common.Address, caller bind.ContractCaller) (*KIP17ServiceChainCaller, error) {
	contract, err := bindKIP17ServiceChain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KIP17ServiceChainCaller{contract: contract}, nil
}

// NewKIP17ServiceChainTransactor creates a new write-only instance of KIP17ServiceChain, bound to a specific deployed contract.
func NewKIP17ServiceChainTransactor(address common.Address, transactor bind.ContractTransactor) (*KIP17ServiceChainTransactor, error) {
	contract, err := bindKIP17ServiceChain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KIP17ServiceChainTransactor{contract: contract}, nil
}

// NewKIP17ServiceChainFilterer creates a new log filterer instance of KIP17ServiceChain, bound to a specific deployed contract.
func NewKIP17ServiceChainFilterer(address common.Address, filterer bind.ContractFilterer) (*KIP17ServiceChainFilterer, error) {
	contract, err := bindKIP17ServiceChain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KIP17ServiceChainFilterer{contract: contract}, nil
}

// bindKIP17ServiceChain binds a generic wrapper to an already deployed contract.
func bindKIP17ServiceChain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KIP17ServiceChainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KIP17ServiceChain *KIP17ServiceChainRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KIP17ServiceChain.Contract.KIP17ServiceChainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KIP17ServiceChain *KIP17ServiceChainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KIP17ServiceChain.Contract.KIP17ServiceChainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KIP17ServiceChain *KIP17ServiceChainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KIP17ServiceChain.Contract.KIP17ServiceChainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KIP17ServiceChain *KIP17ServiceChainCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KIP17ServiceChain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KIP17ServiceChain *KIP17ServiceChainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KIP17ServiceChain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KIP17ServiceChain *KIP17ServiceChainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KIP17ServiceChain.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_KIP17ServiceChain *KIP17ServiceChainCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KIP17ServiceChain.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_KIP17ServiceChain *KIP17ServiceChainSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _KIP17ServiceChain.Contract.BalanceOf(&_KIP17ServiceChain.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_KIP17ServiceChain *KIP17ServiceChainCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _KIP17ServiceChain.Contract.BalanceOf(&_KIP17ServiceChain.CallOpts, owner)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_KIP17ServiceChain *KIP17ServiceChainCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KIP17ServiceChain.contract.Call(opts, out, "bridge")
	return *ret0, err
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_KIP17ServiceChain *KIP17ServiceChainSession) Bridge() (common.Address, error) {
	return _KIP17ServiceChain.Contract.Bridge(&_KIP17ServiceChain.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_KIP17ServiceChain *KIP17ServiceChainCallerSession) Bridge() (common.Address, error) {
	return _KIP17ServiceChain.Contract.Bridge(&_KIP17ServiceChain.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_KIP17ServiceChain *KIP17ServiceChainCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KIP17ServiceChain.contract.Call(opts, out, "getApproved", tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_KIP17ServiceChain *KIP17ServiceChainSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _KIP17ServiceChain.Contract.GetApproved(&_KIP17ServiceChain.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_KIP17ServiceChain *KIP17ServiceChainCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _KIP17ServiceChain.Contract.GetApproved(&_KIP17ServiceChain.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_KIP17ServiceChain *KIP17ServiceChainCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KIP17ServiceChain.contract.Call(opts, out, "isApprovedForAll", owner, operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_KIP17ServiceChain *KIP17ServiceChainSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _KIP17ServiceChain.Contract.IsApprovedForAll(&_KIP17ServiceChain.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_KIP17ServiceChain *KIP17ServiceChainCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _KIP17ServiceChain.Contract.IsApprovedForAll(&_KIP17ServiceChain.CallOpts, owner, operator)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_KIP17ServiceChain *KIP17ServiceChainCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KIP17ServiceChain.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_KIP17ServiceChain *KIP17ServiceChainSession) IsOwner() (bool, error) {
	return _KIP17ServiceChain.Contract.IsOwner(&_KIP17ServiceChain.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_KIP17ServiceChain *KIP17ServiceChainCallerSession) IsOwner() (bool, error) {
	return _KIP17ServiceChain.Contract.IsOwner(&_KIP17ServiceChain.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_KIP17ServiceChain *KIP17ServiceChainCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KIP17ServiceChain.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_KIP17ServiceChain *KIP17ServiceChainSession) Owner() (common.Address, error) {
	return _KIP17ServiceChain.Contract.Owner(&_KIP17ServiceChain.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_KIP17ServiceChain *KIP17ServiceChainCallerSession) Owner() (common.Address, error) {
	return _KIP17ServiceChain.Contract.Owner(&_KIP17ServiceChain.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_KIP17ServiceChain *KIP17ServiceChainCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KIP17ServiceChain.contract.Call(opts, out, "ownerOf", tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_KIP17ServiceChain *KIP17ServiceChainSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _KIP17ServiceChain.Contract.OwnerOf(&_KIP17ServiceChain.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_KIP17ServiceChain *KIP17ServiceChainCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _KIP17ServiceChain.Contract.OwnerOf(&_KIP17ServiceChain.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_KIP17ServiceChain *KIP17ServiceChainCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KIP17ServiceChain.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_KIP17ServiceChain *KIP17ServiceChainSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _KIP17ServiceChain.Contract.SupportsInterface(&_KIP17ServiceChain.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_KIP17ServiceChain *KIP17ServiceChainCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _KIP17ServiceChain.Contract.SupportsInterface(&_KIP17ServiceChain.CallOpts, interfaceId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_KIP17ServiceChain *KIP17ServiceChainTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17ServiceChain.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_KIP17ServiceChain *KIP17ServiceChainSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17ServiceChain.Contract.Approve(&_KIP17ServiceChain.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_KIP17ServiceChain *KIP17ServiceChainTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17ServiceChain.Contract.Approve(&_KIP17ServiceChain.TransactOpts, to, tokenId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_KIP17ServiceChain *KIP17ServiceChainTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KIP17ServiceChain.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_KIP17ServiceChain *KIP17ServiceChainSession) RenounceOwnership() (*types.Transaction, error) {
	return _KIP17ServiceChain.Contract.RenounceOwnership(&_KIP17ServiceChain.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_KIP17ServiceChain *KIP17ServiceChainTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _KIP17ServiceChain.Contract.RenounceOwnership(&_KIP17ServiceChain.TransactOpts)
}

// RequestValueTransfer is a paid mutator transaction binding the contract method 0x3f4c4e3d.
//
// Solidity: function requestValueTransfer(uint256 _uid, address _to, bytes _extraData) returns()
func (_KIP17ServiceChain *KIP17ServiceChainTransactor) RequestValueTransfer(opts *bind.TransactOpts, _uid *big.Int, _to common.Address, _extraData []byte) (*types.Transaction, error) {
	return _KIP17ServiceChain.contract.Transact(opts, "requestValueTransfer", _uid, _to, _extraData)
}

// RequestValueTransfer is a paid mutator transaction binding the contract method 0x3f4c4e3d.
//
// Solidity: function requestValueTransfer(uint256 _uid, address _to, bytes _extraData) returns()
func (_KIP17ServiceChain *KIP17ServiceChainSession) RequestValueTransfer(_uid *big.Int, _to common.Address, _extraData []byte) (*types.Transaction, error) {
	return _KIP17ServiceChain.Contract.RequestValueTransfer(&_KIP17ServiceChain.TransactOpts, _uid, _to, _extraData)
}

// RequestValueTransfer is a paid mutator transaction binding the contract method 0x3f4c4e3d.
//
// Solidity: function requestValueTransfer(uint256 _uid, address _to, bytes _extraData) returns()
func (_KIP17ServiceChain *KIP17ServiceChainTransactorSession) RequestValueTransfer(_uid *big.Int, _to common.Address, _extraData []byte) (*types.Transaction, error) {
	return _KIP17ServiceChain.Contract.RequestValueTransfer(&_KIP17ServiceChain.TransactOpts, _uid, _to, _extraData)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_KIP17ServiceChain *KIP17ServiceChainTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17ServiceChain.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_KIP17ServiceChain *KIP17ServiceChainSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17ServiceChain.Contract.SafeTransferFrom(&_KIP17ServiceChain.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_KIP17ServiceChain *KIP17ServiceChainTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17ServiceChain.Contract.SafeTransferFrom(&_KIP17ServiceChain.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_KIP17ServiceChain *KIP17ServiceChainTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _KIP17ServiceChain.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_KIP17ServiceChain *KIP17ServiceChainSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _KIP17ServiceChain.Contract.SafeTransferFrom0(&_KIP17ServiceChain.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_KIP17ServiceChain *KIP17ServiceChainTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _KIP17ServiceChain.Contract.SafeTransferFrom0(&_KIP17ServiceChain.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_KIP17ServiceChain *KIP17ServiceChainTransactor) SetApprovalForAll(opts *bind.TransactOpts, to common.Address, approved bool) (*types.Transaction, error) {
	return _KIP17ServiceChain.contract.Transact(opts, "setApprovalForAll", to, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_KIP17ServiceChain *KIP17ServiceChainSession) SetApprovalForAll(to common.Address, approved bool) (*types.Transaction, error) {
	return _KIP17ServiceChain.Contract.SetApprovalForAll(&_KIP17ServiceChain.TransactOpts, to, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_KIP17ServiceChain *KIP17ServiceChainTransactorSession) SetApprovalForAll(to common.Address, approved bool) (*types.Transaction, error) {
	return _KIP17ServiceChain.Contract.SetApprovalForAll(&_KIP17ServiceChain.TransactOpts, to, approved)
}

// SetBridge is a paid mutator transaction binding the contract method 0x8dd14802.
//
// Solidity: function setBridge(address _bridge) returns()
func (_KIP17ServiceChain *KIP17ServiceChainTransactor) SetBridge(opts *bind.TransactOpts, _bridge common.Address) (*types.Transaction, error) {
	return _KIP17ServiceChain.contract.Transact(opts, "setBridge", _bridge)
}

// SetBridge is a paid mutator transaction binding the contract method 0x8dd14802.
//
// Solidity: function setBridge(address _bridge) returns()
func (_KIP17ServiceChain *KIP17ServiceChainSession) SetBridge(_bridge common.Address) (*types.Transaction, error) {
	return _KIP17ServiceChain.Contract.SetBridge(&_KIP17ServiceChain.TransactOpts, _bridge)
}

// SetBridge is a paid mutator transaction binding the contract method 0x8dd14802.
//
// Solidity: function setBridge(address _bridge) returns()
func (_KIP17ServiceChain *KIP17ServiceChainTransactorSession) SetBridge(_bridge common.Address) (*types.Transaction, error) {
	return _KIP17ServiceChain.Contract.SetBridge(&_KIP17ServiceChain.TransactOpts, _bridge)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_KIP17ServiceChain *KIP17ServiceChainTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17ServiceChain.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_KIP17ServiceChain *KIP17ServiceChainSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17ServiceChain.Contract.TransferFrom(&_KIP17ServiceChain.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_KIP17ServiceChain *KIP17ServiceChainTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KIP17ServiceChain.Contract.TransferFrom(&_KIP17ServiceChain.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KIP17ServiceChain *KIP17ServiceChainTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _KIP17ServiceChain.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KIP17ServiceChain *KIP17ServiceChainSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _KIP17ServiceChain.Contract.TransferOwnership(&_KIP17ServiceChain.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KIP17ServiceChain *KIP17ServiceChainTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _KIP17ServiceChain.Contract.TransferOwnership(&_KIP17ServiceChain.TransactOpts, newOwner)
}

// KIP17ServiceChainApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the KIP17ServiceChain contract.
type KIP17ServiceChainApprovalIterator struct {
	Event *KIP17ServiceChainApproval // Event containing the contract specifics and raw log

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
func (it *KIP17ServiceChainApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KIP17ServiceChainApproval)
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
		it.Event = new(KIP17ServiceChainApproval)
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
func (it *KIP17ServiceChainApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KIP17ServiceChainApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KIP17ServiceChainApproval represents a Approval event raised by the KIP17ServiceChain contract.
type KIP17ServiceChainApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_KIP17ServiceChain *KIP17ServiceChainFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*KIP17ServiceChainApprovalIterator, error) {

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

	logs, sub, err := _KIP17ServiceChain.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &KIP17ServiceChainApprovalIterator{contract: _KIP17ServiceChain.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_KIP17ServiceChain *KIP17ServiceChainFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *KIP17ServiceChainApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _KIP17ServiceChain.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KIP17ServiceChainApproval)
				if err := _KIP17ServiceChain.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_KIP17ServiceChain *KIP17ServiceChainFilterer) ParseApproval(log types.Log) (*KIP17ServiceChainApproval, error) {
	event := new(KIP17ServiceChainApproval)
	if err := _KIP17ServiceChain.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KIP17ServiceChainApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the KIP17ServiceChain contract.
type KIP17ServiceChainApprovalForAllIterator struct {
	Event *KIP17ServiceChainApprovalForAll // Event containing the contract specifics and raw log

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
func (it *KIP17ServiceChainApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KIP17ServiceChainApprovalForAll)
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
		it.Event = new(KIP17ServiceChainApprovalForAll)
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
func (it *KIP17ServiceChainApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KIP17ServiceChainApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KIP17ServiceChainApprovalForAll represents a ApprovalForAll event raised by the KIP17ServiceChain contract.
type KIP17ServiceChainApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_KIP17ServiceChain *KIP17ServiceChainFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*KIP17ServiceChainApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _KIP17ServiceChain.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &KIP17ServiceChainApprovalForAllIterator{contract: _KIP17ServiceChain.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_KIP17ServiceChain *KIP17ServiceChainFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *KIP17ServiceChainApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _KIP17ServiceChain.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KIP17ServiceChainApprovalForAll)
				if err := _KIP17ServiceChain.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_KIP17ServiceChain *KIP17ServiceChainFilterer) ParseApprovalForAll(log types.Log) (*KIP17ServiceChainApprovalForAll, error) {
	event := new(KIP17ServiceChainApprovalForAll)
	if err := _KIP17ServiceChain.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KIP17ServiceChainOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the KIP17ServiceChain contract.
type KIP17ServiceChainOwnershipTransferredIterator struct {
	Event *KIP17ServiceChainOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *KIP17ServiceChainOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KIP17ServiceChainOwnershipTransferred)
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
		it.Event = new(KIP17ServiceChainOwnershipTransferred)
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
func (it *KIP17ServiceChainOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KIP17ServiceChainOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KIP17ServiceChainOwnershipTransferred represents a OwnershipTransferred event raised by the KIP17ServiceChain contract.
type KIP17ServiceChainOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_KIP17ServiceChain *KIP17ServiceChainFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*KIP17ServiceChainOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _KIP17ServiceChain.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &KIP17ServiceChainOwnershipTransferredIterator{contract: _KIP17ServiceChain.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_KIP17ServiceChain *KIP17ServiceChainFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *KIP17ServiceChainOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _KIP17ServiceChain.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KIP17ServiceChainOwnershipTransferred)
				if err := _KIP17ServiceChain.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_KIP17ServiceChain *KIP17ServiceChainFilterer) ParseOwnershipTransferred(log types.Log) (*KIP17ServiceChainOwnershipTransferred, error) {
	event := new(KIP17ServiceChainOwnershipTransferred)
	if err := _KIP17ServiceChain.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KIP17ServiceChainTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the KIP17ServiceChain contract.
type KIP17ServiceChainTransferIterator struct {
	Event *KIP17ServiceChainTransfer // Event containing the contract specifics and raw log

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
func (it *KIP17ServiceChainTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KIP17ServiceChainTransfer)
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
		it.Event = new(KIP17ServiceChainTransfer)
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
func (it *KIP17ServiceChainTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KIP17ServiceChainTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KIP17ServiceChainTransfer represents a Transfer event raised by the KIP17ServiceChain contract.
type KIP17ServiceChainTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_KIP17ServiceChain *KIP17ServiceChainFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*KIP17ServiceChainTransferIterator, error) {

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

	logs, sub, err := _KIP17ServiceChain.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &KIP17ServiceChainTransferIterator{contract: _KIP17ServiceChain.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_KIP17ServiceChain *KIP17ServiceChainFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *KIP17ServiceChainTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _KIP17ServiceChain.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KIP17ServiceChainTransfer)
				if err := _KIP17ServiceChain.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_KIP17ServiceChain *KIP17ServiceChainFilterer) ParseTransfer(log types.Log) (*KIP17ServiceChainTransfer, error) {
	event := new(KIP17ServiceChainTransfer)
	if err := _KIP17ServiceChain.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ServiceChainKIP17NFTABI is the input ABI used to generate the binding from.
const ServiceChainKIP17NFTABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_uid\",\"type\":\"uint256\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"requestValueTransfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"},{\"name\":\"tokenURI\",\"type\":\"string\"}],\"name\":\"mintWithTokenURI\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_user\",\"type\":\"address\"},{\"name\":\"_startID\",\"type\":\"uint256\"},{\"name\":\"_endID\",\"type\":\"uint256\"}],\"name\":\"registerBulk\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_bridge\",\"type\":\"address\"}],\"name\":\"setBridge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addMinter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceMinter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isMinter\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_bridge\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MinterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MinterRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"}]"

// ServiceChainKIP17NFTBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const ServiceChainKIP17NFTBinRuntime = `608060405234801561001057600080fd5b50600436106101cf5760003560e01c8063715018a61161010457806398650275116100a2578063c87b56dd11610071578063c87b56dd14610722578063e78cea921461073f578063e985e9c514610747578063f2fde38b14610775576101cf565b80639865027514610600578063a22cb46514610608578063aa271e1a14610636578063b88d4fde1461065c576101cf565b80638dd14802116100de5780638dd14802146105a45780638f32d59b146105ca57806395d89b41146105d2578063983b2d56146105da576101cf565b8063715018a6146105625780637a9adac61461056a5780638da5cb5b1461059c576101cf565b80633f4c4e3d116101715780634f6ccce71161014b5780634f6ccce71461044757806350bb4e7f146104645780636352211e1461051f57806370a082311461053c576101cf565b80633f4c4e3d1461036f57806342842e0e146103f457806342966c681461042a576101cf565b8063095ea7b3116101ad578063095ea7b3146102c557806318160ddd146102f357806323b872dd1461030d5780632f745c5914610343576101cf565b806301ffc9a7146101d457806306fdde031461020f578063081812fc1461028c575b600080fd5b6101fb600480360360208110156101ea57600080fd5b50356001600160e01b03191661079b565b604080519115158252519081900360200190f35b6102176107ba565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610251578181015183820152602001610239565b50505050905090810190601f16801561027e5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6102a9600480360360208110156102a257600080fd5b5035610851565b604080516001600160a01b039092168252519081900360200190f35b6102f1600480360360408110156102db57600080fd5b506001600160a01b0381351690602001356108b6565b005b6102fb6109e3565b60408051918252519081900360200190f35b6102f16004803603606081101561032357600080fd5b506001600160a01b038135811691602081013590911690604001356109e9565b6102fb6004803603604081101561035957600080fd5b506001600160a01b038135169060200135610a41565b6102f16004803603606081101561038557600080fd5b8135916001600160a01b03602082013516918101906060810160408201356401000000008111156103b557600080fd5b8201836020820111156103c757600080fd5b803590602001918460018302840111640100000000831117156103e957600080fd5b509092509050610ac3565b6102f16004803603606081101561040a57600080fd5b506001600160a01b03813581169160208101359091169060400135610b90565b6102f16004803603602081101561044057600080fd5b5035610bab565b6102fb6004803603602081101561045d57600080fd5b5035610bff565b6101fb6004803603606081101561047a57600080fd5b6001600160a01b03823516916020810135918101906060810160408201356401000000008111156104aa57600080fd5b8201836020820111156104bc57600080fd5b803590602001918460018302840111640100000000831117156104de57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610c68945050505050565b6102a96004803603602081101561053557600080fd5b5035610ccf565b6102fb6004803603602081101561055257600080fd5b50356001600160a01b0316610d2c565b6102f1610d97565b6102f16004803603606081101561058057600080fd5b506001600160a01b038135169060208101359060400135610e2b565b6102a9610ebb565b6102f1600480360360208110156105ba57600080fd5b50356001600160a01b0316610eca565b6101fb610f36565b610217610f47565b6102f1600480360360208110156105f057600080fd5b50356001600160a01b0316610fa8565b6102f1610ff8565b6102f16004803603604081101561061e57600080fd5b506001600160a01b0381351690602001351515611003565b6101fb6004803603602081101561064c57600080fd5b50356001600160a01b03166110d2565b6102f16004803603608081101561067257600080fd5b6001600160a01b038235811692602081013590911691604082013591908101906080810160608201356401000000008111156106ad57600080fd5b8201836020820111156106bf57600080fd5b803590602001918460018302840111640100000000831117156106e157600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506110e5945050505050565b6102176004803603602081101561073857600080fd5b503561113a565b6102a9611222565b6101fb6004803603604081101561075d57600080fd5b506001600160a01b0381358116916020013516611231565b6102f16004803603602081101561078b57600080fd5b50356001600160a01b031661125f565b6001600160e01b03191660009081526020819052604090205460ff1690565b60098054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156108465780601f1061081b57610100808354040283529160200191610846565b820191906000526020600020905b81548152906001019060200180831161082957829003601f168201915b505050505090505b90565b600061085c826112b2565b61089a57604051600160e51b62461bcd02815260040180806020018281038252602b8152602001806124f7602b913960400191505060405180910390fd5b506000908152600260205260409020546001600160a01b031690565b60006108c182610ccf565b9050806001600160a01b0316836001600160a01b0316141561092d5760408051600160e51b62461bcd02815260206004820181905260248201527f4b495031373a20617070726f76616c20746f2063757272656e74206f776e6572604482015290519081900360640190fd5b336001600160a01b038216148061094957506109498133611231565b61098757604051600160e51b62461bcd0281526004018080602001828103825260378152602001806125226037913960400191505060405180910390fd5b60008281526002602052604080822080546001600160a01b0319166001600160a01b0387811691821790925591518593918516917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591a4505050565b60075490565b6109f333826112cf565b610a3157604051600160e51b62461bcd0281526004018080602001828103825260308152602001806123866030913960400191505060405180910390fd5b610a3c838383611376565b505050565b6000610a4c83610d2c565b8210610a8c57604051600160e51b62461bcd02815260040180806020018281038252602a8152602001806122b8602a913960400191505060405180910390fd5b6001600160a01b0383166000908152600560205260409020805483908110610ab057fe5b9060005260206000200154905092915050565b600e54610adb9033906001600160a01b0316866109e9565b600e54604051600160e01b6315c281770281523360048201818152602483018890526001600160a01b038781166044850152608060648501908152608485018790529416936315c28177938992899289928992919060a401848480828437600081840152601f19601f8201169050808301925050509650505050505050600060405180830381600087803b158015610b7257600080fd5b505af1158015610b86573d6000803e3d6000fd5b5050505050505050565b610a3c838383604051806020016040528060008152506110e5565b610bb533826112cf565b610bf357604051600160e51b62461bcd02815260040180806020018281038252602f815260200180612428602f913960400191505060405180910390fd5b610bfc81611395565b50565b6000610c096109e3565b8210610c4957604051600160e51b62461bcd02815260040180806020018281038252602b8152602001806124a8602b913960400191505060405180910390fd5b60078281548110610c5657fe5b90600052602060002001549050919050565b6000610c73336110d2565b610cb157604051600160e51b62461bcd0281526004018080602001828103825260308152602001806122e26030913960400191505060405180910390fd5b610cbb84846113a7565b610cc583836113c8565b5060019392505050565b6000818152600160205260408120546001600160a01b031680610d2657604051600160e51b62461bcd0281526004018080602001828103825260288152602001806123336028913960400191505060405180910390fd5b92915050565b60006001600160a01b038216610d7657604051600160e51b62461bcd0281526004018080602001828103825260298152602001806124576029913960400191505060405180910390fd5b6001600160a01b0382166000908152600360205260409020610d269061142e565b610d9f610f36565b610de15760408051600160e51b62461bcd02815260206004820181905260248201526000805160206123b6833981519152604482015290519081900360640190fd5b600d546040516000916001600160a01b0316907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600d80546001600160a01b0319169055565b610e33610f36565b610e755760408051600160e51b62461bcd02815260206004820181905260248201526000805160206123b6833981519152604482015290519081900360640190fd5b815b81811015610eb557610eac8482604051806040016040528060078152602001600160c81b667465737455524902815250610c68565b50600101610e77565b50505050565b600d546001600160a01b031690565b610ed2610f36565b610f145760408051600160e51b62461bcd02815260206004820181905260248201526000805160206123b6833981519152604482015290519081900360640190fd5b600e80546001600160a01b0319166001600160a01b0392909216919091179055565b600d546001600160a01b0316331490565b600a8054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156108465780601f1061081b57610100808354040283529160200191610846565b610fb1336110d2565b610fef57604051600160e51b62461bcd0281526004018080602001828103825260308152602001806122e26030913960400191505060405180910390fd5b610bfc81611432565b6110013361147a565b565b6001600160a01b0382163314156110645760408051600160e51b62461bcd02815260206004820152601860248201527f4b495031373a20617070726f766520746f2063616c6c65720000000000000000604482015290519081900360640190fd5b3360008181526004602090815260408083206001600160a01b03871680855290835292819020805460ff1916861515908117909155815190815290519293927f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31929181900390910190a35050565b6000610d26600c8363ffffffff6114c216565b6110f08484846109e9565b6110fc8484848461152c565b610eb557604051600160e51b62461bcd0281526004018080602001828103825260308152602001806123f86030913960400191505060405180910390fd5b6060611145826112b2565b61118357604051600160e51b62461bcd02815260040180806020018281038252602e815260200180612241602e913960400191505060405180910390fd5b6000828152600b602090815260409182902080548351601f6002600019610100600186161502019093169290920491820184900484028101840190945280845290918301828280156112165780601f106111eb57610100808354040283529160200191611216565b820191906000526020600020905b8154815290600101906020018083116111f957829003601f168201915b50505050509050919050565b600e546001600160a01b031681565b6001600160a01b03918216600090815260046020908152604080832093909416825291909152205460ff1690565b611267610f36565b6112a95760408051600160e51b62461bcd02815260206004820181905260248201526000805160206123b6833981519152604482015290519081900360640190fd5b610bfc8161190f565b6000908152600160205260409020546001600160a01b0316151590565b60006112da826112b2565b61131857604051600160e51b62461bcd02815260040180806020018281038252602b815260200180612559602b913960400191505060405180910390fd5b600061132383610ccf565b9050806001600160a01b0316846001600160a01b0316148061135e5750836001600160a01b031661135384610851565b6001600160a01b0316145b8061136e575061136e8185611231565b949350505050565b6113818383836119b3565b61138b8382611afd565b610a3c8282611bf2565b610bfc6113a182610ccf565b82611c30565b6113b18282611c78565b6113bb8282611bf2565b6113c481611daf565b5050565b6113d1826112b2565b61140f57604051600160e51b62461bcd02815260040180806020018281038252602b81526020018061235b602b913960400191505060405180910390fd5b6000828152600b602090815260409091208251610a3c92840190612144565b5490565b611443600c8263ffffffff611df316565b6040516001600160a01b038216907f6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f690600090a250565b61148b600c8263ffffffff611e7716565b6040516001600160a01b038216907fe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb6669290600090a250565b60006001600160a01b03821661150c57604051600160e51b62461bcd0281526004018080602001828103825260228152602001806123d66022913960400191505060405180910390fd5b506001600160a01b03166000908152602091909152604090205460ff1690565b6000806060611543866001600160a01b0316611ee1565b6115525760019250505061136e565b856001600160a01b031663150b7a0260e01b3389888860405160240180856001600160a01b03166001600160a01b03168152602001846001600160a01b03166001600160a01b0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b838110156115de5781810151838201526020016115c6565b50505050905090810190601f16801561160b5780820380516001836020036101000a031916815260200191505b5060408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909a16999099178952518151919890975087965094509250829150849050835b602083106116735780518252601f199092019160209182019101611654565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d80600081146116d5576040519150601f19603f3d011682016040523d82523d6000602084013e6116da565b606091505b50805191935091501580159061171a57508051600160e11b630a85bd0102906020808401919081101561170c57600080fd5b50516001600160e01b031916145b1561172a5760019250505061136e565b856001600160a01b0316636745782b60e01b3389888860405160240180856001600160a01b03166001600160a01b03168152602001846001600160a01b03166001600160a01b0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b838110156117b657818101518382015260200161179e565b50505050905090810190601f1680156117e35780820380516001836020036101000a031916815260200191505b5060408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909a16999099178952518151919890975087965094509250829150849050835b6020831061184b5780518252601f19909201916020918201910161182c565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d80600081146118ad576040519150601f19603f3d011682016040523d82523d6000602084013e6118b2565b606091505b5080519193509150158015906118f257508051600160e01b636745782b0290602080840191908110156118e457600080fd5b50516001600160e01b031916145b156119025760019250505061136e565b5060009695505050505050565b6001600160a01b03811661195757604051600160e51b62461bcd02815260040180806020018281038252602681526020018061226f6026913960400191505060405180910390fd5b600d546040516001600160a01b038084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3600d80546001600160a01b0319166001600160a01b0392909216919091179055565b826001600160a01b03166119c682610ccf565b6001600160a01b031614611a0e57604051600160e51b62461bcd0281526004018080602001828103825260288152602001806124806028913960400191505060405180910390fd5b6001600160a01b038216611a5657604051600160e51b62461bcd0281526004018080602001828103825260238152602001806122956023913960400191505060405180910390fd5b611a5f81611ee7565b6001600160a01b0383166000908152600360205260409020611a8090611f22565b6001600160a01b0382166000908152600360205260409020611aa190611f39565b60008181526001602052604080822080546001600160a01b0319166001600160a01b0386811691821790925591518493918716917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b6001600160a01b038216600090815260056020526040812054611b2790600163ffffffff611f4216565b600083815260066020526040902054909150808214611bc2576001600160a01b0384166000908152600560205260408120805484908110611b6457fe5b906000526020600020015490508060056000876001600160a01b03166001600160a01b031681526020019081526020016000208381548110611ba257fe5b600091825260208083209091019290925591825260069052604090208190555b6001600160a01b0384166000908152600560205260409020805490611beb9060001983016121c2565b5050505050565b6001600160a01b0390911660009081526005602081815260408084208054868652600684529185208290559282526001810183559183529091200155565b611c3a8282611fa2565b6000818152600b602052604090205460026000196101006001841615020190911604156113c4576000818152600b602052604081206113c4916121e6565b6001600160a01b038216611cd65760408051600160e51b62461bcd02815260206004820152601f60248201527f4b495031373a206d696e7420746f20746865207a65726f206164647265737300604482015290519081900360640190fd5b611cdf816112b2565b15611d345760408051600160e51b62461bcd02815260206004820152601b60248201527f4b495031373a20746f6b656e20616c7265616479206d696e7465640000000000604482015290519081900360640190fd5b600081815260016020908152604080832080546001600160a01b0319166001600160a01b038716908117909155835260039091529020611d7390611f39565b60405181906001600160a01b038416906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908290a45050565b600780546000838152600860205260408120829055600182018355919091527fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c6880155565b611dfd82826114c2565b15611e525760408051600160e51b62461bcd02815260206004820152601f60248201527f526f6c65733a206163636f756e7420616c72656164792068617320726f6c6500604482015290519081900360640190fd5b6001600160a01b0316600090815260209190915260409020805460ff19166001179055565b611e8182826114c2565b611ebf57604051600160e51b62461bcd0281526004018080602001828103825260218152602001806123126021913960400191505060405180910390fd5b6001600160a01b0316600090815260209190915260409020805460ff19169055565b3b151590565b6000818152600260205260409020546001600160a01b031615610bfc57600090815260026020526040902080546001600160a01b0319169055565b8054611f3590600163ffffffff611f4216565b9055565b80546001019055565b600082821115611f9c5760408051600160e51b62461bcd02815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b611fac8282611fce565b611fb68282611afd565b6000818152600660205260408120556113c4816120a8565b816001600160a01b0316611fe182610ccf565b6001600160a01b03161461202957604051600160e51b62461bcd0281526004018080602001828103825260248152602001806124d36024913960400191505060405180910390fd5b61203281611ee7565b6001600160a01b038216600090815260036020526040902061205390611f22565b60008181526001602052604080822080546001600160a01b0319169055518291906001600160a01b038516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908390a45050565b6007546000906120bf90600163ffffffff611f4216565b600083815260086020526040812054600780549394509092849081106120e157fe5b9060005260206000200154905080600783815481106120fc57fe5b6000918252602080832090910192909255828152600890915260409020829055600780549061212f9060001983016121c2565b50505060009182525060086020526040812055565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061218557805160ff19168380011785556121b2565b828001600101855582156121b2579182015b828111156121b2578251825591602001919060010190612197565b506121be929150612226565b5090565b815481835581811115610a3c57600083815260209020610a3c918101908301612226565b50805460018160011615610100020316600290046000825580601f1061220c5750610bfc565b601f016020900490600052602060002090810190610bfc91905b61084e91905b808211156121be576000815560010161222c56fe4b495031374d657461646174613a2055524920717565727920666f72206e6f6e6578697374656e7420746f6b656e4f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573734b495031373a207472616e7366657220746f20746865207a65726f20616464726573734b49503137456e756d657261626c653a206f776e657220696e646578206f7574206f6620626f756e64734d696e746572526f6c653a2063616c6c657220646f6573206e6f74206861766520746865204d696e74657220726f6c65526f6c65733a206163636f756e7420646f6573206e6f74206861766520726f6c654b495031373a206f776e657220717565727920666f72206e6f6e6578697374656e7420746f6b656e4b495031374d657461646174613a2055524920736574206f66206e6f6e6578697374656e7420746f6b656e4b495031373a207472616e736665722063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f7665644f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572526f6c65733a206163636f756e7420697320746865207a65726f20616464726573734b495031373a207472616e7366657220746f206e6f6e204b49503137526563656976657220696d706c656d656e7465724b495031374275726e61626c653a2063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f7665644b495031373a2062616c616e636520717565727920666f7220746865207a65726f20616464726573734b495031373a207472616e73666572206f6620746f6b656e2074686174206973206e6f74206f776e4b49503137456e756d657261626c653a20676c6f62616c20696e646578206f7574206f6620626f756e64734b495031373a206275726e206f6620746f6b656e2074686174206973206e6f74206f776e4b495031373a20617070726f76656420717565727920666f72206e6f6e6578697374656e7420746f6b656e4b495031373a20617070726f76652063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f76656420666f7220616c6c4b495031373a206f70657261746f7220717565727920666f72206e6f6e6578697374656e7420746f6b656ea165627a7a723058207f8f319e1745f9ad2fbd7e2330197c753ae01d8e90d01afb71a033060e7465710029`

// ServiceChainKIP17NFTFuncSigs maps the 4-byte function signature to its string representation.
var ServiceChainKIP17NFTFuncSigs = map[string]string{
	"983b2d56": "addMinter(address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"e78cea92": "bridge()",
	"42966c68": "burn(uint256)",
	"081812fc": "getApproved(uint256)",
	"e985e9c5": "isApprovedForAll(address,address)",
	"aa271e1a": "isMinter(address)",
	"8f32d59b": "isOwner()",
	"50bb4e7f": "mintWithTokenURI(address,uint256,string)",
	"06fdde03": "name()",
	"8da5cb5b": "owner()",
	"6352211e": "ownerOf(uint256)",
	"7a9adac6": "registerBulk(address,uint256,uint256)",
	"98650275": "renounceMinter()",
	"715018a6": "renounceOwnership()",
	"3f4c4e3d": "requestValueTransfer(uint256,address,bytes)",
	"42842e0e": "safeTransferFrom(address,address,uint256)",
	"b88d4fde": "safeTransferFrom(address,address,uint256,bytes)",
	"a22cb465": "setApprovalForAll(address,bool)",
	"8dd14802": "setBridge(address)",
	"01ffc9a7": "supportsInterface(bytes4)",
	"95d89b41": "symbol()",
	"4f6ccce7": "tokenByIndex(uint256)",
	"2f745c59": "tokenOfOwnerByIndex(address,uint256)",
	"c87b56dd": "tokenURI(uint256)",
	"18160ddd": "totalSupply()",
	"23b872dd": "transferFrom(address,address,uint256)",
	"f2fde38b": "transferOwnership(address)",
}

// ServiceChainKIP17NFTBin is the compiled bytecode used for deploying new contracts.
var ServiceChainKIP17NFTBin = "0x60806040523480156200001157600080fd5b5060405160208062002b60833981018060405260208110156200003357600080fd5b5051604080518082018252601481527f53657276696365436861696e4b495031374e46540000000000000000000000006020828101919091528251808401909352600883527f4b4950313753434e0000000000000000000000000000000000000000000000008382015283929082908290620000d5907f01ffc9a700000000000000000000000000000000000000000000000000000000906200028c811b901c565b620000ed6380ac58cd60e01b6200028c60201b60201c565b6200010563780e9d6360e01b6200028c60201b60201c565b81516200011a906009906020850190620004da565b5080516200013090600a906020840190620004da565b5062000149635b5e139f60e01b6200028c60201b60201c565b50505050620001656342966c6860e01b6200028c60201b60201c565b62000176336200035b60201b60201c565b6200018e63fac27f4660e01b6200028c60201b60201c565b600d80546001600160a01b0319163317908190556040516001600160a01b0391909116906000907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a3620001f9816001600160a01b0316620003ad60201b62001ee11760201c565b6200026557604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f627269646765206973206e6f74206120636f6e74726163740000000000000000604482015290519081900360640190fd5b600e80546001600160a01b0319166001600160a01b0392909216919091179055506200057f565b7fffffffff0000000000000000000000000000000000000000000000000000000080821614156200031e57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f4b495031333a20696e76616c696420696e746572666163652069640000000000604482015290519081900360640190fd5b7fffffffff00000000000000000000000000000000000000000000000000000000166000908152602081905260409020805460ff19166001179055565b6200037681600c620003b360201b62001df31790919060201c565b6040516001600160a01b038216907f6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f690600090a250565b3b151590565b620003c582826200045760201b60201c565b156200043257604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f526f6c65733a206163636f756e7420616c72656164792068617320726f6c6500604482015290519081900360640190fd5b6001600160a01b0316600090815260209190915260409020805460ff19166001179055565b60006001600160a01b038216620004ba576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602281526020018062002b3e6022913960400191505060405180910390fd5b506001600160a01b03166000908152602091909152604090205460ff1690565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200051d57805160ff19168380011785556200054d565b828001600101855582156200054d579182015b828111156200054d57825182559160200191906001019062000530565b506200055b9291506200055f565b5090565b6200057c91905b808211156200055b576000815560010162000566565b90565b6125af806200058f6000396000f3fe608060405234801561001057600080fd5b50600436106101cf5760003560e01c8063715018a61161010457806398650275116100a2578063c87b56dd11610071578063c87b56dd14610722578063e78cea921461073f578063e985e9c514610747578063f2fde38b14610775576101cf565b80639865027514610600578063a22cb46514610608578063aa271e1a14610636578063b88d4fde1461065c576101cf565b80638dd14802116100de5780638dd14802146105a45780638f32d59b146105ca57806395d89b41146105d2578063983b2d56146105da576101cf565b8063715018a6146105625780637a9adac61461056a5780638da5cb5b1461059c576101cf565b80633f4c4e3d116101715780634f6ccce71161014b5780634f6ccce71461044757806350bb4e7f146104645780636352211e1461051f57806370a082311461053c576101cf565b80633f4c4e3d1461036f57806342842e0e146103f457806342966c681461042a576101cf565b8063095ea7b3116101ad578063095ea7b3146102c557806318160ddd146102f357806323b872dd1461030d5780632f745c5914610343576101cf565b806301ffc9a7146101d457806306fdde031461020f578063081812fc1461028c575b600080fd5b6101fb600480360360208110156101ea57600080fd5b50356001600160e01b03191661079b565b604080519115158252519081900360200190f35b6102176107ba565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610251578181015183820152602001610239565b50505050905090810190601f16801561027e5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6102a9600480360360208110156102a257600080fd5b5035610851565b604080516001600160a01b039092168252519081900360200190f35b6102f1600480360360408110156102db57600080fd5b506001600160a01b0381351690602001356108b6565b005b6102fb6109e3565b60408051918252519081900360200190f35b6102f16004803603606081101561032357600080fd5b506001600160a01b038135811691602081013590911690604001356109e9565b6102fb6004803603604081101561035957600080fd5b506001600160a01b038135169060200135610a41565b6102f16004803603606081101561038557600080fd5b8135916001600160a01b03602082013516918101906060810160408201356401000000008111156103b557600080fd5b8201836020820111156103c757600080fd5b803590602001918460018302840111640100000000831117156103e957600080fd5b509092509050610ac3565b6102f16004803603606081101561040a57600080fd5b506001600160a01b03813581169160208101359091169060400135610b90565b6102f16004803603602081101561044057600080fd5b5035610bab565b6102fb6004803603602081101561045d57600080fd5b5035610bff565b6101fb6004803603606081101561047a57600080fd5b6001600160a01b03823516916020810135918101906060810160408201356401000000008111156104aa57600080fd5b8201836020820111156104bc57600080fd5b803590602001918460018302840111640100000000831117156104de57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610c68945050505050565b6102a96004803603602081101561053557600080fd5b5035610ccf565b6102fb6004803603602081101561055257600080fd5b50356001600160a01b0316610d2c565b6102f1610d97565b6102f16004803603606081101561058057600080fd5b506001600160a01b038135169060208101359060400135610e2b565b6102a9610ebb565b6102f1600480360360208110156105ba57600080fd5b50356001600160a01b0316610eca565b6101fb610f36565b610217610f47565b6102f1600480360360208110156105f057600080fd5b50356001600160a01b0316610fa8565b6102f1610ff8565b6102f16004803603604081101561061e57600080fd5b506001600160a01b0381351690602001351515611003565b6101fb6004803603602081101561064c57600080fd5b50356001600160a01b03166110d2565b6102f16004803603608081101561067257600080fd5b6001600160a01b038235811692602081013590911691604082013591908101906080810160608201356401000000008111156106ad57600080fd5b8201836020820111156106bf57600080fd5b803590602001918460018302840111640100000000831117156106e157600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506110e5945050505050565b6102176004803603602081101561073857600080fd5b503561113a565b6102a9611222565b6101fb6004803603604081101561075d57600080fd5b506001600160a01b0381358116916020013516611231565b6102f16004803603602081101561078b57600080fd5b50356001600160a01b031661125f565b6001600160e01b03191660009081526020819052604090205460ff1690565b60098054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156108465780601f1061081b57610100808354040283529160200191610846565b820191906000526020600020905b81548152906001019060200180831161082957829003601f168201915b505050505090505b90565b600061085c826112b2565b61089a57604051600160e51b62461bcd02815260040180806020018281038252602b8152602001806124f7602b913960400191505060405180910390fd5b506000908152600260205260409020546001600160a01b031690565b60006108c182610ccf565b9050806001600160a01b0316836001600160a01b0316141561092d5760408051600160e51b62461bcd02815260206004820181905260248201527f4b495031373a20617070726f76616c20746f2063757272656e74206f776e6572604482015290519081900360640190fd5b336001600160a01b038216148061094957506109498133611231565b61098757604051600160e51b62461bcd0281526004018080602001828103825260378152602001806125226037913960400191505060405180910390fd5b60008281526002602052604080822080546001600160a01b0319166001600160a01b0387811691821790925591518593918516917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591a4505050565b60075490565b6109f333826112cf565b610a3157604051600160e51b62461bcd0281526004018080602001828103825260308152602001806123866030913960400191505060405180910390fd5b610a3c838383611376565b505050565b6000610a4c83610d2c565b8210610a8c57604051600160e51b62461bcd02815260040180806020018281038252602a8152602001806122b8602a913960400191505060405180910390fd5b6001600160a01b0383166000908152600560205260409020805483908110610ab057fe5b9060005260206000200154905092915050565b600e54610adb9033906001600160a01b0316866109e9565b600e54604051600160e01b6315c281770281523360048201818152602483018890526001600160a01b038781166044850152608060648501908152608485018790529416936315c28177938992899289928992919060a401848480828437600081840152601f19601f8201169050808301925050509650505050505050600060405180830381600087803b158015610b7257600080fd5b505af1158015610b86573d6000803e3d6000fd5b5050505050505050565b610a3c838383604051806020016040528060008152506110e5565b610bb533826112cf565b610bf357604051600160e51b62461bcd02815260040180806020018281038252602f815260200180612428602f913960400191505060405180910390fd5b610bfc81611395565b50565b6000610c096109e3565b8210610c4957604051600160e51b62461bcd02815260040180806020018281038252602b8152602001806124a8602b913960400191505060405180910390fd5b60078281548110610c5657fe5b90600052602060002001549050919050565b6000610c73336110d2565b610cb157604051600160e51b62461bcd0281526004018080602001828103825260308152602001806122e26030913960400191505060405180910390fd5b610cbb84846113a7565b610cc583836113c8565b5060019392505050565b6000818152600160205260408120546001600160a01b031680610d2657604051600160e51b62461bcd0281526004018080602001828103825260288152602001806123336028913960400191505060405180910390fd5b92915050565b60006001600160a01b038216610d7657604051600160e51b62461bcd0281526004018080602001828103825260298152602001806124576029913960400191505060405180910390fd5b6001600160a01b0382166000908152600360205260409020610d269061142e565b610d9f610f36565b610de15760408051600160e51b62461bcd02815260206004820181905260248201526000805160206123b6833981519152604482015290519081900360640190fd5b600d546040516000916001600160a01b0316907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600d80546001600160a01b0319169055565b610e33610f36565b610e755760408051600160e51b62461bcd02815260206004820181905260248201526000805160206123b6833981519152604482015290519081900360640190fd5b815b81811015610eb557610eac8482604051806040016040528060078152602001600160c81b667465737455524902815250610c68565b50600101610e77565b50505050565b600d546001600160a01b031690565b610ed2610f36565b610f145760408051600160e51b62461bcd02815260206004820181905260248201526000805160206123b6833981519152604482015290519081900360640190fd5b600e80546001600160a01b0319166001600160a01b0392909216919091179055565b600d546001600160a01b0316331490565b600a8054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156108465780601f1061081b57610100808354040283529160200191610846565b610fb1336110d2565b610fef57604051600160e51b62461bcd0281526004018080602001828103825260308152602001806122e26030913960400191505060405180910390fd5b610bfc81611432565b6110013361147a565b565b6001600160a01b0382163314156110645760408051600160e51b62461bcd02815260206004820152601860248201527f4b495031373a20617070726f766520746f2063616c6c65720000000000000000604482015290519081900360640190fd5b3360008181526004602090815260408083206001600160a01b03871680855290835292819020805460ff1916861515908117909155815190815290519293927f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31929181900390910190a35050565b6000610d26600c8363ffffffff6114c216565b6110f08484846109e9565b6110fc8484848461152c565b610eb557604051600160e51b62461bcd0281526004018080602001828103825260308152602001806123f86030913960400191505060405180910390fd5b6060611145826112b2565b61118357604051600160e51b62461bcd02815260040180806020018281038252602e815260200180612241602e913960400191505060405180910390fd5b6000828152600b602090815260409182902080548351601f6002600019610100600186161502019093169290920491820184900484028101840190945280845290918301828280156112165780601f106111eb57610100808354040283529160200191611216565b820191906000526020600020905b8154815290600101906020018083116111f957829003601f168201915b50505050509050919050565b600e546001600160a01b031681565b6001600160a01b03918216600090815260046020908152604080832093909416825291909152205460ff1690565b611267610f36565b6112a95760408051600160e51b62461bcd02815260206004820181905260248201526000805160206123b6833981519152604482015290519081900360640190fd5b610bfc8161190f565b6000908152600160205260409020546001600160a01b0316151590565b60006112da826112b2565b61131857604051600160e51b62461bcd02815260040180806020018281038252602b815260200180612559602b913960400191505060405180910390fd5b600061132383610ccf565b9050806001600160a01b0316846001600160a01b0316148061135e5750836001600160a01b031661135384610851565b6001600160a01b0316145b8061136e575061136e8185611231565b949350505050565b6113818383836119b3565b61138b8382611afd565b610a3c8282611bf2565b610bfc6113a182610ccf565b82611c30565b6113b18282611c78565b6113bb8282611bf2565b6113c481611daf565b5050565b6113d1826112b2565b61140f57604051600160e51b62461bcd02815260040180806020018281038252602b81526020018061235b602b913960400191505060405180910390fd5b6000828152600b602090815260409091208251610a3c92840190612144565b5490565b611443600c8263ffffffff611df316565b6040516001600160a01b038216907f6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f690600090a250565b61148b600c8263ffffffff611e7716565b6040516001600160a01b038216907fe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb6669290600090a250565b60006001600160a01b03821661150c57604051600160e51b62461bcd0281526004018080602001828103825260228152602001806123d66022913960400191505060405180910390fd5b506001600160a01b03166000908152602091909152604090205460ff1690565b6000806060611543866001600160a01b0316611ee1565b6115525760019250505061136e565b856001600160a01b031663150b7a0260e01b3389888860405160240180856001600160a01b03166001600160a01b03168152602001846001600160a01b03166001600160a01b0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b838110156115de5781810151838201526020016115c6565b50505050905090810190601f16801561160b5780820380516001836020036101000a031916815260200191505b5060408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909a16999099178952518151919890975087965094509250829150849050835b602083106116735780518252601f199092019160209182019101611654565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d80600081146116d5576040519150601f19603f3d011682016040523d82523d6000602084013e6116da565b606091505b50805191935091501580159061171a57508051600160e11b630a85bd0102906020808401919081101561170c57600080fd5b50516001600160e01b031916145b1561172a5760019250505061136e565b856001600160a01b0316636745782b60e01b3389888860405160240180856001600160a01b03166001600160a01b03168152602001846001600160a01b03166001600160a01b0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b838110156117b657818101518382015260200161179e565b50505050905090810190601f1680156117e35780820380516001836020036101000a031916815260200191505b5060408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909a16999099178952518151919890975087965094509250829150849050835b6020831061184b5780518252601f19909201916020918201910161182c565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d80600081146118ad576040519150601f19603f3d011682016040523d82523d6000602084013e6118b2565b606091505b5080519193509150158015906118f257508051600160e01b636745782b0290602080840191908110156118e457600080fd5b50516001600160e01b031916145b156119025760019250505061136e565b5060009695505050505050565b6001600160a01b03811661195757604051600160e51b62461bcd02815260040180806020018281038252602681526020018061226f6026913960400191505060405180910390fd5b600d546040516001600160a01b038084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3600d80546001600160a01b0319166001600160a01b0392909216919091179055565b826001600160a01b03166119c682610ccf565b6001600160a01b031614611a0e57604051600160e51b62461bcd0281526004018080602001828103825260288152602001806124806028913960400191505060405180910390fd5b6001600160a01b038216611a5657604051600160e51b62461bcd0281526004018080602001828103825260238152602001806122956023913960400191505060405180910390fd5b611a5f81611ee7565b6001600160a01b0383166000908152600360205260409020611a8090611f22565b6001600160a01b0382166000908152600360205260409020611aa190611f39565b60008181526001602052604080822080546001600160a01b0319166001600160a01b0386811691821790925591518493918716917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b6001600160a01b038216600090815260056020526040812054611b2790600163ffffffff611f4216565b600083815260066020526040902054909150808214611bc2576001600160a01b0384166000908152600560205260408120805484908110611b6457fe5b906000526020600020015490508060056000876001600160a01b03166001600160a01b031681526020019081526020016000208381548110611ba257fe5b600091825260208083209091019290925591825260069052604090208190555b6001600160a01b0384166000908152600560205260409020805490611beb9060001983016121c2565b5050505050565b6001600160a01b0390911660009081526005602081815260408084208054868652600684529185208290559282526001810183559183529091200155565b611c3a8282611fa2565b6000818152600b602052604090205460026000196101006001841615020190911604156113c4576000818152600b602052604081206113c4916121e6565b6001600160a01b038216611cd65760408051600160e51b62461bcd02815260206004820152601f60248201527f4b495031373a206d696e7420746f20746865207a65726f206164647265737300604482015290519081900360640190fd5b611cdf816112b2565b15611d345760408051600160e51b62461bcd02815260206004820152601b60248201527f4b495031373a20746f6b656e20616c7265616479206d696e7465640000000000604482015290519081900360640190fd5b600081815260016020908152604080832080546001600160a01b0319166001600160a01b038716908117909155835260039091529020611d7390611f39565b60405181906001600160a01b038416906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908290a45050565b600780546000838152600860205260408120829055600182018355919091527fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c6880155565b611dfd82826114c2565b15611e525760408051600160e51b62461bcd02815260206004820152601f60248201527f526f6c65733a206163636f756e7420616c72656164792068617320726f6c6500604482015290519081900360640190fd5b6001600160a01b0316600090815260209190915260409020805460ff19166001179055565b611e8182826114c2565b611ebf57604051600160e51b62461bcd0281526004018080602001828103825260218152602001806123126021913960400191505060405180910390fd5b6001600160a01b0316600090815260209190915260409020805460ff19169055565b3b151590565b6000818152600260205260409020546001600160a01b031615610bfc57600090815260026020526040902080546001600160a01b0319169055565b8054611f3590600163ffffffff611f4216565b9055565b80546001019055565b600082821115611f9c5760408051600160e51b62461bcd02815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b611fac8282611fce565b611fb68282611afd565b6000818152600660205260408120556113c4816120a8565b816001600160a01b0316611fe182610ccf565b6001600160a01b03161461202957604051600160e51b62461bcd0281526004018080602001828103825260248152602001806124d36024913960400191505060405180910390fd5b61203281611ee7565b6001600160a01b038216600090815260036020526040902061205390611f22565b60008181526001602052604080822080546001600160a01b0319169055518291906001600160a01b038516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908390a45050565b6007546000906120bf90600163ffffffff611f4216565b600083815260086020526040812054600780549394509092849081106120e157fe5b9060005260206000200154905080600783815481106120fc57fe5b6000918252602080832090910192909255828152600890915260409020829055600780549061212f9060001983016121c2565b50505060009182525060086020526040812055565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061218557805160ff19168380011785556121b2565b828001600101855582156121b2579182015b828111156121b2578251825591602001919060010190612197565b506121be929150612226565b5090565b815481835581811115610a3c57600083815260209020610a3c918101908301612226565b50805460018160011615610100020316600290046000825580601f1061220c5750610bfc565b601f016020900490600052602060002090810190610bfc91905b61084e91905b808211156121be576000815560010161222c56fe4b495031374d657461646174613a2055524920717565727920666f72206e6f6e6578697374656e7420746f6b656e4f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573734b495031373a207472616e7366657220746f20746865207a65726f20616464726573734b49503137456e756d657261626c653a206f776e657220696e646578206f7574206f6620626f756e64734d696e746572526f6c653a2063616c6c657220646f6573206e6f74206861766520746865204d696e74657220726f6c65526f6c65733a206163636f756e7420646f6573206e6f74206861766520726f6c654b495031373a206f776e657220717565727920666f72206e6f6e6578697374656e7420746f6b656e4b495031374d657461646174613a2055524920736574206f66206e6f6e6578697374656e7420746f6b656e4b495031373a207472616e736665722063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f7665644f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572526f6c65733a206163636f756e7420697320746865207a65726f20616464726573734b495031373a207472616e7366657220746f206e6f6e204b49503137526563656976657220696d706c656d656e7465724b495031374275726e61626c653a2063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f7665644b495031373a2062616c616e636520717565727920666f7220746865207a65726f20616464726573734b495031373a207472616e73666572206f6620746f6b656e2074686174206973206e6f74206f776e4b49503137456e756d657261626c653a20676c6f62616c20696e646578206f7574206f6620626f756e64734b495031373a206275726e206f6620746f6b656e2074686174206973206e6f74206f776e4b495031373a20617070726f76656420717565727920666f72206e6f6e6578697374656e7420746f6b656e4b495031373a20617070726f76652063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f76656420666f7220616c6c4b495031373a206f70657261746f7220717565727920666f72206e6f6e6578697374656e7420746f6b656ea165627a7a723058207f8f319e1745f9ad2fbd7e2330197c753ae01d8e90d01afb71a033060e7465710029526f6c65733a206163636f756e7420697320746865207a65726f2061646472657373"

// DeployServiceChainKIP17NFT deploys a new Klaytn contract, binding an instance of ServiceChainKIP17NFT to it.
func DeployServiceChainKIP17NFT(auth *bind.TransactOpts, backend bind.ContractBackend, _bridge common.Address) (common.Address, *types.Transaction, *ServiceChainKIP17NFT, error) {
	parsed, err := abi.JSON(strings.NewReader(ServiceChainKIP17NFTABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ServiceChainKIP17NFTBin), backend, _bridge)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ServiceChainKIP17NFT{ServiceChainKIP17NFTCaller: ServiceChainKIP17NFTCaller{contract: contract}, ServiceChainKIP17NFTTransactor: ServiceChainKIP17NFTTransactor{contract: contract}, ServiceChainKIP17NFTFilterer: ServiceChainKIP17NFTFilterer{contract: contract}}, nil
}

// ServiceChainKIP17NFT is an auto generated Go binding around a Klaytn contract.
type ServiceChainKIP17NFT struct {
	ServiceChainKIP17NFTCaller     // Read-only binding to the contract
	ServiceChainKIP17NFTTransactor // Write-only binding to the contract
	ServiceChainKIP17NFTFilterer   // Log filterer for contract events
}

// ServiceChainKIP17NFTCaller is an auto generated read-only Go binding around a Klaytn contract.
type ServiceChainKIP17NFTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ServiceChainKIP17NFTTransactor is an auto generated write-only Go binding around a Klaytn contract.
type ServiceChainKIP17NFTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ServiceChainKIP17NFTFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type ServiceChainKIP17NFTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ServiceChainKIP17NFTSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type ServiceChainKIP17NFTSession struct {
	Contract     *ServiceChainKIP17NFT // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ServiceChainKIP17NFTCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type ServiceChainKIP17NFTCallerSession struct {
	Contract *ServiceChainKIP17NFTCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// ServiceChainKIP17NFTTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type ServiceChainKIP17NFTTransactorSession struct {
	Contract     *ServiceChainKIP17NFTTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// ServiceChainKIP17NFTRaw is an auto generated low-level Go binding around a Klaytn contract.
type ServiceChainKIP17NFTRaw struct {
	Contract *ServiceChainKIP17NFT // Generic contract binding to access the raw methods on
}

// ServiceChainKIP17NFTCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type ServiceChainKIP17NFTCallerRaw struct {
	Contract *ServiceChainKIP17NFTCaller // Generic read-only contract binding to access the raw methods on
}

// ServiceChainKIP17NFTTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type ServiceChainKIP17NFTTransactorRaw struct {
	Contract *ServiceChainKIP17NFTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewServiceChainKIP17NFT creates a new instance of ServiceChainKIP17NFT, bound to a specific deployed contract.
func NewServiceChainKIP17NFT(address common.Address, backend bind.ContractBackend) (*ServiceChainKIP17NFT, error) {
	contract, err := bindServiceChainKIP17NFT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ServiceChainKIP17NFT{ServiceChainKIP17NFTCaller: ServiceChainKIP17NFTCaller{contract: contract}, ServiceChainKIP17NFTTransactor: ServiceChainKIP17NFTTransactor{contract: contract}, ServiceChainKIP17NFTFilterer: ServiceChainKIP17NFTFilterer{contract: contract}}, nil
}

// NewServiceChainKIP17NFTCaller creates a new read-only instance of ServiceChainKIP17NFT, bound to a specific deployed contract.
func NewServiceChainKIP17NFTCaller(address common.Address, caller bind.ContractCaller) (*ServiceChainKIP17NFTCaller, error) {
	contract, err := bindServiceChainKIP17NFT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ServiceChainKIP17NFTCaller{contract: contract}, nil
}

// NewServiceChainKIP17NFTTransactor creates a new write-only instance of ServiceChainKIP17NFT, bound to a specific deployed contract.
func NewServiceChainKIP17NFTTransactor(address common.Address, transactor bind.ContractTransactor) (*ServiceChainKIP17NFTTransactor, error) {
	contract, err := bindServiceChainKIP17NFT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ServiceChainKIP17NFTTransactor{contract: contract}, nil
}

// NewServiceChainKIP17NFTFilterer creates a new log filterer instance of ServiceChainKIP17NFT, bound to a specific deployed contract.
func NewServiceChainKIP17NFTFilterer(address common.Address, filterer bind.ContractFilterer) (*ServiceChainKIP17NFTFilterer, error) {
	contract, err := bindServiceChainKIP17NFT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ServiceChainKIP17NFTFilterer{contract: contract}, nil
}

// bindServiceChainKIP17NFT binds a generic wrapper to an already deployed contract.
func bindServiceChainKIP17NFT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ServiceChainKIP17NFTABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ServiceChainKIP17NFT.Contract.ServiceChainKIP17NFTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ServiceChainKIP17NFT.Contract.ServiceChainKIP17NFTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ServiceChainKIP17NFT.Contract.ServiceChainKIP17NFTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ServiceChainKIP17NFT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ServiceChainKIP17NFT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ServiceChainKIP17NFT.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ServiceChainKIP17NFT.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ServiceChainKIP17NFT.Contract.BalanceOf(&_ServiceChainKIP17NFT.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ServiceChainKIP17NFT.Contract.BalanceOf(&_ServiceChainKIP17NFT.CallOpts, owner)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ServiceChainKIP17NFT.contract.Call(opts, out, "bridge")
	return *ret0, err
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTSession) Bridge() (common.Address, error) {
	return _ServiceChainKIP17NFT.Contract.Bridge(&_ServiceChainKIP17NFT.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTCallerSession) Bridge() (common.Address, error) {
	return _ServiceChainKIP17NFT.Contract.Bridge(&_ServiceChainKIP17NFT.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ServiceChainKIP17NFT.contract.Call(opts, out, "getApproved", tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ServiceChainKIP17NFT.Contract.GetApproved(&_ServiceChainKIP17NFT.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ServiceChainKIP17NFT.Contract.GetApproved(&_ServiceChainKIP17NFT.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ServiceChainKIP17NFT.contract.Call(opts, out, "isApprovedForAll", owner, operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ServiceChainKIP17NFT.Contract.IsApprovedForAll(&_ServiceChainKIP17NFT.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ServiceChainKIP17NFT.Contract.IsApprovedForAll(&_ServiceChainKIP17NFT.CallOpts, owner, operator)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) view returns(bool)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTCaller) IsMinter(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ServiceChainKIP17NFT.contract.Call(opts, out, "isMinter", account)
	return *ret0, err
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) view returns(bool)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTSession) IsMinter(account common.Address) (bool, error) {
	return _ServiceChainKIP17NFT.Contract.IsMinter(&_ServiceChainKIP17NFT.CallOpts, account)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) view returns(bool)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTCallerSession) IsMinter(account common.Address) (bool, error) {
	return _ServiceChainKIP17NFT.Contract.IsMinter(&_ServiceChainKIP17NFT.CallOpts, account)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ServiceChainKIP17NFT.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTSession) IsOwner() (bool, error) {
	return _ServiceChainKIP17NFT.Contract.IsOwner(&_ServiceChainKIP17NFT.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTCallerSession) IsOwner() (bool, error) {
	return _ServiceChainKIP17NFT.Contract.IsOwner(&_ServiceChainKIP17NFT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ServiceChainKIP17NFT.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTSession) Name() (string, error) {
	return _ServiceChainKIP17NFT.Contract.Name(&_ServiceChainKIP17NFT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTCallerSession) Name() (string, error) {
	return _ServiceChainKIP17NFT.Contract.Name(&_ServiceChainKIP17NFT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ServiceChainKIP17NFT.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTSession) Owner() (common.Address, error) {
	return _ServiceChainKIP17NFT.Contract.Owner(&_ServiceChainKIP17NFT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTCallerSession) Owner() (common.Address, error) {
	return _ServiceChainKIP17NFT.Contract.Owner(&_ServiceChainKIP17NFT.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ServiceChainKIP17NFT.contract.Call(opts, out, "ownerOf", tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ServiceChainKIP17NFT.Contract.OwnerOf(&_ServiceChainKIP17NFT.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ServiceChainKIP17NFT.Contract.OwnerOf(&_ServiceChainKIP17NFT.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ServiceChainKIP17NFT.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ServiceChainKIP17NFT.Contract.SupportsInterface(&_ServiceChainKIP17NFT.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ServiceChainKIP17NFT.Contract.SupportsInterface(&_ServiceChainKIP17NFT.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ServiceChainKIP17NFT.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTSession) Symbol() (string, error) {
	return _ServiceChainKIP17NFT.Contract.Symbol(&_ServiceChainKIP17NFT.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTCallerSession) Symbol() (string, error) {
	return _ServiceChainKIP17NFT.Contract.Symbol(&_ServiceChainKIP17NFT.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ServiceChainKIP17NFT.contract.Call(opts, out, "tokenByIndex", index)
	return *ret0, err
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _ServiceChainKIP17NFT.Contract.TokenByIndex(&_ServiceChainKIP17NFT.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _ServiceChainKIP17NFT.Contract.TokenByIndex(&_ServiceChainKIP17NFT.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ServiceChainKIP17NFT.contract.Call(opts, out, "tokenOfOwnerByIndex", owner, index)
	return *ret0, err
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _ServiceChainKIP17NFT.Contract.TokenOfOwnerByIndex(&_ServiceChainKIP17NFT.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _ServiceChainKIP17NFT.Contract.TokenOfOwnerByIndex(&_ServiceChainKIP17NFT.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ServiceChainKIP17NFT.contract.Call(opts, out, "tokenURI", tokenId)
	return *ret0, err
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTSession) TokenURI(tokenId *big.Int) (string, error) {
	return _ServiceChainKIP17NFT.Contract.TokenURI(&_ServiceChainKIP17NFT.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _ServiceChainKIP17NFT.Contract.TokenURI(&_ServiceChainKIP17NFT.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ServiceChainKIP17NFT.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTSession) TotalSupply() (*big.Int, error) {
	return _ServiceChainKIP17NFT.Contract.TotalSupply(&_ServiceChainKIP17NFT.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTCallerSession) TotalSupply() (*big.Int, error) {
	return _ServiceChainKIP17NFT.Contract.TotalSupply(&_ServiceChainKIP17NFT.CallOpts)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTTransactor) AddMinter(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _ServiceChainKIP17NFT.contract.Transact(opts, "addMinter", account)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTSession) AddMinter(account common.Address) (*types.Transaction, error) {
	return _ServiceChainKIP17NFT.Contract.AddMinter(&_ServiceChainKIP17NFT.TransactOpts, account)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTTransactorSession) AddMinter(account common.Address) (*types.Transaction, error) {
	return _ServiceChainKIP17NFT.Contract.AddMinter(&_ServiceChainKIP17NFT.TransactOpts, account)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ServiceChainKIP17NFT.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ServiceChainKIP17NFT.Contract.Approve(&_ServiceChainKIP17NFT.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ServiceChainKIP17NFT.Contract.Approve(&_ServiceChainKIP17NFT.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _ServiceChainKIP17NFT.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _ServiceChainKIP17NFT.Contract.Burn(&_ServiceChainKIP17NFT.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _ServiceChainKIP17NFT.Contract.Burn(&_ServiceChainKIP17NFT.TransactOpts, tokenId)
}

// MintWithTokenURI is a paid mutator transaction binding the contract method 0x50bb4e7f.
//
// Solidity: function mintWithTokenURI(address to, uint256 tokenId, string tokenURI) returns(bool)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTTransactor) MintWithTokenURI(opts *bind.TransactOpts, to common.Address, tokenId *big.Int, tokenURI string) (*types.Transaction, error) {
	return _ServiceChainKIP17NFT.contract.Transact(opts, "mintWithTokenURI", to, tokenId, tokenURI)
}

// MintWithTokenURI is a paid mutator transaction binding the contract method 0x50bb4e7f.
//
// Solidity: function mintWithTokenURI(address to, uint256 tokenId, string tokenURI) returns(bool)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTSession) MintWithTokenURI(to common.Address, tokenId *big.Int, tokenURI string) (*types.Transaction, error) {
	return _ServiceChainKIP17NFT.Contract.MintWithTokenURI(&_ServiceChainKIP17NFT.TransactOpts, to, tokenId, tokenURI)
}

// MintWithTokenURI is a paid mutator transaction binding the contract method 0x50bb4e7f.
//
// Solidity: function mintWithTokenURI(address to, uint256 tokenId, string tokenURI) returns(bool)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTTransactorSession) MintWithTokenURI(to common.Address, tokenId *big.Int, tokenURI string) (*types.Transaction, error) {
	return _ServiceChainKIP17NFT.Contract.MintWithTokenURI(&_ServiceChainKIP17NFT.TransactOpts, to, tokenId, tokenURI)
}

// RegisterBulk is a paid mutator transaction binding the contract method 0x7a9adac6.
//
// Solidity: function registerBulk(address _user, uint256 _startID, uint256 _endID) returns()
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTTransactor) RegisterBulk(opts *bind.TransactOpts, _user common.Address, _startID *big.Int, _endID *big.Int) (*types.Transaction, error) {
	return _ServiceChainKIP17NFT.contract.Transact(opts, "registerBulk", _user, _startID, _endID)
}

// RegisterBulk is a paid mutator transaction binding the contract method 0x7a9adac6.
//
// Solidity: function registerBulk(address _user, uint256 _startID, uint256 _endID) returns()
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTSession) RegisterBulk(_user common.Address, _startID *big.Int, _endID *big.Int) (*types.Transaction, error) {
	return _ServiceChainKIP17NFT.Contract.RegisterBulk(&_ServiceChainKIP17NFT.TransactOpts, _user, _startID, _endID)
}

// RegisterBulk is a paid mutator transaction binding the contract method 0x7a9adac6.
//
// Solidity: function registerBulk(address _user, uint256 _startID, uint256 _endID) returns()
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTTransactorSession) RegisterBulk(_user common.Address, _startID *big.Int, _endID *big.Int) (*types.Transaction, error) {
	return _ServiceChainKIP17NFT.Contract.RegisterBulk(&_ServiceChainKIP17NFT.TransactOpts, _user, _startID, _endID)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTTransactor) RenounceMinter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ServiceChainKIP17NFT.contract.Transact(opts, "renounceMinter")
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTSession) RenounceMinter() (*types.Transaction, error) {
	return _ServiceChainKIP17NFT.Contract.RenounceMinter(&_ServiceChainKIP17NFT.TransactOpts)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTTransactorSession) RenounceMinter() (*types.Transaction, error) {
	return _ServiceChainKIP17NFT.Contract.RenounceMinter(&_ServiceChainKIP17NFT.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ServiceChainKIP17NFT.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTSession) RenounceOwnership() (*types.Transaction, error) {
	return _ServiceChainKIP17NFT.Contract.RenounceOwnership(&_ServiceChainKIP17NFT.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ServiceChainKIP17NFT.Contract.RenounceOwnership(&_ServiceChainKIP17NFT.TransactOpts)
}

// RequestValueTransfer is a paid mutator transaction binding the contract method 0x3f4c4e3d.
//
// Solidity: function requestValueTransfer(uint256 _uid, address _to, bytes _extraData) returns()
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTTransactor) RequestValueTransfer(opts *bind.TransactOpts, _uid *big.Int, _to common.Address, _extraData []byte) (*types.Transaction, error) {
	return _ServiceChainKIP17NFT.contract.Transact(opts, "requestValueTransfer", _uid, _to, _extraData)
}

// RequestValueTransfer is a paid mutator transaction binding the contract method 0x3f4c4e3d.
//
// Solidity: function requestValueTransfer(uint256 _uid, address _to, bytes _extraData) returns()
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTSession) RequestValueTransfer(_uid *big.Int, _to common.Address, _extraData []byte) (*types.Transaction, error) {
	return _ServiceChainKIP17NFT.Contract.RequestValueTransfer(&_ServiceChainKIP17NFT.TransactOpts, _uid, _to, _extraData)
}

// RequestValueTransfer is a paid mutator transaction binding the contract method 0x3f4c4e3d.
//
// Solidity: function requestValueTransfer(uint256 _uid, address _to, bytes _extraData) returns()
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTTransactorSession) RequestValueTransfer(_uid *big.Int, _to common.Address, _extraData []byte) (*types.Transaction, error) {
	return _ServiceChainKIP17NFT.Contract.RequestValueTransfer(&_ServiceChainKIP17NFT.TransactOpts, _uid, _to, _extraData)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ServiceChainKIP17NFT.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ServiceChainKIP17NFT.Contract.SafeTransferFrom(&_ServiceChainKIP17NFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ServiceChainKIP17NFT.Contract.SafeTransferFrom(&_ServiceChainKIP17NFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ServiceChainKIP17NFT.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ServiceChainKIP17NFT.Contract.SafeTransferFrom0(&_ServiceChainKIP17NFT.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ServiceChainKIP17NFT.Contract.SafeTransferFrom0(&_ServiceChainKIP17NFT.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTTransactor) SetApprovalForAll(opts *bind.TransactOpts, to common.Address, approved bool) (*types.Transaction, error) {
	return _ServiceChainKIP17NFT.contract.Transact(opts, "setApprovalForAll", to, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTSession) SetApprovalForAll(to common.Address, approved bool) (*types.Transaction, error) {
	return _ServiceChainKIP17NFT.Contract.SetApprovalForAll(&_ServiceChainKIP17NFT.TransactOpts, to, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTTransactorSession) SetApprovalForAll(to common.Address, approved bool) (*types.Transaction, error) {
	return _ServiceChainKIP17NFT.Contract.SetApprovalForAll(&_ServiceChainKIP17NFT.TransactOpts, to, approved)
}

// SetBridge is a paid mutator transaction binding the contract method 0x8dd14802.
//
// Solidity: function setBridge(address _bridge) returns()
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTTransactor) SetBridge(opts *bind.TransactOpts, _bridge common.Address) (*types.Transaction, error) {
	return _ServiceChainKIP17NFT.contract.Transact(opts, "setBridge", _bridge)
}

// SetBridge is a paid mutator transaction binding the contract method 0x8dd14802.
//
// Solidity: function setBridge(address _bridge) returns()
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTSession) SetBridge(_bridge common.Address) (*types.Transaction, error) {
	return _ServiceChainKIP17NFT.Contract.SetBridge(&_ServiceChainKIP17NFT.TransactOpts, _bridge)
}

// SetBridge is a paid mutator transaction binding the contract method 0x8dd14802.
//
// Solidity: function setBridge(address _bridge) returns()
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTTransactorSession) SetBridge(_bridge common.Address) (*types.Transaction, error) {
	return _ServiceChainKIP17NFT.Contract.SetBridge(&_ServiceChainKIP17NFT.TransactOpts, _bridge)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ServiceChainKIP17NFT.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ServiceChainKIP17NFT.Contract.TransferFrom(&_ServiceChainKIP17NFT.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ServiceChainKIP17NFT.Contract.TransferFrom(&_ServiceChainKIP17NFT.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ServiceChainKIP17NFT.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ServiceChainKIP17NFT.Contract.TransferOwnership(&_ServiceChainKIP17NFT.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ServiceChainKIP17NFT.Contract.TransferOwnership(&_ServiceChainKIP17NFT.TransactOpts, newOwner)
}

// ServiceChainKIP17NFTApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ServiceChainKIP17NFT contract.
type ServiceChainKIP17NFTApprovalIterator struct {
	Event *ServiceChainKIP17NFTApproval // Event containing the contract specifics and raw log

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
func (it *ServiceChainKIP17NFTApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ServiceChainKIP17NFTApproval)
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
		it.Event = new(ServiceChainKIP17NFTApproval)
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
func (it *ServiceChainKIP17NFTApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ServiceChainKIP17NFTApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ServiceChainKIP17NFTApproval represents a Approval event raised by the ServiceChainKIP17NFT contract.
type ServiceChainKIP17NFTApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ServiceChainKIP17NFTApprovalIterator, error) {

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

	logs, sub, err := _ServiceChainKIP17NFT.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ServiceChainKIP17NFTApprovalIterator{contract: _ServiceChainKIP17NFT.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ServiceChainKIP17NFTApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _ServiceChainKIP17NFT.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ServiceChainKIP17NFTApproval)
				if err := _ServiceChainKIP17NFT.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTFilterer) ParseApproval(log types.Log) (*ServiceChainKIP17NFTApproval, error) {
	event := new(ServiceChainKIP17NFTApproval)
	if err := _ServiceChainKIP17NFT.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ServiceChainKIP17NFTApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ServiceChainKIP17NFT contract.
type ServiceChainKIP17NFTApprovalForAllIterator struct {
	Event *ServiceChainKIP17NFTApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ServiceChainKIP17NFTApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ServiceChainKIP17NFTApprovalForAll)
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
		it.Event = new(ServiceChainKIP17NFTApprovalForAll)
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
func (it *ServiceChainKIP17NFTApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ServiceChainKIP17NFTApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ServiceChainKIP17NFTApprovalForAll represents a ApprovalForAll event raised by the ServiceChainKIP17NFT contract.
type ServiceChainKIP17NFTApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ServiceChainKIP17NFTApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ServiceChainKIP17NFT.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ServiceChainKIP17NFTApprovalForAllIterator{contract: _ServiceChainKIP17NFT.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ServiceChainKIP17NFTApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ServiceChainKIP17NFT.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ServiceChainKIP17NFTApprovalForAll)
				if err := _ServiceChainKIP17NFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTFilterer) ParseApprovalForAll(log types.Log) (*ServiceChainKIP17NFTApprovalForAll, error) {
	event := new(ServiceChainKIP17NFTApprovalForAll)
	if err := _ServiceChainKIP17NFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ServiceChainKIP17NFTMinterAddedIterator is returned from FilterMinterAdded and is used to iterate over the raw logs and unpacked data for MinterAdded events raised by the ServiceChainKIP17NFT contract.
type ServiceChainKIP17NFTMinterAddedIterator struct {
	Event *ServiceChainKIP17NFTMinterAdded // Event containing the contract specifics and raw log

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
func (it *ServiceChainKIP17NFTMinterAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ServiceChainKIP17NFTMinterAdded)
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
		it.Event = new(ServiceChainKIP17NFTMinterAdded)
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
func (it *ServiceChainKIP17NFTMinterAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ServiceChainKIP17NFTMinterAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ServiceChainKIP17NFTMinterAdded represents a MinterAdded event raised by the ServiceChainKIP17NFT contract.
type ServiceChainKIP17NFTMinterAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinterAdded is a free log retrieval operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed account)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTFilterer) FilterMinterAdded(opts *bind.FilterOpts, account []common.Address) (*ServiceChainKIP17NFTMinterAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ServiceChainKIP17NFT.contract.FilterLogs(opts, "MinterAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &ServiceChainKIP17NFTMinterAddedIterator{contract: _ServiceChainKIP17NFT.contract, event: "MinterAdded", logs: logs, sub: sub}, nil
}

// WatchMinterAdded is a free log subscription operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed account)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTFilterer) WatchMinterAdded(opts *bind.WatchOpts, sink chan<- *ServiceChainKIP17NFTMinterAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ServiceChainKIP17NFT.contract.WatchLogs(opts, "MinterAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ServiceChainKIP17NFTMinterAdded)
				if err := _ServiceChainKIP17NFT.contract.UnpackLog(event, "MinterAdded", log); err != nil {
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
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTFilterer) ParseMinterAdded(log types.Log) (*ServiceChainKIP17NFTMinterAdded, error) {
	event := new(ServiceChainKIP17NFTMinterAdded)
	if err := _ServiceChainKIP17NFT.contract.UnpackLog(event, "MinterAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ServiceChainKIP17NFTMinterRemovedIterator is returned from FilterMinterRemoved and is used to iterate over the raw logs and unpacked data for MinterRemoved events raised by the ServiceChainKIP17NFT contract.
type ServiceChainKIP17NFTMinterRemovedIterator struct {
	Event *ServiceChainKIP17NFTMinterRemoved // Event containing the contract specifics and raw log

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
func (it *ServiceChainKIP17NFTMinterRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ServiceChainKIP17NFTMinterRemoved)
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
		it.Event = new(ServiceChainKIP17NFTMinterRemoved)
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
func (it *ServiceChainKIP17NFTMinterRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ServiceChainKIP17NFTMinterRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ServiceChainKIP17NFTMinterRemoved represents a MinterRemoved event raised by the ServiceChainKIP17NFT contract.
type ServiceChainKIP17NFTMinterRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinterRemoved is a free log retrieval operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed account)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTFilterer) FilterMinterRemoved(opts *bind.FilterOpts, account []common.Address) (*ServiceChainKIP17NFTMinterRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ServiceChainKIP17NFT.contract.FilterLogs(opts, "MinterRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &ServiceChainKIP17NFTMinterRemovedIterator{contract: _ServiceChainKIP17NFT.contract, event: "MinterRemoved", logs: logs, sub: sub}, nil
}

// WatchMinterRemoved is a free log subscription operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed account)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTFilterer) WatchMinterRemoved(opts *bind.WatchOpts, sink chan<- *ServiceChainKIP17NFTMinterRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ServiceChainKIP17NFT.contract.WatchLogs(opts, "MinterRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ServiceChainKIP17NFTMinterRemoved)
				if err := _ServiceChainKIP17NFT.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
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
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTFilterer) ParseMinterRemoved(log types.Log) (*ServiceChainKIP17NFTMinterRemoved, error) {
	event := new(ServiceChainKIP17NFTMinterRemoved)
	if err := _ServiceChainKIP17NFT.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ServiceChainKIP17NFTOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ServiceChainKIP17NFT contract.
type ServiceChainKIP17NFTOwnershipTransferredIterator struct {
	Event *ServiceChainKIP17NFTOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ServiceChainKIP17NFTOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ServiceChainKIP17NFTOwnershipTransferred)
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
		it.Event = new(ServiceChainKIP17NFTOwnershipTransferred)
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
func (it *ServiceChainKIP17NFTOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ServiceChainKIP17NFTOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ServiceChainKIP17NFTOwnershipTransferred represents a OwnershipTransferred event raised by the ServiceChainKIP17NFT contract.
type ServiceChainKIP17NFTOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ServiceChainKIP17NFTOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ServiceChainKIP17NFT.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ServiceChainKIP17NFTOwnershipTransferredIterator{contract: _ServiceChainKIP17NFT.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ServiceChainKIP17NFTOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ServiceChainKIP17NFT.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ServiceChainKIP17NFTOwnershipTransferred)
				if err := _ServiceChainKIP17NFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTFilterer) ParseOwnershipTransferred(log types.Log) (*ServiceChainKIP17NFTOwnershipTransferred, error) {
	event := new(ServiceChainKIP17NFTOwnershipTransferred)
	if err := _ServiceChainKIP17NFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ServiceChainKIP17NFTTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ServiceChainKIP17NFT contract.
type ServiceChainKIP17NFTTransferIterator struct {
	Event *ServiceChainKIP17NFTTransfer // Event containing the contract specifics and raw log

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
func (it *ServiceChainKIP17NFTTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ServiceChainKIP17NFTTransfer)
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
		it.Event = new(ServiceChainKIP17NFTTransfer)
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
func (it *ServiceChainKIP17NFTTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ServiceChainKIP17NFTTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ServiceChainKIP17NFTTransfer represents a Transfer event raised by the ServiceChainKIP17NFT contract.
type ServiceChainKIP17NFTTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ServiceChainKIP17NFTTransferIterator, error) {

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

	logs, sub, err := _ServiceChainKIP17NFT.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ServiceChainKIP17NFTTransferIterator{contract: _ServiceChainKIP17NFT.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ServiceChainKIP17NFTTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _ServiceChainKIP17NFT.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ServiceChainKIP17NFTTransfer)
				if err := _ServiceChainKIP17NFT.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ServiceChainKIP17NFT *ServiceChainKIP17NFTFilterer) ParseTransfer(log types.Log) (*ServiceChainKIP17NFTTransfer, error) {
	event := new(ServiceChainKIP17NFTTransfer)
	if err := _ServiceChainKIP17NFT.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
