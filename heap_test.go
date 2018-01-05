package heap

import (
	"reflect"
	"testing"
)

func testEq(a, b []int) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestHeap_Insert(t *testing.T) {
	type fields struct {
		h     []int
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
		wantState []int
	}{
		{"Add first item", fields{[]int{0, 0, 0, 0, 0}, 0}, args{5}, true, []int{5, 0, 0, 0, 0}},
		{"No percolate", fields{[]int{5, 0, 0, 0, 0}, 1}, args{6}, true, []int{5, 6, 0, 0, 0}},
		{"Percolate one layer", fields{[]int{5, 6, 0, 0, 0}, 2}, args{4}, true, []int{4, 6, 5, 0, 0}},
		{"Percolate two layers", fields{[]int{4, 6, 5, 0, 0}, 3}, args{2}, true, []int{2, 4, 5, 6, 0}},
		{"Percolate to middle", fields{[]int{2, 4, 5, 6, 0}, 4}, args{3}, true, []int{2, 3, 5, 6, 4}},
		{"No space left", fields{[]int{2, 3, 5, 6, 4}, 5}, args{10}, false, []int{2, 3, 5, 6, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heap := &Heap{
				h:     tt.fields.h,
				count: tt.fields.count,
			}
			if got := heap.Insert(tt.args.item); got != tt.want {
				t.Errorf("Heap.Insert() = %v, want %v", got, tt.want)
			}

			if !testEq(heap.h, tt.wantState) {
				t.Errorf("Heap.Insert() = %v, want state %v", heap.h, tt.wantState)
			}
		})
	}
}

func TestHeap_RemoveMin(t *testing.T) {
	type fields struct {
		h     []int
		count int
	}
	tests := []struct {
		name      string
		fields    fields
		want      int
		want1     bool
		wantState []int
	}{
		{"Nothing to remove", fields{[]int{0, 0, 0, 0, 0}, 0}, 0, false, []int{0, 0, 0, 0, 0}},
		{"Remove left", fields{[]int{2, 3, 5, 6, 4}, 5}, 2, true, []int{3, 4, 5, 6, 4}},
		{"Remove right", fields{[]int{2, 5, 4, 6, 7, 5, 3}, 7}, 2, true, []int{3, 5, 4, 6, 7, 5, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heap := &Heap{
				h:     tt.fields.h,
				count: tt.fields.count,
			}
			got, got1 := heap.RemoveMin()
			if got != tt.want {
				t.Errorf("Heap.RemoveMin() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Heap.RemoveMin() got1 = %v, want %v", got1, tt.want1)
			}
			if !testEq(heap.h, tt.wantState) {
				t.Errorf("Heap.RemoveMin() got = %v, want state %v", heap.h, tt.wantState)
			}
		})
	}
}

func TestSort(t *testing.T) {
	type args struct {
		orig []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Sort already sorted slice", args{[]int{1, 2, 3, 4, 5, 6, 7}}, []int{1, 2, 3, 4, 5, 6, 7}},
		{"Sort empty slice", args{[]int{}}, []int{}},
		{"Sort with duplicates", args{[]int{1, 1, 3, 3, 2, 2}}, []int{1, 1, 2, 2, 3, 3}},
		{"Normal sort", args{[]int{3, 6, 2, 6, 3, 2, 7, 4, 2, 3, 6, 5, 3, 8, 23, 54}}, []int{2, 2, 2, 3, 3, 3, 3, 4, 5, 6, 6, 6, 7, 8, 23, 54}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sort(tt.args.orig); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sort() = %v, want %v", got, tt.want)
			}
		})
	}
}
