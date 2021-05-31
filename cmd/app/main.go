package main

import (
	"cloner/internal/cloner"
	"cloner/internal/git"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/go-playground/validator/v10"
	"io/ioutil"
	"os"
)

const AppVersion ="v0.0.1"

type Config struct {
	Git         *git.Config     `json:"git" validate:"required"`
	Layouts     *cloner.Layouts `json:"layouts" validate:"required,dive,required"`
	ProjectsDir string          `json:"projects_dir" validate:"required"`
	Namespaces  []string        `json:"namespaces" validate:"required,dive,required"`
}

func main() {
	version := flag.Bool("version", false, "display version")
	flag.Parse()

	if *version {
		fmt.Println(AppVersion)
		return
	}

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
		&cloner.Layout{
			Name:        "layout-pub-sub",
			Namespace:   "examples",
			URL:         "https://qcm-git.mbrd.ru/service-platform/examples/layout-pub-sub",
			Description: "Pub-sub MTSBF template",
		},
	)

	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	cfg := &Config{Layouts: layouts, ProjectsDir: homeDir + "/" + "mygo", Namespaces: []string{"common-bank-services"}}
	configFilePath := homeDir + "/" + cloner.ConfigFileName
	var configFile string
	flag.StringVar(&configFile, "configfile", configFilePath, "Init config")
	flag.Parse()

	err = cfg.ReadFromFile(configFile)
	if err != nil {
		panic(`Cannot read config file ` + configFile + " " + err.Error() + "\n\n" +
			`Try to create a config file and set git credentials: ` + "\n" +
			`echo '{"git": {"user":"", "password":""}}' > ~/.clonerconfig`)
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
		return err
	}
	err = json.Unmarshal(content, c)
	if err != nil {
		return err
	}

	return nil
}
