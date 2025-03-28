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

package basedef

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	"github.com/ethereum/go-ethereum/common"
	"github.com/joeqian10/neo3-gogogo/crypto"
	"github.com/joeqian10/neo3-gogogo/helper"
)

func ReadFile(fileName string) ([]byte, error) {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0666)
	if err != nil {
		return nil, fmt.Errorf("ReadFile: open file %s error %s", fileName, err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			logs.Error("ReadFile: File %s close error %s", fileName, err)
		}
	}()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("ReadFile: ioutil.ReadAll %s error %s", fileName, err)
	}
	return data, nil
}

func Hash2Address(chainId uint64, value string) string {
	if chainId == ETHEREUM_CROSSCHAIN_ID {
		addr := common.HexToAddress(value)
		return strings.ToLower(addr.String()[2:])
	} else if chainId == BSC_CROSSCHAIN_ID {
		addr := common.HexToAddress(value)
		return strings.ToLower(addr.String()[2:])
	} else if chainId == ONTEVM_CROSSCHAIN_ID {
		addr := common.HexToAddress(value)
		return strings.ToLower(addr.String()[2:])
	} else if chainId == NEO3_CROSSCHAIN_ID {
		addrHex, _ := hex.DecodeString(value)
		addr := helper.UInt160FromBytes(addrHex)
		address := crypto.ScriptHashToAddress(addr, helper.DefaultAddressVersion)
		return address
	}
	return value
}

func HexReverse(arr []byte) []byte {
	l := len(arr)
	x := make([]byte, 0)
	for i := l - 1; i >= 0; i-- {
		x = append(x, arr[i])
	}
	return x
}

func HexStringReverse(value string) string {
	aa, _ := hex.DecodeString(value)
	bb := HexReverse(aa)
	return hex.EncodeToString(bb)
}

func String2Float64(value string) float64 {
	v, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0
	}
	return v
}

func Int64FromFigure(figure int) int64 {
	x := int64(1)
	for i := 0; i < figure; i++ {
		x *= 10
	}
	return x
}

func Address2Hash(chainId uint64, value string) (string, error) {
	if chainId == ETHEREUM_CROSSCHAIN_ID {
		addr := common.HexToAddress(value)
		return strings.ToLower(addr.String()[2:]), nil
	} else if chainId == BSC_CROSSCHAIN_ID {
		addr := common.HexToAddress(value)
		return strings.ToLower(addr.String()[2:]), nil
	} else if chainId == ONTEVM_CROSSCHAIN_ID {
		addr := common.HexToAddress(value)
		return strings.ToLower(addr.String()[2:]), nil
	} else if chainId == NEO3_CROSSCHAIN_ID {
		scriptHash, err := crypto.AddressToScriptHash(value, helper.DefaultAddressVersion)
		if err != nil {
			return value, err
		}
		addrBytes := scriptHash.ToByteArray()
		address := hex.EncodeToString(addrBytes)
		return address, nil
	}
	return value, nil
}

// lock item_proxy use
func Proxy2Address(chainId uint64, proxy string) string {
	if chainId == NEO3_CROSSCHAIN_ID {
		proxy = HexStringReverse(proxy)
	}
	return Hash2Address(chainId, proxy)
}

func ConfirmEnv(env string) {
	if ENV != env {
		logs.Error("Config env(%s) does not match build env(%s)", env, ENV)
		os.Exit(1)
	}
	logs.Info("Current env: %s", ENV)
}

func GetChainName(id uint64) string {
	switch id {
	case POLY_CROSSCHAIN_ID:
		return "Poly"
	case ETHEREUM_CROSSCHAIN_ID:
		return "Ethereum"
	case BSC_CROSSCHAIN_ID:
		return "Bsc"
	case NEO3_CROSSCHAIN_ID:
		return "Neo3"
	case ONTEVM_CROSSCHAIN_ID:
		return "OntEVM"
	default:
		return fmt.Sprintf("Unknown(%d)", id)
	}
}

func FormatAddr(chain uint64, addr string) string {
	switch chain {
	case NEO3_CROSSCHAIN_ID:
		return addr
	default:
		return common.HexToAddress(addr).String()
	}
}

func FormatTxHash(chain uint64, hash string) string {
	switch chain {
	default:
		return common.HexToHash(hash).String()
	}
}

// Has0xPrefix validates str begins with '0x' or '0X'.
func Has0xPrefix(str string) bool {
	return len(str) >= 2 && str[0] == '0' && (str[1] == 'x' || str[1] == 'X')
}

func IsETHChain(chainId uint64) bool {
	for _, v := range ETH_CHAINS {
		if chainId == v {
			return true
		}
	}
	return false
}
