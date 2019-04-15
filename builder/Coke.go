package builder

type Coke struct {
}

func (this *Coke) Name() string {
	return "可口可乐"
}

func (this *Coke) Packing() Packing {
	return &Bottle{}
}

func (this *Coke) Price() float64 {
	return 15.0

}
