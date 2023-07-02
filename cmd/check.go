/*
Copyright © 2023 Kyle Chadha @kylechadha
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/opencamp-hq/core/client"
	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check [campground_id] [start_date (MM-DD-YYYY)] [end_date (MM-DD-YYYY)]",
	Short: "Check campground availability",
	Long: `Check campground availability. You'll need to get the campground ID
by calling 'opencamp search [campground]' first.

Note that start_date and end_date should be in MM-DD-YYYY format.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Argument validation.
		switch len(args) {
		case 0:
			l.Error("campground_id is a required argument")
			return
		case 1:
			l.Error("start_date in MM-DD-YYYY format is a required argument")
			return
		case 2:
			l.Error("end_date in MM-DD-YYYY format is a required argument")
			return
		}

		campgroundID := args[0]
		startDate := args[1]
		endDate := args[2]

		start, err := time.Parse("01-02-2006", startDate)
		if err != nil {
			l.Error("Unable to parse start date", "err", err)
			return
		}

		end, err := time.Parse("01-02-2006", endDate)
		if err != nil {
			l.Error("Unable to parse end date", "err", err)
			return
		}

		// Check availability.
		c := client.New(l, 10*time.Second)
		sites, err := c.Availability(campgroundID, start, end)
		if err != nil {
			l.Error("Unable to determine campground availability", "err", err)
			return
		}

		if len(sites) == 0 {
			fmt.Println("Sorry we didn't find any available campsites!")
		} else {
			fmt.Println("The following sites are available for those dates:")
			for _, s := range sites {
				fmt.Printf(" - Site %-20s Book at: https://www.recreation.gov/camping/campsites/%s\n", s.Site, s.CampsiteID)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
