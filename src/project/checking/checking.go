package checking

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"os/user"
	"project/constants"
	"project/log"
	"time"
)

func InitialChecking() {
	rand.Seed(time.Now().Unix())

	checkConfigDir()
	CreateTempDir("")
	DeleteAndCreateTempDir(constants.ProjectName)
}

func checkConfigDir() {
	dir := HomeDir() + "/.gonfig"

	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		log.LogError("The ~/.gonfig dir does not exist")
		os.Exit(1)
	}

	file := ConfigFilePath()
	_, err2 := os.Stat(file)
	if os.IsNotExist(err2) {
		log.LogError("The ~/.gonfig/config.yml file does not exist")
		os.Exit(1)
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

/**
 *
 */
func DeleteAndCreateTempDir(project string) {

	tempDir := HomeDir() + "/.gonfig/temp/" + project
	log.LogDebug("Cleanning temporary dir @" + tempDir)

	if err := exec.Command("sh", "-c", "rm -rf "+tempDir).Run(); err != nil {
		log.LogError("Error removing build directory")
	}

	CreateTempDir(project)
}

/**
 *
 */
func CreateTempDir(project string) {

	tempDir := HomeDir() + "/.gonfig/temp/" + project

	_, err := os.Stat(tempDir)
	if os.IsNotExist(err) {
		log.LogDebug("Creating temp dir @" + tempDir)
		os.MkdirAll(tempDir, os.ModePerm)
	}
}
