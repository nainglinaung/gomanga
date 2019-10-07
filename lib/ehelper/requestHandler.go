package ehelper

import (
	"io"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type RequestHandler interface {
	request() *http.Response
	FetchURL() *http.Response
	ParseResponse() (string, string)
}

type FileHandler interface {
	CreateFolderPath() string
	CreateFolder()
	LowerAndReplace()
	Download()
}

type Ehelper struct {
	ErrorHandler
	RequestHandler
	FileHandler
}

type Selector struct {
	Current string
	Next    string
}

// type User struct {
// 	name string
// }

// func (d *Downloader) Download(link) []byte, error {

func (e Ehelper) RequestChapterLink(url string) *http.Response {
	return e.request(url)
}

func (e Ehelper) ParseChapter(body io.Reader, chapterSelector string) string {
	doc, err := goquery.NewDocumentFromReader(body)
	e.CheckError(err)
	token, exist := doc.Find(chapterSelector).Attr("href")
	if exist {
		return token
	}
	return ""

}

func (e Ehelper) request(url string) *http.Response {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	resp, err := client.Get(url)
	e.CheckError(err)
	return resp
}

func (e Ehelper) FetchURL(link string) *http.Response {
	resp := e.request(link)

	if resp.StatusCode == http.StatusOK {
		return resp
	} else {
		return nil
	}
}

func (e Ehelper) ParseResponse(body io.Reader, selector Selector) (string, string) {
	doc, err := goquery.NewDocumentFromReader(body)
	e.CheckError(err)

	// selector.Current
	bodyString, existFlag := doc.Find(selector.Current).Attr("src")
	nextString, existFlag2 := doc.Find(selector.Next).Attr("href")
	// fmt.Println(bodyString)
	// CheckError(existFlag)
	if existFlag && existFlag2 {
		return bodyString, nextString
	} else if existFlag && !existFlag2 {
		return bodyString, ""
	}

	return "", ""
}
