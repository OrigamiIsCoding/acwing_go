package main

import "fmt"

func main() {
	var (
		n float64
	)
	const (
		INF = 1e4 + 10
	)
	fmt.Scanf("%f", &n)

	l, r := -INF, INF
	for r-l > 1e-8 {
		mid := (l + r) / 2
		if mid*mid*mid < n {
			l = mid
		} else {
			r = mid
		}
	}
	fmt.Printf("%.6f", l)
}
