package ds

import (
	"testing"
)

func TestAdd(t *testing.T) {
	list := MyList{data: "1"}
	list.Add(MyList{data: "2"})
	list.Add(MyList{data: "3"})

	if list.Length() != 3 {
		t.Errorf("list.Length() == %v, want %v", list.Length(), 3)
	}
}
