package utils

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var cellSeparator = regexp.MustCompile(`\|\|?`)
var headerCellSeparator = regexp.MustCompile(`!!?`)
var link = regexp.MustCompile(`\[\[[^|]+\|([^|])+\]\]`)
var br = regexp.MustCompile(`(\w+/)</br>(\w+)`)
var pre = regexp.MustCompile(`<pre>(.*)</pre>`)
var code = regexp.MustCompile(`<code>(.*)</code>`)

func ConvertWIKIToMarkdown(sourcefile string, targetfile string) {
	source, err := os.Open(sourcefile)
	if err != nil {
		panic(err)
	}
	defer func(source *os.File) {
		err := source.Close()
		if err != nil {
			panic(err)
		}
	}(source)

	target, err := os.Create(targetfile)
	if err != nil {
		panic(err)
	}

	defer func(target *os.File) {
		err := target.Close()
		if err != nil {
			panic(err)
		}
	}(target)

	scanner := bufio.NewScanner(source)
	writer := bufio.NewWriter(target)
	fmt.Fprintln(writer, "---")
	fmt.Fprintln(writer, "cssclasses: tablepage, tablenowrap, hideproperties, tablegray")
	fmt.Fprintln(writer, "---")

	for scanner.Scan() {
		line := scanner.Text()

		if isWIKITableLine(line) {
			table := buildTable(scanner)
			for _, t := range table {
				_, err := fmt.Fprintln(writer, t)
				if err != nil {
					panic(err)
				}
			}
		} else {
			_, err := fmt.Fprintln(writer, line)
			if err != nil {
				panic(err)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	if err := writer.Flush(); err != nil {
		panic(err)
	}
}

func isWIKITableLine(line string) bool {
	return strings.HasPrefix(line, "{|class=\"wikitable")
}

func buildTable(scanner *bufio.Scanner) []string {
	table := make([]string, 0)
	header := make([]string, 0)
	buildHeader := false
	row := make([]string, 0)

	table = append(table, "<table>")

	for scanner.Scan() {
		line := scanner.Text()

		if isWIKITableLine(line) {
			subTable := buildTable(scanner)
			row[len(row)-1] = row[len(row)-1] + strings.Join(subTable, "")
		} else if line == "|-" {
			if len(header) > 0 && !buildHeader {
				table = append(table, "<tr>")
				for _, h := range header {
					table = append(table, "<th>"+h+"</th>")
				}
				table = append(table, "</tr>")

				buildHeader = true
			} else {
				table = append(table, "<tr>")
				for _, r := range row {
					if pre.Match([]byte(r)) {
						r = pre.ReplaceAllString(r, "$1")
					}
					if code.Match([]byte(r)) {
						r = code.ReplaceAllString(r, "$1")
					}
					table = append(table, "<td>"+r+"</td>")
				}
				table = append(table, "</tr>")

				row = make([]string, 0)
			}
		} else if line == "|}" {
			if len(row) > 0 {
				table = append(table, "<tr>")
				for _, r := range row {
					table = append(table, "<td>"+r+"</td>")
				}
				table = append(table, "</tr>")
			}
			table = append(table, "</table>")
			return table
		} else if strings.HasPrefix(line, "|") {
			if link.Match([]byte(line)) {
				line = link.ReplaceAllString(line, "$2")
			}

			if br.Match([]byte(line)) {
				line = br.ReplaceAllString(line, "$1$2")
			}

			line = strings.ReplaceAll(line, "<title>", "title")

			r := cellSeparator.Split(strings.Trim(line, "|"), -1)
			row = append(row, r...)
		} else if strings.HasPrefix(line, "!") {
			if !buildHeader {
				h := headerCellSeparator.Split(strings.Trim(line, "!"), -1)
				header = append(header, h...)
			} else {
				r := cellSeparator.Split(strings.Trim(line, "!"), -1)
				row = append(row, r...)
			}
		} else {
			if link.Match([]byte(line)) {
				line = link.ReplaceAllString(line, "$2")
			}

			if br.Match([]byte(line)) {
				line = br.ReplaceAllString(line, "$1$2")
			}

			if pre.Match([]byte(line)) {
				line = pre.ReplaceAllString(line, "$1")
			}

			line = strings.ReplaceAll(line, "<title>", "title")
			row[len(row)-1] = row[len(row)-1] + line
		}
	}

	return table
}
