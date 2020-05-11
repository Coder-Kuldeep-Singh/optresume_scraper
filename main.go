package main

import (
	"fmt"

	"github.com/optresume/jobs"
	"github.com/optresume/model"
)

func main() {
	BaseUrl := "https://www.optresume.com/"
	output, err := jobs.AllJobs(BaseUrl, 100)
	if err != nil {
		fmt.Printf("Error to read response {%v}\n", err)
		return
	}
	// for _, list := range output {
	// fmt.Println(output)
	model.Optresume(output)
	// }
	fmt.Println("************************************************************************************************************************************************")
	fmt.Println("************************************************************************************************************************************************")
	fmt.Println("Execution Finished")

}
