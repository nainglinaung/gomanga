package mangatown

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/nainglinaung/gomanga/lib/ehelper"
)

var (
	next         string
	url          string
	folderPath   string
	imageCounter int
	selector     ehelper.Selector
	helper       ehelper.Ehelper
	crawlWg      sync.WaitGroup
	downloadWg   sync.WaitGroup
)

func init() {
	selector.Current = "#viewer > a > img"
	selector.Next = ""
	url = "http://mangatown.com/"
	imageCounter = 1
}

func Execute(manga string, chapter int, output string) {
	manga = helper.LowerAndReplace(manga, " ", "_")
	manga = strings.Replace(manga, "-", "_", -1)
	folderPath = helper.CreateFolderPath(manga, chapter, output)
	helper.CreateFolder(folderPath)
	url := fmt.Sprintf("%smanga/%s/c%03d", url, manga, chapter)
	crawlWg.Add(1)
	start := time.Now()
	downloadChan := make(chan [2]string)
	go func() {
		for data := range downloadChan {
			downloadWg.Add(1)
			go download(data[0], data[1])
		}
	}()
	go crawl(url, downloadChan)
	crawlWg.Wait()
	fmt.Printf("Elasped Time: %#v\n", time.Now().Sub(start).Seconds())
	downloadWg.Wait()
}

func download(imageLink, fullImagePath string) {
	defer downloadWg.Done()
	helper.Download(imageLink, fullImagePath)
}
func crawl(url string, downloadChan chan<- [2]string) {
	defer crawlWg.Done()
	fileLink := fmt.Sprintf("%s/%d.html", url, imageCounter)
	println(fileLink)
	resp := helper.FetchURL(fileLink)
	if resp != nil {
		imageLink, _ := helper.ParseResponse(resp.Body, selector)
		if len(imageLink) > 0 {
			fullImagePath := fmt.Sprintf("%s/%d.jpg", folderPath, imageCounter)
			downloadChan <- [2]string{imageLink, fullImagePath}
			helper.Log(fullImagePath)
			imageCounter++
			crawlWg.Add(1)
			go crawl(url, downloadChan)
		} else {
			close(downloadChan)
		}

	}

}
