package common

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"

	"poly-bridge/basedef"
	"poly-bridge/chainsdk"
	"poly-bridge/conf"
)

var (
	ethereumSdk *chainsdk.EthereumSdkPro

	bscSdk *chainsdk.EthereumSdkPro

	neo3Sdk *chainsdk.Neo3SdkPro

	ontevmSdk *chainsdk.EthereumSdkPro

	sdkMap map[uint64]interface{}
	config *conf.Config
)

func GetSdk(chainId uint64) interface{} {
	return sdkMap[chainId]
}

func SetupChainsSDK(cfg *conf.Config) {
	if cfg == nil {
		panic("Missing config")
	}
	config = cfg
	newChainSdks(cfg)
}

func newChainSdks(config *conf.Config) {
	sdkMap = make(map[uint64]interface{}, 0)
	{
		ethereumConfig := config.GetChainListenConfig(basedef.ETHEREUM_CROSSCHAIN_ID)
		if ethereumConfig == nil {
			panic("chain is invalid")
		}
		urls := ethereumConfig.GetNodesUrl()
		ethereumSdk = chainsdk.NewEthereumSdkPro(urls, ethereumConfig.ListenSlot, ethereumConfig.ChainId)
		sdkMap[basedef.ETHEREUM_CROSSCHAIN_ID] = ethereumSdk
	}
	{
		bscConfig := config.GetChainListenConfig(basedef.BSC_CROSSCHAIN_ID)
		if bscConfig == nil {
			panic("chain is invalid")
		}
		urls := bscConfig.GetNodesUrl()
		bscSdk = chainsdk.NewEthereumSdkPro(urls, bscConfig.ListenSlot, bscConfig.ChainId)
		sdkMap[basedef.BSC_CROSSCHAIN_ID] = bscSdk
	}
	{
		neo3Config := config.GetChainListenConfig(basedef.NEO3_CROSSCHAIN_ID)
		if neo3Config == nil {
			panic("chain is invalid")
		}
		urls := neo3Config.GetNodesUrl()
		neo3Sdk = chainsdk.NewNeo3SdkPro(urls, neo3Config.ListenSlot, neo3Config.ChainId)
		sdkMap[basedef.NEO3_CROSSCHAIN_ID] = neo3Sdk
	}
	{
		ontevmConfig := config.GetChainListenConfig(basedef.ONTEVM_CROSSCHAIN_ID)
		if ontevmConfig == nil {
			panic("ont evm is invalid")
		}
		urls := ontevmConfig.GetNodesUrl()
		ontevmSdk = chainsdk.NewEthereumSdkPro(urls, ontevmConfig.ListenSlot, ontevmConfig.ChainId)
		sdkMap[basedef.ONTEVM_CROSSCHAIN_ID] = ontevmSdk
	}
}

func GetBalance(chainId uint64, hash string) (*big.Int, error) {
	maxBalance := big.NewInt(0)
	maxFun := func(balance *big.Int) {
		if balance != nil && balance.Cmp(maxBalance) > 0 {
			maxBalance = balance
		}
	}
	errMap := make(map[error]bool, 0)
	if chainId == basedef.ETHEREUM_CROSSCHAIN_ID {
		ethereumConfig := config.GetChainListenConfig(basedef.ETHEREUM_CROSSCHAIN_ID)
		if ethereumConfig == nil {
			panic("chain is invalid")
		}
		for _, v := range ethereumConfig.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := ethereumSdk.Erc20Balance(hash, v)
			maxFun(balance)
			errMap[err] = true
		}
	}
	if chainId == basedef.BSC_CROSSCHAIN_ID {
		bscConfig := config.GetChainListenConfig(basedef.BSC_CROSSCHAIN_ID)
		if bscConfig == nil {
			panic("chain is invalid")
		}
		for _, v := range bscConfig.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := bscSdk.Erc20Balance(hash, v)
			maxFun(balance)
			errMap[err] = true
		}
	}
	if chainId == basedef.NEO3_CROSSCHAIN_ID {
		neo3Config := config.GetChainListenConfig(basedef.NEO3_CROSSCHAIN_ID)
		if neo3Config == nil {
			panic("chain is invalid")
		}
		for _, v := range neo3Config.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := neo3Sdk.Nep17Balance(hash, v)
			maxFun(balance)
			errMap[err] = true
		}
	}
	if chainId == basedef.ONTEVM_CROSSCHAIN_ID {
		ontevmConfig := config.GetChainListenConfig(basedef.ONTEVM_CROSSCHAIN_ID)
		if ontevmConfig == nil {
			panic("ontevm is invalid")
		}
		for _, v := range ontevmConfig.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := ontevmSdk.Erc20Balance(hash, v)
			maxFun(balance)
			errMap[err] = true
		}
	}
	if maxBalance.Cmp(big.NewInt(0)) > 0 {
		return maxBalance, nil
	}
	var err error
	for k, _ := range errMap {
		if k == nil {
			return new(big.Int).SetUint64(0), nil
		} else {
			err = k
		}
	}
	return new(big.Int).SetUint64(0), err
}

