package main

import (
	"sync"
	"fmt"
	//"runtime"
)

var wg = sync.WaitGroup{}
var m = new(sync.RWMutex)

var counter = 0
func main(){
	fmt.Println("Hey GORoutines!!")

	// go sayHello()
	// time.Sleep(100 * time.Millisecond)


	// var msg = "Hello with WaitGroup!"
	// wg.Add(1)
	// go func(msg string){
	// 	fmt.Println(msg)
	// 	wg.Done()
	// }(msg)
	//runtime.GOMAXPROCS(100) -> set the threads of the application
	for index := 0; index < 10; index++ {
		wg.Add(2)
		m.RLock()
		go sayNewHello()
		m.Lock()
		go increment()
	}
	wg.Wait()
	
}

func sayNewHello(){
	fmt.Println("Hello", counter, "!!")
	m.RUnlock()
	wg.Done()
}

func increment(){
	counter++
	m.Unlock()
	wg.Done()
}

func sayHello(){
	fmt.Println("Hello!!")
	
}

