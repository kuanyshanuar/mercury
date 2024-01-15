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

## Database migrations:

### PostgreSQL base image:
    - https://gitlab.com/zharzhanov/postgresql.data-migration

## PostgreSQL
`docker run -d \
         -p 5433:5432 \
         --network mercury \
         --name postgres \
         -v ~/postgres_storage:/var/lib/postgresql/data \
         -e POSTGRES_DB=postgres \
         -e POSTGRES_USER=postgres \
         -e POSTGRES_PASSWORD=Lbr6QgHZUdz8sTQXFRPczczpupDAXTAX \
         postgres:14.2`

## Minio
`docker run -d --restart=on-failure:3
            -p 9000:9000 \
            -p 9090:9090 \
            --network mercury \
            --name minio \
            -v ~/minio/data:/data \
            -e "MINIO_ROOT_USER=mercury" \
            -e "MINIO_ROOT_PASSWORD=minio123" \
            quay.io/minio/minio server /data --console-address ":9090"`

## Redis
`docker run -d --restart=on-failure:3 -p 6379:6379 --name redis --network=mercury --rm redis redis-server --save 20 1 --loglevel warning --requirepass eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81`