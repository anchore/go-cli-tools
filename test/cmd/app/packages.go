package app

import (
	"github.com/anchore/go-cli-tools/inject"
	"github.com/anchore/go-cli-tools/log"
)

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

func Packages(_ inject.Container, cc *CatalogerConfig, gc *GoCatalogerConfig, input string) error {
	log.Tracef("packages executed with input: %v enabled: %v search: %v", input, cc.Enabled, gc.NetworkLicenseSearch)
	return nil
}
