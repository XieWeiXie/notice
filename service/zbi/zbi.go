package zbi

type Zbi struct {

}

func (z Zbi) Notice(exp string) error {
	panic("implement me")
}

func NewZbi() Zbi {
	return Zbi{}
}