package flashback

import (
	"log"
	"time"
)

func DoFlashbackByDbScaleTools(sourceUserInfo string, sourceSocket string, targetUserInfo string, targetSocket string) {
	InitMasterConnection(sourceUserInfo, sourceSocket)
	InitSlaveConnection(targetUserInfo, targetSocket)

	// 1.1 断开主备集群的复制，主集群踢出、备集群断开
	RemoveSlaveCluster()
	CloseReplication()

	// 1.2 等待binlog回放完成
	/**
	获取灾备集群的GTID，确保灾备集群的数据全部回放完成
	*/
	for {
		GetSlaveGTIDSet()
		if SlaveStatus.SecondsBehindMaster.Int64 == 0 {
			log.Println("记录gtid，确保灾备集群全部回放完Binlog: ", SlaveStatus.ExecutedGtidSet)
			break
		}
		time.Sleep(3 * time.Second)
	}

	// 1.3 记录备集群GTID和POS位点信息，记录备集群拓扑关系、IP信息

	// 1.4 关闭备集群只读参数，变为read write

	// 1.5 监控线程监控备集群在回放期间的拓扑关系，每5秒一次，如果拓扑发生变化则记录一次GTID和POS位点

	// 2.1 打开备集群只读参数，变为read only

	// 2.2 记录备集群主节点GTID和POS位点信息，

	// 2.3 获取监控线程记录的拓扑关系，判断是否有主从切换，

	// 2.4 根据binlog位点信息、GTID信息调用dbscale_binlog_tool执行闪回动作

	// 2.5 set global gtid_purged="旧GTID"

	// 2.6 重新构建主集群和备集群的复制关系
}
