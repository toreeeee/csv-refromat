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

	var lintCmd = &cobra.Command{
		Use:   "lint",
		Short: "Check csv file for errors",
		Run:   Lint,
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

	lintDelimiter = lintCmd.Flags().StringP("delimiter", "d", ",", "Delimiter used in CSV file")

	rootCmd.AddCommand(reformatCmd)
	rootCmd.AddCommand(lintCmd)

	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		return
	}
}
