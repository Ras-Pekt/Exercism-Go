package lasagnamaster

import "strings"

func PreparationTime(layers []string, avgPrepTime int) int {
	if avgPrepTime == 0 {
		avgPrepTime = 2
	}
	return len(layers) * avgPrepTime
}

func Quantities(layers []string) (int, float64) {
	sauseCount := 0
	noodleCount := 0

	for _, item := range layers {
		if strings.ToLower(item) == "sauce" {
			sauseCount += 1
		}
		if strings.ToLower(item) == "noodles" {
			noodleCount += 1
		}
	}
	return noodleCount * 50, float64(sauseCount) * 0.2
}

func AddSecretIngredient(friendsList, myList []string) {
	lastIngredient := friendsList[len(friendsList)-1]
	idx := len(myList) - 1
	myList[idx] = lastIngredient
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	newRecipe := make([]float64, len(quantities))
	scaleFactor := float64(portions) / 2.0
	for idx, item := range quantities {
		newRecipe[idx] = item * scaleFactor
	}
	return newRecipe
}
