package _460

import "testing"

func TestLFUCache(t *testing.T) {

	//cache := Constructor(3)
	//cache.Put(2, 2)
	//cache.Put(1, 1)
	//cache.Get(2)
	//cache.Get(1)
	//cache.Get(2)
	//cache.Put(3, 3)
	//cache.Put(4, 4)
	//cache.Get(3)
	//cache.Get(2)
	//cache.Get(1)
	//cache.Get(4)

	//cache := Constructor(2)
	//cache.Put(1, 1);
	//cache.Put(2, 2);
	//cache.Get(1);    // returns 1
	//cache.Put(3, 3); // evicts key 2
	//cache.Get(2);    // returns -1 (not found)
	//cache.Get(3);    // returns 3.
	//cache.Put(4, 4); // evicts key 1.
	//cache.Get(1);    // returns -1 (not found)
	//cache.Get(3);    // returns 3
	//cache.Get(4);    // returns 4

	cache := Constructor(1)
	cache.Put(2, 1)
	cache.Get(2)
	cache.Put(3, 2)
	cache.Get(2)
	cache.Get(3)

	//cache := Constructor(2)
	//cache.Put(2, 6)
	//cache.Get(1)
	//cache.Put(1, 5)
	//cache.Put(1, 2)
	//cache.Get(1)
	//cache.Get(2)

	//cache := Constructor(2)
	//cache.Put(2, 1)
	//cache.Put(3, 2)
	//cache.Get(3)
	//cache.Get(2)
	//cache.Put(4, 3)
	//cache.Get(2)
	//cache.Get(3)
	//cache.Get(4)

	//cache := Constructor(3)
	//cache.Put(1, 1)
	//cache.Put(2, 2)
	//cache.Put(3, 3)
	//cache.Put(4, 4)
	//cache.Get(4)
	//cache.Get(3)
	//cache.Get(2)
	//cache.Get(1)
	//cache.Put(5, 5)
	//
	//for i := 1; i <= 5; i++ {
	//    cache.Get(i)
	//}

}

func TestLFUCache2(t *testing.T) {
	cache := Constructor(10)
	cache.Put(10, 13)
	cache.Put(3, 17)
	cache.Put(6, 11)
	cache.Put(10, 5)
	cache.Put(9, 10)
	cache.Get(13)
	cache.Put(2, 19)
	cache.Get(2)
	cache.Get(3)
	cache.Put(5, 25)
	cache.Get(8)
	cache.Put(9, 22)
	cache.Put(5, 5)
	cache.Put(1, 30)
	cache.Get(11)
	cache.Put(9, 12)
	cache.Get(7)
	cache.Get(5)
	cache.Get(8)
	cache.Get(9)
	cache.Put(4, 30)
	cache.Put(9, 3)
	cache.Get(9)
	cache.Get(10)
	cache.Get(10)
	cache.Put(6, 14)
	cache.Put(3, 1)
	cache.Get(3)
	cache.Put(10, 11)
	cache.Get(8)
	cache.Put(2, 14)
	cache.Get(1)
	cache.Get(5)
	cache.Get(4)
	cache.Put(11, 4)
	cache.Put(12, 24)
	cache.Put(5, 18)
	cache.Get(13)
	cache.Put(7, 23)
	cache.Get(8)
	cache.Get(12)
	cache.Put(3, 27)
	cache.Put(2, 12)
	cache.Get(5)
	cache.Put(2, 9)
	cache.Put(13, 4)
	cache.Put(8, 18)
	cache.Put(1, 7)
	cache.Get(6)
	cache.Put(9, 29)
	cache.Put(8, 21)
	cache.Get(5)
	cache.Put(6, 30)
	cache.Put(1, 12)
	cache.Get(10)
	cache.Put(4, 15)
	cache.Put(7, 22)
	cache.Put(11, 26)
	cache.Put(8, 17)
	cache.Put(9, 29)
	cache.Get(5)
	cache.Put(3, 4)
	cache.Put(11, 30)
	cache.Get(12)
	cache.Put(4, 29)
	cache.Get(3)
	cache.Get(9)
	cache.Get(6)
	cache.Put(3, 4)
	cache.Get(1)
	cache.Get(10)
	cache.Put(3, 29)
	cache.Put(10, 28)
	cache.Put(1, 20)
	cache.Put(11, 13)
	cache.Get(3)
	cache.Put(3, 12)
	cache.Put(3, 8)
	cache.Put(10, 9)
	cache.Put(3, 26)
	cache.Get(8)
	cache.Get(7)
	cache.Get(5)
	cache.Put(13, 17)
	cache.Put(2, 27)
	cache.Put(11, 15)
	cache.Get(12)
	cache.Put(9, 19)
	cache.Put(2, 15)
	cache.Put(3, 16)
	cache.Get(1)
	cache.Put(12, 17)
	cache.Put(9, 1)
	cache.Put(6, 19)
	cache.Get(4)
	cache.Get(5)
	cache.Get(5)
	cache.Put(8, 1)
	cache.Put(11, 7)
	cache.Put(5, 2)
}

