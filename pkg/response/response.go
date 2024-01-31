package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type page struct {
	Count     int         `json:"count"`
	PageIndex int         `json:"pageIndex"`
	PageSize  int         `json:"pageSize"`
	List      interface{} `json:"list"`
}

func OK(c *gin.Context, message string, data interface{}) {
	c.AbortWithStatusJSON(http.StatusOK, response{
		Code:    http.StatusOK,
		Message: message,
		Data:    data,
	})
}

func Error(c *gin.Context, code int, err error, msg string) {
	res := response{}
	res.Code = code
	if err != nil {
		res.Message = err.Error()
	}
	if msg != "" {
		res.Message = msg
	}
	c.AbortWithStatusJSON(http.StatusOK, res)
}

func PageOK(c *gin.Context, result interface{}, count int, pageIndex int, pageSize int, msg string) {
	var res page
	res.List = result
	res.Count = count
	res.PageIndex = pageIndex
	res.PageSize = pageSize
	OK(c, msg, res)
}
