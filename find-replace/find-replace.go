package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

var (
	matchString   string
	replaceString string
)

func main() {
	directory := flag.String("path", "", "Path to scan for files")
	match := flag.String("match", "", "String to match")
	replace := flag.String("replace", "", "Replace matched string with this value")
	flag.Parse()

	if *directory == "" {
		log.Println("Usage:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	matchString = *match
	replaceString = *replace

	err := filepath.Walk(*directory, visit)
	if err != nil {
		panic(err)
	}
}

func visit(path string, fi os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if !!fi.IsDir() {
		return nil
	}

	read, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	log.Println(path)

	re := regexp.MustCompile(`(?i)` + matchString)

	contents := re.ReplaceAllString(string(read), replaceString)

	err = ioutil.WriteFile(path, []byte(contents), 0)
	if err != nil {
		panic(err)
	}

	return nil
}
