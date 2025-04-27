package src

import (
	"context"
	"sync"
	"sync/atomic"
)

type GoodQueue struct {
	jq          chan Job
	workers     int
	wg          sync.WaitGroup
	ctx         context.Context
	cancel      context.CancelFunc
	closed      atomic.Bool
	processFunc func(Job)
}

func NewGoodQueue(bufferSize, workers int, processFunc func(Job)) *GoodQueue {
	ctx, cancel := context.WithCancel(context.Background())
	q := &GoodQueue{
		jq:          make(chan Job, bufferSize),
		workers:     workers,
		ctx:         ctx,
		cancel:      cancel,
		processFunc: processFunc,
	}
	q.startWorkers()
	return q
}

func (j *GoodQueue) startWorkers() {
	for i := 0; i < j.workers; i++ {
		j.wg.Add(1)
		go func(workerID int) {
			defer j.wg.Done()
			for {
				select {
				case <-j.ctx.Done():
					return
				case job, ok := <-j.jq:
					if !ok {
						return
					}
					j.processFunc(job)
				}
			}
		}(i)
	}
}

func (j *GoodQueue) Producer(v Job) error {
	if j.closed.Load() {
		return context.Canceled
	}

	select {
	case <-j.ctx.Done():
		return j.ctx.Err()
	case j.jq <- v:
		return nil
	default:
		return context.DeadlineExceeded
	}
}

func (j *GoodQueue) Close() {
	if !j.closed.CompareAndSwap(false, true) {
		return // 이미 닫혔음
	}
	j.cancel()
	close(j.jq)
	j.wg.Wait()
}
