package leetcode

import "testing"

func TestRemoveOuterParentheses(t *testing.T) {
	sample := []string{
		"(()())(())",
		"(()())(())(()(()))",
		"()()",
	}

	want := []string{
		"()()()",
		"()()()()(())",
		"",
	}

	for i, s := range sample {
		out := RemoveOuterParentess(s)
		if out != want[i] {
			t.Error("sample[", i, "] want: ", want[i], " but: ", out)
			t.FailNow()
		}
	}
}
