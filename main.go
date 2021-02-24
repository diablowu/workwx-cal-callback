package main

import (
	"diablowu/workwx-cal-callback/router"
	"diablowu/workwx-cal-callback/wxapi"
	"flag"
	"github.com/gin-gonic/gin"
)

var (
	corpId         = flag.String("corp-id", "", "corp id")
	callBackToken  = flag.String("token", "", "日程应用接收事件的Token")
	callBackAesKey = flag.String("aes-key", "", "日程应用接收事件的EncodingAESKey")
	agentId        = flag.String("agent-id", "", "日程应用的应用AgentId")
	secret         = flag.String("secret", "", "日程应用的Secret")
	bind           = flag.String("bind", ":9191", "listening address")
)

func main() {

	flag.Parse()

	wxapi.Init(func(context *wxapi.Context) {
		context.AesKey = *callBackAesKey
		context.Token = *callBackToken
		context.AgentId = *agentId
		context.Secret = *secret
		context.CorpId = *corpId
	})
	r := gin.New()
	//r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/callback", router.CallbackVerify)
	r.POST("/callback", router.HandleEvent)
	r.POST("/cal/create", router.CreateCal)

	r.Run(*bind)
}
