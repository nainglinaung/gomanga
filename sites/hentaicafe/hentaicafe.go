package hentaicafe

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/nainglinaung/gomanga/lib/ehelper"
)

var (
	url          string
	folderPath   string
	imageCounter int
	selector     ehelper.Selector
	helper       ehelper.Ehelper
	linkLength   int
	wg           sync.WaitGroup
	// ehelper ehelper
)

func init() {
	selector.Current = "img.open"
	selector.Next = "span.next > a"
	url = "http://www.hentai.cafe"
	linkLength = 11
}

func Execute(code int, output string) {

	if output == "" {
		folderPath = fmt.Sprintf("hentai-cafe/%d", code)
	} else {
		folderPath = fmt.Sprintf("%shentai-cafe/%d", output, code)
	}

	helper.CreateFolder(folderPath)
	link := getChapterLink(code)
	// fmt.Println(link)
	pages := getPages(link)
	crawl(getChapterLink(code), pages)
	println("Finished")

}

func getPages(link string) int {
	resp := helper.FetchURL(link)
	doc := helper.Parse(resp.Body)
	token, e := doc.Find(".text").Html()
	helper.CheckError(e)

	s := strings.Split(token, " ")
	pages, _ := strconv.Atoi(s[0])
	return pages

}

func getChapterLink(code int) string {
	chapterLink := fmt.Sprintf("%s/hc.fyi/%d", url, code)
	resp := helper.FetchURL(chapterLink)
	doc := helper.Parse(resp.Body)
	token, exist := doc.Find("a[title='Read']").Attr("href")
	if exist {
		return token
	} else {
		return ""
	}

}

func crawl(link string, limit int) {
	fmt.Println(link)
	currentLinkArray := strings.Split(link, "/")
	imageArrayLength := len(currentLinkArray)
	baseURL := strings.Join(currentLinkArray[:imageArrayLength-1], "/")
	wg.Add(limit - 1)

	for i := 1; i <= limit; i++ {

		go func(i int) {

			resp := helper.FetchURL(fmt.Sprintf("%s/page/%d", baseURL, i))
			if resp != nil {
				imageURL, _ := helper.ParseResponse(resp.Body, selector)
				fullImagePath := fmt.Sprintf("%s/%d.jpg", folderPath, i)
				helper.Download(imageURL, fullImagePath)
				helper.Log(fullImagePath)
			}
			defer wg.Done()

		}(i)
	}
	wg.Wait()

}
