package cmd

import "github.com/spf13/cobra"

var jobCmd = &cobra.Command{
	Use:   "job",
	Short: "Manages jobs (like a calendar)",
}

func init() {
	rootCmd.AddCommand(jobCmd)
}
