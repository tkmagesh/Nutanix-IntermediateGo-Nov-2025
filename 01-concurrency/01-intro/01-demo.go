package main

import (
	"fmt"
	"time"
)

func main() {
	go f1() //scheduling the execution of "f1()" through the scheduler to be executed in "future"

	f2()

	// block the main() function (by performing an IO operation) execution so that the scheduler will schedule the f1() through the OS thread (Cooperative multitasking by the go scheduler)

	// DO NOT do this
	// time.Sleep(2 * time.Second)

	// adjust the time so that we wait until f1 is completed
	time.Sleep(5 * time.Second)
	// fmt.Scanln()
}

func f1() {
	fmt.Println("f1 started")
	// simulate a time consuming operation
	time.Sleep(4 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
