package vminterface

import (
	com "github.com/cryptokass/levm/common"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/ethdb"
)

// NewStateDB - Create a new StateDB using levelDB instead of RAM
func NewStateDB(root common.Hash, dbPath string) (*state.StateDB, ethdb.Database) {

	// open ethdb
	/*edb, err := ethdb.NewLDBDatabase(dbPath, 100, 100)
	db := state.NewDatabase(edb)
	com.PanicErr(err)
	*/

	edb, _ := rawdb.NewLevelDBDatabase(dbPath, 100, 100, "")
	//edb := rawdb.NewMemoryDatabase()
	db := state.NewDatabase(edb)

	// make statedb
	stateDB, err := state.New(root, db)
	com.PanicErr(err)

	return stateDB, edb

}
