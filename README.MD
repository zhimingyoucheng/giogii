### 主备集群数据比对、基线参数比对

```shell
master_info_repository = TABLE
sync_master_info = 10000  io thread 的位置持久化也就没法保证正确
relay_log_recovery = ON 


slave_master_info 表存储 slave io thread  读取主库的 binlog 位置
slave_relay_log_info 表存储 slave sql thread 的工作位置

# 场景：当参数relay_log_recovery设置为on，如果slave节点与主节点直接发生断开，且此时slave宕机，那么slave io thread 之前保存的relay log将会失效
# 会按照sql thread 回放的点位从主库再次获取，此时已经连不上主了，所以slave延迟的数据会全部丢弃。

# 参数比对工具
innodb_thread_concurrency 可以不一样
long_query_time 
innodb_page_cleaners = innodb_buffer_pool_instances
open_files_limit = ulimit -n  8.0.19以后
```


### 使用方法
```shell
# 1）主备集群数据比对使用方法
# s 源端集群信息 si 源端连接信息  t 目标端集群信息  ti 目标端连接信息
./giogii -s 'admin:!QAZ2wsx' -si '172.16.76.105:16310' -t 'admin:!QAZ2wsx' -ti '172.16.128.13:16310'

# 返回值 两个集群之间如果无GTID和POS差异则返回0，任何大于0的值都需要发起二次调用


# 2）参数基线比对使用方法
# -s 管控平台MySQL信息，固定值，不需要改. -si 管控平台连接信息  -t 目标端集群信息  -ti 目标端连接信息 -c 管控平台基准参数模板
./giogii -s 'root:drACgwoqtM' -si '172.17.128.49:13336' -t 'admin:!QAZ2wsx' -ti '172.17.128.13:16310' -c base

# 检查结果会打印到控制台

```

### 编译
```shell
# x86环境
cd giogii/
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

# arm环境
cd giogii/
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 GOARM=7 go build

```

