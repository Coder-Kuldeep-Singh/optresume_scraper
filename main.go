package main

import (
	"fmt"

	"github.com/optresume/jobs"
)

func main() {
	BaseUrl := "https://www.optresume.com/"
	output, err := jobs.AllJobs(BaseUrl, 2)
	if err != nil {
		fmt.Printf("Error to read response {%v}\n", err)
		return
	}
	// for _, list := range output {
	fmt.Println(output)
	// }
	fmt.Println("************************************************************************************************************************************************")
	fmt.Println("************************************************************************************************************************************************")
	fmt.Println("Execution Finished")

}
