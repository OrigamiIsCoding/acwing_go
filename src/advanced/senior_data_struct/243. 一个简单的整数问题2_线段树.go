package main

import "fmt"

const N = 100010

type Node struct {
	l, r     int
	val, add int64
}

func (n *Node) len() int {
	return n.r - n.l + 1
}

func (n *Node) mid() int {
	return (n.l + n.r) / 2
}

func (n *Node) recv(add int64) {
	n.val += int64(n.len()) * add
	n.add += add
}

var (
	s          [N << 2]Node
	a          [N]int
	n, m, l, r int
	d          int64
	op         string
)

func pushUp(u int) {
	s[u].val = s[u<<1].val + s[u<<1|1].val
}

func pushDown(u int) {
	if s[u].add != 0 {
		s[u<<1].recv(s[u].add)
		s[u<<1|1].recv(s[u].add)
		s[u].add = 0
	}
}

func build(u, l, r int) {
	s[u] = Node{
		l: l,
		r: r,
	}

	if l >= r {
		s[u].val = int64(a[l])
	} else {
		mid := s[u].mid()
		build(u<<1, l, mid)
		build(u<<1|1, mid+1, r)
		pushUp(u)
	}
}

func modify(u, l, r int, d int64) {
	if s[u].l >= l && s[u].r <= r {
		s[u].recv(d)
		return
	}
	mid := s[u].mid()
	pushDown(u)
	if mid >= l {
		modify(u<<1, l, r, d)
	}
	if r > mid {
		modify(u<<1|1, l, r, d)
	}
	pushUp(u)
}

func query(u, l, r int) (ans int64) {
	if s[u].l >= l && s[u].r <= r {
		return s[u].val
	}
	mid := s[u].mid()
	pushDown(u)
	if mid >= l {
		ans += query(u<<1, l, r)
	}
	if r > mid {
		ans += query(u<<1|1, l, r)
	}
	return
}

func main() {
	fmt.Scan(&n, &m)
	for i := 1; i <= n; i++ {
		fmt.Scan(&a[i])
	}
	build(1, 1, n)
	for i := 0; i < m; i++ {
		fmt.Scan(&op)
		if op == "C" {
			fmt.Scan(&l, &r, &d)
			modify(1, l, r, d)
		} else {
			fmt.Scan(&l, &r)
			fmt.Println(query(1, l, r))
		}
	}
}
