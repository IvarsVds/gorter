package main

import (
	"flag"
	"fmt"
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
		fmt.Println("No input directory specified")
		os.Exit(1)
	}

	// check if inputDir & outputDir exist on disk
	if _, err := os.Stat(inputDir); os.IsNotExist(err) {
		fmt.Println("Input direcotry doesn't exist")
		os.Exit(1)
	}

	if outputDir != "" {
		if _, err := os.Stat(outputDir); os.IsNotExist(err) {
			// if outputDir doesn't exist on the disk, attempt to create it
			err := os.Mkdir(outputDir, os.ModePerm)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	}

	config := loadConfig()

	files, err := os.ReadDir(inputDir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// sorting
	for k, v := range config.Directories {
		// set working path which contains files to sort
		var wk string
		if outputDir != "" {
			wk = outputDir
		} else {
			wk = inputDir
		}
		od := filepath.Join(wk, k)

		// create directory where to place sorted files
		_ = os.Mkdir(od, os.ModePerm)

		for _, filename := range files {
			for _, ext := range v {
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
					fmt.Println(err)
					os.Exit(1)
				}
				break
			}
		}
	}

	fmt.Println("Files sorted!")
}
