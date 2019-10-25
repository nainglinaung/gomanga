package mangahere

import (
	"fmt"
	"sync"

	"github.com/nainglinaung/gomanga/lib/ehelper"
)

var (
	url          string
	folderPath   string
	imageCounter int
	selector     ehelper.Selector
	wg           sync.WaitGroup
	helper       ehelper.Ehelper
)

func init() {
	selector.Current = "img.img"
	selector.Next = "#sel_page_1 > option"
	url = "http://www.manga-here.club"
	imageCounter = 1

}

// #Execute blah blah
func Execute(manga string, chapter int, output string) {
	manga = helper.LowerAndReplace(manga, " ", "-")
	folderPath = helper.CreateFolderPath(manga, chapter, output)
	helper.CreateFolder(folderPath)
	println(folderPath)
	crawl(fmt.Sprintf("%s/%s/chapter-%d", url, manga, chapter))
}

func GetTotalCount(link string) []string {
	resp := helper.FetchURL(link)
	return helper.ParseTotalCount(resp.Body, selector)
}

func crawl(link string) {
	println(link)
	imageList := GetTotalCount(link)
	imageArrayLength := len(imageList)
	wg.Add(imageArrayLength)

	fmt.Println(imageList)

	for i := 0; i < imageArrayLength; i++ {
		go func(i int) {
			fullURL := fmt.Sprintf("%s%s", url, imageList[i])
			resp := helper.FetchURL(fullURL)
			if resp != nil {
				imageURL, _ := helper.ParseResponse(resp.Body, selector)
				fullImagePath := fmt.Sprintf("%s/%d.jpg", folderPath, i)
				helper.Download(imageURL, fullImagePath)
				helper.Log(fullImagePath)
				defer wg.Done()
			}
		}(i)
	}
	wg.Wait()

}
