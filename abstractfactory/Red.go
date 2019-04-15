package abstractfactory

import "fmt"

type Red struct {

}

func (this *Red)Fill()  {
	fmt.Println("用红色涂满")
}
