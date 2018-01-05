package filewriter

import (
	"os"
	"fmt"
	"bufio"
)

type FileWriter struct {
}

func(fw *FileWriter) Write(content string) {
	destination := "C:\\Users\\steve\\Desktop\\chord-melody-data.txt"

	fileHandle, _ := os.Create(destination)
	writer := bufio.NewWriter(fileHandle)
	defer fileHandle.Close()

	fmt.Fprintln(writer, content)
	writer.Flush()
}