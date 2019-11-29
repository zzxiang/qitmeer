// Copyright (c) 2017-2018 The qitmeer developers
// Copyright (c) 2014-2016 The btcsuite developers
// Copyright (c) 2015-2016 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package params

import (
	"github.com/Qitmeer/qitmeer/common/hash"
	"github.com/Qitmeer/qitmeer/core/protocol"
	"github.com/Qitmeer/qitmeer/core/types"
	"github.com/Qitmeer/qitmeer/core/types/pow"
	"github.com/Qitmeer/qitmeer/ledger"
	"time"
)

// MainNet ------------------------------------------------------------------------

// genesisCoinbaseTx is the coinbase transaction for the genesis blocks for
// the main network.
func buildGenesisCoinbaseTx(net protocol.Network) types.Transaction {
	tx := types.Transaction{
		Version: 1,
		TxIn: []*types.TxInput{
			{
				// Fully null.
				PreviousOut: types.TxOutPoint{
					Hash:     hash.Hash{},
					OutIndex: 0xffffffff,
				},
				SignScript: []byte{
					0x00, 0x00,
				},
				Sequence: 0xffffffff,
			},
		},
		LockTime: 0,
		Expire:   0,
	}
	ledger.Ledger(&tx, net)
	return tx
}

var genesisCoinbaseTx = buildGenesisCoinbaseTx(protocol.MainNet)

// mainnetgenesisMerkleRoot is the hash of the first transaction in the genesis block
// for the main network.
var genesisMerkleRoot = genesisCoinbaseTx.TxHashFull()

// genesisBlock defines the genesis block of the block chain which serves as the
// public transaction ledger for the main network.
//
// The genesis block for mainnet, testnet, and privnet are not evaluated
// for proof of work. The only values that are ever used elsewhere in the
// blockchain from it are:
// (1) The genesis block hash is used as the PrevBlock in params.go.
// (2) The difficulty starts off at the value given by bits.
// (3) The stake difficulty starts off at the value given by SBits.
// (4) The timestamp, which guides when blocks can be built on top of it
//      and what the initial difficulty calculations come out to be.
//
// The genesis block is valid by definition and none of the fields within
// it are validated for correctness.
var genesisBlock = types.Block{
	Header: types.BlockHeader{
		ParentRoot: hash.Hash{},
		TxRoot:     genesisMerkleRoot,
		//UtxoCommitment: types.Hash{},
		//CompactFilter: types.Hash{},
		StateRoot:  hash.Hash{},
		Timestamp:  time.Unix(1561939200, 0), // 2019-07-01 00:00:00 GMT
		Difficulty: 0x1b01ffff,               // Difficulty 32767
		Pow:        &pow.Blake2bd{},
	},
	Transactions: []*types.Transaction{&genesisCoinbaseTx},
}

// genesisHash is the hash of the first block in the block chain for the main
// network (genesis block).
var genesisHash = genesisBlock.BlockHash()

// TestNet ------------------------------------------------------------------------

