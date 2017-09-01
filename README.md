# drone-pypi

[![Build Status](http://beta.drone.io/api/badges/drone-plugins/drone-pypi/status.svg)](http://beta.drone.io/drone-plugins/drone-pypi)
[![Join the chat at https://gitter.im/drone/drone](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/drone/drone)
[![Go Doc](https://godoc.org/github.com/drone-plugins/drone-pypi?status.svg)](http://godoc.org/github.com/drone-plugins/drone-pypi)
[![Go Report](https://goreportcard.com/badge/github.com/drone-plugins/drone-pypi)](https://goreportcard.com/report/github.com/drone-plugins/drone-pypi)
[![](https://images.microbadger.com/badges/image/plugins/pypi.svg)](https://microbadger.com/images/plugins/pypi "Get your own image badge on microbadger.com")

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
