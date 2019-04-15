package factory

import "fmt"

type NullShape struct {
}
func (this *NullShape) Draw()  {
	fmt.Println("没有你想要画得形状")
}
func (this *NullShape)IsNil() bool {
	return 	true
}

