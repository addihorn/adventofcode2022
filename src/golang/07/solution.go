package main

import (
	aocfolders "example/hello/src/golang/07/aocfolder"
	"example/hello/src/golang/aocutils"
	"fmt"
	"strconv"
	"strings"
)

var smallestFolderToDelete *aocfolders.Directory

func main() {

	currentFolder := aocfolders.NewSystem()
	commandsAndOutput := aocutils.ReadInput("input.txt")

	for _, cmdLine := range commandsAndOutput {
		if cmdLine[0] == '$' {
			if len(strings.Split(cmdLine, " cd ")) > 1 {
				currentFolder = currentFolder.ChangeDir(strings.Split(cmdLine, " cd ")[1])
			}
		} else {
			if cmdLine[0:3] == "dir" {
				currentFolder.MkDir(cmdLine[4:len(cmdLine)])
			} else {
				fileDetails := strings.Split(cmdLine, " ")
				fileSize, _ := strconv.ParseInt(fileDetails[0], 10, 64)
				currentFolder.TouchFile(fileDetails[1], fileSize)
			}
		}
	}

	currentFolder = currentFolder.ChangeDir("/")
	fmt.Println("Size of all Files:", currentFolder.FileSizes)

	fmt.Println("Size of directory <= 100000:", getSizeOfSubDirs(*currentFolder))

	currentFolder = currentFolder.ChangeDir("/")
	systemSize := 70000000
	minimalSpaceForUpdate := 30000000
	fmt.Println("Space to allocate:", (int64(systemSize-minimalSpaceForUpdate)-currentFolder.FileSizes)*-1)
	smallestFolderToDelete = currentFolder
	findDirectoryToDelete(*currentFolder, (int64(systemSize-minimalSpaceForUpdate)-currentFolder.FileSizes)*-1)
	fmt.Println("Smallest Directory to Remove:", smallestFolderToDelete)
}

func findDirectoryToDelete(directory aocfolders.Directory, spaceNecessary int64) {
	for _, subDir := range directory.Subdirectories {
		findDirectoryToDelete(*subDir, spaceNecessary)
	}
	if directory.FileSizes < spaceNecessary {
		return
	}

	if directory.FileSizes < smallestFolderToDelete.FileSizes {
		*smallestFolderToDelete = directory
	}

}

func getSizeOfSubDirs(directory aocfolders.Directory) int64 {

	var directorySize int64
	directorySize = 0
	for _, subDir := range directory.Subdirectories {

		directorySize += getSizeOfSubDirs(*subDir)
		if subDir.FileSizes <= 100000 {
			fmt.Println(subDir.Path, subDir.FileSizes, directorySize)
			directorySize += subDir.FileSizes
		}
	}

	return directorySize
}
