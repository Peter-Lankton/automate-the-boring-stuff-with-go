package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"

	"strings"

	"strconv"

	"sort"

	"github.com/opesun/goquery"
)

func main() {
	pwdDir, _ := os.Getwd()
	fLog, err := os.OpenFile(pwdDir+"/log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0640)
	check(err, fLog)

	url := `http://xkcd.com`

	err = os.Mkdir(pwdDir+"/xkcd", 0775)
	check(err, fLog)

	resp, err := http.Get(url)
	log.Printf("URL:	%v", url)
	check(err, fLog)
	defer resp.Body.Close()

	x, err := goquery.Parse(resp.Body)
	check(err, fLog)

	var urlInt []int
	regLastLink, _ := regexp.Compile(`\/[0-9]{1,}\/`)
	for _, i := range x.Find("a").Attrs("href") {
		if regLastLink.MatchString(i) {
			a, err := strconv.Atoi(strings.Trim(i, "/"))
			check(err, fLog)
			urlInt = append(urlInt, a)
		}
	}
	sort.Ints(urlInt)
	fmt.Println(urlInt[len(urlInt)-1])

	for k := urlInt[len(urlInt)-1] + 1; k >= 1; k-- {
		resp, err := http.Get(url + "/" + strconv.Itoa(k))
		log.Printf("URL: 	%v", url+"/"+strconv.Itoa(k))
		check(err, fLog)
		defer resp.Body.Close()

		x, err := goquery.Parse(resp.Body)
		check(err, fLog)

		regStr, _ := regexp.Compile(`comics`)
		for _, i := range x.Find("img").Attrs("src") {
			if regStr.MatchString(i) {

				nameFile := strings.Split(i, "/comics/")

				name := pwdDir + "/xkcd/" + nameFile[1]
				fSave, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0640)
				check(err, fLog)

				respImg, err := http.Get("http:" + i)
				log.Printf("http	%v", "http:"+i)
				check(err, fLog)

				defer respImg.Body.Close()

				bodyImg, err := ioutil.ReadAll(respImg.Body)
				check(err, fLog)

				log.Printf("name:	%v", name)
				fSave.Write(bodyImg)
			}
		}
	}

	log.SetOutput(io.MultiWriter(fLog, os.Stdout))
}

// err check to log
func check(err error, fLog *os.File) {
	if err != nil {
		log.Println(err)
	}
	log.SetOutput(io.MultiWriter(fLog, os.Stdout))
}
