package lists

import "testing"

func TestAdd(t *testing.T) {
	list := NewSinglyLinkedList()
	cur := list.Head
	if cur != nil {
		t.Error("Empty list is not nil!")
	}

	elem := &Element{Value: 1}
	list.Append(elem)
	cur = list.Head
	if cur != elem {
		t.Error("Wrong value, expected 1 got: ", cur.Value)
	}

	elem = &Element{Value: 2}
	list.Append(elem)
	cur = cur.Next()
	if cur != elem {
		t.Error("Wrong value got: ", cur.Value)
	}

	if list.Head == elem {
		t.Error("Wrong value got: ", cur.Value)
	}
}

func TestRemove(t *testing.T) {
	list := NewSinglyLinkedList()
	cur := list.Head
	if cur != nil {
		t.Error("Empty list is not nil!")
	}

	elem1 := &Element{Value: 1}
	elem2 := &Element{Value: 2}
	elem3 := &Element{Value: 3}
	list.Append(elem1)
	list.Append(elem2)
	list.Append(elem3)
	cur = list.Head
	if list.Head.Next().Next().Next() != nil {
		t.Error("Expected 3 nodes. Something went wrong.")
	}

	e := list.Remove(elem2)
	if e == nil {
		t.Error("Should have removed 2nd element but didnt!")
	}
	if list.Head != elem1 && list.Head.Next() != elem3 {
		t.Error("Removed element 2 but linked failed!")
	}
}

func TestInsertAfter(t *testing.T) {
	list := NewSinglyLinkedList()
	cur := list.Head
	if cur != nil {
		t.Error("Empty list is not nil!")
	}

	elem1 := &Element{Value: 1}
	elem2 := &Element{Value: 2}
	elem3 := &Element{Value: 3}
	elem4 := &Element{Value: 4}
	list.Append(elem1)
	list.Append(elem2)
	list.Append(elem4)

	cur = list.Head
	if list.Head.Next().Next() != elem4 {
		t.Error("Expected 4 nodes. Something went wrong.")
	}

	e := list.InsertAfter(elem3, elem2)
	if e == nil {
		t.Error("Should have removed 2nd element but didnt!")
	}
	if list.Head != elem1 && list.Head.Next().Next() != elem3 {
		t.Error("Added element 3 after 2. But didn't insert correctly!")
	}
}

func TestReverse(t *testing.T) {
	listForward := NewSinglyLinkedList()
	listBackwards := NewSinglyLinkedList()

	elem1 := &Element{Value: 1}
	elem2 := &Element{Value: 2}
	elem3 := &Element{Value: 3}
	elem4 := &Element{Value: 4}
	listForward.Append(&Element{Value: 1})
	listForward.Append(&Element{Value: 2})
	listForward.Append(&Element{Value: 3})
	listForward.Append(&Element{Value: 4})

	listBackwards.Append(elem4)
	listBackwards.Append(elem3)
	listBackwards.Append(elem2)
	listBackwards.Append(elem1)

	listReversed := listForward.Reverse()
	cur := listReversed.Head
	if cur.Value.(int) != 4 {
		t.Error("Should be 4!")
	}

	cur = cur.next
	if cur.Value.(int) != 3 {
		t.Error("Should be 3!")
	}

	cur = cur.next
	if cur.Value.(int) != 2 {
		t.Error("Should be 2!")
	}

	cur = cur.next
	if cur.Value.(int) != 1 {
		t.Error("Should be 1!")
	}

	cur = cur.next
	if cur != nil {
		t.Error("Wrong number of elements!")
	}
}

func TestMiddle(t *testing.T) {
	list := NewSinglyLinkedList()

	elem1 := &Element{Value: 1}
	elem2 := &Element{Value: 2}
	elem3 := &Element{Value: 3}
	elem4 := &Element{Value: 4}
	elem5 := &Element{Value: 4}
	list.Append(elem1)
	list.Append(elem2)
	list.Append(elem3)

	mid := list.Middle()
	if mid != elem2 {
		t.Error("Ops. Should have been elem2")
	}

	list.Append(elem4)
	mid = list.Middle()
	if mid != elem2 {
		t.Error("Ops. Should have been elem2")
	}

	list.Append(elem5)
	mid = list.Middle()
	if mid != elem3 {
		t.Error("Ops. Should have been elem3")
	}
}