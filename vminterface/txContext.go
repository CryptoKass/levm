package vminterface

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// NewVMContext will construct a new EVM Context with default values.
// TODO: include gas price variable in params
func NewTxContext(origin common.Address) vm.TxContext {
	return vm.TxContext{
		Origin:   origin,
		GasPrice: big.NewInt(1),
	}
}
