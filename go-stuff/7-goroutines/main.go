package main

import (
	"fmt"
	"sync"
	"time"
)

var WORK_ITEMS = [5]string{"id1", "id2", "id3", "id4", "id5"}
var DB []string

const SLEEP_TIME = time.Duration(1) * time.Second

func main() {
	// Go has real multi-threading. It actually runs everything parallel provided your CPU is multi-core. 
	// Launching a goroutine (synonymous to a thread) is pretty simple. This might be why people prefer Go.
	// It seems to have simplified working with systems where multi-threaded is required.

	t0 := time.Now()

	// This blocks the main goroutine or thread
	process_sync()

	fmt.Printf("Time taken to run on main goroutine, i.e., Blocking: %v\n", time.Since(t0))
	fmt.Printf("\n")

	// WaitGroups are just counters which are used to tell Go's runtime that we need to wait for goroutines to finish
	var wg sync.WaitGroup

	// Run a Goroutine by just adding go before the function call.
	// On its own without a wait-group or channel, the main goroutine doesn't wait for goros to finish.
	t0 = time.Now()

	for _, i := range WORK_ITEMS {
		wg.Add(1)
		go work(i, &wg)
	}

	// This blocks  parent/invoking goroutine (or "awaits") until wg reaches 0.
	wg.Wait()

	fmt.Printf("Time taken to run on sub-goroutines, i.e., Concurrent: %v\n", time.Since(t0))

	// It internally invokes wg.Wait blocking the main goroutine
	// But processes each item in sub-goroutines for concurrency
	// This shows that a function may internally make use of wait-group or channels
	// The invoker will not know if a func will block unless the body of the func is explored.
	t0 = time.Now()

	process_async()

	fmt.Printf("Time taken to run on sub-goroutines, i.e., Concurrent: %v\n", time.Since(t0))
	fmt.Printf("\n")

	// Safe mutations in go-routines require the use of sync.Mutex
	var mutex sync.RWMutex

	for _, item := range WORK_ITEMS {
		// This will cause undefined behavior. Multiple goros are writing to same memory location.
		// Note: Comment out all of the code above before running, otherwise the side-effect won't appear.
		// You may need to run it multiple times to see the side-effect.
		// wg.Go(func() { unsafeWriteToDb(item) })

		// Must pass the reference to mutex. 
		// Value mutex will be different for each call, i.e., each goro will get their own mtx.
		wg.Go(func() { safeWriteToDb(item, &mutex) })
	}

	wg.Wait()

	fmt.Printf("DB state %v; Len: %v; Cap: %v\n", DB, len(DB), cap(DB))
}

func process_sync() {
	for _, i := range WORK_ITEMS {
		work(i, nil)
	}
}

func process_async() {
	var wg sync.WaitGroup

	// sync.WorkGroup provides a convenient abstraction over managing wg counters manually.
	for _, i := range WORK_ITEMS {
		// This syntax lets you run Synchronous or Blocking funcs, concurrently.
		// Without managing wg counters yourself.
		wg.Go(func() { work(i, nil) })
	}

	// This blocks main goroutine (or "awaits") until wg reaches 0.
	wg.Wait()
}

/*
work will block the goroutine invoking this function.
If a wait-group is provided, it will invoke wg.Done. All other ceremony of managing wg is on the client.
*/
func work(workId string, wg *sync.WaitGroup) {
	if wg != nil {
		// TIP: This is also why defer is function-scoped.
		// Block scope would defer it until this block exited which is not what we want.
		defer wg.Done()
	}

	time.Sleep(SLEEP_TIME)
	fmt.Printf("Processed item with workId: %v\n", workId)
}

func unsafeWriteToDb(workItem string) {
	fmt.Printf("Writing %v to DB...\n", workItem)
	time.Sleep(SLEEP_TIME)
	DB = append(DB, workItem)
}

func safeWriteToDb(workItem string, mutex *sync.RWMutex) {
	fmt.Printf("Writing %v to DB...\n", workItem)
	time.Sleep(SLEEP_TIME)

	mutex.Lock()
	DB = append(DB, workItem)
	mutex.Unlock()
}
