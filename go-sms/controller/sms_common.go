package controller

import "github.com/kataras/iris/v12"

type CommonResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ICommonResponse interface {
	Success() CommonResponse
	Fail() CommonResponse
	SetData(data interface{}) CommonResponse
	Send(ctx iris.Context)
}

func (c CommonResponse) Success() CommonResponse {
	c.Code = 200
	c.Message = "success"
	return c
}

func (c CommonResponse) Fail() CommonResponse {
	c.Code = 500
	c.Message = "fail"
	return c
}

func (c CommonResponse) SetData(data interface{}) CommonResponse {
	c.Data = data
	return c
}

func (c CommonResponse) Send(ctx iris.Context) {
	err := ctx.JSON(c)
	if err != nil {
		return
	}
}