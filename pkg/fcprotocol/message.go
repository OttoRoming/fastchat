package fcprotocol

import "math"

type Message interface {
	method() uint16
	Confidential() bool
}

type Request interface {
	Message
	requestTag()
}

type Response interface {
	Message
	responseTag()
}

const (
	requestUptime uint16 = iota
	requestSignUp
	requestLogIn
	requestSendChat
	requestChatHistory
	requestContacts

	responseMOTD uint16 = 1<<15 + iota
	responseSignedIn
	responseChatSent
	responseChatHistory
	responseContacts

	responseError = math.MaxUint16
)

// Uptime related messages
type RequestMOTD struct{}

func (RequestMOTD) method() uint16 {
	return requestUptime
}

func (RequestMOTD) Confidential() bool {
	return false
}

func (RequestMOTD) requestTag() {}

type ResponseMOTD struct {
	MOTD string
}

func (ResponseMOTD) method() uint16 {
	return responseMOTD
}

func (ResponseMOTD) Confidential() bool {
	return false
}

func (ResponseMOTD) responseTag() {}

// Account related messages

type RequestSignUp struct {
	Username string
	Password string
}

func (RequestSignUp) method() uint16 {
	return requestSignUp
}

func (RequestSignUp) Confidential() bool {
	return true
}

func (RequestSignUp) requestTag() {}

type RequestLogin struct {
	Username string
	Password string
}

func (RequestLogin) method() uint16 {
	return requestLogIn
}

func (RequestLogin) Confidential() bool {
	return true
}

func (RequestLogin) requestTag() {}

type ResponseSignedIn struct {
	Token string
}

func (ResponseSignedIn) method() uint16 {
	return responseSignedIn
}

func (ResponseSignedIn) Confidential() bool {
	return false
}

func (ResponseSignedIn) responseTag() {}

// Chat related messages
type RequestSendChat struct {
	Token string

	To      string
	Content string
}

func (RequestSendChat) method() uint16 {
	return requestSendChat
}

func (RequestSendChat) Confidential() bool {
	return false
}

func (RequestSendChat) requestTag() {}

type ResponseMessageSent struct {
}

func (ResponseMessageSent) method() uint16 {
	return responseChatSent
}

func (ResponseMessageSent) Confidential() bool {
	return false
}

func (ResponseMessageSent) responseTag() {}

// Chat history related emthods
type RequestChatHistory struct {
	Token string

	To string
}

func (RequestChatHistory) method() uint16 {
	return requestChatHistory
}

func (RequestChatHistory) Confidential() bool {
	return false
}

func (RequestChatHistory) requestTag() {}

type ResponseChatHistory struct {
	Chats []struct {
		To      string
		Content string
	}
}

func (ResponseChatHistory) method() uint16 {
	return responseChatHistory
}

func (ResponseChatHistory) Confidential() bool {
	return false
}

func (ResponseChatHistory) responseTag() {}

type ResponseError struct {
	Message string
}

func (ResponseError) method() uint16 {
	return responseError
}

func (ResponseError) Confidential() bool {
	return false
}

func (ResponseError) responseTag() {}

type RequestContacts struct {
	Token string
}

func (RequestContacts) method() uint16 {
	return requestContacts
}

func (RequestContacts) Confidential() bool { return false }

func (RequestContacts) requestTag()

type ResponseContacts struct {
	contacts []string
}

func (ResponseContacts) method() uint16 {
	return responseContacts
}

func (ResponseContacts) Confidential() bool { return false }

func (ResponseContacts) requestTag()
