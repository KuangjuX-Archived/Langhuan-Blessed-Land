package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	. "spider/douban"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var (
	BaseUrl = "https://movie.douban.com/top250"
)

// 新增数据
func Add(movies []DoubanMovie) {
	for _, movie := range movies {
		data, _ := json.Marshal(movie)
		ioutil.WriteFile("data.json", data, 0644)
	}
}

// 开始爬取
func Start() {
	var movies []DoubanMovie

	pages := GetPages(BaseUrl)
	for _, page := range pages {
		doc, err := goquery.NewDocument(strings.Join([]string{BaseUrl, page.Url}, ""))
		if err != nil {
			log.Println(err)
		}

		movies = append(movies, ParseMovies(doc)...)
	}
	fmt.Printf("movies: %v\n", movies)
	Add(movies)
}

func main() {
	Start()
}
