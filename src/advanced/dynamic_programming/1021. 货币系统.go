package main

import "fmt"

func main() {
	const N = 3010
	var (
		f       [N]int
		n, m, v int
	)
	f[0] = 1
	fmt.Scan(&n, &m)
	for i := 0; i < n; i++ {
		fmt.Scan(&v)
		for j := v; j <= m; j++ {
			f[j] += f[j-v]
		}
	}
	fmt.Println(f[m])
}
