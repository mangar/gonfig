package main

import (
	"fmt"
	"os"
	"project/checking"
	"project/cmd"
	"project/config"
	"project/constants"
	"project/content"
	"project/log"

	emoji "gopkg.in/kyokomi/emoji.v1"
)

var (
	Version string
	Build   string
	Hash    string
)

func main() {

	if os.Getenv("FlagDebug") == "true" {
		constants.FlagDebug = true
		fmt.Println("Debug mode enabled")
	}

	if len(os.Args) == 2 {

		if os.Args[1] == "help" {
			cmd.Help()

		} else if os.Args[1] == "clear" || os.Args[1] == "clean" {
			cmd.Clear()

		} else if os.Args[1] == "version" {
			printVersion()

		} else {
			log.LogError("Command not found")

		}

	} else if len(os.Args) < 3 {
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

}

func printVersion() {
	out := emoji.Sprint(":checkered_flag: Version: ", Version, "\n:clock2: Build Time: ", Build)
	fmt.Println(out)

}
