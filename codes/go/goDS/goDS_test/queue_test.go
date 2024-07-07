package goDS_test

import (
	"fmt"
	"goDSA/goDS"
	"testing"
)

func TestQueue(t *testing.T) {
	q := goDS.NewQueue[int](10)

	for i := 0; i < 20; i++ {
		q.Enqueue(i)
	}

	for i := 0; i < 20; i++ {
		fmt.Println(q.Dequeue())
	}
}
