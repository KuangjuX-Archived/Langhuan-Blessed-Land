package main

import (
	. "spider/douban"
)

func main() {
	url := "https://movie.douban.com/j/search_subjects?type=movie&tag=%E7%83%AD%E9%97%A8&sort=recommend&page_limit=20&page_start=60"
	content := SendRequest(url)
	Parse(content)
	// fmt.Printf("%v", string(content))
}
