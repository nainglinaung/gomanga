package ehelper

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func (e Ehelper) CreateFolderPath(manga string, chapter int, output string) string {
	var folderPath string
	if output == "" {
		folderPath = fmt.Sprintf("%s/%d", manga, chapter)

	} else {
		folderPath = fmt.Sprintf("%s%s/%d", output, manga, chapter)
	}
	return folderPath
}

func (e Ehelper) CreateFolder(path string) {
	os.MkdirAll(path, os.ModePerm)
}

func (e Ehelper) LowerAndReplace(manga, existing, replacement string) string {
	manga = strings.ToLower(manga)
	manga = strings.Replace(manga, existing, replacement, -1)
	return manga
}

func (e Ehelper) Download(url string, fullImagePath string) {
	resp := e.request(url)
	file, err := os.Create(fullImagePath)
	_, err = io.Copy(file, resp.Body)
	e.CheckError(err)
	defer resp.Body.Close()
}
