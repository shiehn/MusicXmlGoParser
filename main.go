package main

import (
	"io/ioutil"
	"fmt"
	"encoding/xml"
	"github.com/MusicXmlGoParser/xmlparser"
	"flag"
	"os"
	"strings"
	"github.com/MusicXmlGoParser/xmlparser/filewriter"
)

func main() {
	encode := flag.Bool("encode", false, "a bool")
	xmlDir := flag.String("dir", "", "")

	flag.Parse()

	*xmlDir = formatPath(xmlDir)

	fmt.Println("XXXXXXXX: " + *xmlDir)

	output := ""

	if *xmlDir != "" {
		fileNames, err := ioutil.ReadDir(*xmlDir)
		if err != nil {
			fmt.Printf("ERROR opening: %v \n", *xmlDir)
			panic(err)
		}

		for _, name := range fileNames {

			if shouldIgnore(name.Name()) {
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
