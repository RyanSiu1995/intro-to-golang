package main

import (
	"fmt"

	"github.com/RyanSiu1995/intro-to-golang/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println("Error encountered... Use DEBUG flag to see more error message...")
	}
}
