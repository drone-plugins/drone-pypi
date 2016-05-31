package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin"
)

// Params desribes how to upload a Python module to PyPI.
type Params struct {
	Distributions []string `json:"distributions"`
	Password      *string  `json:"password,omitempty"`
	Repository    *string  `json:"repository,omitempty"`
	Username      *string  `json:"username,omitempty"`
	Tag           *string  `json:"tag,omitempty"`
}

func main() {
	w := drone.Workspace{}
	v := Params{}
	plugin.Param("workspace", &w)
	plugin.Param("vargs", &v)
	plugin.MustParse()

	err := v.Deploy(&w)
	if err != nil {
		log.Fatal(err)
	}
}

// Deploy creates a PyPI configuration file and uploads a module.
func (v *Params) Deploy(w *drone.Workspace) error {
	err := v.CreateConfig()
	if err != nil {
		return err
	}
	err = v.UploadDist(w)
	if err != nil {
		return err
	}
	return nil
}

// CreateConfig creates a PyPI configuration file in the home directory of
// the current user.
func (v *Params) CreateConfig() error {
	f, err := os.Create(path.Join(os.Getenv("HOME"), ".pypirc"))
	if err != nil {
		return err
	}
	defer f.Close()
	buf := bufio.NewWriter(f)
	err = v.WriteConfig(buf)
	if err != nil {
		return err
	}
	buf.Flush()
	return nil
}

// UploadDist executes a distutils command to upload a python module.
func (v *Params) UploadDist(w *drone.Workspace) error {
	cmd := v.Upload()
	cmd.Dir = w.Path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fmt.Println("$", strings.Join(cmd.Args, " "))
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

// WriteConfig writes a .pypirc to a supplied io.Writer.
func (v *Params) WriteConfig(w io.Writer) error {
	repository := "https://pypi.python.org/pypi"
	if v.Repository != nil {
		repository = *v.Repository
	}
	username := "guido"
	if v.Username != nil {
		username = *v.Username
	}
	password := "secret"
	if v.Password != nil {
		password = *v.Password
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
func (v *Params) Upload() *exec.Cmd {
	distributions := []string{"sdist"}
	if len(v.Distributions) > 0 {
		distributions = v.Distributions
	}
	args := []string{"setup.py"}
	if v.Tag != nil {
		args = append(args, "egg_info", "-b", v.Tag)
	}
	for i := range distributions {
		args = append(args, distributions[i])
	}
	args = append(args, "upload")
	args = append(args, "-r")
	args = append(args, "pypi")
	return exec.Command("python", args...)
}
