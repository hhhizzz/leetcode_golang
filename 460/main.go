package _460

type ListNode struct {
	data interface{}
	next *ListNode
	last *ListNode
}

type LinkedList struct {
	head *ListNode
	tail *ListNode
}

func (list *LinkedList) addNode(newNode *ListNode) {
	if list.head == nil {
		//the list is empty
		list.head, list.tail = newNode, newNode
	} else {
		newNode.next, list.head.last = list.head, newNode
		list.head = newNode
	}
}

func (list *LinkedList) add(data interface{}) {
	newNode := &ListNode{data: data}
	list.addNode(newNode)
}

func (list *LinkedList) deleteNode(node *ListNode) {
	if list.head == list.tail {
		// only has one node
		list.head, list.tail = nil, nil
	} else if list.head == node {
		list.head = list.head.next
	} else if list.tail == node {
		list.delete()
	} else {
		node.last.next, node.next.last = node.next, node.last
	}
	node.last = nil
	node.next = nil
}

func (list *LinkedList) delete() *ListNode {
	if list.tail == nil {
		return nil
	} else {
		result := list.tail
		list.tail = result.last
		result.last = nil
		if list.tail == nil {
			//the list has only one node
			list.head = nil
		} else {
			list.tail.next = nil
		}
		return result
	}
}

type CountData struct {
	count     int
	valueList *LinkedList
}

type ValueData struct {
	key       int
	value     int
	countData *ListNode
}

type LFUCache struct {
	capacity  interface{}
	table     map[int]*ListNode
	countList *LinkedList
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		capacity:  capacity,
		table:     map[int]*ListNode{},
		countList: &LinkedList{},
	}
}

func (this *LFUCache) update(valueDataNode *ListNode) {
	valueData := valueDataNode.data.(*ValueData)
	countDataNode := valueData.countData
	countData := countDataNode.data.(*CountData)

	valueList := countData.valueList
	if countDataNode.last != nil && countDataNode.last.data.(*CountData).count == countData.count+1 {
		lastCountData := countDataNode.last.data.(*CountData)

		valueList.deleteNode(valueDataNode)
		lastCountData.valueList.addNode(valueDataNode)
		valueData.countData = countDataNode.last

		if valueList.head == nil {
			this.countList.deleteNode(countDataNode)
		}
	} else {
		if valueList.head == valueList.tail {
			countData.count += 1
		} else {
			valueList.deleteNode(valueDataNode)

			lastCountData := CountData{countData.count + 1, &LinkedList{}}
			lastCountData.valueList.addNode(valueDataNode)
			lastCountDataNode := &ListNode{data: &lastCountData, next: countDataNode, last: countDataNode.last}

			if this.countList.head == countDataNode {
				this.countList.head = lastCountDataNode
			} else {
				countDataNode.last.next = lastCountDataNode
			}
			countDataNode.last = lastCountDataNode
			valueData.countData = lastCountDataNode
		}
	}
}

func (this *LFUCache) toString() {
	//current := this.countList.head
	//for current != nil {
	//   fmt.Printf("%d: ", current.data.(*CountData).count)
	//   currentValue := current.data.(*CountData).valueList.head
	//   for currentValue != nil {
	//       fmt.Printf("%d:%d, ", currentValue.data.(*ValueData).key, currentValue.data.(*ValueData).value)
	//       currentValue = currentValue.next
	//   }
	//   fmt.Println()
	//   current = current.next
	//}
	//fmt.Println("************")
}

func (this *LFUCache) Get(key int) int {
	if valueDataNode, ok := this.table[key]; !ok {
		//fmt.Printf("cache.Get(%d)=-1\n", key)
		return -1
	} else {
		this.update(valueDataNode)
		//fmt.Printf("cache.Get(%d)=%d\n", key, valueDataNode.data.(*ValueData).value)
		this.toString()
		return valueDataNode.data.(*ValueData).value
	}
}

func (this *LFUCache) Put(key int, value int) {
	if this.capacity == 0 {
		return
	}

	if oldValueNode, ok := this.table[key]; ok {
		this.update(oldValueNode)
		oldValue := oldValueNode.data.(*ValueData)
		oldValue.value = value

		//fmt.Printf("cache.Put(%d, %d)\n", key, value)
		this.toString()
		return
	}

	if this.countList.tail == nil {
		newCountData := CountData{count: 0, valueList: &LinkedList{}}
		this.countList.add(&newCountData)
	}

	lastReviewCount := this.countList.tail.data.(*CountData)
	lastReviewValueList := lastReviewCount.valueList

	if len(this.table) == this.capacity {
		//delete the last reviewed data
		lastReviewValue := lastReviewValueList.delete().data.(*ValueData)
		delete(this.table, lastReviewValue.key)
		//fmt.Printf("envict the key: %d\n", lastReviewValue.key)

		if lastReviewValueList.head == nil {
			//the countNode is empty, delete the countNode
			this.countList.delete()
		}
		if this.countList.tail != nil {
			lastReviewCount = this.countList.tail.data.(*CountData)
			lastReviewValueList = lastReviewCount.valueList
		}

	}

	if this.countList.tail == nil {
		newCountData := CountData{count: 0, valueList: &LinkedList{}}
		newCountDataNode := ListNode{data: &newCountData}
		this.countList.head, this.countList.tail = &newCountDataNode, &newCountDataNode
	} else if lastReviewCount.count != 0 {
		newCountData := CountData{count: 0, valueList: &LinkedList{}}
		newCountDataNode := ListNode{data: &newCountData, last: this.countList.tail}
		this.countList.tail.next = &newCountDataNode
		this.countList.tail = &newCountDataNode
	}

	//update the lastReview data
	lastReviewCount = this.countList.tail.data.(*CountData)
	lastReviewValueList = lastReviewCount.valueList

	newValueData := ValueData{
		key:       key,
		value:     value,
		countData: this.countList.tail,
	}

	lastReviewValueList.add(&newValueData)
	this.table[key] = lastReviewValueList.head

	//fmt.Printf("cache.Put(%d, %d)\n", key, value)
	this.toString()
}
