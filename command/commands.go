package command

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
)

func Init() {
	var rootCmd = &cobra.Command{
		Use:   "csv",
		Short: "CSV utility program",
	}

	var reformatCmd = &cobra.Command{
		Use:   "format",
		Short: "Reformat csv file",
		Run:   Format,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("file path is required")
			}
			return nil
		},
	}

	outputFile = reformatCmd.Flags().StringP("output", "o", "", "Output CSV file")
	delimiter = reformatCmd.Flags().StringP("delimiter", "d", ",", "Delimiter for CSV file")
	outputDelimiter = reformatCmd.Flags().StringP("delimiter-out", "s", ",", "Delimiter for CSV file")
	writeToFile = reformatCmd.Flags().BoolP("write", "w", false, "Write to file")

	rootCmd.AddCommand(reformatCmd)

	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		return
	}
}
