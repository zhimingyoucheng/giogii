package check

import (
	"fmt"
	"giogii/src/mapper"
	"log"
)

var BaseParameter mapper.SqlScaleOperator
var ClusterParameter mapper.SqlScaleOperator
var TargetSocket string

func InitCheckParameterConf(sourceUserInfo string, sourceSocket string, sourceDatabase string, targetUserInfo string, targetSocket string, targetDatabase string) {
	s, t := mapper.InitConfig(sourceUserInfo, sourceSocket, sourceDatabase, targetUserInfo, targetSocket, targetDatabase)
	BaseParameter = &s
	ClusterParameter = &t
	TargetSocket = targetSocket
}

func DoCheckParameter() {
	// select name,value,type from configuration_items where configuration_id = "d992bc11-fe27-4e03-a355-4ed325c7ca23";

	// init base template
	// select i.name,i.value,i.type from configuration_items as i inner join configuration as c on c.uuid = i.configuration_id where c.name = "base";
	var strSql = "select i.name,i.value,i.type from configuration_items as i inner join configuration as c on c.uuid = i.configuration_id where c.name = ?"
	configuration := BaseParameter.DoQueryParseParameter(strSql, "base")
	for i := 0; i < len(configuration); i++ {
		if configuration[i].Type == "dbscale" {
			strSql = fmt.Sprintf("dbscale show options like '%s'", configuration[i].Name)
			value := ClusterParameter.DoQueryParseValue(strSql)
			if value == "TRUE" {
				value = "1"
			} else if value == "FALSE" {
				value = "0"
			}
			if value != configuration[i].Value {
				log.Println(fmt.Sprintf("实例[%s]参数：%s 基准值为：%s,实际值为：%s", TargetSocket, configuration[i].Name, configuration[i].Value, value))
			}
		}
	}
}
