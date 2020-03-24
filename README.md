# MatterhornApiService
Web API service for Matterhorn implemented with Go.

## Requirements
To build and run MatterhornApiService locally you will need:
* Go 1.13 or higher
* golang-migrate CLI tool (installed using Homebrew on macOS or Scoop on Windows)
* Docker Desktop (used to host a local instance of mysql-server and run openapi-generator)

## Setup
This section documents how to set up your environment to run MatterhornApiService.

### Go Runtime
MatterhornApiService is built to run on Go v1.13.

### Configuring Environment Variables
`MATTERHORN_ENV` - Defines the environment configuration to be used. Expected values include `dev` and `prod`. This determines the database endpoint and schema name is used.

`MATTERHORN_DB_PASSWORD` - Specifies the password used to connect to the database instance. User name is expected to be "admin" for now.

### Build
Use the Go CLI to build the application
`go build`

### Run
Use the Go CLI to run the application.
`go run .`

### Test
Use the Go CLI to execute unit tests.
`go test`

## Adding new APIs
MatterhornApiService follows a spec-first approach to defining API endpoints. To add, remove, or modify an API, first update the OpenAPI spec in `api/openapi.yaml`. Next, execute the VSCode task named "Generate Server Stubs". This will regenerate server stubs based on the updated OpenAPI spec without overwriting any existing implementation (as configured in `.openapi-generator-ignore`). You will need to modify `main.go` to create your *Service and *Controller structs and initialize them with any injected dependencies.

## CI/CD
MatterhornApiService is configured to deploy to AWS Elastic Beanstalk. `Buildfile`, and `Procfile` are provided to describe how to run and build the application to Elastic Beanstalk. Before running the application, Elastic Beanstalk will execute the `build.sh` script; this is done instead of deploying a pre-compiled binary because it needs to be compiled for the OS and architecture of the target host.

GitHub Actions for this repository are configured to automatically deploy to Matterhorn's production Elastic Beanstalk environment in response to commits being pushed to `master`.

## Database Migrations
MatterhornApiService uses [golang-migrate](https://github.com/golang-migrate/migrate) to execute database migrations on startup.

VSCode tasks are defined for migrating the development database or rolling back the most recently applied migration. These tasks require both that the `MATTERHORN_DB_PASSWORD` environment variable is set and that the golang-migrate CLI tool is installed. 

You can install this tool on Windows by first installing the [Scoop](https://scoop.sh/) package manager and then running `scoop install migrate`. On macOS, use homebrew and run `brew install golang-migrate`.