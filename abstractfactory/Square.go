package abstractfactory

import "fmt"

type Square struct {

}
func (this *Square)Draw()  {
	fmt.Println("画正方形")
}
