package abstractfactory

import "fmt"

type Circle struct {

}
func (this *Circle) Draw()  {
	fmt.Println("画圆形")
}
