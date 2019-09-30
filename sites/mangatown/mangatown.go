package mangatown

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var (
	selector     string
	next         string
	url          string
	folderPath   string
	imageCounter int
)

func init() {
	selector = "div#viewer > a > img"
	url = "http://mangatown.com/"
	imageCounter = 1
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func fetchURL(link string) string {

	var bodyString string

	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	resp, err := client.Get(link)

	checkError(err)

	if resp.StatusCode == http.StatusOK {
		if doc, err := goquery.NewDocumentFromReader(resp.Body); err == nil {
			bodyString, _ = doc.Find(selector).Attr("src")
			return bodyString
		}
	}

	return "nil"
}

func Execute(manga string, chapter int) {
	manga = strings.ToLower(manga)
	manga = strings.Replace(manga, " ", "_", -1)
	manga = strings.Replace(manga, "-", "_", -1)

	fmt.Println(manga)
	folderPath = fmt.Sprintf("%s/%d", manga, chapter)
	os.MkdirAll(folderPath, os.ModePerm)
	link := fmt.Sprintf("%s/manga/%s/c%03d", url, manga, chapter)
	fmt.Println(link)
	crawl(link)

}

func download(url string) {

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := client.Get(url)
	defer resp.Body.Close()
	checkError(err)
	// folderPath

	fullImagePath := fmt.Sprintf("%s/%d.jpg", folderPath, imageCounter)
	imageCounter++
	fmt.Println(fullImagePath)
	file, err := os.Create(fullImagePath)

	_, err = io.Copy(file, resp.Body)
	checkError(err)

}

func crawl(url string) {

	link := fmt.Sprintf("%s/%d.html", url, imageCounter)
	imageURL := fetchURL(link)

	if imageURL != "nil" {
		download(imageURL)
		crawl(url)
	}

	// crawl(url)

}
