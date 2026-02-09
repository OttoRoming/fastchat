package fcprotocol

type Message interface {
	method() uint16
}

const (
	// Account Related methods
	methodReqSignUp uint16 = iota
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

	methodLimit
)

// Account related messages

type ReqSignUp struct {
	Username string
	Password string
}

func (ReqSignUp) method() uint16 {
	return methodReqSignUp
}

type ReqLogin struct {
	Username string
	Password string
}

func (ReqLogin) method() uint16 {
	return methodReqLogIn
}

type AckSignedIn struct {
	Token string
}

func (AckSignedIn) method() uint16 {
	return methodAckSignedin
}

type ErrUsernameInUse struct{}

func (ErrUsernameInUse) method() uint16 {
	return methodErrUsernameInUse
}

// Chat related messages
type ReqSendMessage struct {
	ID      string
	To      string
	Content string
}

func (ReqSendMessage) method() uint16 {
	return methodReqSendChat
}

type AckMessageSent struct {
	id string
}

func (AckMessageSent) method() uint16 {
	return methodAckChatSent
}

// Chat history related emthods
type ReqGetHistory struct {
	with string
}

func (ReqGetHistory) method() uint16 {
	return methodReqGetHistory
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

// Generic
type ErrAccountNotInUse struct{}

func (ErrAccountNotInUse) method() uint16 {
	return methodErrAccountNotInUse
}
