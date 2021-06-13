package syncexamples

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func WaitGroupExample() {
	var wg sync.WaitGroup

	sayHello := func(name string) {
		defer wg.Done() // decrements counter
		fmt.Println("Hello, " + name)
	}

	log.Print(time.Now())

	wg.Add(2)

	go sayHello("Mallory")
	go sayHello("Jonathan")

	wg.Wait() // join point

	log.Print(time.Now())
}
