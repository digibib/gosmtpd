package main

import (
	"log"
	"net"
	"net/smtp"
)

// fake Auth by returning empty
type nillAuth struct{}

func (n nillAuth) Start(server *smtp.ServerInfo) (proto string, toServer []byte, err error) {
	return
}
func (n nillAuth) Next(fromServer []byte, more bool) (toServer []byte, err error) {
	return
}

func getForwardHost() string {
	return net.JoinHostPort(*forwardhost, *forwardport)
}

func forwardEmail(client *MailConnection) {
	err := smtp.SendMail(*forwardsmtp, nillAuth{}, client.From, []string{client.To}, []byte(client.Data))
	if err != nil {
		log.Fatal("ERROR %v", err)
	}
}
