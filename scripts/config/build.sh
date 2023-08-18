#!/bin/bash

PWD2=`pwd`
echo $PWD2
PluralityServer="$GOPATH/src/github.com/fedigram/fedigram-server"

echo "build document ..."
cd ${PluralityServer}/service/document
go build

echo "build auth_session ..."
cd ${PluralityServer}/service/auth_session
go build

echo "build sync ..."
cd ${PluralityServer}/messenger/sync
go build

echo "build upload ..."
cd ${PluralityServer}/messenger/upload
go build


echo "build auth_key ..."
cd ${PluralityServer}/access/auth_key
go build

echo "build biz_server ..."
cd ${PluralityServer}/messenger/biz_server
go build

echo "build session ..."
cd ${PluralityServer}/access/session
go build

echo "build frontend ..."
cd ${PluralityServer}/access/frontend
go build

cd ${PWD2}
