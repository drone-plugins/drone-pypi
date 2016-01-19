[![Build Status](http://beta.drone.io/api/badges/drone-plugins/drone-pypi/status.svg)](http://beta.drone.io/drone-plugins/drone-pypi)
[![Docker Pulls](https://img.shields.io/docker/pulls/plugins/drone-pypi.svg)](https://hub.docker.com/r/plugins/drone-pypi/)

# drone-pypi

[![Build Status](http://beta.drone.io/api/badges/drone-plugins/drone-pypi/status.svg)](http://beta.drone.io/drone-plugins/drone-pypi)
[![](https://badge.imagelayers.io/plugins/drone-pypi:latest.svg)](https://imagelayers.io/?images=plugins/drone-pypi:latest 'Get your own badge on imagelayers.io')

Drone plugin for publishing to the Python package index

## Usage

Upload a source distribution to PyPI

```sh
./drone-pypi <<EOF
{
	"workspace": {
		"path": "/drone/my-module-py"
	}
	"vargs": {
		"username": "guido",
		"password": "secret"
	}
}
EOF
```

Upload a source distribution and a wheel to PyPI

```sh
./drone-pypi <<EOF
{
	"workspace": {
		"path": "/drone/my-module-py"
	}
	"vargs": {
		"distributions": ["sdist", "bdist_wheel"],
		"username": "guido",
		"password": "secret"
	}
}
EOF
```

Upload a source distribution to a private PyPI server, e.g. [simplepypi][]

```sh
./drone-pypi <<EOF
{
	"workspace": {
		"path": "/drone/my-module-py"
	}
	"vargs": {
		"repository": "https://pypi.example.com"
	}
}
EOF
```

[simplepypi]: https://github.com/steiza/simplepypi

## Docker

Build the Docker container using the `netgo` build tag to eliminate
the CGO dependency:

```sh
CGO_ENABLED=0 go build -a -tags netgo
docker build --rm=true -t plugins/drone-pypi .
```
