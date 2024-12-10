package model

// STATES FOR HANDLING USER MESSAGE REPLIES

type States struct {
	CurrentState State
}

type State string

var (
	NilState State = "nil-state"
	WaitingForNextMessage State = "waiting-for-next-message"
)
