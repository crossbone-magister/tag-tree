package logic

import (
	"os"
	"strconv"
	"testing"

	"github.com/crossbone-magister/tag-tree/data"
	"github.com/crossbone-magister/tag-tree/output"
)

func TestPruneEmptyBranches(t *testing.T) {
	root := data.NewNode("root")
	child1 := data.AppendNewChild(root, "tag1", 0)
	child1.Duration += 10
	data.AppendNewChild(root, "tag2", 0)
	PruneEmptyBranches(root)
	if len(root.Children) > 1 {
		t.Error("Children were not pruned")
	}
	if root.Children[0].Tag != "tag1" {
		t.Error("Expected children tag1 is not present")
	}
}

func TestPruneToMaxDepth(t *testing.T) {
	root := data.NewNode("root")
	node := root
	for i := 1; i <= 10; i++ {
		node = data.AppendNewChild(node, strconv.Itoa(i), 0)
	}
	output.PrintTree(root, os.Stdout)
	var maxDepth uint64 = 3
	PruneToMaxDepth(root, maxDepth)
	output.PrintTree(root, os.Stdout)
	node = root
	for i := 0; i < int(maxDepth); i++ {
		if len(node.Children) <= 0 {
			t.Errorf("Node at level [%d] has no children", i)
		}
		node = node.Children[0]
	}
	println(node.Tag)
	if len(node.Children) > 0 {
		t.Error("Latest node was not pruned.")
	}
}
