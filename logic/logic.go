package logic

import (
	"log"
	"slices"
	"time"

	"github.com/crossbone-magister/tag-tree/data"
	"github.com/crossbone-magister/timewlib"
)

func AddIntervalToTree(node *data.TagTreeNode, interval timewlib.Interval, tagList []string, priority uint8) {
	log.Printf("Visiting [%p]-[%s] with remaining tags: %s", node, node.Tag, tagList)
	if len(tagList) <= 0 {
		log.Printf("Node [%p]-[%s] is a leaf.\n", node, node.Tag)
		node.Duration += interval.Duration()
	} else if len(node.Children) <= 0 {
		child := data.AppendNewChild(node, tagList[0], priority)
		log.Printf("Appending child [%p]-[%s] to node [%p]-[%s]\n", child, child.Tag, node, node.Tag)
		AddIntervalToTree(child, interval, tagList[1:], priority)
	} else {
		var next *data.TagTreeNode
		nextIndex := -1
		log.Println("Searching for child node with tag in remaining tags...")
		for _, child := range node.Children {
			//If there's more than one child with the lowest priority, this search takes the last one
			for index, tag := range tagList {
				if tag == child.Tag && (nextIndex < 0 || child.Priority <= next.Priority) {
					next = child
					nextIndex = index
				}
			}
		}
		if nextIndex >= 0 {
			log.Printf("Found child [%s] at index [%d], visiting...\n", next.Tag, nextIndex)
			AddIntervalToTree(next, interval, slices.Delete(tagList, nextIndex, nextIndex+1), priority)
		} else {
			log.Printf("Child not found. Adding new child [%s] to [%s]\n", tagList[0], node.Tag)
			child := data.AppendNewChild(node, tagList[0], priority)
			AddIntervalToTree(child, interval, tagList[1:], priority)
		}
	}
}

func CalculateTreeTotal(node *data.TagTreeNode) time.Duration {
	total := node.Duration
	for _, child := range node.Children {
		total += CalculateTreeTotal(child)
	}
	node.Duration = total
	return node.Duration
}

func PruneEmptyBranches(node *data.TagTreeNode) {
	node.Children = slices.DeleteFunc(node.Children, func(child *data.TagTreeNode) bool {
		log.Printf("Should we prune [%s]-[%s]? %t\n", child.Tag, child.Duration.String(), child.Duration.Seconds() <= 0)
		return child.Duration.Seconds() <= 0
	})
	for _, child := range node.Children {
		PruneEmptyBranches(child)
	}
}

func PruneToMaxDepth(node *data.TagTreeNode, maxDepth uint64) {
	pruneToMaxDepth(node, 0, maxDepth)
}

func pruneToMaxDepth(node *data.TagTreeNode, currentDepth, maxDepth uint64) {
	if currentDepth == maxDepth {
		node.Children = make([]*data.TagTreeNode, 0)
	} else {
		for _, child := range node.Children {
			pruneToMaxDepth(child, currentDepth+1, maxDepth)
		}
	}
}
