package ds

import (
	"fmt"
)

type MyList struct {
	data string
	next *MyList
}

func (list *MyList) Add(l MyList) {
	l.next = list.next
	list.next = &l
}

func (list *MyList) Append(l MyList) {
	for {
		if list.next == nil {
			list.next = &l
			break
		}

		list = list.next
	}
}

func (list *MyList) Show() {
	for list.next != nil {
		fmt.Println(list.data)

		list = list.next
	}
}

func (list *MyList) Length() int {
	length := 0

	for list.next != nil {
		length++
	}

	return length
}
