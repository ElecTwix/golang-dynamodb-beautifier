package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	readFile, err := os.Open("raw.txt")
	data, err := os.Create("data.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	var newarr []string = []string{}

	totalsize := len(fileLines) - 1

	for index, line := range fileLines {

		line = strings.TrimSpace(line)
		if check(line) {
			continue
		}
		line = strings.Trim(line, ",")
		arr := strings.Split(line, ":")
		arr[1] = strings.TrimSpace(arr[1])

		if check2(arr[1]) {
			continue
		}

		line = editlayout(arr[0], arr[1])

		if index == totalsize {
			line = strings.TrimRight(line, ",")
		}

		newarr = append(newarr, line)
		data.WriteString(line + "\n")
	}

	data.Close()

}

func editlayout(index string, value string) string {

	str := fmt.Sprintf("\"%s\": { \"S\": \"%s\"},", index, value)

	return str

}

func check(str string) bool {
	if len(str) < 1 {
		return true
	}
	if str[0] == '/' {
		return true
	}
	return false
}

func check2(str string) bool {
	size := len(str)
	if size < 2 && size > 0 {
		if str[0] == '0' {
			return true
		}
	}
	return false
}

func start(str []string) []string {

	return str
}
