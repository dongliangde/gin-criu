package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Data    interface{} `json:"data" comment:"Return the data"`     //返回数据
	Code    int         `json:"code" comment:"The response status"` //响应状态
	Message string      `json:"msg" comment:"The response message"` //响应消息
}

func ReturnJson(Context *gin.Context, httpCode int, dataCode int, msg string, data interface{}) {

	Context.JSON(httpCode, gin.H{
		"code": dataCode,
		"msg":  msg,
		"data": data,
	})
}

// Json characters are returned in standard JSON format
func ReturnJsonFromString(Context *gin.Context, httpCode int, jsonStr string) {
	Context.Header("Content-Type", "application/json; charset=utf-8")
	Context.String(httpCode, jsonStr)
}

func Success(c *gin.Context, msg string, data interface{}) {
	ReturnJson(c, http.StatusOK, http.StatusOK, msg, data)
}

func Fail(c *gin.Context, dataCode int, msg string, data interface{}) {
	ReturnJson(c, http.StatusBadRequest, dataCode, msg, data)
	c.Abort()
}
