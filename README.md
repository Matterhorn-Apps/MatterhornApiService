# MatterhornApiService
![](https://github.com/Matterhorn-Apps/MatterhornApiService/workflows/build/badge.svg)
Web API service for Matterhorn implemented with Go.

## Requirements
To build and run MatterhornApiService locally you will need:
* Go 1.14 or higher
* golang-migrate CLI tool (used to execute database migrations - see installation instructions below)
* Docker Desktop (used to host a local instance of mysql-server)

## Setup
This section documents how to set up your environment to run MatterhornApiService.

### Go Runtime
MatterhornApiService is built to run on Go v1.14.

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
MatterhornApiService follows a spec-first approach to defining API endpoints using GraphQL. 

To add, remove, or modify an API, first add any new types to the schema document `graph/schema.graphqls`. Next, run `go generate ./...` to generate new and updated Go types and resolver implementation stubs. Then implement any new resolvers in `graph/schema.resolvers.go`. If you need to define a modified Go type to use in place of the one that was auto-generated, you can define it in `graph/model`, as has been done for `graph/model/user.go`. One useful reason to do this is to replace a reference to a "child" entity with an ID; this will trigger generation of a dependent resolver that will only execute to retrieve the corresponding field if it is requested.

## CI/CD
MatterhornApiService is configured to deploy to AWS Elastic Beanstalk.

GitHub Actions for this repository are configured to automatically deploy to Matterhorn's production Elastic Beanstalk environment in response to commits being pushed to `master`.

## Database Migrations
MatterhornApiService uses [golang-migrate](https://github.com/golang-migrate/migrate) to execute database migrations on startup.

VSCode tasks are defined for migrating the development database or rolling back the most recently applied migration. These tasks require both that the `MATTERHORN_DB_PASSWORD` environment variable is set and that the golang-migrate CLI tool is installed. 

You can install this tool on Windows by first installing the [Scoop](https://scoop.sh/) package manager and then running `scoop install migrate`. On macOS, use homebrew and run `brew install golang-migrate`.