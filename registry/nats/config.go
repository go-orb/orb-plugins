package nats

import (
	"crypto/tls"
	"fmt"

	"go-micro.dev/v5/config"
	"go-micro.dev/v5/config/source/cli"
	"go-micro.dev/v5/registry"
	"go-micro.dev/v5/types"
)

const name = "nats"

// Defaults.
//
//nolint:gochecknoglobals
var (
	DefaultAddresses  = []string{"nats://localhost:4222"}
	DefaultQueryTopic = "micro.registry.nats.query"
	DefaultWatchTopic = "micro.registry.nats.watch"
)

func init() {
	//nolint:errcheck
	_ = cli.Flags.Add(cli.NewFlag(
		"registry_addresses",
		DefaultAddresses,
		cli.ConfigPathSlice([]string{"registry", "addresses"}),
		cli.Usage("Registry addresses."),
	))

	if err := registry.Plugins.Add(name, registry.ProviderFunc(ProvideRegistryNATS)); err != nil {
		panic(err)
	}
}

// Config provides configuration for the MDNS registry.
type Config struct {
	registry.Config `yaml:",inline"`

	Addresses []string    `json:"addresses,omitempty" yaml:"addresses,omitempty"`
	Secure    bool        `json:"secure,omitempty" yaml:"secure,omitempty"`
	TLSConfig *tls.Config `json:"-" yaml:"-"`

	Quorum int `json:"quorum,omitempty" yaml:"quorum,omitempty"`

	QueryTopic string `json:"queryTopic,omitempty" yaml:"queryTopic,omitempty"`
	WatchTopic string `json:"watchTopic,omitempty" yaml:"watchTopic,omitempty"`
}

// NewConfig creates a new config object.
func NewConfig(
	serviceName types.ServiceName,
	datas types.ConfigData,
	opts ...registry.Option,
) (Config, error) {
	cfg := Config{
		Config: registry.NewConfig(),
	}

	cfg.ApplyOptions(opts...)

	sections := types.SplitServiceName(serviceName)
	if err := config.Parse(append(sections, registry.ComponentType), datas, &cfg); err != nil {
		return cfg, fmt.Errorf("parse config: %w", err)
	}

	return cfg, nil
}

// ApplyOptions applies a set of options to the config.
func (c *Config) ApplyOptions(opts ...registry.Option) {
	for _, o := range opts {
		o(c)
	}
}

// WithAddress sets the NATS server addresses.
func WithAddress(n ...string) registry.Option {
	return func(c registry.ConfigType) {
		cfg, ok := c.(*Config)
		if ok {
			cfg.Addresses = n
		} else {
			panic(fmt.Sprintf("wrong type: %T", c))
		}
	}
}

// WithSecure defines if we want a secure connection to nats.
func WithSecure(n bool) registry.Option {
	return func(c registry.ConfigType) {
		cfg, ok := c.(*Config)
		if ok {
			cfg.Secure = n
		}
	}
}

// WithTLSConfig defines the TLS config to use for the secure connection.
func WithTLSConfig(n *tls.Config) registry.Option {
	return func(c registry.ConfigType) {
		cfg, ok := c.(*Config)
		if ok {
			cfg.TLSConfig = n
		}
	}
}

// WithQuorum sets the NATS quorum.
func WithQuorum(n int) registry.Option {
	return func(c registry.ConfigType) {
		cfg, ok := c.(*Config)
		if ok {
			cfg.Quorum = n
		}
	}
}

// WithQueryTopic sets the NATS query topic.
func WithQueryTopic(n string) registry.Option {
	return func(c registry.ConfigType) {
		cfg, ok := c.(*Config)
		if ok {
			cfg.QueryTopic = n
		}
	}
}

// WithWatchTopic sets the NATS watch topic.
func WithWatchTopic(n string) registry.Option {
	return func(c registry.ConfigType) {
		cfg, ok := c.(*Config)
		if ok {
			cfg.WatchTopic = n
		}
	}
}
