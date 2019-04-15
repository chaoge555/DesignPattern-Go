package main

import "designPatter/builder"

func main() {
	mealBuilder := builder.MealBuilder{}
	vegMeal := mealBuilder.PrepareNonVegMeal()
	vegMeal.ShowItems()

	nonVegMeal := mealBuilder.PrepareVegMeal()
	nonVegMeal.ShowItems()
}
