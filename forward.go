package main

import (
	"bytes"
	"crypto/tls"
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

func forwardSmtp(client *MailConnection) {

	log.Println(*forwardsmtp)
	c, err := smtp.Dial(*forwardsmtp)

	host, _, _ := net.SplitHostPort(*forwardsmtp)
	tlc := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}
	if err := c.StartTLS(tlc); err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	// Set the sender and recipient.
	c.Mail(client.From)
	c.Rcpt(client.To)
	// Send the email body.
	wc, err := c.Data()
	if err != nil {
		log.Fatal(err)
	}
	defer wc.Close()
	buf := bytes.NewBufferString(client.Data)
	if _, err = buf.WriteTo(wc); err != nil {
		log.Fatal(err)
	}
}
