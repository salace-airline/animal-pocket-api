#!/bin/bash
sudo yum install docker -y
python3-pip install docker-compose
sudo dnf -y install docker-compose
sudo service docker start
sudo usermod -a -G docker ec2-user
sudo chmod 666 /var/run/docker.sock
sudo ln -s /usr/local/bin/docker-compose /usr/bin/docker-compose
sudo service docker start
