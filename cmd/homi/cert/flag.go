package cert

import "github.com/urfave/cli"

var (
	genRootCertFlag = cli.BoolFlag{
		Name:  "genrootcert",
		Usage: "generate root private key and certificate",
	}
	genClientCertFlag = cli.BoolFlag{
		Name:  "genclientcert",
		Usage: "generate client private key and certificate",
	}
	genCANameFlag = cli.StringFlag{
		Name:  "caname",
		Usage: "CA name in certificate to be generated",
	}
	genOrgNameFlag = cli.StringFlag{
		Name:  "orgname",
		Usage: "organization name in certificate to be generated",
	}
)
