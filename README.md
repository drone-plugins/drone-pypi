# drone-pypi

[![Build Status](http://cloud.drone.io/api/badges/drone-plugins/drone-pypi/status.svg)](http://cloud.drone.io/drone-plugins/drone-pypi)
[![Gitter chat](https://badges.gitter.im/drone/drone.png)](https://gitter.im/drone/drone)
[![Join the discussion at https://discourse.drone.io](https://img.shields.io/badge/discourse-forum-orange.svg)](https://discourse.drone.io)
[![Drone questions at https://stackoverflow.com](https://img.shields.io/badge/drone-stackoverflow-orange.svg)](https://stackoverflow.com/questions/tagged/drone.io)
[![](https://images.microbadger.com/badges/image/plugins/pypi.svg)](https://microbadger.com/images/plugins/pypi "Get your own image badge on microbadger.com")
[![Go Doc](https://godoc.org/github.com/drone-plugins/drone-pypi?status.svg)](http://godoc.org/github.com/drone-plugins/drone-pypi)
[![Go Report](https://goreportcard.com/badge/github.com/drone-plugins/drone-pypi)](https://goreportcard.com/report/github.com/drone-plugins/drone-pypi)

Drone Plugin for PyPi publishing with [twine](https://pypi.org/project/twine/). For the usage information and a listing of the available options please take a look at [the docs](http://plugins.drone.io/plugins/pypi/).

## Build

Build the binary with the following commands:

```Shell
export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0
export GO111MODULE=on

go build -v -a -tags netgo -o release/linux/amd64/drone-pypi
```

## Docker

Build the Docker image with the following commands:

```Shell
docker build \
  --label org.label-schema.build-date=$(date -u +"%Y-%m-%dT%H:%M:%SZ") \
  --label org.label-schema.vcs-ref=$(git rev-parse --short HEAD) \
  --file docker/Dockerfile.linux.amd64 --tag plugins/pypi .
```

## Usage

```Shell
docker run --rm \
  -e PLUGIN_USERNAME=jdoe \
  -e PLUGIN_PASSWORD=my_secret \
  -e PLUGIN_SKIP_BUILD=false \
  -v $(pwd):$(pwd) \
  -w $(pwd) \
  plugins/pypi
```
