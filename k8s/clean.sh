#!/bin/bash

rm -rf ~/.minikube
minikube stop
minikube delete --all
yes | docker system prune -a
yes | docker volume prune
