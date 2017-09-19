package config

import (
	"fmt"
	"os"
	"os/user"
)

func InitialChecking() {
	checkConfigDir()
}

func checkConfigDir() {
	dir := HomeDir() + "/.gonfig"

	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		panic("The ~/.gonfig dir does not exist")
	}

	file := HomeDir() + "/.gonfig/config.json"
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
