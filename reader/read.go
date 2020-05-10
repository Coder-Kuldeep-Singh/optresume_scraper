package reader

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func ReadAll(selector string, resp *http.Response) (string, error) {
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return "", err
	}

	output := ""
	doc.Find(selector).Each(func(index int, item *goquery.Selection) {
		contents, _ := item.Html()
		output = contents
	})
	return output, nil
}
