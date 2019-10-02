package nhentai

import (
	"fmt"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/nainglinaung/gomanga/lib/ehelper"
)

var (
	selector     string
	next         string
	url          string
	folderPath   string
	imageCounter int
)

func init() {
	selector = "#image-container > a > img"
	url = "https://nhentai.net"
}

func Execute(code int, output string) {

	// folderPath = ehelper.CreateFolderPath(manga, chapter, output)

	if output == "" {
		folderPath = fmt.Sprintf("%d", code)
	} else {
		folderPath = fmt.Sprintf("%s/%d", output, code)
	}

	ehelper.CreateFolder(folderPath)
	crawl(fmt.Sprintf("%s/g/%d", url, code), 1)

}

func crawl(url string, counter int) {

	// if currentChapter == chapter {
	fileLink := fmt.Sprintf("%s/%d", url, counter)
	imageURL := fetchURL(fileLink)
	if imageURL != "nil" {
		fullImagePath := fmt.Sprintf("%s/%d.jpg", folderPath, counter)
		ehelper.Download(imageURL, fullImagePath)
		counter++
		crawl(url, counter)
	}

}

func fetchURL(link string) string {

	var bodyString string

	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	resp, err := client.Get(link)

	ehelper.CheckError(err)

	if resp.StatusCode == http.StatusOK {
		if doc, err := goquery.NewDocumentFromReader(resp.Body); err == nil {
			bodyString, _ = doc.Find(selector).Attr("src")
			return bodyString
		}
	}

	return "nil"
}
