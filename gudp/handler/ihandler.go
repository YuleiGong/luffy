package handler

type IHandler interface{}

func NewHandler() IHandler {
	return &Handler{}
}
