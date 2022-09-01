package check

import (
	"fmt"
	"giogii/src/mapper"
	"log"
	"strings"
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
			value := strings.ToLower(ClusterParameter.DoQueryParseValue(strSql))
			if value == "true" {
				value = "1"
			} else if value == "false" {
				value = "0"
			}
			if value != strings.ToLower(configuration[i].Value) {
				log.Println(fmt.Sprintf("[实例 %s]参数：%s 基准值为：%s,实际值为：%s", TargetSocket, configuration[i].Name, configuration[i].Value, value))
			}
		} else if configuration[i].Type == "mysql" {

			if configuration[i].Name == "performance-schema-instrument" {
				continue
			}

			if strings.Contains(configuration[i].Name, "performance-schema-consumer") {
				strSql = fmt.Sprintf("select * from performance_schema.setup_consumers where name = ?")
				index := strings.Index(configuration[i].Name, "consumer")
				args := strings.ReplaceAll(configuration[i].Name[index+9:], "-", "_")
				consumer := ClusterParameter.DoQueryParseConsumers(strSql, args)
				if consumer.Enabled == "YES" {
					consumer.Enabled = "on"
				} else if consumer.Enabled == "NO" {
					consumer.Enabled = "off"
				}

				if consumer.Enabled != configuration[i].Value {
					log.Println(fmt.Sprintf("[实例 %s]参数：%s 基准值为：%s,实际值为：%s", TargetSocket, configuration[i].Name, configuration[i].Value, consumer.Enabled))
				}
			} else {
				strSql = fmt.Sprintf("show variables like '%s'", configuration[i].Name)
				value := strings.ToLower(ClusterParameter.DoQueryParseValue(strSql))
				baseValue := strings.ToLower(configuration[i].Value)
				if value == "on" {
					value = "1"
				} else if value == "off" {
					value = "0"
				}
				if baseValue == "on" {
					baseValue = "1"
				} else if baseValue == "off" {
					baseValue = "0"
				}
				if value != baseValue {
					log.Println(fmt.Sprintf("[实例 %s]参数：%s 基准值为：%s,实际值为：%s", TargetSocket, configuration[i].Name, configuration[i].Value, value))
				}
			}

		}
	}
}
