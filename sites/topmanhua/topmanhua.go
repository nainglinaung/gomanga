package topmanhua

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

// https://www.isekaiscan.com/manga/tomb-raider-king/chapter-1
// https://isekaiscan.com/manga/tomb-raider-king/chapter-1/
// https://www.isekaiscan.com/manga/tomb-raider-king/chapter-28/

func init() {
	selector.Current = ".page-break > img"
	baseURL = "http://topmanhua.com"
}

// #Execute blah blah
func Execute(manga string, chapter int, output string) {
	manga = helper.LowerAndReplace(manga, " ", "-")
	folderPath = helper.CreateFolderPath(manga, chapter, output)
	helper.CreateFolder(folderPath)
	crawl(fmt.Sprintf("%s/manhua/%s/chapter-%d", baseURL, manga, chapter), chapter)
}

func crawl(link string, chapter int) {

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
