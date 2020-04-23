package cmd

import (
	"fmt"

	example "github.com/RyanSiu1995/intro-to-golang/internal"
	"github.com/spf13/cobra"
)

func init() {
	var exampleCmd = &cobra.Command{
		Use:   "version",
		Short: "Show the current version of gcb-visualizer",
		Run: func(cmd *cobra.Command, args []string) {
			interactive, _ := cmd.Flags().GetBool("interactive")
			example.Run(interactive)
		},
	}
	version.Flags().BoolP("interactive", "i", true, "Execute the example with interactive mode")
	rootCmd.AddCommand(exampleCmd)
}
