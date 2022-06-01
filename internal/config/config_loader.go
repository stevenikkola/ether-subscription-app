package config

import (
	"encoding/json"
	"io/ioutil"
)

func getConfig[T any](filepath string, config T) T {
	fileContents, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(fileContents, &config)
	if err != nil {
		panic(err)
	}
	return config
}

func GetBloxrouteConfig() BloxrouteConfig {
	var bloxrouteConfig BloxrouteConfig
	return getConfig("resources/bloXroute_config.json", bloxrouteConfig)
}

func GetMainConfig() MainConfig {
	var mainConfig MainConfig
	return getConfig("resources/main_config.json", mainConfig)
}

type BloxrouteConfig struct {
	WebsocketsCloudApiBaseUri string
	AuthorizationHeader       string
}

type MainConfig struct {
	SubscriptionFilters SubscriptionFilters
}

type SubscriptionFilters struct {
	ToAddress   string
	FromAddress string
}
