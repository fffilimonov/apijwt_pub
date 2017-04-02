package hyper

import (
	"path/filepath"

	"github.com/hyperhq/hypercli/cli"
	"github.com/hyperhq/hypercli/cliconfig"
	flag "github.com/hyperhq/hypercli/pkg/mflag"
)

var clientFlags = &cli.ClientFlags{FlagSet: new(flag.FlagSet), Common: commonFlags}

func init() {
	clientFlags.PostParse = func() {
		clientFlags.Common.PostParse()
		clientFlags.Common.TrustKey = filepath.Join(cliconfig.ConfigDir(), defaultTrustKeyFile)
	}
}
