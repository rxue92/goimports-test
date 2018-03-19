package functions

import "fmt"

func Loop(n int) {
	for i := range make([]interface{}, n) {
		fmt.Printf("loop:%d\n", i)
	}
	fmt.Println("looping finished")
}

func dirtyFunc() {
	fmt.Println("should be formatted by goimports")
}
