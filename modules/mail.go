package modules

import (
	"log"
	"net/smtp"

	"github.com/tvpsh2020/anime-crawler/config"
)

func mailToDest(title string) {
	auth := smtp.PlainAuth("", config.SMTP.Setting.Username, config.SMTP.Setting.Password, config.SMTP.Setting.Server)
	to := []string{config.SMTP.Setting.To}
	msg := []byte("To: " + config.SMTP.Setting.To + "\r\n" +
		"Subject: 你關注的動畫有新作品上架囉!!\r\n" +
		"\r\n" +
		"如標題，不囉嗦。\r\n" + "標題內容 : " + title + "\r\n")
	err := smtp.SendMail(config.SMTP.Setting.Server+":"+config.SMTP.Setting.Port, auth, config.SMTP.Setting.From, to, msg)
	if err != nil {
		log.Fatal(err)
	}
}
