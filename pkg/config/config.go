package config

import "github.com/jinzhu/configor"

type Config struct {
	AppName string
	Port    string
	DB      struct {
		Mssql []struct {
			Enabled  bool   `default:"true"`
			Host     string `default:"localhost"`
			Port     string `default:"5532"`
			Username string `default:"localhost"`
			Password string `default:"password"`
			Database string `default:"prueba"`
			Nameconn string `default:"prueba"`
		}
	}
}

func NewConfig() (*Config, error) {
	c := &Config{}
	err := configor.Load(c, "config.yml")
	if err != nil {
		return nil, err
	}
	return c, nil
}
