package un_pack

import "gudp/message"

type IUnPack interface {
	UnPack() message.Message
}

func NewUnPack() IUnPack {
	return &UnPack{}
}
