package fcprotocol

type Message interface {
	method() uint16
	Confidential() bool
}

const (
	// Uptime related methods
	requestUptime uint16 = iota
	responseUptime

	// Account related methods
	requestSignUp
	requestLogIn
	responseSignedIn
	errorUsernameTaken

	// Chat related methods
	requestSendChat
	responseChatSent

	// Chat history related methods
	requestChatHistory
	responseChatHistory

	// Generic
	errorAccountNotFound
	errorFailedRead
)

// Uptime related messages
type RequestUptime struct{}

func (RequestUptime) method() uint16 {
	return requestUptime
}

func (RequestUptime) Confidential() bool {
	return false
}

type ResponseUptime struct {
	Uptime string
}

func (ResponseUptime) method() uint16 {
	return responseUptime
}

func (ResponseUptime) Confidential() bool {
	return false
}

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

type ResponseSignedIn struct {
	Token string
}

func (ResponseSignedIn) method() uint16 {
	return responseSignedIn
}

func (ResponseSignedIn) Confidential() bool {
	return false
}

type ErrorUsernameTaken struct{}

func (ErrorUsernameTaken) method() uint16 {
	return errorUsernameTaken
}

func (ErrorUsernameTaken) Confidential() bool {
	return false
}

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

type ResponseMessageSent struct {
}

func (ResponseMessageSent) method() uint16 {
	return responseChatSent
}

func (ResponseMessageSent) Confidential() bool {
	return false
}

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

// Generic
type ErrorAccountNotFound struct{}

func (ErrorAccountNotFound) method() uint16 {
	return errorAccountNotFound
}

func (ErrorAccountNotFound) Confidential() bool {
	return false
}

type ErrorFailedRead struct {
	Message string
}

func (ErrorFailedRead) method() uint16 {
	return errorFailedRead
}

func (ErrorFailedRead) Confidential() bool {
	return false
}
