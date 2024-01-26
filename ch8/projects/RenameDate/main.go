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
