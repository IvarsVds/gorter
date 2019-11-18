package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Name:    "gorter",
		Version: "1.0.0",
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "Ivars Veidelis",
				Email: "ivars.veidelis@protonmail.com",
			},
		},
		Usage:                "Sort and organize files in directories",
		EnableBashCompletion: true,
		Action: func(c *cli.Context) error {
			var inputDir string
			var outDir string

			if c.NArg() > 0 {
				inputDir = c.Args().Get(0)
				outDir = c.Args().Get(1)
			} else {
				log.Fatal("No input directory specified")
			}

			// check if inputDir and outDir exist on disk
			if _, err := os.Stat(inputDir); os.IsNotExist(err) {
				log.Fatal("Input directory doesn't exist")
			}

			if outDir != "" {
				if _, err := os.Stat(outDir); os.IsNotExist(err) {
					// if outDir doesn't exist, attempt to create it
					err := os.Mkdir(outDir, os.ModePerm)
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

			var workDir string
			var outPath string

			// refactor this for max efficiency
			for _, directory := range config.Directories {
				if len(outDir) > 0 {
					workDir = outDir
					outPath = filepath.Join(outDir, directory.Name)
				} else {
					workDir = inputDir
					outPath = filepath.Join(inputDir, directory.Name)
				}

				err := os.Mkdir(outPath, os.ModePerm)
				if err != nil {
					//fmt.Println(err)
				}

				for _, filename := range files {
					for _, ext := range directory.Ext {
						// check for file extension match
						if filename.IsDir() == false {
							fname := filename.Name()
							sInd := len(fname) - len(ext)
							if fname[sInd:] == ext {
								err := os.Rename(filepath.Join(workDir, fname), filepath.Join(outPath, fname))
								if err != nil {
									log.Fatal(err)
								}

								break
							}
						}
					}
				}

			}
			fmt.Println("Files within", inputDir, "sorted!")
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
