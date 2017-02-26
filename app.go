package main

import (
	"github.com/zhangweilun/gor"
	"github.com/zhangweilun/goavi/common"
	"log"
	"github.com/zhangweilun/goxmlpath"
	"strings"
	"fmt"
	"strconv"
)

/**
* 
* @author willian
* @created 2017-02-26 13:23
* @email 18702515157@163.com  
**/
func main() {

	ro := &gor.Request_options{UserAgent:common.RandomUserAgent()}

	res, err := gor.Get("http://www.icpmp.com/fanhao/daquanmp1v.html", ro)

	if err != nil {
		log.Fatalln(err)
	}

	page, err := goxmlpath.ParseHTML(res.RawResponse.Body)
	if err != nil {
		log.Fatal(err)
	}
	res.RawResponse.Body.Close()
	total_xpath := goxmlpath.MustCompile("/html/body/div[5]/div/span")

	total_unchecked, ok := total_xpath.String(page)

	if !ok {
		log.Fatalln("xpath 解析出错")
	}

	start := strings.Index(total_unchecked, "/")
	total := strings.TrimSpace(total_unchecked[start+1:])

	fmt.Println(total)

	total_page, err := strconv.ParseInt(total, 10, 0)
	if err != nil {
		log.Fatalln(err)
	}
	//fmt.Println(total_page)


	urls := make([]string, total_page)

	for i:=1;i <= int(total_page); i++ {
		urls[i-1] = "http://www.icpmp.com/fanhao/daquanmp"+strconv.Itoa(i)+"v.html"
	}

	for _,v := range urls {
		ro := &gor.Request_options{UserAgent:common.RandomUserAgent()}
		res, err := gor.Get(v, ro)
		//fmt.Println(res.String())
		if err != nil {
			fmt.Println("get请求出错")
			continue
		}
		page, err := goxmlpath.ParseHTML(res.RawResponse.Body)
		if err != nil {
			log.Fatal(err)
		}
		res.RawResponse.Body.Close()
		titles := goxmlpath.MustCompile(`/html/body/div[6]/div/div/ul/notempty/li/h4/a/text()`)
		names := goxmlpath.MustCompile(`/html/body/div[6]/div/div/ul/notempty/li/p/a/text()`)
		images := goxmlpath.MustCompile(`/html/body/div[6]/div/div/ul/notempty/li/a/span/img/@lz_src`)
		title_array := common.PutData(titles, page)
		name_array := common.PutData(names, page)
		image_array := common.PutData(images, page)
		iter := common.Zip(title_array, name_array, image_array)
		for tuple := iter(); tuple != nil; tuple = iter() {
			fmt.Println(tuple[0])
			fmt.Println(tuple[1])
			fmt.Println(tuple[2])
		}

	}


}
