package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"bytes"
)

func main() {
	var phone string = os.Args[1]
	var subject string = os.Args[2]
	var body string = os.Args[3]

	_ = subject

	url := getUrl(phone, body)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}

func getUrl(phone, body string) string {
	buf := bytes.Buffer{}
	buf.WriteString("http://10.10.11.241/sms.php?phone=")
	buf.WriteString(phone)
	buf.WriteString("&content=")
	buf.WriteString(body)
	return buf.String()
}

func convToStr(body []string) (str string) {
	for _, v:= range body {
		str += v
	}
	return
}
