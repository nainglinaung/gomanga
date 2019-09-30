package mangareader

import (
	"fmt"
	"os"
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
	manga = strings.ToLower(manga)
	manga = strings.Replace(manga, " ", "-", -1)

	if output == "" {
		folderPath = fmt.Sprintf("%s/%d", manga, chapter)
	} else {
		folderPath = fmt.Sprintf("%s/%s/%d", output, manga, chapter)
	}

	os.MkdirAll(folderPath, os.ModePerm)
	link := fmt.Sprintf("%s/%s/%d", url, manga, chapter)
	crawl(link, chapter)
}

func crawl(link string, chapter int) {

	currentChapter, err := strconv.Atoi(strings.Split(link, "/")[4])
	ehelper.CheckError(err)

	if currentChapter == chapter {
		resp := ehelper.FetchURL(link)
		if resp != nil {
			imageURL, nextLink := ehelper.ParseResponse(resp.Body, selector)
			if len(link) > 0 {
				fullImagePath := fmt.Sprintf("%s/%d.jpg", folderPath, imageCounter)
				nextLink = fmt.Sprintf("%s%s", url, nextLink)
				ehelper.Download(imageURL, fullImagePath)
				imageCounter++
				fmt.Println(nextLink)
				crawl(nextLink, chapter)
			}
		}
	}

}
