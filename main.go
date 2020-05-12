package main

import (
	"fmt"

	"github.com/optresume/jobs"
	"github.com/optresume/model"
)

func main() {
	BaseUrl := "https://www.optresume.com/"
	output, err := jobs.AllJobs(BaseUrl, 1)
	if err != nil {
		fmt.Printf("Error to read response {%v}\n", err)
		return
	}
	model.Optresume(output)
	fmt.Println("************************************************************************************************************************************************")
	fmt.Println("************************************************************************************************************************************************")
	fmt.Println("Execution Finished")
}
