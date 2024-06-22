package command

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
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

	var generateCmd = &cobra.Command{
		Use:   "generate [amount] [output]",
		Short: "Generate csv file with randomized data",
		Run:   Generate,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("amount is required")
			}
			_, err := strconv.Atoi(args[0])
			if err != nil {
				return errors.New("amount should be of type int")
			}
			if len(args) < 2 {
				return errors.New("output path is required")
			}
			return nil
		},
	}

	generateDelimiter = generateCmd.Flags().StringP("delimiter", "d", ",", "Delimiter used in CSV file")
	generateSimple = generateCmd.Flags().BoolP("simple", "s", false, "Generate simple csv file")

	rootCmd.AddCommand(reformatCmd)
	rootCmd.AddCommand(lintCmd)
	rootCmd.AddCommand(generateCmd)

	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		return
	}
}
