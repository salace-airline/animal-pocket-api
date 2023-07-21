#!/bin/bash
cd /home/ec2-user/animalpocketresources/api
docker-compose build --no-cache
docker-compose up -d
