package main

type element struct {
	parent, target string
	attr           map[string]string
}

type config struct {
	url, name   string
	image, next element
}

var configs = map[string]config{
	"mangareader": {
		name: "mangareader",
	},
	"nhentai": {
		name: "nhentai",
	},
	"mangatown": {
		name: "mangatown",
	},
	"hentaicafe": {
		name: "hentaicafe",
	},
	"hentainexus": {
		name: "hentainexus",
	},
	"mangazuki": {
		name: "mangazuki",
	},
	"mangapanda": {
		name: "mangapanda",
	},
	"isekaiscan": {
		name: "isekaiscan",
	},
	"mngdoom": {
		name: "mngdoom",
	},
	"topmanhua": {
		name: "topmanhua",
	},
	"mangahere": {
		name: "mangahere",
	},
	"manytoon": {
		name: "manytoon",
	},
	"manhwa18": {
		name: "manhwa18",
	},
	"manhwaclub": {
		name: "manhwaclub",
	},
}
