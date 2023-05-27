package main

import "fmt"

func main() {
	var (
		n, x int
	)
	fmt.Scanf("%d", &n)

	lowbit := func(x int) int {
		return x & -x
	}

	count := func(x int) int {
		cnt := 0
		for ; x != 0; x -= lowbit(x) {
			cnt++
		}
		return cnt
	}

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &x)
		fmt.Printf("%d ", count(x))
	}
}
