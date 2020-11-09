package tools

import (
	com "github.com/sledro/levm/common"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/trie"
)

func StartTrieDB(edb ethdb.Database) *trie.Trie {
	tdb := trie.NewDatabase(edb)
	tr, err := trie.New(common.Hash{}, tdb)
	com.PanicErr(err)
	return tr
}
