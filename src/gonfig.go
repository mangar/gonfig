package main

import (
	"fmt"
	"os"
	"project/checking"
	"project/config"
	"project/constants"
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

		if constants.FlagDebug {
			fmt.Println("Project Name:", constants.ProjectName)
			fmt.Println("Output:", constants.Output)
			fmt.Println("Content:")
		}

		for i, dir := range os.Args {
			if i > 2 {
				if constants.FlagDebug {
					fmt.Println(" -", dir)
				}
				constants.Directories = append(constants.Directories, dir)
			}
		}

	}

	checking.InitialChecking()
	config.GetConfig()

}
