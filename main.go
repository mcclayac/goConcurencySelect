package main

import (
	"fmt"
	"time"
)

func emit(wordChannel chan string, done chan bool) {
	words := []string{"the", "quikc", "brown", "fox"}
	defer close(wordChannel)
	i := 0

	t := time.NewTimer(2 * time.Second)

	for {

		select {
		case wordChannel <- words[i]:
			i++
			if i == len(words) {
				i = 0
			}
		case <-done:
			fmt.Printf("got Done Message\n")
			//close(done)
			done <- true // channels are bi-directional
			// bad if I took out.. program will block
			//fmt.Printf("got Done Message\n")
			//close(wordChannel)
			return

		case <-t.C:
			fmt.Printf("\n\nTimer Fired\n\n")
			return
		}

	}
	//fmt.Printf("\nCompleted\n")
	//
	//close(wordChannel)

}


func main() {

	wordChannel := make(chan string)
	doneChannel := make(chan bool)

	go emit(wordChannel, doneChannel)

	//for i:= 0; i < 200; i++ {
	//	fmt.Printf("word: %s\t\t", <-wordChannel)
	//}

	for word := range wordChannel {
		fmt.Printf("word: %s\t\t", word)
	}

	// Channels are bi-directional

	fmt.Printf("\n\n")
	//doneChannel <- true
	//<- doneChannel   // will not terminate unless something is recieved

}