func GetTotalSupply(chainId uint64, hash string) (*big.Int, error) {
	if chainId == basedef.ETHEREUM_CROSSCHAIN_ID {
		ethereumConfig := config.GetChainListenConfig(basedef.ETHEREUM_CROSSCHAIN_ID)
		if ethereumConfig == nil {
			panic("chain is invalid")
		}
		return ethereumSdk.Erc20TotalSupply(hash)
	}
	if chainId == basedef.BSC_CROSSCHAIN_ID {
		bscConfig := config.GetChainListenConfig(basedef.BSC_CROSSCHAIN_ID)
		if bscConfig == nil {
			panic("chain is invalid")
		}
		return bscSdk.Erc20TotalSupply(hash)
	}

	if chainId == basedef.NEO3_CROSSCHAIN_ID {
		neo3Config := config.GetChainListenConfig(basedef.NEO3_CROSSCHAIN_ID)
		if neo3Config == nil {
			panic("chain is invalid")
		}
		return neo3Sdk.Nep17TotalSupply(hash)
	}

	if chainId == basedef.ONTEVM_CROSSCHAIN_ID {
		ontevmConfig := config.GetChainListenConfig(basedef.ONTEVM_CROSSCHAIN_ID)
		if ontevmConfig == nil {
			panic("ontevm is invalid")
		}
		return ontevmSdk.Erc20TotalSupply(hash)
	}
	return new(big.Int).SetUint64(0), nil
}

type ProxyBalance struct {
	Amount    *big.Int
	ItemName  string
	ItemProxy string
}

func GetProxyBalance(chainId uint64, hash string, proxy string) (*big.Int, error) {
	switch chainId {
	case basedef.ETHEREUM_CROSSCHAIN_ID:
		return ethereumSdk.Erc20Balance(hash, proxy)

	case basedef.BSC_CROSSCHAIN_ID:
		return bscSdk.Erc20Balance(hash, proxy)

	case basedef.NEO3_CROSSCHAIN_ID:
		return neo3Sdk.Nep17Balance(hash, proxy)

	case basedef.ONTEVM_CROSSCHAIN_ID:
		return ontevmSdk.Erc20Balance(hash, proxy)

	default:
		return new(big.Int).SetUint64(0), nil
	}
}

func GetNftOwner(chainId uint64, asset string, tokenId int) (owner common.Address, err error) {
	switch chainId {
	case basedef.ETHEREUM_CROSSCHAIN_ID:
		return ethereumSdk.GetNFTOwner(asset, big.NewInt(int64(tokenId)))

	case basedef.BSC_CROSSCHAIN_ID:
		return bscSdk.GetNFTOwner(asset, big.NewInt(int64(tokenId)))

	default:
		return common.Address{}, fmt.Errorf("has nat func with chain:%v", chainId)
	}
}

func GetBoundLockProxy(lockProxies []string, srcTokenHash, DstTokenHash string, srcChainId, dstChainId uint64) (string, error) {
	if sdk, exist := sdkMap[dstChainId]; exist {
		if value, ok := sdk.(*chainsdk.EthereumSdkPro); ok {
			return value.GetBoundLockProxy(lockProxies, srcTokenHash, DstTokenHash, srcChainId)
		}
	}
	return "", fmt.Errorf("chain %d is not ethereum based", dstChainId)
}
