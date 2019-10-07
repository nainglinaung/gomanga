package hentaicafe

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
	linkLength   int
	// ehelper ehelper
)

func init() {
	selector.Current = "img.open"
	selector.Next = "span.next > a"
	url = "http://www.hentai.cafe"
	imageCounter = 1
	linkLength = 11
}

func Execute(code int, output string) {

	if output == "" {
		folderPath = fmt.Sprintf("hentai-cafe/%d", code)
	} else {
		folderPath = fmt.Sprintf("%shentai-cafe/%d", output, code)
	}

	helper.CreateFolder(folderPath)
	crawl(getChapterLink(code))

}

func getChapterLink(code int) string {
	chapterLink := fmt.Sprintf("%s/hc.fyi/%d", url, code)
	resp := helper.RequestChapterLink(chapterLink)
	return fmt.Sprintf("%s%s", helper.ParseChapter(resp.Body, "a[title='Read']"), "page/1")

}

func crawl(link string) {
	currentLinkArray := strings.Split(link, "/")
	currentLinkLength := len(currentLinkArray)

	if linkLength == currentLinkLength {
		resp := helper.FetchURL(link)
		if resp != nil {
			imageURL, _ := helper.ParseResponse(resp.Body, selector)
			defer resp.Body.Close()
			fmt.Println(imageURL)
			if len(imageURL) > 0 {
				pageLength, _ := strconv.Atoi(currentLinkArray[linkLength-1])
				fullImagePath := fmt.Sprintf("%s/%d.jpg", folderPath, pageLength)
				fmt.Println(fullImagePath)
				pageLength++
				currentLinkArray[linkLength-1] = fmt.Sprintf("%d", pageLength)
				nextLink := strings.Join(currentLinkArray, "/")
				println(nextLink)
				helper.Download(imageURL, fullImagePath)
				crawl(nextLink)
			}
		}
	}

}
