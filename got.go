package main

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"net/http"
	"os"
)

func getTranslate(text string) {

	//有道翻译api
	APIKEY := "1712456283"
	KEYFROM := "gaodb712"
	APIURL := "http://fanyi.youdao.com/openapi.do"

	// 命令行颜色
	red := "\033[31;1m %s \033[1;m"
	green := "\033[32;1m %s \033[1;m"
	blue := "\033[34;1m %s \033[1;m"
	// purple := "\033[35;1m %s \033[1;m"
	yellow := "\033[33;1m %s \033[1;m"
	dark_green := "\033[36;1m %s \033[1;m"

	myurl := APIURL + "?keyfrom=" + KEYFROM + "&key=" + APIKEY + "&type=data&doctype=json&version=1.1&q=" + text
	response, _ := http.Get(myurl)
	robots, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("error")
	}

	response.Body.Close()
	js, err := simplejson.NewJson([]byte(string(robots)))
	if err != nil {
		panic("json format error")
	}
	if _, ok := js.CheckGet("basic"); !ok {
		fmt.Printf(red, "未查询到该词！！！")
	} else {
		fmt.Printf(dark_green, "单词：\n\t")
		word, _ := js.Get("query").String()
		fmt.Printf(green, word)

		fmt.Println()
		fmt.Printf(dark_green, "翻译：\n\t")
		translation, _ := js.Get("translation").Array()
		for _, v := range translation {
			fmt.Printf(blue, v)
		}

		fmt.Println()
		fmt.Printf(dark_green, "详细：\n\t")
		explains, _ := js.Get("basic").Get("explains").Array()
		for _, v := range explains {
			fmt.Printf(blue, v)
		}

		fmt.Println()
		fmt.Printf(dark_green, "网络短语：\n")
		webExplains, _ := js.Get("web").Array()

		for _, v := range webExplains {
			if vv, ok := v.(map[string]interface{}); ok {
				fmt.Printf("\t"+yellow, vv["key"])
				fmt.Println(vv["value"])

			}

		}

	}

}

func main() {

	getTranslate(os.Args[1])
}
