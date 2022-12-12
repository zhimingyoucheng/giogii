# 开始
##使用方法
1）主备集群位点数据比对
s 源端集群信息 si 源端连接信息  t 目标端集群信息  ti 目标端连接信息
返回值 两个集群之间如果无GTID和POS差异则返回0
```shell
./giogii -s 'admin:!QAZ2wsx' -si '172.16.76.105:16310' -t 'admin:!QAZ2wsx' -ti '172.16.128.13:16310'
```
2）参数基线比对使用方法
-s 管控平台MySQL信息，固定值，不需要改. -si 管控平台连接信息  -t 目标端集群信息  -ti 目标端连接信息 -c 管控平台基准参数模板名称
```shell
./giogii -s 'root:drACgwoqtM' -si '172.17.128.49:13336' -t 'admin:!QAZ2wsx' -ti '172.17.128.13:16310' -c base
```
3）锁监控使用方法
-s  需要监控的集群用户信息，格式为 username:password
-si 需要监控的集群连接信息，格式为 ip:port
-m  'm' 固定写法
```shell
./giogii -s 'admin:!QAZ2wsx' -si '172.17.139.26:16310' -m 'm'
```

## 编译
### x86环境
```shell
cd giogii/
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
```
### arm环境
```shell
cd giogii/
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 GOARM=7 go build
```
