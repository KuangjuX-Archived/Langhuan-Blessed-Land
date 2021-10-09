package douban

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func SendRequest(url string) []byte {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "PostmanRuntime/7.26.10")
	req.Header.Add("Cookie", "Cookie_1=value; bid=AoIf_jbuNcc")

	res, err := client.Do(req)
	if res != nil {
		fmt.Printf("[Error] %v", err)
	}
	defer res.Body.Close()

	content, _ := ioutil.ReadAll(res.Body)
	return content
}

func Parse(data []byte) {
	var t interface{}
	json.Unmarshal(data, &t)
	content := t.(map[string]interface{})
	// fmt.Printf("[Content] %v\n", content)
	Subject := content["subjects"]
	// fmt.Printf("[subjects] %v\n", Subject)
	SubjectList := Subject.([]interface{})
	for i := 0; i < len(SubjectList); i++ {
		res := SubjectList[i].(map[string]interface{})
		fmt.Printf("[Debug] title: %v  rate: %v\n", res["title"].(string), res["rate"].(string))
		// fmt.Printf("%v: %v\n", i, SubjectList[i].(map[string]interface{}))
	}
}
