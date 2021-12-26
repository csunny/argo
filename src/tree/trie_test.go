package tree

import "testing"

func TestNewTrie(t *testing.T) {
	type args struct {
		rootHash []byte
	}

	storage, _ := NewMemoryStorage()
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			"nil",
			args{nil},
			nil,
			false,
		},
		{
			"root",
			args{[]byte("hello")},
			nil,
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTrie(tt.args.rootHash, storage, false)
			if err != nil {
				t.Errorf("New trieNode error, error = %v", err)
			}
			if got == nil {
				return
			}

			t.Log(got)
		})
	}
}
