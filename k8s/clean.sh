#!/bin/bash

minikube stop
minikube delete --all
yes | docker system prune -a
yes | docker volume prune