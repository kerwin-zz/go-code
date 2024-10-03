package main

/*
	给你一个大小为 m x n 的二维矩阵 grid 。你需要判断每一个格子 grid[i][j] 是否满足：
	· 如果它下面的格子存在，那么它需要等于它下面的格子，也就是	grid[i][j] == grid[i + 1][j] 。
	· 如果它右边的格子存在，那么它需要不等于它右边的格子，也就是	grid[i][j] != grid[i][j + 1] 。
	如果 所有 格子都满足以上条件，那么返回 true ，否则返回 false 。
*/

func satisfiesConditions(grid [][]int) bool {
	for i, row := range grid {
		for j, x := range row {
			if j > 0 && x == row[j-1] || i > 0 && x != grid[i-1][j] {
				return false
			}
		}
	}
	return true
}