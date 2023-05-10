package indexer

import (
	"github.com/tendermint/tendermint/types"
)

//go:generate ../../scripts/mockery_generate.sh BlockIndexer

// BlockIndexer defines an interface contract for indexing block events.
type BlockIndexer interface {
	// Has returns true if the given height has been indexed. An error is returned
	// upon database query failure.
	Has(height int64) (bool, error)

	// Index indexes BeginBlock and EndBlock events for a given block by its height.
	Index(types.EventDataNewBlockHeader) error
}
