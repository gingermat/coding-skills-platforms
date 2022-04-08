package allergies

var allAllergens = []string{"eggs", "peanuts", "shellfish", "strawberries", "tomatoes", "chocolate", "pollen", "cats"}

func Allergies(score uint) (allergens []string) {
	for _, allergen := range allAllergens {
		if ok := AllergicTo(score, allergen); ok {
			allergens = append(allergens, allergen)
		}
	}

	return allergens
}

func AllergicTo(score uint, name string) bool {
	for i, allergen := range allAllergens {
		mask := uint(1 << i)
		if allergen == name && score&mask > 0 {
			return true
		}
	}

	return false
}
