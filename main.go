package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ProfOak/flag2"
	"github.com/nainglinaung/gomanga/sites/hentaicafe"
	"github.com/nainglinaung/gomanga/sites/hentainexus"
	"github.com/nainglinaung/gomanga/sites/isekaiscan"
	"github.com/nainglinaung/gomanga/sites/mangahere"
	"github.com/nainglinaung/gomanga/sites/mangapanda"
	"github.com/nainglinaung/gomanga/sites/mangareader"
	"github.com/nainglinaung/gomanga/sites/mangatown"
	"github.com/nainglinaung/gomanga/sites/mangazuki"
	"github.com/nainglinaung/gomanga/sites/mngdoom"
	"github.com/nainglinaung/gomanga/sites/nhentai"
	"github.com/nainglinaung/gomanga/sites/topmanhua"
)

var (
	site    string
	manga   string
	chapter string
	output  string
)

func checkParameter(options flag2.Options) (string, string, string, string) {

	output = fmt.Sprintf("%v", options["output"])
	site = fmt.Sprintf("%v", options["site"])
	manga = fmt.Sprintf("%v", options["manga"])
	chapter = fmt.Sprintf("%v", options["chapter"])

	return output, site, manga, chapter

}

func init() {

	f := flag2.NewFlag()
	fmt.Println("GoManga Starting:")
	f.AddString("s", "site", "Which mangawebsite do you want to fetch", "mangareader")
	f.AddString("m", "manga", "Which manga do you want to fetch", "bleach")
	f.AddString("c", "chapter", "Which chapter do you want to fetch", "482")
	f.AddString("o", "output", "where do you want to save", "")
	options, _ := f.Parse(os.Args)

	output, site, manga, chapter = checkParameter(options)
}

func main() {

	chapter, err := strconv.Atoi(chapter)

	if err != nil {
		panic(err)
	}

	switch name := configs[site].name; name {
	case "mangareader":
		mangareader.Execute(manga, chapter, output)
	case "nhentai":
		nhentai.Execute(chapter, output)
	case "mangatown":
		mangatown.Execute(manga, chapter, output)
	case "hentaicafe":
		hentaicafe.Execute(chapter, output)
	case "hentainexus":
		hentainexus.Execute(chapter, output)
	case "mangazuki":
		mangazuki.Execute(manga, chapter, output)
	case "mangapanda":
		mangapanda.Execute(manga, chapter, output)
	case "isekaiscan":
		isekaiscan.Execute(manga, chapter, output)
	case "mngdoom":
		mngdoom.Execute(manga, chapter, output)
	case "topmanhua":
		topmanhua.Execute(manga, chapter, output)
	case "mangahere":
		mangahere.Execute(manga, chapter, output)
	}

}
