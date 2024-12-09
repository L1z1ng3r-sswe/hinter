package main

import (
	"context"
	"fmt"
	"net/http"

	"golang.org/x/sync/errgroup"
)

const (
	gosLimit = 100
)

func main() {
	urls := []string{
		"https://jalap.",
		"https://youtube.com",
		"https://www.google.com",
		"https://wildberries.com",
		"https://discord.com",
		"https://channel.com",
		"https://casio.com",
	}

	urlsCh := make(chan string, len(urls))

	// Create an errgroup with context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	// Launch workers
	for i := 0; i < gosLimit; i++ {
		g.Go(func() error {
			for {
				select {
				case <-ctx.Done():
					// Stop processing if the context is canceled
					return ctx.Err()
				case url, ok := <-urlsCh:
					if !ok {
						return nil // Channel is closed, exit
					}

					fmt.Println("Fetching: ", url)
					if err := fetch(url); err != nil {
						cancel() // Cancel all other goroutines
						return err
					}
					fmt.Println("Fetched: ", url)
				}
			}
		})
	}

	// Feed URLs into the channel
	for _, url := range urls {
		urlsCh <- url
	}
	close(urlsCh)

	// Wait for all workers to finish
	if err := g.Wait(); err != nil {
		fmt.Println("Error during the fetch: ", err)
	} else {
		fmt.Println("URLs fetching finished")
	}
}

func fetch(url string) error {
	_, err := http.Get(url)
	return err
}
