package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func processLine(text string, output *[]string) {
	s := strings.Fields(text)

	if strings.HasPrefix(s[0], "#") {
		s = s[1:]
	}

	o := fmt.Sprint(strings.Join(s[:], ","))
	*output = append(*output, o)
}

func main() {
	var output []string

	infile, err := os.Open("log.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer infile.Close()

	reader := bufio.NewScanner(infile)
	for i := 0; reader.Scan(); i++ {
		if i == 0 {
			continue
		}
		line := reader.Text()
		processLine(line, &output)
	}

	if err := reader.Err(); err != nil {
		log.Fatal(err)
	}

	outfile, err := os.OpenFile("output.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer outfile.Close()

	writer := bufio.NewWriter(outfile)

	for _, data := range output {
		_, _ = writer.WriteString(data + "\n")
	}

	writer.Flush()
	fmt.Println("done.")
}
