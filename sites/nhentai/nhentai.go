package nhentai

import (
	"fmt"
	"io"
	"net/http"
	"os"
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
	imageCounter = 1
}

func Execute(code int, output string) {

	if output == "" {
		folderPath = fmt.Sprintf("%d", code)
	} else {
		folderPath = fmt.Sprintf("%s/%d", output, code)
	}

	os.MkdirAll(folderPath, os.ModePerm)
	link := fmt.Sprintf("%s/g/%d", url, code)
	crawl(link, imageCounter)

}

func crawl(url string, counter int) {

	// if currentChapter == chapter {
	imageURL := fmt.Sprintf("%s/%d", url, counter)
	imageURL = fetchURL(imageURL)
	if imageURL != "nil" {
		download(imageURL)
		counter++
		crawl(url, counter)
	}

}

func download(url string) {

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := client.Get(url)
	ehelper.CheckError(err)
	defer resp.Body.Close()

	fullImagePath := fmt.Sprintf("%s/%d.jpg", folderPath, imageCounter)
	fmt.Println(fullImagePath)
	imageCounter++
	// fmt.Println(fullImagePath)
	file, err := os.Create(fullImagePath)

	_, err = io.Copy(file, resp.Body)
	ehelper.CheckError(err)

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
