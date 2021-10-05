package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "https://movie.douban.com/j/search_subjects?type=movie&tag=%25E7%2583%25AD%25E9%2597%25A8&sort=rank&page_limit=20&page_start=40"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("User-Agent", "PostmanRuntime/7.26.10")
	req.Header.Add("Cookie", "Cookie_1=value; bid=AoIf_jbuNcc")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print("print.\n")
	fmt.Printf("content: %v\n", string(body))
	// fmt.Println(string(body))
}
