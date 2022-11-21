package cmd

import (
	"github.com/johnreitano/numi/app"
)

func initSDKConfig() {
	app.SetAddressPrefixesInSDKConfig().Seal()
}
