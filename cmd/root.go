/*
Copyright © 2023 Kyle Chadha @kylechadha
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "opencamp",
	Short: "Determine recreation.gov campground availability",
	Long: `
 ______  ______  ______  __   __  ______  ______  __    __  ______
/\  __ \/\  == \/\  ___\/\ "-.\ \/\  ___\/\  __ \/\ "-./  \/\  == \
\ \ \/\ \ \  _-/\ \  __\\ \ \-.  \ \ \___\ \  __ \ \ \-./\ \ \  _-/
 \ \_____\ \_\   \ \_____\ \_\\"\_\ \_____\ \_\ \_\ \_\ \ \_\ \_\
  \/_____/\/_/    \/_____/\/_/ \/_/\/_____/\/_/\/_/\/_/  \/_/\/_/

		
This command line tool allows you to check whether a campground managed by
recreation.gov has any available sites. First search for the campground you're interested
in with 'opencamp search [campground]', and note the campground ID of the correct match.

Then check its availability with 'opencamp check [id] [start_date] [end_date]'. Note
that start_date and end_date should be in MM-DD-YYYY format. If there are any available
sites, we'll let you know.

Finally, you can continuously poll availability to see if any cancelations occur. Do
this with 'opencamp poll [id] [start_date] [end_date] --interval=10m.' The application will
run continuously and check availability (eg) every 10 minutes until a campsite becomes
available or today's date is passed the start_date.
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var verbose bool

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose logging output")
}
