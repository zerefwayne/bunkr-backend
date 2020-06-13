package config

import (
	"log"

	"github.com/sendgrid/sendgrid-go"
)

func (c *Config) ConnectSendGrid() {

	c.SendGrid = sendgrid.NewSendClient(c.Env.SendGridEnv.Key)
	log.Println("sendgrid	| sendgrid connected successfully!", c.Env.SendGridEnv.Key)

}
