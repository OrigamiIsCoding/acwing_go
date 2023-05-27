package main

import "fmt"

func main() {
	const (
		N = 1e5 + 10
	)
	var (
		a, b [N]int
		n, m int
	)
	fmt.Scanf("%d %d", &n, &m)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}
	for i := 0; i < m; i++ {
		fmt.Scanf("%d", &b[i])
	}

	i := 0
	for j := 0; i < n && j < m; j++ {
		if a[i] == b[j] {
			i++
		}
	}
	if i == n {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
