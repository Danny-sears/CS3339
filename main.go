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
const MAXopcodeSize = 11

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
	if len(line) >= MAXopcodeSize {
		line = line[:MAXopcodeSize]

		// This switch statement finds the opcode instruction
		switch line {

		// Checks 11-bit opcode cases FIRST, then narrows down to 9, 8, and 6-bit opcodes
		case "10001010000":
			fmt.Println("AND")
		case "10001011000":
			fmt.Println("ADD") // Currently for testing
		case "10101010000":
			fmt.Println("ORR")
		case "11001011000":
			fmt.Println("SUB")
		case "11010011010":
			fmt.Println("LSR")
		case "11010011011":
			fmt.Println("LSL")
		case "11111000000":
			fmt.Println("STUR")
		case "11111000010":
			fmt.Println("LDUR")
		case "11010011100":
			fmt.Println("ASR")
		case "11101010000":
			fmt.Println("EOR")
		} // STILL NEEDS THE REST OF THE OPCODES
		fmt.Println(line) // Currently for testing
	}

}
