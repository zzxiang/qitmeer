/*
 * Copyright (c) 2017-2020 The qitmeer developers
 */

package token

import (
	"fmt"
	"github.com/Qitmeer/qitmeer/core/dbnamespace"
	"github.com/Qitmeer/qitmeer/core/json"
	"github.com/Qitmeer/qitmeer/core/serialization"
	"github.com/Qitmeer/qitmeer/core/types"
	"github.com/Qitmeer/qitmeer/database"
)

// tokenState specifies the token balance of the current block.
// the updates are written in the same order as the tx in the block, which is
// used to verify the correctness of the token balance
type TokenState struct {
	PrevStateID uint32
	Types       TokenTypesMap
	Balances    TokenBalancesMap
	Updates     []ITokenUpdate
}

func (ts *TokenState) GetTokenBalances() []json.TokenBalance {
	tbs := []json.TokenBalance{}
	for k, v := range ts.Balances {
		tb := json.TokenBalance{CoinId: uint16(k), CoinName: k.Name(), Balance: v.Balance, LockedMeer: v.LockedMeer}
		tbs = append(tbs, tb)
	}
	return tbs
}

// Serialize function will serialize the token state into byte slice
func (ts *TokenState) Serialize() ([]byte, error) {
	// total number of bytes to serialize
	serializeSize := serialization.SerializeSizeVLQ(uint64(ts.PrevStateID))

	serializeSize += serialization.SerializeSizeVLQ(uint64(len(ts.Balances)))
	for id, b := range ts.Balances {
		// sanity check
		if id == types.MEERID || b.Balance < 0 || b.LockedMeer < 0 {
			return nil, fmt.Errorf("invalid token balance {%v, %v}", id, b)
		}
		serializeSize += serialization.SerializeSizeVLQ(uint64(id))
		serializeSize += serialization.SerializeSizeVLQ(uint64(b.Balance))
		serializeSize += serialization.SerializeSizeVLQ(uint64(b.LockedMeer))
	}
	serializeSize += serialization.SerializeSizeVLQ(uint64(len(ts.Updates)))
	serialized := make([]byte, serializeSize)
	offset := 0
	offset = serialization.PutVLQ(serialized, uint64(ts.PrevStateID))

	offset += serialization.PutVLQ(serialized[offset:], uint64(len(ts.Balances)))
	for id, b := range ts.Balances {
		offset += serialization.PutVLQ(serialized[offset:], uint64(id))
		offset += serialization.PutVLQ(serialized[offset:], uint64(b.Balance))
		offset += serialization.PutVLQ(serialized[offset:], uint64(b.LockedMeer))
	}

	offset += serialization.PutVLQ(serialized[offset:], uint64(len(ts.Updates)))

	for _, v := range ts.Updates {
		uSerialized, err := v.Serialize()
		if err != nil {
			return nil, err
		}
		serialized = append(serialized, uSerialized...)
	}
	return serialized, nil
}

// Deserialize function will deserializes token state from the byte slice
func (ts *TokenState) Deserialize(data []byte) (int, error) {
	prevStateID, offset := serialization.DeserializeVLQ(data)
	if offset == 0 {
		return offset, fmt.Errorf("unexpected end of data while reading prevStateID")
	}
	// Deserialize the balance.
	var balances map[types.CoinID]TokenBalance
	numOfBalances, bytesRead := serialization.DeserializeVLQ(data[offset:])
	if bytesRead == 0 {
		return offset, fmt.Errorf("unexpected end of data while reading number of balances")
	}
	offset += bytesRead

	if numOfBalances > 0 {
		balances = make(map[types.CoinID]TokenBalance, numOfBalances)
		for i := uint64(0); i < numOfBalances; i++ {
			// token id
			derId, bytesRead := serialization.DeserializeVLQ(data[offset:])
			if bytesRead == 0 {
				return offset, fmt.Errorf("unexpected end of data while reading token id at balances{%d}", i)
			}
			offset += bytesRead

			// token balance
			balance, bytesRead := serialization.DeserializeVLQ(data[offset:])
			if bytesRead == 0 {
				return offset, fmt.Errorf("unexpected end of data while reading balance at balances{%d}", i)
			}
			offset += bytesRead

			// locked meer
			lockedMeer, bytesRead := serialization.DeserializeVLQ(data[offset:])
			if bytesRead == 0 {
				return offset, fmt.Errorf("unexpected end of data while reading balance at balances{%d}", i)
			}
			offset += bytesRead

			id := types.CoinID(uint16(derId))
			balances[id] = TokenBalance{int64(balance), int64(lockedMeer)}
		}
	}
	updates := []ITokenUpdate{}
	numOfUpdates, bytesRead := serialization.DeserializeVLQ(data[offset:])
	if bytesRead == 0 {
		return offset, fmt.Errorf("unexpected end of data while reading number of balances")
	}
	offset += bytesRead

	if numOfUpdates > 0 {
		for i := uint64(0); i < numOfUpdates; i++ {
			//type
			tu := TokenUpdate{}
			bytesRead, err := tu.Deserialize(data[offset:])
			if err != nil {
				return offset, err
			}
			update := NewTokenUpdate(tu.GetType())
			bytesRead, err = update.Deserialize(data[offset:])
			if err != nil {
				return offset, err
			}
			offset += bytesRead
			updates = append(updates, update)
		}
	}
	ts.PrevStateID = uint32(prevStateID)
	ts.Balances = balances
	ts.Updates = updates

	return offset, nil
}

