package main

import (
	"fmt"
	"sort"
)

type Point struct {
	x, y int
}

type Points []Point

const N = 5010

var (
	p       = make(Points, 0)
	f       [N]int
	n, x, y int
)

func (a Points) Len() int      { return len(a) }
func (a Points) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a Points) Less(i, j int) bool {
	if a[i].x != a[j].x {
		return a[i].x < a[j].x
	}
	return a[i].y < a[j].y
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x, &y)
		p = append(p, Point{
			x: x,
			y: y,
		})
	}

	sort.Sort(&p)
	len := 0
	for i := 0; i < n; i++ {
		l, r := 0, len
		for l < r {
			mid := (l + r + 1) / 2
			if f[mid] > p[i].y {
				r = mid - 1
			} else {
				l = mid
			}
		}
		len = max(len, l+1)
		f[l+1] = p[i].y
	}
	fmt.Println(len)
}
