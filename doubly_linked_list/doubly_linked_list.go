package doubly_linked_list

type Node struct {
	Value any
	Key   string
	Next  *Node
	Prev  *Node
}

type DoublyLinkedList struct {
	Head *Node
	Tail *Node
}

func CreateDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{
		Head: nil,
		Tail: nil,
	}
}

func (dll *DoublyLinkedList) RemoveNode(node *Node) *Node {
	nodePrev := node.Prev
	nodeNext := node.Next

	if nodePrev != nil {
		nodePrev.Next = nodeNext
	} else {
		dll.Head = nodeNext
	}

	if nodeNext != nil {
		nodeNext.Prev = nodePrev
	} else {
		dll.Tail = nodePrev
	}

	return node
}

func (dll *DoublyLinkedList) RemoveTail() *Node {
	prevTail := dll.Tail

	if dll.Tail == nil {
		return nil
	}

	if dll.Tail.Prev != nil {
		dll.Tail.Prev.Next = nil
	} else {
		dll.Head = nil
	}

	dll.Tail = dll.Tail.Prev

	return prevTail
}

func (dll *DoublyLinkedList) RemoveHead() *Node {
	prevHead := dll.Head

	if dll.Head == nil {
		return nil
	}

	if dll.Head.Next != nil {
		dll.Head = dll.Head.Next
		dll.Head.Prev = nil
	} else {
		dll.Head = nil
		dll.Tail = nil
	}

	return prevHead
}

func (dll *DoublyLinkedList) AddToHead(value any, key string) {
	newNode := &Node{
		Value: value,
		Key:   key,
		Next:  dll.Head,
		Prev:  nil,
	}

	if dll.Head != nil {
		dll.Head.Prev = newNode
	} else {
		dll.Tail = newNode
	}

	dll.Head = newNode
}
