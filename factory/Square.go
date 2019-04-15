package factory

import "fmt"

type Square struct {
}
func(this *Square)Draw() {
	fmt.Println("画正方形")
}
func (this *Square)IsNil() bool {
	return false
}