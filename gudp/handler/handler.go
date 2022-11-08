package handler

import (
	"fmt"
	"gudp/message"
)

type Handler struct{}

func (h *Handler) Do(msg *message.Message) {

	fmt.Printf("%s", msg.GetMessage())
	fmt.Printf("%s", msg.GetClient)
}
