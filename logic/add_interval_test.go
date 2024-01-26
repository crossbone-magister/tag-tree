package logic

import (
	"testing"

	"github.com/crossbone-magister/tag-tree/data"
	"github.com/crossbone-magister/timewlib"
)

var testInterval = timewlib.NewInterval(9, 0, 10, 0)

func TestAddIntervalToTreeLeafNode(t *testing.T) {
	node := data.NewNode("leaf")
	AddIntervalToTree(node, *testInterval, make([]string, 0), 0)
	if node.Duration.Hours() != 1 {
		t.Errorf("Duration of one hour hasn't been added. Current duration for node is %f", node.Duration.Hours())
	}
}

func TestAddIntervalToTreeNewTagOnLeafNode(t *testing.T) {
	node := data.NewNode("tag1")
	AddIntervalToTree(node, *testInterval, []string{"tag11"}, 0)
	if len(node.Children) != 1 && node.Children[0].Tag != "tag11" {
		t.Error("New children wasn't added to node")
	}
}

func TestAddIntervalToTreeExistingChildOfNode(t *testing.T) {
	root := data.NewNode("root")
	data.AppendNewChild(root, "tag1", 0)
	AddIntervalToTree(root, *testInterval, []string{"tag1"}, 0)
	if root.Children[0].Duration.Hours() < 1 {
		t.Errorf("Duration of one hour hasn't been added. Current duration for node is %f", root.Children[0].Duration.Hours())
	}
}

func TestAddIntervalToTreeNewChildOfNodeWithChildren(t *testing.T) {
	root := data.NewNode("root")
	data.AppendNewChild(root, "tag1", 0)
	AddIntervalToTree(root, *testInterval, []string{"tag2"}, 0)
	if len(root.Children) < 2 && root.Children[1].Tag != "tag2" {
		t.Error("New children wasn't added to node")
	}
}
