package main

import "fmt"

func main() {
	const (
		N = 1e5 + 10
	)
	var (
		a, b    [N]int
		n, m, x int
	)
	fmt.Scanf("%d %d %d", &n, &m, &x)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}
	for i := 0; i < m; i++ {
		fmt.Scanf("%d", &b[i])
	}

	for i, j := 0, m-1; i < n && j >= 0; {
		v := a[i] + b[j]
		if v > x {
			j--
		} else if v < x {
			i++
		} else {
			fmt.Println(i, j)
			break
		}
	}
}
