## Pre requirements:
Here are the minimum dependencies required to run the application:
1. golang 1.18+
2. gomocks: `go install github.com/golang/mock/mockgen@v1.6.0`
3. docker-compose

## How to:

### Mock services
Run  
`$ make testmocks`

### Run unit-tests
Run  
`$ make test`

### Run application with postgresql database
Run  
`$ make start-app-postgresql`

## Project structure:
1. /cmd
   - /main.go - main application starter
2. /internal - contains the application codebase
   - /mocks - generated mocks which is described in Makefile:testmocks
3. /migrations - contains database migrations
4. /scripts - useful scripts
5. /pkg - packages
6. version
   - version.sh - scripts for increasing the version of the application that is called from the pipeline
   - version - contains the current version of the application

