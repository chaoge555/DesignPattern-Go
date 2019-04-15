package abstractfactory

type AbstractFactory interface {
	 GetShape() Shape
	 GetColor() Color
}
