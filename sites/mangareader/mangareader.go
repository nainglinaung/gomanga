package mangareader

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nainglinaung/gomanga/lib/ehelper"
)

var (
	url          string
	folderPath   string
	imageCounter int
	selector     ehelper.Selector
	helper       ehelper.Ehelper
)

func init() {
	selector.Current = "img#img"
	selector.Next = "span.next > a"
	url = "http://www.mangareader.net"
	imageCounter = 1

}

// #Execute blah blah
func Execute(manga string, chapter int, output string) {
	manga = helper.LowerAndReplace(manga, " ", "-")
	folderPath = helper.CreateFolderPath(manga, chapter, output)
	helper.CreateFolder(folderPath)
	println(folderPath)
	crawl(fmt.Sprintf("%s/%s/%d", url, manga, chapter), chapter)
}

func crawl(link string, chapter int) {

	currentChapter, err := strconv.Atoi(strings.Split(link, "/")[4])
	helper.CheckError(err)
	if currentChapter == chapter {
		resp := helper.FetchURL(link)
		if resp != nil {
			imageURL, nextLink := helper.ParseResponse(resp.Body, selector)
			defer resp.Body.Close()
			if len(link) > 0 {
				fullImagePath := fmt.Sprintf("%s/%d.jpg", folderPath, imageCounter)
				nextLink = fmt.Sprintf("%s%s", url, nextLink)
				helper.Download(imageURL, fullImagePath)
				helper.Log(fullImagePath)
				imageCounter++
				crawl(nextLink, chapter)
			}
		}
	}

}
