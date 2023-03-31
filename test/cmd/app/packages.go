package app

import (
	"github.com/anchore/go-cli-tools/inject"
	"github.com/anchore/go-cli-tools/log"
)

type PackageBasicConfig struct {
	InputSource string
}

type CatalogerConfig struct {
	Enabled []string
}

func DefaultCatalogerConfig() CatalogerConfig {
	return CatalogerConfig{
		Enabled: []string{
			"one",
			"two",
		},
	}
}

type GoCatalogerConfig struct {
	NetworkLicenseSearch bool
}

func Packages(c inject.Container, bc *PackageBasicConfig, cc *CatalogerConfig, gc *GoCatalogerConfig) error {
	log.Tracef("packages executed with: %v %v %v", bc.InputSource, cc.Enabled, gc.NetworkLicenseSearch)
	return nil
}
