package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v any) *ListItem
	PushBack(v any) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value any
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	front *ListItem
	back  *ListItem
	len   int
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Back() *ListItem {
	if l.len == 0 {
		return nil
	}

	return l.back
}

func (l *list) Front() *ListItem {
	if l.len == 0 {
		return nil
	}
	return l.front
}
func (l *list) PushFront(v any) *ListItem {
	l.len++

	var newListItem ListItem

	newListItem.Value = v
	newListItem.Prev = l.front
	newListItem.Next = nil

	if l.front != nil {
		l.front.Next = &newListItem
	}

	if l.back == nil {
		l.back = &newListItem
	}

	l.front = &newListItem

	return &newListItem
}

func (l *list) PushBack(v any) *ListItem {
	l.len++

	var newListItem ListItem

	newListItem.Value = v
	newListItem.Prev = nil
	newListItem.Next = l.back

	if l.back != nil {
		l.back.Prev = &newListItem
	}

	if l.front == nil {
		l.front = &newListItem
	}
	l.back = &newListItem

	return &newListItem
}

func (l *list) Remove(i *ListItem) {
	l.len--
	prevEl := i.Prev
	nextEl := i.Next
	prevEl.Next = nextEl
	nextEl.Prev = prevEl
}

func (l *list) MoveToFront(i *ListItem) {
}

func NewList() List {
	return new(list)
}
