package list

import "testing"

func TestSkipList(t *testing.T) {
	skipList := NewSkipList(3, 0.5)

	sample := []int{3, 6, 7, 9, 12, 19, 17, 26, 21, 25}

	for _, v := range sample {
		skipList.Insert(v)
	}

	skipList.Print()

	skipList.Search(19)

	skipList.Delete(19)
	skipList.Print()
}
