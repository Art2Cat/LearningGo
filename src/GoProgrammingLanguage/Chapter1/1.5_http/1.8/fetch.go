package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

/**
练习 1.8： 修改fetch这个范例，如果输入的url参数没有 http:// 前缀的话，
为这个url加上该前缀。你可能会用到strings.HasPrefix这个函数。
*/
func main() {
	fetchBody()
}

func fetchBody() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		//http.Get函数是创建HTTP请求的函数
		resp, err := http.Get(url)
		//处理错误，有错误则打印出错误信息，并退出
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		//fmt.Print(resp.Status)

		var bs io.Writer
		c, err := io.Copy(bs, resp.Body)
		fmt.Println(c)
		//ioutil.ReadAll函数从response中读取到全部内容；其结果保存在变量b中
		//b, err := ioutil.ReadAll(resp.Body)
		//读取完毕后关闭resp的Body流，防止资源泄露
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		//最后打印出获取到的结果
		fmt.Printf("%s", bs)
	}
}
