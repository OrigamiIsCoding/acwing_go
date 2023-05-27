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

	mul := func(a []int, b int) []int {
		if b == 0 || (len(a) == 0 && a[0] == 0) {
			return []int{0}
		}
		c := make([]int, 0)
		t := 0
		for _, v := range a {
			t += v * b
			c = append(c, t%10)
			t /= 10
		}
		for t > 0 {
			c = append(c, t%10)
			t /= 10
		}
		return c
	}

	c := mul(a, b)
	for i := len(c) - 1; i >= 0; i-- {
		fmt.Print(c[i])
	}
}
