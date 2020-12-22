#!/usr/bin/env bash

# todo(yumcoder) change dc ip
# sed -i '/ipAddress = /c\ipAddress = 127.0.0.1' a.txt
# todo(yumcoder) change folder path for nbfs

docker start mysql-docker redis-docker etcd-docker

PluralityServer="$GOPATH/src/github.com/PluralityNET/PluralityServer"


echo "build document ..."
cd ${PluralityServer}/service/document
go build
./document &
sleep 1

echo "build auth_session ..."
cd ${PluralityServer}/service/auth_session
go build
./auth_session &
sleep 1

echo "build sync ..."
cd ${PluralityServer}/messenger/sync
go build
./sync &
sleep 1

echo "build upload ..."
cd ${PluralityServer}/messenger/upload
go build
./upload &
sleep 1


echo "build auth_key ..."
cd ${PluralityServer}/access/auth_key
go build
./auth_key &

echo "build biz_server ..."
cd ${PluralityServer}/messenger/biz_server
go build
./biz_server &
sleep 1

echo "build session ..."
cd ${PluralityServer}/access/session
go build
./session &
sleep 1

echo "build frontend ..."
cd ${PluralityServer}/access/frontend
go build
./frontend &
sleep 1

echo "***** wait *****"
wait
