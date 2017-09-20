package cmd

import (
	"os/exec"
	"project/checking"
	"project/log"
)

func Clear() {
	tempDir := checking.HomeDir() + "/.gonfig/temp/"

	log.LogDebug("Cleanning temporary dir: " + tempDir)

	if err := exec.Command("sh", "-c", "rm -rf "+tempDir).Run(); err != nil {
		log.LogError("Error removing build directory")
	}

	checking.CreateTempDir("")

	log.LogClean("Temporary dir empty.")
}