// testNetGenesisCoinbaseTx is the coinbase transaction for the genesis blocks for
// the test network.
func buildTestNetGenesisCoinbaseTx(net protocol.Network) types.Transaction {
	tx := types.Transaction{
		Version: 1,
		TxIn: []*types.TxInput{
			{
				// Fully null.
				PreviousOut: types.TxOutPoint{
					Hash:     hash.Hash{},
					OutIndex: 0xffffffff,
				},
				SignScript: []byte{
					0x54, 0x68, 0x65, 0x20, 0x48, 0x4c, 0x43, 0x20, 0x54, 0x6f, 0x6b, 0x65,
					0x6e, 0x20, 0x44, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x65, 0x64, 0x20,
					0x32, 0x30, 0x30, 0x32, 0x38, 0x37, 0x39, 0x31, 0x31, 0x20, 0x37, 0x64,
					0x33, 0x38, 0x34, 0x36, 0x61, 0x62, 0x31, 0x36, 0x64, 0x66, 0x61, 0x35,
					0x31, 0x31, 0x30, 0x37, 0x39, 0x36, 0x66, 0x36, 0x62, 0x35, 0x66, 0x33,
					0x31, 0x33, 0x36, 0x65, 0x64, 0x66, 0x64, 0x63, 0x36, 0x64, 0x39, 0x66,
					0x64, 0x64, 0x37, 0x62, 0x33, 0x33, 0x36, 0x62, 0x30, 0x36, 0x37, 0x66,
					0x66, 0x38, 0x34, 0x33, 0x35, 0x33, 0x35, 0x36, 0x63, 0x30, 0x38, 0x64,
					0x36, 0x61, 0x0a,
					0x54, 0x68, 0x6f, 0x73, 0x65, 0x20, 0x77, 0x68, 0x6f, 0x2c, 0x20, 0x77,
					0x68, 0x65, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x79, 0x20, 0x73, 0x70, 0x65,
					0x6e, 0x64, 0x2c, 0x20, 0x61, 0x72, 0x65, 0x20, 0x6e, 0x6f, 0x74, 0x20,
					0x65, 0x78, 0x74, 0x72, 0x61, 0x76, 0x61, 0x67, 0x61, 0x6e, 0x74, 0x20,
					0x61, 0x6e, 0x64, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x6e, 0x69, 0x67, 0x67,
					0x61, 0x72, 0x64, 0x6c, 0x79, 0x2c, 0x20, 0x62, 0x75, 0x74, 0x20, 0x68,
					0x6f, 0x6c, 0x64, 0x20, 0x61, 0x20, 0x6a, 0x75, 0x73, 0x74, 0x20, 0x28,
					0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x29, 0x20, 0x62, 0x65, 0x74,
					0x77, 0x65, 0x65, 0x6e, 0x20, 0x74, 0x68, 0x6f, 0x73, 0x65, 0x20, 0x28,
					0x65, 0x78, 0x74, 0x72, 0x65, 0x6d, 0x65, 0x73, 0x29, 0x2e, 0x20, 0x28,
					0x51, 0x75, 0x72, 0x61, 0x6e, 0x20, 0x32, 0x35, 0x3a, 0x36, 0x37, 0x29,
					0x0a,
				},
				Sequence: 0xffffffff,
			},
		},
		LockTime: 0,
		Expire:   0,
	}
	ledger.Ledger(&tx, net)
	return tx
}
//
var testNetGenesisCoinbaseTx = buildTestNetGenesisCoinbaseTx(protocol.TestNet)

// testNetGenesisMerkleRoot is the hash of the first transaction in the genesis block
// for the test network.
var testNetGenesisMerkleRoot = testNetGenesisCoinbaseTx.TxHashFull()

// testNetGenesisBlock defines the genesis block of the block chain which
// serves as the public transaction ledger for the test network (version 3).
var testNetGenesisBlock = types.Block{
	Header: types.BlockHeader{
		Version:    2,
		ParentRoot: hash.Hash{},
		TxRoot:     testNetGenesisMerkleRoot,
		Timestamp:  time.Unix(1574726400, 0), // 2019-11-26 00:00:00 GMT
		Difficulty: 0x1c1fffff,
		Pow:        &pow.Blake2bd{},
	},
	Transactions: []*types.Transaction{&testNetGenesisCoinbaseTx},
}

// testNetGenesisHash is the hash of the first block in the block chain for the
// test network.
var testNetGenesisHash = testNetGenesisBlock.BlockHash()

// PrivNet -------------------------------------------------------------------------

