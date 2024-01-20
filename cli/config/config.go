package config

import (
	"net"

	"github.com/go-ini/ini"
)

type Config struct {
	APIURL    string `ini:"api_url"`
	Subnet    string `ini:"subnet"`
	NewPrefix int    `ini:"new_prefix"`
}

func (c *Config) GetAPIURL() string {
	return c.APIURL
}

func (c *Config) GetNewPrefix() int {
	return c.NewPrefix
}

func (c *Config) GetSubnet() *net.IPNet {
	_, network, err := net.ParseCIDR(c.Subnet)
	if err != nil {
		return nil
	}

	return network
}

func Parse(filePath string) (*Config, error) {
	cfg, err := ini.Load(filePath)
	if err != nil {
		return nil, err
	}

	config := new(Config)

	err = cfg.MapTo(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
