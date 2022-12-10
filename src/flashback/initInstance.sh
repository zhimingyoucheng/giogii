#!/bin/bash
kill -9 `ps -ef | grep mysqld | grep clonebackup| grep -v $$ | awk '{print $2}' | xargs`
rm -rf /data/mysqldata/clonebackup/
tmp=`ls /data/mysqldata/`
cnfdir=`find /data/mysqldata/ -name *.conf`
socket=`find /data/mysqldata/ -name mysql.sock`
mkdir -p /data/mysqldata/clonebackup/{logfile,dbdata,tmp,socket,pid}
cp ${cnfdir} /data/mysqldata/clonebackup/
cnfdir=`find /data/mysqldata/clonebackup/ -name *.conf`
sed "s/${tmp}/clonebackup/g" ${cnfdir} > /data/mysqldata/clonebackup/new.conf
echo "port = 18000" >> /data/mysqldata/clonebackup/new.conf
echo "innodb_buffer_pool_size = 512MB" >> /data/mysqldata/clonebackup/new.conf
/data/app/mysql-8.0.26/bin/mysqld --defaults-file=/data/mysqldata/clonebackup/new.conf --initialize-insecure --user=mysql
nohup /data/app/mysql-8.0.26/bin/mysqld_safe --defaults-file=/data/mysqldata/clonebackup/new.conf --user=mysql --datadir=/data/mysqldata/clonebackup/dbdata > /data/mysqldata/clonebackup/logfile/out.log 2>&1 &
ls