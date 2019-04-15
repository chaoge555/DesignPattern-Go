package abstractfactory

type ColorFactory struct {

}
func (this * ColorFactory)GetColorFactory(color string) Color  {
	if color=="Red" {
		return &Red{}
	}
	if color=="Green" {
		return &Green{}
	}
	if color =="Blue" {
		return &Blue{}
	}
	return nil
}
func (this * ColorFactory)GetShapeFactory(shape string) Shape  {
	return nil
}
