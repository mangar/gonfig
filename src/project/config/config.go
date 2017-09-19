package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"project/checking"
	"project/constants"
	"project/log"
)

type ConfigSetup struct {
	Project string `json:project`
	Repo    string `json:repo`
}

var Configs = make([]ConfigSetup, 0)
var Config = ConfigSetup{}

func GetConfig() ConfigSetup {
	data, _ := ioutil.ReadFile(checking.ConfigFilePath())
	json.Unmarshal([]byte(data), &Configs)

	for _, config := range Configs {
		if config.Project == constants.ProjectName {
			Config = config
		}
	}

	if Config.Project == "" {
		log.LogError("Repository not found in config.json")
		os.Exit(1)
	}

	return Config
}
