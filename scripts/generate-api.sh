#!/usr/bin/env bash

# Declare variables.
serviceName=$1
serviceVersion=$2

# Check if the service name specified.
if [ "$serviceName" = '' ]; then
    echo "service name required"
    exit 1
fi

# Check if the service version specified.
if [ "$serviceVersion" = '' ]; then
    echo "service version not found, default value v1 applied"
    serviceVersion="v1"
fi

if [ "$serviceName" = 'self' ]; then
    # Generate self api's
    export PROTO_TARGET_DIRECTORY="./api/proto/v1"
    docker-compose -f docker-compose.protoc.yml up --remove-orphans
else
    # Define target.
    target="./third_party/mercuryapis/github.com.dochq/$serviceName/$serviceVersion/"
    # rm -rf $target
    # mkdir -p $target
    # chmod -R 777 $target

    # Copy proto files from the cloud storage.
    # gsutil -m cp -r gs://dochq-protos/$serviceName/$serviceVersion/*.proto $target

    # Generate api's
    export PROTO_TARGET_DIRECTORY="$target"
    docker-compose -f docker-compose.protoc.yml up --remove-orphans
fi
