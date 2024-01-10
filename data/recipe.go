package data

type Cocktails struct {
	Drinks []DrinkRecipe `json:"drinks"`
}

type DrinkRecipe struct {
	Name         string `json:"strDrink"`
	Instructions string `json:"strInstructions"`
	Ingredient1  string `json:"strIngredient1"`
	Ingredient2  string `json:"strIngredient2"`
	Ingredient3  string `json:"strIngredient3"`
	Ingredient4  string `json:"strIngredient4"`
	Ingredient5  string `json:"strIngredient5"`
	Measurement1 string `json:"strMeasure1"`
	Measurement2 string `json:"strMeasure2"`
	Measurement3 string `json:"strMeasure3"`
	Measurement4 string `json:"strMeasure4"`
	Measurement5 string `json:"strMeasure5"`
}
