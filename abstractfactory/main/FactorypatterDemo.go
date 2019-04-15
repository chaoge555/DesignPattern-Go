package main

import "designPatter/abstractfactory"

func main() {
	colorFactory := abstractfactory.ColorFactory{}
	red := colorFactory.GetColorFactory("Red")
	red.Fill()

	green :=colorFactory.GetColorFactory("Green")
	green.Fill()

	blue :=colorFactory.GetColorFactory("Blue")
	blue.Fill()

	shapeFactory := abstractfactory.ShapeFactory{}
	circle := shapeFactory.GetShapeFactory("Circle")
	circle.Draw()

	rectangle :=shapeFactory.GetShapeFactory("Rectangle")
	rectangle.Draw()

	square :=shapeFactory.GetShapeFactory("Square")
	square.Draw()
}