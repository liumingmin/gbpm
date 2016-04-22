package models

import (
	"time"
)

type GBpmDefProcess struct {
    Id  string  `orm:"size(32);pk;"`
    OrderNo  int  `orm:"null;default(0);"`
    Code  string  `orm:"null;size(32);index;"`
    Name  string  `orm:"null;size(255);"`
    CreateDate  time.Time  `orm:"null;type(date);auto_now_add;"`
    UpdateDate  time.Time  `orm:"null;type(date);auto_now;"`
    CreateUser  string  `orm:"null;size(36);"`
    UpdateUser  string  `orm:"null;size(36);"`
    Remark  string  `orm:"null;size(255);"`
    Ver  int  `orm:"null;default(1);"`
}

type GBpmDefTransition struct {
    Id  string  `orm:"size(32);pk;"`
    ProcessId  string  `orm:"size(32);"`
    Name  string  `orm:"null;size(64);"`
    PreNodeId  string  `orm:"null;size(32);"`
    NextNodeId  string  `orm:"null;size(32);"`
    AllowBack int  `orm:"null;default(0);"`
    CreateDate  time.Time  `orm:"null;type(date);auto_now_add;"`
    UpdateDate  time.Time  `orm:"null;type(date);auto_now;"`
    Ver  int  `orm:"null;default(1);"`
}

type GBpmDefNode struct {
    Id  string  `orm:"size(32);pk;"`
    ProcessId  string  `orm:"size(32);"`
    Name  string  `orm:"null;size(255);"`
    Kind  int  `orm:"null;default(1);"`
    ApprovalRoleId  string  `orm:"null;size(32);"`
    ApprovalUserId  string  `orm:"null;size(32);"`
    JobId  string  `orm:"null;size(255);"`
    CreateDate  time.Time  `orm:"null;type(date);auto_now_add;"`
    UpdateDate  time.Time  `orm:"null;type(date);auto_now;"`
    Remark  string  `orm:"null;size(255);"`
    Ver  int  `orm:"null;default(1);"`
}

type GBpmDefJob struct {
    Id  string  `orm:"size(32);pk;"`
    OrderNo  int  `orm:"null;default(0);"`
    ProcessId  string  `orm:"null;size(32);"`
    Name  string  `orm:"null;size(255);"`
    TriggerFunc  string  `orm:"null;size(255);"`
    CreateDate  time.Time  `orm:"null;type(date);auto_now_add;"`
    UpdateDate  time.Time  `orm:"null;type(date);auto_now;"`
    Remark  string  `orm:"null;size(255);"`
    Ver  int  `orm:"null;default(1);"`
}

type GBpmRuExecution struct {
    Id  string  `orm:"size(32);pk;"`
    Pid  string  `orm:"null;size(32);"`
    Name  string  `orm:"null;size(255);"`
    ProcessId  string  `orm:"null;size(32);"`
    ProcessInstanceId  string  `orm:"null;size(32);"`
    CurrNodeId  string  `orm:"null;size(32);"`
    State int `orm:"null;default(1);"`
    BizModelName  string  `orm:"null;size(255);"`
    BizEntityId  string  `orm:"null;size(255);"`
    CreateDate  time.Time  `orm:"null;type(date);auto_now_add;"`
    UpdateDate  time.Time  `orm:"null;type(date);auto_now;"`
    CreateUser  string  `orm:"null;size(36);"`
    UpdateUser  string  `orm:"null;size(36);"`
    Remark  string  `orm:"null;size(255);"`
    Ver  int  `orm:"null;default(1);"`
}

type GBpmRuTask struct {
    Id  string  `orm:"size(32);pk;"`
    Name  string  `orm:"null;size(255);"`
    ProcessId  string  `orm:"null;size(32);"`
    ProcInstanceId  string  `orm:"null;size(32);"`
    ExecutionId  string  `orm:"null;size(32);"`
    NodeId  string  `orm:"null;size(32);"`
    ApprovalRoleId  string  `orm:"null;size(32);"`
    ApprovalUserId  string  `orm:"null;size(32);"`
    JobId  string  `orm:"null;size(32);"`
    CreateDate  time.Time  `orm:"null;type(date);auto_now_add;"`
    UpdateDate  time.Time  `orm:"null;type(date);auto_now;"`
    CreateUser  string  `orm:"null;size(36);"`
    UpdateUser  string  `orm:"null;size(36);"`
    Remark  string  `orm:"null;size(255);"`
    Ver  int  `orm:"null;default(1);"`
    Status  int  `orm:"null;default(0);"`
}


