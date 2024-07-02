FROM golang:alpine as builder

ENV GO111MODULE=on

ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /app

COPY . .

RUN go mod tidy
RUN cp -r /app/web/dist /app/resource/
RUN go env;CGO_ENABLED=0 go build -ldflags '-s -w' -o hfc main.go

FROM alpine:latest

ENV TZ Asia/Shanghai
ENV SERVER_PORT 8081
ENV SERVER_ADDR 0.0.0.0:$SERVER_PORT
ENV TIME_ZONE=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TIME_ZONE /etc/localtime && echo $TIME_ZONE > /etc/timezone

WORKDIR /usr/local/hfc
RUN touch console.yaml

COPY --from=builder /app/4hfire ./

EXPOSE $SERVER_PORT

RUN apk add tzdata
RUN cp /usr/share/zoneinfo/${TZ} /etc/localtime
RUN echo ${TZ} > /etc/timezone
ENTRYPOINT ./hfc