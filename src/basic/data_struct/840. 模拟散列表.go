package main

import (
	"fmt"
)

func main() {
	const (
		N = 100007
	)
	var (
		h, ne, e  [N]int
		op        string
		n, x, idx int
	)

	init := func(n int) {
		for i := 0; i < n; i++ {
			h[i] = -1
		}
	}

	hash := func(x int) int {
		return ((x % N) + N) % N
	}

	insert := func(x int) {
		e[idx] = x
		i := hash(x)
		ne[idx] = h[i]
		h[i] = idx
		idx++
	}

	query := func(x int) bool {
		i := hash(x)
		for j := h[i]; j != -1; j = ne[j] {
			if e[j] == x {
				return true
			}
		}
		return false
	}

	init(N)
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%s %d", &op, &x)
		if op == "I" {
			insert(x)
		} else {
			if query(x) {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		}
	}
}
