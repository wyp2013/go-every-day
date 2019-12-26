package utils

import (
	"github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	HashrateDB struct {
		User         string `yaml:"user"`
		Password     string `yaml:"password"`
		Addr         string `yaml:"addr"`
		DB           string `yaml:"db"`
		MaxOpenConns int    `yaml:"maxOpenConns"`
		MaxIdleConns int    `yaml:"maxIdleConns"`
	}

	Redis struct {
		Addr     string `yaml:"addr"`
		Password string `yaml:"password"`
		DB       int    `yaml:"db"`
	}

	Log struct {
		Level      string `yaml:"level"`
		ShowSource bool   `yaml:"showSource"`
		Path       string `yaml:"path"`
	}

	Health struct {
		Port            int `yaml:"level"`
		IntervalSeconds int `yaml:"intervalSeconds"`
	}

	PoolConfig string `yaml:"poolConfig"`
}

func (c *Config) GetPoolHashrateDBConn() string {
	conn := c.HashrateDB.User + ":" + c.HashrateDB.Password + "@tcp(" + c.HashrateDB.Addr + ")/" + c.HashrateDB.DB + "?parseTime=true"
	return conn
}

var cfg Config

func InitConfig(cfgFile string) (*Config, error) {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			log.Fatal(err)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName("config.yaml")
	}
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err == nil {
		log.WithField("config", viper.ConfigFileUsed()).Info("Using specified config file")
		if err := viper.Unmarshal(&cfg); err != nil {
			log.WithFields(log.Fields{"config": viper.ConfigFileUsed(), "error": err,}).Fatal("Error parsing config file")
			return nil, err
		}
	} else {
		log.WithFields(log.Fields{"config": viper.ConfigFileUsed(), "error": err,}).Fatal("Error parsing config file")
		return nil, err
	}

	return &cfg, nil
}

func GetConfig() *Config {
	if &cfg == nil {
		panic("config is not initial, should call InitConfig function")
	}
	return &cfg
}
