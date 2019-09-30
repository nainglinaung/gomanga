package ehelper

import (
	"io"
	"net/http"
	"os"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Selector struct {
	Current string
	Next    string
}

var (
	client *http.Client
)

func init() {
	client = &http.Client{
		Timeout: 30 * time.Second,
	}
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func request(url string) *http.Response {
	resp, err := client.Get(url)
	CheckError(err)
	return resp
}

func FetchURL(link string) *http.Response {
	resp := request(link)
	if resp.StatusCode == http.StatusOK {
		return resp
	} else {
		return nil
	}
}

// func ParseResponse(body io.Reader, selector string, next string) {

// }

func ParseResponse(body io.Reader, selector Selector) (string, string) {
	doc, err := goquery.NewDocumentFromReader(body)
	CheckError(err)
	bodyString, existFlag := doc.Find(selector.Current).Attr("src")
	nextString, existFlag2 := doc.Find(selector.Next).Attr("href")

	if existFlag && existFlag2 {
		return bodyString, nextString
	} else if existFlag && !existFlag2 {
		return bodyString, ""
	}

	return "", ""
}

func Download(url string, fullImagePath string) {

	resp := request(url)
	// fmt.Println(fullImagePath)
	file, err := os.Create(fullImagePath)
	_, err = io.Copy(file, resp.Body)
	CheckError(err)

}
