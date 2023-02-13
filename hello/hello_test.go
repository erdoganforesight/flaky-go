package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestExample(t *testing.T) {
	// Passed test
	t.Run("passed test", func(t *testing.T) {
		if 1+1 == 2 {
			t.Log("1+1 equals 2 as expected")
		} else {
			t.Error("1+1 does not equal 2")
		}
	})

	// // Failed test
	t.Run("failed test", func(t *testing.T) {
		if 1+1 == 3 {
			t.Log("1+1 equals 3 as expected")
		} else {
			t.Error("1+1 does not equal 3")
		}
	})

	// Skipped test
	t.Run("skipped test", func(t *testing.T) {
		t.Skip("skipping this test")
	})

	// Flaky test
	t.Run("flaky test", func(t *testing.T) {
		min, max := 1, 10
		rand.Seed(time.Now().UnixNano())
		count := rand.Intn(max - min)
		fmt.Print("Name: ", count)
		if 0 <= count && count < 3 {
			t.Log("test passed on the third execution")
		} else {
			t.Error("test failed on execution number", count)
		}
	})
}
