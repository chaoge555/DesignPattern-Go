package builder

type ChickenBurger struct {
}

func (this *ChickenBurger) Name() string {
	return "鸡肉汉堡"
}

func (this *ChickenBurger) Packing() Packing {
	return &Wrapper{}
}

func (this *ChickenBurger) Price() float64 {
	return 30.0
}
