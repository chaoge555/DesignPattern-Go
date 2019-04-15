package builder

type Item interface {
	Name() string;
	Packing() Packing;
	Price() float64;
}
