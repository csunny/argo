package stack

type Item interface {
}

// ItemStack the stack of items
type ItemStack struct {
	items []Item
}

// New Create a new ItemStack
func (s *ItemStack) New() *ItemStack {
	s.items = []Item{}
	return s
}

// Push adds an Item to the top of the stack
func (s *ItemStack) Push(t Item) {
	s.items = append(s.items, t)
}

// Pop removes an Item from the top of the stack
func (s *ItemStack) Pop() *Item {
	item := s.items[len(s.items)-1] // 后进先出
	s.items = s.items[0:len(s.items)-1]
	return &item

}
