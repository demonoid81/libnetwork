package libnetwork

import (
	"github.com/demonoid81/libnetwork/drivers/bridge"
	"github.com/demonoid81/libnetwork/drivers/host"
	"github.com/demonoid81/libnetwork/drivers/ipvlan"
	"github.com/demonoid81/libnetwork/drivers/macvlan"
	"github.com/demonoid81/libnetwork/drivers/null"
	"github.com/demonoid81/libnetwork/drivers/overlay"
	"github.com/demonoid81/libnetwork/drivers/remote"
)

func getInitializers(experimental bool) []initializer {
	in := []initializer{
		{bridge.Init, "bridge"},
		{host.Init, "host"},
		{ipvlan.Init, "ipvlan"},
		{macvlan.Init, "macvlan"},
		{null.Init, "null"},
		{overlay.Init, "overlay"},
		{remote.Init, "remote"},
	}
	return in
}
