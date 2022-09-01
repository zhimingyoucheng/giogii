package mapper

import "giogii/src/entity"

type ParameterOperator interface {
	DoQueryParseParameter(sqlStr string, args string) (c []entity.Configuration)
	DoClose()
}

func (sqlScaleStruct *SqlStruct) DoQueryParseParameter(sqlStr string, args string) (c []entity.Configuration) {
	rows := sqlScaleStruct.doPrepareQuery(sqlStr, args)
	defer func() {
		rows.Close()
	}()

	for rows.Next() {
		var configuration entity.Configuration
		rows.Scan(&configuration.Name, &configuration.Value, &configuration.Type)
		c = append(c, configuration)
	}
	return
}
