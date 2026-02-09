package main

import (
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/log"
)

type model struct {
	choices  []string         // items on the to-do list
	cursor   int              // which to-do list item our cursor is pointing at
	selected map[int]struct{} // which to-do items are selected
}

const (
	colorBg0     = "#1d1d1d"
	colorBg1     = "#232323"
	colorBg2     = "#5e5e5e"
	colorFg      = "#e8e8e8"
	colorPrimary = "#d82934"
)

func login() error {
	const (
		signIn = "Sign in"
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
	err := login()
	if err != nil {
		log.Fatal(err)
	}
}
