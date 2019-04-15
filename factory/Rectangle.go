package factory

import "fmt"

type Rectangle struct {
}
func (this *Rectangle) Draw()  {
	fmt.Println("画矩形")
}
func (this *Rectangle)IsNil() bool {
	return false
}
