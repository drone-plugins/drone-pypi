local PipelineTesting = {
  kind: "pipeline",
  name: "testing",
  platform: {
    os: "linux",
    arch: "amd64",
  },
  services: [
      {
          name: "pypiserver",
          image: "pypiserver/pypiserver",
          pull: "always",
          entrypoint: [
              "pypi-server",
              "-P",
              ".",
              "-a",
              ".",
              "-p",
              "8080",
              "/data/packages"
          ],
      },
  ],
  steps: [
    {
      name: "vet",
      image: "golang:1.11",
      pull: "always",
      environment: {
        GO111MODULE: "on",
      },
      commands: [
        "go vet ./...",
      ],
    },
    {
      name: "test",
      image: "golang:1.11-alpine",
      pull: "always",
      environment: {
        GO111MODULE: "on",
        PLUGIN_REPOSITORY: "http://pypiserver:8080/",
        PLUGIN_DISTRIBUTIONS: "sdist",
        PLUGIN_USERNAME: "demo",
        PLUGIN_PASSWORD: "demo",
      },
      commands: [
        "apk --no-cache add -U python3 git",
        "pip3 install --no-cache-dir --upgrade pip setuptools wheel six twine",
        "go test -cover ./...",
      ],
    },
  ],
  trigger: {
    branch: [ "master" ],
  },
};

local PipelineBuild(os="linux", arch="amd64") = {
  kind: "pipeline",
  name: os + "-" + arch,
  platform: {
    os: os,
    arch: arch,
  },
  steps: [
    {
      name: "build-push",
      image: "golang:1.11",
      pull: "always",
      environment: {
        CGO_ENABLED: "0",
        GO111MODULE: "on",
      },
      commands: [
        "go build -v -ldflags \"-X main.build=${DRONE_BUILD_NUMBER}\" -a -o release/" + os + "/" + arch + "/drone-pypi",
      ],
      when: {
        event: [ "push", "pull_request" ],
      },
    },
    {
      name: "build-tag",
      image: "golang:1.11",
      pull: "always",
      environment: {
        CGO_ENABLED: "0",
        GO111MODULE: "on",
      },
      commands: [
        "go build -v -ldflags \"-X main.version=${DRONE_TAG##v} -X main.build=${DRONE_BUILD_NUMBER}\" -a -o release/" + os + "/" + arch + "/drone-pypi",
      ],
      when: {
        event: [ "tag" ],
      },
    },
    {
      name: "executable",
      image: "golang:1.11",
      pull: "always",
      commands: [
        "./release/" + os + "/" + arch + "/drone-pypi --help",
      ],
    },
    {
      name: "dryrun",
      image: "plugins/docker:" + os + "-" + arch,
      pull: "always",
      settings: {
        dry_run: true,
        tags: os + "-" + arch,
        dockerfile: "docker/Dockerfile." + os + "." + arch,
        repo: "plugins/pypi",
        username: { "from_secret": "docker_username" },
        password: { "from_secret": "docker_password" },
      },
      when: {
        event: [ "pull_request" ],
      },
    },
    {
      name: "publish",
      image: "plugins/docker:" + os + "-" + arch,
      pull: "always",
      settings: {
        auto_tag: true,
        auto_tag_suffix: os + "-" + arch,
        dockerfile: "docker/Dockerfile." + os + "." + arch,
        repo: "plugins/pypi",
        username: { "from_secret": "docker_username" },
        password: { "from_secret": "docker_password" },
      },
      when: {
        event: [ "push", "tag" ],
      },
    },
  ],
  depends_on: [
    "testing",
  ],
  trigger: {
    branch: [ "master" ],
  },
};

local PipelineNotifications = {
  kind: "pipeline",
  name: "notifications",
  platform: {
    os: "linux",
    arch: "amd64",
  },
  steps: [
    {
      name: "manifest",
      image: "plugins/manifest:1",
      pull: "always",
      settings: {
        username: { "from_secret": "docker_username" },
        password: { "from_secret": "docker_password" },
        spec: "docker/manifest.tmpl",
        ignore_missing: true,
      },
    },
    {
      name: "microbadger",
      image: "plugins/webhook:1",
      pull: "always",
      settings: {
        url: { "from_secret": "microbadger_url" },
      },
    },
  ],
  depends_on: [
    "linux-amd64",
    "linux-arm64",
    "linux-arm",
  ],
  trigger: {
    branch: [ "master" ],
    event: [ "push", "tag" ],
  },
};

[
  PipelineTesting,
  PipelineBuild("linux", "amd64"),
  PipelineBuild("linux", "arm64"),
  PipelineBuild("linux", "arm"),
  PipelineNotifications,
]
