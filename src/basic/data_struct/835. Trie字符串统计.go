package main

import "fmt"

func main() {
	const N = 1e5 + 10
	var (
		son   [N][26]int
		cnt   [N * 26]int
		idx   = 1
		n     int
		s, op string
	)

	insert := func(s string) {
		p := 0
		for i := 0; i < len(s); i++ {
			c := int(s[i] - 'a')
			if son[p][c] == 0 {
				son[p][c] = idx
				idx++
			}
			p = son[p][c]
		}
		cnt[p]++
	}

	query := func(s string) int {
		p := 0
		for i := 0; i < len(s); i++ {
			c := int(s[i] - 'a')
			if son[p][c] == 0 {
				return 0
			}
			p = son[p][c]
		}
		return cnt[p]
	}

	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%s %s", &op, &s)
		if op == "I" {
			insert(s)
		} else {
			fmt.Println(query(s))
		}
	}
}
