package conf

import (
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
)

var loaded = false
var configure Config

type Config struct {
	Brokers []Broker `toml:"broker"`
}

type Broker struct {
	Name      string
	APIKey    string `toml:"APIKey"`
	SecretKey string
}

func LoadConfig(path string) {
	if _, err := toml.DecodeFile(path, &configure); err != nil {
		//	return errors.New("Failed to load config : " + err.Error())
		return
	}
	loaded = true
}

func GetConfig() Config {
	if !loaded {
		logrus.Panic("config was not loaded, but get requested..")
	}
	return configure
}

func GetBrokers() map[string]Broker {
	if !loaded {
		logrus.Panic("config was not loaded, but get requested..")
	}
	return confToBrokerMap(configure.Brokers)
}
