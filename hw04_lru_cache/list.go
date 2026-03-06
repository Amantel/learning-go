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
	newListItem.Prev = nil
	newListItem.Next = l.front

	if l.front != nil {
		l.front.Prev = &newListItem
	}

	l.front = &newListItem

	if l.back == nil {
		l.back = l.front
	}

	return &newListItem
}

func (l *list) PushBack(v any) *ListItem {
	l.len++

	var newListItem ListItem

	newListItem.Value = v
	newListItem.Prev = l.back
	newListItem.Next = nil

	if l.back != nil {
		l.back.Next = &newListItem
	}

	l.back = &newListItem

	if l.front == nil {
		l.front = l.back
	}

	return &newListItem
}

func (l *list) Remove(i *ListItem) {
	l.len--

	prevEl := i.Prev
	nextEl := i.Next
	if prevEl != nil {
		prevEl.Next = nextEl
	}
	if nextEl != nil {
		nextEl.Prev = prevEl
	}

	if i == l.front {
		l.front = nextEl
	}
	if i == l.back {
		l.back = prevEl
	}

}

func (l *list) MoveToFront(i *ListItem) {
	// соединяем там где будет дырка
	// if i.Next != nil {
	// 	i.Next.Prev = i.Prev
	// }
	// if i.Prev != nil {
	// 	i.Prev.Next = i.Next
	// }

	l.Remove(i)

	l.PushFront(i.Value)
}

func NewList() List {
	return new(list)
}
