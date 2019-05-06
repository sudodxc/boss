package engine

import (
	"city/fetcher"
	"log"
)

func Run(seeds ...Request) { //shixian types struct request duilie
	var requests []Request    //
	for _, r := range seeds { //for it
		requests = append(requests, r) //append it
	}
	for len(requests) > 0 { //if it not nil
		r := requests[0]        //get first
		requests = requests[1:] //
		log.Printf("Fetching %s:\n", r.Url)
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetch:error"+"fetching url:%s:%v", r.Url, err)
			continue
		}
		parseResults := r.ParseFunc(body)
		requests = append(requests, parseResults.Requests...) //make your all request input requests

		for _, item := range parseResults.Items {
			log.Printf("Got items %v", item)
		}
	}
}
