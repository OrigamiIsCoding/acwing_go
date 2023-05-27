package main

import "fmt"

func main() {
	const (
		N = 1e5 + 10
	)
	var (
		head         = -1
		e, ne        [N]int
		op           string
		n, k, x, idx int
	)
	fmt.Scanf("%d", &n)

	insert := func(x int) {
		e[idx] = x
		ne[idx] = head
		head = idx
		idx++
	}

	insertK := func(k, x int) {
		e[idx] = x
		ne[idx] = ne[k]
		ne[k] = idx
		idx++
	}

	remove := func(k int) {
		if k == -1 {
			head = ne[head]
		} else {
			ne[k] = ne[ne[k]]
		}
	}

	for i := 0; i < n; i++ {
		fmt.Scanf("%s", &op)
		if op == "H" {
			fmt.Scanf("%d", &x)
			insert(x)
		} else if op == "I" {
			fmt.Scanf("%d %d", &k, &x)
			insertK(k-1, x)
		} else {
			fmt.Scanf("%d", &k)
			remove(k - 1)
		}
	}
	for i := head; i != -1; i = ne[i] {
		fmt.Printf("%d ", e[i])
	}
}
