package internal

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

// Define a custom type called out
// It will have all the functions from []string
type out []string

// This is the polymorphism implementation in Go
// Similar to typeclass in Haskell, trait in Rust or Inheritance in Java/Python
// fmt.Println() accepts a type called interface{}
// Inside fmt.Println, it will convert the interface{} to a type called Stringer
// In Stringer, it is supposed that the String() method is implemented
func (s out) String() string {
	var str string
	for _, i := range s {
		str += fmt.Sprintf("%s\n", i)
	}
	return str
}

func Run(argv ...bool) {
	interactive := false
	if len(argv) > 0 {
		interactive = argv[0]
	}
	var numberOfInput int
	if _, err := fmt.Scanln(&numberOfInput); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if interactive {
		fmt.Println("How many string you have?")
	}
	reader := bufio.NewReader(os.Stdin)
	var output out
	for i := 0; i < numberOfInput; i++ {
		fmt.Printf("What is your %d string?", i)
		s, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		splitted := strings.Fields(s)
		sort.Sort(sort.Reverse(sort.StringSlice(splitted)))
		output = append(output, strings.Join(splitted, " "))
	}
	fmt.Print(output)
}
