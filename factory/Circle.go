package factory

import "fmt"


type Circle struct {
}
func (this *Circle)Draw()  {
	fmt.Println("画圆形")

}
func (this *Circle)IsNil() bool {
	return false
}