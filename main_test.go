package main

import "os"

func Example() {
	os.Args = []string{"", "Salut ça va ?"}
	main()
	// Output: slt sa va?
}
