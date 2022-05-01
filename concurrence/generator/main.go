package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func title(urls ...string) <-chan string {
	channel := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println(err)
				return
			}
			html, _ := ioutil.ReadAll(resp.Body)
			title, _ := regexp.Compile("<title>(.*?)<\\/title>")
			channel <- title.FindStringSubmatch(string(html))[1]
		}(url)
	}
	return channel
}

// go func(ur string) {
// 	fmt.Println("url:", ur)
// 	response, resError := http.Get(ur)
// 	html, err := ioutil.ReadAll(response.Body)

// 	if resError != nil {
// 		fmt.Println(resError)
// 	}

// 	if err != nil {
// 		panic(err)
// 	}

// 	title, _ := regexp.Compile("<title>(.*?)<\\/title>")
// 	channel <- title.FindStringSubmatch(string(html))[1]
// }(url)
func main() {
	t1 := title("https://www.baidu.com", "https://www.google.com")
	t2 := title("https://www.gmail.com", "https://www.amazon.com")

	fmt.Printf("%s | %s\n", <-t1, <-t2)
	fmt.Printf("%s | %s\n", <-t1, <-t2)
}
