package config

type PostProcess interface {
	PostProcess() error
}
