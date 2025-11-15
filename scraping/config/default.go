package config

import (
	"encoding/json"
	"fmt"
)

var DefaultConfig ScrapersConfig = ScrapersConfig{
	Scrapers: []ScraperCfg{
		NewRozamrynyCfg(),
		NewOlomoucCfg("M.D. Original 1869", "MD-Original-1869-id2208"),
		NewOlomoucCfg("Bistro Paulus", "Bistro-Paulus-6806"),
	},
}

func GetConfig(cfg []byte) (ScrapersConfig, error) {
	if cfg == nil {
		return DefaultConfig, nil
	}

	var c ScrapersConfig
	err := json.Unmarshal(cfg, &c)
	if err != nil {
		return ScrapersConfig{}, fmt.Errorf("could not unmarshal config: %v", err)
	}

	return c, nil
}
