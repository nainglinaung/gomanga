package mngdoom

import (
	"fmt"
	"strings"
	"sync"

	"github.com/nainglinaung/gomanga/lib/ehelper"
)

var (
	baseURL    string
	folderPath string
	selector   ehelper.Selector
	helper     ehelper.Ehelper
	wg         sync.WaitGroup
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

func crawl(link string) {

	resp := helper.FetchURL(link)
	if resp != nil {

		imageArray := helper.ParseChapterArray(resp.Body, selector)
		imageArrayLength := len(imageArray)
		wg.Add(imageArrayLength)

		for i := 0; i < imageArrayLength; i++ {
			go func(i int) {
				imageURL := strings.TrimSpace(imageArray[i])
				fullImagePath := fmt.Sprintf("%s/%d.jpg", folderPath, i)
				helper.Download(imageURL, fullImagePath)
				helper.Log(fullImagePath)
				defer wg.Done()
			}(i)
		}

	}
	wg.Wait()

}
