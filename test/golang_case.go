package main

import (
	// "fmt"
	// "io"

	"log"

	"github.com/levigross/grequests"
)

func Example_basicGet() {
	// This is a very basic GET request
	resp, err := grequests.Get("http://httpbin.org/get", nil)
	defer resp.Close() //显式定义关闭响应，对于用协程方法来请求的话很重要
	if err != nil {
		log.Println(err)
	}

	if resp.Ok != true {
		log.Println("Request did not return OK")
	}

	log.Println(resp.String())
}

func main() {
	Example_basicGet()
}
