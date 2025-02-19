package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 返回的结果：
type Result struct {
	Time time.Time   `json:"time"`
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// 成功
func Success(c *gin.Context, data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	res := Result{}
	res.Time = time.Now()
	res.Code = int(ApiCode.SUCCESS)
	res.Msg = ApiCode.GetMessage(ApiCode.SUCCESS)
	res.Data = data

	c.JSON(http.StatusOK, res)
}

// 出错
func Error(c *gin.Context, code int, data interface{}) {
	res := Result{}
	res.Time = time.Now()
	res.Code = code
	res.Msg = ApiCode.GetMessage(ApiCode.FAILED)
	res.Data = data
	c.JSON(http.StatusOK, res)
}
