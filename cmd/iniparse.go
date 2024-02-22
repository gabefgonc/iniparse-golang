package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gabefgonc/iniparse-golang/pkg/iniparse"
)

func main() {
	switch len(os.Args) {
	case 1:
		fmt.Fprintln(os.Stderr, "no filename provided")
		os.Exit(1)
	case 2:
		printAllSections()
	default:
		makeQuery()
	}
}

func printAllSections() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	result, err := iniparse.Parse(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	for section, items := range result {
		fmt.Printf("Section %s:\n", section)
		for key, value := range items {
			fmt.Printf("    %s == %s \n", key, value)
		}
	}
}

func makeQuery() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	result, err := iniparse.Parse(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	query := os.Args[2]
	if len(strings.Split(query, ".")) == 2 {
		value, err := iniparse.QueryItem(result, query)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		fmt.Println(value)
		return
	}
	items, err := iniparse.QuerySection(result, os.Args[2])
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	for key, value := range items {
		fmt.Printf("%s == %s\n", key, value)
	}
}
