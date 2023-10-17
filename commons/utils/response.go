package utils

import "github.com/gin-gonic/gin"

type httpContext struct {
	c *gin.Context
}

type Response struct {
	Code        int         `json:"code"`
	Success     bool        `json:"success"`
	Message     string      `json:"message"`
	MessageCode string      `json:"message_code"`
	Data        interface{} `json:"data"`
	Errors      interface{} `json:"errors,omitempty"`
}

type PaginationResponse struct {
	Response
	PaginationResult
}

func ToJSON(c *gin.Context) *httpContext {
	return &httpContext{c}
}

func (hc httpContext) CustomResponse(code int, success bool, msg string, msgCode string, data interface{}, errors interface{}) {
	hc.c.JSON(code, &Response{
		Code:        code,
		Success:     success,
		Message:     msg,
		MessageCode: msgCode,
		Data:        data,
		Errors:      errors,
	})
}

func (hc httpContext) PaginationResponse(resp *PaginationResponse) {
	hc.c.JSON(resp.Code, resp)
}
