package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var inputDir string
var outputDir string

const (
	inputUsage  string = "directory containing sortable files"
	outputUsage string = "directory where to create directories / place files"
)

func init() {
	flag.StringVar(&inputDir, "i", "", inputUsage+" (shorthand)")
	flag.StringVar(&inputDir, "inputdir", "", inputUsage)
	flag.StringVar(&outputDir, "o", "", outputUsage+" (shorthand)")
	flag.StringVar(&outputDir, "outputdir", "", outputUsage)
}

func main() {
	flag.Parse()
	// take first argument as input and output dir
	// shortuct for lazy people aka me
	if flag.NArg() > 0 {
		if inputDir == "" {
			inputDir = flag.Args()[0]
		}
	}

	if inputDir == "" {
		log.Fatal("No input directory specified")
	}

	// check if inputDir & outputDir exist on disk
	if _, err := os.Stat(inputDir); os.IsNotExist(err) {
		log.Fatal("Input direcotry doesn't exist")
	}

	if outputDir != "" {
		if _, err := os.Stat(outputDir); os.IsNotExist(err) {
			// if outputDir doesn't exist on the disk, attempt to create it
			err := os.Mkdir(outputDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	config := loadConfig()

	files, err := ioutil.ReadDir(inputDir)
	if err != nil {
		log.Fatal(err)
	}

	// sorting
	for _, dir := range config.Directories {
		// set working path which contains files to sort
		var wk string
		if outputDir != "" {
			wk = outputDir
		} else {
			wk = inputDir
		}
		od := filepath.Join(wk, dir.Name)

		// create directory where to place sorted files
		_ = os.Mkdir(od, os.ModePerm)

		for _, filename := range files {
			for _, ext := range dir.Ext {
				if filename.IsDir() {
					continue
				}

				fn := filename.Name()
				si := len(fn) - len(ext)
				// check if slice bounds are out of range
				if si <= 0 {
					continue
				}

				fext := fn[si:]
				if fext != ext {
					continue
				}
				// if file ext matches, move the file
				err := os.Rename(filepath.Join(wk, fn), filepath.Join(od, fn))
				if err != nil {
					log.Fatal(err)
				}
				break
			}
		}
	}
	// make some folder resolutions, so it doesn't print out .
	fmt.Printf("Files within %s sorted!\n", inputDir)
}
