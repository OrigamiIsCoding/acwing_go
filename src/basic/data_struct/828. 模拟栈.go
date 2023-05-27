package main

import "fmt"

func main() {
	const (
		N = 1e5 + 10
	)
	var (
		stk       [N]int
		top, m, x int
		op        string
	)

	init := func() {
		top = -1
	}

	push := func(x int) {
		top++
		stk[top] = x
	}

	pop := func() {
		top--
	}

	empty := func() bool {
		return top == -1
	}

	query := func() int {
		return stk[top]
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
			fmt.Printf("%d\n", query())
		}
	}
}
