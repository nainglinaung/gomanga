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
	LogHandler
	FileHandler
}

type Selector struct {
	Current string
	Next    string
}

func (e Ehelper) request(url string) *http.Response {
	client := &http.Client{
		Timeout: 600 * time.Second,
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

func (e Ehelper) Parse(body io.Reader) *goquery.Document {
	doc, err := goquery.NewDocumentFromReader(body)
	e.CheckError(err)
	return doc
}

func (e Ehelper) ParseResponse(body io.Reader, selector Selector) (string, string) {
	doc, err := goquery.NewDocumentFromReader(body)
	e.CheckError(err)
	bodyString, existFlag := doc.Find(selector.Current).Attr("src")
	nextString, existFlag2 := doc.Find(selector.Next).Attr("href")
	if existFlag && existFlag2 {
		return bodyString, nextString
	} else if existFlag && !existFlag2 {
		return bodyString, ""
	}

	return "", ""
}
