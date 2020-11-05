package stackqueue

import "testing"

func TestQueue(t *testing.T) {
	in := []int{1, 2, 3, 4, 5}

	q := NewQueue(WithQueueCapacity(4))

	for i := 0; i < 4; i++ {
		ok, err := q.Add(in[i])
		if err != nil {
			t.Error("queue add element failed: ", in[i], err)
			t.FailNow()
		}

		if !ok {
			t.Error("queue add false ")
			t.FailNow()
		}
	}

	ok, err := q.Add(in[4])
	if err == nil {
		t.Error("queue add element want: err, but: nil")
		t.FailNow()
	}

	if ok {
		t.Error("queue add want false, but true ")
		t.FailNow()
	}

	d, err := q.Remove()
	if err != nil {
		t.Error("queue remove failed: ", err)
		t.FailNow()
	}

	if d.(int) != 1 {
		t.Error("queue remove want: 1, but: ", d.(int))
		t.FailNow()
	}

	d = q.Peek()
	if d == nil {
		t.Error("queue peek failed with nil")
		t.FailNow()
	}

	if d.(int) != 2 {
		t.Error("queue peek want: 2, but: ", d.(int))
		t.FailNow()
	}
}
