package hentainexus

import (
	"fmt"

	"github.com/nainglinaung/gomanga/lib/ehelper"
)

var (
	next       string
	url        string
	folderPath string
	helper     ehelper.Ehelper
	selector   ehelper.Selector
)

func init() {
	selector.Current = "img#currImage"
	url = "https://hentainexus.com"
}

func Execute(code int, output string) {

	if output == "" {
		folderPath = fmt.Sprintf("%d", code)
	} else {
		folderPath = fmt.Sprintf("%s%d", output, code)
	}

	helper.CreateFolder(folderPath)

	crawl(fmt.Sprintf("%s/read/%d", url, code), 1)
}

func crawl(link string, counter int) {

	fileLink := fmt.Sprintf("%s/%03d", link, counter)
	resp := helper.FetchURL(fileLink)
	if resp != nil {
		imageLink, _ := helper.ParseResponse(resp.Body, selector)
		if len(imageLink) > 0 {
			fullImagePath := fmt.Sprintf("%s/%d.jpg", folderPath, counter)
			helper.Download(imageLink, fullImagePath)
			helper.Log(fullImagePath)
			counter++
			crawl(link, counter)

		}
	}
}
