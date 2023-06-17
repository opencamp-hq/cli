package email

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"net/smtp"
	"os"
	"strings"
	"syscall"
	"text/template"
	"time"

	"github.com/opencamp-hq/core/models"
	"golang.org/x/crypto/ssh/terminal"
)

type SMTPSender struct {
	template *template.Template
	host     string
	port     string
	email    string
	password string
}

func NewSMTPSender() (*SMTPSender, error) {
	t, err := template.New("email").Parse(emailTemplate)
	if err != nil {
		return nil, errors.New("Internal error: Unable to parse email template")
	}

	fmt.Println("In order to get notified by email, please specify your email SMTP details")
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("SMTP Server: ")
	host, err := reader.ReadString('\n')
	if err != nil {
		return nil, errors.New("Unable to read smtp server")
	}

	fmt.Print("SMTP Port: ")
	port, err := reader.ReadString('\n')
	if err != nil {
		return nil, errors.New("Unable to read smtp port")
	}

	fmt.Print("Email address: ")
	email, err := reader.ReadString('\n')
	if err != nil {
		return nil, errors.New("Unable to read email address")
	}

	fmt.Print("Password: ")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return nil, errors.New("Unable to read password")
	}
	password := string(bytePassword)

	return &SMTPSender{
		template: t,
		host:     strings.TrimSuffix(host, "\n"),
		port:     strings.TrimSuffix(port, "\n"),
		email:    strings.TrimSuffix(email, "\n"),
		password: strings.TrimSuffix(password, "\n"),
	}, nil
}

func (s SMTPSender) Send(cg models.Campground, start, end time.Time, sites models.Campsites) error {
	var buf bytes.Buffer

	auth := smtp.PlainAuth("", s.email, s.password, s.host)

	// message :=

	return nil
}
