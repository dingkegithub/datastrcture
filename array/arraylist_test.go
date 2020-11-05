package array

import "testing"

func TestArrayList(t *testing.T) {
	want := []int{1, 2, 3, 4, 5, 4, 3, 2, 1}

	arrList := NewArrayList()
	for _, v := range want {
		arrList.Add(v)
	}

	if arrList.Size() != 9 {
		t.Log("size want: 9, but: ", arrList.Size())
		t.FailNow()
	}

	ok, err := arrList.Remove(5)
	if err != nil {
		t.Log("remove interface err: ", err)
		t.FailNow()
	}

	if arrList.Size() != 8 {
		t.Log("size want: 8, but: ", arrList.Size())
		t.FailNow()
	}

	if !ok {
		t.Log("delete 5, want true, but: ", ok)
		t.FailNow()
	}

	if idx := arrList.Index(5); idx != -1 {
		t.Log("5 index, want -1, but: ", idx)
		t.FailNow()
	}

	ok, err = arrList.AddIdx(4, 5)
	if !ok || err != nil {
		t.Log("AddIdx 4->5 failed", ok, err)
		t.FailNow()
	}

	if arrList.Size() != 9 {
		t.Log("size want: 9, but: ", arrList.Size())
		t.FailNow()
	}

	ok, err = arrList.AddIdx(4, 6)
	if !ok || err != nil {
		t.Log("AddIdx failed", ok, err)
		t.FailNow()
	}

	if arrList.Size() != 10 {
		t.Log("size want: 10, but: ", arrList.Size())
		t.FailNow()
	}

	if idx := arrList.Index(5); idx != 5 {
		t.Log("5 index, want 5, but: ", idx)
		t.FailNow()
	}

	ok, err = arrList.AddIdx(4, 6)
	if !ok || err != nil {
		t.Log("AddIdx failed", ok, err)
		t.FailNow()
	}

	if arrList.Size() != 11 {
		t.Log("size want: 11, but: ", arrList.Size())
		t.FailNow()
	}

	if idx := arrList.Index(1); idx != 0 {
		t.Log("9 index, want 9, but: ", idx)
		t.FailNow()
	}

	arrList.Trim()
	if arrList.Size() != arrList.Capacity() {
		t.Log("capacity != size")
		t.FailNow()
	}
}
