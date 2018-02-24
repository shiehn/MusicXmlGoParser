package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"github.com/MusicXmlGoParser/xmlparser"
	"github.com/MusicXmlGoParser/xmlparser/filewriter"
)

func main() {
	encode := flag.Bool("encode", false, "a bool")
	xmlDir := flag.String("dir", "", "")
	flag.Parse()

	if *xmlDir == "" {
		fmt.Fprint(os.Stderr, "\nUsage Example: 'go run main.go -encode=true -dir=/path/to/dataset'\n")
		os.Exit(1)
	}

	*xmlDir = formatPath(xmlDir)
	fmt.Println("DIRECTORY: " + *xmlDir)

	output := ""
		fileNames, err := ioutil.ReadDir(*xmlDir)
		if err != nil {
			fmt.Printf("ERROR opening: %v \n", *xmlDir)
			panic(err)
		}

		if len(fileNames) < 1 {
			fmt.Printf("NO FILES FOUND IN %v\n", *xmlDir)
		}

		fmt.Println("WTF")

		if !hasDirectories(fileNames) {
			fmt.Printf("The Dataset dir must contain a dir for each song which inturn contains .xml files for each key transpositions\n")
			os.Exit(1)
		}

		for _, name := range fileNames {
			if shouldIgnore(name.Name()) {
				fmt.Println("Ignoring unsuportted file: " + name.Name())
				continue
			}

			if name.IsDir() {
				fmt.Println("*****************************************************")
				fmt.Println("*****************************************************")
				fmt.Println("*****************************************************")
				fmt.Printf("DIRECTORY: %v \n \n", *xmlDir+"\\"+name.Name())

				innerFiles, err := ioutil.ReadDir(*xmlDir + "\\" + name.Name())
				if err != nil {
					panic(err)
				}

				for _, songName := range innerFiles {
					if songName.IsDir() {
						fmt.Println("")
						fmt.Printf("WARNING: '%v' is a directory. SKIPPING! \n\n", songName.Name())
						continue
					}

					fmt.Printf("FILE: %v \n", *xmlDir+"\\"+name.Name()+"\\"+songName.Name())
					musicXML, err := ioutil.ReadFile(*xmlDir + "\\" + name.Name() + "\\" + songName.Name())
					var xmlDoc xmlparser.MXLDoc
					err = xml.Unmarshal(musicXML, &xmlDoc)
					if err != nil {
						fmt.Printf("ERROR PARSING MUSIC XML: %v \n", err)
						os.Exit(1)
					}

					parser := xmlparser.Parser{
						MusicXml: xmlDoc,
					}

					audioStr, err := parser.Parse()
					output = output + audioStr

					if *encode {
						if err != nil {
							panic(err)
						}

						fmt.Printf("OUTPUT: %s \n\n", audioStr)
					} else {
						if err == nil {
							fmt.Printf("LOOKS GOOD! \n")
							continue
						}

						fmt.Printf("ERROR: %v \n", err)
					}
				}
		}
	}

	fw := filewriter.FileWriter{}
	fw.Write(output)
}

func shouldIgnore(fileName string) bool {
	ignoreList := []string{".git", ".idea"}

	for _, listItem := range ignoreList {
		if strings.Contains(listItem, fileName) {
			return true
		}
	}

	return false
}

func formatPath(xmlDir *string) string {

	fmt.Println("VVVVV:" + *xmlDir)

	if *xmlDir == "" || len(*xmlDir) < 1 {
		message := fmt.Sprintf("ERROR malformed directory: %v \n", xmlDir)
		panic(message)
	}

	return strings.TrimSuffix(*xmlDir, "\\")
}

func hasDirectories(fileNames []os.FileInfo) bool {
	count := 0
		for _, name := range fileNames {
			if name.IsDir() {
				count++
			}
		}
	return count > 0
}