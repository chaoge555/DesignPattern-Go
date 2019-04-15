package abstractfactory

type ShapeFactory struct {

}
func (this * ShapeFactory)GetColorFactory(color string) Color  {
	return nil
}
func (this * ShapeFactory)GetShapeFactory(shape string) Shape  {
	if shape=="Circle" {
		return &Circle{}
	}
	if shape =="Rectangle" {
		return &Rectangle{}
	}
	if shape=="Square" {
		return &Square{}
	}
	return nil
}