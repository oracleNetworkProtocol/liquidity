package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName is the name of the liquidity module
	ModuleName = "liquidity"

	// RouterKey is the message router key for the liquidity module
	RouterKey = ModuleName

	// StoreKey is the default store key for the liquidity module
	StoreKey = ModuleName

	// QuerierRoute is the querier route for the liquidity module
	QuerierRoute = ModuleName

	// PoolCoinDenomPrefix is the prefix used for liquidity pool coin representation
	PoolCoinDenomPrefix = "pool"
)

var (
	// param key for global Liquidity Pool IDs
	GlobalLiquidityPoolIDKey = []byte("globalLiquidityPoolId")

	PoolKeyPrefix                  = []byte{0x11}
	PoolByReserveAccIndexKeyPrefix = []byte{0x12}

	PoolBatchIndexKeyPrefix = []byte{0x21} // Last PoolBatchIndex
	PoolBatchKeyPrefix      = []byte{0x22}

	PoolBatchDepositMsgStateIndexKeyPrefix  = []byte{0x31}
	PoolBatchWithdrawMsgStateIndexKeyPrefix = []byte{0x32}
	PoolBatchSwapMsgStateIndexKeyPrefix     = []byte{0x33}
	PoolDepositSuccessMsgsKeyPrefix         = []byte{0x34}
	PoolDepositSuccessMsgIndexKeyPrefix     = []byte{0x35}
	PoolSwapSuccessMsgKeyPrefix             = []byte{0x36}
	PoolSwapSuccessMsgIndexKeyPrefix        = []byte{0x37}
)

// GetPoolSwapSuccessMsgsKeyPrefix returns kv indexing key of the latest index value of the msg index
func GetPoolSwapSuccessMsgsPrefix(poolID uint64) []byte {
	key := make([]byte, 9)
	key[0] = PoolSwapSuccessMsgKeyPrefix[0]
	copy(key[1:9], sdk.Uint64ToBigEndian(poolID))
	return key
}

func GetPoolSwapSuccessMsgsAddressPrefix(poolID uint64, depositorAcc sdk.AccAddress) []byte {
	key := GetPoolSwapSuccessMsgsPrefix(poolID)

	key = append(key, depositorAcc.Bytes()...)
	return key
}

// GetPoolDepositSuccessMsgPrefix returns prefix of deposit success message in the pool's
func GetPoolSwapSuccessMsgIndexPrefix(poolID, msgIndex uint64) []byte {
	key := make([]byte, 17)
	key[0] = PoolSwapSuccessMsgIndexKeyPrefix[0]
	copy(key[1:9], sdk.Uint64ToBigEndian(poolID))
	copy(key[9:17], sdk.Uint64ToBigEndian(msgIndex))
	return key
}

// GetPoolSwapSuccessMsgsKeyPrefix returns kv indexing key of the latest index value of the msg index
func GetPoolSwapSuccessMsgsAddressIndexPrefix(poolID uint64, swapRequesterAddress sdk.AccAddress, msgIndex uint64) []byte {
	key := GetPoolSwapSuccessMsgsPrefix(poolID)
	key = append(key, swapRequesterAddress.Bytes()...)
	key = append(key, sdk.Uint64ToBigEndian(msgIndex)...)
	return key
}

// GetPoolDepositSuccessMsgPrefix returns prefix of deposit success message in the pool's
func GetPoolDepositSuccessMsgsPrefix(poolID uint64) []byte {
	key := make([]byte, 9)
	key[0] = PoolDepositSuccessMsgsKeyPrefix[0]
	copy(key[1:9], sdk.Uint64ToBigEndian(poolID))
	return key
}

// GetPoolDepositSuccessMsgPrefix returns prefix of deposit success message in the pool's
func GetPoolDepositSuccessMsgIndexPrefix(poolID, msgIndex uint64) []byte {
	key := make([]byte, 17)
	key[0] = PoolDepositSuccessMsgIndexKeyPrefix[0]
	copy(key[1:9], sdk.Uint64ToBigEndian(poolID))
	copy(key[9:17], sdk.Uint64ToBigEndian(msgIndex))
	return key
}

func GetPoolDepositSuccessMsgsAddressPrefix(poolID uint64, depositorAcc sdk.AccAddress) []byte {
	key := GetPoolDepositSuccessMsgsPrefix(poolID)

	key = append(key, depositorAcc.Bytes()...)
	return key
}

