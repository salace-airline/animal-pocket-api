#!/bin/bash

# get Github variables
export DB_USER=$DB_USER
export DB_PASSWORD=$DB_PASSWORD
export DB_NAME=$DB_NAME

# run
cd /home/ec2-user/animalpocketresources/api
docker-compose build --no-cache
docker-compose up -d