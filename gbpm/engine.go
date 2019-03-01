package gbpm

import (
	"errors"

	"fmt"

	"github.com/liumingmin/gbpm/common"
	"github.com/liumingmin/gbpm/dao"
)

func StartInstance(processCode, userId, instanceId string, issubmit bool) error {
	var ormer = dao.NewOrm()
	processId := dao.FindProcessByUserId(ormer, processCode, userId)
	if processId == "" {
		return errors.New(common.ErrMsgCreateProcess)
	}

	createNodes := fmt.Sprintf(`insert into %s_node(Id,InstanceId,OrderNo,UserId,Kind) 
		select UUID(),'%s',OrderNo,UserId,Kind from bpmnode 
		where ProcessId = '%s'`, processCode, instanceId, processId)
	ormer.Raw(createNodes).Exec()

	if issubmit {
		signalStart := fmt.Sprintf(`update %s_node set Token=1 where InstanceId='%s' and OrderNo =
			(select * from (select min(OrderNo) from %s_node where InstanceId='%s') tmp)`,
			processCode, instanceId, processCode, instanceId)
		ormer.Raw(signalStart).Exec()
	}

	return nil
}

func SignalInstance(processCode, userId, instanceId, msg string) (bool, error) {
	var ormer = dao.NewOrm()
	nodes := dao.FindNodeOrderNoByUserId(ormer, processCode, instanceId)

	if len(nodes) == 0 {
		return false, errors.New(common.ErrMsgSignalEnd)
	}

	orderNo := -1
	for _, node := range nodes {
		if node.UserId == userId && node.Token == 1 {
			orderNo = node.OrderNo
			break
		}
	}

	if orderNo < 0 {
		orderNo = 1 //TODO min
	}

	signal1 := fmt.Sprintf(`update %s_node set Token=0,Done=1,Msg='%s' where InstanceId='%s' and OrderNo = %d and Token=1`,
		processCode, msg, instanceId, orderNo)

	signal2 := fmt.Sprintf(`update %s_node set Token=1 where InstanceId='%s' and OrderNo = %d `,
		processCode, instanceId, orderNo+1)

	ormer.Raw(signal1).Exec()
	res, _ := ormer.Raw(signal2).Exec()
	num, _ := res.RowsAffected()
	return num == 0, nil
}

//func Min(){
//
//}
