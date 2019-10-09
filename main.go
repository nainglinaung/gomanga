package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ProfOak/flag2"
	"github.com/nainglinaung/gomanga/sites/hentaicafe"
	"github.com/nainglinaung/gomanga/sites/hentainexus"
	"github.com/nainglinaung/gomanga/sites/mangareader"
	"github.com/nainglinaung/gomanga/sites/mangatown"
	"github.com/nainglinaung/gomanga/sites/nhentai"
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
	}

}
