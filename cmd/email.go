package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

func setupEmail() error {
	fmt.Println("In order to get notified by email, please specify your email smtp details")
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("SMTP Server: ")
	smtpServer, err := reader.ReadString('\n')
	if err != nil {
		return errors.New("Unable to read smtp server")
	}

	fmt.Print("SMTP Port: ")
	smtpPort, err := reader.ReadString('\n')
	if err != nil {
		return errors.New("Unable to read smtp port")
	}

	fmt.Print("Email address: ")
	email, err := reader.ReadString('\n')
	if err != nil {
		return errors.New("Unable to read email address")
	}

	fmt.Print("Password: ")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return errors.New("Unable to read password")
	}
	password := string(bytePassword)

	smtpServer = strings.TrimSuffix(smtpServer, "\n")
	smtpPort = strings.TrimSuffix(smtpPort, "\n")
	email = strings.TrimSuffix(email, "\n")
	password = strings.TrimSuffix(password, "\n")

	fmt.Println(smtpServer, smtpPort, email, password)
	return nil
}
