# drone-pypi

[![Build Status](http://beta.drone.io/api/badges/drone-plugins/drone-pypi/status.svg)](http://beta.drone.io/drone-plugins/drone-pypi)
[![Join the discussion at https://discourse.drone.io](https://img.shields.io/badge/discourse-forum-orange.svg)](https://discourse.drone.io)
[![Go Report Card](https://goreportcard.com/badge/github.com/drone-plugins/drone-pypi)](https://goreportcard.com/report/github.com/drone-plugins/drone-pypi)

Basic Drone Plugin for PyPi publishes. The plugin use twine to upload packages.

## Build

Build the binary with the following commands:

```shell
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -tags netgo -o release/linux/amd64/drone-pypi
docker build --rm -t plugins/drone-pypi .
```

## Usage

```shell
docker run --rm \
  -e PLUGIN_USERNAME=jdoe \
  -e PLUGIN_PASSWORD=my_secret \
  -e PLUGIN_SKIP_BUILD=false \
  -v $(pwd):$(pwd) \
  -w $(pwd) \
  plugins/drone-pypi
```
