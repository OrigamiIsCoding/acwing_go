package main

import "fmt"

const N = 1e5 + 10

var (
	h             [N]int
	n, m, t, size int
)

func up(u int) {
	for t := u / 2; t != 0 && h[t] < h[u]; t = u / 2 {
		h[t], h[u] = h[u], h[t]
		u /= 2
	}
}

func down(u int) {
	t := u
	if i := u * 2; i <= size && h[t] < h[i] {
		t = i
	}
	if i := u*2 + 1; i <= size && h[t] < h[i] {
		t = i
	}
	if t != u {
		h[t], h[u] = h[u], h[t]
		down(t)
	}
}

func push(x int) {
	size++
	h[size] = x
	up(size)
}

func top() int {
	return h[1]
}

func updateTop(x int) {
	h[1] = x
	down(1)
}

func pop() {
	h[1], h[size] = h[size], h[1]
	size--
	down(1)
}

func main() {
	fmt.Scanf("%d %d", &n, &m)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &t)
		if size < m {
			push(t)
		} else if t < top() {
			updateTop(t)
		}
	}
	for i := 0; i < m; i++ {
		pop()
	}
	for i := 1; i <= m; i++ {
		fmt.Printf("%d ", h[i])
	}
}
