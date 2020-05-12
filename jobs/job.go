package jobs

import (
	re "github.com/optresume/results"
)

// Version 1 function
// func JobsByTitle(URL string, jobTitle string, pages int, wg *sync.WaitGroup) ([]Jobs, error) {
// 	defer wg.Done()
// 	results := []Jobs{}
// 	// var results []string

// 	for i := 1; i <= pages; i++ {
// 		pageNumber := strconv.Itoa(i)
// 		URL = URL + jobTitle + "-jobs?&page_number=" + pageNumber
// 		resp, err := client.Request(URL)
// 		if err != nil {
// 			log.Printf("Proxy Failed {%v}\n", err)
// 		}

// 		// output, err := reader.ReadAll(".careerfy-row", resp)
// 		// if err != nil {
// 		// 	fmt.Printf("Error to get response {%v}\n", err)
// 		// }
// 		// fmt.Println("*******************************************************************************************************************")
// 		// results = append(results, output)

// 		doc, err := goquery.NewDocumentFromResponse(resp)
// 		if err != nil {
// 			return nil, err
// 		}
// 		sel := doc.Find("ul.careerfy-row")
// 		for i := range sel.Nodes {
// 			item := sel.Eq(i)
// 			title := item.Find("div.careerfy-list-option > h3")
// 			company := item.Find("div.careerfy-list-option > ul > li:nth-child(1) > a")
// 			location := item.Find("div.careerfy-list-option > ul > li:nth-child(2)")
// 			types := item.Find("div.careerfy-list-option > ul > li:nth-child(3)")
// 			view := item.Find("div.careerfy-job-userlist > a")
// 			description := item.Find("div.card-description")
// 			postedtime := item.Find("div.careerfy-job.careerfy-joblisting-classic > ul > li:nth-child(1) > div > div > ul > li")

// 			T := title.Text()
// 			T = strings.TrimSpace(T)

// 			C := company.Text()
// 			C = strings.TrimSpace(C)

// 			L := location.Text()
// 			L = strings.TrimSpace(L)

// 			TY := types.Text()
// 			TY = strings.TrimSpace(TY)

// 			V, _ := view.Attr("href")
// 			V = strings.TrimSpace(V)

// 			D := description.Text()
// 			D = strings.TrimSpace(D)

// 			P := postedtime.Text()
// 			P = strings.TrimSpace(P)
// 			result := Jobs{
// 				T,
// 				C,
// 				L,
// 				TY,
// 				V,
// 				D,
// 				P,
// 			}
// 			results = append(results, result)
// 		}

// 	}
// 	return results, nil
// }

// Running Function
func JobsByTitle(URL string, jobTitle string, pages int) ([]re.Jobs, error) {
	URL = URL + jobTitle + "-jobs"
	results, err := re.ExtracterData(URL, pages)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func JobsBySkill(URL string, Skill string, pages int) ([]re.Jobs, error) {
	URL = URL + Skill + "-jobs"
	results, err := re.ExtracterData(URL, pages)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func JobsByCategories(URL string, Categories string, pages int) ([]re.Jobs, error) {
	URL = URL + Categories + "-jobs"
	results, err := re.ExtracterData(URL, pages)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func AllJobs(URL string, pages int) ([]re.Jobs, error) {
	URL = URL + "job-search"
	results, err := re.ExtracterData(URL, pages)
	if err != nil {
		return nil, err
	}
	return results, nil
}
