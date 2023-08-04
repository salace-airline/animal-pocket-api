#!/bin/bash
cd /home/ec2-user/animalpocketresources/api
envsubst < docker-compose.yml.template > docker-compose.yml
docker-compose build --no-cache
docker-compose up -d
