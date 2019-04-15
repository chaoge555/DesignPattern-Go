package builder

import "fmt"

type Meal struct {
	Items []Item
}

func (this *Meal) AddItem(item Item) {
	this.Items = append(this.Items, item)
}

func (this *Meal) GetCost() float64 {
	cost := 0.0

	/*	for index,val:=range this.Items  {
			cost=val.Price()+cost
		}*/

	for i := 0; i < len(this.Items); i++ {
		cost = cost + this.Items[i].Price()
	}
	return cost
}

func (this *Meal) ShowItems() {
	for i := 0; i < len(this.Items); i++ {
		fmt.Println("商品", this.Items[i].Name())
		fmt.Println("包装", this.Items[i].Packing().Pack())
		fmt.Println("价钱", this.Items[i].Price(), "元")
	}
	fmt.Println("一共：", this.GetCost())
	fmt.Println("")

}
