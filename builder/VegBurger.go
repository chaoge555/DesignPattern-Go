package builder

type VegBurger struct {
}

func (this *VegBurger) Name() string {
	return "蔬菜汉堡"
}

func (this *VegBurger) Packing() Packing {
	return &Wrapper{}
}

func (this *VegBurger) Price() float64 {
	return 25.0
}
