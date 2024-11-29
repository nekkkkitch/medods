package main

import (
	"fmt"
	"log"
	"net/http"
)

var ( // change it after testGetUUID!
	userID = "588481e5-7df9-4744-82c6-990d2966c588"
)

func main() {
	client := &http.Client{}
	request, err := http.NewRequest("GET", fmt.Sprintf("http://localhost:8082/tokens?id=%v", userID), nil)
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("X-Access-Token:", resp.Header.Get("X-Access-Token"))
	fmt.Println("X-Refresh-Token:", resp.Header.Get("X-Refresh-Token"))
}
