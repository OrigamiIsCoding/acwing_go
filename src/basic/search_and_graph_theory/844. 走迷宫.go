package main

import "fmt"

func main() {
	const (
		N = 110
	)
	var (
		a, dist      [N][N]int
		q            [N * N][2]int
		dx, dy       = [4]int{1, -1, 0, 0}, [4]int{0, 0, 1, -1}
		n, m, hh, tt int
	)
	fmt.Scanf("%d %d", &n, &m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scanf("%d", &a[i][j])
		}
	}

	dist[0][0] = 1
	q[tt] = [2]int{0, 0}

	for tt >= hh {
		u := q[hh]
		hh++
		x, y := u[0], u[1]
		for i := 0; i < 4; i++ {
			xx, yy := x+dx[i], y+dy[i]
			if xx >= 0 && xx < n && yy >= 0 && yy < m && a[xx][yy] == 0 && dist[xx][yy] == 0 {
				dist[xx][yy] = dist[x][y] + 1
				tt++
				q[tt] = [2]int{xx, yy}
			}
		}
	}
	fmt.Println(dist[n-1][m-1] - 1)
}
