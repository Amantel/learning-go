package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func Run(tasks []Task, n, m int) error {
	var wg sync.WaitGroup
	ch := make(chan Task, len(tasks))
	var errorCount int32

	for _, task := range tasks {
		ch <- task
	}

	close(ch)

	for i := 0; i < n; i++ {
		wg.Add(1)

		go func() {

			for task := range ch {

				func() {
					defer func() {
						if r := recover(); r != nil {
							atomic.AddInt32(&errorCount, 1)
						}
					}()
					error := task()
					if error != nil {
						atomic.AddInt32(&errorCount, 1)
					}
				}()

				if atomic.LoadInt32(&errorCount) > int32(m) {
					break
				}
			}
			wg.Done()
		}()
	}

	wg.Wait()

	if errorCount > int32(m) {
		return ErrErrorsLimitExceeded
	} else {
		return nil
	}
}
