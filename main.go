package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ProfOak/flag2"
	"github.com/nainglinaung/gomanga/sites/mangareader"
	"github.com/nainglinaung/gomanga/sites/nhentai"
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

	chapter, err := strconv.Atoi(chapter)

	if err != nil {
		panic(err)
	}
	// greet.test()
	// Create Manga Directory
	// os.MkdirAll(manga, os.ModePerm)
	if configs[site].name == "mangareader" {
		mangareader.Execute(manga, chapter)
	} else if configs[site].name == "nhentai" {
		nhentai.Execute(chapter)
	}

}
