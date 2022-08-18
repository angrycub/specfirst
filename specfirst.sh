#!/bin/bash

if [ "$1" == "start" ]; then
# Run APICurito Container
echo "==> Starting Apicurito"
docker run --rm --name=apicurito -d -p 8080:8080 apicurio/apicurito-ui

# Run Nomad Dev Agent
echo "==> Starting Nomad Dev Agent"
nomad agent -dev > /tmp/specfirst.nomad.log 2>&1 &
nomadPID=$!
echo ${nomadPID} > /tmp/specfirst.nomad.pid
disown ${nomadPID}

# Run Mockoon
echo "==> Starting Mockoon"
docker run -d --name=mockoon \
--mount type=bind,source=$(pwd)/environments,target=/environments,readonly \
--mount type=bind,source=$(pwd)/spec,target=/spec,readonly \
-p 5656:5656 mockoon/cli:latest --data /environments/mockoon.json --port 5656

# Open Webpage
open http://127.0.0.1:5656

fi

if [ "$1" == "stop" ]; then
docker stop apicurito
docker stop mockoon
nomadPID=$(cat /tmp/specfirst.nomad.pid)

kill $nomadPID
killEC=$?
if [ $killEC -ne 0 ]; then
    echo "Couldn't kill Nomad dev agent (pid ${nomadPID})..."
else
    rm -f /tmp/specfirst.nomad.pid
fi

fi
