package main

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	//	Find the torrent

	//	If we can't find it, emit an error and exit

	//	If we can find it, get the hash
	torrentHash := os.Args[1]

	//	Build the request to delete:
	form := url.Values{}
	form.Add("hashes", torrentHash)
	_, err := http.Post("http://localhost:4040/command/deletePerm", "application/x-www-form-urlencoded", strings.NewReader(form.Encode()))
	if err != nil {
		log.Fatal(err)
	}
}
