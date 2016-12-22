package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

/**
练习 1.7： 函数调用io.Copy(dst, src)会从src中读取内容，并将读到的结果写入到dst中，
使用这个函数替代掉例子中的ioutil.ReadAll来拷贝响应结构体到os.Stdout，
避免申请一个缓冲区（例子中的b）来存储。记得处理io.Copy返回结果中的错误。
*/
func main() {
	fetchBody()
}

func fetchBody() {
	for _, url := range os.Args[1:] {
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
