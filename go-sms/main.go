package main

import (
	"LHXHL/go-sms/controller"
	"LHXHL/go-sms/model"
	"LHXHL/go-sms/service"
	"github.com/kataras/iris/v12"
	"github.com/spf13/viper"
)

func main() {
	readEnv()
	app := iris.Default()
	iris.RegisterOnInterrupt(func() {
		model.DbNow.Close()
	})
	registerApp(app)
	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<h1>Hello qqw!<h1>")
	})
	model.InitDB("sms.db")
	service.InitTxSms()
	app.Listen(":" + viper.GetString("PORT"))
}

func registerApp(app *iris.Application) {
	sms := app.Party("/sms")
	sms.Post("/new", controller.NewCode)
	sms.Post("/check", controller.CheckCode)
	sms.Get("/total", controller.Total)
	sms.Post("/clear", controller.FlushAll)

}

func readEnv() {
	viper.SetConfigFile("config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
