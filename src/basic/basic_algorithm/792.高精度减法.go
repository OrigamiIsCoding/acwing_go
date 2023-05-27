package main

import "fmt"

func main() {
	var (
		A, B string
		a, b = make([]int, 0), make([]int, 0)
		c    []int
	)
	fmt.Scanf("%s\n%s", &A, &B)
	for i := len(A) - 1; i >= 0; i-- {
		a = append(a, int(A[i]-'0'))
	}
	for i := len(B) - 1; i >= 0; i-- {
		b = append(b, int(B[i]-'0'))
	}

	cmp := func(a, b []int) bool {
		if len(a) != len(b) {
			return len(a) > len(b)
		}
		for i := len(a) - 1; i >= 0; i-- {
			if a[i] != b[i] {
				return a[i] > b[i]
			}
		}
		return true
	}

	sub := func(a, b []int) []int {
		c := make([]int, 0)
		t := 0
		for i, v := range a {
			t += v
			if i < len(b) {
				t -= b[i]
			}
			c = append(c, (t+10)%10)
			if t < 0 {
				t = -1
			} else {
				t = 0
			}
		}
		for len(c) > 1 && c[len(c)-1] == 0 {
			c = c[:len(c)-1]
		}
		return c
	}

	if cmp(a, b) {
		c = sub(a, b)
	} else {
		c = sub(b, a)
		fmt.Print("-")
	}
	for i := len(c) - 1; i >= 0; i-- {
		fmt.Print(c[i])
	}
}
