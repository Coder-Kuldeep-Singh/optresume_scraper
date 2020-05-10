package main

import (
	"fmt"
	"sync"

	"github.com/optresume/jobs"
)

func main() {
	var wg sync.WaitGroup
	BaseUrl := "https://www.optresume.com/"
	wg.Add(1)
	go func() {
		output, err := jobs.JobsByTitle(BaseUrl, "dot2wnet-developer", 1, &wg)
		if err != nil {
			fmt.Printf("Error to read response {%v}\n", err)
			return
		}
		for _, list := range output {
			fmt.Println(list.Description)
			fmt.Println("************************************************************************************************************************************************")
			fmt.Println("************************************************************************************************************************************************")
		}

	}()
	wg.Wait()
	fmt.Println("Execution Finished")

}
