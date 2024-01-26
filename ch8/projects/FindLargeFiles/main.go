package main

import (
	"fmt"
	"io"
	"os"
)

const LG_FILE_SIZE = 104857600

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Provide the directory to search")
	}
	if len(os.Args) > 1 {
		FindLargeFiles(os.Args[1])
	}
}

func FindLargeFiles(dir string) {
	dh, err := os.Open(dir)
	if err != nil {
		return
	}
	defer dh.Close()
	for {
		fis, err := dh.Readdir(10)
		if err == io.EOF {
			break
		}
		for _, fi := range fis {
			if fi.Size() > LG_FILE_SIZE {
				fmt.Printf("%s/%s\t%d MB\n", dir, fi.Name(), fi.Size()/1048576)
			}

			if fi.IsDir() {
				FindLargeFiles(dir + "/" + fi.Name())
			}
		}
	}
}
