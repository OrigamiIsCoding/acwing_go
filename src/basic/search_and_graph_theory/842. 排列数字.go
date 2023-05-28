package main

import "fmt"

var (
	n int
	q = make([]int, 0)
)

func dfs(u, state int) {
	if u == n {
		for _, v := range q {
			fmt.Printf("%d ", v)
		}
		fmt.Println()
		return
	}

	for i := 1; i <= n; i++ {
		if ((state >> i) & 1) == 0 {
			q = append(q, i)
			dfs(u+1, state|1<<i)
			q = q[:len(q)-1]
		}
	}
}

func main() {
	fmt.Scanf("%d", &n)
	dfs(0, 0)
}
