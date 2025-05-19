package main

import (
	"log"
	"net/http"

	"codebyte.cy/pressminer/src/app/client"
)

func main() {
	handler := http.HandlerFunc(client.CrawlerServer)
	log.Fatal(http.ListenAndServe(":5000", handler))
}
