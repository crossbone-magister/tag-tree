package output

import (
	"fmt"
	"io"
	"log"
	"slices"
	"strings"

	"github.com/crossbone-magister/tag-tree/data"
)

const c_BRANCH_CHARACTER = "├── "
const c_BRANCH_CHARACTER_LEAF = "└── "
const c_PADDING_PATTERN = "│   "
const c_PADDING_PATTERN_LEAF = "    "

func printTree(node *data.TagTreeNode, prefix string, writer io.Writer) {
	log.Printf("Printing children of node [%s] with prefix [%s]\n", node.Tag, prefix)
	slices.SortFunc(node.Children, sortByTagAlphabetically)
	for i, child := range node.Children {
		branch := c_BRANCH_CHARACTER
		padding := c_PADDING_PATTERN
		if i == len(node.Children)-1 {
			branch = c_BRANCH_CHARACTER_LEAF
			padding = c_PADDING_PATTERN_LEAF
		}
		fmt.Fprintf(writer, "%s%s%s - %s\n", prefix, branch, child.Tag, child.Duration.String())
		printTree(child, prefix+padding, writer)
	}
}

func sortByTagAlphabetically(a, b *data.TagTreeNode) int {
	return strings.Compare(a.Tag, b.Tag)
}

func PrintTree(node *data.TagTreeNode, writer io.Writer) {
	fmt.Fprintln(writer, node.Duration.String())
	printTree(node, "", writer)
}
