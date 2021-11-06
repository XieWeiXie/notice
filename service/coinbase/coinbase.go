package coinbase

type CoinBase struct {

}

func (c CoinBase) Notice(exp string) error {
	panic("implement me")
}

func NewCoinBase() CoinBase{
	return CoinBase{}
}
