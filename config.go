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
}
