package mngdoom

import (
	"fmt"
	"io"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/nainglinaung/gomanga/lib/ehelper"
)

var (
	baseURL    string
	folderPath string
	selector   ehelper.Selector
	helper     ehelper.Ehelper
	wg         sync.WaitGroup
	total      []string
)

func init() {
	selector.Current = ".chapter-page2 > option"
	selector.Next = "span.next > a"
	baseURL = "http://www.mngdoom.com"
}

// #Execute blah blah
func Execute(manga string, chapter int, output string) {
	manga = helper.LowerAndReplace(manga, " ", "-")
	folderPath = helper.CreateFolderPath(manga, chapter, output)
	helper.CreateFolder(folderPath)
	crawl(fmt.Sprintf("%s/%s/%d", baseURL, manga, chapter))
}

func parseChapterArray(body io.Reader, Selector ehelper.Selector) []string {
	doc := helper.Parse(body)
	doc.Find(selector.Current).Each(func(i int, s *goquery.Selection) {
		data, exist := s.Attr("value")
		if exist {
			total = append(total, data)
		} else {
			fmt.Println("not existed")
		}
	})
	return total
}

func crawl(link string) {

	resp := helper.FetchURL(link)
	if resp != nil {

		imageArray := parseChapterArray(resp.Body, selector)
		imageArrayLength := len(imageArray)
		wg.Add(imageArrayLength)

		for i := 0; i < imageArrayLength; i++ {
			go func(i int) {
				defer wg.Done()
				imageURL := strings.TrimSpace(imageArray[i])
				fullImagePath := fmt.Sprintf("%s/%d.jpg", folderPath, i)
				helper.Download(imageURL, fullImagePath)
				helper.Log(fullImagePath)
			}(i)
		}

	}
	wg.Wait()

}
