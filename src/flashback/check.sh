#!/bin/bash
username=$1
password=$2
socket=`find /data/mysqldata/clonebackup/ -name mysql.sock`
/data/app/mysql-8.0.26/bin/mysql -u${username} -p${password} -S ${socket} -e "shutdown;"
sleep 2s
mysqld_safe=`ps -ef | grep mysqld_safe |  grep -v $$ |  awk '{print $8" "$9" "$10" "$11" "$12" "}'`
a=`echo ${mysqld_safe} | awk '{print $1" "$2" "$3" "$4" "$5" "}'`
kill -9 `ps -ef | grep mysqld | grep 16| grep -v $$ | awk '{print $2 " " }' | xargs`
string=`ls /data/mysqldata/`
array=(${string// /})
rm -rf /data/mysqldata/${array}/dbdata_bak
mv /data/mysqldata/${array}/dbdata /data/mysqldata/${array}/dbdata_bak
sleep 2s
mv /data/mysqldata/clonebackup/dbdata /data/mysqldata/${array}/
sleep 2s
nohup ${a} > /data/mysqldata/${array}/logfile/out.log 2>&1 &
rm -rf /data/mysqldata/clonebackup
