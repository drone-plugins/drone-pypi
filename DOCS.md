Use the PyPI plugin to deploy a Python package to a PyPI server. You can
override the default configuration with the following parameters:

* `repository` - The repository name (optional)
* `username` - The username to login with (optional)
* `password` - A password to login with (optional)
* `distributions` - A list of distribution types to deploy (optional)

## Example

The following is a sample configuration in your .drone.yml file:

```yaml
publish:
  pypi:
    repository: https://pypi.python.org/pypi
    username: guido
    password: secret
    distributions:
      - sdist
      - bdist_wheel
```

### Upload `bdist_wheel`

In some cases you may want to upload a wheel additionally to the source
distribution to some PyPI server, for this use case we expose the following
additional parameter:

* `distributions` - List of distribution types

Example configuration that uploads `sdist` and `bdist_wheel`:

```yaml
publish:
  pypi:
    username: guido
    password: secret
    distributions:
      - sdist
      - bdist_wheel
```

### Private registry

In some cases you may want to upload a package to a private PyPI server, for
this use case we expose the following additional parameter:

* `repository` - Repository URL

Example configuration that uploads to a private PyPI server:

```yaml
publish:
  pypi:
    repository: https://pypi.example.com
    username: guido
    password: secret
```
