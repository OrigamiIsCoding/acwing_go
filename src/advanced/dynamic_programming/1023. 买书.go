package main

import "fmt"

func main() {
	const N = 1010
	var (
		f [N]int
		w = [4]int{10, 20, 50, 100}
		n int
	)
	fmt.Scan(&n)
	f[0] = 1
	for _, v := range w {
		for i := v; i <= n; i++ {
			f[i] += f[i-v]
		}
	}
	fmt.Println(f[n])
}
