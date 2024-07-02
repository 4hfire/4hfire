#!/bin/bash

set -e

#代码模板
tpl="./scripts/template"

# echo "开始创建库模型：${dbname} "
# echo "start create database model >>>>>"
# goctl model mysql datasource -url="${username}:${passwd}@tcp(${host}:${port})/${dbname}" -table="*"  -dir="./model" -cache=true --style=go_zero
# goctl model mysql ddl -src="./scripts/db/authenticator.sql" -dir="./model" -cache=true --style=go_zero

#构建api服务代码
echo "start create api server >>>>>"
goctl api go -api ./api/desc/*.api -dir ./api --style=go_zero --home="./scripts/template"

#构建rpc服务代码
# echo "start create rpc server >>>>>"
# goctl rpc protoc ./rpc/pb/console.proto --go_out=./rpc --go-grpc_out=./rpc --zrpc_out=./rpc --style=go_zero

# goctl api plugin -plugin goctl-swagger="swagger -filename cabinet.json -host 127.0.0.2 -basepath /api" -api ./api/desc/cabinet.api -dir .
# docker run --rm -p 8084:8080 -e SWAGGER_JSON=/foo/cabinet.json -v $PWD:/foo swaggerapi/swagger-ui
# goctl api plugin -plugin goctl-swagger="swagger -filename ./deploy/apidoc/doc/console.json -host 127.0.0.1 -basepath /api" -api ./api/desc/console.api -dir .
# grpcui -plaintext localhost:6858