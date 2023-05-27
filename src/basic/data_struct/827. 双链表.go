package main

import "fmt"

const (
	N = 1e5 + 10
)

var (
	head, tail, idx int
	ne, pe, e       [N]int
	n, k, x         int
	op              string
)

func init() {
	head = 0
	tail = 1
	ne[head] = tail
	pe[head] = -1
	ne[tail] = -1
	pe[tail] = head
	idx = 2
}

func pushLeft(x int) {
	pushLeftK(ne[head], x)
}

func pushRight(x int) {
	pushRightK(pe[tail], x)
}

func remove(k int) {
	pe[ne[k]] = pe[k]
	ne[pe[k]] = ne[k]
}

func pushLeftK(k, x int) {
	e[idx] = x
	ne[idx] = k
	pe[idx] = pe[k]
	ne[pe[k]] = idx
	pe[k] = idx
	idx++
}

func pushRightK(k, x int) {
	pushLeftK(ne[k], x)
}

func main() {
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%s", &op)
		if op == "L" {
			fmt.Scanf("%d", &x)
			pushLeft(x)
		} else if op == "R" {
			fmt.Scanf("%d", &x)
			pushRight(x)
		} else if op == "D" {
			fmt.Scanf("%d", &k)
			remove(k + 1)
		} else if op == "IL" {
			fmt.Scanf("%d %d", &k, &x)
			pushLeftK(k+1, x)
		} else {
			fmt.Scanf("%d %d", &k, &x)
			pushRightK(k+1, x)
		}
	}
	for i := ne[head]; i != tail; i = ne[i] {
		fmt.Printf("%d ", e[i])
	}
}
