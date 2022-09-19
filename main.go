package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"sync"
	"sync/atomic"
)

var (
	mutexCounter int32
	mx           sync.Mutex
	wgMutex      sync.WaitGroup

	mutexErrorCounter int32
	mxError           sync.Mutex

	channelCounter int32
	ch             chan struct{}
	wgChannel      sync.WaitGroup

	atomicCounter int32
	wgAtomic      sync.WaitGroup

	counter = 10000
)

func incMutexCounter() {
	mx.Lock()
	mutexCounter++
	mx.Unlock()
}

func incMutexErrorCounter() {
	mxError.Lock()
	mutexErrorCounter++
	mxError.Unlock()
}

func incChannelCounter() {
	ch <- struct{}{}
	channelCounter++
	<-ch
}

func incAtomicCounter() {
	atomic.AddInt32(&atomicCounter, 1)
}

func main() {
	fmt.Println("runMutexExample start")
	runMutexExample()
	fmt.Println(mutexCounter)
	fmt.Println("runMutexExample end")
	fmt.Println("")

	fmt.Println("runMutexErrorExample start")
	runMutexErrorExample()
	fmt.Println(mutexErrorCounter)
	fmt.Println("runMutexErrorExample end")
	fmt.Println("")

	fmt.Println("runChannelExample start")
	runChannelExample()
	fmt.Println(channelCounter)
	fmt.Println("runChannelExample end")
	fmt.Println("")

	fmt.Println("runAtomicExample start")
	runAtomicExample()
	fmt.Println(atomicCounter)
	fmt.Println("runAtomicExample end")
}

func runMutexExample() {
	mx = sync.Mutex{}
	wgMutex = sync.WaitGroup{}

	wgMutex.Add(counter)
	for i := 0; i < counter; i++ {
		go func() {
			incMutexCounter()
			wgMutex.Done()
		}()
	}
	wgMutex.Wait()
}

func runMutexErrorExample() {
	mxError = sync.Mutex{}
	g := new(errgroup.Group)

	for i := 0; i < counter; i++ {
		g.Go(func() error {
			incMutexErrorCounter()
			return nil
		})
	}
	g.Wait()
}

func runChannelExample() {
	ch = make(chan struct{}, 1)
	wgChannel = sync.WaitGroup{}

	wgChannel.Add(counter)
	for i := 0; i < counter; i++ {
		go func() {
			incChannelCounter()
			wgChannel.Done()
		}()
	}
	wgChannel.Wait()
}

func runAtomicExample() {
	wgAtomic = sync.WaitGroup{}

	wgAtomic.Add(counter)
	for i := 0; i < counter; i++ {
		go func() {
			incAtomicCounter()
			wgAtomic.Done()
		}()
	}
	wgAtomic.Wait()
}
