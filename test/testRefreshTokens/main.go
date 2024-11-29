package main

import (
	"fmt"
	"log"
	"net/http"
)

var ( // change it after testGetTokens!
	access  = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiI1ODg0ODFlNS03ZGY5LTQ3NDQtODJjNi05OTBkMjk2NmM1ODhAMTcyLjE5LjAuMSIsImV4cCI6MTczNjQ5NjA4M30.rp9UmPut92-QbnV8hKGAYAVCaZpLGmMXesvgspmoElbQzOs932jXZyc0xo7RKLSEMBQRU7viWr8UGdI-kkiPImnQWLy4IvgI9oUj6AL2WcggZqU6pTvwUqNvE7VSjqHi5-M1gUg-SQOtNUNVjLVyJD0znGXw2Ukty2IpzJLbbQgEzCKPH8JwFrOmC0XLHo3krRNblNXOt8OZdl1b_FLum1O9bbcTH41xN5A4gm0PhscIP3c4IwGiP3fgrXA0zIgpiGKERWjI_n-P-k375tmEHgv3p7KWCEBMyMTxJQHv_d7EZ8XA-U_CWoobkDwRj1zpsGEE6BcEtBJkJkGGFPXvPA"
	refresh = "MGRkNzkxMjMtZGEyYS00ZGUzLThmMjQtOWFkNTJlYWMxZDU2"
)

func main() {
	client := &http.Client{}
	request, err := http.NewRequest("GET", "http://localhost:8082/refreshtokens", nil)
	request.Header.Add("X-Access-Token", access)
	request.Header.Add("X-Refresh-Token", refresh)
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
