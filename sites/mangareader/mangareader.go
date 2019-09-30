package mangareader

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
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

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func init() {
	selector = "img#img"
	next = "span.next > a"
	url = "http://www.mangareader.net"
	imageCounter = 1
}

// #Execute blah blah
func Execute(manga string, chapter int, output string) {
	manga = strings.ToLower(manga)
	manga = strings.Replace(manga, " ", "-", -1)
	fmt.Println(output)
	fmt.Println(manga)

	if output == "" {
		folderPath = fmt.Sprintf("%s/%d", manga, chapter)
	} else {
		folderPath = fmt.Sprintf("%s%s/%d", output, manga, chapter)
	}

	fmt.Println(folderPath)
	os.MkdirAll(folderPath, os.ModePerm)
	link := fmt.Sprintf("%s/%s/%d", url, manga, chapter)
	crawl(link, chapter)
}

func fetchURL(link string) (string, string) {

	var bodyString string

	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	resp, err := client.Get(link)

	checkError(err)

	if resp.StatusCode == http.StatusOK {
		if doc, err := goquery.NewDocumentFromReader(resp.Body); err == nil {
			bodyString, _ = doc.Find(selector).Attr("src")
			nexthref, _ := doc.Find(next).Attr("href")
			nextLink := fmt.Sprintf("%s%s", url, nexthref)
			return bodyString, nextLink
		}
	}

	return "nil", "nil"
}

func crawl(url string, chapter int) {

	currentChapter, err := strconv.Atoi(strings.Split(url, "/")[4])

	checkError(err)

	if currentChapter == chapter {
		imageURL, nextURL := fetchURL(url)
		// fmt.Println(imageURL)
		download(imageURL)
		crawl(nextURL, chapter)
	}

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
