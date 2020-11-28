package pointer

import "fmt"

func Pointer() {
	i, j := 1, 2

	p := &i
	q := &j

	*p = 10
	*q = 20

	fmt.Println(i, j)
}
