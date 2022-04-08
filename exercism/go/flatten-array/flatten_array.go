package flatten

func Flatten(input interface{}) []interface{} {
	output := []interface{}{}

	switch v := input.(type) {
	case []interface{}:
		for _, el := range v {
			output = append(output, Flatten(el)...)
		}
	default:
		if v != nil {
			output = append(output, v)
		}
	}

	return output
}
