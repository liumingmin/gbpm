package dao

import (
	"github.com/astaxie/beego/orm"
	"github.com/liumingmin/goutils/utils"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	regmodels()

	var models = utils.GetRegModels()
	orm.RegisterModelWithPrefix("", models...)

	initDb()
}
