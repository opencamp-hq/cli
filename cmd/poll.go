/*
Copyright © 2023 Kyle Chadha @kylechadha
*/
package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/inconshreveable/log15"
	"github.com/opencamp-hq/core/client"
	"github.com/spf13/cobra"
)

// pollCmd represents the poll command
var pollCmd = &cobra.Command{
	Use:   "poll [campground_id] [start_date (MM-DD-YYYY)] [end_date (MM-DD-YYYY)]",
	Short: "Continuously polls campground availability",
	Long: `Continously poll campground availability. You'll need to get the campground ID
	by calling 'rgn search [campground]' first.
	
	Note that start_date and end_date should be in MM-DD-YYYY format.`,
	Run: func(cmd *cobra.Command, args []string) {
		l := log15.New()
		l.SetHandler(log15.StreamHandler(os.Stdout, log15.TerminalFormat()))

		if verbose {
			l.SetHandler(log15.LvlFilterHandler(log15.LvlDebug, l.GetHandler()))
		} else {
			l.SetHandler(log15.LvlFilterHandler(log15.LvlInfo, l.GetHandler()))
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

		i, err := time.ParseDuration(interval)
		if err != nil {
			l.Error("Unable to parse interval", "err", err)
			return
		}

		if i < time.Minute || i > 24*time.Hour {
			l.Error("Polling interval is out of bounds. Choose a time between 1m and 24h")
			return
		}

		c := client.New(l, 10*time.Second)
		ctx := context.Background()
		sites, err := c.Poll(ctx, campgroundID, start, end, i)
		if err != nil {
			l.Error("Encountered an error while polling", "err", err)
			return
		}

		fmt.Println("The following sites are available for those dates:")
		for _, s := range sites {
			fmt.Printf(" - Site %-15s Book at: https://www.recreation.gov/camping/campsites/%s\n", s.Site, s.CampsiteID)
		}
	},
}

var interval string

func init() {
	rootCmd.AddCommand(pollCmd)

	pollCmd.Flags().StringVar(&interval, "interval", "10m", "polling interval. Specify a time between 1m and 24h")
}
