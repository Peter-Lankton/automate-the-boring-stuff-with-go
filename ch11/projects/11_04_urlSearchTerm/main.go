package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/jackdanger/collectlinks"
)

func main() {

	pwdDir, _ := os.Getwd()
	fLog, err := os.OpenFile(pwdDir+"/log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0640)
	check(err, fLog)

	resp, err := http.Get("https://automatetheboringstuffwithgo.com")
	check(err, fLog)
	defer resp.Body.Close()

	links := collectlinks.All(resp.Body)

	for _, link := range links {
		fmt.Println(link)
	}
}

// err check to log
func check(err error, fLog *os.File) {
	if err != nil {
		log.Fatalln(err)
	}
	log.SetOutput(io.MultiWriter(fLog, os.Stdout))
}
