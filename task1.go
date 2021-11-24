package main

import (
	"fmt"
	"sync"
)

// This programs demonstrates how a channel can be used for sending and
// receiving by any number of goroutines. It also shows how  the select
// statement can be used to choose one out of several communications.
func main() {
	people := []string{"Anna", "Bob", "Cody", "Dave", "Eva"}
	match := make(chan string, 1) // Make room for one unmatched send.
	wg := new(sync.WaitGroup)
	//var wg sync.WaitGroup
	wg.Add(len(people))
	for _, name := range people {
		go Seek(name, match, wg)
	}
	wg.Wait()
	select {
	case name := <-match:
		fmt.Printf("No one received %s’s message.\n", name)
	default:
		// There was no pending send operation.
	}
}

// Seek either sends or receives, whichever possible, a name on the match
// channel and notifies the wait group when done.
func Seek(name string, match chan string, wg *sync.WaitGroup) {
	select {
	case peer := <-match:
		fmt.Printf("%s sent a message to %s.\n", peer, name)
	case match <- name:
		// Wait for someone to receive my message.
	}
	wg.Done()
}

//What happens if you remove the go-command from the Seek call in the main function?

//seek blir inte kallad i goroutinen, istället en åt gången, programmet väntar alltså
//på att varje seek anrop ska avslutas, men inget händer "i praktiken"

//what happens if you switch the declaration wg := new(sync.Waitgroup) to var wg sync.WaitGroup
// and the parameter wg *sync.WaitGroup to wg sync.WaitGroup

//alla goroutiner "instanser" av seek kan komma åt samma wg på ett "bra" sätt
//markerar som done vid avslut, att ändra koden skapar error pga type i seek argumenter
//alltså det ska va en wg och inte en var

//What happens if you remove the buffer on the channel match?

// när man tar bort buffert måste man ha "rätt" antal sendare å recievare
//eftersom det inte är så, blir det deadlock

//What happens if you remove the default-case from the case-statement in the main function?

//Default är empty ju, så inget händer
