package fileio

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

//ReadWebsitesFromFile ...
func ReadWebsitesFromFile() []string {

	rawFile, err := os.Open("../data/websites.txt")
	var websites []string

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	file := bufio.NewReader(rawFile)

	for {
		line, err := file.ReadString('\n')
		line = strings.TrimSpace(line)
		websites = append(websites, line)

		if err == io.EOF {
			break
		} else if err != nil && err != io.EOF {
			fmt.Println(err)
			os.Exit(-1)
		}
	}

	rawFile.Close()
	return websites
}
