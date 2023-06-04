package main

import "fmt"

func main() {
	const N = 10010
	var (
		f       [N]int
		n, m, a int
	)
	fmt.Scan(&n, &m)
	f[0] = 1
	for t := 0; t < n; t++ {
		fmt.Scan(&a)
		for i := m; i >= a; i-- {
			f[i] += f[i-a]
		}
	}
	fmt.Println(f[m])
}
