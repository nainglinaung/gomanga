package manhwaclub

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
	total      []string
	selector   ehelper.Selector
	helper     ehelper.Ehelper
	wg         sync.WaitGroup
)

func init() {
	selector.Current = ".wp-manga-chapter-img"
	selector.Next = "span.next > a"
	baseURL = "https://manhwa.club"
}

func Execute(manga string, chapter int, output string) {
	manga = helper.LowerAndReplace(manga, " ", "-")
	folderPath = helper.CreateFolderPath(manga, chapter, output)
	helper.CreateFolder(folderPath)
	crawl(fmt.Sprintf("%s/manhwa/%s/chapter-%d", baseURL, manga, chapter), chapter)
}

func parseChapterArray(body io.Reader, Selector ehelper.Selector) []string {
	doc := helper.Parse(body)

	doc.Find(selector.Current).Each(func(i int, s *goquery.Selection) {
		data, exist := s.Attr("data-src")

		if exist {
			total = append(total, data)
		} else {
			fmt.Println("not existed")
		}
	})
	return total
}
func crawl(link string, chapter int) {

	fmt.Println(link)

	resp := helper.FetchURL(link)
	if resp != nil {

		imageArray := parseChapterArray(resp.Body, selector)
		imageArrayLength := len(imageArray)
		fmt.Println(imageArrayLength)
		wg.Add(imageArrayLength)

		for i := 0; i < imageArrayLength; i++ {
			go func(i int) {
				defer wg.Done()
				imageURL := strings.TrimSpace(imageArray[i])
				fmt.Println(imageURL)
				fullImagePath := fmt.Sprintf("%s/%d.jpg", folderPath, i)
				helper.Download(imageURL, fullImagePath)
				helper.Log(fullImagePath)
			}(i)
		}

	}
	wg.Wait()

}
