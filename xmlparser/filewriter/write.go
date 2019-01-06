package filewriter

import (
	"bufio"
	"fmt"
	"os"
)

type FileWriter struct {
}

func (fw *FileWriter) Write(content string, destination string) { 

	fileHandle, _ := os.Create(destination)
	writer := bufio.NewWriter(fileHandle)
	defer fileHandle.Close()

	fmt.Fprintln(writer, content)
	writer.Flush()
}
