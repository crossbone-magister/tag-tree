package output_test

import (
	"bytes"
	"testing"

	"github.com/crossbone-magister/tag-tree/data"
	"github.com/crossbone-magister/tag-tree/output"
)

func TestOutput(t *testing.T) {
	root := data.TagTreeNode{
		Tag:      "root",
		Duration: 0,
		Children: []*data.TagTreeNode{
			{
				Tag:      "project1",
				Duration: 0,
				Children: []*data.TagTreeNode{
					{
						Tag:      "activity11",
						Duration: 0,
						Children: make([]*data.TagTreeNode, 0),
					},
					{
						Tag:      "activity13",
						Duration: 0,
						Children: make([]*data.TagTreeNode, 0),
					},
					{
						Tag:      "activity12",
						Duration: 0,
						Children: make([]*data.TagTreeNode, 0),
					},
				},
			},
			{
				Tag:      "project2",
				Duration: 0,
				Children: []*data.TagTreeNode{
					{
						Tag:      "activity21",
						Duration: 0,
						Children: make([]*data.TagTreeNode, 0),
					},
					{
						Tag:      "activity22",
						Duration: 0,
						Children: make([]*data.TagTreeNode, 0),
					},
				},
			},
		},
	}
	var buff bytes.Buffer
	output.PrintTree(&root, &buff)
	var actual = buff.String()
	var expected = `0s
├── project1 - 0s
│   ├── activity11 - 0s
│   ├── activity12 - 0s
│   └── activity13 - 0s
└── project2 - 0s
    ├── activity21 - 0s
    └── activity22 - 0s
`
	if expected != actual {
		t.Errorf("Actual output [%s] is different from expected [%s]", actual, expected)
	}
}
