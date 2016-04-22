package main

import (
	"fmt"
	_ "github.com/astaxie/beego"
	"github.com/liumingmin/gbpm/gbpm"
	_ "github.com/liumingmin/gbpm/models"

	"tsrv/dbase"
)

func main() {
	Test()
}

//----------------------------------------------------------------------------------------------------------------------


func Test() {

	engine := &gbpm.GBpmEngine{}
	engine.Init(dbase.NewOrm(""))
	engine.LoadInstanceExecs()

		_,err := engine.StartProcess("simple",nil)

		if err != nil{
			fmt.Print(err)
			return
		}
		//engine.Transition(procInst.Id, "t2")

	//engine.Transition("5718969654126b25b8000001","t4")
}