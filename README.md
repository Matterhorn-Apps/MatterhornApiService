# MatterhornApiService
Web API service for Matterhorn implemented with Go.

## Setup

### Compile
`go build .`

### Run
`go run server.go`

## CI/CD

MatterhornApiService is configured to deploy to AWS Elastic Beanstalk. `Buildfile`, and `Procfile` are provided to describe how to run and build the application to Elastic Beanstalk. Before running the application, Elastic Beanstalk will execute the `build.sh` script; this is done instead of deploying a pre-compiled binary because it needs to be compiled for the OS and architecture of the target host.

GitHub Actions for this repository are configured to automatically deploy to Matterhorn's production Elastic Beanstalk environment in response to commits being pushed to `master`.