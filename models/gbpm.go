package models

import (
	"time"
)

type BpmProcess struct {
	Id         string    `orm:"size(32);pk;"`
	Code       string    `orm:"null;size(32);index;"`
	Name       string    `orm:"null;size(160);"`
	Remark     string    `orm:"null;size(255);"`
	CreateDate time.Time `orm:"null;type(datetime);auto_now_add;"`
}

type BpmProcessLink struct {
	Id        string `orm:"size(32);pk;"`
	ProcessId string `orm:"null;size(32);index;"`
	UserId    string `orm:"null;size(32);index;"`
}

type BpmNode struct {
	Id        string `orm:"size(32);pk;"`
	OrderNo   int    `orm:"null;default(0);"`
	ProcessId string `orm:"null;size(32);index;"`
	UserId    string `orm:"null;size(32);index;"`
	Kind      int    `orm:"null;default(0);"` //0 node 1 end
	Remark    string `orm:"null;size(255);"`
}

type BpmCarbon struct {
	Id        string `orm:"size(32);pk;"`
	OrderNo   int    `orm:"null;default(0);"`
	ProcessId string `orm:"null;size(32);index;"`
	UserId    string `orm:"null;size(32);index;"`
}

////////////////////////////////////////////////////////////////////////////////////////////////

type InstanceNode struct {
	Id         string `orm:"size(64);pk;"`
	OrderNo    int    `orm:"null;default(0);"`
	InstanceId string `orm:"null;size(32);index;"`
	UserId     string `orm:"null;size(32);index;"`
	Kind       int    `orm:"null;default(0);"` //0 node 1 end
	Token      int    `orm:"null;default(0);"` //0 no 1 get
	Done       int    `orm:"null;default(0);"` //0 no 1 done
	Msg        string `orm:"null;size(255);"`
	Remark     string `orm:"null;size(255);"`
}

type InstanceCarbon struct {
	Id         string `orm:"size(32);pk;"`
	OrderNo    int    `orm:"null;default(0);"`
	InstanceId string `orm:"null;size(32);index;"`
	UserId     string `orm:"null;size(32);index;"`
}
