package main

import (
	"fmt"
)

var (
	A, B string
	a, b = make([]int, 0), make([]int, 0)
)

func add(a, b []int) []int {
	c := make([]int, 0)
	t := 0
	for i := 0; i < len(a) || i < len(b); i++ {
		if i < len(a) {
			t += a[i]
		}
		if i < len(b) {
			t += b[i]
		}
		c = append(c, t%10)
		t /= 10
	}
	if t != 0 {
		c = append(c, t)
	}
	return c
}

func main() {
	fmt.Scanf("%s\n%s", &A, &B)
	for i := len(A) - 1; i >= 0; i-- {
		a = append(a, int(A[i]-'0'))
	}
	for i := len(B) - 1; i >= 0; i-- {
		b = append(b, int(B[i]-'0'))
	}

	c := add(a, b)
	for i := len(c) - 1; i >= 0; i-- {
		fmt.Print(c[i])
	}
}
