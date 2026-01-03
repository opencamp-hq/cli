/*
Copyright © 2023 Kyle Chadha @kylechadha
*/
package cmd

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands.
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

Both the check and poll commands accept a --notify flag which can be set to email or
(in the future) text. Please note you will be required to set your notification platform's
credentials, ie: your smtp configuration or twilio API key for email and text, respectively.
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
				log.Fatalf("Unable to read config file: %v", err)
			}
		}

		// Configure logger.
		var level slog.Level
		if viper.GetBool("verbose") {
			level = slog.LevelDebug
		} else {
			level = slog.LevelInfo
		}
		l = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level}))

		if viper.GetBool("verbose") {
			configMap := viper.AllSettings()
			var config []any
			for k, v := range configMap {
				// TODO: If we wanted to redact the password here, we'd need to recursively
				// process v in the case where it's a map or other nestable data structure.
				config = append(config, k, v)
			}
			l.Debug("Running in debug mode", config...)
		}
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints build version",
	Long:  "Prints build version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("opencamp version %s\n", version)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

var l *slog.Logger
var version = "dev"

func init() {
	var verbose bool
	var config string
	var notify string

	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose logging output")
	rootCmd.PersistentFlags().StringVarP(&config, "config", "c", "", "config file location")
	rootCmd.PersistentFlags().StringVarP(&notify, "notify", "n", "", "specify 'email' or 'text' if you would like to receive an email or text if availability is found")

	rootCmd.AddCommand(versionCmd)
}
