package main

import (
	"cloner/internal/cloner"
	"cloner/internal/git"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"io/ioutil"
	"log"
	"os"
)

const ConfigFileName = ".clonerconfig"

type Config struct {
	Git         *git.Config     `json:"git" validate:"required"`
	Layouts     *cloner.Layouts `json:"layouts" validate:"required,dive,required"`
	ProjectsDir string          `json:"projects_dir" validate:"required"`
	Namespaces  []string        `json:"namespaces" validate:"required,dive,required"`
}

func main() {
	layouts := cloner.NewLayouts(
		&cloner.Layout{
			Name:        "layout-grpc",
			Namespace:   "examples",
			URL:         "https://qcm-git.mbrd.ru/service-platform/examples/layout-grpc",
			Description: "Grpc MTSBF template",
		},
		&cloner.Layout{
			Name:        "layout-http",
			Namespace:   "examples",
			URL:         "https://qcm-git.mbrd.ru/service-platform/examples/layout-http",
			Description: "Http MTSBF template",
		},
	)

	cfg := &Config{Layouts: layouts, ProjectsDir: "mygo", Namespaces: []string{"common-bank-services"}}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	configFilePath := homeDir + "/" + ConfigFileName
	err = cfg.ReadFromFile(configFilePath)
	if err != nil {
		panic("Cannot read config file " + configFilePath + " " + err.Error())
	}

	validate := validator.New()
	err = validate.Struct(cfg)
	if err != nil {
		panic(fmt.Sprintf("Config %s not valid: %s", configFilePath, err.Error()))
	}

	gitClient := git.New(cfg.Git)

	cl := cloner.New(cfg.Layouts, gitClient, cfg.ProjectsDir, cfg.Namespaces)
	fmt.Println(cloner.Logo)
	err = cl.Start()
	if err != nil {
		panic(err)
	}
}

func (c *Config) ReadFromFile(filePath string) error {
	content, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(content, c)
	if err != nil {
		return err
	}

	return nil
}
