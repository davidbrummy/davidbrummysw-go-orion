#!/bin/bash

docker build . -t davidbrummysw-go-orion-service1.0

docker run -it -p 80:8080 --env-file dev.env --network orion --name davidbrummysw-go-orion-service davidbrummysw-go-orion-service1.0       