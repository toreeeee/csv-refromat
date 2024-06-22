package command

import (
	"csv/table"
	"csv/table/table_row"
	"csv/utils/console"
	"fmt"
	"github.com/spf13/cobra"
	"io/fs"
	"math/rand/v2"
	"os"
	"strconv"
)

var generateDelimiter *string

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ "

func randInt(min int, max int) int {
	return min + rand.IntN(max-min)
}

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.IntN(len(letterBytes))]
	}
	return string(b)
}

func Generate(cmd *cobra.Command, args []string) {
	amountOfRows, _ := strconv.ParseInt(args[0], 10, 32)
	outPath := args[1]

	if len(*generateDelimiter) != 1 {
		console.Error("Input delimiter should be 1 character: '%s' given\n", *delimiter)
		os.Exit(1)
	}

	t := table.Table{
		Rows: make([]table_row.TableRow, 0),
		Headings: table_row.TableRow{
			Cols: []string{"1st", "2nd", "3rd", "4th"},
		},
	}

	for i := 0; i < int(amountOfRows); i++ {
		t.Rows = append(t.Rows, table_row.TableRow{
			Cols: []string{
				randStringBytes(randInt(10, 16)),
				randStringBytes(randInt(10, 16)),
				randStringBytes(randInt(10, 16)),
				randStringBytes(randInt(10, 16)),
			},
		})
	}

	encoded := t.Encode(*generateDelimiter)

	err := os.WriteFile(outPath, []byte(encoded), fs.ModePerm)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	console.Success("Generated file: %s\n", outPath)

	//
	//startTime := time.Now()
	//
	//reformatted := reformatFile(inputFile)
	//
	//fmt.Printf("took %s\n", time.Since(startTime))
	//
	//var outPath string
	//if len(*outputFile) != 0 {
	//	outPath = *outputFile
	//} else {
	//	outPath = inputFile
	//}
	//
	//if *writeToFile {
	//	err := os.WriteFile(outPath, []byte(reformatted), fs.ModePerm)
	//	if err != nil {
	//		fmt.Println(err)
	//		os.Exit(1)
	//	}
	//
	//	fmt.Println("File has been updated")
	//} else {
	//	fmt.Println(reformatted)
	//}
}
