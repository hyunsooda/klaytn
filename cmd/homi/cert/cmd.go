package cert

import (
	"errors"

	"gopkg.in/urfave/cli.v1"
)

var (
	CertCommand = cli.Command{
		Name:        "cert",
		Usage:       "Certificate generator",
		Description: "Cert command generate private key and certificate",
		Action:      gen,
		Flags: []cli.Flag{
			genRootCertFlag,
			genClientCertFlag,
			genCANameFlag,
			genOrgNameFlag,
		},
	}
	ErrEmptyCAName  = errors.New("CA name is not given")
	ErrEmptyOrgName = errors.New("Organization name is not given")
)

func getCANameFlag(ctx *cli.Context) (string, error) {
	caName := ctx.String(genCANameFlag.Name)
	if len(caName) == 0 {
		return "", ErrEmptyCAName
	}
	return caName, nil
}

func getOrgNameFlag(ctx *cli.Context) (string, error) {
	caName := ctx.String(genOrgNameFlag.Name)
	if len(caName) == 0 {
		return "", ErrEmptyOrgName
	}
	return caName, nil
}

func gen(ctx *cli.Context) error {
	if ctx.Bool(genRootCertFlag.Name) {
		caName, err := getCANameFlag(ctx)
		if err != nil {
			return err
		}
		return GenRootCert(caName)
	} else if ctx.Bool(genClientCertFlag.Name) {
		caName, err := getOrgNameFlag(ctx)
		if err != nil {
			return err
		}
		return GenClientCert(caName)
	}
	return nil
}
