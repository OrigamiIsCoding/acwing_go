package main

import "fmt"

func main() {
	const N = 1e5 + 10
	var (
		s, p string
		ne   [N]int
		n, m int
	)
	fmt.Scanf("%d\n%s\n%d\n%s", &n, &p, &m, &s)
	s = " " + s
	p = " " + p
	for i, j := 2, 0; i <= n; i++ {
		for j != 0 && p[j+1] != p[i] {
			j = ne[j]
		}
		if p[j+1] == p[i] {
			j++
		}
		ne[i] = j
	}
	for i, j := 1, 0; i <= m; i++ {
		for j != 0 && p[j+1] != s[i] {
			j = ne[j]
		}
		if p[j+1] == s[i] {
			j++
		}
		if j == n {
			fmt.Printf("%d ", i-n)
			j = ne[j]
		}
	}
}
