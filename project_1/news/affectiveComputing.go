package news

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)


type res struct {
	Result float64 `json:"result"`
}


// 情感计算结果，积极 & 消极 分别保存在一个map[string]string中 存放顺序：先积极，后消极
func AffectiveComputingResult(m map[string]string) (map[string]string, map[string]string){
	pos := make(map[string]string)
	neg := make(map[string]string)
	for title := range m {
		score := callComputingApi(title)
		if score > 0.5 { // 积极标题
			pos[title] = m[title]
		}
		if score < -0.5 { // 消极标题
			neg[title] = m[title]
		}
	}
	return pos, neg
}


// 调用情感计算api，获取新闻标题得分
func callComputingApi(words string) float64{
	//fmt.Println(words)
	client := &http.Client{Timeout: 0}
	url := "http://baobianapi.pullword.com:9091/get.php"
	contentType := "application/json"
	var body []byte
	body, err := json.Marshal(words)
	if err != nil {
		fmt.Println("marshal failed!")
	}
	resp, err := client.Post(url, contentType, bytes.NewReader(body))
	if err != nil {
		fmt.Println("api call failed!")
	}
	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)

	//fmt.Println(string(content))

	r := res{}

	jsonErr := json.Unmarshal(content, &r)
	if jsonErr != nil {
		fmt.Println("json unmarshal failed!")
	}
	//fmt.Println(r.Result)
	return r.Result
}
