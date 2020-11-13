package leetcode

import "testing"

func TestFizzBuzz(t *testing.T) {

	in := 16

	want := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz", "16"}

	out := FizzBuzzL0(in)

	if len(out) != len(want) {
		t.Error("array size want: ", len(want), " but out: ", len(out))
		t.FailNow()
	}

	for i := 0; i < len(want); i++ {
		if out[i] != want[i] {
			t.Error("want[", i, "]: ", want[i], " but out[", i, "]=", out[i])
			t.FailNow()
		}
	}
}
