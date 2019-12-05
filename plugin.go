package main

import (
	"log"
	"os"
	"os/exec"
	"strings"

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
}

func (p Plugin) buildCommand() *exec.Cmd {
	// Set the default of distributions in here
	// as CLI package still has issues with string slice defaults
	distributions := []string{"sdist"}
	// Sanitize bad dist values
	inputDists := make([]string, 0, len(p.Distributions))
	for _, dist := range p.Distributions {
		if strings.TrimSpace(dist) == "" {
			continue
		}
		inputDists = append(inputDists, dist)
	}
	if len(inputDists) > 0 {
		distributions = p.Distributions
	}
	args := []string{p.SetupFile}
	for i := range distributions {
		args = append(args, distributions[i])
	}
	return exec.Command("python3", args...)
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
	args = append(args, "dist/*")

	return exec.Command("twine", args...)
}

// Exec runs the plugin - doing the necessary setup.py modifications
func (p *Plugin) Exec() error {
	// If a setup.py is in a subdirectory, we need to change to that directory first
	// so the correct files are packaged.
	pathParts := strings.Split(p.SetupFile, string(os.PathSeparator))
	if len(pathParts) > 1 {
		pathParts := pathParts[0 : len(pathParts)-1]
		packageDir := strings.Join(pathParts, string(os.PathSeparator))
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}
		defer os.Chdir(cwd)
		err = os.Chdir(packageDir)
		if err != nil {
			return errors.Wrap(err, "Failed to chdir to "+packageDir)
		}
		// Now change the setup file value as well, since it's relative and we're in
		// the correct place.
		p.SetupFile = "setup.py"
	}

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
