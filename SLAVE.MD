## 
Slave_IO_State: Waiting for source to send event
Master_Host: 172.16.76.106
Master_User: dbscale_internal
Master_Port: 16315
Connect_Retry: 60
Master_Log_File: greatdb-bin.000002
Read_Master_Log_Pos: 94405068
Relay_Log_File: mysql-relay.000006
Relay_Log_Pos: 94381185
Relay_Master_Log_File: greatdb-bin.000002
Slave_IO_Running: Yes
Slave_SQL_Running: Yes
Replicate_Do_DB: 
Replicate_Ignore_DB: 
Replicate_Do_Table: 
Replicate_Ignore_Table: 
Replicate_Wild_Do_Table: 
Replicate_Wild_Ignore_Table:
Last_Errno: 0
Last_Error:
Skip_Counter: 0
Exec_Master_Log_Pos: 94405068
Relay_Log_Space: 94381829
Until_Condition: None
Until_Log_File:
Until_Log_Pos: 0
Master_SSL_Allowed: No
Master_SSL_CA_File:
Master_SSL_CA_Path:
Master_SSL_Cert:
Master_SSL_Cipher:
Master_SSL_Key:
Seconds_Behind_Master: 0
Master_SSL_Verify_Server_Cert: No
Last_IO_Errno: 0
Last_IO_Error:
Last_SQL_Errno: 0
Last_SQL_Error:
Replicate_Ignore_Server_Ids:
Master_Server_Id: 11
Master_UUID: ca44b126-1843-11ed-ae88-5254008903de
Master_Info_File: mysql.slave_master_info
SQL_Delay: 0
SQL_Remaining_Delay: NULL
Slave_SQL_Running_State: Replica has read all relay log; waiting for more updates
Master_Retry_Count: 86400
Master_Bind:
Last_IO_Error_Timestamp:
Last_SQL_Error_Timestamp:
Master_SSL_Crl:
Master_SSL_Crlpath:
Retrieved_Gtid_Set: ca44b126-1843-11ed-ae88-5254008903de:1-32:34-133
Executed_Gtid_Set: ca44b126-1843-11ed-ae88-5254008903de:1-133,
e37b67a2-183a-11ed-8e9f-52540012644d:1-3,
e3aa3eeb-183a-11ed-b503-5254008da967:1-43
Auto_Position: 1
Replicate_Rewrite_DB:
Channel_Name:
Master_TLS_Version:
Master_public_key_path:
Get_master_public_key: 0
Network_Namespace:


# 显示二进制文件以及文件大小（需要开启二进制日志记录功能）
SHOW {BINARY | MASTER} LOGS

# 显示二进制文件的执行过程
SHOW BINLOG EVENTS [IN 'log_name'] [FROM pos] [LIMIT [offset,] row_count]

# 显示MySQL当前支持哪些字符集
SHOW CHARACTER SET [like_or_where]

# 显示MySQL支持字符集的排序规则
SHOW COLLATION [like_or_where]

# 显示表的列信息（等同于DESC，需要先创建表）
SHOW [FULL] COLUMNS FROM tbl_name [FROM db_name] [like_or_where]

# 显示已经创建的库，创建时的语句
SHOW CREATE DATABASE db_name

# 显示已经创建的事件，创建时的语句
SHOW CREATE EVENT event_name

# 显示已经创建的函数，创建时的语句
SHOW CREATE FUNCTION func_name

# 显示已经创建的存储过程，创建时的语句
SHOW CREATE PROCEDURE proc_name

# 显示已经创建的表，创建时的语句
SHOW CREATE TABLE tbl_name

# 显示已经创建的触发器，创建时的语句
SHOW CREATE TRIGGER trigger_name

# 显示已经创建的视图，创建时的语句
SHOW CREATE VIEW view_name

# 显示mysql中所有数据库的名称
SHOW DATABASES [like_or_where]

# 显示存储引擎的详细信息
SHOW ENGINE engine_name {STATUS | MUTEX}

# 显示数据库支持的存储引擎和默认存储引擎
SHOW [STORAGE] ENGINES

# 显示最后一个执行语句所产生的错误信息
SHOW ERRORS [LIMIT [offset,] row_count]

# 显示事件信息
SHOW EVENTS

# 服务器内部调试，显示一个指定存储的内部实现的表示形式过程
SHOW FUNCTION CODE func_name

# 显示存储函数信息（需要先创建存储函数）
SHOW FUNCTION STATUS [like_or_where]

# 显示指定用户拥有的权限
SHOW GRANTS FOR user

# 显示表索引信息（需要先创建索引）
SHOW INDEX FROM tbl_name [FROM db_name]

# 显示Master当前正在使用的二进制信息
SHOW MASTER STATUS

# 列举在表缓存中当前被打开的非TEMPORARY表
SHOW OPEN TABLES [FROM db_name] [like_or_where]

# 显示MySQL插件信息
SHOW PLUGINS

# 服务器内部调试，显示一个指定存储的内部实现的表示形式过程
SHOW PROCEDURE CODE proc_name

# 显示存储过程信息（需要先创建存储过程）
SHOW PROCEDURE STATUS [like_or_where]

# 显示MySQL所支持的所有权限，及权限可操作的对象
SHOW PRIVILEGES

# 显示系统中正在运行的所有进程，普通用户只能查看自己的进行信息
SHOW [FULL] PROCESSLIST

# 显示当前会话执行语句资源使用情况
SHOW PROFILE [types] [FOR QUERY n] [OFFSET n] [LIMIT n]

# 显示当前会话执行语句资源使用情况
SHOW PROFILES

# 显示relaylog事件信息（需要先做主从复制）
SHOW RELAYLOG EVENTS [IN 'log_name'] [FROM pos] [LIMIT [offset,] row_count]

# 显示Master主机上已注册的复制主机列表（需要先做主从复制）
SHOW SLAVE HOSTS

# 显示Slave主机状态信息（需要先做主从复制）
SHOW SLAVE STATUS [FOR CHANNEL channel]

# 显示MySQL状态信息
SHOW [GLOBAL | SESSION] STATUS [like_or_where]

# 显示表属性信息
SHOW TABLE STATUS [FROM db_name] [like_or_where]

# 显示当前数据库中所有表的名称
SHOW [FULL] TABLES [FROM db_name] [like_or_where]

# 显示触发器信息（需要先创建触发器）
SHOW TRIGGERS [FROM db_name] [like_or_where]

# 显示MySQL变量信息
SHOW [GLOBAL | SESSION] VARIABLES [like_or_where]

# 显示最后一个执行语句所产生的警告信息
SHOW WARNINGS [LIMIT [offset,] row_count]




