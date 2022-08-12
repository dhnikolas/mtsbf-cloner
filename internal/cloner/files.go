package cloner

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func getTmpDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	tmpDir := homeDir + "/" + TmpDirectory
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
	err := os.RemoveAll(dir)
	return err
}

func createDir(path string) (string, error) {
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return "", err
	}

	return path, nil
}

func copyDir(source, destination string) error {
	err := filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		relPath := strings.Replace(path, source, "", 1)
		if relPath == "" {
			return nil
		}
		if info.IsDir() {
			return os.Mkdir(filepath.Join(destination, relPath), 0755)
		} else {
			data, err := ioutil.ReadFile(filepath.Join(source, relPath))
			if err != nil {
				return err
			}
			return ioutil.WriteFile(filepath.Join(destination, relPath), data, 0777)
		}
	})

	return err
}

func openWith(app, path string) error {
	command, args := makeCommand(app, path)
	cmd := exec.Command(command, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("Error: %s, %s ", out, err.Error())
	}

	return nil
}

func makeCommand(app, path string) (string, []string) {
	switch runtime.GOOS {
	case "darwin":
		return "open", []string{"-na", getDarwinAppName(app), "--args", path}
	default:
		return app, []string{path}
	}
}

func getDarwinAppName(app string) string {
	switch app {
	case "goland":
		return "GoLand.app"
	case "vscode":
		return "Visual Studio.app"
	default:
		return app
	}
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
