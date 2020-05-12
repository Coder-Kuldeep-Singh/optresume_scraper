package results

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/optresume/client"
)

type Jobs struct {
	ResultRank  int
	Title       string
	Company     string
	Location    string
	Type        string
	View        string
	Description string
	PostedTime  string
}

// Main function which used to extract main content
func ExtracterData(URL string, pages int) ([]Jobs, error) {
	results := []Jobs{}
	for i := 1; i <= pages; i++ {
		pageNumber := strconv.Itoa(i)
		URL = URL + "?&page_number=" + pageNumber
		resp, err := client.Request(URL)
		if err != nil {
			log.Printf("Proxy Failed {%v}\n", err)
		}
		doc, err := goquery.NewDocumentFromResponse(resp)
		if err != nil {
			return nil, err
		}

		sel := doc.Find("li.careerfy-column-12")

		rank := 1
		for i := range sel.Nodes {
			item := sel.Eq(i)
			title := item.Find("div.careerfy-list-option > h3")
			company := item.Find("div.careerfy-list-option > ul > li:nth-child(1) > a")
			location := item.Find("div.careerfy-list-option > ul > li:nth-child(2)")
			types := item.Find("div.careerfy-list-option > ul > li:nth-child(3)")
			view := item.Find("div.careerfy-job-userlist > a")
			description := item.Find("div.card-description")
			postedtime := item.Find("ul.careerfy-column-12 > li")

			T := title.Text()
			T = strings.TrimSpace(T)

			C := company.Text()
			C = strings.TrimSpace(C)

			L := location.Text()
			L = strings.TrimSpace(L)

			TY := types.Text()
			TY = strings.TrimSpace(TY)

			V, _ := view.Attr("href")
			V = strings.TrimSpace(V)

			D := description.Text()
			D = strings.TrimSpace(D)

			P := postedtime.Text()
			P = strings.TrimSpace(P)

			fmt.Printf("Title {%s}\n", T)
			fmt.Printf("Company {%s}\n", C)
			fmt.Printf("Location {%s}\n", L)
			fmt.Printf("Type {%s}\n", TY)
			fmt.Printf("View {%s}\n", V)
			fmt.Printf("Description {%s}\n", D)
			fmt.Printf("PostedTime {%v}\n", P)

			if V != "" && V != "#" {
				result := Jobs{
					rank,
					T,
					C,
					L,
					TY,
					V,
					D,
					P,
				}
				results = append(results, result)
				rank += 1
			}
		}

	}
	return results, nil
}
