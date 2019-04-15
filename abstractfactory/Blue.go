package abstractfactory

import "fmt"

type Blue struct {
}

func (this *Blue) Fill() {
	fmt.Println("用蓝色涂满")
}
