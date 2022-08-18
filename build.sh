#!/bin/bash

echo "Making the rfc artifacts..."
cd rfc
./widderdoc.sh
cd ..

echo "Making the Golang generated code..."
cd golang
./openapi-generator.sh
cd ..

