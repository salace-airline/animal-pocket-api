#!/bin/bash

# get Github variables
export DB_USER=$SECRET_DB_USER
export DB_PASSWORD=$SECRET_DB_PASSWORD
export DB_NAME=$SECRET_DB_NAME
envsubst < docker-compose.yml.template > docker-compose.yml

# run
cd /home/ec2-user/animalpocketresources/api
docker-compose build --no-cache
docker-compose up -d
