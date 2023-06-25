# gopeek

A Command Line Interface for "peeking" at CSV files.

It allows users to view the
  - first ten lines of a CSV file
  - the entire file
  - a specified number of lines from a CSV file in the terminal.


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
