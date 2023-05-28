package main

import "fmt"

func main() {
	const N = 1e5 + 10
	var (
		stk       [N]int
		top, n, t int
	)
	stk[0] = -1
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &t)
		for top != -1 && stk[top] >= t {
			top--
		}
		fmt.Printf("%d ", stk[top])
		top++
		stk[top] = t
	}
}
