package main

import (
	"fmt"
	"strings"
)

func main() {
	const End = "12345678x"
	var (
		s, c   string
		cnt    = make(map[string]int, 0)
		q      []string
		dx, dy = [4]int{1, -1, 0, 0}, [4]int{0, 0, 1, -1}
	)
	for i := 0; i < 9; {
		_, err := fmt.Scanf("%s", &c)
		if err == nil {
			s += c
			i++
		}
	}
	q = append(q, s)

	cnt[s] = 0
	for len(q) != 0 {
		u := q[0]
		q = q[1:]

		if u == End {
			break
		}

		idx := strings.IndexRune(u, 'x')
		x, y := idx/3, idx%3

		for i := 0; i < 4; i++ {
			xx, yy := x+dx[i], y+dy[i]
			if xx >= 0 && xx < 3 && yy >= 0 && yy < 3 {
				j := xx*3 + yy
				ns := []rune(u)
				ns[idx], ns[j] = ns[j], ns[idx]
				ne := string(ns)
				_, ok := cnt[ne]
				if ok {
					continue
				}
				cnt[ne] = cnt[u] + 1
				q = append(q, ne)
			}
		}
	}
	v, ok := cnt[End]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println(-1)
	}
}
