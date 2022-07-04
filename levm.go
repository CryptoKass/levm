package levm

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/eth/tracers/logger"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/params"

	vmi "github.com/sledro/levm/vminterface"
)

// LEVM is a container for the go-ethereum EVM
// with methods to create and call contracts.
//
// LEVM contains the two most important objects
// for interacting with the EVM: stateDB and
// vm.EVM. The LEVM should be created with the
// LEVM.New() method, unless you know what you
// doing.
type LEVM struct {
	stateDB *state.StateDB
	evm     *vm.EVM
	edb     ethdb.Database
}

// New creates a new instace of the LEVM
func New(dbPath string, blockNumber *big.Int, origin common.Address) *LEVM {
	// create blank LEVM instance:
	lvm := LEVM{}

	// setup storage using dbpath
	lvm.stateDB, lvm.edb = vmi.NewStateDB(common.Hash{}, dbPath)

	// update the evm - creates new EVM
	lvm.NewEVM(blockNumber, origin)

	return &lvm
}

// NewEVM creates a fresh evm instance with
// new origin and blocknumber and time.
// This method recreates the contained EVM while
// keeping the stateDB the same.
func (lvm *LEVM) NewEVM(blockNumber *big.Int, origin common.Address) {

	// create contexted for the evm context
	chainContext := vmi.NewChainContext(origin)
	blockContext := vmi.NewBlockContext(origin, origin, blockNumber, chainContext)
	txContext := vmi.NewTxContext(origin)

	tcr := logger.NewStructLogger(&logger.Config{})

	// create vm config
	vmConfig := vm.Config{Debug: true, Tracer: tcr, ExtraEips: []int{150, 1052, 1884}, NoBaseFee: true}

	// create the evm
	lvm.evm = vm.NewEVM(blockContext, txContext, lvm.stateDB, params.MainnetChainConfig, vmConfig)
}

// DeployContract will create and deploy a new
// contract from the contract data.
func (lvm *LEVM) DeployContract(fromAddr common.Address, contractData []byte) ([]byte, common.Address, uint64, error) {

	// Get reference to the transaction sender
	contractRef := vm.AccountRef(fromAddr)
	leftOver := big.NewInt(0)

	return lvm.evm.Create(
		contractRef,
		contractData,
		lvm.stateDB.GetBalance(fromAddr).Uint64(),
		leftOver,
	)
}

// CallContract - make a call to a Contract Method
// using prepacked Inputs. To use ABI directly try
// lvm.CallContractABI()
func (lvm *LEVM) CallContract(callerAddr, contractAddr common.Address, value *big.Int, inputs []byte) ([]byte, error) {
	// Get reference to the transaction sender
	callerRef := vm.AccountRef(callerAddr)
	output, gas, err := lvm.evm.Call(
		callerRef,
		contractAddr,
		inputs,
		lvm.stateDB.GetBalance(callerAddr).Uint64(),
		value,
	)
	lvm.stateDB.SetBalance(callerAddr, big.NewInt(0).SetUint64(gas))
	return output, err
}

// CallContractABI - make a call to a Contract Method
// using the ABI.
func (lvm *LEVM) CallContractABI(callerAddr, contractAddr common.Address, value *big.Int, abiObject abi.ABI, funcName string, args ...interface{}) ([]byte, error) {

	inputs, err := abiObject.Pack(funcName, args...)
	if err != nil {
		return nil, err
	}

	callerRef := vm.AccountRef(callerAddr)
	output, gas, err := lvm.evm.Call(
		callerRef,
		contractAddr,
		inputs,
		lvm.stateDB.GetBalance(callerAddr).Uint64(),
		value,
	)
	lvm.stateDB.SetBalance(callerAddr, big.NewInt(0).SetUint64(gas))
	return output, err
}
