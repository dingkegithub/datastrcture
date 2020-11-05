package stackqueue

import "testing"

func TestDqueue(t *testing.T) {
	in := []int{1, 2, 3}
	dq := NewDequeue(WithDqueueCapacity(3))

	if err := dq.AddFirst(in[0]); err != nil {
		t.Error("add 1 first: ", err)
		t.FailNow()
	}

	if err := dq.AddLast(in[1]); err != nil {
		t.Error("add 2 last: ", err)
		t.FailNow()
	}

	if err := dq.AddFirst(in[2]); err != nil {
		t.Error("add 3 first err: ", err)
		t.FailNow()
	}

	d := dq.PeekFirst()
	if d == nil {
		t.Error("peek 3 nil")
		t.FailNow()
	}

	if d.(int) != in[2] {
		t.Error("peek want: ", in[2], " but: ", d.(int))
		t.FailNow()
	}

	d = dq.PeekLast()
	if d == nil {
		t.Error("peek 2 nil")
		t.FailNow()
	}

	if d.(int) != in[1] {
		t.Error("peek want: ", in[1], " but: ", d.(int))
		t.FailNow()
	}

	// 3,1,2
	d = dq.PollLast()
	if d == nil {
		t.Error("poll last nil")
		t.FailNow()
	}

	if d.(int) != in[1] {
		t.Error("poll want: ", in[1], " but: ", d.(int))
		t.FailNow()
	}

	// 3,1
	d = dq.PollLast()
	if d == nil {
		t.Error("poll last nil")
		t.FailNow()
	}

	if d.(int) != in[0] {
		t.Error("poll want: ", in[0], " but: ", d.(int))
		t.FailNow()
	}

	l := dq.PeekLast()
	f := dq.PeekFirst()

	if l == nil || f == nil {
		t.Error("peek nonnil, but: ", l, f)
		t.FailNow()
	}

	if l.(int) != f.(int) {
		t.Error("peek last != first", l, f)
		t.FailNow()
	}

	d, err := dq.RemoveFirst()
	if err != nil {
		t.Error("remove first failed: ", err)
		t.FailNow()
	}

	if d.(int) != f.(int) {
		t.Error("peek last != first", d, f)
		t.FailNow()
	}

	_, err = dq.RemoveLast()
	if err == nil {
		t.Error("RemoveLast want: ", err, "but nil")
		t.FailNow()
	}
}
