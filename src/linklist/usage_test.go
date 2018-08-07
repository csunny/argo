package linklist

import "testing"

func TestLinkNode_NewJosphuseRing(t *testing.T) {
	head := LinkNode{}

	itertor := head.NewJosphuseRing(5)
	t.Error(itertor)
	if itertor != nil {
		t.Error(itertor.Payload)
		itertor = itertor.Next
	}


}
