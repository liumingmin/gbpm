package gbpm

import (
	"fmt"
	"errors"

	"github.com/astaxie/beego/orm"

	"github.com/liumingmin/gbpm/models"
)



func GetGBpmModels()  []interface{}  {
	return []interface{}{
		new(models.GBpmDefProcess),
		new(models.GBpmDefTransition),
		new(models.GBpmDefNode),
		new(models.GBpmDefJob),
		new(models.GBpmRuExecution),
		new(models.GBpmRuTask),
	}
}

type GBpmEngine struct {
	 processes map[string]*GBpmProcess
	 procExecutions map[string]*GBpmProcessExecution

	 processCodeMap map[string]*GBpmProcess

	 taskInstances map[string]*models.GBpmRuTask

	 ormer orm.Ormer
	 isInited bool
}

func (this *GBpmEngine) Init(ormer orm.Ormer) {
	if this.isInited {
		return;
	}

	this.isInited = true

	this.taskInstances = make(map[string]*models.GBpmRuTask)

	this.processes = make(map[string]*GBpmProcess)
	this.procExecutions = make(map[string]*GBpmProcessExecution)
	this.processCodeMap = make(map[string]*GBpmProcess)

	this.ormer = ormer //dbase.NewOrm("")

	var procDefs = make([]*models.GBpmDefProcess,0,10)

	getGbpmData(this.ormer, "GBpmDefProcess", &procDefs, "")

	for _,defProc := range procDefs{
		var defTrans = make([]*models.GBpmDefTransition,0,10)
		var defNodes = make([]*models.GBpmDefNode,0,10)

		getGbpmData(this.ormer, "GBpmDefTransition", &defTrans, defProc.Id)
		getGbpmData(this.ormer, "GBpmDefNode", &defNodes, defProc.Id)

		process := createBpmProcess(defProc,defTrans,defNodes)

		this.processes[process.defProcess.Id] = process
		this.processCodeMap[process.defProcess.Code] = process
	}
}

func (this *GBpmEngine) LoadInstanceExecs() {
	var ruExcetions = make([]*models.GBpmRuExecution,0,100)

	getGbpmData(this.ormer, "GBpmRuExecution", &ruExcetions, "")

	for _,ruExcetion := range ruExcetions {
		if process,isok := this.processes[ruExcetion.ProcessId];isok{
			instExec := &GBpmProcessExecution{}
			instExec.process = process
			instExec.ruExcetion = ruExcetion
			instExec.init()

			this.procExecutions[ruExcetion.Id] = instExec
		}
	}
}

func getGbpmData(ormer orm.Ormer, modelName string, models interface{}, procId string){
	query := ormer.QueryTable(modelName)

	var err error
	if procId == ""{
		_, err = query.All(models)
	}else {
		_, err = query.Filter("ProcessId",procId).All(models)
	}

	if err != nil{
		fmt.Print("数据库获取流程信息失败")
		//return
	}
}

func (this *GBpmEngine) StartProcess(procCode string, params map[string]string) ( *GBpmProcessExecution,error){
	if process,isok := this.processCodeMap[procCode];isok{
		procExecution := createExecution(process)


		return procExecution,nil
	}

	return nil,errors.New("can not found process by code")
}

func (this *GBpmEngine) saveExecution(exection *GBpmProcessExecution){
	if _,isok := this.procExecutions[exection.ruExcetion.Id];isok{
		this.ormer.Update(exection.ruExcetion)
	}else{
		this.ormer.Insert(exection.ruExcetion)
		this.procExecutions[exection.ruExcetion.Id] = exection
	}
}


func (this *GBpmEngine) Transition(executionId string, taskId string) error {
	if procExecution,isok := this.procExecutions[executionId];isok{
		if procExecution.transition(taskId) {
			switch procExecution.currProcNode.defNode.Kind {
				case Kgbpm_node_normal:
					this.transToNormal(procExecution)
					break
				case Kgbpm_node_task:
					this.transToTask(procExecution)
					break
				case Kgbpm_node_end:
					this.transToEnd(procExecution)
					break
				case Kgbpm_node_fork:
					this.transToFork(procExecution)
			}

			this.ormer.Update(procExecution.ruExcetion)
		}


		return nil
	}

	return errors.New("not found process instance by id")
}

func  (this *GBpmEngine) transToNormal(execution *GBpmProcessExecution)  {

}

func  (this *GBpmEngine) transToTask(execution *GBpmProcessExecution)  {
	taskInst := createTaskInstance(execution)

	this.taskInstances[taskInst.Id] = taskInst
	this.ormer.Insert(taskInst)
}

func  (this *GBpmEngine) transToEnd(execution *GBpmProcessExecution)  {
	execution.ruExcetion.State = Kgbpm_exec_inactive
}

func  (this *GBpmEngine) transToFork(execution *GBpmProcessExecution)  {
	for _,subProcNode := range execution.currProcNode.nextNodes{
		subExecution :=  createSubExecution(execution)
		subExecution.transition(subProcNode.defNode.Id)

		this.transToTask(subExecution)

		this.saveExecution(subExecution)
	}

	execution.ruExcetion.State = Kgbpm_exec_suspend
	this.saveExecution(execution)
}

//func  (this *GBpmEngine) transToJoin(node *models.GBpmDefNode, executionId string)  {
//	taskInst := &models.GBpmRuTask{}
//	taskInst.Id =  common.NewUuidV1()
//	taskInst.Name = node.Name
//	taskInst.ProcessId = node.ProcessId
//	taskInst.ProcInstanceId = executionId
//	taskInst.NodeId = node.Id
//
//	this.ormer.Insert(taskInst)
//}


