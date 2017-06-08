package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"net/url"
	"strings"
)

//	Set up our flags
var (
	baseURL     = flag.String("apiUrl", "http://localhost:4040", "The base url for the qbitTorrent UI website")
	torrentPath = flag.String("file", "", "The full path to the torrent to remove")
	torrentHash = flag.String("hash", "", "The hash for the torrent to remove")
)

func main() {
	//	Parse the command line for flags:
	flag.Parse()

	hashToLookFor := *torrentHash

	//	If we don't have a file and we don't have a hash, get out:
	if *torrentPath == "" && hashToLookFor == "" {
		log.Fatal("[ERROR] No file or hash specified")
	}

	//	If we don't have a hash, we need to look for
	//	the list of torrents we have and try to find a
	//	match using the name
	if hashToLookFor == "" {
		//	Query to get all torrents
		torrents := []Torrent{}
		res, err := http.Get(*baseURL + "/query/torrents")
		if res != nil {
			defer res.Body.Close()
		}
		if err != nil {
			log.Fatalf("[ERROR] %v", err)
		}

		//	Decode the return object
		if err = json.NewDecoder(res.Body).Decode(&torrents); err != nil {
			log.Fatalf("[ERROR] %v", err)
		}

		//	Loop through and try to find our torrent
		for _, torrent := range torrents {
			//	If we can find the torrent name in the file string, we have a match
			if strings.Contains(*torrentPath, torrent.Name) {
				hashToLookFor = torrent.Hash
				break
			}
		}

		//	If we can't find it, emit an error and exit
		if hashToLookFor == "" {
			log.Fatalf("[ERROR] Can't find the torrent hash for the file %v", *torrentPath)
		}
	}

	//	Build the request:
	form := url.Values{}
	form.Add("hashes", hashToLookFor)

	//	Delete the torrent
	_, err := http.Post(*baseURL+"/command/deletePerm", "application/x-www-form-urlencoded", strings.NewReader(form.Encode()))
	if err != nil {
		log.Fatalf("[ERROR] %v", err)
	}
}
