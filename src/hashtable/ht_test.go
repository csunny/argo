package hashtable

import (
	"testing"
	"fmt"
)

func populateHashTable(count int, start int) *HashTable {
	dict := HashTable{}
	for i := start; i < (start + count); i++ {
		dict.Put(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))
	}
	return &dict
}

func TestPut(t *testing.T) {
	dict := populateHashTable(3, 0)
	if size := dict.Size(); size != 3 {
		t.Errorf("Test failed, expected 3 and got %d", size)
	}

	dict.Put("magic", "shuai") //should not add a new one, just change the existing one
	if size := dict.Size(); size != 4 {
		t.Errorf("wrong count, expected 4 and got %d", size)
	}

	dict.Put("magic1", "haha")
	if size := dict.Size(); size != 5 {
		t.Errorf("Test Failed, expected 5 and got %d", size)
	}

}
