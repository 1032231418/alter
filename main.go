package main

import (
	"alter/config"
	"alter/pkg"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"os"
)

var (
	GVA_CONFIG config.Server
)

func main() {
	//接受参数
	args := os.Args //获取用户输入的所有参数
	//args的类型是[]string

	configfile := args[1]
	t := args[2]       //获取输入的第一个参数
	subject := args[3] //获取输入的第二个参数
	body := args[4]    //获取输入的第二个参数

	v := viper.New()
	v.SetConfigFile(configfile)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&GVA_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&GVA_CONFIG); err != nil {
		fmt.Println(err)
	}

	switch {
	case t == "mail":
		if subject != "" && body != "" {
			log.Println(t, "", "##", "通道发送")
			pkg.SendMail(GVA_CONFIG, subject, body)
		} else {
			log.Println(t, "未传入body信息", "##", "")
		}
	case t == "dingding":
		if body != "" {
			log.Println(t, "", "##", "通道发送")
			pkg.SendDingding(GVA_CONFIG, body)
		} else {
			log.Println(t, "未传入body信息", "##", "")
		}
	case t == "weixin":
		if body != "" {
			log.Println(t, "", "##", "通道发送")
			pkg.SendWeixin(GVA_CONFIG, body)
		} else {
			log.Println(t, "未传入body信息", "##", "")
		}
	default:
		fmt.Println("没有接收合法参数，使用-h 获得帮助信息。")
	}
}
