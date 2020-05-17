package api

import (
	"github.com/demonoid81/libnetwork/drivers/bridge"
	"github.com/demonoid81/libnetwork/netlabel"
)

func GetOpsMap(bridgeName, defaultMTU string) map[string]string {
	if defaultMTU == "" {
		return map[string]string{
			bridge.BridgeName: bridgeName,
		}
	}
	return map[string]string{
		bridge.BridgeName:  bridgeName,
		netlabel.DriverMTU: defaultMTU,
	}
}
