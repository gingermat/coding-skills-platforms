package lasagna

func PreparationTime(layers []string, time int) int {
	if time == 0 {
		time = 2
	}
	return len(layers) * time
}

func Quantities(layers []string) (int, float64) {
	var (
		noodle int
		sauce  float64
	)

	for _, el := range layers {
		switch el {
		case "noodles":
			noodle += 50
		case "sauce":
			sauce += .2
		}
	}

	return noodle, sauce
}

func AddSecretIngredient(friendList, myList []string) []string {
	return append(myList, friendList[len(friendList)-1])
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	result := make([]float64, len(quantities))

	for i, q := range quantities {
		result[i] = q / float64(2) * float64(portions)
	}

	return result
}
