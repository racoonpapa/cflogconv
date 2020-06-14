package main

import (
	"bufio"
	"strings"

	"compress/gzip"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"flag"
)

var header string = "date,time,x-edge-location,sc-bytes,c-ip,cs-method,cs(Host),cs-uri-stem,sc-status,cs(Referer),cs(User-Agent),cs-uri-query,cs(Cookie),x-edge-result-type,x-edge-request-id,x-host-header,cs-protocol,cs-bytes,time-taken,x-forwarded-for,ssl-protocol,ssl-cipher,x-edge-response-result-type,cs-protocol-version,fle-status,fle-encrypted-fields,c-port,time-to-first-byte,x-edge-detailed-result-type,sc-content-type,sc-content-len,sc-range-start,sc-range-end"

func processLine(text string, output *[]string) {
	if strings.HasPrefix(text, "#") {
		return
	}
	s := strings.Fields(text)
	if len(s) == 0 {
		return
	}
	o := fmt.Sprint(strings.Join(s[:], ","))
	*output = append(*output, o)
}

func gunzip(src string) ([]byte, error) {
	f, err := os.Open(src)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	gzf, err := gzip.NewReader(f)
	if err != nil {
		return nil, err
	}
	defer gzf.Close()

	fileContents, err := ioutil.ReadAll(gzf)
	if err != nil {
		return nil, err
	}

	return fileContents, nil
}

func main() {
	inputDir := flag.String("i", "", "input directory path")
	outputFilename := flag.String("o", "output.csv", "output CSV file")

	flag.Parse()

	if *inputDir == "" {
		flag.Usage()
		return
	}

	output := []string{header}

	files, err := ioutil.ReadDir(*inputDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if !f.IsDir() && strings.HasSuffix(f.Name(), ".gz") {
			filePath := filepath.Join(*inputDir, f.Name())
			if content, err := gunzip(filePath); err == nil {
				lines := strings.Split(string(content), "\n")
				for _, line := range lines {
					processLine(line, &output)
				}
			}
		}
	}

	outfile, err := os.OpenFile(*outputFilename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer outfile.Close()

	writer := bufio.NewWriter(outfile)

	writer.WriteString(strings.Join(output, "\n"))
	writer.Flush()

	fmt.Println("Ok, done.")
}
