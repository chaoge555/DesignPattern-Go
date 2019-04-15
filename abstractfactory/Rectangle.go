package abstractfactory

import "fmt"

type Rectangle struct {

}
func (this *Rectangle)Draw()  {
	fmt.Println("画矩形")
}
