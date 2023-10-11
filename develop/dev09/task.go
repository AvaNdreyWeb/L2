package main

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"io"
	"log"
	"mime"
	"net/http"
	"net/url"
	"os"
	"path"
	"regexp"
	"strings"
	"sync"
)

func main() {
	uri := os.Args[1]
	dst := "./download/"
	maxDepth := 1

	download(uri, dst, maxDepth)
}

func download(uri, dst string, maxDepth int) {
	if maxDepth < 0 {
		return
	}

	if _, e := url.ParseRequestURI(uri); e != nil {
		return
	}

	log.Println("Download: ", uri)

	res, err := http.Get(uri)
	if err != nil {
		log.Println(err)
		return
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return
	}

	filename := getFilename(uri, res.Header)
	out, err := os.Create(dst + filename)
	if err != nil {
		log.Println(err)
		return
	}
	defer out.Close()
	out.Write(data)

	links := getParseLinks(data)
	var wg sync.WaitGroup
	for _, link := range links {
		wg.Add(1)
		go func(link string) {
			defer wg.Done()
			download(link, dst, maxDepth-1)
		}(link)
	}
	wg.Wait()
}

func getFilename(uri string, header http.Header) string {
	mediaType := createFilename(header)
	baseName := path.Base(uri)

	if mediaType == "" {
		return "Error mediaType"
	}

	newName := cut(cut(baseName, "#", 0), "?", 0)

	if path.Ext(newName) == "" && mediaType != "" {
		return newName + "." + mediaType
	}

	return newName
}

func createFilename(header http.Header) string {
	contentType := header.Get("Content-Type")
	mType, _, _ := mime.ParseMediaType(contentType)
	mediaType := cut(mType, "/", 1)
	return mediaType
}

func cut(s, sep string, i int) string {
	if strings.Contains(s, sep) {
		return strings.Split(s, sep)[i]
	}
	return s
}

func getParseLinks(data []byte) []string {
	reg := regexp.MustCompile(`(http|https):\/\/([\w\-_]+(?:(?:\.[\w\-_]+)+))([\w\-\.,@?^=%&amp;:/~\+#]*[\w\-\@?^=%&amp;/~\+#])?`)
	result := reg.FindAll(data, -1)

	subUrls := make([]string, len(result))
	for i := 0; i < len(result); i++ {
		subUrls = append(subUrls, string(result[i]))
	}

	return subUrls
}
