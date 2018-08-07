package queue

import "testing"

var q ItemQueue

func initQueue() *ItemQueue  {
	if q.items == nil{
		q = ItemQueue{}
		q.New()
	}

	return &q
}

func TestItemQueue_Enqueue(t *testing.T) {
	q := initQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	if size := q.Size(); size != 3{
		t.Errorf("wrong count, the correct count is 3 but got %d", size)
	}
}

func TestItemQueue_Dequeue(t *testing.T) {
	q.Dequeue()

	if size := q.Size(); size != 2{
		t.Errorf("test failed, the corrected value is 2, but got %d", size)
	}

	q.Dequeue()
	q.Dequeue()
	if !q.IsEmpty(){
		t.Errorf("the queue should be empty.")
	}
}