func (ts *TokenState) Commit() {
	for _, tu := range ts.Updates {
		if bu, ok := tu.(*BalanceUpdate); ok {
			ts.Balances.UpdateBalance(bu)
		}
	}
}

// dbPutTokenState put a token balance record into the token state database.
// the key is the provided block hash
func DBPutTokenState(dbTx database.Tx, bid uint32, ts *TokenState) error {
	// Serialize the current token state.
	serializedData, err := ts.Serialize()
	if err != nil {
		return err
	}
	// Store the current token balance record into the token state database.
	meta := dbTx.Metadata()
	bucket := meta.Bucket(dbnamespace.TokenBucketName)
	var serializedID [4]byte
	dbnamespace.ByteOrder.PutUint32(serializedID[:], bid)
	return bucket.Put(serializedID[:], serializedData)
}

// dbFetchTokenState fetch the token balance record from the token state database.
// the key is the input block hash.
func DBFetchTokenState(dbTx database.Tx, bid uint32) (*TokenState, error) {
	// if it is genesis hash, return empty tokenState directly
	if bid == 0 {
		return &TokenState{}, nil
	}
	// Fetch record from the token state database by block hash
	meta := dbTx.Metadata()
	bucket := meta.Bucket(dbnamespace.TokenBucketName)

	var serializedID [4]byte
	dbnamespace.ByteOrder.PutUint32(serializedID[:], bid)
	v := bucket.Get(serializedID[:])
	if v == nil {
		return nil, fmt.Errorf("tokenstate db can't find record from block id : %v", bid)
	}
	// deserialize the fetched token state record
	ts := TokenState{}
	_, err := ts.Deserialize(v)
	return &ts, err
}

func DBRemoveTokenState(dbTx database.Tx, id uint32) error {
	bucket := dbTx.Metadata().Bucket(dbnamespace.TokenBucketName)
	var serializedID [4]byte
	dbnamespace.ByteOrder.PutUint32(serializedID[:], id)

	key := serializedID[:]
	return bucket.Delete(key)
}

func CheckUnMintUpdate(update *BalanceUpdate) error {
	if update.Typ != types.TxTypeTokenUnmint {
		return fmt.Errorf("checkUnMintUpdate : wrong update type %v", update.Typ)
	}
	if err := checkUpdateCommon(update); err != nil {
		return err
	}
	return nil
}

func CheckMintUpdate(update *BalanceUpdate) error {
	if update.Typ != types.TxTypeTokenMint {
		return fmt.Errorf("checkUnMintUpdate : wrong update type %v", update.Typ)
	}
	if err := checkUpdateCommon(update); err != nil {
		return err
	}
	return nil
}

func checkUpdateCommon(update *BalanceUpdate) error {
	if !types.IsKnownCoinID(update.TokenAmount.Id) {
		return fmt.Errorf("checkUpdateCommon : unknown token id %v", update.TokenAmount.Id.Name())
	}
	if update.TokenAmount.Value <= 0 {
		return fmt.Errorf("checkUpdateCommon : wrong token amount : %v", update.TokenAmount.Value)
	}
	if update.MeerAmount <= 0 {
		return fmt.Errorf("checkUpdateCommon : wrong meer amount : %v", update.MeerAmount)
	}
	return nil
}
