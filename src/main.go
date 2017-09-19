package main

import (
	"project/checking"
	"project/config"
)

func main() {

	checking.InitialChecking()
	config.GetConfig()

}
