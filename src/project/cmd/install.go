package cmd

import (
	"io/ioutil"
	"os"
	"project/checking"
	"project/log"
)

func Install() {

	createHome()
	createConfigFile()
	log.LogSuccess("Installation complete!")

}

func createHome() {

	tempDir := checking.HomeDir() + "/.gonfig"

	_, err := os.Stat(tempDir)
	if os.IsNotExist(err) {
		log.LogDebug("Creating home dir @" + tempDir)
		os.MkdirAll(tempDir, os.ModePerm)
	} else {
		log.LogDebug("Home dir already created. Nothing to do")
	}

}

func createConfigFile() {

	tempDir := checking.HomeDir() + "/.gonfig/config.json"
	_, err := os.Stat(tempDir)

	if os.IsNotExist(err) {
		log.LogDebug("Creating config.json @" + tempDir)

		_content := []byte("[\n" +
			"	{\n" +
			"		\"project\":\"ProjectName\",\n" +
			"		\"repo\" : \"https://USERNAME:APP_PASSWORD@bitbucket.org/REPOSITORY_NAME.git\"\n" +
			"	}\n" +
			"]\n")

		err := ioutil.WriteFile(tempDir, _content, 0644)
		if err != nil {
			log.LogError("Error writing on file")
		}

	} else {
		log.LogDebug("Config file already there. Nothing to do")
	}

}
