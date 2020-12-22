#!/bin/bash

PWD2=`pwd`
echo $PWD2
PluralityServer="$GOPATH/src/github.com/PluralityNET/PluralityServer"

echo "build document ..."
cd ${PluralityServer}/service/document
# go build
nohup ./document >> ${PWD2}/document.log 2>&1 &
sleep 1

echo "build auth_session ..."
cd ${PluralityServer}/service/auth_session
# go build
nohup ./auth_session >> ${PWD2}/auth_session.log 2>&1 &
sleep 1

echo "build sync ..."
cd ${PluralityServer}/messenger/sync
# go build
nohup ./sync >> ${PWD2}/sync.log 2>&1 &
sleep 1

echo "build upload ..."
cd ${PluralityServer}/messenger/upload
# go build
nohup ./upload >> ${PWD2}/upload.log 2>&1 &
sleep 1


echo "build auth_key ..."
cd ${PluralityServer}/access/auth_key
# go build
nohup ./auth_key >> ${PWD2}/auth_key.log 2>&1 &
sleep 1

echo "build biz_server ..."
cd ${PluralityServer}/messenger/biz_server
# go build
nohup ./biz_server >> ${PWD2}/biz_server.log 2>&1 &
sleep 1

echo "build session ..."
cd ${PluralityServer}/access/session
# go build
nohup ./session >> ${PWD2}/session.log 2>&1 &
sleep 1

echo "build frontend ..."
cd ${PluralityServer}/access/frontend
# go build
nohup ./frontend >> ${PWD2}/frontend.log 2>&1 &
sleep 1

cd $PWD2

