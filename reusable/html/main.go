package html

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// Title returns the title of the HTML document.
func Title(urls ...string) <-chan string {
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
