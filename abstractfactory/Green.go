package abstractfactory

import "fmt"

type Green struct {

}
func (this *Green)Fill()  {
	fmt.Println("用绿色涂满")
}
