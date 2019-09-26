package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ProfOak/flag2"
	"github.com/nainglinaung/gomanga/sites/mangareader"
)

var (
	site    string
	manga   string
	chapter string
)

func checkParameter(options flag2.Options) (string, string, string) {

	site = fmt.Sprintf("%v", options["site"])
	manga = fmt.Sprintf("%v", options["manga"])
	chapter = fmt.Sprintf("%v", options["chapter"])

	return site, manga, chapter

}

func init() {

	f := flag2.NewFlag()

	f.AddString("s", "site", "Which mangawebsite do you want to fetch", "mangareader")
	f.AddString("m", "manga", "Which manga do you want to fetch", "bleach")
	f.AddString("c", "chapter", "Which chapter do you want to fetch", "482")

	options, _ := f.Parse(os.Args)

	site, manga, chapter = checkParameter(options)

}

func main() {

	// currentParser = configs[site]

	chapter, err := strconv.Atoi(chapter)

	if err != nil {
		panic(err)
	}

	// Create Manga Directory
	// os.MkdirAll(manga, os.ModePerm)

	if configs[site].name == "mangareader" {
		mangareader.Execute(manga, chapter)
	}

	// fullURL := fmt.Sprintf("%s/%s/%s", currentParser.Url, manga, chapter)
	// fmt.Println(fullURL)
	// crawl(fullURL)

	// mangahere.FetchURL("asdasd")
	// link := fetchNext(currentParser.Next)
	// NextLink := Next.Target

	// fmt.Println(fullURL)

	// parser := parserMap[*link]

	// fmt.Println(parser)
	// fmt.Println(parserMap["http://www.ecchi-manga.net"])

	// var content []byte

	// if counts, err := f.Read(content); err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println(counts)
	// }

	// fmt.Println(string(content))

	// tokenizer := html.Parse(resp.Body)

	// // ioutil.ReadAll(resp.Body)

	// for {
	// 	tt := tokenizer.Next()

	// 	if tt == html.ErrorToken {
	// 		// ...
	// 		continue

	// 	}
	// 	fmt.Println("LL")
	// 	fmt.Println(tokenizer.Token())
	// 	fmt.Println("NN")

	// }

	// fmt.Println(body)
}
