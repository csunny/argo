package queue

type Item interface {

}


// Item the type of the queue
type ItemQueue struct {
	items []Item
}

type ItemQueuer interface {
	New() ItemQueue
	Enqueue(t Item)
	Dequeue() Item
	IsEmpty() bool
	Size() int
}

// New creates a new ItemQueue
func (s *ItemQueue) New() *ItemQueue{
	s.items = []Item{}
	return s
}