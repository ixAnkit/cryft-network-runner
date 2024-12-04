package api

import (
	"github.com/cryft-labs/coreth/plugin/evm"
	"github.com/ixAnkit/cryftgo/api/admin"
	"github.com/ixAnkit/cryftgo/api/health"
	"github.com/ixAnkit/cryftgo/api/info"
	"github.com/ixAnkit/cryftgo/api/ipcs"
	"github.com/ixAnkit/cryftgo/api/keystore"
	"github.com/ixAnkit/cryftgo/indexer"
	"github.com/ixAnkit/cryftgo/vms/avm"
	"github.com/ixAnkit/cryftgo/vms/platformvm"
)

// Issues API calls to a node
// TODO: byzantine api. check if appropriate. improve implementation.
type Client interface {
	PChainAPI() platformvm.Client
	XChainAPI() avm.Client
	XChainWalletAPI() avm.WalletClient
	CChainAPI() evm.Client
	CChainEthAPI() EthClient // ethclient websocket wrapper that adds mutexed calls, and lazy conn init (on first call)
	InfoAPI() info.Client
	HealthAPI() health.Client
	IpcsAPI() ipcs.Client
	KeystoreAPI() keystore.Client
	AdminAPI() admin.Client
	PChainIndexAPI() indexer.Client
	CChainIndexAPI() indexer.Client
	// TODO add methods
}
