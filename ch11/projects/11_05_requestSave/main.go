package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	pwdDir, _ := os.Getwd()
	fLog, err := os.OpenFile(pwdDir+"/log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0640)
	check(err, fLog)

	fSave, err := os.OpenFile(pwdDir+"/save.txt", os.O_CREATE|os.O_WRONLY, 0600)
	check(err, fLog)

	url := "http://www.gutenberg.org/cache/epub/1112/pg1112.txt"

	resp, err := http.Get(url)
	check(err, fLog)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	check(err, fLog)

	fSave.Write(body)
}

// err check to log
func check(err error, fLog *os.File) {
	if err != nil {
		log.Fatalln(err)
	}
	log.SetOutput(io.MultiWriter(fLog, os.Stdout))
}
