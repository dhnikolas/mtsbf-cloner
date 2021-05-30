package cloner

import (
	"cloner/internal/git"
	"fmt"
	"os"
)

type Cloner struct {
	GitClient *git.Git
	Layouts     *Layouts
	ProjectsDir string
	Namespaces []string
}

func New(layouts *Layouts, gitClient *git.Git, projectsDir string, namespaces []string) *Cloner {
	return &Cloner{Layouts: layouts, GitClient: gitClient, ProjectsDir: projectsDir, Namespaces: namespaces}
}

func (c *Cloner) Start() error {

	tmpDir, err := createTmpDir()
	if err != nil {
		return err
	}
	defer func() {
		err := removeTmpDir()
		if err != nil {
			fmt.Println(err)
		}
	}()

	l, err := c.AskLayout()
	if err != nil {
		return err
	}

	ch := make(chan error)

	go func(ch chan error) {
		err = c.GitClient.Clone(l.URL, tmpDir)
		ch <- err
	}(ch)

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	projectName, err := c.AskProjectName()
	if err != nil {
		return err
	}
	err =  <- ch
	if err != nil {
		return err
	}

	projectDir, err := createDir(homeDir + "/" + c.ProjectsDir + "/" + projectName)
	if err != nil {
		return err
	}
	err = copyDir(tmpDir+"/", projectDir)
	if err != nil {
		return err
	}

	err = removeDir(projectDir + "/.git")
	if err != nil {
		return err
	}

	namespace, err := c.AskNamespace()
	if err != nil {
		return err
	}

	err = replaceInDirectory(projectDir, l.Name, projectName)
	if err != nil {
		return err
	}

	err = replaceInDirectory(projectDir, l.Namespace, namespace)
	if err != nil {
		return err
	}

	app, err := c.AskOpenWith()
	if err != nil {
		return err
	}

	if app != "none" {
		err = openWith(app, projectDir)
		if err != nil {
			fmt.Println(err)
		}
	}

	return nil
}



