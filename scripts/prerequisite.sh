#!/usr/bin/env bash

# install docker
# Install Docker for Ubuntu ---see---> https://docs.docker.com/install/linux/docker-ce/ubuntu/
#   for Ubuntu 18.04 ---see---> https://linuxconfig.org/how-to-install-docker-on-ubuntu-18-04-bionic-beaver
# Install Docker for Mac ---see---> https://docs.docker.com/docker-for-mac/install/
# Install Docker for Windows ---see---> https://docs.docker.com/docker-for-windows/install/#start-docker-for-windows

echo "run etcd-docker..."
docker run --name etcd-docker -d -p 2379:2379 -p 2380:2380 appcelerator/etcd
echo "run mysql-docker..."
docker run --name mysql-docker -p 3306:3306 -e MYSQL_ALLOW_EMPTY_PASSWORD=yes -d mysql:5.7
echo "run redis-docker..."
docker run --name redis-docker -p 6379:6379 -d redis
echo "clone PluralityServer..."
mkdir ${GOPATH}/src/github.com/PluralityNET/
cd ${GOPATH}/src/github.com/PluralityNET/
git clone https://github.com/fedigram/fedigram-server.git

echo "create db schema ..."
docker exec -it mysql-docker sh -c 'exec mysql -u root -p -e"CREATE DATABASE PluralityServer;"'
docker exec -i mysql-docker mysql --user=root PluralityServer < ${GOPATH}/src/github.com/fedigram/fedigram-server/scripts/PluralityServer.sql
echo "OK"
