package pkg

import (
	"alter/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Message struct {
	MsgType string `json:"msgtype"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
}

func SendWeixin(g config.Server, body string) {
	var m Message
	m.MsgType = "text"
	m.Text.Content = fmt.Sprintf("Weinxin告警通知:\t %s", body)
	jsons, err := json.Marshal(m)
	if err != nil {
		log.Fatalln("Weixin:", "发送失败", "##", "SendMessage Marshal failed:", err)
		return
	}
	resp := string(jsons)
	client := &http.Client{}
	req, err := http.NewRequest("POST", g.Weixin.Webhook, strings.NewReader(resp))
	if err != nil {
		log.Fatalln("Weixin:", "发送失败", "##", "Send Weixin Failed!Err:", err)
	} else {
		log.Println("Weixin:", "发送成功", "##", "Send Weixin Successfully!")
	}

	req.Header.Set("Content-Type", "application/json")
	r, err := client.Do(req)
	if err != nil {
		log.Fatalln("Weixin:", "发送失败", "##", "SendMessage client Do failed:", err)
		return
	}
	defer r.Body.Close()
	_, err = ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln("Weixin:", "发送失败", "##", "SendMessage ReadAll Body failed:", err)
		return
	}
}
