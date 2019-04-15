package builder

type Pepsi struct {
}

func (this *Pepsi) Name() string {
	return "百事可乐"
}

func (this *Pepsi) Packing() Packing {
	return &Bottle{}
}

func (this *Pepsi) Price() float64 {
	return 13.0
}
