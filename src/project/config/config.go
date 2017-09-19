package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"project/checking"
	"project/constants"
)

type ConfigSetup struct {
	Project string `json:project`
	Repo    string `json:repo`
}

var Config = ConfigSetup{}

func GetConfig() ConfigSetup {
	data, _ := ioutil.ReadFile(checking.ConfigFilePath())
	json.Unmarshal([]byte(data), &Config)

	if constants.FlagDebug {
		fmt.Println("Config:", Config)
	}

	return Config
}
