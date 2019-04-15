package factory



type Shapefactory struct {
	
}
func (this *Shapefactory) GetShape (name string)  Shape {
	if name=="Circle" {
		return &Circle{}
	}
	if name=="Rectangle" {
		return &Rectangle{}
	}
	if name=="Square" {
		return &Square{}
	}
	return &NullShape{}
}
