package series

func All(length int, data string) []string {
	if length > len(data) {
		return nil
	}

	var result []string

	for i := 0; i < len(data); i++ {
		if end := i + length; end <= len(data) {
			chunk := data[i:end]
			result = append(result, chunk)
		}
	}
	return result
}

func UnsafeFirst(length int, data string) string {
	if length <= len(data) {
		return data[0:length]
	}

	return ""
}
