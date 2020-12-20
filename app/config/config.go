package config

import (
	"github.com/tkanos/gonfig"
)

// Configuration is...
type Configuration struct {
	DBUSERNAME string
	DBPASSWORD string
	DBPORT     string
	DBHOST     string
	DBNAME     string
}

// GetConfig is...
func GetConfig() Configuration {
	conf := Configuration{}
	gonfig.GetConf("config/config.json", &conf)
	return conf
}
