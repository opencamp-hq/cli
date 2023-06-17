package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/opencamp-hq/core/notify"
	"golang.org/x/crypto/ssh/terminal"
)

func GetSMTPConfig() (*notify.SMTPConfig, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("In order to get notified by email, please specify your email SMTP details")

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

	return &notify.SMTPConfig{
		Host:     strings.TrimSuffix(host, "\n"),
		Port:     strings.TrimSuffix(port, "\n"),
		Email:    strings.TrimSuffix(email, "\n"),
		Password: strings.TrimSuffix(password, "\n"),
	}, nil
}
