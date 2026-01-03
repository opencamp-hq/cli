/*
Copyright © 2023 Kyle Chadha @kylechadha
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/opencamp-hq/core/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var searchCmd = &cobra.Command{
	Use:   "search [campground]",
	Short: "Search for campgrounds. Use this to get the campground ID",
	Long: `Search for campgrounds. Use this to get the campground ID.

Note that for the time being you must search by the name of a campground,
not a park area or city.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			l.Error("campground name is a required argument")
			return
		}

		// Search campgrounds.
		c := client.New(l, 10*time.Second)
		campgrounds, err := c.Suggest(args[0])
		if err != nil {
			l.Error("Unable to retrieve campground info", "err", err)
			return
		}
		if len(campgrounds) == 0 {
			fmt.Println("Sorry, no campgrounds with that name were found.")
			return
		}

		if viper.GetBool("json") {
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

func init() {
	rootCmd.AddCommand(searchCmd)

	var json bool
	searchCmd.Flags().BoolVar(&json, "json", false, "output in JSON format")
}
