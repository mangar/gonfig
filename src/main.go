package main

import (
	"fmt"
	"project/config"
)

var ConfigSetup = config.ConfigSetup{}

func main() {

	config.InitialChecking()
	ConfigSetup := config.GetConfig()

	fmt.Println("ConfigSetup:", ConfigSetup)
}
