package main

import (
	"fmt"
	"sync"

	"github.com/optresume/jobs"
	"github.com/optresume/model"
)

func main() {
	var wg sync.WaitGroup
	BaseUrl := "https://www.optresume.com/"
	wg.Add(1)
	go func() {

		output, err := jobs.AllJobs(BaseUrl, 100, &wg)
		if err != nil {
			fmt.Printf("Error to read response {%v}\n", err)
			return
		}
		// for _, list := range output {
		// fmt.Println(output)
		wg.Add(1)
		go func() {
			model.Optresume(output, &wg)
		}()
	}()
	// }
	wg.Wait()
	fmt.Println("************************************************************************************************************************************************")
	fmt.Println("************************************************************************************************************************************************")
	fmt.Println("Execution Finished")
}