func TestUpdate1(t *testing.T) {
	cache := Constructor(2)
	cache.Put(1, 1)
	for i := 0; i < 10; i++ {
		cache.Get(1)
	}
	countData := cache.countList.head.data.(*CountData)
	if countData.count != 10 {
		t.Errorf("the count data need to be 10")
	}
	if countData.valueList.head.data.(*ValueData).key != 1 {
		t.Errorf("the value data need to be 1")
	}
	if countData.valueList.head != countData.valueList.tail {
		t.Errorf("there must be one valueData in the count list")
	}
}

func TestUpdate2(t *testing.T) {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)

	cache.Get(1)

	countData := cache.countList.tail.data.(*CountData)
	if countData.count != 0 {
		t.Errorf("the count need to be 0")
	}
	valueData := countData.valueList.head.data.(*ValueData)
	if valueData.key != 2 {
		t.Errorf("the value data need to be 2")
	}

	if countData.valueList.head != countData.valueList.tail {
		t.Errorf("there must be one valueData in the count list")
	}

	cache.Get(2)
	countData = cache.countList.tail.data.(*CountData)
	if cache.countList.head != cache.countList.tail {
		t.Errorf("there must be only one countData in the list")
	}
	if countData.count != 1 {
		t.Errorf("the count need to be 1")
	}
	if countData.valueList.head == countData.valueList.tail {
		t.Errorf("there must be more than valueData in the count list")
	}

}

func TestUpdate3(t *testing.T) {
	cache := Constructor(3)

	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)

	cache.Get(3)
	cache.Get(3)
	cache.Get(2)

	current := cache.countList.head
	i := 2
	for current != nil {
		if current.data.(*CountData).count != i {
			t.Errorf("the count need to be %d\n", i)
		}
		current = current.next
		i--
	}

	cache.Get(2)

	valueList := cache.countList.head.data.(*CountData).valueList
	if valueList.head.data.(*ValueData).key != 2 {
		t.Errorf("the value need to be 2")
	}
	if valueList.head.next.data.(*ValueData).key != 3 {
		t.Errorf("the value need to be 3")
	}
	if cache.countList.tail.data.(*CountData).valueList.head.data.(*ValueData).key != 1 {
		t.Errorf("the value need to be 1")
	}

	cache.Get(1)
	valueList = cache.countList.tail.data.(*CountData).valueList
	if valueList.head.data.(*ValueData).key != 1 {
		t.Errorf("the value need to be 1")
	}

	cache.Get(1)
	valueList = cache.countList.head.data.(*CountData).valueList
	if valueList.head.data.(*ValueData).key != 1 {
		t.Errorf("the value need to be 1")
	}
	if valueList.head.next.data.(*ValueData).key != 2 {
		t.Errorf("the value need to be 2")
	}
	if valueList.tail.data.(*ValueData).key != 3 {
		t.Errorf("the value need to be 3")
	}

}

func TestList(t *testing.T) {
	list := LinkedList{}
	data := ValueData{value: 1}
	list.add(data)
	list.delete()
	if list.head != nil || list.tail != nil {
		t.Errorf("expected nil list")
	}
	data = ValueData{value: 1}
	list.add(data)
	data = ValueData{value: 2}
	list.add(data)
	list.delete()
	if list.head != list.tail {
		t.Errorf("expected head and tail to be pointed to same address")
	}
	if v, ok := list.head.data.(ValueData); ok {
		if v.value != 2 {
			t.Errorf("deleted the wrong node")
		}
	} else {
		t.Errorf("error casting data to ValueData")
	}
}
