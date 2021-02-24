package router

import (
	"diablowu/workwx-cal-callback/wxapi"
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"log"
)

func CallbackVerify(c *gin.Context) {
	sign := c.Query("msg_signature")
	timestamp := c.Query("timestamp")
	nonce := c.Query("nonce")
	echostr := c.Query("echostr")

	crypto := wxapi.NewWXBizMsgCrypt(wxapi.DefaultAPI.Context.Token, wxapi.DefaultAPI.Context.AesKey, wxapi.DefaultAPI.Context.CorpId, wxapi.XmlType)
	echostrExplain, err := crypto.VerifyURL(sign, timestamp, nonce, echostr)
	if err != nil {
		panic(err)
	}

	c.Data(200, "text/html", []byte(echostrExplain))

}

func HandleEvent(c *gin.Context) {
	sign := c.Query("msg_signature")
	timestamp := c.Query("timestamp")
	nonce := c.Query("nonce")
	crypto := wxapi.NewWXBizMsgCrypt(wxapi.DefaultAPI.Context.Token, wxapi.DefaultAPI.Context.AesKey, wxapi.DefaultAPI.Context.CorpId, wxapi.XmlType)
	buf, err := c.GetRawData()
	if err != nil {
		panic(err)
	}
	if buf, err := crypto.DecryptMsg(sign, timestamp, nonce, buf); err != nil {
		panic(err)
	} else {
		var eventData = new(wxapi.CalEvent)
		if err := xml.Unmarshal(buf, eventData); err != nil {
			panic(err)
		} else {
			log.Printf("event :\n %s \n", eventData.String())
		}
	}
}

func CreateCal(c *gin.Context) {

	owner := c.PostForm("owner")
	readOnly := c.PostForm("readonly")
	title := c.PostForm("title")

	bs, err := wxapi.NewCalendar(owner, title, readOnly)
	if err != nil {
		c.Error(err)
	} else {
		c.Data(200, "application/json", bs)
	}

}
