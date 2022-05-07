package news

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

// GetNews 获取百度新闻主页新闻信息
func GetNews() map[string]string{
	url := "http://news.baidu.com/"
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("获取百度新闻首页失败！")
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		fmt.Printf("status code error: %d", res.StatusCode)
	}
	// 使用goquery解析html
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("goquery load failed!")
	}

	// 获取热点新闻的标题以及对应的link
	// todo 暂不考虑新闻名相同的情况
	news := make(map[string]string)
	// class hotnews 下对应的新闻
	doc.Find(".hotnews").Each(func(i int, selection *goquery.Selection) {
		selection.Find("a").Each(func(i int, s *goquery.Selection) {
			title := s.Text()
			//fmt.Println(title)
			link, _ := s.Attr("href")
			//fmt.Println(link)
			news[title] = link
		})
	})
	// class ulist & focuslistnews 下对应的新闻
	doc.Find(".ulist.focuslistnews").Each(func(i int, selection *goquery.Selection) {
		selection.Find("a").Each(func(i int, s *goquery.Selection) {
			title := s.Text()
			//fmt.Println(title)
			link, _ := s.Attr("href")
			//fmt.Println(link)
			news[title] = link
		})
	})

	return news
}

