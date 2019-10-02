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

type Selector struct {
	Current string
	Next    string
}

// func init() {

// }

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func CreateFolderPath(manga string, chapter int, output string) string {
	var folderPath string
	if output == "" {
		folderPath = fmt.Sprintf("%s%d", manga, chapter)

	} else {
		folderPath = fmt.Sprintf("%s%s/%d", output, manga, chapter)
	}

	return folderPath
}

func CreateFolder(path string) {
	os.MkdirAll(path, os.ModePerm)
}

func LowerAndReplace(manga, existing, replacement string) string {
	manga = strings.ToLower(manga)
	manga = strings.Replace(manga, existing, replacement, -1)
	return manga
}

func request(url string) *http.Response {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	// fmt.Println(url)
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

func ParseResponse(body io.Reader, selector Selector) (string, string) {
	doc, err := goquery.NewDocumentFromReader(body)
	CheckError(err)

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

func Download(url string, fullImagePath string) {
	// println(url)
	// println(fullImagePath)
	resp := request(url)
	file, err := os.Create(fullImagePath)
	_, err = io.Copy(file, resp.Body)
	CheckError(err)

}
