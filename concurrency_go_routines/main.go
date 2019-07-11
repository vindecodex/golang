package main

import (
	"fmt"
	"time"
)

/*
Concurrency is execution of a function in any order, partial or in parallel
Concurrency in golang is Go routine
To achieve a routine you will just need to add go on the left side of a function
*/

func a() {
	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}
	time.Sleep(time.Second * 2) // we set a sleep of 2 * millisecond
}

func main() {
	go a()                      // was set a sleep
	time.Sleep(time.Second * 2) // was set a sleep
	for i := 0; i < 100; i++ {  // a() and this will run parallel because we set their both time the same
		fmt.Println(i)
	}
}

// Remove the go before invocation of a() to see the difference :)
/*
It is fine to create a lot of go routines because it is a lite weight and wont harm you program :)
*/
