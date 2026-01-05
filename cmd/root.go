package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "GlimmerWorksBuilder",
	Short: "A build tool developed for GlimmerWorks.",
	Long: `GlimmerWorksBuilder is a build tool developed for GlimmerWorks. 
GlimmerWorksBuilder can pull the latest git repository into GlimmerWorks. 
GlimmerWorksBuilder can also install/uninstall so libraries into Android projects.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
