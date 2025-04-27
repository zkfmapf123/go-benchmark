package src

import (
	"testing"

	"github.com/google/uuid"
)

func BenchmarkConcurrent(b *testing.B) {
	q := NewQueue(1000)
	done := make(chan bool)

	// Consumer 시작
	go func() {
		q.Consumer()
		done <- true
	}()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Producer(Job{
			Idx: uuid.New().String(),
		})
	}

	// 채널을 닫아서 Consumer 종료
	close(q.Jq)
	<-done
}

func BenchmarkGoodConcurrent(b *testing.B) {

	q := NewGoodQueue(1000, 10, Process)
	done := make(chan bool)

	go func() {
		q.Close()
		done <- true
	}()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Producer(Job{
			Idx: uuid.New().String(),
		})
	}

	<-done
}
