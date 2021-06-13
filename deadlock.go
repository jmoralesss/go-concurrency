package main

import (
	"fmt"
	"sync"
	"time"
)

type value struct {
	mu    sync.Mutex
	value int
}

/*
	Conditions for a Deadlock

		Mutual Exclusion
			A concurrent process holds exclusive rights to a resource at any time

		Wait for Condition
			A concurrent process most simultaneosly hold a resource and be waiting
			for an additional resource

		No Preemption
			A resource held by a concurrent process can only be released by that
			process

		Circular Wait
			A concurrent process (P1) must be waiting on a chain of other concurrent
			processes (P2), which are in turn waiting on it (P1)

	A deadlock can be avoided by ensuring that at least one condition is not true
*/

func DeadlockExample() {
	var wg sync.WaitGroup

	printSum := func(v1, v2 *value) {
		defer wg.Done()

		v1.mu.Lock()
		defer v1.mu.Unlock()

		time.Sleep(2 * time.Second)
		v2.mu.Lock()
		defer v2.mu.Unlock()

		fmt.Printf("sum=%v\n", v1.value+v2.value)
	}

	var a, b value

	wg.Add(2)

	go printSum(&a, &b)
	go printSum(&b, &a)
	wg.Wait()
}
