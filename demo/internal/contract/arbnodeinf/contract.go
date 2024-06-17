// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package arbnodeinf

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ArbnodeinfMetaData contains all meta data concerning the Arbnodeinf contract.
var ArbnodeinfMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"blockL1Num\",\"inputs\":[{\"name\":\"l2BlockNum\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"l1BlockNum\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"constructOutboxProof\",\"inputs\":[{\"name\":\"size\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"leaf\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"send\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proof\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"estimateRetryableTicket\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"deposit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"l2CallValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"excessFeeRefundAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callValueRefundAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"findBatchContainingBlock\",\"inputs\":[{\"name\":\"blockNum\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"batch\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"gasEstimateComponents\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"contractCreation\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"gasEstimate\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasEstimateForL1\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"baseFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"l1BaseFeeEstimate\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"gasEstimateL1Component\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"contractCreation\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"gasEstimateForL1\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"baseFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"l1BaseFeeEstimate\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"getL1Confirmations\",\"inputs\":[{\"name\":\"blockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"confirmations\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"l2BlockRangeForL1\",\"inputs\":[{\"name\":\"blockNum\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"firstBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"lastBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"legacyLookupMessageBatchProof\",\"inputs\":[{\"name\":\"batchNum\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"index\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"proof\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"path\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"l2Sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"l1Dest\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"l2Block\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"l1Block\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"calldataForL1\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nitroGenesisBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"number\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"}]",
}

// ArbnodeinfABI is the input ABI used to generate the binding from.
// Deprecated: Use ArbnodeinfMetaData.ABI instead.
var ArbnodeinfABI = ArbnodeinfMetaData.ABI

// Arbnodeinf is an auto generated Go binding around an Ethereum contract.
type Arbnodeinf struct {
	ArbnodeinfCaller     // Read-only binding to the contract
	ArbnodeinfTransactor // Write-only binding to the contract
	ArbnodeinfFilterer   // Log filterer for contract events
}

// ArbnodeinfCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArbnodeinfCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbnodeinfTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArbnodeinfTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbnodeinfFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArbnodeinfFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbnodeinfSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArbnodeinfSession struct {
	Contract     *Arbnodeinf       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArbnodeinfCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArbnodeinfCallerSession struct {
	Contract *ArbnodeinfCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ArbnodeinfTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArbnodeinfTransactorSession struct {
	Contract     *ArbnodeinfTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ArbnodeinfRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArbnodeinfRaw struct {
	Contract *Arbnodeinf // Generic contract binding to access the raw methods on
}

// ArbnodeinfCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArbnodeinfCallerRaw struct {
	Contract *ArbnodeinfCaller // Generic read-only contract binding to access the raw methods on
}

// ArbnodeinfTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArbnodeinfTransactorRaw struct {
	Contract *ArbnodeinfTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArbnodeinf creates a new instance of Arbnodeinf, bound to a specific deployed contract.
func NewArbnodeinf(address common.Address, backend bind.ContractBackend) (*Arbnodeinf, error) {
	contract, err := bindArbnodeinf(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Arbnodeinf{ArbnodeinfCaller: ArbnodeinfCaller{contract: contract}, ArbnodeinfTransactor: ArbnodeinfTransactor{contract: contract}, ArbnodeinfFilterer: ArbnodeinfFilterer{contract: contract}}, nil
}

// NewArbnodeinfCaller creates a new read-only instance of Arbnodeinf, bound to a specific deployed contract.
func NewArbnodeinfCaller(address common.Address, caller bind.ContractCaller) (*ArbnodeinfCaller, error) {
	contract, err := bindArbnodeinf(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbnodeinfCaller{contract: contract}, nil
}

// NewArbnodeinfTransactor creates a new write-only instance of Arbnodeinf, bound to a specific deployed contract.
func NewArbnodeinfTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbnodeinfTransactor, error) {
	contract, err := bindArbnodeinf(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbnodeinfTransactor{contract: contract}, nil
}

// NewArbnodeinfFilterer creates a new log filterer instance of Arbnodeinf, bound to a specific deployed contract.
func NewArbnodeinfFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbnodeinfFilterer, error) {
	contract, err := bindArbnodeinf(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbnodeinfFilterer{contract: contract}, nil
}

// bindArbnodeinf binds a generic wrapper to an already deployed contract.
func bindArbnodeinf(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbnodeinfABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Arbnodeinf *ArbnodeinfRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Arbnodeinf.Contract.ArbnodeinfCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Arbnodeinf *ArbnodeinfRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Arbnodeinf.Contract.ArbnodeinfTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Arbnodeinf *ArbnodeinfRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Arbnodeinf.Contract.ArbnodeinfTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Arbnodeinf *ArbnodeinfCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Arbnodeinf.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Arbnodeinf *ArbnodeinfTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Arbnodeinf.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Arbnodeinf *ArbnodeinfTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Arbnodeinf.Contract.contract.Transact(opts, method, params...)
}

// BlockL1Num is a free data retrieval call binding the contract method 0x6f275ef2.
//
// Solidity: function blockL1Num(uint64 l2BlockNum) view returns(uint64 l1BlockNum)
func (_Arbnodeinf *ArbnodeinfCaller) BlockL1Num(opts *bind.CallOpts, l2BlockNum uint64) (uint64, error) {
	var out []interface{}
	err := _Arbnodeinf.contract.Call(opts, &out, "blockL1Num", l2BlockNum)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// BlockL1Num is a free data retrieval call binding the contract method 0x6f275ef2.
//
// Solidity: function blockL1Num(uint64 l2BlockNum) view returns(uint64 l1BlockNum)
func (_Arbnodeinf *ArbnodeinfSession) BlockL1Num(l2BlockNum uint64) (uint64, error) {
	return _Arbnodeinf.Contract.BlockL1Num(&_Arbnodeinf.CallOpts, l2BlockNum)
}

// BlockL1Num is a free data retrieval call binding the contract method 0x6f275ef2.
//
// Solidity: function blockL1Num(uint64 l2BlockNum) view returns(uint64 l1BlockNum)
func (_Arbnodeinf *ArbnodeinfCallerSession) BlockL1Num(l2BlockNum uint64) (uint64, error) {
	return _Arbnodeinf.Contract.BlockL1Num(&_Arbnodeinf.CallOpts, l2BlockNum)
}

// ConstructOutboxProof is a free data retrieval call binding the contract method 0x42696350.
//
// Solidity: function constructOutboxProof(uint64 size, uint64 leaf) view returns(bytes32 send, bytes32 root, bytes32[] proof)
func (_Arbnodeinf *ArbnodeinfCaller) ConstructOutboxProof(opts *bind.CallOpts, size uint64, leaf uint64) (struct {
	Send  [32]byte
	Root  [32]byte
	Proof [][32]byte
}, error) {
	var out []interface{}
	err := _Arbnodeinf.contract.Call(opts, &out, "constructOutboxProof", size, leaf)

	outstruct := new(struct {
		Send  [32]byte
		Root  [32]byte
		Proof [][32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Send = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Root = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Proof = *abi.ConvertType(out[2], new([][32]byte)).(*[][32]byte)

	return *outstruct, err

}

// ConstructOutboxProof is a free data retrieval call binding the contract method 0x42696350.
//
// Solidity: function constructOutboxProof(uint64 size, uint64 leaf) view returns(bytes32 send, bytes32 root, bytes32[] proof)
func (_Arbnodeinf *ArbnodeinfSession) ConstructOutboxProof(size uint64, leaf uint64) (struct {
	Send  [32]byte
	Root  [32]byte
	Proof [][32]byte
}, error) {
	return _Arbnodeinf.Contract.ConstructOutboxProof(&_Arbnodeinf.CallOpts, size, leaf)
}

// ConstructOutboxProof is a free data retrieval call binding the contract method 0x42696350.
//
// Solidity: function constructOutboxProof(uint64 size, uint64 leaf) view returns(bytes32 send, bytes32 root, bytes32[] proof)
func (_Arbnodeinf *ArbnodeinfCallerSession) ConstructOutboxProof(size uint64, leaf uint64) (struct {
	Send  [32]byte
	Root  [32]byte
	Proof [][32]byte
}, error) {
	return _Arbnodeinf.Contract.ConstructOutboxProof(&_Arbnodeinf.CallOpts, size, leaf)
}

// FindBatchContainingBlock is a free data retrieval call binding the contract method 0x81f1adaf.
//
// Solidity: function findBatchContainingBlock(uint64 blockNum) view returns(uint64 batch)
func (_Arbnodeinf *ArbnodeinfCaller) FindBatchContainingBlock(opts *bind.CallOpts, blockNum uint64) (uint64, error) {
	var out []interface{}
	err := _Arbnodeinf.contract.Call(opts, &out, "findBatchContainingBlock", blockNum)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// FindBatchContainingBlock is a free data retrieval call binding the contract method 0x81f1adaf.
//
// Solidity: function findBatchContainingBlock(uint64 blockNum) view returns(uint64 batch)
func (_Arbnodeinf *ArbnodeinfSession) FindBatchContainingBlock(blockNum uint64) (uint64, error) {
	return _Arbnodeinf.Contract.FindBatchContainingBlock(&_Arbnodeinf.CallOpts, blockNum)
}

// FindBatchContainingBlock is a free data retrieval call binding the contract method 0x81f1adaf.
//
// Solidity: function findBatchContainingBlock(uint64 blockNum) view returns(uint64 batch)
func (_Arbnodeinf *ArbnodeinfCallerSession) FindBatchContainingBlock(blockNum uint64) (uint64, error) {
	return _Arbnodeinf.Contract.FindBatchContainingBlock(&_Arbnodeinf.CallOpts, blockNum)
}

// GetL1Confirmations is a free data retrieval call binding the contract method 0xe5ca238c.
//
// Solidity: function getL1Confirmations(bytes32 blockHash) view returns(uint64 confirmations)
func (_Arbnodeinf *ArbnodeinfCaller) GetL1Confirmations(opts *bind.CallOpts, blockHash [32]byte) (uint64, error) {
	var out []interface{}
	err := _Arbnodeinf.contract.Call(opts, &out, "getL1Confirmations", blockHash)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetL1Confirmations is a free data retrieval call binding the contract method 0xe5ca238c.
//
// Solidity: function getL1Confirmations(bytes32 blockHash) view returns(uint64 confirmations)
func (_Arbnodeinf *ArbnodeinfSession) GetL1Confirmations(blockHash [32]byte) (uint64, error) {
	return _Arbnodeinf.Contract.GetL1Confirmations(&_Arbnodeinf.CallOpts, blockHash)
}

// GetL1Confirmations is a free data retrieval call binding the contract method 0xe5ca238c.
//
// Solidity: function getL1Confirmations(bytes32 blockHash) view returns(uint64 confirmations)
func (_Arbnodeinf *ArbnodeinfCallerSession) GetL1Confirmations(blockHash [32]byte) (uint64, error) {
	return _Arbnodeinf.Contract.GetL1Confirmations(&_Arbnodeinf.CallOpts, blockHash)
}

// L2BlockRangeForL1 is a free data retrieval call binding the contract method 0x48e7f811.
//
// Solidity: function l2BlockRangeForL1(uint64 blockNum) view returns(uint64 firstBlock, uint64 lastBlock)
func (_Arbnodeinf *ArbnodeinfCaller) L2BlockRangeForL1(opts *bind.CallOpts, blockNum uint64) (struct {
	FirstBlock uint64
	LastBlock  uint64
}, error) {
	var out []interface{}
	err := _Arbnodeinf.contract.Call(opts, &out, "l2BlockRangeForL1", blockNum)

	outstruct := new(struct {
		FirstBlock uint64
		LastBlock  uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FirstBlock = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.LastBlock = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// L2BlockRangeForL1 is a free data retrieval call binding the contract method 0x48e7f811.
//
// Solidity: function l2BlockRangeForL1(uint64 blockNum) view returns(uint64 firstBlock, uint64 lastBlock)
func (_Arbnodeinf *ArbnodeinfSession) L2BlockRangeForL1(blockNum uint64) (struct {
	FirstBlock uint64
	LastBlock  uint64
}, error) {
	return _Arbnodeinf.Contract.L2BlockRangeForL1(&_Arbnodeinf.CallOpts, blockNum)
}

// L2BlockRangeForL1 is a free data retrieval call binding the contract method 0x48e7f811.
//
// Solidity: function l2BlockRangeForL1(uint64 blockNum) view returns(uint64 firstBlock, uint64 lastBlock)
func (_Arbnodeinf *ArbnodeinfCallerSession) L2BlockRangeForL1(blockNum uint64) (struct {
	FirstBlock uint64
	LastBlock  uint64
}, error) {
	return _Arbnodeinf.Contract.L2BlockRangeForL1(&_Arbnodeinf.CallOpts, blockNum)
}

// LegacyLookupMessageBatchProof is a free data retrieval call binding the contract method 0x89496270.
//
// Solidity: function legacyLookupMessageBatchProof(uint256 batchNum, uint64 index) view returns(bytes32[] proof, uint256 path, address l2Sender, address l1Dest, uint256 l2Block, uint256 l1Block, uint256 timestamp, uint256 amount, bytes calldataForL1)
func (_Arbnodeinf *ArbnodeinfCaller) LegacyLookupMessageBatchProof(opts *bind.CallOpts, batchNum *big.Int, index uint64) (struct {
	Proof         [][32]byte
	Path          *big.Int
	L2Sender      common.Address
	L1Dest        common.Address
	L2Block       *big.Int
	L1Block       *big.Int
	Timestamp     *big.Int
	Amount        *big.Int
	CalldataForL1 []byte
}, error) {
	var out []interface{}
	err := _Arbnodeinf.contract.Call(opts, &out, "legacyLookupMessageBatchProof", batchNum, index)

	outstruct := new(struct {
		Proof         [][32]byte
		Path          *big.Int
		L2Sender      common.Address
		L1Dest        common.Address
		L2Block       *big.Int
		L1Block       *big.Int
		Timestamp     *big.Int
		Amount        *big.Int
		CalldataForL1 []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Proof = *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)
	outstruct.Path = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.L2Sender = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.L1Dest = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.L2Block = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.L1Block = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Amount = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.CalldataForL1 = *abi.ConvertType(out[8], new([]byte)).(*[]byte)

	return *outstruct, err

}

// LegacyLookupMessageBatchProof is a free data retrieval call binding the contract method 0x89496270.
//
// Solidity: function legacyLookupMessageBatchProof(uint256 batchNum, uint64 index) view returns(bytes32[] proof, uint256 path, address l2Sender, address l1Dest, uint256 l2Block, uint256 l1Block, uint256 timestamp, uint256 amount, bytes calldataForL1)
func (_Arbnodeinf *ArbnodeinfSession) LegacyLookupMessageBatchProof(batchNum *big.Int, index uint64) (struct {
	Proof         [][32]byte
	Path          *big.Int
	L2Sender      common.Address
	L1Dest        common.Address
	L2Block       *big.Int
	L1Block       *big.Int
	Timestamp     *big.Int
	Amount        *big.Int
	CalldataForL1 []byte
}, error) {
	return _Arbnodeinf.Contract.LegacyLookupMessageBatchProof(&_Arbnodeinf.CallOpts, batchNum, index)
}

// LegacyLookupMessageBatchProof is a free data retrieval call binding the contract method 0x89496270.
//
// Solidity: function legacyLookupMessageBatchProof(uint256 batchNum, uint64 index) view returns(bytes32[] proof, uint256 path, address l2Sender, address l1Dest, uint256 l2Block, uint256 l1Block, uint256 timestamp, uint256 amount, bytes calldataForL1)
func (_Arbnodeinf *ArbnodeinfCallerSession) LegacyLookupMessageBatchProof(batchNum *big.Int, index uint64) (struct {
	Proof         [][32]byte
	Path          *big.Int
	L2Sender      common.Address
	L1Dest        common.Address
	L2Block       *big.Int
	L1Block       *big.Int
	Timestamp     *big.Int
	Amount        *big.Int
	CalldataForL1 []byte
}, error) {
	return _Arbnodeinf.Contract.LegacyLookupMessageBatchProof(&_Arbnodeinf.CallOpts, batchNum, index)
}

// NitroGenesisBlock is a free data retrieval call binding the contract method 0x93a2fe21.
//
// Solidity: function nitroGenesisBlock() pure returns(uint256 number)
func (_Arbnodeinf *ArbnodeinfCaller) NitroGenesisBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Arbnodeinf.contract.Call(opts, &out, "nitroGenesisBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NitroGenesisBlock is a free data retrieval call binding the contract method 0x93a2fe21.
//
// Solidity: function nitroGenesisBlock() pure returns(uint256 number)
func (_Arbnodeinf *ArbnodeinfSession) NitroGenesisBlock() (*big.Int, error) {
	return _Arbnodeinf.Contract.NitroGenesisBlock(&_Arbnodeinf.CallOpts)
}

// NitroGenesisBlock is a free data retrieval call binding the contract method 0x93a2fe21.
//
// Solidity: function nitroGenesisBlock() pure returns(uint256 number)
func (_Arbnodeinf *ArbnodeinfCallerSession) NitroGenesisBlock() (*big.Int, error) {
	return _Arbnodeinf.Contract.NitroGenesisBlock(&_Arbnodeinf.CallOpts)
}

// EstimateRetryableTicket is a paid mutator transaction binding the contract method 0xc3dc5879.
//
// Solidity: function estimateRetryableTicket(address sender, uint256 deposit, address to, uint256 l2CallValue, address excessFeeRefundAddress, address callValueRefundAddress, bytes data) returns()
func (_Arbnodeinf *ArbnodeinfTransactor) EstimateRetryableTicket(opts *bind.TransactOpts, sender common.Address, deposit *big.Int, to common.Address, l2CallValue *big.Int, excessFeeRefundAddress common.Address, callValueRefundAddress common.Address, data []byte) (*types.Transaction, error) {
	return _Arbnodeinf.contract.Transact(opts, "estimateRetryableTicket", sender, deposit, to, l2CallValue, excessFeeRefundAddress, callValueRefundAddress, data)
}

// EstimateRetryableTicket is a paid mutator transaction binding the contract method 0xc3dc5879.
//
// Solidity: function estimateRetryableTicket(address sender, uint256 deposit, address to, uint256 l2CallValue, address excessFeeRefundAddress, address callValueRefundAddress, bytes data) returns()
func (_Arbnodeinf *ArbnodeinfSession) EstimateRetryableTicket(sender common.Address, deposit *big.Int, to common.Address, l2CallValue *big.Int, excessFeeRefundAddress common.Address, callValueRefundAddress common.Address, data []byte) (*types.Transaction, error) {
	return _Arbnodeinf.Contract.EstimateRetryableTicket(&_Arbnodeinf.TransactOpts, sender, deposit, to, l2CallValue, excessFeeRefundAddress, callValueRefundAddress, data)
}

// EstimateRetryableTicket is a paid mutator transaction binding the contract method 0xc3dc5879.
//
// Solidity: function estimateRetryableTicket(address sender, uint256 deposit, address to, uint256 l2CallValue, address excessFeeRefundAddress, address callValueRefundAddress, bytes data) returns()
func (_Arbnodeinf *ArbnodeinfTransactorSession) EstimateRetryableTicket(sender common.Address, deposit *big.Int, to common.Address, l2CallValue *big.Int, excessFeeRefundAddress common.Address, callValueRefundAddress common.Address, data []byte) (*types.Transaction, error) {
	return _Arbnodeinf.Contract.EstimateRetryableTicket(&_Arbnodeinf.TransactOpts, sender, deposit, to, l2CallValue, excessFeeRefundAddress, callValueRefundAddress, data)
}

// GasEstimateComponents is a paid mutator transaction binding the contract method 0xc94e6eeb.
//
// Solidity: function gasEstimateComponents(address to, bool contractCreation, bytes data) payable returns(uint64 gasEstimate, uint64 gasEstimateForL1, uint256 baseFee, uint256 l1BaseFeeEstimate)
func (_Arbnodeinf *ArbnodeinfTransactor) GasEstimateComponents(opts *bind.TransactOpts, to common.Address, contractCreation bool, data []byte) (*types.Transaction, error) {
	return _Arbnodeinf.contract.Transact(opts, "gasEstimateComponents", to, contractCreation, data)
}

// GasEstimateComponents is a paid mutator transaction binding the contract method 0xc94e6eeb.
//
// Solidity: function gasEstimateComponents(address to, bool contractCreation, bytes data) payable returns(uint64 gasEstimate, uint64 gasEstimateForL1, uint256 baseFee, uint256 l1BaseFeeEstimate)
func (_Arbnodeinf *ArbnodeinfSession) GasEstimateComponents(to common.Address, contractCreation bool, data []byte) (*types.Transaction, error) {
	return _Arbnodeinf.Contract.GasEstimateComponents(&_Arbnodeinf.TransactOpts, to, contractCreation, data)
}

// GasEstimateComponents is a paid mutator transaction binding the contract method 0xc94e6eeb.
//
// Solidity: function gasEstimateComponents(address to, bool contractCreation, bytes data) payable returns(uint64 gasEstimate, uint64 gasEstimateForL1, uint256 baseFee, uint256 l1BaseFeeEstimate)
func (_Arbnodeinf *ArbnodeinfTransactorSession) GasEstimateComponents(to common.Address, contractCreation bool, data []byte) (*types.Transaction, error) {
	return _Arbnodeinf.Contract.GasEstimateComponents(&_Arbnodeinf.TransactOpts, to, contractCreation, data)
}

// GasEstimateL1Component is a paid mutator transaction binding the contract method 0x77d488a2.
//
// Solidity: function gasEstimateL1Component(address to, bool contractCreation, bytes data) payable returns(uint64 gasEstimateForL1, uint256 baseFee, uint256 l1BaseFeeEstimate)
func (_Arbnodeinf *ArbnodeinfTransactor) GasEstimateL1Component(opts *bind.TransactOpts, to common.Address, contractCreation bool, data []byte) (*types.Transaction, error) {
	return _Arbnodeinf.contract.Transact(opts, "gasEstimateL1Component", to, contractCreation, data)
}

// GasEstimateL1Component is a paid mutator transaction binding the contract method 0x77d488a2.
//
// Solidity: function gasEstimateL1Component(address to, bool contractCreation, bytes data) payable returns(uint64 gasEstimateForL1, uint256 baseFee, uint256 l1BaseFeeEstimate)
func (_Arbnodeinf *ArbnodeinfSession) GasEstimateL1Component(to common.Address, contractCreation bool, data []byte) (*types.Transaction, error) {
	return _Arbnodeinf.Contract.GasEstimateL1Component(&_Arbnodeinf.TransactOpts, to, contractCreation, data)
}

// GasEstimateL1Component is a paid mutator transaction binding the contract method 0x77d488a2.
//
// Solidity: function gasEstimateL1Component(address to, bool contractCreation, bytes data) payable returns(uint64 gasEstimateForL1, uint256 baseFee, uint256 l1BaseFeeEstimate)
func (_Arbnodeinf *ArbnodeinfTransactorSession) GasEstimateL1Component(to common.Address, contractCreation bool, data []byte) (*types.Transaction, error) {
	return _Arbnodeinf.Contract.GasEstimateL1Component(&_Arbnodeinf.TransactOpts, to, contractCreation, data)
}

