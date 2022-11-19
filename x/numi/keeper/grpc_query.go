package keeper

import (
	"github.com/johnreitano/numi/x/numi/types"
)

var _ types.QueryServer = Keeper{}
