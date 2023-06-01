package main

import "fmt"

func main() {
	const N = 110
	var (
		s          [N]int
		f          [2][N]int
		n, m, v, w int
	)

	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	fmt.Scanf("%d %d", &n, &m)
	for i := 1; i <= n; i++ { // 枚举组
		fmt.Scanf("%d", &s[i])
		for j := 0; j < s[i]; j++ { // 枚举组里面的物品
			fmt.Scanf("%d %d", &v, &w)
			for k := 0; k <= m; k++ { // 枚举体积
				f[i&1][k] = max(f[i&1][k], f[(i-1)&1][k])
				if k-v >= 0 {
					f[i&1][k] = max(f[i&1][k], f[(i-1)&1][k-v]+w)
				}
			}
		}
	}
	fmt.Println(f[n&1][m])
}
