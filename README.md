# MatterhornApiService
Web API service for Matterhorn implemented with Go.

## Setup

### Database
MatterhornApiService expects to connect to a MySQL-compatible database instance. You will need to set some environment variables to define how to connect.
`MATTERHORN_DB_ENDPOINT` - Specifies the host name of the database instance, including the port. E.g., http://mydatabase.com:3306.
`MATTERHORN_DB_PASSWORD` - Specifies the password used to connect to the database instance. User name is expected to be "admin" for now.

### Compile
`go build`

### Run
`go run`

### Unit Tests
`go test`

## CI/CD

MatterhornApiService is configured to deploy to AWS Elastic Beanstalk. `Buildfile`, and `Procfile` are provided to describe how to run and build the application to Elastic Beanstalk. Before running the application, Elastic Beanstalk will execute the `build.sh` script; this is done instead of deploying a pre-compiled binary because it needs to be compiled for the OS and architecture of the target host.

GitHub Actions for this repository are configured to automatically deploy to Matterhorn's production Elastic Beanstalk environment in response to commits being pushed to `master`.