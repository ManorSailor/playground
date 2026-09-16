package main

import (
	"fmt"
	"time"
)

type Task func()

// Using Channel here because it guarantees two things:
// FIFO & Blocks without busy-looping, go already handles it for us internally

// Processed ONE task at time, new additions are processed on next round trip
type MacroTaskQ chan Task

// Contains tasks having priority in execution over MacroQ tasks
// Processed until empty, any new additions are processed right away
// Processed only when the Stack is EMPTY
// A Slice because this is only ever modified by the event loop goroutine
// Using golang's concurrency features (go func()) will cause race conditions
type MicroTaskQ []Task

type IEventLoop interface {
	RunUntilCompletion()
	QueueMicroTask(Task)
	SetTimeout(Task, int)
}

type EventLoop struct {
	macroQ       MacroTaskQ
	microQ       MicroTaskQ
	macroTaskCtr uint
}

func NewLoop() IEventLoop {
	// Need to use buffered channels here otherwise a deadlock would be raised
	// NOT A FIX. JUST DELAYING THE INEVITABLE.
	return &EventLoop{
		macroQ:       make(MacroTaskQ, 5),
		microQ:       make(MicroTaskQ, 0, 5),
		macroTaskCtr: 0,
	}
}

func (evl *EventLoop) QueueMicroTask(task Task) {
	evl.microQ = append(evl.microQ, task)
}

func (evl *EventLoop) RunUntilCompletion() {
	for evl.macroTaskCtr > 0 || len(evl.microQ) > 0 {
		for {
			if len(evl.microQ) == 0 {
				fmt.Printf("No more micro tasks queued; Moving on...\n")
				break
			}

			micTask := evl.microQ[0]
			evl.microQ = evl.microQ[1:]

			micTask()
		}

		if evl.macroTaskCtr == 0 {
			continue
		}

		// This is not identical to how it is defined in the HTML Spec.
		// We don't wait indefinitely for macro tasks to appear. We wait for EITHER.
		macTask, isOpen := <-evl.macroQ
		if !isOpen {
			break
		}

		macTask()
		evl.macroTaskCtr--
	}

}

func (evl *EventLoop) SetTimeout(callback Task, delay int) {
	if delay < 0 {
		delay = 0
	}

	evl.macroTaskCtr++

	go func() {
		time.Sleep(time.Millisecond * time.Duration(delay))
		evl.macroQ <- callback
	}()
}

func main() {
	ev := NewLoop()

	ev.SetTimeout((func() {
		ev.QueueMicroTask((func() {
			fmt.Printf("Executing MicroTask 2\n")
		}))

		fmt.Printf("Executing MacroTask 2\n")

		ev.SetTimeout((func() {
			fmt.Printf("Executing MacroTask 2.1\n")
		}), 500)
	}), 1000)

	ev.QueueMicroTask((func() {
		fmt.Printf("Executing MicroTask 1\n")

		ev.QueueMicroTask((func() {
			fmt.Printf("Executing MicroTask 1.1\n")
		}))

		ev.SetTimeout((func() {
			fmt.Printf("Executing MacroTask 1\n")

			ev.QueueMicroTask((func() {
				fmt.Printf("Executing MicroTask 1.2\n")
			}))
		}), 500)
	}))

	ev.RunUntilCompletion()
}
