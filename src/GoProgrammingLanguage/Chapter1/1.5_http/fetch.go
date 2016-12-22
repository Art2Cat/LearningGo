package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//go build fetch.go
//./fetch http://www.qq.com
func main() {
	for _, url := range os.Args[1:] {
		//http.Get函数是创建HTTP请求的函数
		resp, err := http.Get(url)
		//处理错误，有错误则打印出错误信息，并退出
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		//fmt.Print(resp.Status)

		//ioutil.ReadAll函数从response中读取到全部内容；其结果保存在变量b中
		b, err := ioutil.ReadAll(resp.Body)
		//读取完毕后关闭resp的Body流，防止资源泄露
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		//最后打印出获取到的结果
		fmt.Printf("%s", b)
	}
}
