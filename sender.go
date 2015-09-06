package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	var phone string = os.Args[1]
	var subject string = os.Args[2]
	var body string = os.Args[3]

	_ = subject

	res, err := http.Get("http://10.10.11.241/sms.php?phone="+phone+"&content="+body)
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