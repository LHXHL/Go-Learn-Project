package controller

import (
	"LHXHL/go-sms/service"
	"github.com/kataras/iris/v12"
	"github.com/spf13/viper"
)

type NewCodeRequest struct {
	Phone string `json:"phone"`
}

func NewCode(ctx iris.Context) {
	var req NewCodeRequest
	err := ctx.ReadJSON(&req)
	if err != nil {
		return
	}
	phone := req.Phone
	if !service.Sms.ValidPhone(phone) {
		CommonResponse{}.Fail().SetData("not a valid phone").Send(ctx)
		return
	}
	if service.Sms.IfExist(phone) {
		CommonResponse{}.Fail().SetData(
			iris.Map{
				"phone": phone,
				"msg":   "already send code",
			}).Send(ctx)
		return
	}

	_, err = service.Sms.GenCode(phone)
	if err != nil {
		return
	}
	CommonResponse{}.Success().SetData(
		iris.Map{
			"phone":  phone,
			"status": "success",
			"expire": viper.GetString("EXPIRE") + "s",
		}).Send(ctx)
}

type CheckCodeRequest struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}

func CheckCode(ctx iris.Context) {
	var req CheckCodeRequest
	err := ctx.ReadJSON(&req)
	if err != nil {
		return
	}
	phone := req.Phone
	code := req.Code
	if service.Sms.ValidCode(phone, code) {
		CommonResponse{}.Success().SetData("pass").Send(ctx)
		return
	}
	CommonResponse{}.Fail().SetData("The phone code is not valid")
}

func Total(ctx iris.Context) {
	CommonResponse{}.Success().SetData(iris.Map{
		"total": service.Sms.Total(),
	}).Send(ctx)
}

func FlushAll(ctx iris.Context) {
	service.Sms.ClearAll()
	CommonResponse{}.Success().SetData("flush success").Send(ctx)
}
