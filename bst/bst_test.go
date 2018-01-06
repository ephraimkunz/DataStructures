package bst

import (
	"reflect"
	"testing"
)

func TestNode_Exists(t *testing.T) {
	type fields struct {
		left  *Node
		right *Node
		value int
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
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Node{
				left:  tt.fields.left,
				right: tt.fields.right,
				value: tt.fields.value,
			}
			if got := n.Exists(tt.args.item); got != tt.want {
				t.Errorf("Node.Exists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_Remove(t *testing.T) {
	type fields struct {
		left  *Node
		right *Node
		value int
	}
	type args struct {
		item int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Node
		want1  bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Node{
				left:  tt.fields.left,
				right: tt.fields.right,
				value: tt.fields.value,
			}
			got, got1 := n.Remove(tt.args.item)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Node.Remove() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Node.Remove() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNode_Add(t *testing.T) {
	type fields struct {
		left  *Node
		right *Node
		value int
	}
	type args struct {
		item int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Node
		want1  bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Node{
				left:  tt.fields.left,
				right: tt.fields.right,
				value: tt.fields.value,
			}
			got, got1 := n.Add(tt.args.item)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Node.Add() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Node.Add() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNode_IsBST(t *testing.T) {
	type fields struct {
		left  *Node
		right *Node
		value int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Node{
				left:  tt.fields.left,
				right: tt.fields.right,
				value: tt.fields.value,
			}
			if got := n.IsBST(); got != tt.want {
				t.Errorf("Node.IsBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
