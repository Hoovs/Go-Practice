package lists

import "sync"

type Element struct {
	Value interface{}
	next  *Element
}

func (e *Element) Next() *Element {
	return e.next
}

type SinglyLinkedList struct {
	Head *Element
	size int
	lock sync.Mutex
}

func NewSinglyLinkedList() *SinglyLinkedList{
	list := new(SinglyLinkedList)
	list.Head = nil
	list.size = 0
	list.lock = sync.Mutex{}
	return list
}

func (l *SinglyLinkedList) InsertAfter(v interface{}, e *Element) *Element {
	l.lock.Lock()
	defer l.lock.Unlock()

	cur := l.Head
	for cur != e && cur.next != nil {
		cur = cur.next
	}

	if cur != e {
		return nil
	}

	tmp := cur.next
	cur.next = e
	e.next = tmp

	return e
}

func (l *SinglyLinkedList) Size() int {
	return l.size
}

func (l *SinglyLinkedList) Append(e *Element) {
	l.lock.Lock()
	defer l.lock.Unlock()

	if l.Head == nil {
		l.Head = e
	} else {
		cur := l.Head
		for cur.next != nil {
			cur = cur.next
		}
		cur.next = e
	}

	l.size++
}

func (l *SinglyLinkedList) Remove(e *Element) *Element {
	l.lock.Lock()
	defer l.lock.Unlock()

	if l.size != 0 {
		cur := l.Head
		var prev *Element
		for cur != e && cur.next != nil {
			prev = cur
			cur = cur.next
		}

		if cur != e {
			return nil
		}

		prev.next = cur.next
		cur.next = nil

		return cur
	}
	return nil
}

func (l *SinglyLinkedList) Reverse() *SinglyLinkedList {
	l.lock.Lock()
	defer l.lock.Unlock()

	cur := l.Head

	var next, prev *Element

	for cur != nil {
		next = cur.next
		cur.next = prev
		prev = cur
		cur = next
	}

	l.Head = prev

	return l
}

func (l *SinglyLinkedList) Middle() *Element {
	c1, c2 := l.Head, l.Head

	for c1.next != nil && c1.next.next != nil {
		c1 = c1.next.next
		c2 = c2.next
	}

	return c2
}