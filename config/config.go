package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	NodeIPs  []string `yaml:"nodeIPs"`
	DbIP     string   `yaml:"dbIP"`
	DbPort   int      `yaml:"dbPort"`
	DbUser   string   `yaml:"dbUser"`
	DbPwd    string   `yaml:"dbPwd"` // 非安全实践，正确做法交互输入
	LogLevel string   `yaml:"logLevel"`
}

var (
	config       *Config
	ChainVersion string
)

func LoadConfig(configPath string) error {
	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Printf("yamlFile.Get err #%v\n ", err)
		return err
	}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		fmt.Printf("Unmarshal: %v\n", err)
		return err
	}

	return nil
}

func GetConfig() *Config {
	return config
}
