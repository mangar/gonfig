package config

import (
	"encoding/json"
	"io/ioutil"
	"project/checking"
	"project/log"
)

type ConfigSetup struct {
	Project string `json:project`
	Repo    string `json:repo`
}

var Config = ConfigSetup{}

func GetConfig() ConfigSetup {
	data, _ := ioutil.ReadFile(checking.ConfigFilePath())
	json.Unmarshal([]byte(data), &Config)

	log.LogDebug("Config: " + Config.Project + " / " + Config.Repo)

	return Config
}
