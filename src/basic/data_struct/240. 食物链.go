package main

const N = 5e4 + 10

var (
	p, d           [N]int
	n, k, op, x, y int
)

func initP(n int) {
	for i := 0; i <= n; i++ {
		p[i] = i
		d[i] = 0
	}
}

func find(x int) int {
	if x != p[x] {
		t := find(p[x])
		d[x] += d[p[x]]
		p[x] = t
	}
	return p[x]
}

func main() {

}
