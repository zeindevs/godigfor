package main

import (
	"container/list"
	"fmt"
	"log"
	"os"
	"path"
)

type FileNode struct {
	FullPath string
	Info     os.FileInfo
}

func insertSorted(fileList *list.List, fileNode FileNode) {
	if fileList.Len() == 0 {
		fileList.PushFront(fileNode)
		return
	}

	for element := fileList.Front(); element != nil; element = element.Next() {
		if fileNode.Info.ModTime().Before(element.Value.(FileNode).Info.ModTime()) {
			fileList.InsertBefore(fileNode, element)
			return
		}
	}

	fileList.PushBack(fileNode)
}

func GetFilesInDirRecursivelyBySize(fileList *list.List, dpath string) {
	dirFiles, err := os.ReadDir(dpath)
	if err != nil {
		log.Println("Error reading directory:", err.Error())
	}

	for _, dirFile := range dirFiles {
		fullpath := path.Join(dpath, dirFile.Name())
		dirInfo, _ := dirFile.Info()
		if dirFile.IsDir() {
			GetFilesInDirRecursivelyBySize(fileList, path.Join(dpath, dirFile.Name()))
		} else if dirInfo.Mode().IsRegular() {
			insertSorted(fileList, FileNode{FullPath: fullpath, Info: dirInfo})
		}

	}
}

func main() {
	fileList := list.New()
	GetFilesInDirRecursivelyBySize(fileList, "/")

	for element := fileList.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value.(FileNode).Info.ModTime())
		fmt.Printf("%s\n", element.Value.(FileNode).FullPath)
	}
}
