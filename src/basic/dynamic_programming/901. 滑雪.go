package main

import "fmt"

const N = 310

var (
	g, w   [N][N]int
	r, c   int
	dx, dy = [4]int{1, -1, 0, 0}, [4]int{0, 0, 1, -1}
)

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func f(x, y int) int {
	if g[x][y] == 0 {
		g[x][y] = 1
		for i := 0; i < 4; i++ {
			xx, yy := x+dx[i], y+dy[i]
			if xx >= 0 && xx < r && yy >= 0 && yy < c && w[x][y] > w[xx][yy] {
				g[x][y] = max(g[x][y], f(xx, yy)+1)
			}
		}
	}
	return g[x][y]
}

func main() {
	fmt.Scan(&r, &c)
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			fmt.Scan(&w[i][j])
		}
	}
	ans := 0
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			ans = max(ans, f(i, j))
		}
	}
	fmt.Println(ans)
}
