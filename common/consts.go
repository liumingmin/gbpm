package common

const GRunDebug string = "debug"
const GRunRelease string = "release"

const GMasterDB string = "default"
const GSlaveDB string = "slave"

//ajax
const ReqData = "Data"
const ReqUpdateCol = "DataUpdateColumn"
const ReqMobileSecret = "__secret"

const ReqProjectId = "ProjectId"
const ReqRelateModels = "RelateModels"

const RespCode = "Code"
const RespMsg = "Msg"
const RespContent = "Content"

const TreeRootId = "__root__"

//page
const PageRespOk = "isOk"
const PageRespCreate = "isCreate"
const PageRespEdit = "isEdit"
const PageRespMsg = "msg"
const PageRespContent = "data"
const PageRespInputTag = "inputTag"

const PageCaptchaId = "__captchaId"
const PageCaptchaCode = "__captchaCode"

const ClientType = "clientType"
const ClientTypeMobile = "mobile"
const JwtKey = "Dv12NpJfAr5WQ2St93r73UrAw5I2GHfVAy7MRgbqpknBxfB8LUdboIQGVI3KWIGw"

const SuperAdminId = "admin0000"

const (
	ErrMsgCreateProcess = "用户没有发起流程权限"
	ErrMsgSignalEnd     = "用户的步骤已经结束"
)
