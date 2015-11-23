Use the PyPI plugin to deploy a Python package to a PyPI server.

* `repository` - The repository name (optional)
* `username` - The username to login with (optional)
* `password` - A password to login with (optional)
* `distributions` - A list of distribution types to deploy (optional)

The following is an example configuration for your .drone.yml:

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
