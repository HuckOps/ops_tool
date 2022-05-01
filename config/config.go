package config

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

var Config ConfigMap

type ConfigMap struct {
	MySQL MySQL `yaml:"mysql"`
}

type MySQL struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	DB       string `json:"db"`
}

func InitConfig() {
	config := ConfigMap{}
	configFile, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	if err := yaml.Unmarshal(configFile, &config); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	Config = config
}
