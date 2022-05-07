package server

import (
	"encoding/json"
	"fmt"
	"internship/project_1/news"
	"io"
	"net/http"
)

type News struct {
	Title string
	Link string
}

type Data struct {
	ResultNews []News
}

func GetPositiveNews(writer http.ResponseWriter, requst *http.Request) {
	baiduNews := news.GetNews()
	pos, _ := news.AffectiveComputingResult(baiduNews)
	data := new(Data)
	for news := range pos {
		n := News{
			Title: news,
			Link: pos[news],
		}
		data.ResultNews = append(data.ResultNews, n)
	}

	res, _ := json.Marshal(data)
	io.WriteString(writer, string(res))
}

func GetNegativeNews(writer http.ResponseWriter, requst *http.Request) {
	baiduNews := news.GetNews()
	_, neg := news.AffectiveComputingResult(baiduNews)
	data := new(Data)
	for news := range neg {
		n := News{
			Title: news,
			Link: neg[news],
		}
		data.ResultNews = append(data.ResultNews, n)
	}

	res, _ := json.Marshal(data)
	io.WriteString(writer, string(res))
}

func ServerStart() {
	// rout
	http.HandleFunc("/news/positive", GetPositiveNews)
	http.HandleFunc("/news/negative", GetNegativeNews)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("web服务启动异常!")
	} else {
		fmt.Println("web服务启动成功!")
	}
}
