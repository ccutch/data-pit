package main

import (
	"fmt"

	"github.com/ccutch/data-pit/users"
)

func main() {
	for i := 0; i < 1; i++ {
		res := users.CreateTestUser()
		fmt.Println(res)
	}
}
