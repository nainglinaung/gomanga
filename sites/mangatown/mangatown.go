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
)

func init() {
	selector.Current = "#viewer > a > img"
	selector.Next = ""
	url = "http://mangatown.com/"
	imageCounter = 1
}

func Execute(manga string, chapter int, output string) {
	manga = ehelper.LowerAndReplace(manga, " ", "_")
	manga = strings.Replace(manga, "-", "_", -1)
	folderPath = ehelper.CreateFolderPath(manga, chapter, output)
	url := fmt.Sprintf("%smanga/%s/c%03d", url, manga, chapter)
	crawl(url)
}

func crawl(url string) {
	fileLink := fmt.Sprintf("%s/%d.html", url, imageCounter)
	resp := ehelper.FetchURL(fileLink)

	if resp != nil {
		imageLink, _ := ehelper.ParseResponse(resp.Body, selector)
		if len(imageLink) > 0 {
			fullImagePath := fmt.Sprintf("%s/%d.jpg", folderPath, imageCounter)
			fmt.Println(fullImagePath)
			ehelper.Download(imageLink, fullImagePath)
			imageCounter++
			crawl(url)
		}

	}
}
