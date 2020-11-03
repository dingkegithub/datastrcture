package leetcode

import (
	"math/rand"
	"testing"
)

func TestMoveZeros(t *testing.T) {
	idx := rand.Intn(len(sample))

	testSample := sample[idx]

	MoveZero(testSample)
	if Equal(testSample, moveZeroWant[idx]) {
		t.FailNow()
	}
}
