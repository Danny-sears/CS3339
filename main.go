package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

// Text file to be opened and read
var txtfile = "rtest1_bin_translated.txt"

// Text file to be written to
var txtfile_output = ""

const bitlineSize = 32
const opcodeSize = 11

func main() {
	//fmt.Println("Hello, World!")
	// This line opens the file and checks for errors
	openfile, err := os.Open(txtfile)
	if err != nil {
		log.Fatal(err)
	}

	// Don't forget to close the file once we're done
	defer openfile.Close()

	// This scanner reads the file line by line
	scanner := bufio.NewScanner(openfile)

	// This loops through every line in a file until the file ends
	for scanner.Scan() {
		// This happens to each line
		fullline := scanner.Text()
		fmt.Printf("line: %s\n", fullline)
		//formatLine(fullline)
		defineOpcode(fullline)

		if err == io.EOF {
			break
		}
	}

}

// This function will format the translated lines to be written into the output file
func formatLine(line string) {
	fmt.Println(line) // Currently for testing
}

// This function reads the first 11 bits of each line and uses a switch statement
// to find the correct opcode
func defineOpcode(line string) {
	if len(line) >= opcodeSize {
		line = line[:opcodeSize]

		// This switch statement finds the opcode instruction
		switch line {
		case "10001011000":
			fmt.Println("ADD") // Currently for testing
		}
		fmt.Println(line) // Currently for testing
	}

}
