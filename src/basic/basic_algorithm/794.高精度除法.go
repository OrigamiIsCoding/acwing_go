package main

import "fmt"

func main() {
	var (
		A string
		a = make([]int, 0)
		b int
	)
	fmt.Scanf("%s\n%d", &A, &b)
	for i := len(A) - 1; i >= 0; i-- {
		a = append(a, int(A[i]-'0'))
	}

	reverse := func(a []int) []int {
		for i := 0; i < len(a)/2; i++ {
			j := len(a) - i - 1
			a[i], a[j] = a[j], a[i]
		}
		return a
	}

	div := func(a []int, b int) ([]int, int) {
		c := make([]int, 0)
		t := 0

		for _, v := range reverse(a) {
			t *= 10
			t += v
			c = append(c, t/b)
			t %= b
		}
		c = reverse(c)
		for len(c) > 1 && c[len(c)-1] == 0 {
			c = c[:len(c)-1]
		}
		return c, t
	}

	c, p := div(a, b)
	for i := len(c) - 1; i >= 0; i-- {
		fmt.Print(c[i])
	}
	fmt.Printf("\n%d", p)
}
