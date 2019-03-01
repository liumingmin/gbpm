package models

import "time"

type SysUser struct {
	Id         string    `orm:"size(32);pk;"`
	OrderNo    int       `orm:"null;default(0);"`
	UserName   string    `orm:"null;size(64);"`
	RealName   string    `orm:"null;size(64);"`
	Password   string    `orm:"null;size(255);"`
	IsEnable   int       `orm:"null;default(1);"`
	CreateDate time.Time `orm:"null;type(date);auto_now_add;"`
	UpdateDate time.Time `orm:"null;type(date);auto_now;"`
	PictureId  string    `orm:"null;size(32);"`
	Remark     string    `orm:"null;size(255);"`
	Ver        int       `orm:"null;default(1);"`
}
