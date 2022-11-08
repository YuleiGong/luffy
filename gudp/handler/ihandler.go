package handler

import "gudp/message"

type IHandler interface {
	Do(*message.Message)
}

func NewHandler() IHandler {
	return &Handler{}
}
