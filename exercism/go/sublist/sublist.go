package sublist

import "reflect"

type Relation string

var (
	EQUAL     Relation = "equal"
	UNEQUAL   Relation = "unequal"
	SUBLIST   Relation = "sublist"
	SUPERLIST Relation = "superlist"
)

func Sublist(listOne, listTwo []int) Relation {
	if reflect.DeepEqual(listOne, listTwo) {
		return EQUAL
	} else if isSublist(listOne, listTwo) {
		return SUBLIST
	} else if isSublist(listTwo, listOne) {
		return SUPERLIST
	}

	return UNEQUAL
}

func isSublist(listOne, listTwo []int) bool {
	for i := 0; i <= len(listTwo)-len(listOne); i++ {
		if reflect.DeepEqual(listTwo[i:i+len(listOne)], listOne) {
			return true
		}
	}
	return false
}
