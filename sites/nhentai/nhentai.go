package nhentai

import (
	"fmt"

	"github.com/nainglinaung/gomanga/lib/ehelper"
)

var (
	selector     ehelper.Selector
	next         string
	url          string
	folderPath   string
	imageCounter int
	helper       ehelper.Ehelper
)

func init() {
	selector.Current = "#image-container > a > img"
	url = "https://nhentai.net"
}

func Execute(code int, output string) {

	if output == "" {
		folderPath = fmt.Sprintf("%d", code)
	} else {
		folderPath = fmt.Sprintf("%s%d", output, code)
	}

	helper.CreateFolder(folderPath)
	crawl(fmt.Sprintf("%s/g/%d", url, code), 1)
}

func crawl(link string, counter int) {

	fileLink := fmt.Sprintf("%s/%d", link, counter)
	resp := helper.FetchURL(fileLink)

	if resp != nil {
		imageLink, _ := helper.ParseResponse(resp.Body, selector)
		fullImagePath := fmt.Sprintf("%s/%d.jpg", folderPath, counter)
		helper.Download(imageLink, fullImagePath)
		helper.Log(fullImagePath)
		counter++
		crawl(link, counter)
	}

}
