package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/ecarter202/timebox/models"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(dayCmd)

	// sub commands
	dayCmd.AddCommand(dayCreateCmd)
	dayCmd.AddCommand(dayUpdateCmd)
	dayCmd.AddCommand(dayShowCmd)
	dayCmd.AddCommand(dayDelCmd)
}

var dayCmd = &cobra.Command{
	Use:   "day",
	Short: "Manage days",
}

var dayCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new day",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Creating a new day")
	},
}

var dayUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a day",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Updating a day")
	},
}

var dayShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show a day",
	Run: func(cmd *cobra.Command, args []string) {
		// NOTE:
		// the day object will be pulled from the database
		start1 := time.Now()
		end1 := start1.Add(time.Minute * 25)
		start2 := end1
		end2 := start2.Add(time.Minute * 10)
		start3 := start1.Add(time.Hour * 2)
		end3 := start3.Add(time.Minute * 5)

		day := models.Day{
			Blocks: []models.Block{
				{
					Task: models.Task{
						Name: "Stare...",
					},
					Start: start1,
					End:   end1,
				},
				{
					Task: models.Task{
						Name: "TPS Reports",
					},
					Start: start2,
					End:   end2,
				},
				{
					Task: models.Task{
						Name: "Chotchkies",
					},
					Start: start3,
					End:   end3,
				},
			},
		}

		table := tablewriter.NewWriter(os.Stdout)

		var (
			header  string
			headers []string
		)
		for _, block := range day.Blocks {
			header = fmt.Sprintf("%s->%s", block.Start.Format("3:04PM"), block.End.Format("3:04PM"))
			headers = append(headers, header)
		}

		table.SetHeader(headers)

		// append row
		var row []string
		for _, block := range day.Blocks {
			row = append(row, block.Task.Name)
		}
		table.Append(row)
		// end append row

		table.Render()
	},
}

var dayDelCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a day",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Deleting a day")
	},
}
