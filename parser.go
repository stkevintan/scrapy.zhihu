package main

import (
	"io/ioutil"
	"log"

	"./browser"
	"github.com/pelletier/go-toml"
)

const configPath = "./config.toml"

//Config is define of config
type Config struct {
	Account     browser.Account
	MysqlConfig MysqlConfig
	TopicNames  []string
}

var content []byte

func init() {
	ret, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalf("Cannot Parse the config file: %s\n", configPath)
		return
	}
	content = ret
}

//Parser parse the config
func Parser() (*Config, error) {
	config := &Config{}
	err := toml.Unmarshal(content, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
