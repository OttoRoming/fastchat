package main

import (
	"net"

	"github.com/OttoRoming/fastchat/cmd/fcserver/data"
	"github.com/OttoRoming/fastchat/pkg/fcprotocol"
	"github.com/alexedwards/argon2id"
	"github.com/charmbracelet/log"
)

const (
	address = "localhost:4040"
)

var (
	responseDatabaseError = fcprotocol.ResponseError{
		Message: "internal database error",
	}
	responseServerError = fcprotocol.ResponseError{
		Message: "internal server error",
	}
)

func handleSignUp(request *fcprotocol.RequestSignUp) fcprotocol.Response {
	previousAccount, err := data.GetAccountByUsername(request.Username)
	if err != nil {
		log.Error("database error", "err", err)
		return responseDatabaseError
	}
	if previousAccount != nil {
		return fcprotocol.ResponseError{
			Message: "account with username already exsists",
		}
	}

	hash, err := argon2id.CreateHash(request.Password, argon2id.DefaultParams)
	if err != nil {
		log.Error("failed to hash password", "err", err)
		return responseServerError
	}

	account, err := data.AddAccount(request.Username, hash)
	if err != nil {
		log.Error("database error", "err", err)
		return responseDatabaseError
	}
	log.Info("created account", "id", account.ID, "username", account.Username)

	return fcprotocol.ResponseSignedIn{
		Token: account.Token,
	}
}

func handleLogIn(request *fcprotocol.RequestLogin) fcprotocol.Response {
	account, err := data.GetAccountByUsername(request.Username)
	if err != nil {
		return responseDatabaseError
	}
	if account == nil {
		return fcprotocol.ResponseError{
			Message: "account with username not found",
		}
	}

	match, _, err := argon2id.CheckHash(
		request.Password,
		account.Password,
	)
	if err != nil {
		log.Error("failed to check password", "err", err)
		return responseServerError
	}
	if !match {
		return fcprotocol.ResponseError{
			Message: "incorrect password",
		}
	}

	return fcprotocol.ResponseSignedIn{
		Token: account.Token,
	}
}

func handleRequest(request fcprotocol.Request) fcprotocol.Response {
	if !request.Confidential() {
		log.Info("received request", "request", request)
	}

	switch r := request.(type) {
	case *fcprotocol.RequestMOTD:
		return fcprotocol.ResponseMOTD{
			MOTD: "Welcome to the fcserver",
		}
	case *fcprotocol.RequestSignUp:
		return handleSignUp(r)
	case *fcprotocol.RequestLogin:
		return handleLogIn(r)
	default:
		return fcprotocol.ResponseError{
			Message: "unknown request method",
		}
	}
}

func handleConnection(conn net.Conn) {
	for {
		fcprotocol.HandleRequest(handleRequest, conn)
	}
}

func main() {
	data.LoadDB()
	defer data.Close()

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Server listening", "address", address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Warn("connection failed", "err", err)
			continue
		}

		go handleConnection(conn)
	}
}
