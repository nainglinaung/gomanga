package mangatown

import (
	"fmt"
	"strings"

	"github.com/nainglinaung/gomanga/lib/ehelper"
)

var (
	next         string
	url          string
	folderPath   string
	imageCounter int
	selector     ehelper.Selector
	helper       ehelper.Ehelper
)

func init() {
	selector.Current = "#viewer > a > img"
	selector.Next = ""
	url = "http://mangatown.com/"
	imageCounter = 1
}

func Execute(manga string, chapter int, output string) {
	manga = helper.LowerAndReplace(manga, " ", "_")
	manga = strings.Replace(manga, "-", "_", -1)
	folderPath = helper.CreateFolderPath(manga, chapter, output)
	helper.CreateFolder(folderPath)
	url := fmt.Sprintf("%smanga/%s/c%03d", url, manga, chapter)
	crawl(url)
}

func crawl(url string) {
	fileLink := fmt.Sprintf("%s/%d.html", url, imageCounter)
	println(fileLink)
	resp := helper.FetchURL(fileLink)

	if resp != nil {
		imageLink, _ := helper.ParseResponse(resp.Body, selector)
		if len(imageLink) > 0 {
			fullImagePath := fmt.Sprintf("%s/%d.jpg", folderPath, imageCounter)
			fmt.Println(fullImagePath)
			helper.Download(imageLink, fullImagePath)
			imageCounter++
			crawl(url)
		}

	}
}
