/*
Copyright © 2023 Kyle Chadha @kylechadha
*/
package cmd

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/opencamp-hq/core/client"
	"github.com/opencamp-hq/core/models"
	"github.com/opencamp-hq/core/notify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var pollCmd = &cobra.Command{
	Use:   "poll [campground_id] [start_date (MM-DD-YYYY)] [end_date (MM-DD-YYYY)]",
	Short: "Continuously polls campground availability",
	Long: `Continously poll campground availability. You'll need to get the campground ID
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

		// Flag validation.
		intervalFlag := viper.GetString("interval")
		interval, err := time.ParseDuration(intervalFlag)
		if err != nil {
			l.Error("Unable to parse interval", "err", err)
			return
		}

		if interval < time.Minute || interval > 24*time.Hour {
			l.Error("Polling interval is out of bounds. Choose a time between 1m and 24h")
			return
		}

		notifyFlag := viper.GetString("notify")
		var e *notify.SMTPSender
		if len(notifyFlag) > 0 {
			switch strings.ToLower(notifyFlag) {
			case "email":
				cfg, err := GetSMTPConfig()
				if err != nil {
					l.Warn("Unable to get SMTP config. Email notifications will not be sent", "err", err)
				}

				e, err = notify.NewSMTPSender(*cfg)
				if err != nil {
					l.Warn("Unable to setup SMTP sender. Email notifications will not be sent", "err", err)
				}

			case "sms":

			default:
				l.Error("Unknown notification mechanism. Please specify 'email' or 'sms'")
				return
			}
		}

		// Poll availability.
		c := client.New(l, 10*time.Second)
		ctx := context.Background()
		sites, err := c.Poll(ctx, campgroundID, start, end, interval)
		if err != nil {
			l.Error("Encountered an error while polling", "err", err)
			return
		}

		// Notify the user.
		fmt.Println("Just in! The following sites are now available for those dates:")
		for _, s := range sites {
			fmt.Printf(" - Site %-20s Book at: https://www.recreation.gov/camping/campsites/%s\n", s.Site, s.CampsiteID)
		}
		fmt.Print("\n")

		cg, err := c.SearchByID(campgroundID)
		if err != nil {
			cg = &models.Campground{EntityID: campgroundID}
			l.Warn("Unable to pull campground data for rich notifications", "err", err)
		}

		if len(notifyFlag) > 0 {
			switch strings.ToLower(notifyFlag) {
			case "email":
				cfg, _ := GetSMTPConfig()
				err = e.Send(cfg.Email, cg, startDate, endDate, sites)
				if err != nil {
					l.Error("Unable to send email", "err", err)
					return
				}
				l.Info("Notification email sent")

			case "sms":
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(pollCmd)

	var interval string
	pollCmd.Flags().StringVar(&interval, "interval", "10m", "polling interval. Specify a time between 1m and 24h")
}
