package serializer

import "github.com/gin-gonic/gin"

// 三位数的错误编码为复用http原本含义
// 五位数错误编码为应用自定义错误
// 五开头的五位数错误编码为服务器端错误
// 四开头的五位数错误编码为客户端错误
const (
	// CodeCheckLogin 未登录
	CodeCheckLogin = 401

	// CodeNoRightErr 未授权访问
	CodeNoRightErr = 403

	// CodeDBError 数据库操作失败
	CodeDBError = 50001

	//C odeEncryptError 加密失败
	CodeEncryptError = 50002

	// CodeParamErr 其他错误
	CodeParamErr = 40001
)

// Response 基础序列化器
type Response struct {
	Status 	int 		`json:"code"`
	Data 	interface{} `json:"data"`
	Msg		string		`json:"msg"`
	Error	string		`json:"error"`
}

// DataList 基础列表结构
type DataList struct {
	Items	interface{}	`json:"items"`
	Total	uint		`json:"total"`
}

// TrackedErrorResponse 有追踪信息的错误响应
type TrackedErrorResponse struct {
	Response
	TrackID		string	`json:track_id:`
}

// BuildListResponse 列表构建器
func BuildListResponse(items interface{}, total uint) Response {
	return Response{
		Data:	DataList{
			Items:	items,
			Total:	total,
			},
	}
}

func CheckLogin() Response {
	return Response{
		Status:	CodeCheckLogin,
		Msg:	"未登录",
	}
}

// Err 通常错误处理
func Err(errCode int, msg string, err error) Response {
	res := Response {
		Status:	errCode,
		Msg:	msg,
	}

	// 生产环境隐藏底层报错
	if err != nil && gin.Mode() != gin.ReleaseMode {
		res.Error = err.Error()
	}
	return res
}

// DBErr 数据库操作失败
func DBErr(msg string, err error) Response {
	if msg == "" {
		msg = "数据库操作失败"
	}
	return Err(CodeDBError, msg, err)
}

// ParamErr 各种参数错误
func ParamErr(msg string, err error) Response {
	if msg == "" {
		msg = "参数错误"
	}
	return Err(CodeParamErr, msg, err)
}