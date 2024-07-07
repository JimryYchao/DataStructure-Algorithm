package goDS_test

import (
	"fmt"
	"goDSA/goDS"
	"testing"
)

func TestStack(t *testing.T) {
	s := goDS.NewStack[int](10)

	for i := 0; i < 20; i++ {
		s.Push(i)
	}
	for i := 0; i < 10; i++ {
		fmt.Println(s.Pop())
	}
}
