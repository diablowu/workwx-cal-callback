package wxapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

const wxAPIEndpoint = "https://qyapi.weixin.qq.com"

type PersonalCalendar struct {
	Calendar struct {
		Organizer   string `json:"organizer"`
		ReadOnly    int    `json:"readonly"`
		Default     int    `json:"set_as_default"`
		Summary     string `json:"summary"`
		Color       string `json:"color"`
		Description string `json:"description"`
	} `json:"calendar"`
	AgentID string `json:"agentid"`
}

type CalEvent struct {
	To         string `xml:"ToUserName"`
	From       string `xml:"FromUserName"`
	createTime string `xml:"CreateTime"`
	MsgType    string `xml:"MsgType"`
	Event      string `xml:"Event"`
	CalId      string `xml:"CalId"`
	ScheduleId string `xml:"ScheduleId"`
}

func (e CalEvent) String() string {
	bs, _ := json.MarshalIndent(e, "", "  ")
	return string(bs)
}

func NewCalendar(owner, title, readonly string) ([]byte, error) {
	cal := new(PersonalCalendar)
	cal.Calendar.Color = "#FF3030"
	cal.Calendar.ReadOnly, _ = strconv.Atoi(readonly)
	cal.Calendar.Default = 0
	cal.Calendar.Organizer = owner
	cal.Calendar.Summary = title
	cal.Calendar.Description = title
	cal.AgentID = DefaultAPI.Context.AgentId

	bs, _ := json.Marshal(cal)
	resp, err := http.Post(fmt.Sprintf(wxAPIEndpoint+"/cgi-bin/oa/calendar/add?access_token=%s", DefaultAPI.Context.AccessToken), "application/json", bytes.NewReader(bs))

	if err != nil {
		return nil, err
	} else {
		bs, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		} else {
			return bs, nil
		}
	}
}

