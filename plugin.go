package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"strings"
)

type Plugin struct {
	Repository    string   `json:"repository,omitempty"`
	Username      string   `json:"username,omitempty"`
	Password      string   `json:"password,omitempty"`
	Distributions []string `json:"distributions"`
}

func (p *Plugin) Exec() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	return p.Deploy(dir)
}

// Deploy creates a PyPI configuration file and uploads a module.
func (p *Plugin) Deploy(w string) error {
	err := p.CreateConfig()
	if err != nil {
		return err
	}
	err = p.UploadDist(w)
	if err != nil {
		return err
	}
	return nil
}

// CreateConfig creates a PyPI configuration file in the home directory of
// the current user.
func (p *Plugin) CreateConfig() error {
	f, err := os.Create(path.Join(os.Getenv("HOME"), ".pypirc"))
	if err != nil {
		return err
	}
	defer f.Close()
	buf := bufio.NewWriter(f)
	err = p.WriteConfig(buf)
	if err != nil {
		return err
	}
	buf.Flush()
	return nil
}

// UploadDist executes a distutils command to upload a python module.
func (p *Plugin) UploadDist(w string) error {
	cmd := p.Upload()
	cmd.Dir = w
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fmt.Println("$", strings.Join(cmd.Args, " "))
	fmt.Println(w)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

// WriteConfig writes a .pypirc to a supplied io.Writer.
func (p *Plugin) WriteConfig(w io.Writer) error {
	repository := "https://pypi.python.org/pypi"
	if p.Repository != "" {
		repository = p.Repository
	}
	username := "guido"
	if p.Username != "" {
		username = p.Username
	}
	password := "secret"
	if p.Password != "" {
		password = p.Password
	}
	_, err := io.WriteString(w, fmt.Sprintf(`[distutils]
index-servers =
    pypi

[pypi]
repository: %s
username: %s
password: %s
`, repository, username, password))
	return err
}

// Upload creates a distutils upload command.
func (p *Plugin) Upload() *exec.Cmd {
	distributions := []string{"sdist"}
	if len(p.Distributions) > 0 {
		distributions = p.Distributions
	}
	args := []string{"setup.py"}
	for i := range distributions {
		args = append(args, distributions[i])
	}
	args = append(args, "upload")
	args = append(args, "-r")
	args = append(args, "pypi")
	return exec.Command("python", args...)
}
