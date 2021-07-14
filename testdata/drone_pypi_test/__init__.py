# Only version of setuptools<45 evaluates the code of module to read the
# __version__ attribute during the build operation.
# setuptools>=45 uses ast (Abstract Syntax Trees) to read the __version__
# attribute without evaluate the code in order to prevent error with
# third party import

import setuptools


if tuple((int(e) if e.isdigit() else e) for e in setuptools.__version__.split('.')) < (45, 0, 0):
    raise Exception("The drone-plugin should use setuptools version >= 45")


__version__ = '0.1.0'
