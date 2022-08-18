#!/bin/bash 

# Requires Docker
rm -rf out/*
docker run --rm \
  -v "${PWD}:/local" \
  -v "${PWD}/../spec:/spec" \
  openapitools/openapi-generator-cli generate \
    -i /spec/openapi.yaml \
    -g go \
    -o /local/out/go

docker run --rm \
  -v "${PWD}:/local" \
  -v "${PWD}/../spec:/spec" \
  openapitools/openapi-generator-cli generate \
    -i /spec/openapi.yaml \
    -g go-server \
    -o /local/out/go-server

docker run --rm \
  -v "${PWD}:/local" \
  -v "${PWD}/../spec:/spec" \
  openapitools/openapi-generator-cli generate \
    -i /spec/openapi.yaml \
    -g go-echo-server \
    -o /local/out/go-echo-server


# oapi-generator
# go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@master 
