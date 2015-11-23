from setuptools import setup
import subprocess


def get_version():
    try:
        git = subprocess.Popen(
            ["git", "describe", "--long"],
            stdout=subprocess.PIPE,
            stderr=subprocess.PIPE)
    except Exception:
        return "0.0.0"
    val = git.communicate()[0]
    if git.returncode != 0:
        return "0.0.0"
    l = val.strip().split("-")
    return l[0] + "." + l[1]


setup(
    name="drone-pypi",
    version=get_version(),
    description="Module for testing drone-pypi.",
    url="http://github.com/drone-plugins/drone-pypi",
    packages=["drone_pypi"],
    maintainer="Drone Contributors",
    maintainer_email="support@drone.io",
)
