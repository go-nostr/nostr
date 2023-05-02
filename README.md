# @go-nostr/nostr

⚠️ **WARNING: This project is actively under development and no stable releases have been made yet.** ⚠️

Please be aware that the code and documentation may contain unfinished features or inconsistencies. We appreciate your interest in the project and encourage you to check back for updates on progress towards a stable release.

Roadmap:

| Status      | Milestone       | Description                                                     |
|-------------|-----------------|-----------------------------------------------------------------|
| In-progress | `nostr`         | core package for use by external applications and executable    |
| In-progress | `internal/cli`  | command-line client interface for Nostr client                  |
| To-do       | `internal/cmd`  | relay server with persistent data repositories                  |
| To-do       | `internal/docs` | documentation site built with Hugo                              |
| To-do       | `internal/web`  | web client and relay administrator interface built with Angular |

## Overview

All-in-one Angular/Go/Hugo monorepo for the Nostr protocol, featuring comprehensive documentation, an all-in-one Go executable and reusable packages with an Angular web client.

## Getting Started

### Prerequisites

Before starting development, make sure to install the following dependencies:

- [Install Angular](https://angular.io/guide/setup-local)
- [Install Docker](https://docs.docker.com/engine/install/)
- [Install Go](https://go.dev/doc/install)
- [Install Hugo](https://gohugo.io/installation/)
- [Install NPM and NodeJS](https://docs.npmjs.com/downloading-and-installing-node-js-and-npm)

### Installation

After prerequisites have been installed, use the following shell commands from the project directory to install dependencies:

#### Install Go dependencies

Fetch and install all Go dependencies for the project. The `-d` flag ensures that the packages are downloaded only, without installing them.

```shell
go get -d ./...
```

#### Install NPM dependencies

Install all NPM dependencies for the monorepo workspaces. The -ws flag ensures that the installation is performed across all workspaces.

```shell
npm i -ws
```

## Usage

To instantly set up and run the Nostr application, simply execute the following Docker command. This will pull the latest Nostr image, map the necessary ports, and get your local instance up and running:

```shell
docker run -p 3001:3001 -p 4200:4200 ghcr.io/go-nostr/nostr:latest
```

### Build Docker image

Create a Docker image for the application by using the Dockerfile in the repository. This command will build the image and tag it with the name 'nostr'.

```shell
docker build -t nostr -f internal/config/Dockerfile .
```

### Build Docker services

Build all the services defined in the docker-compose.yml file. This command will create Docker images for each service and store them locally.

```shell
docker-compose -f ./internal/config/docker-compose.yml build --no-cache
```

### Build Hugo site

Build the Hugo documentation site maintained in the `internal/docs` directory. This command will build the Hugo the documentation site for distribution, placing the output in the `internal/docs/public` directory.

```shell
hugo -s internal/docs    # exclude drafts
```

```shell
hugo -s internal/docs -D # include drafts
```

### Build NPM packages

Compile and build all NPM packages in the monorepo, considering their interdependencies. The -ws flag ensures that the build is performed across all workspaces.

```shell
npm run build -ws
```

### Format Go files

Format all Go source files in the repository by running the Go formatter. This ensures consistent coding style across the project.

```shell
go fmt ./...
```

### Generate Go dependencies

Automatically generate code for Go dependencies, such as mocks, based on the source files. This command should be executed whenever there are changes to the Go source files.

```shell
go generate ./...
```

### Run Docker image

Run the Docker image defined in the Dockerfile file. This command will run the image and display the logs in the console.

```shell
docker run -p 3001:3001 -p 4200:4200 nostr
```

### Run Docker services

Start all the Docker services defined in the docker-compose.yml file. This command will run the containers and display the logs in the console.

```shell
docker-compose -f ./internal/config/docker-compose.yml up --build
```

### Run Go tests

Execute all Go tests in the repository, including unit and integration tests. This command will also display a summary of the test results.

```shell
go test ./...
```

### Run Hugo server

Build and serve the Hugo documentation site maintained in the `internal/docs` directory. This command will start the Hugo server and generate the documentation site, which can be accessed in a web browser.

```shell
hugo server -s internal/docs    # exclude drafts
```

```shell
hugo server -s internal/docs -D # include drafts
```

### Run NPM tests

Run all NPM tests in the monorepo, considering their interdependencies. The -ws flag ensures that the tests are executed across all workspaces.

```shell
npm test -ws
```

## Contributing

If you're an Angular or Go developer looking to help advance the Nostr protocol, we'd love your help! To get started, make sure you're familiar with the Nostr protocol and comfortable [creating pull requests](https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/proposing-changes-to-your-work-with-pull-requests/creating-a-pull-request) and [issues](https://docs.github.com/en/issues/tracking-your-work-with-issues/creating-an-issue).

Before contributing, please take a moment to familiarize yourself with our project's vision by reading our [VISION](./VISION.md) document. We also suggest contributors follow our [CODE_OF_CONDUCT](./CODE_OF_CONDUCT.md) to maintain good vibes.

## Deployment

The deployment process is handled by GitHub Actions, with separate workflows for Angular, Docker, Go, and Hugo. When a push is made, the Angular workflow runs, building the application, running tests, and linting the code. The built Angular application is then archived into a dist.tar file, which is uploaded as a build artifact.

On pushes to the main branch, the Docker workflow builds and pushes a Docker image to the GitHub Container Registry (GHCR), supporting both amd64 and arm64 platforms. The Go workflow runs in parallel, checking Go code formatting, building NPM packages, and running Go tests. The Hugo workflow, also triggered on pushes to the main branch, builds NPM packages and generates the Hugo site, uploading the artifacts and deploying them to GitHub Pages. This process ensures the codebase remains up-to-date and deployable, making it easy for current and future contributors to understand the project's deployment process.

| Badge                                                                                                                                                    | Description                                                                        |
|----------------------------------------------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------|
| [![Angular](https://github.com/go-nostr/nostr/actions/workflows/angular.yml/badge.svg)](https://github.com/go-nostr/nostr/actions/workflows/angular.yml) | Builds, tests, and lints the Angular application, and uploads the build artifacts. |
| [![Docker](https://github.com/go-nostr/nostr/actions/workflows/docker.yml/badge.svg)](https://github.com/go-nostr/nostr/actions/workflows/docker.yml)    | Builds and pushes Docker images to the GitHub Container Registry.                  |
| [![Go](https://github.com/go-nostr/nostr/actions/workflows/go.yml/badge.svg)](https://github.com/go-nostr/nostr/actions/workflows/go.yml)                | Checks Go code formatting, generates Go dependencies, and runs Go tests.           |
| [![Hugo](https://github.com/go-nostr/nostr/actions/workflows/hugo.yml/badge.svg)](https://github.com/go-nostr/nostr/actions/workflows/hugo.yml)          | Builds NPM packages, generates the Hugo site, and deploys to GitHub Pages.         |

## Built-with

This project is built with the following core technologies:

- [Angular](https://angular.io/)
- [Docker](https://docker.com/)
- [Go](https://go.dev/)
- [Hugo](https://gohugo.io/)
- [NPM](https://www.npmjs.com/)
- [NodeJS](https://nodejs.org/en)

## Acknowledgements

- [@peaceiris](https://github.com/peaceiris) for the [Hugo GitHub Action](https://github.com/peaceiris/actions-hugo)

## Authors

- [@chantzlarge](https://github.com/chantzlarge)
- [@vavrajosef](https://github.com/vavrajosef)

## Versioning

This project uses [semantic versioning](https://semver.org) and [GitHub](https://docs.github.com/en/repositories/releasing-projects-on-github/managing-releases-in-a-repository) to publish new versions and manage releases.

## License

This project is licensed under the terms of the [LICENSE](./LICENSE).

## References

For additional materials helpful for contributors or users, please see:

- [@nostr-protocol/nips](https://github.com/nostr-protocol/nips)
- [@nhooyr/websocket](https://github.com/nhooyr/websocket)
