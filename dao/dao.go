package dao

import (
	"fmt"
	"os"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/liumingmin/gbpm/common"
	"github.com/liumingmin/goutils/conf"
	"github.com/liumingmin/goutils/log4go"
)

func NewOrm() orm.Ormer {
	var result = orm.NewOrm()
	result.Using(common.GMasterDB)
	return result
}

func NewReadOrm() orm.Ormer {
	var result = orm.NewOrm()
	result.Using(common.GSlaveDB)
	return result
}

func initDb() {
	dbs := conf.Conf.DATABASES
	if dbs == nil {
		fmt.Fprintf(os.Stderr, "No database configuration")
		return
	}

	orm.RegisterDriver("mysql", orm.DRMySQL)

	for _, _db := range dbs {
		err := initMySQL(&_db)
		if err != nil {
			log4go.Error("Init MySQL failed. key: %v, err: %v", _db.KEY, err)
		}
	}

	orm.Debug = (conf.Conf.SERVER.RUNMODE == common.GRunDebug)
	orm.DefaultRowsLimit = -1
	orm.DefaultTimeLoc, _ = time.LoadLocation("Asia/Shanghai")

}

func initMySQL(dbconf *conf.Database) (err error) {

	connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbconf.USER, dbconf.PASSWORD, dbconf.HOST, dbconf.NAME)
	dbparmas := "?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"

	err = orm.RegisterDataBase(dbconf.KEY, "mysql", connStr+dbparmas)

	if err != nil {
		log4go.Error("ConnStr=%v, err=%v", connStr, err)
		return err
	} else {
		log4go.Info("ConnStr=%v", connStr)
	}

	maxIdle := dbconf.Ext("maxIdle", 0).(int)
	maxOpen := dbconf.Ext("maxOpen", 16).(int)

	orm.SetMaxIdleConns(dbconf.KEY, maxIdle)
	orm.SetMaxOpenConns(dbconf.KEY, maxOpen)

	syncdb := dbconf.Ext("syncDb", false).(bool)
	if syncdb {
		err = orm.RunSyncdb(dbconf.KEY, false, true)
		if err != nil {
			log4go.Error("InitDb@orm.RunSyncdb: err=%v", err)
		}
	}

	return
}
