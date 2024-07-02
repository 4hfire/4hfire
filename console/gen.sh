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
#goctl api go -api ./api/desc/*.api -dir ./api --style=go_zero --home=${tpl}

#构建rpc服务代码
# echo "start create rpc server >>>>>"
# goctl rpc protoc ./rpc/pb/console.proto --go_out=./rpc --go-grpc_out=./rpc --zrpc_out=./rpc --style=go_zero

#生成swagger文档
 goctl api plugin -plugin goctl-swagger="swagger -filename ./deploy/apidoc/doc/console.json -host 192.168.31.32:24397 -basepath /api/v1/console " -api ./api/desc/console.api -dir .
 docker run --rm -p 8083:8080 -e SWAGGER_JSON=/doc/console.json -v $PWD/deploy/apidoc/doc:/doc swaggerapi/swagger-ui