package stackqueue

import "testing"

func TestStack(t *testing.T) {
	in := []int{3, 4, 6, 1, 2}

	stack := NewStack()
	if !stack.Empty() {
		t.Error("stack want: empty true, but: empty ", stack.Empty())
		t.FailNow()
	}

	for _, v := range in {
		stack.Push(v)
	}

	if stack.Empty() {
		t.Error("stack want: empty false, but: empty ", stack.Empty())
		t.FailNow()
	}

	d, err := stack.Peek()
	if err != nil {
		t.Error("stack Peek error: ", err)
		t.FailNow()
	}

	if d.(int) != 2 {
		t.Error("stack Peek want 2 but: ", d.(int))
		t.FailNow()
	}

	// pop operation
	d, err = stack.Pop()
	if err != nil {
		t.Error("stack Pop error: ", err)
		t.FailNow()
	}

	if d.(int) != 2 {
		t.Error("stack Pop want 2 but: ", d.(int))
		t.FailNow()
	}

	// search
	if idx := stack.Search(1); idx != 3 {
		t.Error("stack Search want 3 but: ", idx)
		t.FailNow()
	}

	for i := 3; i >= 0; i-- {
		d, err := stack.Pop()
		if err != nil {
			t.Error("stack pop ", i, " failed", err)
			t.FailNow()
		}

		if in[i] != d.(int) {
			t.Error("stack[", i, "] want: ", in[i], " but: ", d.(int))
			t.FailNow()
		}
	}

	_, err = stack.Pop()
	if err == nil {
		t.Error("pop empty stack, want error, but nil")
		t.FailNow()
	}
}
