func floodFill(image [][]int, sr int, sc int, color int) [][]int {
    orgColor := image[sr][sc]
	if orgColor == color {
		return image
	}

	rows := len(image)
	cols := len(image[0])

	var dfs func(r, c int)
	dfs = func(r, c int) {
		if !(0 <= r && r < rows) || !(0 <= c && c < cols) || image[r][c] != orgColor {
			return
		}

		image[r][c] = color
		dfs(r+1,c)
		dfs(r-1,c)
		dfs(r,c+1)
		dfs(r,c-1)
	}

	dfs(sr, sc)
	return image
}
