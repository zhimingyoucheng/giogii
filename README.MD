

# 主机群和灾备集群数据比对工具，主要用于主集群和灾备集群之间的切换功能，对比集群之间的增量差异

# 架构
## 无状态程序，提供调用接口 
## 客户端与服务端采用RPC调用
## server 端 使用grpc框架

master_info_repository = TABLE
sync_master_info = 10000  io thread 的位置持久化，也就没法保证正确，
relay_log_recovery = ON 


slave_master_info 表存储 slave io thread  读取主库的 binlog 位置
slave_relay_log_info 表存储 slave sql thread 的工作位置

场景：当参数relay_log_recovery设置为on，如果slave节点与主节点直接发生断开，且此时slave宕机，那么slave io thread 之前保存的relay log将会失效
会按照sql thread 回放的点位从主库再次获取，此时已经连不上主了，所以slave延迟的数据会全部丢弃。



# 实现
## 1）对比主集群和灾备集群的GTID
### 1.1）begin  全部去读取主集群和灾备集群的主节点
### 1.2) 一个时间范围内的比对，比如1秒内的比对，连接主集群和灾备集群要同时在200ms之内同时返回
### 1.3）连接超时