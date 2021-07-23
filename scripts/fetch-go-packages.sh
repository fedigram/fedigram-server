#!/usr/bin/env bash

echo entered $0

get () {
  local cl="go mod download $1"
  echo "starting $cl..."
  $cl
  echo "completed $cl"
}

unset GOPATH

go env

get go.etcd.io/etcd@v3.5.0
get google.golang.org/grpc
#get go.etcd.io/etcd/clientv3
#get go.etcd.io/etcd/mvcc/mvccpb
#get google.golang.org/grpc/naming
get github.com/BurntSushi/toml
get github.com/bwmarrin/snowflake
get github.com/disintegration/imaging
get github.com/go-sql-driver/mysql
get github.com/gogo/protobuf/proto
get github.com/golang/protobuf/proto
get github.com/golang/glog
get github.com/golang/protobuf/ptypes/any
get github.com/gomodule/redigo/redis
get github.com/grpc-ecosystem/go-grpc-middleware
get github.com/grpc-ecosystem/go-grpc-middleware/auth
get github.com/grpc-ecosystem/go-grpc-middleware/tags
get github.com/grpc-ecosystem/go-grpc-middleware/util/metautils
get github.com/jmoiron/sqlx
get golang.org/x/net/context
get google.golang.org/grpc
get google.golang.org/grpc/codes
get google.golang.org/grpc/grpclog
get google.golang.org/grpc/metadata
get google.golang.org/grpc/status

echo leaving $0

