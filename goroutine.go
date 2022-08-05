package main

/*
import (
	"fmt"
	"time"
)

// notice we've not changed anything in this function
// when compared to our previous sequential program
func compute(value int) {
	for i := 0; i < value; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}

func main() {
	fmt.Println("Goroutine Tutorial")

	// notice how we've added the 'go' keyword
	// in front of both our compute function calls
	go compute(10)
	go compute(10)
	// we make our anonymous function concurrent using `go`
	go func() {
		fmt.Println("Executing my Concurrent anonymous function")
	}()
	// we have to once again block until our anonymous goroutine
	// has finished or our main() function will complete without
	// printing anything
	var input string
	fmt.Scanln(&input)
}
*/
