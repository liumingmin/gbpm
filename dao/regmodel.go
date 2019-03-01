package dao

import (
	"github.com/liumingmin/gbpm/models"
	"github.com/liumingmin/goutils/utils"
)

func regmodels() {
	utils.RegisterModels(
		new(models.SysUser),

		new(models.BpmProcess),
		new(models.BpmProcessLink),
		new(models.BpmNode),
		new(models.BpmCarbon),

		new(models.TestNode),
	)
}
