package cmd

import "fmt"

func Help() {

	fmt.Println("")
	fmt.Println("usage: gonfig project_name output_dir directory_content1 [ directory_content2 ... ]")
	fmt.Println("")
	fmt.Println("Ex.:")
	fmt.Println(" gonfig ProjectName ../output microservices/all/development microservices/users/development")
	fmt.Println("")
	fmt.Println("The above example will look for the project: ProjectName in ~/.gonfig/config.json, copy the content from the folders: ")
	fmt.Println("microservices/all/development microservices/users/development into ../output dir.")
	fmt.Println("")
	fmt.Println("options:")
	fmt.Println(" clear    perform the cleaning on temporary dir. The temporary dir is located at ~/.gonfig/ ")
	fmt.Println(" help     show this content")
	fmt.Println(" install  create the basic structure for configruation: ~/.gonfig dir and the config.json basic file")
	fmt.Println(" version  display the version and build time")
	fmt.Println("")

}
