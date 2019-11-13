package mangareader

import (
	"fmt"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/nainglinaung/gomanga/lib/ehelper"
)

var (
	url          string
	folderPath   string
	imageCounter int
	selector     ehelper.Selector
	wg           sync.WaitGroup
	helper       ehelper.Ehelper
	total        []string
)

func init() {
	selector.Current = "img#img"
	selector.Next = "#pageMenu > option"
	url = "http://www.mangareader.net"
	imageCounter = 1

}

// #Execute blah blah
func Execute(manga string, chapter int, output string) {
	manga = helper.LowerAndReplace(manga, " ", "-")
	folderPath = helper.CreateFolderPath(manga, chapter, output)
	helper.CreateFolder(folderPath)
	println(folderPath)
	crawl(fmt.Sprintf("%s/%s/%d", url, manga, chapter))
}

func GetTotalCount(link string) []string {
	resp := helper.FetchURL(link)
	doc := helper.Parse(resp.Body)
	doc.Find(selector.Next).Each(func(i int, s *goquery.Selection) {
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

	imageList := GetTotalCount(link)
	imageArrayLength := len(imageList)
	wg.Add(imageArrayLength)

	for i := 0; i < imageArrayLength; i++ {
		go func(i int) {
			defer wg.Done()
			fullURL := fmt.Sprintf("%s%s", url, imageList[i])
			resp := helper.FetchURL(fullURL)
			if resp != nil {
				imageURL, _ := helper.ParseResponse(resp.Body, selector)
				fullImagePath := fmt.Sprintf("%s/%d.jpg", folderPath, i)
				helper.Download(imageURL, fullImagePath)
				helper.Log(fullImagePath)
			}
		}(i)
	}
	wg.Wait()

}
