package server

import (
	"fmt"
	"log"
	"strings"

	"github.com/emersion/go-sasl"
	"github.com/emersion/go-smtp"
	"gitlab.com/meta-node/mail/core/entities"
	"gitlab.com/meta-node/mail/helper"
)

type MailHandler struct {
	Flag  chan error
	Email *entities.Email
	Auth  sasl.Client
}

func (mh *MailHandler) SendmailHandler() {
	for _, receiver := range mh.Email.To {
		conn := helper.ClassifyEmail(receiver)
		switch conn.Classify {
		case "System":
			{
				fmt.Println("hí hí hí hí hí hí hí")
				go mh.sendSystemEmail(conn)
			}
		case "Outside":
			{
				fmt.Println("hố hố hố ................")
				go mh.sendOutsideEmail(conn)
			}
		}
	}
}

func (mh *MailHandler) sendSystemEmail(conn *helper.Connection) {
	// address := fmt.Sprintf("%v:%v", conn.Host, conn.Port)
	// content := strings.NewReader(mh.Email.Content)
	// fmt.Println(address)
	// if err := smtp.SendMail(address, mh.Auth, mh.Email.From, conn.Receiver, content); err != nil {
	// 	log.Println(err)
	// }
}

func (mh *MailHandler) sendOutsideEmail(conn *helper.Connection) {
	address := fmt.Sprintf("%v:%v", conn.Host, conn.Port)
	fmt.Println(mh.Auth)
	content := strings.NewReader(mh.Email.Content)
	if err := smtp.SendMail(address, mh.Auth, mh.Email.From, conn.Receiver, content); err != nil {
		log.Println("lỗi này:", err)
	}
}

// func (mh MailHandler) ReceivemailHandler() {}
