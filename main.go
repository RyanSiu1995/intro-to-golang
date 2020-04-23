package main

import (
	"fmt"

	"github.com/RyanSiu1995/intro-to-go/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println("Error encountered... Use DEBUG flag to see more error message...")
	}
}
