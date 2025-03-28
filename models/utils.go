/*
 * Copyright (C) 2020 The poly network Authors
 * This file is part of The poly network library.
 *
 * The  poly network  is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The  poly network  is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 * You should have received a copy of the GNU Lesser General Public License
 * along with The poly network .  If not, see <http://www.gnu.org/licenses/>.
 */

package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"poly-bridge/basedef"
	"poly-bridge/utils/decimal"
)

type Request struct {
	JsonRpc string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	Id      uint        `json:"id"`
}

type RPCError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type Response struct {
	Error  *RPCError       `json:"error"`
	ID     int             `json:"id"`
	Result json.RawMessage `json:"result,omitempty"`
}

// TxnReceipt is the receipt obtained over JSON/RPC from the ethereum client
type TxnReceipt struct {
	BlockHash         *common.Hash    `json:"blockHash"`
	BlockNumber       *hexutil.Big    `json:"blockNumber"`
	ContractAddress   *common.Address `json:"contractAddress"`
	CumulativeGasUsed *hexutil.Big    `json:"cumulativeGasUsed"`
	GasUsed           *hexutil.Big    `json:"gasUsed"`
	TransactionHash   *common.Hash    `json:"transactionHash"`
	TransactionIndex  *hexutil.Uint   `json:"transactionIndex"`
	From              *common.Address `json:"from"`
	To                *common.Address `json:"to"`
	Status            *hexutil.Big    `json:"status"`
	L1BlockNumber     *hexutil.Big    `json:"l1BlockNumber"`
}

type chainCache struct {
	ChainLogo        string
	ChainExplorerUrl string
	ChainFeeName     string
	ChainFeeLogo     string
}

var chainsNamesCache = map[uint64]string{}
var chainsCache = map[uint64]chainCache{}

func Init(chains []*Chain, chainFees []*ChainFee) {
	for _, chain := range chains {
		chainsNamesCache[chain.ChainId] = chain.Name
		for _, chainFee := range chainFees {
			if chain.ChainId == chainFee.ChainId {
				chainsCache[chain.ChainId] = chainCache{
					chain.ChainLogo,
					chain.ChainExplorerUrl,
					chainFee.TokenBasicName,
					chainFee.TokenBasic.Meta,
				}
			}
		}
		if _, ok := chainsCache[chain.ChainId]; !ok {
			chainsCache[chain.ChainId] = chainCache{
				chain.ChainLogo,
				chain.ChainExplorerUrl,
				"",
				"",
			}
		}
	}
}

func ChainId2Name(id uint64) string {
	name, ok := chainsNamesCache[id]
	if ok {
		return name
	}
	return fmt.Sprintf("%v", id)
}

func ChainId2ChainCache(id uint64) chainCache {
	cache, ok := chainsCache[id]
	if ok {
		return cache
	}
	return chainCache{}
}

type BigInt struct {
	big.Int
}

func NewBigIntFromInt(value int64) *BigInt {
	x := new(big.Int).SetInt64(value)
	return NewBigInt(x)
}

func NewBigInt(value *big.Int) *BigInt {
	return &BigInt{Int: *value}
}

func (bigInt *BigInt) Value() (driver.Value, error) {
	if bigInt == nil {
		return "null", nil
	}
	return bigInt.String(), nil
}

func (bigInt *BigInt) Scan(v interface{}) error {
	value, ok := v.([]byte)
	if !ok {
		return fmt.Errorf("type error, %v", v)
	}
	str := string(value)
	if str == "null" || str == "nil" || str == "<nil>" || str == "" {
		return nil
	}
	data, ok := new(big.Int).SetString(str, 10)
	if !ok {
		return fmt.Errorf("not a valid big integer: %s", value)
	}
	bigInt.Int = *data
	return nil
}

func FormatAmount(precision uint64, amount *BigInt) string {
	precision_new := decimal.NewFromBigInt(big.NewInt(1), int32(precision))
	amount_new := decimal.NewFromBigInt(&amount.Int, 0)
	return amount_new.Div(precision_new).String()
}

func FeePrecison(chain uint64) int {
	switch chain {
	case basedef.NEO3_CROSSCHAIN_ID:
		return 8
	default:
		return 18
	}
}

func FormatFee(chain uint64, fee *BigInt) string {
	fee_new := decimal.NewFromBigInt(&fee.Int, 0)

	switch chain {
	case basedef.ONTEVM_CROSSCHAIN_ID:
		precision_new := decimal.New(1, 18)
		return fee_new.Div(precision_new).String() + " ONG(evm)"

	case basedef.ETHEREUM_CROSSCHAIN_ID:
		precision_new := decimal.New(1, 18)
		return fee_new.Div(precision_new).String() + " ETH"

	case basedef.BSC_CROSSCHAIN_ID:
		precision_new := decimal.New(1, 18)
		return fee_new.Div(precision_new).String() + " BNB"

	case basedef.NEO3_CROSSCHAIN_ID:
		precision_new := decimal.New(1, 8)
		return fee_new.Div(precision_new).String() + " GAS"

	default:
		precision_new := decimal.New(int64(1), 0)
		return fee_new.Div(precision_new).String()

	}
}

func TxType2Name(ty uint32) string {
	return "cross chain transfer"
}
func Precent(a uint64, b uint64) string {
	c := float64(a) / float64(b)
	return fmt.Sprintf("%.2f%%", c*100)
}

func NullToZero(a **BigInt) {
	if *a == nil {
		*a = NewBigInt(new(big.Int).SetInt64(0))
	}
}

func FormatString(data string) string {
	if len(data) > 64 {
		return data[:64]
	}
	return data
}

func FormatAssert(data string) string {
	if len(data) > 120 {
		return data[:120]
	}
	return data
}

func Format8190(data string) string {
	if len(data) > 8190 {
		return data[:8190]
	}
	return data
}

func GetTokenType(chainId uint64, standard uint8) string {
	tokenType := ""
	switch standard {
	case TokenTypeErc20:
		tokenType = "20"
	case TokenTypeErc721:
		tokenType = "721"
	default:
		tokenType = "20"
	}
	switch chainId {
	case basedef.ETHEREUM_CROSSCHAIN_ID:
		return "ERC" + "-" + tokenType

	case basedef.BSC_CROSSCHAIN_ID:
		return "BEP" + "-" + tokenType

	case basedef.NEO3_CROSSCHAIN_ID:
		if standard == TokenTypeErc721 {
			return "NFT"
		}
		return "NEP17"

	default:
		return "ERC" + "-" + tokenType
	}
}
