package jobs

import (
	"log"

	"github.com/optresume/client"
	"github.com/optresume/reader"
)

type Jobs struct {
	Title       string
	Compnay     string
	Location    string
	Type        string
	View        string
	Description string
	PostedTime  string
}

func JobsByTitle(URL string, jobTitle, pages string) (string, error) {
	URL = URL + jobTitle + "-jobs?page_number=" + pages
	resp, err := client.Request(URL)
	if err != nil {
		log.Printf("Proxy Failed {%v}\n", err)
	}

	output, err := reader.ReadAll(".careerfy-row", resp)
	if err != nil {
		return output, err
	}
	return output, err
}
