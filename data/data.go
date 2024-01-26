package data

import (
	"log"
	"time"
)

type TagTreeNode struct {
	Tag      string
	Duration time.Duration
	Priority uint8
	Children []*TagTreeNode
}

func NewNode(tag string) *TagTreeNode {
	return &TagTreeNode{
		Tag:      tag,
		Duration: 0,
		Priority: 0,
		Children: make([]*TagTreeNode, 0),
	}
}

func AppendNewChild(parent *TagTreeNode, tag string, priority uint8) *TagTreeNode {
	child := NewNode(tag)
	child.Priority = priority
	log.Printf("Appending [%p] to [%p]", child, parent)
	parent.Children = append(parent.Children, child)
	return child
}
