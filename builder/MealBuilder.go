package builder

type MealBuilder struct {
}

func (this *MealBuilder) PrepareVegMeal() Meal {
	meal := Meal{}
	meal.AddItem(&VegBurger{})
	meal.AddItem(&Coke{})
	return meal
}

func (this *MealBuilder) PrepareNonVegMeal() Meal {
	meal := Meal{}
	meal.AddItem(&ChickenBurger{})
	meal.AddItem(&Pepsi{})
	return meal
}
