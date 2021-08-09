package pkg

import (
	"alter/config"
	"github.com/labulaka521/crocodile/common/notify/dingding"
	"log"
)

func SendDingding(g config.Server, body string) {
	ding := dingding.NewDing(g.Dingding.Webhook, 1, "")
	err := ding.Send([]string{""}, "DingDing告警通知:", body)
	if err != nil {
		log.Fatalln("dingding:", "发送失败", "##", "Send Dingding Failed!Err:", err)
	} else {
		log.Println("dingding:", "发送成功", "##", "Send Dingding Successfully!")
	}
}
