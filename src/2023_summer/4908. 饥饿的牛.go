package main

import (
	"fmt"
)

func main() {
	var (
		n                              int
		t, d, b, timestamp, count, ans int64
	)

	min := func(a, b int64) int64 {
		if a < b {
			return a
		} else {
			return b
		}
	}

	fmt.Scan(&n, &t)
	for i := 0; i < n; i++ {
		fmt.Scan(&d, &b)

		// d 天收到 b
		// (timestamp - t) =
		c := min(count, d-timestamp)
		ans += c
		count += b - c
		timestamp = d
	}
	fmt.Println(ans + min(count, t-timestamp+1))
}
