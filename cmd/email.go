package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"syscall"

	"github.com/opencamp-hq/core/notify"
	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh/terminal"
)

func GetSMTPConfig() (*notify.SMTPConfig, error) {
	cfg := &notify.SMTPConfig{
		Host:     viper.GetString("smtp.host"),
		Port:     viper.GetString("smtp.port"),
		Email:    viper.GetString("smtp.email"),
		Password: viper.GetString("smtp.password"),
	}
	if cfg.Valid() {
		return cfg, nil
	}

	var err error
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("In order to get notified by email, please specify your email SMTP details\n")

	fmt.Print("SMTP Server: ")
	cfg.Host, err = reader.ReadString('\n')
	if err != nil {
		return nil, errors.New("Unable to read smtp server")
	}

	fmt.Print("SMTP Port: ")
	cfg.Port, err = reader.ReadString('\n')
	if err != nil {
		return nil, errors.New("Unable to read smtp port")
	}

	fmt.Print("Email address: ")
	cfg.Email, err = reader.ReadString('\n')
	if err != nil {
		return nil, errors.New("Unable to read email address")
	}

	fmt.Print("Password: ")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return nil, errors.New("Unable to read password")
	}
	cfg.Password = string(bytePassword)
	fmt.Print("\n\n")

	return cfg, nil
}
