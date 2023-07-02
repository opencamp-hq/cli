/*
Copyright © 2023 Kyle Chadha @kylechadha
*/
package cmd

import (
	"os"
	"strings"

	"github.com/inconshreveable/log15"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
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
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Configure Viper with flags.
		cmd.Flags().VisitAll(func(f *pflag.Flag) {
			// Remove dashes (if any are present) from the flag name.
			configName := strings.ReplaceAll(f.Name, "-", "")
			viper.BindPFlag(configName, f)
		})

		// Configure Viper with env vars.
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		viper.AutomaticEnv()

		// Configure Viper with config file.
		configFilepath := viper.GetString("config")
		if configFilepath != "" {
			viper.SetConfigFile(configFilepath)
		} else {
			viper.SetConfigName("config")
			viper.AddConfigPath(".")
		}

		if err := viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
				l.Crit("Unable to read config file: %s", "err", err)
				os.Exit(1)
			}
		}

		// Configure logger.
		l = log15.New()
		l.SetHandler(log15.StreamHandler(os.Stdout, log15.TerminalFormat()))

		if viper.GetBool("verbose") {
			l.SetHandler(log15.LvlFilterHandler(log15.LvlDebug, l.GetHandler()))

			configMap := viper.AllSettings()
			var config []interface{}
			for k, v := range configMap {
				config = append(config, k, v)
			}
			l.Debug("Running in debug mode", config...)
		} else {
			l.SetHandler(log15.LvlFilterHandler(log15.LvlInfo, l.GetHandler()))
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var l log15.Logger

func init() {
	var verbose bool
	var config string
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose logging output")
	rootCmd.PersistentFlags().StringVarP(&config, "config", "c", "", "config file location")
}
