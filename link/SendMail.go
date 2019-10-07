package main

import (
	"fmt"
	"github.com/unknwon/com"
	"github.com/astaxie/beego"
	"github.com/scorredoira/email"
	"net/mail"
	"net/smtp"
	"time"
)

func main() {
	from := beego.AppConfig.String("from")
	password := beego.AppConfig.String("password")
	smtpport := beego.AppConfig.String("smtpport")
	smtpaddress := beego.AppConfig.String("smtpaddress")
	tolist := beego.AppConfig.String("tolist")
	subject := beego.AppConfig.String("subject")

	body := beego.AppConfig.String("subject")
	fmt.Println(body)
	m := email.NewMessage(subject, body)

	m.From = mail.Address{Name: "一花一世界", Address: from}
	m.To = []string{tolist}
	fmt.Println(m.To)

	//检查附件是否存在
	attachpath := beego.AppConfig.String("attachpath")
	attachfilename := beego.AppConfig.String("attachfilename")
	currentdate := time.Now().Format("20060102")
	attachfile := attachpath + "\\" + attachfilename + "-" + currentdate + "_" + "v1.0" + ".rar"
	fmt.Println("附件名:" + attachfile)

	if isexist := com.IsExist(attachfile); isexist == false {
		fmt.Println("no attach")
		return
	}

	if err := m.Attach(attachfile); err != nil {
		fmt.Println("append attach error")
		return
	}

	auth := smtp.PlainAuth("", from, password, smtpaddress)
	err := email.Send(smtpport, auth, m)
	if err != nil {
		fmt.Println("send mail error!")
		fmt.Println(err)
	} else {
		fmt.Println("send mail success!")
	}

}
