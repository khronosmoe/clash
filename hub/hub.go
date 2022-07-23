package hub

import (
	"github.com/Dreamacro/clash/config"
	"github.com/Dreamacro/clash/hub/executor"
	"github.com/Dreamacro/clash/hub/route"
)

type Option func(*config.Config)

func WithExternalUI(externalUI string) Option {
	return func(cfg *config.Config) {
		cfg.General.ExternalUI = externalUI
	}
}

func WithExternalController(externalController string) Option {
	return func(cfg *config.Config) {
		cfg.General.ExternalController = externalController
	}
}

func WithSecret(secret string) Option {
	return func(cfg *config.Config) {
		cfg.General.Secret = secret
	}
}

func WithExternalTLS(externalTLS bool) Option {
	return func(cfg *config.Config) {
		cfg.General.ExternalTLS = externalTLS
	}
}

func WithTLSCertPath(tLSCertPath string) Option {
	return func(cfg *config.Config) {
		cfg.General.TLSCertPath = tLSCertPath
	}
}

func WithTLSKeyPath(tLSKeyPath string) Option {
	return func(cfg *config.Config) {
		cfg.General.TLSKeyPath = tLSKeyPath
	}
}

// Parse call at the beginning of clash
func Parse(options ...Option) error {
	cfg, err := executor.Parse()
	if err != nil {
		return err
	}

	for _, option := range options {
		option(cfg)
	}

	if (cfg.General.ExternalTLS) {
		route.SetTLS(cfg.General.TLSCertPath,cfg.General.TLSKeyPath)
	}

	if cfg.General.ExternalUI != "" {
		route.SetUIPath(cfg.General.ExternalUI)
	}

	if cfg.General.ExternalController != "" {
		go route.Start(cfg.General.ExternalController, cfg.General.Secret)
	}

	executor.ApplyConfig(cfg, true)
	return nil
}
