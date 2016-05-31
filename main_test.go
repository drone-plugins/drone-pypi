package main

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/drone/drone-go/drone"
)

// TestPublish checks if this module can successfully publish a PyPI
// package. A simple module is included in the `testdata` directory.
//
// To run this test against the PyPI test server:
//
// 1. register a new account (https://wiki.python.org/moin/TestPyPI)
// 2. Export DRONE_PYPI_PATH, DRONE_PYPI_REPOSITORY, DRONE_PYPI_USERNAME,
//    DRONE_PYPI_PASSWORD, and DRONE_PYPI_DISTRIBUTIONS
// 3. Run the test suite
//
// For example:
//
//     $ export DRONE_PYPI_PATH=testdata
//     $ export DRONE_PYPI_REPOSITORY=https://testpypi.python.org/pypi
//     $ export DRONE_PYPI_USERNAME=drone_pypi_test
//     $ export DRONE_PYPI_PASSWORD=$uper$ecretPassword
//     $ export DRONE_PYPI_DISTRIBUTIONS=sdist
//     $ go test -run TestPublish
//
// > NOTE: PyPI will refuse to upload the same version of a module twice,
// > however setup.py still returns zero to the shell so this appears as a
// > successful test.
func TestPublish(t *testing.T) {
	w := drone.Workspace{Path: os.Getenv("DRONE_PYPI_PATH")}
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
	err := v.Deploy(&w)
	if err != nil {
		t.Error(err)
	}
}

// TestConfig checks if a PyPI configuration file can be generated.
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

// TestUpload checks if a distutils upload command can be properly
// formatted.
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
		c := v.Upload()
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

// TestUploadTagged checks if a distutils upload command can be properly
// formatted with egg_info tag.
func TestUploadTagged(t *testing.T) {
	v := Params{Tag: "1234"}
	c := v.Upload()
	exp := []string{"python", "setup.py", "egg_info", "-b", "1234", "sdist", "upload", "-r", "pypi"}
	if len(c.Args) != len(exp) {
		t.Errorf("Expected %d, got %d", len(exp), len(c.Args))
	}
	for i := range c.Args {
		if c.Args[i] != exp[i] {
			t.Errorf("Expected %s, got %s", strings.Join(exp, " "), strings.Join(c.Args, " "))
		}
	}
}

func sPtr(s string) *string {
	return &s
}
