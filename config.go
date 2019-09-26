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
		url: "http://www.mangareader.net", name: "mangareader",
		image: element{
			parent: "div", target: "img#img",
			attr: map[string]string{"id": "img"},
		},
		next: element{
			parent: "div", target: "span.next > a",
			attr: map[string]string{"id": "imgholder"},
		},
	},
}
