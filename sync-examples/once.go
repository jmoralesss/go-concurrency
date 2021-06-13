package syncexamples

import (
	"fmt"
	"sync"
)

func OnceExample() {
	var count int

	increment := func() {
		count++
	}

	decrement := func() {
		count--
	}

	var once sync.Once

	once.Do(increment)
	once.Do(decrement) // ignored

	fmt.Printf("Count: %d\n", count) // output is "Count: 1"
}
