package main

import "fmt"

func main() {
	const N = 1e5 + 10
	var (
		q            [N]int
		op           string
		m, x, hh, tt int
	)
	init := func() {
		tt = -1
		hh = 0
	}
	push := func(x int) {
		tt++
		q[tt] = x
	}
	pop := func() {
		hh++
	}
	empty := func() bool {
		return tt < hh
	}
	query := func() int {
		return q[hh]
	}
	init()
	fmt.Scanf("%d", &m)
	for i := 0; i < m; i++ {
		fmt.Scanf("%s", &op)
		if op == "push" {
			fmt.Scanf("%d", &x)
			push(x)
		} else if op == "pop" {
			pop()
		} else if op == "empty" {
			if empty() {
				fmt.Println("YES")
			} else {
				fmt.Println("NO")
			}
		} else {
			fmt.Println(query())
		}
	}
}
