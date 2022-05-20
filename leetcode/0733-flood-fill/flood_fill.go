package flood_fill

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if oldColor := image[sr][sc]; oldColor != newColor {
		fillRood(image, sr, sc, oldColor, newColor)
	}

	return image
}

func fillRood(image [][]int, sr int, sc int, oldColor int, newColor int) {
	if sr < 0 || sc < 0 || sr >= len(image) || sc >= len(image[0]) {
		return
	}

	if image[sr][sc] != oldColor {
		return
	}

	image[sr][sc] = newColor

	fillRood(image, sr+1, sc, oldColor, newColor)
	fillRood(image, sr-1, sc, oldColor, newColor)
	fillRood(image, sr, sc+1, oldColor, newColor)
	fillRood(image, sr, sc-1, oldColor, newColor)
}
