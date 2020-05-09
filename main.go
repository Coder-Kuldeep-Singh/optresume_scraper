package main

import (
	"fmt"
	"log"

	"github.com/optresume/client"
)

func main() {
	resp, err := client.Request("https://google.com/")
	if err != nil {
		log.Println("Proxy Failed", err)
	}
	fmt.Println(resp)
}
