#!/bin/bash

docker network create -d bridge mongo-network
docker pull mongo:latest
mkdir /mnt/mongodb
docker docker run --name mongodb -e MONGO_INITDB_ROOT_USERNAME=admin -e MONGO_INITDB_ROOT_PASSWORD=ABcd_1234 -v /mnt/mongodb:/data/db --network mongo-network -p 27017:27017 -itd mongo:latest

docker pull mongo-express:latest
docker run --name mongo_express --network mongo-network -e ME_CONFIG_MONGODB_SERVER=mongodb -e ME_CONFIG_MONGODB_ADMINUSERNAME=admin -e ME_CONFIG_MONGODB_ADMINPASSWORD=ABcd_1234 -e ME_CONFIG_BASICAUTH_USERNAME=admin -e ME_CONFIG_BASICAUTH_PASSWORD=ABcd_1234 -p 8081:8081 -itd mongo-express:latest