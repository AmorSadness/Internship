package test

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"testing"
)

func Test_goquery(t *testing.T) {


	//goQuery 解析html
	//获取http返回结果
	url := "http://news.baidu.com/"
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("获取网页失败！")
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		fmt.Printf("status code error: %d", res.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("goQuery load fail")
	}

	//find item
	//打印获取的全部html信息
	fmt.Println(doc.Html())


	//判断元素是否存在
	exist := doc.Find("div").HasClass("hotnews")
	if exist {
		fmt.Println("存在该元素")
	} else {
		fmt.Println("不存在该元素")
	}
	fmt.Println(doc.Find(".hotnews").Text())
	doc.Find(".hotnews").Each(func(i int, selection *goquery.Selection) {
		title := selection.Find("a").Text()
		link := selection.Find("a")
		fmt.Println(link)
		fmt.Printf("Review %d: %s\n", i, title)
	})

	doc.Find("a")

	doc.Find("a").Each(func(i int, selection *goquery.Selection) {
		//title := selection.Find("a").Text()
		//link := selection.Find("a")
		//fmt.Println(link)
		//fmt.Printf("Review %d: %s\n", i, title)
		fmt.Println(selection.Attr("href"))
	})

	//goquery 基本使用
	//元素选择器 a / p / div等html基本元素 Find("div")
	//id选择器 #紧跟id值 Find("#id值")
	//元素 和 id 可以组合  Find("div#id值")    貌似可以理解为元素具有属性
	//class选择器 Find(".class")
	//元素 和 class 可以组合 Find("div.class")
	//属性(值)选择器 Find("div[class=name]")
	//多属性筛选 Find("div[id][lang=mkm]")
	//子元素筛选 Find("parent>child") 筛选父元素下符合条件的一级子元素   Find("body>div")
	//全部子元素，中间用空格 Find("body div")
	//prev + next 相邻元素选择器（严格）
	//prev ~ next （同一父元素）
	//内容过滤器
	//
	//
	//获取热点新闻的标题以及对应的link

	doc.Find(".hotnews").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Html())
		//fmt.Println(selection.Find("a").Attr("href"))
	})

	hotnews := doc.Find(".hotnews")

	hotnews.Find("a").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
		link, _ := selection.Attr("href")
		fmt.Println(link)
	})


	doc.Find(".ulist.focuslistnews").Each(func(i int, selection *goquery.Selection) {

		selection.Find("a").Each(func(i int, s *goquery.Selection) {
			fmt.Println(s.Text())
			link, _ := s.Attr("href")
			fmt.Println(link)
		})
	})
}
