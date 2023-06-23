package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"github.com/jedib0t/go-pretty/v6/table"
	"reflect"
)

func main() {

	var usage = `
	Usage: gopeek file-name.csv [options]

	A simple CLI to peek CSV files in the terminal.
	The first 10 lines are printed by default.

	Options and Arguments
	-f : full file
	-n <number>: first n lines
	
	Examples:
	gopeek my-csv-file.csv
		Prints the first 10 lines.

	gopeek my-csv-file.csv -f
		Prints the entire file.
		
	gopeek my-csv-file.csv -n 100
		Prints the first 100 lines.
	`

	if len(os.Args) == 1 {
		fmt.Println(usage)
		os.Exit(0)
	}

	var args []string
	var notargs []string
	for i := 0; i < len(os.Args); i++ {
		if os.Args[i][0] == '-' || i == 0 {
			notargs = append(notargs, os.Args[i])
		} else {
			args = append(args, os.Args[i])
		}
	}

	fileName := args[0]

	// type ErrorHandling int

	// const (
	// 	ContinueOnError ErrorHandling = iota // Return a descriptive error.
	// 	ExitOnError                          // Call os.Exit(2) or for -h/-help Exit(0).
	// 	PanicOnError                         // Call panic with a descriptive error.
	// )
	
	os.Args = notargs
	allLines := flag.Bool("f", false, "bool if entire file is to be printed")
	numRecords := flag.Int("n", 10, "number of lines to be printed")
	
	flag.Parse()

	if reflect.TypeOf(*allLines) != reflect.TypeOf(true) || reflect.TypeOf(*numRecords) != reflect.TypeOf(0) {
		fmt.Println("hdfngjdf")
		os.Exit(0)
	}



	fd, error := os.Open(fileName)

	if error != nil {
		usageAndExit("Could not open file")
	}

	fileReader := csv.NewReader(fd)
	fileReader.LazyQuotes = true
	records, error := fileReader.ReadAll()

	if error != nil {
		usageAndExit("Could not read the file")
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)


	if *allLines {
		// fmt.Println(records)
		recordsPrinted := 0
		for recordsPrinted < len(records) {
			currRecord := records[recordsPrinted]
			// fmt.Println(currRecord)
			t.AppendRow(table.Row{currRecord[0], currRecord[1], currRecord[2]})
			t.AppendSeparator()
			recordsPrinted++
		}
	} else {
		recordsPrinted := 0
		if len(records) < *numRecords {
			*numRecords = len(records)
		}
		for recordsPrinted < *numRecords {
			currRecord := records[recordsPrinted]
			t.AppendRow(table.Row{currRecord[0], currRecord[1], currRecord[2]})
			t.AppendSeparator()
			recordsPrinted++
		}
	}
	// t.SetAutoIndex(true)
	t.Render()

}

func usageAndExit(msg string) {
	if msg != "" {
		fmt.Fprint(os.Stderr, msg)
		fmt.Fprintf(os.Stderr, "\n")
	}

	os.Exit(0)
}
