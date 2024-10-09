package main

// project load test api <endpoint>/uuid
// increment by 10 every time it is called

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
)

var (
	url   = flag.String("url", "http://localhost:8080/uuid", "url to test")
	inc   = flag.Int("inc", 5, "increment value")
	limit = flag.Int("limit", 10000, "limit")
	wg    sync.WaitGroup
)

func sendRequest(url string) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
}

func main() {
	flag.Parse()

	// simulate requests incrementing by inc flag
	for i := 1; i <= *limit; i += *inc {
		fmt.Printf("Sending %d requests...\n", i)
		for j := 0; j < i; j++ {
			wg.Add(1)
			go sendRequest(*url)
		}
		wg.Wait()
	}

}
