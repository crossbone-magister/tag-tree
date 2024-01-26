package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/crossbone-magister/tag-tree/config"
	"github.com/crossbone-magister/tag-tree/data"
	"github.com/crossbone-magister/tag-tree/logic"
	"github.com/crossbone-magister/tag-tree/output"
	"github.com/crossbone-magister/timewlib"
)

const c_HIERARCHY_PRIORITY = 0
const c_INTERVAL_PRIORITY = 1

func main() {
	var err error
	if raw, err := timewlib.Parse(os.Stdin); err == nil {
		handleLogging(raw.Configuration)
		hierarchy := config.ExtractInitialHierarchy(raw.Configuration)
		if intervals, err := timewlib.Process(raw.Intervals); err == nil {
			root := data.NewNode("root")
			log.Println("---HIERARCHY CREATION START---")
			for _, tags := range hierarchy {
				logic.AddIntervalToTree(root, *timewlib.NewInterval(0, 0, 0, 0), tags, c_HIERARCHY_PRIORITY)
			}
			log.Println("---HIERARCHY CREATION END---")
			log.Println("---INTERVALS ADDITION START---")
			for _, interval := range intervals {
				logic.AddIntervalToTree(root, interval, interval.Tags, c_INTERVAL_PRIORITY)
			}
			log.Println("---INTERVALS ADDITION END---")
			logic.CalculateTreeTotal(root)
			if maxDepth, err := config.RetrieveMaxDepth(raw.Configuration); err == nil {
				logic.PruneToMaxDepth(root, maxDepth)
			} else {
				switch err.(type) {
				case *config.ConfigurationKeyNotFound:
					log.Println(err)
				default:
					fmt.Fprintf(os.Stderr, "Error reading max depth configuration: %s\n", err)
				}
			}
			if _, ok := raw.Configuration[config.PRUNE_CONFIG_KEY]; ok {
				logic.PruneEmptyBranches(root)
			}
			output.PrintTree(root, os.Stdout)
		}
	}
	if err != nil {
		printErrorAndExit(err)
	}
}

func handleLogging(rawConfig map[string]string) {
	if !config.IsDebugEnabled(rawConfig) {
		log.SetOutput(io.Discard)
	}
}

func printErrorAndExit(err error) {
	fmt.Printf("Error while parsing: %s\n", err)
	os.Exit(1)
}
