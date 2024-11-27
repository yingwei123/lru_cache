package doubly_linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoublyLinkedList(t *testing.T) {
	t.Run("test add node to empty list", testAddNodeEmptyList)
	t.Run("test add node to non empty list", testAddNodeToNonEmptyList)
	t.Run("test remove node", testRemoveNode)
	t.Run("test remove tail", testRemoveTail)
	t.Run("test remove head", testRemoveHead)
}

func testAddNodeEmptyList(t *testing.T) {
	dll := CreateDoublyLinkedList()
	dll.AddToHead(1, "1")

	assert.Equal(t, 1, dll.Head.Value)
	assert.Equal(t, 1, dll.Tail.Value)
}

func testAddNodeToNonEmptyList(t *testing.T) {
	dll := CreateDoublyLinkedList()
	dll.AddToHead(1, "1")
	dll.AddToHead(2, "2")

	assert.Equal(t, 2, dll.Head.Value)
	assert.Equal(t, 1, dll.Tail.Value)
}

func testRemoveNode(t *testing.T) {
	dll := CreateDoublyLinkedList()
	dll.AddToHead(1, "1")
	dll.AddToHead(2, "2")

	prevHead := dll.RemoveNode(dll.Head)

	assert.Equal(t, 1, dll.Head.Value)
	assert.Equal(t, 2, prevHead.Value)
}

func testRemoveTail(t *testing.T) {
	dll := CreateDoublyLinkedList()
	dll.AddToHead(1, "1")
	dll.AddToHead(2, "2")

	prevTail := dll.RemoveTail()

	assert.Equal(t, 1, prevTail.Value)
	assert.Equal(t, 2, dll.Head.Value)
}

func testRemoveHead(t *testing.T) {
	dll := CreateDoublyLinkedList()
	dll.AddToHead(1, "1")
	dll.AddToHead(2, "2")

	prevHead := dll.RemoveHead()

	assert.Equal(t, 2, prevHead.Value)
	assert.Equal(t, 1, dll.Head.Value)
}
