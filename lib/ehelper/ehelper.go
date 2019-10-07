package ehelper

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type ErrorHandler interface {
	CheckError()
}

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

type User struct {
	name string
}

// func (d *Downloader) Download(link) []byte, error {

func (e Ehelper) CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func (e Ehelper) CreateFolderPath(manga string, chapter int, output string) string {
	var folderPath string
	if output == "" {
		folderPath = fmt.Sprintf("%s/%d", manga, chapter)

	} else {
		folderPath = fmt.Sprintf("%s%s/%d", output, manga, chapter)
	}
	return folderPath
}

func (e Ehelper) CreateFolder(path string) {
	os.MkdirAll(path, os.ModePerm)
}

func (e Ehelper) LowerAndReplace(manga, existing, replacement string) string {
	manga = strings.ToLower(manga)
	manga = strings.Replace(manga, existing, replacement, -1)
	return manga
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
	// .Attr("src")
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

func (e Ehelper) Download(url string, fullImagePath string) {
	resp := e.request(url)
	file, err := os.Create(fullImagePath)
	_, err = io.Copy(file, resp.Body)
	e.CheckError(err)
}
