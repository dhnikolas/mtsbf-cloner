package git

import (
	git_client "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

type Git struct {
	config *Config
}

type Config struct {
	Url      string `json:"url"`
	User     string `json:"user" validate:"required"`
	Password string `json:"password" validate:"required"`
	Path     string `json:"path"`
}

func New(cfg *Config) *Git {
	return &Git{config: cfg}
}

func (g *Git) Clone (path, destination string) error {
	_, err := git_client.PlainClone(destination, false, &git_client.CloneOptions{
		Auth: &http.BasicAuth{
			Username: g.config.User,
			Password: g.config.Password,
		},
		URL:      path,
	})

	return err
}



