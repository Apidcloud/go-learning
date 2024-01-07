package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://amazon.com",
		"http://stackoverflow.com",
	}

	c := make(chan string)

	for _, link := range links {
		fmt.Println("preparing")

		go checkLinkStatus(link, c)

		//fmt.Printf("Status of %s is %d\n", link, status)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLinkStatus(link, c)
		}(l)
	}

	// infinite loop, where each iteration is a blocking call
	// waiting for a message from the channel.
	/* for l := range c {
		go checkLinkStatus(l, c)
	} */

	// equivalent to the following:
	// infite loop. But <-c is blocking, so it's not like the for
	// is contantly iterating. It only moves to the next iteration
	// when we receive a message (blocking call) from the channel
	/* for {
		go checkLinkStatus(<-c, c)
	} */
}

func checkLinkStatus(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		c <- link
		return
	}

	fmt.Println(link, "seems to be up")
	c <- link
}
