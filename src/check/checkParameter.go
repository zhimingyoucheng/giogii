package check

import "giogii/src/mapper"

var BaseParameter mapper.ParameterOperator
var ClusterParameter mapper.ParameterOperator

func InitCheckParameterConf(sourceUserInfo string, sourceSocket string, sourceDatabase string, targetUserInfo string, targetSocket string, targetDatabase string) {
	s, t := mapper.InitConfig(sourceUserInfo, sourceSocket, sourceDatabase, targetUserInfo, targetSocket, targetDatabase)
	BaseParameter = &s
	ClusterParameter = &t
}

func DoCheckParameter() {
	// select name,value,type from configuration_items where configuration_id = "d992bc11-fe27-4e03-a355-4ed325c7ca23";

	// init base template
	// select i.name,i.value,i.type from configuration_items as i inner join configuration as c on c.uuid = i.configuration_id where c.name = "base";
	var strSql = "select i.name,i.value,i.type from configuration_items as i inner join configuration as c on c.uuid = i.configuration_id where c.name = ?"
	configuration := BaseParameter.DoQueryParseParameter(strSql, "base")
	for i := 0; i < len(configuration); i++ {
	}
}
