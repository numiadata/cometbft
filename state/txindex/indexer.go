package txindex

import (
	"github.com/numiadata/tools/pubsub"
)

// XXX/TODO: These types should be moved to the indexer package.

//go:generate ../../scripts/mockery_generate.sh TxIndexer

// TxIndexer interface defines methods to index and search transactions.
type TxIndexer interface {
	// AddBatch analyzes, indexes and stores a batch of transactions.
	AddBatch(b *pubsub.Batch) error

	// Index analyzes, indexes and stores a single transaction.
	Index(result *pubsub.TxResult) error

	// Get returns the transaction specified by hash or nil if the transaction is not indexed
	// or stored.
	Get(hash []byte) (*pubsub.TxResult, error)
}
