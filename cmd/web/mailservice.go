package main

import (
	"log"
	"time"

	mail "github.com/xhit/go-simple-mail/v2"

	"github.com/youssef-aly1996/bookings/internal/models"
)

func listenForMail()  {
	go func() {
		m := <-appConfig.MailChan
		sendMail(m)
	}()
}

func sendMail(m models.MailModel)  {
	server := mail.NewSMTPClient()
	server.Host = "localhost"
	server.Port = 1025
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second
	client, err := server.Connect()
	if err != nil {
		log.Println(err)
	}
	email := mail.NewMSG()
	email.SetFrom(m.From).AddTo(m.To).SetSubject(m.Subject)
	email.SetBody(mail.TextHTML, m.Content)

	err = email.Send(client)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("email is sent")
	}
}