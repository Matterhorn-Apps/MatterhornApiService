# MatterhornApiService
Web API service for Matterhorn implemented with Go.

## Setup
This section documents how to set up your environment to run MatterhornApiService.

### Go Runtime
MatterhornApiService is built to run on Go v1.14.

### Configuring Environment Variables
`MATTERHORN_ENV`            - Defines the environment configuration to be used. Expected values include `dev` and `prod`. This determines the database endpoint and schema name is used.
`MATTERHORN_DB_PASSWORD`    - Specifies the password used to connect to the database instance. User name is expected to be "admin" for now.

### Compile
`go build`

### Run
`go run .`

### Test
`go test`

## CI/CD
MatterhornApiService is configured to deploy to AWS Elastic Beanstalk. `Buildfile`, and `Procfile` are provided to describe how to run and build the application to Elastic Beanstalk. Before running the application, Elastic Beanstalk will execute the `build.sh` script; this is done instead of deploying a pre-compiled binary because it needs to be compiled for the OS and architecture of the target host.

GitHub Actions for this repository are configured to automatically deploy to Matterhorn's production Elastic Beanstalk environment in response to commits being pushed to `master`.

## Database Migrations
MatterhornApiService uses [https://github.com/golang-migrate/migrate](golang-migrate) to execute database migrations on startup.