func GetPoolDepositSuccessMsgAddressIndexPrefix(poolID uint64, depositorAcc sdk.AccAddress, msgIndex uint64) []byte {
	key := GetPoolDepositSuccessMsgsPrefix(poolID)
	key = append(key, depositorAcc.Bytes()...)
	key = append(key, sdk.Uint64ToBigEndian(msgIndex)...)
	return key
}

// GetPoolKey returns kv indexing key of the pool
func GetPoolKey(poolID uint64) []byte {
	key := make([]byte, 9)
	key[0] = PoolKeyPrefix[0]
	copy(key[1:], sdk.Uint64ToBigEndian(poolID))
	return key
}

// GetPoolByReserveAccIndexKey returns kv indexing key of the pool indexed by reserve account
func GetPoolByReserveAccIndexKey(reserveAcc sdk.AccAddress) []byte {
	return append(PoolByReserveAccIndexKeyPrefix, reserveAcc.Bytes()...)
}

// GetPoolBatchIndexKey returns kv indexing key of the latest index value of the pool batch
func GetPoolBatchIndexKey(poolID uint64) []byte {
	key := make([]byte, 9)
	key[0] = PoolBatchIndexKeyPrefix[0]
	copy(key[1:9], sdk.Uint64ToBigEndian(poolID))
	return key
}

// GetPoolBatchKey returns kv indexing key of the pool batch indexed by pool id
func GetPoolBatchKey(poolID uint64) []byte {
	key := make([]byte, 9)
	key[0] = PoolBatchKeyPrefix[0]
	copy(key[1:9], sdk.Uint64ToBigEndian(poolID))
	return key
}

// GetPoolBatchDepositMsgStatesPrefix returns prefix of deposit message states in the pool's latest batch for iteration
func GetPoolBatchDepositMsgStatesPrefix(poolID uint64) []byte {
	key := make([]byte, 9)
	key[0] = PoolBatchDepositMsgStateIndexKeyPrefix[0]
	copy(key[1:9], sdk.Uint64ToBigEndian(poolID))
	return key
}

// GetPoolBatchWithdrawMsgsPrefix returns prefix of withdraw message states in the pool's latest batch for iteration
func GetPoolBatchWithdrawMsgsPrefix(poolID uint64) []byte {
	key := make([]byte, 9)
	key[0] = PoolBatchWithdrawMsgStateIndexKeyPrefix[0]
	copy(key[1:9], sdk.Uint64ToBigEndian(poolID))
	return key
}

// GetPoolBatchSwapMsgStatesPrefix returns prefix of swap message states in the pool's latest batch for iteration
func GetPoolBatchSwapMsgStatesPrefix(poolID uint64) []byte {
	key := make([]byte, 9)
	key[0] = PoolBatchSwapMsgStateIndexKeyPrefix[0]
	copy(key[1:9], sdk.Uint64ToBigEndian(poolID))
	return key
}

// GetPoolBatchDepositMsgStateIndexKey returns kv indexing key of the latest index value of the msg index
func GetPoolBatchDepositMsgStateIndexKey(poolID, msgIndex uint64) []byte {
	key := make([]byte, 17)
	key[0] = PoolBatchDepositMsgStateIndexKeyPrefix[0]
	copy(key[1:9], sdk.Uint64ToBigEndian(poolID))
	copy(key[9:17], sdk.Uint64ToBigEndian(msgIndex))
	return key
}

// GetPoolBatchWithdrawMsgStateIndexKey returns kv indexing key of the latest index value of the msg index
func GetPoolBatchWithdrawMsgStateIndexKey(poolID, msgIndex uint64) []byte {
	key := make([]byte, 17)
	key[0] = PoolBatchWithdrawMsgStateIndexKeyPrefix[0]
	copy(key[1:9], sdk.Uint64ToBigEndian(poolID))
	copy(key[9:17], sdk.Uint64ToBigEndian(msgIndex))
	return key
}

// GetPoolBatchSwapMsgStateIndexKey returns kv indexing key of the latest index value of the msg index
func GetPoolBatchSwapMsgStateIndexKey(poolID, msgIndex uint64) []byte {
	key := make([]byte, 17)
	key[0] = PoolBatchSwapMsgStateIndexKeyPrefix[0]
	copy(key[1:9], sdk.Uint64ToBigEndian(poolID))
	copy(key[9:17], sdk.Uint64ToBigEndian(msgIndex))
	return key
}
