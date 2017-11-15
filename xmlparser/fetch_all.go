package xmlparser

import (
	"io/ioutil"
	"log"
	"fmt"
	"os"
	"strings"
)

func FetchAll() []os.FileInfo {
	files, err := ioutil.ReadDir("C:\\GoWorkspace\\src\\github.com\\MusicXmlGoParser\\testassets\\")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}

	//empty := []string{"one", "two", "three"}
	return filterForXml(files)
}

func filterForXml(files []os.FileInfo) (returnArray []os.FileInfo) {
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".xml"){
			returnArray = append(returnArray, f)
		}
	}
	return
}
