package _707

import (
	"fmt"
	"testing"
)

func TestMyLinkedList(t *testing.T) {
	list := Constructor()
	list.AddAtIndex(-1, 0)

	if list.Get(-1) != -1 {
		t.Errorf("Wrong at 1")
	}
	list.DeleteAtIndex(-1)

	list = Constructor()
	list.AddAtHead(1)
	list.AddAtTail(3)
	list.AddAtIndex(1, 2)
	if list.Get(1) != 2 {
		t.Errorf("Wrong at 2")
	}
	list.DeleteAtIndex(1)
	if list.Get(1) != 3 {
		t.Errorf("Wrong at 2")
	}

	list = Constructor()
	list.AddAtHead(1)
	list.AddAtTail(3)
	list.AddAtIndex(1, 2)
	if list.Get(1) != 2 {
		t.Errorf("Wrong at 3")
	}
	list.DeleteAtIndex(0)
	if list.Get(0) != 2 {
		t.Errorf("Wrong at 3")
	}

	list = Constructor()
	list.AddAtIndex(-1, 0)

	if list.Get(0) != 0 {
		t.Errorf("Wrong at 4")
	}
	list.DeleteAtIndex(0)

	list = Constructor()
	list.AddAtHead(1)
	list.AddAtTail(3)
	list.AddAtIndex(1, 2)
	if list.Get(-1) != -1 {
		t.Errorf("Wrong at 5")
	}
	list.DeleteAtIndex(1)
	if list.Get(-3) != -1 {
		t.Errorf("Wrong at 5")
	}

	list = Constructor()
	list.Get(0) //-1
	list.AddAtIndex(1, 2)
	list.Get(0) //-1
	list.Get(1) //-1
	list.AddAtIndex(0, 1)
	list.Get(0)              //1
	fmt.Println(list.Get(1)) //-1
}
