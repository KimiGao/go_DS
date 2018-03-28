// Package ds provides ...
package ds

import (
	"testing"
)

func TestAddBT(t *testing.T) {
	tree := &BiTree{data: 6}
	tree.AddBT(BiTree{data: 2})

	if tree.left.data != 2 {
		t.Errorf("tree.Add() is error")
	}
}
