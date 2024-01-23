Pre requirements:
Here are the minimum dependencies required to run the application:

golang 1.18+
gomocks: go install github.com/golang/mock/mockgen@v1.6.0
docker-compose
How to:
Mock services
Run
$ make testmocks

Run unit-tests
Run
$ make test

Run application with postgresql database
Run
$ make start-app-postgresql

Project structure:
/cmd
/main.go - main application starter
/internal - contains the application codebase
/mocks - generated mocks which is described in Makefile:testmocks
/migrations - contains database migrations
/scripts - useful scripts
/pkg - packages
version
version.sh - scripts for increasing the version of the application that is called from the pipeline
version - contains the current version of the application
