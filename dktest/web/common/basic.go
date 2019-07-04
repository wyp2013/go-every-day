package common

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-xorm/xorm"
	"github.com/labstack/echo"
)

/*
 * 自定义上下文
 * log句柄
 * xorm句柄
 *
 */
type CustomContext struct {
	echo.Context `json:"-"`
	DbEngine     *xorm.Engine     `json:"-"`
	RequestId    string
	ActionId     string
	Body         []byte `json:"-"`
	WriteLogId   int
	RspCode      int
	StartTime    time.Time
	ErrMsg       string      `json:"-"`
}


type Output struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ConsulResult struct {
	Errno  int         `json:"errno"`
	Errmsg string      `json:"errmsg"`
	Desc   interface{} `json:"data"`
}

func ParseConsulResult(result string) (*ConsulResult, error) {
	r := ConsulResult{}
	err := json.Unmarshal([]byte(result), &r)

	if err != nil {
		return nil, err
	}

	return &r, nil
}

func newOutput(code int, data interface{}) *Output {
	msg, ok := ErrnoDesc[code]
	if !ok {
		msg = "Unknown Error"
	}

	return &Output{Code: code, Message: msg, Data: data}
}


func (e *CustomContext) SendRsp(code int, data interface{}) error {
	ret := newOutput(code, data)
	e.RspCode = code

	return e.JSON(http.StatusOK, ret)
}

func (e *CustomContext) SendRspErrLock(taskId int) error {
	data := fmt.Sprintf("{\"taskId\" : \"%d\"}", taskId)
	return e.SendRsp(ERR_LOCK, data)
}

//new begin
func NewOutput(code int, data interface{}, traceId string, errMsg string) *Output {
	msg, ok := ErrnoDesc[code]
	if !ok {
		msg = "Unknown Error"
	}
	if errMsg != "" {
		msg = errMsg
	}

	msg = fmt.Sprintf("%s: traceId[%s]", msg, traceId)

	return &Output{Code: code, Message: msg, Data: data}
}

func (e *CustomContext) SendRspInternal(code int, data interface{}, errMsg string) error {
	ret := NewOutput(code, data, e.RequestId, errMsg)

	e.RspCode = code

	return e.JSON(http.StatusOK, ret)
}

func (e *CustomContext) SendRspOK(data interface{}) error {
	return e.SendRspInternal(ERR_OK, data, "")
}

func (e *CustomContext) SendRspErr(code int, data interface{}) error {
	return e.SendRspInternal(code, data, "")
}

func (e *CustomContext) SendRspErrMsg(code int, data interface{}) error {
	return e.SendRspInternal(code, data, e.ErrMsg)
}

func (e *CustomContext) SetCommonErrMsg(errMsg string) {
	e.ErrMsg = errMsg
	return
}

func (e *CustomContext) GetRequstID() string {
	//return e.Response().Header().Get(echo.HeaderXRequestID)
	return e.RequestId
}


func ObjToStr(data interface{}) string {
	s, _ := json.Marshal(data)
	return string(s)
}

