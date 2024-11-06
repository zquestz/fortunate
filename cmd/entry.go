package cmd

import (
	"fmt"
	"os"
	"unicode"

	"github.com/zquestz/fortunate/config"
	"github.com/zquestz/fortunate/fyneapp"

	"github.com/spf13/cobra"
)

// EntryCmd is the main command for the application.
var EntryCmd = &cobra.Command{
	Use:   "fortunate",
	Short: "A beautiful motivation/fortune app for Linux.",
	Long:  `A beautiful motivation/fortune app for Linux.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := performCommand(cmd, args)
		if err != nil {
			bail(err)
		}
	},
}

func init() {
	err := config.AppConfig.Load()
	if err != nil {
		bail(fmt.Errorf("failed to load configuration: %w", err))
	}

	prepareFlags()
}

func bail(err error) {
	fmt.Fprintf(os.Stderr, "[Error] %s\n", capitalize(err.Error()))
	os.Exit(1)
}

func capitalize(str string) string {
	if len(str) == 0 {
		return ""
	}
	tmp := []rune(str)
	tmp[0] = unicode.ToUpper(tmp[0])
	return string(tmp)
}

func prepareFlags() {
	EntryCmd.PersistentFlags().BoolVar(
		&config.AppConfig.DisplayVersion, "version", false, "display version")
}

// Where all the work happens.
func performCommand(cmd *cobra.Command, args []string) error {
	if config.AppConfig.DisplayVersion {
		fmt.Printf("%s %s\n", config.AppName, config.Version)
		return nil
	}

	if len(args) != 0 {
		help := cmd.HelpFunc()
		help(cmd, args)
	}

	fyneapp.Run()

	return nil
}
