package main

import (
	"log"
	"os/exec"
	"path"
	"path/filepath"

	"github.com/pkg/errors"
)

// Plugin defines the PyPi plugin parameters
type Plugin struct {
	Repository    string
	Username      string
	Password      string
	SetupFile     string
	Distributions []string
	SkipBuild     bool
	DistDir       string
}

func (p Plugin) buildCommand() *exec.Cmd {
	// Set the default of distributions in here
	// as CLI package still has issues with string slice defaults
	distributions := []string{"sdist"}
	if len(p.Distributions) > 0 {
		distributions = p.Distributions
	}
	dirToWork := path.Dir(p.SetupFile)
	args := []string{"setup.py"}
	for i := range distributions {
		args = append(args, distributions[i])
	}
	cmd := exec.Command("python3", args...)
	cmd.Dir = dirToWork
	return cmd
}

func (p Plugin) uploadCommand() *exec.Cmd {
	args := []string{}
	args = append(args, "upload")
	args = append(args, "--repository-url")
	args = append(args, p.Repository)
	args = append(args, "--username")
	args = append(args, p.Username)
	args = append(args, "--password")
	args = append(args, p.Password)
	args = append(args, filepath.Join(p.DistDir, "/*"))

    cmd := exec.Command("twine", args...)
    cmd.Dir = path.Dir(p.SetupFile)
	return cmd
}

// Exec runs the plugin - doing the necessary setup.py modifications
func (p Plugin) Exec() error {
	if !p.SkipBuild {
		out, err := p.buildCommand().CombinedOutput()
		if err != nil {
			return errors.Wrap(err, string(out))
		}
		log.Printf("Output: %s", out)
	}

	out, err := p.uploadCommand().CombinedOutput()
	if err != nil {
		return errors.Wrap(err, string(out))
	}
	log.Printf("Output: %s", out)

	return nil
}
