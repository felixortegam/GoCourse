package main

import (
	"io"
	"net/http"
	"os"
	"sync"
)

func main() {
	sites := []string{
		"https://www.google.com",
		"https://drive.google.com",
		"https://maps.google.com",
		"https://hangouts.google.com",
	}

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		for _, site := range sites {
			func(site string) {
				res, err := http.Get(site)
				if err != nil {
				}

				io.WriteString(os.Stdout, res.Status+"\n")
			}(site)
		}

		wg.Done()
	}()

	wg.Wait()
}
