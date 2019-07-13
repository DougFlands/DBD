package main

import (
	"fetchDbd/fetch"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json

	key := ""
	value := ""
	resp := ""

	for k, v := range r.Form {
		key = k
		value = strings.Join(v, "")
	}


	if key == "page" {
		result, err := fetch.Fetchurl(value)
		if err != nil {
			fmt.Println(err)
		}
		str := string(result)
		resp = str[10 : len(str)-2]
	}

	fmt.Fprintf(w, resp) //这个写入到w的是输出到客户端的
}

func main() {
	http.HandleFunc("/", sayhelloName) //设置访问的路由
	err := http.ListenAndServe(":3000", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
