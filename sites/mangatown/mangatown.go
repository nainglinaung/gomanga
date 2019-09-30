package mangatown

import (
	"fmt"
	"os"
	"strings"

	"github.com/nainglinaung/gomanga/lib/ehelper"
)

type Selector struct {
	current string
	next    string
}

var (
	next         string
	url          string
	folderPath   string
	imageCounter int
	selector     Selector
)

func init() {
	selector.current = "#viewer > a > img"
	selector.next = ""
	url = "http://mangatown.com/"
	imageCounter = 1
}

func Execute(manga string, chapter int, output string) {
	manga = strings.ToLower(manga)
	manga = strings.Replace(manga, " ", "_", -1)
	manga = strings.Replace(manga, "-", "_", -1)

	if output == "" {
		folderPath = fmt.Sprintf("%s/%d", manga, chapter)
	} else {
		folderPath = fmt.Sprintf("%s%s/%d", output, manga, chapter)
	}

	folderPath = fmt.Sprintf("%s/%d", manga, chapter)
	os.MkdirAll(folderPath, os.ModePerm)
	link := fmt.Sprintf("%smanga/%s/c%03d", url, manga, chapter)
	fmt.Println(link)
	crawl(link)
}

func crawl(url string) {

	link := fmt.Sprintf("%s/%d.html", url, imageCounter)
	resp := ehelper.FetchURL(link)

	if resp != nil {
		fmt.Println(selector)
		// link := ehelper.ParseResponse(resp.Body, selector)
		// if len(link) > 0 {
		// 	fullImagePath := fmt.Sprintf("%s/%d.jpg", folderPath, imageCounter)
		// 	ehelper.Download(link, fullImagePath)
		// 	imageCounter++
		// 	link = fmt.Sprintf("%s/%d.html", url, imageCounter)
		// 	crawl(url)
		// }

	}
}
