package logic

import (
	"testing"
	"time"

	"github.com/crossbone-magister/tag-tree/data"
)

var expectedTotals = map[string]float64{
	"root":       60 * 60,
	"project1":   40 * 60,
	"activity11": 30 * 60,
	"project2":   20 * 60,
	"activity21": 10 * 60,
	"activity22": 10 * 60,
}

func TestCalculateTreeTotal(t *testing.T) {
	zeroDuration, err := time.ParseDuration("0s")
	if err != nil {
		t.Fatal("Error preparing test data", err)
	}
	tenMinutes, err := time.ParseDuration("10m")
	if err != nil {
		t.Fatal("Error preparing test data", err)
	}
	thirtyMinutes, err := time.ParseDuration("30m")
	if err != nil {
		t.Fatal("Error preparing test data", err)
	}
	root := data.TagTreeNode{
		Tag:      "root",
		Duration: zeroDuration,
		Children: []*data.TagTreeNode{
			{
				Tag:      "project1",
				Duration: tenMinutes,
				Children: []*data.TagTreeNode{
					{
						Tag:      "activity11",
						Duration: thirtyMinutes,
						Children: make([]*data.TagTreeNode, 0),
					},
				},
			},
			{
				Tag:      "project2",
				Duration: zeroDuration,
				Children: []*data.TagTreeNode{
					{
						Tag:      "activity21",
						Duration: tenMinutes,
						Children: make([]*data.TagTreeNode, 0),
					},
					{
						Tag:      "activity22",
						Duration: tenMinutes,
						Children: make([]*data.TagTreeNode, 0),
					},
				},
			},
		},
	}
	CalculateTreeTotal(&root)
	verifyTotals(&root, expectedTotals, t)
}

func verifyTotals(node *data.TagTreeNode, expectedTotals map[string]float64, t *testing.T) {
	if expected, ok := expectedTotals[node.Tag]; ok {
		if node.Duration.Seconds() != expected {
			t.Errorf("Duration of node [%s] is [%f] instead of [%f]", node.Tag, node.Duration.Seconds(), expected)
			for _, child := range node.Children {
				verifyTotals(child, expectedTotals, t)
			}
		}
	} else {
		t.Fatal("Missing expected key:", node.Tag)
	}
}
