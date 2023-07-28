#!/bin/bash
cd /home/ec2-user/animalpocketresources/api
sudo pacman -S libxcrypt-compat
docker-compose build --no-cache
docker-compose up -d
