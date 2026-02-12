package main

import (
	"fmt"
	"net"

	"github.com/OttoRoming/fastchat/pkg/fcprotocol"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/log"
)

const (
	colorBg0     = "#1d1d1d"
	colorBg1     = "#232323"
	colorBg2     = "#5e5e5e"
	colorFg      = "#e8e8e8"
	colorPrimary = "#d82934"
)

func login() error {
	const (
		signIn = "Sign up"
		logIn  = "Log in"
	)

	var (
		chosenOption string
		username     string
		password     string
	)

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Choose how to log in").
				Options(
					huh.NewOption(logIn, logIn),
					huh.NewOption(signIn, signIn),
				).
				Value(&chosenOption),
			huh.NewInput().Title("Username").Value(&username),
			huh.NewInput().Title("Password").Value(&password).EchoMode(huh.EchoModePassword),
		),
	)

	return form.Run()
}

func selectUser() error {
	var username string
	form := huh.NewInput().
		Title("Start a chat with").
		Value(&username)
	return form.Run()
}

func main() {
	var disclaimerConfirm bool
	huh.NewConfirm().
		Title("All traffic (including passwords) are unencrypted. Your traffic using this application will most likely be caputerd and analyzed by your ISP, government agencies, your network administrator and more. Are you sure?").
		Affirmative("Yes!").
		Negative("No.").
		Value(&disclaimerConfirm).Run()

	if !disclaimerConfirm {
		log.Info("disclaimer not confirmed, exiting")
		return
	}

	conn, err := net.Dial("tcp", "localhost:4040")

	message := fcprotocol.RequestUptime{}
	err = fcprotocol.SendMessage(message, conn)
	if err != nil {
		log.Fatal("failed to send message", "err", err)
	}
	log.Info("message sent", "message", message)

	var msg fcprotocol.Message
	msg, err = fcprotocol.ReadMessage(conn)

	fmt.Printf("msg: %v\n", msg)

	err = login()
	if err != nil {
		log.Fatal(err)
	}
}
