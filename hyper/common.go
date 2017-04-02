package hyper

import (
	"github.com/docker/go-connections/tlsconfig"
	"github.com/hyperhq/hypercli/cli"
	"github.com/hyperhq/hypercli/cliconfig"
	flag "github.com/hyperhq/hypercli/pkg/mflag"
)

const (
	defaultTrustKeyFile = "key.json"
	defaultCaFile       = "ca.pem"
	defaultKeyFile      = "key.pem"
	defaultCertFile     = "cert.pem"
	tlsVerifyKey        = "tlsverify"
)

var (
	commonFlags     = &cli.CommonFlags{FlagSet: new(flag.FlagSet)}
	dockerCertPath  = ""
	dockerTLSVerify = false
)

func init() {
	dockerCertPath = cliconfig.ConfigDir()
	commonFlags.PostParse = postParseCommon

	var tlsOptions tlsconfig.Options
	commonFlags.TLSOptions = &tlsOptions
}

func postParseCommon() {
	commonFlags.TLS = true
	tlsOptions := commonFlags.TLSOptions
	tlsOptions.InsecureSkipVerify = true
}
