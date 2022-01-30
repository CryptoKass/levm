package tools

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/trie"
	com "github.com/sledro/levm/common"
)

func StartTrieDB(edb ethdb.Database) *trie.Trie {
	tdb := trie.NewDatabase(edb)
	tr, err := trie.New(common.Hash{}, tdb)
	com.PanicErr(err)
	return tr
}
