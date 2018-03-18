package main

import (
	"fmt"
)

type MyList struct {
	data string
	next *MyList
}

func (list *MyList) add(l MyList) {
	l.next = list.next
	list.next = &l
}

func (list *MyList) append(l MyList) {
	for {
		if list.next == nil {
			list.next = &l
			break
		}

		list = list.next
	}
}

func (list *MyList) show() {
	for list.next != nil {
		fmt.Println(list.data)

		list = list.next
	}
}

func main() {
	myList := MyList{data: "1"}
	myList.append(MyList{data: "2"})
	myList.append(MyList{data: "3"})
	myList.append(MyList{data: "4"})
	myList.append(MyList{data: "5"})
	myList.append(MyList{data: "6"})

	myList.show()
}
