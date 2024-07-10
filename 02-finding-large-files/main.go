package main

import (
	"container/list"
	"fmt"
	"log"
	"os"
	"path"
)

type FileNode struct {
	FilePath string
	Info     os.FileInfo
}

func insertSorted(fileList *list.List, fileNode FileNode) {
	if fileList.Len() == 0 {
		fileList.PushFront(fileNode)
		return
	}
	for element := fileList.Front(); element != nil; element = element.Next() {
		if fileNode.Info.Size() < element.Value.(FileNode).Info.Size() {
			fileList.InsertBefore(fileNode, element)
			return
		}
	}
	fileList.PushBack(fileNode)
}

func getFilesInDirRecursivelyBySize(fileList *list.List, dpath string) {
	dirFiles, err := os.ReadDir(dpath)
	if err != nil {
		log.Println("Error reading directory:", err.Error())
	}

	for _, dirFile := range dirFiles {
		dirInfo, _ := dirFile.Info()
		fullpath := path.Join(dpath, dirFile.Name())
		if dirFile.IsDir() {
			getFilesInDirRecursivelyBySize(fileList, path.Join(dpath, dirFile.Name()))
		} else if dirInfo.Mode().IsRegular() {
			insertSorted(fileList, FileNode{FilePath: fullpath, Info: dirInfo})
		}
	}
}

func humanSize(size int64) string {
	return fmt.Sprintf("%.2f", float64(size/1024/1024))
}

func main() {
	fileList := list.New()
	getFilesInDirRecursivelyBySize(fileList, "/")

	for element := fileList.Front(); element != nil; element = element.Next() {
		fmt.Printf("%s\n", humanSize(element.Value.(FileNode).Info.Size()))
		fmt.Printf("%s\n", element.Value.(FileNode).FilePath)
	}
}