var privNetGenesisCoinbaseTx = types.Transaction{
	Version: 1,
	TxIn: []*types.TxInput{
		{
			PreviousOut: types.TxOutPoint{
				Hash:     hash.Hash{},
				OutIndex: 0xffffffff,
			},
			Sequence: 0xffffffff,
			SignScript: []byte{
				0x04, 0xff, 0xff, 0x00, 0x1d, 0x01, 0x04, 0x45, /* |.......E| */
				0x54, 0x68, 0x65, 0x20, 0x54, 0x69, 0x6d, 0x65, /* |The Time| */
				0x73, 0x20, 0x30, 0x33, 0x2f, 0x4a, 0x61, 0x6e, /* |s 03/Jan| */
				0x2f, 0x32, 0x30, 0x30, 0x39, 0x20, 0x43, 0x68, /* |/2009 Ch| */
				0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x6f, 0x72, /* |ancellor| */
				0x20, 0x6f, 0x6e, 0x20, 0x62, 0x72, 0x69, 0x6e, /* | on brin| */
				0x6b, 0x20, 0x6f, 0x66, 0x20, 0x73, 0x65, 0x63, /* |k of sec|*/
				0x6f, 0x6e, 0x64, 0x20, 0x62, 0x61, 0x69, 0x6c, /* |ond bail| */
				0x6f, 0x75, 0x74, 0x20, 0x66, 0x6f, 0x72, 0x20, /* |out for |*/
				0x62, 0x61, 0x6e, 0x6b, 0x73, /* |banks| */
			},
		},
	},
	TxOut: []*types.TxOutput{
		{
			Amount: 0x00000000,
			PkScript: []byte{
				0x41, 0x04, 0x67, 0x8a, 0xfd, 0xb0, 0xfe, 0x55, /* |A.g....U| */
				0x48, 0x27, 0x19, 0x67, 0xf1, 0xa6, 0x71, 0x30, /* |H'.g..q0| */
				0xb7, 0x10, 0x5c, 0xd6, 0xa8, 0x28, 0xe0, 0x39, /* |..\..(.9| */
				0x09, 0xa6, 0x79, 0x62, 0xe0, 0xea, 0x1f, 0x61, /* |..yb...a| */
				0xde, 0xb6, 0x49, 0xf6, 0xbc, 0x3f, 0x4c, 0xef, /* |..I..?L.| */
				0x38, 0xc4, 0xf3, 0x55, 0x04, 0xe5, 0x1e, 0xc1, /* |8..U....| */
				0x12, 0xde, 0x5c, 0x38, 0x4d, 0xf7, 0xba, 0x0b, /* |..\8M...| */
				0x8d, 0x57, 0x8a, 0x4c, 0x70, 0x2b, 0x6b, 0xf1, /* |.W.Lp+k.| */
				0x1d, 0x5f, 0xac, /* |._.| */
			},
		},
	},
	LockTime: 0,
	Expire:   0,
}

// privNetGenesisMerkleRoot is the hash of the first transaction in the genesis
// block for the simulation test network.  It is the same as the merkle root for
// the main network.
var privNetGenesisMerkleRoot = privNetGenesisCoinbaseTx.TxHashFull()

var zeroHash = hash.ZeroHash

// privNetGenesisBlock defines the genesis block of the block chain which serves
// as the public transaction ledger for the simulation test network.
var privNetGenesisBlock = types.Block{
	Header: types.BlockHeader{
		ParentRoot: zeroHash,
		TxRoot:     privNetGenesisMerkleRoot,
		StateRoot: hash.Hash([32]byte{ // Make go vet happy.
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		}),
		Timestamp:  time.Unix(1530833717, 0), // 2018-07-05 23:35:17 GMT
		Difficulty: 0x207fffff,               // 545259519
		Pow:        &pow.Blake2bd{},
	},
	Transactions: []*types.Transaction{&privNetGenesisCoinbaseTx},
}

// privNetGenesisHash is the hash of the first block in the block chain for the
// private test network.
var privNetGenesisHash = privNetGenesisBlock.BlockHash()

// TestMixNet ------------------------------------------------------------------------

//
var testPowNetGenesisCoinbaseTx = types.Transaction{}

// testNetGenesisMerkleRoot is the hash of the first transaction in the genesis block
// for the test network.
var testPowNetGenesisMerkleRoot = testPowNetGenesisCoinbaseTx.TxHashFull()

// testNetGenesisBlock defines the genesis block of the block chain which
// serves as the public transaction ledger for the test network (version 3).
var testPowNetGenesisBlock = types.Block{
	Header: types.BlockHeader{
		Version:    16,
		ParentRoot: hash.Hash{},
		TxRoot:     testPowNetGenesisMerkleRoot,
		Timestamp:  time.Unix(1547735581, 0), // 2019-01-17 14:33:12 GMT
		Difficulty: 0x1e00ffff,
		Pow:        &pow.Blake2bd{},
	},
	Transactions: []*types.Transaction{&testPowNetGenesisCoinbaseTx},
	Parents:      []*hash.Hash{},
}

// testNetGenesisHash is the hash of the first block in the block chain for the
// test network.
var testPowNetGenesisHash = testPowNetGenesisBlock.BlockHash()
