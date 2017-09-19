package main

import (
	"fmt"
	"os"
	"project/checking"
	"project/config"
	"project/constants"
	"project/content"
	"project/log"
)

func main() {

	if os.Getenv("FlagDebug") == "true" {
		constants.FlagDebug = true
		fmt.Println("Debug mode enabled")
	}

	if len(os.Args) < 3 {
		fmt.Println("Inform the (1)project name, (2) destination and the (3) diretories to be downloaded")
		fmt.Println("Ex.:")
		fmt.Println("go run main.go projectDemo ./ microservices/all/development microservices/users/development")
		os.Exit(1)

	} else {

		constants.ProjectName = os.Args[1]
		constants.Output = os.Args[2]

		log.LogDebug("Project Name: " + constants.ProjectName)
		log.LogDebug("Output: " + constants.Output)
		log.LogDebug("Content:")

		for i, dir := range os.Args {
			if i > 2 {
				log.LogDebug(" -" + dir)
				constants.Directories = append(constants.Directories, dir)
			}
		}

	}

	checking.InitialChecking()
	config.GetConfig()

	if content.GitClone() {
		log.LogDebug("Content cloned!")
		if content.DistributeContent() {
			log.LogSuccess("Content Distributed")
		} else {
			log.LogError("Ups... found some problem distributing the content.")
		}
	}

}
