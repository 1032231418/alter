package pkg

import (
	"alter/config"
	"bufio"
	"bytes"
	"fmt"
	"gopkg.in/gomail.v2"
	"log"
	"mime"
	"strconv"
	"strings"
	"sync"
)

func MailTo(mailTo string, subject, body string, g config.Server, wg *sync.WaitGroup) {
	defer wg.Done()
	mailConn := map[string]string{
		"username": g.Mail.Username,
		"authCode": g.Mail.Password,
		"host":     g.Mail.Host,
		"port":     g.Mail.Port,
	}

	var metrics bytes.Buffer
	metrics.WriteString(body)
	scanner := bufio.NewScanner(&metrics)
	bodyHtml := "<h1>Alter Messages:<br/>"
	for scanner.Scan() {
		bodyHtml += fmt.Sprintf("%s<br/>", scanner.Text())
	}
	port, _ := strconv.Atoi(mailConn["port"])
	m := gomail.NewMessage()
	m.SetHeader("From", mime.QEncoding.Encode("UTF-8", "Alter")+"<"+mailConn["username"]+">")
	m.SetHeader("To", mailTo)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", bodyHtml)
	d := gomail.NewDialer(mailConn["host"], port, mailConn["username"], mailConn["authCode"])
	err := d.DialAndSend(m)
	if err != nil {
		log.Fatalln("Mail To:", mailTo, "##", "Send Email Failed!Err:", err)
	} else {
		log.Println("Mail To:", mailTo, "##", "Send Email Successfully!")
	}
}

func SendMail(g config.Server, subject, body string) {
	var wg sync.WaitGroup
	receiver := strings.Split(g.Mailto.Mails, ",")
	for _, mail := range receiver {
		wg.Add(1)
		go MailTo(mail, subject, body, g, &wg)
	}
	wg.Wait()
}
