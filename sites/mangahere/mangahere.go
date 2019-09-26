package mangahere

import (
	"fmt"
	"net/http"
	"time"
)

func FetchURL(link string) (string, string) {

	var bodyString string

	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	resp, err := client.Get(link)

	if err != nil {
		panic(err)
	}

	// if resp.StatusCode == http.StatusOK {
	// 	if doc, err := goquery.NewDocumentFromReader(resp.Body); err == nil {
	// 		bodyString, _ = doc.Find(currentParser.Image.Target).Attr("src")
	// 		nexthref, _ := doc.Find(currentParser.Next.Target).Attr("href")
	// 		nextLink := fmt.Sprintf("%s%s", currentParser.Url, nexthref)
	// 		return bodyString, nextLink
	// 	}
	// }

	return "nil", "nil"
}

// func extractDomain(link string) (string, error) {
// 	u, err := url.Parse(link)
// 	if err != nil {
// 		return "", err
// 	}
// 	return u.Host, nil
// }

func crawl(url string) {

	_, nextURL := FetchURL(url)
	fmt.Println(nextURL)
	// fmt.Println(nextURL)
	// fmt.Println(url2)

	// chapter finish
	crawl(nextURL)

}

// type ImageFetcher interface {
// 	GetImage() io.Reader
// }

// type MangaReaderFetcher struct {
// 	Html []byte
// }

// func (mr *MangaReaderFetcher) GetImage() io.Reader {
// 	mr.
// }

// src, exist := doc.Find("#img").Attr("src")

// if !exist {
// 	panic(errors.New("Image not found"))
// }

// resp, err = client.Get(src)

// if err != nil {
// 	panic(err)
// }

// func crawl(url) {

// 	var img []byte
// 	resp.Body.Read(img)
// 	defer resp.Body.Close()

// 	f, err := os.OpenFile("test.jpg", os.O_RDWR|os.O_CREATE, 0755)
// 	len, err := io.Copy(f, resp.Body)

// 	if err != nil {
// 		panic(err)
// 	}

// 	defer f.Close()

// 	fmt.Println(len)

// }
