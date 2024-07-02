# rpc name
Name: {{.serviceName}}.rpc

# if dev mode, ListenOn use config value, if not use auto generate
ListenOn: 127.0.0.1:8080

# etcd config
Etcd:
  Hosts:
  - 127.0.0.1:2379
  Key: {{.serviceName}}.rpc
