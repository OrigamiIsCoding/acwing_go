package main

import (
	"fmt"
)

const N = 1e5 + 10

var (
	h, hp, ph          [N]int
	n, k, x, size, idx int
	op                 string
)

func swap(i, j int) {
	ph[hp[i]], ph[hp[j]] = ph[hp[j]], ph[hp[i]]
	hp[i], hp[j] = hp[j], hp[i]
	h[i], h[j] = h[j], h[i]
}

func up(u int) {
	for t := u / 2; t != 0 && h[t] > h[u]; t = u / 2 {
		swap(t, u)
		u = t
	}
}

func down(u int) {
	for {
		t := u
		if i := u * 2; i <= size && h[t] > h[i] {
			t = i
		}
		if i := u*2 + 1; i <= size && h[t] > h[i] {
			t = i
		}
		if t != u {
			swap(t, u)
			u = t
		} else {
			break
		}
	}
}

func push(x int) {
	size++
	idx++
	h[size] = x
	ph[idx] = size
	hp[size] = idx
	up(size)
}

func top() int {
	return h[1]
}

func pop() {
	swap(1, size)
	size--
	down(1)
}

func remove(k int) {
	i := ph[k]
	swap(i, size)
	size--
	up(i)
	down(i)
}

func update(k int, x int) {
	i := ph[k]
	h[i] = x
	up(i)
	down(i)
}

func main() {
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%s", &op)
		if op == "I" {
			fmt.Scanf("%d", &x)
			push(x)
		} else if op == "PM" {
			fmt.Println(top())
		} else if op == "DM" {
			pop()
		} else if op == "D" {
			fmt.Scanf("%d", &k)
			remove(k)
		} else {
			fmt.Scanf("%d %d", &k, &x)
			update(k, x)
		}
	}
}
