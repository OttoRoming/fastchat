package fcprotocol

type Message interface {
	method() uint16
	Confidential() bool
}

const (
	// Uptime related methods
	methodReqUptime uint16 = iota
	methodAckUptime

	// Account related methods
	methodReqSignUp
	methodReqLogIn
	methodAckSignedin
	methodErrUsernameInUse

	// Chat related methods
	methodReqSendChat
	methodAckChatSent

	// Chat history related methods
	methodReqGetHistory
	methodAckHistory

	// Generic
	methodErrAccountNotInUse
	methodErrFailedRead

	methodLimit
)

// Uptime related messages
type ReqUptime struct{}

func (ReqUptime) method() uint16 {
	return methodReqUptime
}

func (ReqUptime) Confidential() bool {
	return false
}

type AckUptime struct {
	Uptime string
}

func (AckUptime) method() uint16 {
	return methodAckUptime
}

func (AckUptime) Confidential() bool {
	return false
}

// Account related messages

type ReqSignUp struct {
	Username string
	Password string
}

func (ReqSignUp) method() uint16 {
	return methodReqSignUp
}

func (ReqSignUp) Confidential() bool {
	return true
}

type ReqLogin struct {
	Username string
	Password string
}

func (ReqLogin) method() uint16 {
	return methodReqLogIn
}

func (ReqLogin) Confidential() bool {
	return true
}

type AckSignedIn struct {
	Token string
}

func (AckSignedIn) method() uint16 {
	return methodAckSignedin
}

func (AckSignedIn) Confidential() bool {
	return false
}

type ErrUsernameInUse struct{}

func (ErrUsernameInUse) method() uint16 {
	return methodErrUsernameInUse
}

func (ErrUsernameInUse) Confidential() bool {
	return false
}

// Chat related messages
type ReqSendMessage struct {
	Token string

	ID      string
	To      string
	Content string
}

func (ReqSendMessage) method() uint16 {
	return methodReqSendChat
}

func (ReqSendMessage) Confidential() bool {
	return false
}

type AckMessageSent struct {
	ID string
}

func (AckMessageSent) method() uint16 {
	return methodAckChatSent
}

func (AckMessageSent) Confidential() bool {
	return false
}

// Chat history related emthods
type ReqGetHistory struct {
	Token string

	With string
}

func (ReqGetHistory) method() uint16 {
	return methodReqGetHistory
}

func (ReqGetHistory) Confidential() bool {
	return false
}

type AckHistory struct {
	Chats []struct {
		To      string
		Content string
	}
}

func (AckHistory) method() uint16 {
	return methodAckHistory
}

func (AckHistory) Confidential() bool {
	return false
}

// Generic
type ErrAccountNotInUse struct{}

func (ErrAccountNotInUse) method() uint16 {
	return methodErrAccountNotInUse
}

func (ErrAccountNotInUse) Confidential() bool {
	return false
}

type ErrFailedRead struct {
	Message string
}

func (ErrFailedRead) method() uint16 {
	return methodErrFailedRead
}

func (ErrFailedRead) Confidential() bool {
	return false
}
