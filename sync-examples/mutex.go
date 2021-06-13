package syncexamples

import (
	"fmt"
	"sync"
)

func MutexExample() {
	var count int
	var lock sync.Mutex

	increment := func() {
		lock.Lock()
		defer lock.Unlock()

		beforeIncrement := count
		count = count + 1
		fmt.Println("Incrementing: " + fmt.Sprint(beforeIncrement) + " -> " + fmt.Sprint(count))
	}

	decrement := func() {
		lock.Lock()
		defer lock.Unlock()

		beforeDecrement := count
		count = count - 1
		fmt.Println("Decrementing: " + fmt.Sprint(beforeDecrement) + " -> " + fmt.Sprint(count))
	}

	var arithmetic sync.WaitGroup

	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)

		go func() {
			defer arithmetic.Done()
			increment()
		}()
	}

	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)

		go func() {
			defer arithmetic.Done()
			decrement()
		}()
	}

	arithmetic.Wait()
}
