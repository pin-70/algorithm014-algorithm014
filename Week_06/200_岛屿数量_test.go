package Week_06

// 深度优先算法
func numIslands(grid [][]byte) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' { //这里注意要用'1'
				res++
				helper(grid, i, j)
			}
		}
	}
	return res
}

func helper(grid [][]byte, i, j int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	helper(grid, i-1, j)
	helper(grid, i+1, j)
	helper(grid, i, j-1)
	helper(grid, i, j+1)
}
