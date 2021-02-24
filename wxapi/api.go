package wxapi

import (
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"io/ioutil"
	"net/http"
)

type Context struct {
	Token       string
	AesKey      string
	AgentId     string
	Secret      string
	CorpId      string
	AccessToken string
}

type WxAPI struct {
	Context *Context
}

var DefaultAPI *WxAPI

func Init(callback func(*Context)) {
	DefaultAPI = new(WxAPI)
	ctx := new(Context)
	callback(ctx)
	initAccessToken(ctx)
	DefaultAPI.Context = ctx

	spew.Dump(DefaultAPI)
}

func initAccessToken(ctx *Context) {
	if resp, err := http.Get(fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s", ctx.CorpId, ctx.Secret)); err != nil {
		panic(err)
	} else {
		bs, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		var ret map[string]interface{} = make(map[string]interface{})
		err = json.Unmarshal(bs, &ret)
		if err != nil {
			panic(err)
		}
		ctx.AccessToken = ret["access_token"].(string)
	}
}
