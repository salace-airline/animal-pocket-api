#!/bin/bash
sudo yum update -y
sudo dnf update -y
sudo yum install -y libxcrypt-compat
sudo yum install docker -y
sudo curl -L https://github.com/docker/compose/releases/download/3.8.0/docker-compose-`uname -s`-`uname -m` \
    >~/docker-compose
chmod +x ~/docker-compose
sudo mv ~/docker-compose /usr/local/bin/docker-compose
sudo service docker start
sudo usermod -a -G docker ec2-user
sudo chmod 666 /var/run/docker.sock
sudo ln -s /usr/local/bin/docker-compose /usr/bin/docker-compose
sudo service docker restart
