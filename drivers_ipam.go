package libnetwork

import (
	"github.com/demonoid81/libnetwork/drvregistry"
	"github.com/demonoid81/libnetwork/ipamapi"
	builtinIpam "github.com/demonoid81/libnetwork/ipams/builtin"
	nullIpam "github.com/demonoid81/libnetwork/ipams/null"
	remoteIpam "github.com/demonoid81/libnetwork/ipams/remote"
	"github.com/demonoid81/libnetwork/ipamutils"
)

func initIPAMDrivers(r *drvregistry.DrvRegistry, lDs, gDs interface{}, addressPool []*ipamutils.NetworkToSplit) error {
	builtinIpam.SetDefaultIPAddressPool(addressPool)
	for _, fn := range [](func(ipamapi.Callback, interface{}, interface{}) error){
		builtinIpam.Init,
		remoteIpam.Init,
		nullIpam.Init,
	} {
		if err := fn(r, lDs, gDs); err != nil {
			return err
		}
	}

	return nil
}
