package checking

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"project/constants"
)

func InitialChecking() {
	checkConfigDir()
	CreateTempDir("")
	DeleteAndCreateTempDir(constants.ProjectName)
}

func checkConfigDir() {
	dir := HomeDir() + "/.gonfig"

	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		panic("The ~/.gonfig dir does not exist")
	}

	file := ConfigFilePath()
	_, err2 := os.Stat(file)
	if os.IsNotExist(err2) {
		panic("The ~/.gonfig/config.yml file does not exist")
	}

}

/**
 * Returns the HomeDir of the logged in user
 */
func HomeDir() string {
	var homeDir string
	usr, err := user.Current()

	if err == nil {
		homeDir = usr.HomeDir

	} else {
		homeDir = os.Getenv("HOME")
	}

	return homeDir
}

func printHelp() {
	fmt.Println("HELP")
	os.Exit(1)
}

/**
 * Returns the Config file path (~/.gonfig/config.json)
 */
func ConfigFilePath() string {
	return HomeDir() + "/.gonfig/config.json"
}

func DeleteAndCreateTempDir(project string) {

	tempDir := HomeDir() + "/.gonfig/temp/" + project
	if constants.FlagDebug {
		fmt.Println("Cleanning temporary dir @", tempDir)
	}

	if err := exec.Command("sh", "-c", "rm -rf "+tempDir).Run(); err != nil {
		if constants.FlagDebug {
			fmt.Printf("Error removing build directory: %s\n", err)
		}
	}

	CreateTempDir(project)
}

func CreateTempDir(project string) {

	tempDir := HomeDir() + "/.gonfig/temp/" + project

	_, err := os.Stat(tempDir)
	if os.IsNotExist(err) {
		if constants.FlagDebug {
			fmt.Println("Creating temp dir @", tempDir)
		}
		os.MkdirAll(tempDir, os.ModePerm)
	}

	// return HomeDir() + "/.gonfig/temp" + project
}
