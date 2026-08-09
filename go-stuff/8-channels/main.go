package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

const MAX_CHANNEL_CAP = 5

func main() {
	// Channels are thread safe across goroutines. Consider them an alternative to mutex when you don't need as much control.
	// unbuff_chan := make(chan int)
	buff_chann := make(chan int, 3)

	// Write to channel like this.
	// This blocks the invoking goro until some other goro reads from it.
	// If you run this as-is, go will throw a deadlock at your face.
	buff_chann <- 1

	// This will never run because the main goro will keep waiting for parent goro to read from chan.
	// There is no parent goro for main controlled by us.
	fmt.Printf("Value of my (oni)-chan: %v\n", <-buff_chann)

	// Correct method of using channels:

	// Use buffered Channels with predefined length.
	// It allows writers to continue writing to it continuously until the it's full.
	// Note: The buffered channel capacity is FIXED. It doesn't grow beyond of what you specify in make.
	// Note: `nil` until something is written.
	buff_chan := make(chan int, MAX_CHANNEL_CAP)

	// Reading from `nil` channels block forever.
	// <-buff_chan

	// Spawn one goro to execute writeToChan concurrently
	go writeToChanAndClose(buff_chan, 1)

	// read from channel in the parent/main goroutine.
	// This blocks until the channel is closed.
	readFromChan(buff_chan, 1)

	// reading from a closed channel returns undefined flavor of respective type
	// It doesn't causes any compile-time or runtime errors.
	v, isOpen := <-buff_chan
	fmt.Printf("Value from closed buff_chan: %v; Is buff_chan open: %v\n", v, isOpen)

	// Writing is unsafe tho.xR
	// The following will cause a panic at runtime, compiles fine.
	// buff_chan <- 80

	fmt.Printf("\n")

	singleReaderMultiWriter()

	fmt.Printf("\n")

	singleWriterMultiReader()

	fmt.Printf("\n")

	channelSync()
}

// Function showcasing multiple channel writers but a single reader
// The order of writes cannot be guaranteed as its up-to the scheduler which goro(s) get to write next.
// Note: This clears that FIFO or queue behavior cannot be achieved when there are multiple writers.
// Channels are merely a construct to pass data in a thread-safe manner where multiple goro(s) could be writing to.
func singleReaderMultiWriter() {
	buff_chan := make(chan int, MAX_CHANNEL_CAP)
	wg := sync.WaitGroup{}

	// Spawn 3 goroutines to execute writeToChan concurrently
	for i := range 3 {
		wg.Go(func() { writeToChan(buff_chan, i+1) })
	}

	// Closing responsibility is now on us.
	// Because one goro might finish before the other and close the channel.
	// Hence, we spawn a goro and Wait. Why Wait in goro and not in invoking function's body?
	// Because it will block execution of this function and we will run into deadlock as no one would be reading from channel.
	go func() {
		wg.Wait()
		close(buff_chan)
	}()

	// read from channel in the parent/main goroutine.
	// This blocks until the channel is closed.
	readFromChan(buff_chan, 1)
}

// Function showcasing multiple channel readers but a single writer
// The order of reads is also non-deterministic. If a channel has multiple consumers,
// It cannot be guaranteed what value will be received by which channel.
func singleWriterMultiReader() {
	buff_chan := make(chan int, MAX_CHANNEL_CAP)
	wg := sync.WaitGroup{}

	wg.Go(func() { writeToChanAndClose(buff_chan, 1) })

	for i := range 3 {
		wg.Go(func() { readFromChan(buff_chan, i+1) })
	}

	wg.Wait()
}

// Function showcasing consumer pattern using select which is like switch for channels
// Multiple Channels serving different content being consumed & handled at a single place
// The order of consumption is non-deterministic.
func channelSync() {
	ch1 := make(chan int)
	ch2 := make(chan rune)

	go func() {
		ch1 <- 1
		close(ch1)
	}()

	go func() {
		ch2 <- '😒'
		close(ch2)
	}()

	// We have added a for because we wanted to exhaust reading from all channels
	// The condition is basically there to ensure no channels exist, this is the while <cond>
	// Without the condition, it is an infinite loop as reading from closed channel never blocks.
	for ch1 != nil || ch2 != nil {
		// Select *selects* an case randomly provided the channel has resolved then exits
		select {
		case v, ok := <-ch1:
			if !ok {
				ch1 = nil
				continue
			}

			fmt.Printf("Channel 1 responded with: %v\n", v)

		case v, ok := <-ch2:
			if !ok {
				ch2 = nil
				continue
			}

			fmt.Printf("Channel 2 responded with: %c\n", v)
		}
	}

}

func writeToChanAndClose(c chan<- int, channelId int) {
	defer close(c)
	writeToChan(c, channelId)
}

func writeToChan(c chan<- int, channelId int) {
	defer fmt.Printf("Writing to Channel #%v: %v finished.\n", channelId, c)

	// Random between 0-5, i.e., 0-4
	// +1 because don't want 0 and want 5.
	multiplier := rand.IntN(5) + 1

	for i := range MAX_CHANNEL_CAP {
		// This will keep writing until channel is full or MAX_CHANNEL_CAP is reached.
		// If the channel is full but MAX_CHANNEL_CAP is > chan_cap, it will block until
		// more room is made in the channel by the reader(s).
		c <- i * multiplier
	}
}

func readFromChan(c <-chan int, channelId int) {
	// This is invalid - results in compile-time error as it's a receive-only channel.
	// Sender owns the lifecycle of channel. Not receiver.
	// defer close(c)

	// Range provides convenience over reading from channels.
	// The convenience mainly lies in decoupling the reader from knowing the length of chan beforehand.
	// It blocks whenever the channel is empty.
	// Exits once the channel has been closed AND there are no more values to be consumed.
	for chan_val := range c {
		fmt.Printf("Reader #%v received %v\n", channelId, chan_val)
		time.Sleep(time.Second / 4)
	}

	// Note this is typically how range knows there are no more values to be read
	// for {
	// 	v, hasMore := <-c

	// 	if !hasMore {
	// 		break
	// 	}

	// 	fmt.Printf("Reader #%v received %v\n", channelId, v)
	// 	time.Sleep(time.Second / 4)
	// }
}
