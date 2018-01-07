package bst

import (
	"testing"
)

func TestBST_Insert(t *testing.T) {
	type fields struct {
		root  *node
		count int
	}
	type args struct {
		item int
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      bool
		wantCount int
		wantState string
	}{
		{"Fails on duplicate", fields{&node{5, nil, nil}, 1}, args{5}, false, 1, "5"},
		{"Inserts on empy", fields{nil, 0}, args{5}, true, 1, "5"},
		{"Inserts on left", fields{&node{5, &node{3, nil, nil}, nil}, 2}, args{1}, true, 3, "1 3 5"},
		{"Inserts on right", fields{&node{5, &node{3, nil, nil}, &node{7, nil, nil}}, 3}, args{8}, true, 4, "3 5 7 8"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bst := &BST{
				root:  tt.fields.root,
				count: tt.fields.count,
			}
			if got := bst.Insert(tt.args.item); got != tt.want {
				t.Errorf("BST.Insert() = %v, want %v", got, tt.want)
			}

			if bst.count != tt.wantCount {
				t.Errorf("BST.Insert() = %v, want count %v", bst.count, tt.wantCount)
			}

			if bst.statePrint() != tt.wantState {
				t.Errorf("BST.Insert() = %v, want state %v", bst.statePrint(), tt.wantState)
			}
		})
	}
}

func TestBST_Exists(t *testing.T) {
	type fields struct {
		root  *node
		count int
	}
	type args struct {
		item int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"Empty bst", fields{nil, 0}, args{5}, false},
		{"Not exists", fields{&node{5, &node{3, nil, nil}, &node{7, &node{6, nil, nil}, nil}}, 4}, args{8}, false},
		{"Exists", fields{&node{5, &node{3, nil, nil}, &node{7, &node{6, nil, nil}, nil}}, 4}, args{7}, true},
		{"Exists", fields{&node{5, &node{3, nil, nil}, &node{7, &node{6, nil, nil}, nil}}, 4}, args{6}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bst := &BST{
				root:  tt.fields.root,
				count: tt.fields.count,
			}
			if got := bst.Exists(tt.args.item); got != tt.want {
				t.Errorf("BST.Exists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBST_Remove(t *testing.T) {
	type fields struct {
		root  *node
		count int
	}
	type args struct {
		item int
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      bool
		wantCount int
		wantState string
	}{
		{"Remove from empty tree", fields{nil, 0}, args{5}, false, 0, ""},
		{"Remove root, replace left", fields{&node{5, &node{3, nil, nil}, nil}, 2}, args{5}, true, 1, "3"},
		{"Remove root, replace right", fields{&node{5, nil, &node{7, nil, nil}}, 2}, args{5}, true, 1, "7"},
		{"Remove root, replace none", fields{&node{5, nil, nil}, 1}, args{5}, true, 0, ""},
		{"Remove non root, replace right", fields{&node{5, &node{3, nil, nil}, &node{7, &node{6, nil, nil}, nil}}, 4}, args{7}, true, 3, "3 5 6"},
		{"Remove non root, complicated1", fields{&node{6, &node{0, &node{-1, nil, nil}, &node{3, &node{2, &node{1, nil, nil}, nil}, &node{4, nil, &node{5, nil, nil}}}}, nil}, 8}, args{0}, true, 7, "-1 1 2 3 4 5 6"},
		{"Remove non root, complicated2", fields{&node{6, &node{0, &node{-1, nil, nil}, &node{3, &node{2, &node{1, nil, nil}, nil}, &node{4, nil, &node{5, nil, nil}}}}, nil}, 8}, args{3}, true, 7, "-1 0 1 2 4 5 6"},
		{"Remove non root, complicated3", fields{&node{6, &node{0, &node{-1, nil, nil}, &node{3, nil, &node{4, nil, &node{5, nil, nil}}}}, nil}, 6}, args{3}, true, 5, "-1 0 4 5 6"},
		{"Remove non root, complicated4", fields{&node{6, &node{0, nil, &node{3, &node{2, &node{1, nil, nil}, nil}, &node{4, nil, &node{5, nil, nil}}}}, nil}, 7}, args{0}, true, 6, "1 2 3 4 5 6"},
		{"Remove non root, complicated5", fields{&node{6, &node{0, &node{-1, nil, nil}, &node{3, &node{2, &node{1, nil, nil}, nil}, &node{4, nil, &node{5, nil, nil}}}}, nil}, 8}, args{5}, true, 7, "-1 0 1 2 3 4 6"},
		{"Remove non root, complicated6", fields{&node{6, &node{0, &node{-1, nil, nil}, &node{3, &node{2, &node{1, nil, nil}, nil}, &node{4, nil, &node{5, nil, nil}}}}, nil}, 8}, args{1}, true, 7, "-1 0 2 3 4 5 6"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bst := &BST{
				root:  tt.fields.root,
				count: tt.fields.count,
			}
			if got := bst.Remove(tt.args.item); got != tt.want {
				t.Errorf("BST.Remove() = %v, want %v", got, tt.want)
			}

			if bst.count != tt.wantCount {
				t.Errorf("BST.Insert() = %v, want count %v", bst.count, tt.wantCount)
			}

			if bst.statePrint() != tt.wantState {
				t.Errorf("BST.Insert() = %v, want state %v", bst.statePrint(), tt.wantState)
			}
		})
	}
}
