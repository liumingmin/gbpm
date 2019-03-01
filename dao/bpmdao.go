package dao

import (
	"github.com/astaxie/beego/orm"
	"github.com/liumingmin/goutils/utils"
)

func FindProcessByUserId(ormer orm.Ormer, code, userId string) string {
	var lists []orm.ParamsList

	ormer.Raw(`select t1.Id from BpmProcess t1 ,BpmProcessLink t2 where t1.Id=t2.ProcessId 
					and t1.Code = ? and t2.UserId = ?`, code, userId).ValuesList(&lists)

	if len(lists) > 0 {
		return lists[0][0].(string)
	}

	return ""
}

func FindNodeOrderNoByUserId(ormer orm.Ormer, code, instanceId string) interface{} {
	ms := utils.CreateModels(code + "_node")
	ormer.QueryTable(code+"_node").Filter("instanceId", instanceId).All(ms)
	return ms
}
