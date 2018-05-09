package main

import "fmt"
import "sync"

var wg = new(sync.WaitGroup)

func main() {
	fmt.Println("Hello channels!!")
	fmt.Println()
	ch := make(chan int)
	// will not work, because we have only one receives but 5 senders.. //DEADLOCK
	// go func () {
	// 	i := <- ch
	// 	fmt.Println(i)
	// 	wg.Done()
	// }()
	fmt.Println("Sending from one goroutine and receiving form the other")
	// for index := 0; index < 5; index++ {
	// 	wg.Add(2)

	// 	go func () {
	// 		i := <- ch
	// 		fmt.Println(i)
	// 		wg.Done()
	// 	}()

	// 	go func() {
	// 		ch <- 42
	// 		wg.Done()
	// 	}()
	// }

	fmt.Println("Biderectional sending / receiving \n\n")

	// for index := 0; index < 5; index++ {
	// 	wg.Add(2)

	// 	go func () {
	// 		i := <- ch
	// 		fmt.Println(i)
	// 		ch <- 99
	// 		wg.Done()
	// 	}()

	// 	go func() {
	// 		ch <- 42
	// 		fmt.Println(<-ch)
	// 		wg.Done()
	// 	}()
	// }

	fmt.Println("Send only and receive only channels")
	// for index := 0; index < 5; index++ {
	// 	wg.Add(2)

	// 	go func (ch <-chan int) { //receive only
	// 		i := <- ch
	// 		fmt.Println(i)
	// 		wg.Done()
	// 	}(ch)

	// 	go func(ch chan<- int) { //send only
	// 		ch <- 42
	// 		wg.Done()
	// 	}(ch)
	// }

	fmt.Println("Receiver for more than one values\n")
	ch = make(chan int, 50)
	wg.Add(2)
	go func(ch <-chan int) { //receive only
		// i := <-ch
		// fmt.Println(i)
		// i = <-ch
		// fmt.Println(i)
		for i := range ch {
			fmt.Println(i)
		}

		wg.Done()
	}(ch)

	go func(ch chan<- int) { //send only
		ch <- 42
		ch <- 27
		close(ch) //we have to close the channel because otherwise the channel wait for another values.
		wg.Done()
	}(ch)
	wg.Wait()
}
