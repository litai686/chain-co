module github.com/ipfs-force-community/chain-co

go 1.16

require (
	github.com/dtynn/dix v0.1.0
	github.com/filecoin-project/go-address v0.0.6
	github.com/filecoin-project/go-bitfield v0.2.4
	github.com/filecoin-project/go-data-transfer v1.15.1
	github.com/filecoin-project/go-fil-markets v1.20.1-v16-2
	github.com/filecoin-project/go-jsonrpc v0.1.5
	github.com/filecoin-project/go-state-types v0.1.10
	github.com/filecoin-project/lotus v1.14.0-rc1
	github.com/filecoin-project/venus-auth v1.6.0
	github.com/gbrlsnchs/jwt/v3 v3.0.1
	github.com/google/uuid v1.3.0
	github.com/hashicorp/go-multierror v1.1.1
	github.com/hashicorp/golang-lru v0.5.4
	github.com/ipfs-force-community/metrics v1.0.0
	github.com/ipfs-force-community/venus-common-utils v0.0.0-20210714031758-ea0e25ff0ec4
	github.com/ipfs/go-block-format v0.0.3
	github.com/ipfs/go-cid v0.1.0
	github.com/ipfs/go-log/v2 v2.5.1
	github.com/ipfs/go-metrics-interface v0.0.1
	github.com/libp2p/go-libp2p-core v0.15.1
	github.com/urfave/cli/v2 v2.3.0
	github.com/whyrusleeping/pubsub v0.0.0-20190708150250-92bcb0691325
	go.opencensus.io v0.23.0
	go.uber.org/fx v1.15.0
	go.uber.org/zap v1.21.0
)

replace github.com/filecoin-project/filecoin-ffi => ./extern/filecoin-ffi

replace github.com/filecoin-project/go-jsonrpc => github.com/ipfs-force-community/go-jsonrpc v0.1.4-0.20211201033628-fc1430d095f6

replace github.com/filecoin-project/lotus => github.com/ipfs-force-community/lotus v0.8.1-0.20220711022050-878b48b822b5
