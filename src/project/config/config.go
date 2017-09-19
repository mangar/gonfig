package config

import "fmt"

type ConfigSetup struct {
	Project    string `json:project`
	Repository int    `json:repository`
}

func GetConfig() ConfigSetup {
	fmt.Println("GetConfig")

	return ConfigSetup{}
}
