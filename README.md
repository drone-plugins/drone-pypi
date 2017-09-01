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
docker run --rm               \
  -e PLUGIN_USERNAME=username \
  -e PLUGIN_PASSWORD=password \
  -v $(pwd):$(pwd) \
  -w $(pwd) \
  plugins/pypi
```

Upload a source distribution to a private PyPI server, e.g. [simplepypi][]

```sh
docker run --rm                                  \
  -e PLUGIN_USERNAME=username                    \
  -e PLUGIN_PASSWORD=password                    \
  -e PLUGIN_RESPOSITORY=https://pypi.example.com \
  -v $(pwd):$(pwd)                               \
  -w $(pwd)                                      \
  plugins/pypi
```

[simplepypi]: https://github.com/steiza/simplepypi

## Build

Build the binary with the following command:

```sh
go build
```

## Docker
  		  
Build the Docker image with the following commands:
  		  
```
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -tags netgo -o release/linux/amd64/drone-pypi
docker build --rm -t plugins/pypi .
```

