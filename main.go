package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	// open the input file
	inputFile, err := os.Open("input.txt")
	if err != nil { // if there is an error
		fmt.Println("Error opening input file:", err)
		return
	}
	// close the file when it's done
	defer inputFile.Close()

	// create a reegular expression to match senctence ends
	sentenceEndRegex := regexp.MustCompile(`[.!?](\s|$)`)

	// scanner to read the input file line by line
	scanner := bufio.NewScanner(inputFile)

	// initialize counter file for the output file
	fileCounter := 1

	// loopthrough the lines in the input file
	for scanner.Scan() {
		// read the line from the input file
		line := scanner.Text()

		// split lines into sentences
		sentences := sentenceEndRegex.Split(line, -1)

		// loop through the sentences
		for _, sentence := range sentences {
			// skip empty sentences
			if len(strings.TrimSpace(sentence)) == 0 {
				continue
			}

			//  create a new output file
			outputFile, err := os.Create(fmt.Sprintf("output_%d.txt", fileCounter))
			if err != nil {
				fmt.Println("Error creating output file:", err)
				return
			}
			defer outputFile.Close()

			// write the sentence to the output file
			_, err = outputFile.WriteString(sentence + "\n")
			if err != nil {
				fmt.Println("Error writing to output file:", err)
				return
			}
			// increment the file counter
			fileCounter++
		}
	}

	// check for errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}
	fmt.Println("processing complete")
}
