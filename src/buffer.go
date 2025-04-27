package src

import "fmt"

type JobQueue struct {
	Jq         chan Job
	bufferSize int
}

type Job struct {
	Idx string `json:"idx"`
}

// Write Only : chan <- Job
// Read Only :  <- chan Job
func NewQueue(bufferSize int) JobQueue {
	return JobQueue{
		Jq:         make(chan Job, bufferSize),
		bufferSize: bufferSize,
	}
}

func (j JobQueue) Producer(v Job) {
	j.Jq <- v
	fmt.Printf("Producer buffer size: %d/%d\n", len(j.Jq), j.bufferSize)
}

func (j JobQueue) Consumer() {
	for v := range j.Jq {
		process(v)
		fmt.Printf("Consumer buffer size: %d/%d\n", len(j.Jq), j.bufferSize)
	}
}

func process(job Job) {
	fmt.Println("consumer : ", job)
}
