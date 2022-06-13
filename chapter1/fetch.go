package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		//if !strings.HasPrefix(url, "https://") {
		//	url = "https://" + url
		//}
		resp, err := http.Get(url)
		//fmt.Printf("Status Code %d\n", resp.StatusCode)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		writer := os.Stdout
		reader := io.Reader(resp.Body)
		written, err := io.Copy(writer, reader)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%d", written)
	}
}
