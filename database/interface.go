package database

import (
	"github.com/ninjadotorg/cash-prototype/common"
)

// DB provides the interface that is used to store blocks.
type DB interface {
	StoreBlock(interface{}, byte) error
	FetchBlock(*common.Hash) ([]byte, error)
	HasBlock(*common.Hash) (bool, error)
	FetchAllBlocks() ([][]*common.Hash, error)
	FetchChainBlocks(byte) ([]*common.Hash, error)

	StoreBestBlock(interface{}, byte) error
	FetchBestState(byte) ([]byte, error)

	StoreNullifiers([]byte, string, byte) error
	FetchNullifiers(string, byte) ([][]byte, error)
	HasNullifier([]byte, string, byte) (bool, error)
	StoreCommitments([]byte, string, byte) error
	FetchCommitments(string, byte) ([][]byte, error)
	HasCommitment([]byte, string, byte) (bool, error)

	StoreBlockIndex(*common.Hash, int32, byte) error
	GetIndexOfBlock(*common.Hash) (int32, byte, error)
	GetBlockByIndex(int32, byte) (*common.Hash, error)

	StoreFeeEstimator([]byte, byte) error
	GetFeeEstimator(byte) ([]byte, error)

	Close() error
}
