package leetcode

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestClimbStairs(t *testing.T) {
	rand.Seed(time.Now().Unix())

	for i := 1; i < 23; i++ {
		// l1 := ClimbStairsLevel1(i)
		// l2 := ClimbStairsLevel2(i)
		l3 := ClimbStairsLevel3(i)
		l4 := ClimbStairsLevel4(i)
		// l5 := ClimbStairsLevel5(i)

		fmt.Println("in:    ", i)
		fmt.Println("want:  ", fiblq[i])
		//fmt.Println("out:   ", l3, l4, l5)
		fmt.Println("out:   ", l3, l4)

		if l3 != l4 || l4 != fiblq[i] {
			// if l5 != l3 || l3 != l4 || l4 != fiblq[i] {
			t.FailNow()
		}
	}
}
