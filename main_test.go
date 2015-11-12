package main

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/drone/drone-plugin-go/plugin"
)

func TestDeploy(t *testing.T) {
	w := plugin.Workspace{
		Path: os.Getenv("DRONE_PYPI_PATH"),
	}
	repository := os.Getenv("DRONE_PYPI_REPOSITORY")
	username := os.Getenv("DRONE_PYPI_USERNAME")
	password := os.Getenv("DRONE_PYPI_PASSWORD")
	v := Params{
		Repository:    &repository,
		Username:      &username,
		Password:      &password,
		Distributions: strings.Split(os.Getenv("DRONE_PYPI_DISTRIBUTIONS"), " "),
	}
	if w.Path == "" {
		t.Skip("DRONE_PYPI_PATH not set")
	}
	err := deploy(&w, &v)
	if err != nil {
		t.Error(err)
	}
}

func sPtr(s string) *string {
	return &s
}

func TestConfig(t *testing.T) {
	testdata := []struct {
		repository *string
		username   *string
		password   *string
		exp        string
	}{
		{
			nil,
			nil,
			nil,
			`[distutils]
index-servers =
    pypi

[pypi]
repository: https://pypi.python.org/pypi
username: guido
password: secret
`,
		},
		{
			sPtr("https://pypi.example.com"),
			nil,
			nil,
			`[distutils]
index-servers =
    pypi

[pypi]
repository: https://pypi.example.com
username: guido
password: secret
`,
		},
		{
			nil,
			sPtr("jqhacker"),
			sPtr("supersecret"),
			`[distutils]
index-servers =
    pypi

[pypi]
repository: https://pypi.python.org/pypi
username: jqhacker
password: supersecret
`,
		},
	}
	for i, data := range testdata {
		v := Params{
			Repository:    data.repository,
			Username:      data.username,
			Password:      data.password,
			Distributions: []string{},
		}
		var b bytes.Buffer
		v.WriteConfig(&b)
		if b.String() != data.exp {
			t.Errorf("Case %d: Expected %s, got %s\n", i, data.exp, b.String())
		}
	}
}

func TestUpload(t *testing.T) {
	testdata := []struct {
		distributions []string
		exp           []string
	}{
		{
			[]string{},
			[]string{"python", "setup.py", "sdist", "upload", "-r", "pypi"},
		},
		{
			[]string{"sdist", "bdist_wheel"},
			[]string{"python", "setup.py", "sdist", "bdist_wheel", "upload", "-r", "pypi"},
		},
	}
	for i, data := range testdata {
		v := Params{Distributions: data.distributions}
		c, err := v.Upload()
		if err != nil {
			t.Error(err)
		}
		if len(c.Args) != len(data.exp) {
			t.Errorf("Case %d: Expected %d, got %d", i, len(data.exp), len(c.Args))
		}
		for i := range c.Args {
			if c.Args[i] != data.exp[i] {
				t.Errorf("Case %d: Expected %s, got %s", i, strings.Join(data.exp, " "), strings.Join(c.Args, " "))
			}
		}
	}
}
