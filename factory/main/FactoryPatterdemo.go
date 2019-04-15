package main

import (
	"designPatter/factory"
	"fmt"
)

func main()  {
	 shapeFactory := factory.Shapefactory{}
	 shape1 :=  shapeFactory.GetShape("Circle")
	shape2 :=  shapeFactory.GetShape("Rectangle")
	shape3 :=  shapeFactory.GetShape("Square")
	shape4 :=  shapeFactory.GetShape("Diamond")
	shape5 :=  shapeFactory.GetShape("")
	shape1.Draw()
	shape2.Draw()
	shape3.Draw()
	shape4.Draw()
	shape5.Draw()
	fmt.Println("shape1空状态：",shape1.IsNil())
	fmt.Println("shape2空状态：",shape2.IsNil())
	fmt.Println("shape3空状态：",shape3.IsNil())
	fmt.Println("shape4空状态：",shape4.IsNil())
	fmt.Println("shape5空状态：",shape5.IsNil())
}