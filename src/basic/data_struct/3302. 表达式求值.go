package main

import (
	"fmt"
	"unicode"
)

func main() {
	const N = 1e5 + 10
	var (
		expr       string
		numStk     [N]int
		opStk      [N]rune
		nTop, oTop int
	)
	orderMap := map[rune]int{
		'+': 1,
		'-': 1,
		'*': 2,
		'/': 2,
	}

	eval := func() {
		op, a, b := opStk[oTop], numStk[nTop], numStk[nTop-1]
		oTop--
		nTop--
		if op == '+' {
			b += a
		} else if op == '-' {
			b -= a
		} else if op == '*' {
			b *= a
		} else {
			b /= a
		}
		numStk[nTop] = b
	}

	nTop, oTop = -1, -1
	fmt.Scanf("%s", &expr)
	for i := 0; i < len(expr); i++ {
		r := rune(expr[i])
		if unicode.IsNumber(r) {
			num := 0
			for ; i < len(expr) && unicode.IsNumber(rune(expr[i])); i++ {
				num = num*10 + int(expr[i]-'0')
			}
			nTop++
			numStk[nTop] = num
			i--
		} else {
			if r == '(' {
				oTop++
				opStk[oTop] = r
			} else if r == ')' {
				for opStk[oTop] != '(' {
					eval()
				}
				oTop--
			} else {
				for oTop >= 0 && orderMap[r] <= orderMap[opStk[oTop]] {
					eval()
				}
				oTop++
				opStk[oTop] = r
			}
		}
	}
	for oTop != -1 {
		eval()
	}
	fmt.Println(numStk[nTop])
}
