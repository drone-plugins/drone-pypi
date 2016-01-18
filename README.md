# drone-pypi

[![Build Status](http://beta.drone.io/api/badges/drone-plugins/drone-pypi/status.svg)](http://beta.drone.io/drone-plugins/drone-pypi)
[![Coverage Status](https://aircover.co/badges/drone-plugins/drone-pypi/coverage.svg)](https://aircover.co/drone-plugins/drone-pypi)
[![](https://badge.imagelayers.io/plugins/drone-pypi:latest.svg)](https://imagelayers.io/?images=plugins/drone-pypi:latest 'Get your own badge on imagelayers.io')

Drone plugin to publish a Python package to a PyPI server

## Binary

Build the binary using `make`:

```
make deps build
```

### Example

```sh
./drone-pypi <<EOF
{
    "repo": {
        "clone_url": "git://github.com/drone/drone",
        "owner": "drone",
        "name": "drone",
        "full_name": "drone/drone"
    },
    "system": {
        "link_url": "https://beta.drone.io"
    },
    "build": {
        "number": 22,
        "status": "success",
        "started_at": 1421029603,
        "finished_at": 1421029813,
        "message": "Update the Readme",
        "author": "johnsmith",
        "author_email": "john.smith@gmail.com"
        "event": "push",
        "branch": "master",
        "commit": "436b7a6e2abaddfd35740527353e78a227ddcb2c",
        "ref": "refs/heads/master"
    },
    "workspace": {
        "root": "/drone/src",
        "path": "/drone/src/github.com/drone/drone"
    },
    "vargs": {
        "repository": "https://pypi.example.com",
        "distributions": [
            "sdist",
            "bdist_wheel"
        ],
        "username": "guido",
        "password": "secret"
    }
}
EOF
```

## Docker

Build the container using `make`:

```
make deps docker
```

### Example

```sh
docker run -i plugins/drone-pypi <<EOF
{
    "repo": {
        "clone_url": "git://github.com/drone/drone",
        "owner": "drone",
        "name": "drone",
        "full_name": "drone/drone"
    },
    "system": {
        "link_url": "https://beta.drone.io"
    },
    "build": {
        "number": 22,
        "status": "success",
        "started_at": 1421029603,
        "finished_at": 1421029813,
        "message": "Update the Readme",
        "author": "johnsmith",
        "author_email": "john.smith@gmail.com"
        "event": "push",
        "branch": "master",
        "commit": "436b7a6e2abaddfd35740527353e78a227ddcb2c",
        "ref": "refs/heads/master"
    },
    "workspace": {
        "root": "/drone/src",
        "path": "/drone/src/github.com/drone/drone"
    },
    "vargs": {
        "repository": "https://pypi.example.com",
        "distributions": [
            "sdist",
            "bdist_wheel"
        ],
        "username": "guido",
        "password": "secret"
    }
}
EOF
```
