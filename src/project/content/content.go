package content

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"project/checking"
	"project/config"
	"project/constants"
	"project/log"
)

/**
 * Clone the content into the temp dir.
 */
func GitClone() bool {
	ret := true

	log.LogGit("Clonning content")

	cloneDir := checking.HomeDir() + "/.gonfig/temp/" + constants.ProjectName
	log.LogDebug("Clonning into dir: " + cloneDir)

	command := "git clone " + config.Config.Repo + " " + cloneDir
	log.LogDebug(":: " + command)

	cmd := exec.Command("sh", "-c", command)
	err := cmd.Run()
	if err != nil {
		log.LogError("Problems clonning repository")
		ret = false
	}

	return ret
}

/**
 *
 */
func DistributeContent() bool {
	ret := true
	for _, dir := range constants.Directories {
		log.LogDebug("Copying " + dir)

		fromDir := checking.HomeDir() + "/.gonfig/temp/" + constants.ProjectName + "/" + dir
		log.LogDebug("  fromDir: " + fromDir)
		log.LogDebug("  toDir: " + constants.Output)

		copy_folder(fromDir, constants.Output)
	}

	return ret
	// return true
}

func copy_folder(source string, dest string) (err error) {

	sourceinfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	err = os.MkdirAll(dest, sourceinfo.Mode())
	if err != nil {
		return err
	}

	directory, _ := os.Open(source)

	objects, err := directory.Readdir(-1)

	for _, obj := range objects {

		sourcefilepointer := source + "/" + obj.Name()

		destinationfilepointer := dest + "/" + obj.Name()

		if obj.IsDir() {
			err = copy_folder(sourcefilepointer, destinationfilepointer)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			err = copy_file(sourcefilepointer, destinationfilepointer)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
	return
}

func copy_file(source string, dest string) (err error) {
	sourcefile, err := os.Open(source)
	if err != nil {
		return err
	}

	defer sourcefile.Close()

	destfile, err := os.Create(dest)
	if err != nil {
		return err
	}

	defer destfile.Close()

	_, err = io.Copy(destfile, sourcefile)
	if err == nil {
		sourceinfo, err := os.Stat(source)
		if err != nil {
			err = os.Chmod(dest, sourceinfo.Mode())
		}

	}

	return
}
