package main

import "fmt"

const (
	N = 200010
	M = N << 2
)

type Node struct {
	l, r int
	max  int
}

func (n Node) mid() int {
	return (n.l + n.r) / 2
}

var (
	s          [M]Node
	m, p, l, t int
	op         string
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func build(u, l, r int) {
	s[u] = Node{
		l: l,
		r: r,
	}
	if l < r {
		mid := s[u].mid()
		build(u<<1, l, mid)
		build(u<<1|1, mid+1, r)
	}
}

func pushUp(u int) {
	s[u].max = max(s[u<<1].max, s[u<<1|1].max)
}

func query(u, l, r int) int {
	if s[u].l >= l && s[u].r <= r {
		return s[u].max
	}
	mid := s[u].mid()
	ans := 0
	if l <= mid {
		ans = query(u<<1, l, r)
	}
	if r > mid {
		ans = max(ans, query(u<<1|1, l, r))
	}
	return ans
}

func modify(u, i, x int) {
	if s[u].l >= i && s[u].r <= i {
		s[u].max = x
		return
	}
	mid := s[u].mid()
	if i <= mid {
		modify(u<<1, i, x)
	} else {
		modify(u<<1|1, i, x)
	}
	pushUp(u)
}

func main() {
	fmt.Scan(&m, &p)
	build(1, 1, m)
	n, a := 0, 0
	for i := 0; i < m; i++ {
		fmt.Scan(&op)
		if op == "Q" {
			fmt.Scan(&l)
			a = query(1, n-l+1, n)
			fmt.Println(a)
		} else {
			fmt.Scan(&t)
			n++
			modify(1, n, (t+a)%p)
		}
	}
}
