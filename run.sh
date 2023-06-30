#!/bin/bash
cd /home/ec2-user/animalpocketresources
docker-compose build --no-cache
docker-compose up -d