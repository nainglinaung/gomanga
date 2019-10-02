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
)

func init() {
	selector.Current = "img#img"
	selector.Next = "span.next > a"
	url = "http://www.mangareader.net"
	imageCounter = 1
}

// #Execute blah blah
func Execute(manga string, chapter int, output string) {
	manga = ehelper.LowerAndReplace(manga, " ", "-")
	folderPath = ehelper.CreateFolderPath(manga, chapter, output)
	ehelper.CreateFolder(folderPath)
	println(folderPath)
	crawl(fmt.Sprintf("%s/%s/%d", url, manga, chapter), chapter)
}

func crawl(link string, chapter int) {
	println(link)
	currentChapter, err := strconv.Atoi(strings.Split(link, "/")[4])
	ehelper.CheckError(err)

	if currentChapter == chapter {
		resp := ehelper.FetchURL(link)
		if resp != nil {
			imageURL, nextLink := ehelper.ParseResponse(resp.Body, selector)
			defer resp.Body.Close()
			if len(link) > 0 {
				fullImagePath := fmt.Sprintf("%s/%d.jpg", folderPath, imageCounter)
				nextLink = fmt.Sprintf("%s%s", url, nextLink)
				ehelper.Download(imageURL, fullImagePath)
				imageCounter++
				crawl(nextLink, chapter)
			}
		}
	}

}
