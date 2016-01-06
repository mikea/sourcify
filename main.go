package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func readFileLines(fileName string) ([]string, error) {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(content), "\n"), nil
}

func printContext(fileName string, line int) error {
	if fileExists(fileName) {
		fileLines, err := readFileLines(fileName)
		if err != nil {
			return err
		}

		ctxSize := 5
		startIdx := max(0, line-ctxSize)
		endIdx := min(line+ctxSize, len(fileLines)-1)
		for i, l := range fileLines[startIdx:endIdx] {
			separator := ":"
			if i == ctxSize {
				separator = ">:"
			}
			fmt.Printf("%d %2s %s\n", i+startIdx+1, separator, l)
		}
	}

	return nil
}

func main() {
	lineRegexp := regexp.MustCompile(`^(?P<fname>[^:]+):(?P<line>[0-9]+)(:(?P<column>[0-9]+))?$`)
	regexpNames := map[string]int{}
	for i, name := range lineRegexp.SubexpNames() {
		regexpNames[name] = i
	}
	fnameIdx := regexpNames["fname"]
	lineIdx := regexpNames["line"]

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		matches := lineRegexp.FindAllStringSubmatch(line, -1)
		if matches != nil {
			for _, match := range matches {
				fileName := match[fnameIdx]
				lineNumber, err := strconv.Atoi(match[lineIdx])
				if err != nil {
					panic(err)
				}
				// make it 0-based
				lineNumber = lineNumber - 1
				printContext(fileName, lineNumber)
			}
		}

		os.Stdout.Sync()
	}
}
