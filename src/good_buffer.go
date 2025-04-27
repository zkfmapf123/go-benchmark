package src

import (
	"context"
	"sync"
)

type Priority int

const (
	Low Priority = iota
	Medium
	High
)

type JobStatus string

const (
	Pending    JobStatus = "pending"
	Processing JobStatus = "processing"
	Completed  JobStatus = "completed"
	Failed     JobStatus = "failed"
	Timeout    JobStatus = "timeout"
)

type GoodQueue struct {
	jq      chan Job
	workers int
	wg      sync.WaitGroup
	ctx     context.Context
	cancel  context.CancelFunc
}

func NewGoodQueue(bufferSize, workers int) *GoodQueue {
	ctx, cancel := context.WithCancel(context.Background())
	q := &GoodQueue{
		jq:      make(chan Job, bufferSize),
		workers: workers,
		ctx:     ctx,
		cancel:  cancel,
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
					process(job)
				}
			}
		}(i)
	}
}

func (j *GoodQueue) Producer(v Job) error {
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
	j.cancel()
	close(j.jq)
	j.wg.Wait()
}
