# Chapter 8. Files

In this chapter we're going to focus on working with files. In particular, we'll  focus on the management of local files 
on your laptop or computer. 

Think of all the mundane tasks you typically do with file management; backing up files, removing large files, renaming
a bunch of files, and selectively backing up files. 

Let's get going.

## Back Up Files to a Zip

```go
package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("A path must be provided such as './path/to/folder'")
	}
	if len(os.Args) > 1 {
		fP, err := filepath.Abs(os.Args[1])
		if err != nil {
			fmt.Errorf("ERR %s", err)
		}

		var zipFileName string
		n := 1
		for true {
			zipFileName = fP + "_" + strconv.Itoa(n) + ".zip"
			if _, err := os.Stat(zipFileName); os.IsNotExist(err) {
				break
			}
			n++
		}
		zipDir(os.Args[1], zipFileName)
	}
}

func zipDir(source, target string) error {
	zipFile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	archive := zip.NewWriter(zipFile)
	defer archive.Close()

	info, err := os.Stat(source)
	if err != nil {
		return nil
	}

	var baseDir string
	if info.IsDir() {
		baseDir = filepath.Base(source)
	}

	filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		if baseDir != "" {
			header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, source))
		}

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = io.Copy(writer, file)
		return err
	})

	return err
}

```

## Find Files taking up lots of Space

```go
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


```



## Rename Dates on files

```go
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		reference := ``
		fmt.Println(reference)
	}
	if len(os.Args) > 1 {
		regStr, _ := regexp.Compile(`(\d\d|\d)-(\d\d|\d)-(\d\d\d\d)`)
		dir, err := os.Open(os.Args[1])

		if err != nil {
			return
		}
		defer dir.Close()
		fileInfos, err := dir.Readdir(-1)
		if err != nil {
			return
		}
		for _, fi := range fileInfos {
			strFile := fi.Name()
			fmt.Println(strFile)
			if regStr.MatchString(strFile) {
				findRegexp := regStr.FindStringSubmatch(strFile)
				findRegexp[1], findRegexp[2] = findRegexp[2], findRegexp[1]
				var newDateInFilename string
				for i := 1; i < len(findRegexp); i++ {
					newDateInFilename += findRegexp[i]
					if i != len(findRegexp)-1 {
						newDateInFilename += "."
					}
				}

				strNewFile := strings.Replace(strFile, findRegexp[0], newDateInFilename, -1)

				filePath, err := filepath.Abs(fi.Name())
				if err != nil {
					return
				}
				newFilePath := strings.Replace(filePath, strFile, strNewFile, -1)
				fmt.Println("old file path", filePath)
				fmt.Println(newFilePath)
				os.Rename(filePath, newFilePath)

			}
		}
	}
}

```

## Conclusion