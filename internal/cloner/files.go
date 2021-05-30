package cloner

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func getTmpDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	tmpDir := homeDir + "/.clonertmp"
	return tmpDir, nil
}

func createTmpDir() (string, error) {
	tmpDir, err := getTmpDir()
	if err != nil {
		return "", err
	}
	return createDir(tmpDir)
}

func removeTmpDir() error {
	tmpDir, err := getTmpDir()
	if err != nil {
		return err
	}

	return removeDir(tmpDir)
}

func removeDir(dir string) error {
	cmd := exec.Command("rm", "-rf", dir)
	_, err := cmd.CombinedOutput()
	return err
}

func createDir(path string) (string, error) {
	cmd := exec.Command("mkdir", path)
	_, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return path, nil
}

func copyDir(source, destination string) error {
	cmd := exec.Command("cp", "-R", source, destination)
	_, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	return nil
}

func openWith(app, path string) error {
	cmd := exec.Command(app, path)
	_, err := cmd.CombinedOutput()
	return err
}

func replaceInDirectory(dir, old, new string) error {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if !!info.IsDir() {
			return nil
		}

		read, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}

		newContents := strings.Replace(string(read), old, new, -1)

		err = ioutil.WriteFile(path, []byte(newContents), 0)
		if err != nil {
			panic(err)
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
