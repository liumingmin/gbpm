package gbpm

import (
	_ "fmt"
	_ "errors"

	_ "github.com/astaxie/beego/orm"

	"github.com/liumingmin/gbpm/fsm"
	"github.com/liumingmin/gbpm/models"
	"github.com/liumingmin/gbpm/common"
)

//type GBpmProcessInstance struct {
//	id string
//	process *GBpmProcess
//	executions []*GBpmProcessExecution
//}

type GBpmProcessExecution struct {
	process *GBpmProcess
	currProcNode *GBpmProcessNode

	ruExecution *models.GBpmRuExecution

	parentExecution *GBpmProcessExecution
	childrenExecution map[string]*GBpmProcessExecution

	machine *fsm.Machine
}

func (this *GBpmProcessExecution) CurrentState() fsm.State {
	if this.ruExecution != nil{
		return fsm.State(this.ruExecution.CurrNodeId)
	}
	return fsm.State("")
}

func (this *GBpmProcessExecution) SetState(s fsm.State)    {
	if this.ruExecution != nil{
		this.ruExecution.CurrNodeId = string(s)
	}
}

func (this *GBpmProcessExecution) transition(nodeId string)  bool  {
	if procNode,isok := this.process.processNodes[nodeId];isok{
		err := this.machine.Transition(fsm.State(nodeId))
		if err == nil{
			this.currProcNode = procNode

			return true
		}
	}

	return false
}


func (this *GBpmProcessExecution) init() {
	if this.machine == nil {
		this.machine = &fsm.Machine{Subject: this}
	}

	this.machine.Rules = this.process.rules
}

func createExecution(process *GBpmProcess) *GBpmProcessExecution{
	ruExcetion := &models.GBpmRuExecution{}
	ruExcetion.Id = common.NewUuidV1()
	ruExcetion.Pid = ""
	ruExcetion.ProcessId = process.defProcess.Id
	ruExcetion.ProcessInstanceId = ruExcetion.Id
	ruExcetion.Name = process.defProcess.Name
	ruExcetion.State = Kgbpm_exec_active
	ruExcetion.CurrNodeId = process.startProcNode.defNode.Id

	procExec := &GBpmProcessExecution{}
	procExec.process = process
	procExec.ruExecution = ruExcetion
	procExec.currProcNode = process.startProcNode
	procExec.parentExecution = nil
	procExec.childrenExecution = make(map[string]*GBpmProcessExecution)
	procExec.init()

	return procExec
}

func createSubExecution(parentExecution *GBpmProcessExecution) *GBpmProcessExecution{
	ruExcetion := &models.GBpmRuExecution{}
	ruExcetion.Id = common.NewUuidV1()
	ruExcetion.Pid = parentExecution.ruExecution.Id
	ruExcetion.ProcessId = parentExecution.process.defProcess.Id
	ruExcetion.ProcessInstanceId = parentExecution.ruExecution.ProcessInstanceId
	ruExcetion.Name = parentExecution.process.defProcess.Name
	ruExcetion.State = Kgbpm_exec_active
	ruExcetion.CurrNodeId = parentExecution.currProcNode.defNode.Id

	procExec := &GBpmProcessExecution{}
	procExec.process = parentExecution.process
	procExec.ruExecution = ruExcetion
	procExec.currProcNode = parentExecution.currProcNode
	procExec.parentExecution = parentExecution
	procExec.childrenExecution = make(map[string]*GBpmProcessExecution)
	procExec.init()

	parentExecution.childrenExecution[ruExcetion.Id] = procExec

	return procExec
}

func createTaskInstance(execution *GBpmProcessExecution) *models.GBpmRuTask {
	taskInst := &models.GBpmRuTask{}
	taskInst.Id =  common.NewUuidV1()
	taskInst.NodeId = execution.currProcNode.defNode.Id
	taskInst.Name = execution.currProcNode.defNode.Name
	taskInst.ProcessId = execution.currProcNode.defNode.ProcessId
	taskInst.ProcInstanceId = execution.ruExecution.ProcessInstanceId
	taskInst.ExecutionId = execution.ruExecution.Id

	return taskInst
}