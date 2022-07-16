package cloner

import (
	"errors"
	"github.com/manifoldco/promptui"
	"regexp"
)

func (c *Cloner) AskLayout() (*Layout, error) {
	prompt := promptui.Select{
		Label: "Select layout",
		Items: c.Layouts.GetNames(),
	}

	_, layoutName, err := prompt.Run()
	if err != nil {
		return nil, err
	}
	layout, ok := c.Layouts.Get(layoutName)
	if !ok {
		return nil, errors.New("Layout not exist: " + layoutName)
	}

	return layout, nil
}

func (c *Cloner) AskProjectName() (string, error) {
	var isLetter = regexp.MustCompile(`^[a-z-]{3,50}$`).MatchString
	validate := func(input string) error {
		if !isLetter(input) {
			return errors.New("Project name must have more than 3 characters and contains only lowercase letters and - ")
		}
		fullPath := c.ProjectsDir + "/" + input
		ok, err := exists(fullPath)
		if err != nil {
			return err
		}

		if ok {
			return errors.New("Files already exists " + fullPath)
		}

		return nil
	}
	prompt := promptui.Prompt{
		Label:    "New Project name",
		Validate: validate,
	}

	result, err := prompt.Run()

	if err != nil {
		return "", err
	}

	return result, nil
}

func (c *Cloner) AskNamespace() (string, error) {
	prompt := promptui.Select{
		Label: "Select namespace",
		Items: append(c.Namespaces, "other"),
	}

	_, namespace, err := prompt.Run()
	if err != nil {
		return "", err
	}

	if namespace == "other" {
		var isLetter = regexp.MustCompile(`^[a-z-]{1,50}$`).MatchString
		validate := func(input string) error {
			if !isLetter(input) {
				return errors.New("Namespace must have more than 5 characters and contains only lowercase letters and - ")
			}
			return nil
		}
		prompt := promptui.Prompt{
			Label:    "Other Namespace",
			Validate: validate,
		}
		namespace, err = prompt.Run()
		if err != nil {
			return "", err
		}
	}

	return namespace, nil
}

func (c *Cloner) AskOpenWith() (string, error) {
	prompt := promptui.Select{
		Label: "Select application",
		Items: []string{"goland", "vscode", "none"},
	}

	_, app, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return app, nil
}
