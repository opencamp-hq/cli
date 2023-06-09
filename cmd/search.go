/*
Copyright © 2023 Kyle Chadha @kylechadha
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/inconshreveable/log15"
	"github.com/opencamp-hq/core/client"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search [campground]",
	Short: "Search for campgrounds. Use this to get the campground ID",
	Long: `Search for campgrounds. Use this to get the campground ID.

Note that for the time being you must search by the name of a campground,
not a park area or city.`,
	Run: func(cmd *cobra.Command, args []string) {
		l := log15.New()
		l.SetHandler(log15.StreamHandler(os.Stdout, log15.TerminalFormat()))

		if verbose {
			l.SetHandler(log15.LvlFilterHandler(log15.LvlDebug, l.GetHandler()))
		} else {
			l.SetHandler(log15.LvlFilterHandler(log15.LvlInfo, l.GetHandler()))
		}

		c := client.New(l, 10*time.Second)
		campgrounds, err := c.Search(args[0])
		if err != nil {
			l.Error("Unable to retrieve campground info", "err", err)
			return
		}
		if len(campgrounds) == 0 {
			fmt.Println("Sorry, no campgrounds with that name were found.")
			return
		}

		if jsonOutput {
			bytes, err := json.MarshalIndent(campgrounds, "", "  ")
			if err != nil {
				l.Error("Unable to print output as JSON", "err", err)
				return
			}

			fmt.Println(string(bytes))
		} else {
			var longestName int
			var longestCityState int
			for _, cg := range campgrounds {
				lenName := len(cg.Name)
				lenCityState := len(fmt.Sprintf("%s, %s", cg.City, cg.StateCode))
				if lenName > longestName {
					longestName = lenName
				}
				if lenCityState > longestCityState {
					longestCityState = lenCityState
				}
			}

			for _, cg := range campgrounds {
				fmt.Printf("- %-*s %-*s ID: %s\n", longestName+5, cg.Name, longestCityState+5, fmt.Sprintf("%s, %s", cg.City, cg.StateCode), cg.EntityID)
			}
		}
	},
}

var jsonOutput bool

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().BoolVar(&jsonOutput, "json", false, "output in JSON format")
}
