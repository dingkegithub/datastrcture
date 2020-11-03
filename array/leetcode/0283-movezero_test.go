package leetcode

import (
	"math/rand"
	"testing"
	"time"
)

func TestMoveZeros(t *testing.T) {
	rand.Seed(time.Now().Unix())
	idx := rand.Intn(len(sample))

	testSample := sample[idx]
	PrintArrays("sample", testSample)
	PrintArrays("want  ", moveZeroWant[idx])

	MoveZero(testSample)
	PrintArrays("out   ", moveZeroWant[idx])
	if !Equal(testSample, moveZeroWant[idx]) {
		t.FailNow()
	}
}
