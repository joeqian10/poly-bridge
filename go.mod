module poly-bridge

go 1.14

require (
	github.com/beego/beego/v2 v2.0.1
	github.com/btcsuite/btcd v0.22.0-beta
	github.com/btcsuite/goleveldb v1.0.0
	github.com/cosmos/cosmos-sdk v0.39.1
	github.com/devfans/cogroup v1.1.0
	github.com/ethereum/go-ethereum v1.9.25
	github.com/go-redis/redis v6.14.2+incompatible
	github.com/go-telegram-bot-api/telegram-bot-api/v5 v5.0.0-20221020003552-4126fa611266
	github.com/hashicorp/golang-lru v0.5.4
	github.com/howeyc/gopass v0.0.0-20210920133722-c8aef6fb66ef
	github.com/joeqian10/neo3-gogogo v1.1.2
	github.com/ontio/ontology v1.17.0
	github.com/ontio/ontology-crypto v1.2.1
	github.com/ontio/ontology-go-sdk v1.12.4
	github.com/polynetwork/bridge-common v0.0.0-20210730081758-77b97bbeb305
	github.com/polynetwork/poly v1.3.1
	github.com/polynetwork/poly-go-sdk v0.0.0-20210114035303-84e1615f4ad4
	github.com/stretchr/testify v1.8.2
	github.com/urfave/cli v1.22.4
	gorm.io/driver/mysql v1.0.3
	gorm.io/gorm v1.20.8
)

replace (
	github.com/block-vision/sui-go-sdk => github.com/yandc/sui-go-sdk v1.0.5
	github.com/cosmos/cosmos-sdk => github.com/Switcheo/cosmos-sdk v0.39.2-0.20200814061308-474a0dbbe4ba
)
