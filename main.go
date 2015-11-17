package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/drone/drone-plugin-go/plugin"
)

type Params struct {
	Distributions []string `json:"distributions"`
	Password      *string  `json:"password,omitempty"`
	Repository    *string  `json:"repository,omitempty"`
	Username      *string  `json:"username,omitempty"`
}

func main() {
	w := plugin.Workspace{}
	v := Params{}
	plugin.Param("workspace", &w)
	plugin.Param("vargs", &v)
	plugin.MustParse()

	err := deploy(&w, &v)
	if err != nil {
		log.Fatal(err)
	}
}

func deploy(w *plugin.Workspace, v *Params) error {
	err := createConfig(v)
	if err != nil {
		return err
	}
	err = uploadDist(w, v)
	if err != nil {
		return err
	}
	return nil
}

func createConfig(v *Params) error {
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

func uploadDist(w *plugin.Workspace, v *Params) error {
	cmd, err := v.Upload()
	if err != nil {
		return err
	}
	cmd.Dir = w.Path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fmt.Println("$", strings.Join(cmd.Args, " "))
	err = cmd.Run()
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

// Upload creates a setuptools upload command.
func (v *Params) Upload() (*exec.Cmd, error) {
	distributions := []string{"sdist"}
	if len(v.Distributions) > 0 {
		distributions = v.Distributions
	}
	args := []string{"python", "setup.py"}
	for i := range distributions {
		args = append(args, distributions[i])
	}
	args = append(args, "upload")
	args = append(args, "-r")
	args = append(args, "pypi")
	return command(args)
}

// Command builds a command using a variable length argument list.
func command(args []string) (*exec.Cmd, error) {
	name := args[0]
	cmd := &exec.Cmd{
		Path: name,
		Args: args,
	}
	if filepath.Base(name) == name {
		lp, err := exec.LookPath(name)
		if err != nil {
			return nil, err
		}
		cmd.Path = lp
	}
	return cmd, nil
}
