package gbpm

import (
	_ "fmt"
	_ "errors"

	_ "github.com/astaxie/beego/orm"

	"github.com/liumingmin/gbpm/fsm"
	"github.com/liumingmin/gbpm/models"
	_ "github.com/liumingmin/gbpm/common"
)

const (
	Kgbpm_node_start     	= iota
	Kgbpm_node_normal
	Kgbpm_node_task
	Kgbpm_node_end
	Kgbpm_node_fork
	Kgbpm_node_join
)

const (
	Kgbpm_exec_inactive     	= iota
	Kgbpm_exec_active
	Kgbpm_exec_suspend
)

type GBpmProcess struct {
	defProcess *models.GBpmDefProcess
	defTrans []*models.GBpmDefTransition
	defNodes []*models.GBpmDefNode

	processNodes map[string]*GBpmProcessNode
	startProcNode *GBpmProcessNode

	rules  *fsm.Ruleset
}

type GBpmProcessNode struct {
	defNode *models.GBpmDefNode
	nextNodes map[string]*GBpmProcessNode
}

func (this *GBpmProcess) init()  {
	for _,defNode := range this.defNodes{
		procNode := &GBpmProcessNode{}
		procNode.defNode = defNode
		procNode.nextNodes = make(map[string]*GBpmProcessNode)

		this.processNodes[defNode.Id] = procNode

		if defNode.Kind == Kgbpm_node_start {
			this.startProcNode = procNode
		}
	}

	this.rules = &fsm.Ruleset{}

	for _,tran := range this.defTrans{
		this.rules.AddTransition(fsm.T{fsm.State(tran.PreNodeId), fsm.State(tran.NextNodeId)})
		if tran.AllowBack == 1 {
			this.rules.AddTransition(fsm.T{fsm.State(tran.NextNodeId),fsm.State(tran.PreNodeId)})
		}

		if preNode,isok := this.processNodes[tran.PreNodeId];isok{
			if nextNode,isok2 := this.processNodes[tran.NextNodeId];isok2{
				preNode.nextNodes[nextNode.defNode.Id] = nextNode
			}
		}

		if tran.AllowBack == 1 {
			if nextNode,isok := this.processNodes[tran.NextNodeId];isok{
				if preNode,isok2 := this.processNodes[tran.PreNodeId];isok2{
					nextNode.nextNodes[preNode.defNode.Id] = preNode
				}
			}
		}
	}
}

func createBpmProcess(defProcess *models.GBpmDefProcess, defTrans []*models.GBpmDefTransition,
	defNodes []*models.GBpmDefNode) *GBpmProcess {

	process := &GBpmProcess{}
	process.defProcess =  defProcess
	process.defTrans = defTrans
	process.defNodes = defNodes
	process.init()

	return process
}