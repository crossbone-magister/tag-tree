package data

import "testing"

func TestNewNode(t *testing.T) {
	node := NewNode("tag1")
	if node.Tag != "tag1" {
		t.Error("tag1 != ", node.Tag)
	}
	if len(node.Children) > 0 {
		t.Error("Node children array is not empty")
	}
}

func TestAppendNewChild(t *testing.T) {
	parent := NewNode("parent")
	child := AppendNewChild(parent, "child", 1)
	if child.Tag != "child" {
		t.Error("child !=", child.Tag)
	}
	if len(child.Children) > 0 {
		t.Error("Node children array is not empty")
	}
	if child.Priority != 1 {
		t.Error("Node children priority is not set correctly")
	}
}
