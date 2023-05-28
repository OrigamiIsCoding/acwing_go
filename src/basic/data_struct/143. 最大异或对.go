package main

import (
	"fmt"
)

func main() {
	const N = 1e5 + 10
	var (
		son       [N * 31][2]int
		idx       = 1
		n, a, ans int
	)

	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	query := func(x int) int {
		ans, p := 0, 0
		for i := 31; i >= 0; i-- {
			v := (x >> i & 1)
			nv := 0
			if v == 0 {
				nv = 1
			}
			if son[p][nv] != 0 {
				ans |= (1 << i)
				p = son[p][nv]
			} else {
				p = son[p][v]
			}
		}
		return ans
	}

	insert := func(x int) {
		p := 0
		for i := 31; i >= 0; i-- {
			v := (x >> i & 1)
			if son[p][v] == 0 {
				son[p][v] = idx
				idx++
			}
			p = son[p][v]
		}
	}

	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a)
		ans = max(query(a), ans)
		insert(a)
	}
	fmt.Println(ans)
}
