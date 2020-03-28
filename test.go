/* The point of this is that i'm supposed to read a number of votes from a text file. The file counts the names and the amount of votes
and then it outputs the result.

The desired output should be this

Amber Graham: 3
Brian Martin : 2

But the code keeps outputting something else. It prints the name each time with the number one, each time.

In this regard I'm not sure what it is i'm doing wrong. */

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func datastring(fileName string) ([]string, error) {

	var lines []string

	file, err := os.Open(fileName)

	if err != nil {

		return nil, err
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		line := scanner.Text()
		lines = append(lines, line)

	}

	err = file.Close()
	if err != nil {
		return nil, err
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return lines, nil
}

func main() {

	lines, err := datastring("votes.txt")

	if err != nil {
		log.Fatal(err)
	}

	var names []string
	var counts []int

	for _, line := range lines {

		matched := false

		for i, name := range names {

			if name == line {

				counts[i]++
				matched = true
			}

		}

		if matched == false {

			names = append(names, line)
			counts = append(counts, 1)
		}

	}

	for i, name := range names {

		fmt.Printf("%s: %d\n ", name, counts[i])
	}
}
