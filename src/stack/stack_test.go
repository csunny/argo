package stack

import "testing"

var s ItemStack

func initStack() *ItemStack  {
	if s.items == nil{
		s = ItemStack{}
		s.New()
	}

	return &s
}

func TestItemStack_Push(t *testing.T) {
	s := initStack()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	if size := len(s.items); size != 3{
		t.Errorf("test failed ")
	}
}


func TestItemStack_Pop(t *testing.T) {
	s.Pop()

	if size := len(s.items); size !=2 {
		t.Errorf("test failed, excepted 2 and got %d", size)
	}

	s.Pop()
	s.Pop()

	if size := len(s.items); size != 0{
		t.Errorf("test failed, excepted 0 and got %d", size)
	}
}
