package main

import (
	"fmt"
	"time"
)

type Semaphore struct {
	sem chan struct{}
}

func New(num int) *Semaphore {
	return &Semaphore{
		sem: make(chan struct{}, num),
	}
}

func (s *Semaphore) Acquire() {
	s.sem <- struct{}{}
}

func (s *Semaphore) Release() {
	<-s.sem
}

func main() {
	sem := New(3)
	done := make(chan bool, 1)
	totalprocess := 10
	for i := 0; i <= 10; i++ {
		sem.Acquire()
		go func(v int) {
			defer sem.Release()
			func(n int) {
				fmt.Println(time.Now())
				time.Sleep(2 * time.Second)
			}(v)
			if v == totalprocess {
				done <- true
			}
		}(i)
	}
	fmt.Println(<-done)
}
