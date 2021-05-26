package xkcoin

type XKCoin struct {

}

func (X XKCoin) Notice(exp string) error {
	panic("implement me")
}

func NewXKCoin() XKCoin {
	return XKCoin{}
}