package un_pack

type IUnPack interface{}

func NewUnPack() IUnPack {
	return &UnPack{}
}
