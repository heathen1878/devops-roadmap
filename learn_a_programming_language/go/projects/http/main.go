package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	r, err := http.Get("https://google.com")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {

		// byte slice with 99999 elements
		b := make([]byte, 99999)

		// Uses the read function which is part
		// of the Body interface ReadCloser.
		r.Body.Read(b)

		// convert the byte slice to a string
		fmt.Println(string(b))
	}
}
