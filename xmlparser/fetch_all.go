package xmlparser

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func FetchAll() []os.FileInfo {
	files, err := ioutil.ReadDir("C:\\gocode\\src\\github.com\\MusicXmlGoParser\\testassets\\")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}

	return filterForXml(files)
}

func filterForXml(files []os.FileInfo) (returnArray []os.FileInfo) {
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".xml") {
			returnArray = append(returnArray, f)
		}
	}
	return
}
