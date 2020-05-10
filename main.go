package main

import (
	"fmt"
	"log"

	"github.com/optresume/jobs"
)

func main() {
	BaseUrl := "https://www.optresume.com/"
	output, err := jobs.JobsByTitle(BaseUrl, "golang-developer", "2")

	if err != nil {
		fmt.Printf("Error to read response {%v}\n", err)
		return
	}

	if output == "" {
		log.Println("WARNING: output string is blank!!!")
		return
	}
	fmt.Println(output)

}
