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
			html, htmlErr := ioutil.ReadAll(resp.Body)

			if htmlErr != nil {
				channel <- "Error reading HTML"
			}

			title, titleErr := regexp.Compile("<title>(.*?)<\\/title>")

			if titleErr != nil {
				channel <- "Error parsing HTML"
			}

			channel <- title.FindStringSubmatch(string(html))[1]
		}(url)
	}
	return channel
}
