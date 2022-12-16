# 开始

## 使用方法

1）主备集群位点数据比对, -s 源端集群信息, -si 源端连接信息, -t 目标端集群信息, -ti 目标端连接信息,两个集群之间如果无GTID和POS差异则返回0

```shell
./giogii -s 'admin:!QAZ2wsx' -si '172.16.76.105:16310' -t 'admin:!QAZ2wsx' -ti '172.16.128.13:16310'
```

2）参数基线比对使用方法,-s 管控平台MySQL信息,固定值不需要改, -si 管控平台连接信息, -t 目标端集群信息, -ti 目标端连接信息, -c 管控平台基准参数模板名称

```shell
./giogii -s 'root:drACgwoqtM' -si '172.17.128.49:13336' -t 'admin:!QAZ2wsx' -ti '172.17.128.13:16310' -c base
```

3）锁监控使用方法,-s 需要监控的集群用户信息,格式为 username:password,-si 需要监控的集群连接信息,格式为 ip:port -m m 固定写法

```shell
./giogii -s 'admin:!QAZ2wsx' -si '172.17.139.26:16310' -m m
```

4）灾备集群flashback使用方法,该方法使用的是clone slave节点, -u ssh用户名称, -p ssh用户密码, -f 闪回动作启停, start执行闪回动作准备阶段, stop执行闪回动作后续流程,
执行start和stop直接的时间业务是可以对灾备集群进行写入操作. -s 主集群信息, -si 主集群连接信息, -t 灾备集群信息, -ti 灾备集群连接信息

```shell
./giogii -t 'admin:!QAZ2wsx' -ti '172.17.139.26:16310' -s 'admin:!QAZ2wsx' -si '172.17.139.26:16320'  -u mysql -p mysql -f start
./giogii -t 'admin:!QAZ2wsx' -ti '172.17.139.26:16310' -s 'admin:!QAZ2wsx' -si '172.17.139.26:16320'  -u mysql -p mysql -f stop
```

5）灾备集群flashback使用方法,该方法使用的是proxy flashback工具执行, -u ssh用户名称, -p ssh用户密码, -f 闪回动作启停, begin执行闪回动作准备阶段, end执行闪回动作后续流程,
执行begin和end直接的时间业务是可以对灾备集群进行写入操作. -s 主集群信息, -si 主集群连接信息, -t 灾备集群信息, -ti 灾备集群连接信息

```shell
./giogii -t 'admin:!QAZ2wsx' -ti '172.17.139.26:16310' -s 'admin:!QAZ2wsx' -si '172.17.139.26:16320'  -u mysql -p mysql -f begin
./giogii -t 'admin:!QAZ2wsx' -ti '172.17.139.26:16310' -s 'admin:!QAZ2wsx' -si '172.17.139.26:16320'  -u mysql -p mysql -f end
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
