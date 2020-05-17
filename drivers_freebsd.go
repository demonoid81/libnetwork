package libnetwork

import (
	"github.com/demonoid81/libnetwork/drivers/null"
	"github.com/demonoid81/libnetwork/drivers/remote"
)

func getInitializers(experimental bool) []initializer {
	return []initializer{
		{null.Init, "null"},
		{remote.Init, "remote"},
	}
}
